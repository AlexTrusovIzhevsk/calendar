package api

import (
	"calendar/bl/queries"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetEventsByDateDto struct {
	Date string
}

type GetEventsByDateItemDto struct {
	Id     int
	UserId string
	Date   string
	Event  string
}

type GetEventsByDateListDto struct {
	Events []GetEventsByDateItemDto
}

func getEventsByDateHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	var dto GetEventsByDateDto

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	query := queries.NewGetEventsByDateHandler()

	list, err := query.GetEventsByDate(queries.GetEventsByDateQuery(dto))
	if err != nil {
		returnError(err, writer)
		return
	}

	items := make([]GetEventsByDateItemDto, 0, len(list))
	for _, event := range list {
		items = append(items, GetEventsByDateItemDto(event))
	}

	result := GetEventsByDateListDto{
		Events: items,
	}

	jsonBytes, _ := json.Marshal(result)
	writer.Write(jsonBytes)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
