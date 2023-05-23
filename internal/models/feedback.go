// Filename: internal/models/equipments.go
package models

import (
	"database/sql"
	"fmt"
)

// The EquipmentType model will represent a single equipment type in our equipment_types table
type Feedback struct {
	FeedbackID     string
	FirstName      string
	LastName       string
	FeedbackMessage       string
}

// The EquipmentModel type will encapsulate the
// DB connection pool that will be initialized
// in the main() function
type FeedbackModel struct {
	DB *sql.DB
}

// The Display() function retrieves all equipment types from the database
func (m *FeedbackModel) Display() ([]Feedback, error) {
	query := `
	SELECT feedback.feedback_id, users.first_name, users.last_name, feedback.feedback_message
	FROM feedback
	INNER JOIN users ON feedback.users_id = users.users_id;

	`
	rows, err := m.DB.Query(query)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	var feedbacks []Feedback

	// Iterate over the rows and create a slice of structs
	for rows.Next() {
		var feedback Feedback
		err := rows.Scan(&feedback.FeedbackID, &feedback.FirstName, &feedback.LastName, &feedback.FeedbackMessage)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		feedbacks = append(feedbacks, feedback)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	return feedbacks, nil
}
