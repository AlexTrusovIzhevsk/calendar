package commands

import (
	"calendar/bl"
	"calendar/dal"
	"errors"
	"strconv"
)

type UpdateEventCommand struct {
	Id     int
	UserId string
	Date   string
	Event  string
}

type UpdateEventHandler struct {
	eventRepository dal.EventRepository
}

func NewUpdateEventHandler() *UpdateEventHandler {
	return &UpdateEventHandler{
		eventRepository: dal.GetEventRepository(),
	}
}

func (h *UpdateEventHandler) UpdateEvent(command UpdateEventCommand) (string, error) {
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

	entity, err := h.eventRepository.GetEventById(command.Id)
	if err != nil {
		return "", err
	}
	entity.UserId = *userId
	entity.Event = command.Event
	entity.Date = *time

	err = h.eventRepository.SaveEvent(entity)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(entity.Id), nil
}
