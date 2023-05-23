// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave, noSurf)

	// from here
	router.Handler(http.MethodGet, "/", dynamicMiddleware.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/about", dynamicMiddleware.ThenFunc(app.about))
	router.Handler(http.MethodGet, "/login", dynamicMiddleware.ThenFunc(app.loginform))
	router.Handler(http.MethodPost, "/login", dynamicMiddleware.ThenFunc(app.loginformSubmit))
	router.Handler(http.MethodGet, "/register", dynamicMiddleware.ThenFunc(app.register))
	router.Handler(http.MethodPost, "/register", dynamicMiddleware.ThenFunc(app.registerSubmit))
	router.Handler(http.MethodGet, "/feedback", dynamicMiddleware.ThenFunc(app.feedbackDisplay))
	router.Handler(http.MethodPost, "/feedback", dynamicMiddleware.ThenFunc(app.feedbackFormSubmit))
	router.Handler(http.MethodGet, "/viewusers", dynamicMiddleware.ThenFunc(app.displayUsers))
	router.Handler(http.MethodGet, "/users/update", dynamicMiddleware.ThenFunc(app.updateRecord))
	router.Handler(http.MethodPost, "/users/update", dynamicMiddleware.ThenFunc(app.updateRecord))

	router.Handler(http.MethodGet, "/update", dynamicMiddleware.ThenFunc(app.updateRecord))

	//protected routes
	protected := dynamicMiddleware.Append(app.requireAuthenticationMiddleware)
	router.Handler(http.MethodGet, "/equipment", dynamicMiddleware.ThenFunc(app.equipment))
	//router.Handler(http.MethodGet, "/edit_equipment", dynamicMiddleware.ThenFunc(app.equipmentUpdate))
	//router.Handler(http.MethodPost, "/edit_equipment", dynamicMiddleware.ThenFunc(app.equipmentUpdateSubmit))
	router.Handler(http.MethodGet, "/user", protected.ThenFunc(app.userPortal))
	router.Handler(http.MethodPost, "/user", protected.ThenFunc(app.userPortalFormSubmit))
	router.Handler(http.MethodGet, "/admin", dynamicMiddleware.ThenFunc(app.adminPortal))
	router.Handler(http.MethodPost, "/admin", dynamicMiddleware.ThenFunc(app.adminPortalFormSubmit))
	//stop here

	//tidy up the middleware chain
	standardMiddleware := alice.New(
		app.recoverPanicMiddleware,
		app.logRequestMiddleware,
		securityHeadersMiddleware,
	)

	return standardMiddleware.Then(router)
}
