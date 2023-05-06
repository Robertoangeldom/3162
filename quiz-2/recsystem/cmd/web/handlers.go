// handlers.go
package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/MejiaFrancis/3161/3162/quiz-2/recsystem/internal/models"
	"github.com/justinas/nosurf"
	//"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl", nil)

}
func (app *application) loginform(w http.ResponseWriter, r *http.Request) {flash := app.sessionManager.PopString(r.Context(), "flash")
//render
data := &templateData{ //putting flash into template data
	Flash: flash,
	CSRFToken: nosurf.Token(r),
}
RenderTemplate(w, "login.page.tmpl", data)
}

//loginformSubmit

func (app *application) loginformSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	//lets write the data to the table
	id, roles_id, err := app.user.Authenticate(email, password)
	log.Println(id, roles_id ,err)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			RenderTemplate(w, "login.page.tmpl", nil)
		}
		return
	}
	
	//add the user to the session cookie
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		return
	}
	// add an authenticate entry
	if roles_id == 2{
	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
	http.Redirect(w, r, "/user", http.StatusSeeOther)
	}else if roles_id == 1{
	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "register.page.tmpl", data)
}

//registerSubmit

func (app *application) registerSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fname := r.PostForm.Get("firstname") //"name" is the name of the form
	lname := r.PostForm.Get("lastname")
	age := r.PostForm.Get("age")
	phone := r.PostForm.Get("phone")
	address := r.PostForm.Get("address")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	//lets write the data to the table
	err := app.user.Insert(fname, lname , age, address, phone, email, password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			RenderTemplate(w, "signup.page.tmpl", nil)
		}
	}
		app.sessionManager.Put(r.Context(), "flash", "SignupWassuccessful")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) reserve(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "reservation.page.tmpl", nil)

}

//reserveFormSubmit

func (app *application) reserveFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "reservation.page.tmpl", nil)

}

func (app *application) userPortal(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "userPortal.page.tmpl", nil)

}

//userPortalFormSubmit
func (app *application) userPortalFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "userPortal.page.tmpl", nil)

}

func (app *application) adminPortal(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "adminPortal.page.tmpl", nil)

}

func (app *application) adminPortalFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "adminPortal.page.tmpl", nil)

}

func (app *application) feedback(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "feedback.page.tmpl", nil)
}

//feedbackFormSubmit
func (app *application) feedbackFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "feedback.page.tmpl", nil)

}

// func (app *application) pollReplyShow(w http.ResponseWriter, r *http.Request) {
// 	reservation, err := app.questions.Get()
// 	if err != nil {
// 		return
// 	}
// 	data := &templateData{
// 		Reservation: reservation,
// 	}
// 	RenderTemplate(w, "poll.page.tmpl", data)
// }

// func (app *application) submitlogin(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "bad request", http.StatusBadRequest)
// 		return
// 	}
// 	response := r.PostForm.Get("emotion")
// 	question_id := r.PostForm.Get("id")
// 	identifier, err := strconv.ParseInt(question_id, 10, 64)
// 	if err != nil {
// 		return
// 	}
// 	_, err = app.responses.Insert(response, identifier)
// 	if err != nil {
// 		http.Error(w,
// 			http.StatusText(http.StatusInternalServerError),
// 			http.StatusInternalServerError)
// 		return
// 	}
// }

// This code stores the different options in a variabe then stores using the Insert function... for radio button.
// func (app *application) optionsCreateSubmit(w http.ResponseWriter, r *http.Request) {
// 	// get the four options
// 	r.ParseForm()
// 	option_1 := r.PostForm.Get("option_1")
// 	option_2 := r.PostForm.Get("option_2")

// 	// save the options
// 	_, err := app.reservations.Insert(option_1, option_2,)
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	}
// }
