package models

import "time"

// Heroes struct
type Hero struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Superpower string `json:"superpower"`
}

// Villains struct
type Villain struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	NemesisID int    `json:"nemesis_id"`
}

// CriminalReports struct
type CriminalReport struct {
	ID              int       `json:"id"`
	HeroID          int       `json:"hero_id"`
	VillainID       int       `json:"villain_id"`
	Description     string    `json:"description"`
	TimeOfIncident  time.Time `json:"time_of_incident"`
}

