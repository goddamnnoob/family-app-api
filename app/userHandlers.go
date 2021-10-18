package app

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		fmt.Fprintf(w, err.AsMessage().Message)
	}
	json.NewEncoder(w).Encode(users)
}
