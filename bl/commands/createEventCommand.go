package commands

import (
	"calendar/bl"
	"calendar/dal"
	"errors"
	"strconv"
)

type CreateEventCommand struct {
	UserId string
	Date   string
	Event  string
}

type CreateEventHandler struct {
	eventRepository dal.EventRepository
}

func NewCreateEventHandler() *CreateEventHandler {
	return &CreateEventHandler{
		eventRepository: dal.GetEventRepository(),
	}
}

func (h *CreateEventHandler) CreateEvent(command CreateEventCommand) (string, error) {
	if command.Event == "" {
		return "", errors.New("event is required")
	}

	userId, err := bl.ReadUuid(command.UserId)
	if err != nil {
		return "", err
	}

	time, err := bl.ReadTime(command.Date)
	if err != nil {
		return "", err
	}

	entity := dal.EventEntity{
		Id:     0,
		Event:  command.Event,
		Date:   *time,
		UserId: *userId,
	}

	err = h.eventRepository.CreateEvent(&entity)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(entity.Id), nil
}
