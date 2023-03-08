package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title          string
	Description    string
	Location       string
	StartTime      string
	EndTime        string
	Owner          string
	Subscribers    []User `gorm:"many2many:event_subscribers;"`
	MaxSubscribers int
}
