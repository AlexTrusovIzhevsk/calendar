package api

import (
	"calendar/bl/commands"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateEventDto struct {
	UserId string
	Date   string
	Event  string
}

func createEventHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	var dto CreateEventDto

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	handler := commands.NewCreateEventHandler()

	result, err := handler.CreateEvent(commands.CreateEventCommand(dto))
	if err != nil {
		returnError(err, writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	fmt.Fprint(writer, SuccessDto{Result: result})
}
