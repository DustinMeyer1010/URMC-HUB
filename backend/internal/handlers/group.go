package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/gorilla/mux"
)

func AddUsersToGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	group := vars["group"]

	var modify models.ModifyMembers

	err := json.NewDecoder(r.Body).Decode(&modify)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ad.AddUsersToGroup(group, modify.Members)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func RemoveUsersFromGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	group := vars["group"]

	fmt.Println("here")

	var modify models.ModifyMembers

	err := json.NewDecoder(r.Body).Decode(&modify)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = ad.RemoveUsersFromGroup(group, modify.Members)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
