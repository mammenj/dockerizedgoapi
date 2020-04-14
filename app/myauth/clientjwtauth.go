package myauth

import (
	"myapp/model"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)
// JWTClient jwclient
type JWTClient struct {
	tokenClaim string
	tokenAuth  *jwtauth.JWTAuth
}

//New JWTClient object
func (JWTClient) New(user *model.User) string {

	var roles []string

	for _, role := range user.Roles {
		if role.Roleid != "" {
			roles = append(roles, role.Roleid)
		}
	}

	tokenClaim := jwtgo.MapClaims{
		"user_id": user.Userid,
		"expiry":  time.Now().Add(time.Hour * 1).Unix(),
		"roles":  roles,
	}
	//

	jwt := &JWTClient{
		tokenClaim: "roles",
		tokenAuth:  jwtauth.New("HS256", []byte("mysecret"), nil),
	}

	return jwt.encode(tokenClaim)
}

func (jwt *JWTClient) encode(claims jwtgo.MapClaims) string {
	_, tokenString, _ := jwt.tokenAuth.Encode(claims)
	return tokenString
}
