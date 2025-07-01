package models

import (
	"time"
)

// TODO : ACT NO 3 
type Event struct {
	ID 			int 		`json:"id_event"`
	Title 		string 		`json:"title"`
	Organizer 	string		`json:"organizer"`
	Description	string		`json:"description"`
	Date		time.Time	`json:"date"`
	Location	string		`json:"location"`
	Price		float64		`json:"price"`
	Capacity	int			`json:"capacity"`
	ImageURL	string		`json:"image_url"`
	Status		string		`json:"status"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}