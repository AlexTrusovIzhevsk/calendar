package commands

import (
	"calendar/dal"
	"strconv"
)

type DeleteEventCommand struct {
	Id int
}

type DeleteEventHandler struct {
	eventRepository dal.EventRepository
}

func NewDeleteEventHandler() *DeleteEventHandler {
	return &DeleteEventHandler{
		eventRepository: dal.GetEventRepository(),
	}
}

func (h *DeleteEventHandler) DeleteEvent(command DeleteEventCommand) (string, error) {
	err := h.eventRepository.DeleteEvent(command.Id)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(command.Id), nil
}
