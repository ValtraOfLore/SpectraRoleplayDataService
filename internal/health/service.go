package health

import (
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (Service) GetHealth() HealthStatus {
	tm := time.Now().Format(time.RFC1123)

	return HealthStatus{
		Message:     "API Live",
		CurrentTime: tm,
	}
}
