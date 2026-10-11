package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"uuid"

	"github.com/redis/go-redis/v9"
)

const defaultHoldTTL = 1 * time.Minute

// K -> V
// seat:{movieID}:{seatID} -> session JSON (TTL = held, no TTL = confirmed)
// session:{sessionID} -> seat key (reverse loookup)
type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{
		rdb: rdb,
	}
}

// builds the reverse lookup key
func sessionKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

func (s *RedisStore) Book(b Booking) error {
	err := s.hold(b)
	if err != nil {
		return err
	}

	return nil
}

func (s *RedisStore) ListBookings(movieID string) []Booking {
	return nil
}

func (s *RedisStore) hold(b Booking) error {
	id := uuid.New().String()
	key := fmt.Sprintf("seat:%s:%s", b.MovieID, b.SeatID)

	b.ID = id
	val, _ := json.Marshal(b)

	resp := s.rdb.SetArgs(context.Background(), key, val, redis.SetArgs{
		Mode: "NX",
		TTL:  defaultHoldTTL,
	})
	if resp.Val() != "OK" {
		return ErrSeatAlreadyBooked
	}

	s.rdb.Set(context.Background(), sessionKey(id), key, defaultHoldTTL)

	// printing booked tiket details
	b.ID = id
	data, _ := json.Marshal(b)
	log.Printf("booked ticket details, %s", data)

	return nil
}
