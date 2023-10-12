package interfaces

import "net/http"

type HealthCheckController interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type PingService interface {
	Ping() error
}

type PingRepository interface {
	Ping() error
}
