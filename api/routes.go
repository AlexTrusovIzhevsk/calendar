package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", getEventsByDateHandler)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Calendar app"))
	if err != nil {
		fmt.Println(err)
	}
}

type SuccessDto struct {
	Result string
}

type ErrorDto struct {
	Error string
}

func returnError(err error, writer http.ResponseWriter) {
	errorDto := ErrorDto{Error: err.Error()}
	jsonBytes, _ := json.Marshal(errorDto)
	errorBody := string(jsonBytes)
	http.Error(writer, errorBody, http.StatusInternalServerError)
}
