package helpers

import (
	"fmt"
	"net/http"
	"people/application"
	"runtime/debug"
)

// ServerError helper writes an error message and stack trace
// to the errorLog then sends a generic 500 Internal Server Error
// to the user
//func ServerError(application *config.Application) func(w http.ResponseWriter, err error) {
func ServerError(app *application.Application, w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// ClientError helper sends a specific status code and corresponding desc
// to the user
func ClientError(app *application.Application, w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound wraps around ClientError and returns a 404 response to user
func NotFound(app *application.Application, w http.ResponseWriter) {
	ClientError(app, w, http.StatusNotFound)
}
