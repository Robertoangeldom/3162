// handlers.go
package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

func (app *application) loginform(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	// render
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "login.page.tmpl", data)
}

// loginformSubmit
func (app *application) loginformSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	//lets write the data to the table
	id, roles_id, err := app.user.Authenticate(email, password)
	log.Println(id, roles_id, err)
	log.Println(id, roles_id, err)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		return
	}

	//add the user to the session cookie
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		return
	}
	// add an authenticate entry
	if roles_id == 2 {
		app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	} else if roles_id == 1 {
		app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		if roles_id == 2 {
			app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
			http.Redirect(w, r, "/user", http.StatusSeeOther)
		} else if roles_id == 1 {
			app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}
	}
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "register.page.tmpl", data)
}

// registerSubmit
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
	err := app.user.Insert(email, fname, lname, age, address, phone, password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			RenderTemplate(w, "signup.page.tmpl", nil)
		}
	}
	app.sessionManager.Put(r.Context(), "flash", "SignupWassuccessful")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	app.sessionManager.Put(r.Context(), "flash", "SignupWassuccessful")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// func (app *application) reserve(w http.ResponseWriter, r *http.Request) {
// 	RenderTemplate(w, "reservation.page.tmpl", nil)
// }

// // reserveFormSubmit
//
//	func (app *application) reserveFormSubmit(w http.ResponseWriter, r *http.Request) {
//		RenderTemplate(w, "reservation.page.tmpl", nil)
//	}
func (app *application) userPortal(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "userPortal.page.tmpl", data)

}

func (app *application) equipment(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/equipments.page.tmpl", "./ui/html/base.layout.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Println(ts)
	equipmentTypes, err := app.equipments.Display()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Println(equipmentTypes)
	err = ts.Execute(w, equipmentTypes)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// userPortalFormSubmit
// userPortalFormSubmit
func (app *application) userPortalFormSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	date := r.PostForm.Get("date") //"name" is the name of the form
	time := r.PostForm.Get("time")
	duration := r.PostForm.Get("duration")
	count := r.PostForm.Get("count")
	notes := r.PostForm.Get("notes")
	//lets write the data to the table
	err := app.reservations.Insert(date, time, duration, count, notes)
	if err != nil {
		if errors.Is(err, models.ErrInvalid) {
			RenderTemplate(w, "reservation.page.tmpl", nil)
		}
	}
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

// feedbackFormSubmit
// feedbackFormSubmit
func (app *application) feedbackFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "feedback.page.tmpl", nil)

}

//Display all users on Admin dashboard

func (app *application) displayUsers(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("./ui/html/view.users.tmpl", "./ui/html/base.layout.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Println(ts)
	userTypes, err := app.user.Display()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Println(userTypes)
	err = ts.Execute(w, userTypes)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// func (app *application) updateRecord(w http.ResponseWriter, r *http.Request) {
// 	// Get the ID of the record to update from the request URL parameter

// 	ts, err := template.ParseFiles("update.user.page.tmpl", "./ui/html/base.layout.tmpl")
// 	//RenderTemplate(w, "", nil)
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	// start here
// 	r.ParseForm()
// 	fname := r.PostForm.Get("firstname") //"name" is the name of the form
// 	lname := r.PostForm.Get("lastname")
// 	age := r.PostForm.Get("age")
// 	phone := r.PostForm.Get("phone")
// 	address := r.PostForm.Get("address")
// 	email := r.PostForm.Get("email")
// 	password := r.PostForm.Get("password")
// 	//lets write the data to the table
// 	userTypes, err  := app.user.Update(email, fname, lname, age, address, phone,roles, password, actvated)
// 	if err != nil {
// 		if errors.Is(err, models.ErrDuplicateEmail) {
// 			RenderTemplate(w, "signup.page.tmpl", nil)
// 		}
// 	}
// 	app.sessionManager.Put(r.Context(), "flash", "SignupWassuccessful")
// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
// 	app.sessionManager.Put(r.Context(), "flash", "SignupWassuccessful")
// 	http.Redirect(w, r, "/login", http.StatusSeeOther)

// 	//end here
// 	log.Println(ts)

// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	log.Println(userTypes)
// 	err = ts.Execute(w, userTypes)
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// }
func (app *application) updateRecord(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the record to update from the request URL parameter

	id := r.FormValue("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	RenderTemplate(w, "update.user.page.tmpl", nil)

	ts, err := template.ParseFiles("update.user.page.tmpl", "./ui/html/base.layout.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Update the record in the database
	if r.Method == "POST" {
		r.ParseForm()
		fname := r.PostForm.Get("firstname")
		lname := r.PostForm.Get("lastname")
		age, err := strconv.Atoi(r.PostForm.Get("age"))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		phone := r.PostForm.Get("phone")
		address := r.PostForm.Get("address")
		email := r.PostForm.Get("email")
		password := []byte(r.PostForm.Get("password"))
		activated, err := strconv.ParseBool(r.PostForm.Get("activated"))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		roles, err := strconv.Atoi(r.PostForm.Get("roles"))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = app.user.Update(id64, email, fname, lname, age, address, phone, roles, password, activated)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Redirect to the user's profile page
		http.Redirect(w, r, fmt.Sprintf("/profile/%s", id), http.StatusSeeOther)
		return
	}

	// Get the user record from the database
	user, err := app.user.GetByID(id64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Render the template
	err = ts.Execute(w, user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

//recent code - now

// func (app *application) updateRecord(w http.ResponseWriter, r *http.Request) {
// 	// //Get the ID of the record to update from the request URL parameter
// 	id := r.FormValue("id")
// 	fmt.Println(id)
// 	id64, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	// //Get the user record from the database
// 	user, err := app.user.GetByID(id64)
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	// // Update the record in the database
// 	// if r.Method == "POST" {
// 	// 	r.ParseForm()
// 	// 	fname := r.PostForm.Get("firstname")
// 	// 	lname := r.PostForm.Get("lastname")
// 	// 	age, err := strconv.Atoi(r.PostForm.Get("age"))
// 	// 	if err != nil {
// 	// 		log.Println(err.Error())
// 	// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	// 		return
// 	// 	}
// 	// 	phone := r.PostForm.Get("phone")
// 	// 	address := r.PostForm.Get("address")
// 	// 	email := r.PostForm.Get("email")
// 	// 	password := []byte(r.PostForm.Get("password"))
// 	// 	activated, err := strconv.ParseBool(r.PostForm.Get("activated"))
// 	// 	if err != nil {
// 	// 		log.Println(err.Error())
// 	// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	// 		return
// 	// 	}
// 	// 	roles, err := strconv.Atoi(r.PostForm.Get("roles"))
// 	// 	if err != nil {
// 	// 		log.Println(err.Error())
// 	// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	// 		return
// 	// 	}

// 	// 	err = app.user.Update(id64, email, fname, lname, age, address, phone, roles, password, activated)
// 	// 	if err != nil {
// 	// 		log.Println(err.Error())
// 	// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	// 		return
// 	// 	}

// 	// 	// Redirect to the user's profile page
// 	// 	http.Redirect(w, r, fmt.Sprintf("/profile/%s", id), http.StatusSeeOther)
// 	// 	return
// 	// }
// 	data := &templateData{ //putting flash into template data
// 		User: user,
// 	}

// 	RenderTemplate(w, "update.user.page.tmpl", data)
// }
