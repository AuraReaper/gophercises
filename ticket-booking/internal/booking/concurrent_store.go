package booking

import "sync"

type ConcurrentStore struct {
	bookings map[string]Booking // SeatId -> Booking Struct
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurrentStore) Book(b Booking) error {
	s.Lock()
	defer s.Unlock()
	if _, exsists := s.bookings[b.SeatID]; exsists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b
	return nil
}

func (s *ConcurrentStore) ListBookings(movieID string) []Booking {
	var movieBooking []Booking
	s.Lock()
	defer s.Unlock()
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			movieBooking = append(movieBooking, b)
		}
	}

	return movieBooking
}
