package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"

	"myapp/app/myauth"
	"myapp/model"
	"myapp/repository"
)

func (app *App) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.ListUsers(app.db)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}

	if users == nil {
		fmt.Fprint(w, "[]")
		return
	}

	dtos := users.ToDto()
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		return
	}
}

func (app *App) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	userf := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(userf); err != nil {
		app.logger.Warn().Err(err).Msg("Unable to decode the JSON User")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		return
	}

	user, err := repository.CreateUser(app.db, userf)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataCreationFailure)
		return
	}

	app.logger.Info().Msgf("New user created: %d", user.Name)
	w.WriteHeader(http.StatusCreated)
}

//HandleLoginUser get the User by username/password
func (app *App) HandleLoginUser(w http.ResponseWriter, r *http.Request) {

	userid := chi.URLParam(r, "userid")
	password := chi.URLParam(r, "password")

	app.logger.Log().Msg("Userid is ::" + userid + "::")
	user, err := repository.GetUserByUserid(app.db, userid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}
	app.logger.Log().Msgf("User is :: ", user)
	success := myauth.ComparePasswords(user.Password, []byte(password))
	if !success {
		app.logger.Log().Msg("password is NOT Correct")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	app.logger.Warn().Msg("DEBUG JWT: <<<<<<<<<>>>>>>>User: " + user.Userid)
	jwtoken := myauth.JWTClient{}.New(user)
	log.Println("DEBUG JWT:", jwtoken)
	if err := json.NewEncoder(w).Encode(jwtoken); err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		return
	}

}

func (app *App) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || id == 0 {
		app.logger.Info().Msgf("can not parse ID: %v", id)

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user, err := repository.GetUser(app.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}

	dto := user.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		return
	}
}
