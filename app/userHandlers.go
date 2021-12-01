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
	familyMembers, err := uh.service.GetAllFamilyMembers(userId)
	if err != nil || familyMembers == nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, familyMembers)
	}
}

func (uh UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)
	userid, err := uh.service.CreateUser(user)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}
	writeResponse(w, http.StatusOK, userid)
}

func (uh UserHandlers) GetUserByUserId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userid")
	user, err := uh.service.GetUserByUserId(userId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}
	writeResponse(w, http.StatusOK, user)
}

func (uh UserHandlers) SearchUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	searchText := r.URL.Query()["search"][0]
	key := r.URL.Query()["key"][0] // name,location,phone
	users, err := uh.service.SearchUser(key, searchText)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}
	writeResponse(w, http.StatusOK, users)
}

func (uh UserHandlers) FindRelationship(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := uh.service.FindRelationship()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
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
