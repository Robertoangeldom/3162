// Filename: internal/models/reservation.go
package models

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//Create a user
var (
	ErrNoRecord = errors.New("no matching record found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateEmail = errors.New("duplicate email")
)

type User struct{
	ID int64
	Email string
	FirstName string
	LastName string
	Age int
	Address string
	Phone string
	Roles int
	Password []byte
	Activated bool
	CreatedAt time.Time
}

type UserModel struct{
	DB *sql.DB
}


func (m *UserModel) Insert(email, fname, lname , age, address, phone, password string) error {
	var int_age int
	//lets first hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	int_age, err = strconv.Atoi(age)
	if err != nil{
		return err
	}
	query := `
	INSERT INTO users (email, first_name, last_name, age, user_address, phone_number, user_password)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = m.DB.ExecContext(ctx, query, email, fname, lname, int_age, address, phone, hashedPassword)
	if err != nil {
		switch {
		case err.Error() == `pgx: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, int, error) {
	var id int
	var roles_id int
	var hashedPassword []byte

	query := `
		SELECT users_id, user_password, roles_id 
		FROM users
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(&id, &hashedPassword, &roles_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, 0, ErrInvalidCredentials
		} else {
			return 0, 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, 0, ErrInvalidCredentials
		} else {
			return 0, 0, err
		}
	}

	return id, roles_id, nil
}