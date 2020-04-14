package middleware

import (
	"log"
	"myapp/app/myauth"
	"net/http"

	"github.com/casbin/casbin"
)

func Authorizer(e *casbin.Enforcer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			//user, _, _ := r.BasicAuth()
			myjwt := myauth.JWT{}.New()
			roles, err := myjwt.Authenticate(r)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			}

			method := r.Method
			path := r.URL.Path

			// looop the rolese here
			authorized := false
			for _, role := range roles {
				log.Println("DEBUG JWT: Authorizer ::Role: " + role + " ::path: " + path + " ::method: " + method)
				if e.Enforce(role, path, method) {
					authorized = true
					break
				} else {
					authorized = false
				}
			}
			if authorized {
				next.ServeHTTP(w, r)
			} else {
				log.Println("DEBUG JWT: Authorizer :: USER/Role is NOT authorized")
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			}
		}

		return http.HandlerFunc(fn)
	}
}
