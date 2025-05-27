package store

import (
	"database/sql"
	"time"
)

type Appointments struct {
	ID        int64     `json:"id"`
	Date      time.Time `json:"created_at"`
	ClientID  int64     `json:"client_id"`
	ServiceID int64     `json:"service_id"`
	// TODO: change the status to an enum
	// * status can be 'pending', 'confirmed' and 'canceled'
	Status string `json:"status"`
}

type AppointmentsStore struct {
	db *sql.DB
}
