package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Contact struct {
	Id          int    `json:"id"`
	Name        string `json:"name" gorm:"uniqueIndex,not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
}

func AllContacts(db *gorm.DB, q string) ([]Contact, error) {
	var contacts []Contact

	if q == "" {
		db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}}).Find(&contacts)
		return contacts, nil
	}
	db.Where("phone_number LIKE ?", "%"+q+"%").Order("name ASC").Find(&contacts)
	fmt.Println("q:", q)

	return contacts, nil
}

func GetById(db *gorm.DB, id int) Contact {
	var contact Contact

	db.First(&contact, id)

	return contact
}
