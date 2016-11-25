package RESTServer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Accountjson struct {
	ID       string `json:"id,omitempty"`
	Password string `json:"pw,omitempty"`
	Name     string `json:"name,omitempty"`
}

var m = NewMemoryDataAccess()

func GetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	getID := params["id"]
	account, err := m.GetAccount(getID)
	if err != nil {
		fmt.Println("GetAccount Error!")
		return
	}
	json.NewEncoder(w).Encode(account)
}

func PutHandler(w http.ResponseWriter, r *http.Request) {

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var newAccount Account

	_ = json.NewDecoder(r.Body).Decode(&newAccount)
	_, err := m.CreateAccount(newAccount)
	if err != nil {
		fmt.Printf("apiPostHandler : CreateAccount Error!")
	}
	json.NewEncoder(w).Encode(newAccount)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

}
