// handlers.go
package main

import (
	"errors"
	"html/template"
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
			RenderTemplate(w, "register.page.tmpl", nil)
		}
	}
	app.sessionManager.Put(r.Context(), "flash", "SignupWassuccessful")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) userPortal(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "userPortal.page.tmpl", data)

}

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
func (app *application) equipment(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("./ui/html/equipments.page.tmpl", "./ui/html/base.layout.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	equipmentTypes, err := app.equipments.Display()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	data := struct {
		EquipmentTypes []models.EquipmentType
	}{
		EquipmentTypes: equipmentTypes,
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

/*unc (app *application) equipmentUpdate(w http.ResponseWriter, r *http.Request){
		ts, err := template.ParseFiles("./ui/html/equ_admin.page.tmpl", "./ui/html/base.layout.tmpl")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		equipmentTypes, err := app.equipments.Display()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data := struct {
			EquipmentTypes []models.EquipmentType
		}{
			EquipmentTypes: equipmentTypes,
		}

		err = ts.Execute(w, data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
}

func (app *application) equipmentUpdateSubmit(w http.ResponseWriter, r *http.Request) {
	log.Println("hello there")
	r.ParseForm()
	status := r.PostForm.Get("status") //"name" is the name of the form
	available:= r.PostForm.Get("available")
	id := r.PostForm.Get("id")
	button:= r.PostForm.Get("myButton")
	//lets write the data to the table
	if button == "delete" {
		err := app.equipments.Delete(id)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err := app.equipments.Update(status, available, id)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}


func (app *application) test(w http.ResponseWriter, r *http.Request) {
	log.Println("hello there")
}*/



func (app *application) adminPortal(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/equ_admin.page.tmpl", "./ui/html/base.layout.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	equipmentTypes, err := app.equipments.Display()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := struct {
		EquipmentTypes []models.EquipmentType
	}{
		EquipmentTypes: equipmentTypes,
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Println("hello there")
}
func (app *application) adminPortalFormSubmit(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests to this handler
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse the form values
    err := r.ParseForm()
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Get the ID of the equipment
    id := r.FormValue("id")

    // Get the button that was clicked
    button := r.FormValue("myButton")

    // Get the status and availability values based on the button clicked
    var status string
    var availability string
    switch button {
    case "update":
        status = r.FormValue("status")
        availability = r.FormValue("available")
    case "delete":
        // Call function to delete equipment from database
        // ...
		err := app.equipments.Delete(id)
		if err != nil {
			log.Println("an error occured when passing the data", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
        http.Redirect(w, r, "/admin", http.StatusSeeOther)
        return
    default:
        http.Error(w, "Invalid button value", http.StatusBadRequest)
        return
    }
	err = app.equipments.Update(status, availability, id)
	if err != nil {
		log.Println("an error occured when passing the data", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
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
