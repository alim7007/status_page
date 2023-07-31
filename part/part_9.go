package part

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleConnection).Methods("GET", "OPTIONS")
	http.ListenAndServe("127.0.0.1:8282", router)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var resultT ResultT
	rez := GetResultData()
	if len(rez.Incidents) == 0 || len(rez.SMS) == 0 || len(rez.MMS) == 0 || len(rez.VoiceCall) == 0 || len(rez.Email) == 0 || len(rez.Support) == 0 {
		resultT.Status = false
	} else {
		resultT.Status = true
	}
	if resultT.Status == false {
		resultT.Error = "Error on collect data"
	} else {
		resultT.Data = rez
	}
	bytes, err := json.Marshal(resultT)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
