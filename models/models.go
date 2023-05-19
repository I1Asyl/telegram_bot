package models

import "time"

type User struct {
	NationalId string
	FirstName  string
	LastName   string
	Language   string
	BirthDate  time.Time
	Email      string
	Phone      string
	Gender     string
}

type Patient struct {
	PatientId  int
	NationalId string
	Weight     float64
	Height     float64
	BloodType  string
	Allergies  string
	Conditions string
}

type Doctor struct {
	DoctorId       int
	NationalId     string
	Specialization string
}

type Connection struct {
	ChatId     int64
	NationalId string
	Question   string
}
