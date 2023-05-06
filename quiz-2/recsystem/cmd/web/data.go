// Filename: cmd/web/data.go
package main

import (
	"github.com/MejiaFrancis/3161/3162/quiz-2/recsystem/internal/models"
)

type templateData struct {
	Reservation *models.Reservation
	User 		*models.User
	Flash 	 	string
	CSRFToken 	string
}
