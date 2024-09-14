package domain

import (
	"time"
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Organization string    `json:"organization"`
	Rating       Rating    `json:"rating"`
	Date         time.Time `json:"date"`
	ImageURL     string    `json:"image_url"`
	Capacity     int       `json:"capacity"`
	Price        float64   `json:"price"`
	PartnerID    int       `json:"partner_id"`
	Spots        []Spots   `json:"spots"`
	Tickets      []Tickets `json:"tickets"`
}
