package bl

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func ReadUuid(uuidString string) (*uuid.UUID, error) {
	parsedUuid, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, err
	}
	return &parsedUuid, nil
}

func ReadTime(timeString string) (*time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02", timeString)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}

func ToTimeString(t *time.Time) (string, error) {
	if t == nil {
		return "", fmt.Errorf("time is nil")
	}
	return t.Format("2006-01-02"), nil
}
