package booking

type MemoryStore struct {
	bookings map[string]Booking // SeatId -> Booking Struct
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	if _, exsists := s.bookings[b.SeatID]; exsists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b
	return nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	var movieBooking []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			movieBooking = append(movieBooking, b)
		}
	}

	return movieBooking
}
