package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bersennaidoo/farmstyle/application/rest/problems"
	"github.com/bersennaidoo/farmstyle/domain/models"
	"github.com/bersennaidoo/farmstyle/infrastructure/repositories/mongo"
	"github.com/kataras/golog"
)

type UserHandler struct {
	userRepository *mongo.UserRepository
	log            *golog.Logger
}

func NewUserHandler(userRepository *mongo.UserRepository, log *golog.Logger) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
		log:            log,
	}
}

func (u *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.NewUser
	err := decoder.Decode(&user)
	if err != nil {
		ErrorResponse(problems.FailedToParseJson(problems.ProblemJson{
			Detail: err.Error(),
		}))(w, r)
		return
	}
	res, createErr := u.userRepository.AddUser(user)
	if createErr != nil {
		ErrorResponse(createErr.(*problems.ProblemJson))(w, r)
		return
	}
	writeJson(201, res)(w, r)
}
