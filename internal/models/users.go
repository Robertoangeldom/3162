// Filename: internal/models/reservation.go
package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Create a user
var (
	ErrNoRecord           = errors.New("no matching record found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateEmail     = errors.New("duplicate email")
)

type User struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	Age       int
	Address   string
	Phone     string
	Roles     int
	Password  []byte
	Activated bool
	CreatedAt time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(email, fname, lname, age, address, phone, password string) error {
	var int_age int
	//lets first hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	int_age, err = strconv.Atoi(age)
	if err != nil {
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

// The Display() function retrieves all users from the database
func (m *UserModel) Display() ([]User, error) {
	query := `
	Select users_id, first_name, last_name, phone_number, activated
    From users limit 3;
	`
	rows, err := m.DB.Query(query)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	var userTypes []User

	// Iterate over the rows and create a slice of structs
	for rows.Next() {
		var userType User
		err := rows.Scan(&userType.ID, &userType.FirstName, &userType.LastName, &userType.Phone, &userType.Password)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		userTypes = append(userTypes, userType)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	return userTypes, nil
}

// Updates the individual user in the row

func (m *UserModel) Update(id int64, email string, fname string, lname string, age int, address string, phone string, roles int, password []byte, activated bool) error {
	//var int_age int
	//lets first hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	// int_age, err = strconv.Atoi(age)
	// if err != nil {
	// 	return err
	// }

	query := `UPDATE users SET email=$1, first_name=$2, last_name=$3, age=$4, address=$5, phone=$6, roles=$7, password=$8, activated=$9 WHERE id=$10`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = m.DB.ExecContext(ctx, query, email, fname, lname, age, address, phone, roles, hashedPassword, activated)
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
func (u *UserModel) GetByID(id int64) (*User, error) {
	query := "SELECT id, email, firstname, lastname, age, address, phone, roles, password, activated FROM users WHERE id = $1"

	user := &User{}
	err := u.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Age, &user.Address, &user.Phone, &user.Roles, &user.Password, &user.Activated)
	if err != nil {
		fmt.Println("Error scanning row:", err)
		return nil, err
	}

	return user, nil
}
