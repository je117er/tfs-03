package application

import (
	"log"
	"people/models"
)

type Application struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	PersonStorage *models.PersonStorage
}
