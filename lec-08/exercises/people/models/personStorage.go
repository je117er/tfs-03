package models

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrPersonNotFound = errors.New("person not found")
)

// PersonStorage wraps a gorm.DB connection pool
type PersonStorage struct {
	DB *gorm.DB
}

func NewPersonStorage(db *gorm.DB) PersonStorage {
	return PersonStorage{
		DB: db,
	}
}

// Insert adds a new person into the database.
func (p PersonStorage) Insert(person PersonDBModel) error {
	result := p.DB.Create(&person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ByID gets a user from the database.
func (p PersonStorage) ByID(id uint64) (*PersonDBModel, error) {
	var person PersonDBModel
	result := p.DB.First(&person, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrPersonNotFound
		}
		return nil, result.Error
	}
	return &person, nil
}

// Update modifies a user
func (p *PersonStorage) Update(person *PersonDBModel) error {
	result := p.DB.Save(person)
	return result.Error
}

// Delete removes a user
func (p *PersonStorage) Delete(id uint64) error {
	result := p.DB.Delete(&PersonDBModel{}, id)
	return result.Error
}

// All returns all users
func (p *PersonStorage) All() ([]PersonDBModel, error) {
	var people []PersonDBModel
	result := p.DB.Find(&people)
	if result.Error != nil {
		return nil, result.Error
	}
	return people, nil
}
