package model

type Apartment struct {
	ID            int    // ID is the unique identifier for the apartment.
	ApartmentName string // ApartmentName represents the name of the apartment and has a maximum length of 20 characters.
	Description   string
}
