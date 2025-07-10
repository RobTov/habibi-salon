package db

import (
	"context"
	"log"

	"github.com/RobTov/habibi-salon/internal/store"
)

func Seed(store store.Storage) {
	ctx := context.Background()

	services := generateServices()
	for _, service := range services {
		if err := store.Services.Create(ctx, service); err != nil {
			log.Println("Error creating a service: ", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateServices() []*store.Services {
	services := []*store.Services{
		{
			Name:        "Uñas nuevas",
			Description: "Lorem ipsum dolor sit amet",
			Price:       1500,
			IsActive:    true,
		},
		{
			Name:        "Rellenos",
			Description: "Lorem ipsum dolor sit amet",
			Price:       1000,
			IsActive:    true,
		},
		{
			Name:        "Decoraciones",
			Description: "Lorem ipsum dolor sit amet",
			Price:       50,
			IsActive:    true,
		},
		{
			Name:        "Keratina",
			Description: "Lorem ipsum dolor sit amet",
			Price:       1200,
			IsActive:    true,
		},
		{
			Name:        "Alisado",
			Description: "Lorem ipsum dolor sit amet",
			Price:       1000,
			IsActive:    true,
		},
		{
			Name:        "Tintes",
			Description: "Lorem ipsum dolor sit amet",
			Price:       600,
			IsActive:    true,
		},
		{
			Name:        "Podología",
			Description: "Lorem ipsum dolor sit amet",
			Price:       1500,
			IsActive:    true,
		},
		{
			Name:        "Masaje Facial",
			Description: "Lorem ipsum dolor sit amet",
			Price:       1000,
			IsActive:    true,
		},
		{
			Name:        "Servicio Inactivo",
			Description: "Lorem ipsum dolor sit amet",
			Price:       1000,
			IsActive:    false,
		},
	}

	return services
}
