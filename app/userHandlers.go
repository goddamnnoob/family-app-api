package app

import (
	"encoding/json"
	"net/http"

	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/service"
	"github.com/julienschmidt/httprouter"
)

type UserHandlers struct {
	service service.UserService
}

func (uh UserHandlers) getAllFamilyMembers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userid")
	users, err := uh.service.GetAllFamilyMembers(userId)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	}
	writeResponse(w, http.StatusAccepted, users)
}

func (uh UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)
	userid, err := uh.service.CreateUser(user)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	}
	writeResponse(w, http.StatusAccepted, userid+" Sucessful ")
}

func (uh UserHandlers) GetUserByUserId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userid")
	users, err := uh.service.GetUserByUserId(userId)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	}
	writeResponse(w, http.StatusAccepted, users)
}

func writeResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		panic(err)
	}
}
