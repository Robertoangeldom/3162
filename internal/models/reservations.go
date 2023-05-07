// Filename: internal/models/reservation.go
// this file is used to show the fields of a reservation
package models

import (
	"context"
	"database/sql"
	"time"
)

// The Question model will represent a single question in our questions table
type Reservation struct {
	ReservationID   int64
	UserID          int64
	ReservationDate time.Time // date needs to be extracted
	Duration        int8      // expected time in minutes
	PeopleCount     int8      // number of people for the session
	Notes           string
	Approval        bool
	CreatedAt       time.Time
}

// The QuestionModel type will encapsulate the
// DB connection pool that will be initialized
// in the main() function
type ReservationModel struct {
	DB *sql.DB
}

// The Insert() function stores a question into the  table
func (m *ReservationModel) Insert(duration int, peopleCount int, notes string) (int64, error) {
	// id will be used to stored the unique identifier returned by
	// PostgreSQL after adding the row to the table
	var id int64
	statement :=
		`
							INSERT INTO reservations(duration, peopleCount, notes )
							VALUES($1, $2, $3)
							RETURNING reservation_id
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, duration, peopleCount, notes).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *ReservationModel) Get() (*Reservation, error) {
	var res Reservation

	statement :=
		`
							SELECT reservations_id, user_id, people_count
							FROM reservations
							ORDER BY RANDOM()
							LIMIT 1
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&res.ReservationID, &res.UserID, &res.Duration, &res.PeopleCount, &res.Notes, &res.Approval, &res.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
