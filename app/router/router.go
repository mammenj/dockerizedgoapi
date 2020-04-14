package router

import (
	"errors"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/go-chi/chi"
	"github.com/titpetric/factory/resputil"

	"myapp/app/app"
	"myapp/app/myauth"

	"myapp/app/requestlog"
	"myapp/app/router/middleware"
)

// New router created with App
func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()
	login := myauth.JWT{}.New()
	r.Use(login.Verifier())
	// Protected API endpoints
	r.Group(func(mux chi.Router) {
		// Error out on invalid/empty JWT here
		{
			mux.Use(login.Authenticator())
			e := casbin.NewEnforcer("/myapp/authz_model.conf", "/myapp/authz_policy.csv")
			mux.Use(middleware.Authorizer(e))
			mux.Route("/api/v1", func(r chi.Router) {
				r.Use(middleware.ContentTypeJson)
				r.Method("GET", "/books", requestlog.NewHandler(a.HandleListBooks, l))
				r.Method("GET", "/books", requestlog.NewHandler(a.HandleListBooks, l))
				r.Method("POST", "/books", requestlog.NewHandler(a.HandleCreateBook, l))
				r.Method("GET", "/books/{id}", requestlog.NewHandler(a.HandleReadBook, l))
				r.Method("PUT", "/books/{id}", requestlog.NewHandler(a.HandleUpdateBook, l))
				r.Method("DELETE", "/books/{id}", requestlog.NewHandler(a.HandleDeleteBook, l))
			})

			mux.Route("/api/v2", func(r chi.Router) {
				r.Use(middleware.ContentTypeJson)
				r.Method("GET", "/users", requestlog.NewHandler(a.HandleListUsers, l))
				r.Method("GET", "/users/{id}", requestlog.NewHandler(a.HandleGetUser, l))
				r.Method("POST", "/users", requestlog.NewHandler(a.HandleCreateUser, l))

			})
		}
	})

	// Public API endpoints
	r.Group(func(mux chi.Router) {
		// Print info about claim
		mux.Get("/info", func(w http.ResponseWriter, r *http.Request) {
			owner := login.Decode(r)
			resputil.JSON(w, owner, errors.New("not logged in"))
		})

		mux.Method("POST", "/token/{userid}/{password}", requestlog.NewHandler(a.HandleLoginUser, l))
		mux.Get("/healthz/liveness", app.HandleLive)
		mux.Method("GET", "/healthz/readiness", requestlog.NewHandler(a.HandleReady, l))
		mux.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))
	})

	return r
}
