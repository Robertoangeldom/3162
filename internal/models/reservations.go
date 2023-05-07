// Filename: internal/models/reservation.go
// this file is used to show the fields of a reservation
package models

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

var (
	ErrResreved = errors.New("this date has already been taken")
	ErrInvalid = errors.New("something went wrong while creating reservation")
)

// The Question model will represent a single question in our questions table
type Reservation struct {
	ReservationID   int64
	UserID          int64
	ReservationDate string// date needs to be extracted
	ReservationTime string
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
func (m *ReservationModel) Insert(date, tm, duration, count, notes string) error {
    var int_count int
    int_count, err := strconv.Atoi(count)
    if err != nil{
        return err
    }
    query := `
    INSERT INTO reservations (reservation_date, reservation_time, duration, people_count, notes)
    VALUES ($1, $2, $3, $4, $5)
    `
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    _, err = m.DB.ExecContext(ctx, query, date, tm, duration, int_count, notes)
    if err != nil {
        switch {
		case err.Error() == `pgx: duplicate key value violates unique constraint "reservations_date_time_key"`:
			return ErrResreved
		default:
			return err
		}
    }
    return nil
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
