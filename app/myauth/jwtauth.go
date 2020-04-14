package myauth

import (
	"fmt"
	"log"
	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/pkg/errors"

	"github.com/titpetric/factory/resputil"
)

type JWT struct {
	tokenClaim string
	tokenAuth  *jwtauth.JWTAuth
}

//New JWT object
func (JWT) New() *JWT {
	jwt := &JWT{
		tokenClaim: "roles",
		tokenAuth:  jwtauth.New("HS256", []byte("mysecret"), nil),
	}

	// tokenClaim := jwtgo.MapClaims{
	// 	"user_id": "john",
	// 	"expiry":  time.Now().Add(time.Hour * 1).Unix(),
	// 	"roles":   "admin",
	// }
	// //
	// log.Println("DEBUG JWT:", jwt.Encode(tokenClaim))
	return jwt
}

func (jwt *JWT) Encode(claims jwtgo.MapClaims) string {
	_, tokenString, _ := jwt.tokenAuth.Encode(claims)
	return tokenString

}

func (jwt *JWT) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(jwt.tokenAuth)
}

func (jwt *JWT) Decode(r *http.Request) []string {
	val, _ := jwt.Authenticate(r)
	return val
}

func (jwt *JWT) Authenticate(r *http.Request) ([]string, error) {
	token, claims, err := jwtauth.FromContext(r.Context())
	log.Println("DEBUG JWT: Token:: ", token)
	log.Println("DEBUG JWT: claims:: ", claims)
	if claims == nil {
		return []string{}, errors.New("No roles found in token")
	}
	myclaims := claims[jwt.tokenClaim]
	if myclaims == nil {
		return []string{}, errors.New("No roles found in claims")
	}
	_, ok := myclaims.([]interface{})
	if !ok {
		return []string{}, errors.New("Invalid roles found in claims")
	}

	var tempclaims []interface{} = myclaims.([]interface{})
	roles := make([]string, len(tempclaims))
	for i, v := range tempclaims {
		roles[i] = fmt.Sprint(v)
	}

	log.Println("DEBUG JWT: Token valid claims:: ", roles)

	if err != nil || token == nil {
		return roles, errors.Wrap(err, "Empty or invalid JWT")
	}
	if !token.Valid {
		return roles, errors.New("Invalid JWT")
	}

	if claims == nil {
		return roles, errors.New("No claims found")
	}

	return roles, nil
}

func (jwt *JWT) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO Authenticate with DB, then set the JWT Token
			_, err := jwt.Authenticate(r)
			if err != nil {
				resputil.JSON(w, err)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
