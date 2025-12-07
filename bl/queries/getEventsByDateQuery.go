package queries

import (
	"calendar/bl"
	"calendar/dal"
	"time"
)

type GetEventsByDateQuery struct {
	Date string
}

type GetEventsByDateResult struct {
	Id     int
	UserId string
	Date   string
	Event  string
}

type GetEventsByDateHandler struct {
	eventRepository dal.EventRepository
}

func NewGetEventsByDateHandler() *GetEventsByDateHandler {
	return &GetEventsByDateHandler{
		eventRepository: dal.GetEventRepository(),
	}
}

func (h *GetEventsByDateHandler) GetEventsByDate(command GetEventsByDateQuery) ([]GetEventsByDateResult, error) {
	date, err := bl.ReadTime(command.Date)
	if err != nil {
		return nil, err
	}

	year, month, day := date.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	end := time.Date(year, month, day, 23, 59, 59, 1e9-1, date.Location())

	list := h.eventRepository.GetEventsBetween(start, end)

	result := make([]GetEventsByDateResult, len(list))
	for i, event := range list {
		date, err := bl.ToTimeString(&event.Date)
		if err != nil {
			return nil, err
		}

		result[i] = GetEventsByDateResult{
			Id:     event.Id,
			Date:   date,
			Event:  event.Event,
			UserId: event.UserId.String(),
		}
	}

	return result, nil
}
