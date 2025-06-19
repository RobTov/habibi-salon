package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Appointments struct {
	ID        int64     `json:"id"`
	Date      time.Time `json:"date"`
	ClientID  int64     `json:"client_id"`
	ServiceID int64     `json:"service_id"`
	// TODO: change the status to an enum
	// * status can be 'pending', 'confirmed' and 'canceled'
	Status string `json:"status"`
}

type AppointmentsStore struct {
	db *sql.DB
}

func (s *AppointmentsStore) Create(ctx context.Context, appointment *Appointments) error {
	query := `
	INSERT INTO appointments (date, client_id, service_id, status)
	VALUES ($1, $2, $3);
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, appointment.Date, appointment.ClientID, appointment.ServiceID, appointment.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *AppointmentsStore) GetAll(ctx context.Context) ([]Appointments, error) {
	query := `
	SELECT id, date, client_id, service_id, status
	FROM appointments; 
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	appointments := []Appointments{}
	for rows.Next() {
		var a Appointments

		err := rows.Scan(
			&a.ID,
			&a.Date,
			&a.ClientID,
			&a.ServiceID,
			&a.Status,
		)

		if err != nil {
			return nil, err
		}

		appointments = append(appointments, a)
	}

	return appointments, nil
}

func (s *AppointmentsStore) GetByID(ctx context.Context, appointmentID int64) (*Appointments, error) {
	query := `
	SELECT id, date, client_id, service_id, status
	FROM appointments
	WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var appointment Appointments
	err := s.db.QueryRowContext(ctx, query, appointmentID).Scan(
		&appointment.ID,
		&appointment.Date,
		&appointment.ClientID,
		&appointment.ServiceID,
		&appointment.Status,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}

	}

	return &appointment, nil
}

func (s *AppointmentsStore) Delete(ctx context.Context, appointmentID int64) error {
	query := `
	DELETE FROM appointments WHERE id = $1;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, appointmentID)
	if err != nil {
		return err
	}

	return nil
}
