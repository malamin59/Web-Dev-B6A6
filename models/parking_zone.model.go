package models

import "gorm.io/gorm"

type ParkingZone struct {
	gorm.Model

	Name            string `gorm:"size:100;not null" json:"name"`
	Location        string `gorm:"size:255;not null" json:"location"`
	TotalSpots      int    `gorm:"not null" json:"total_spots"`
	AvailableSpots  int    `gorm:"not null" json:"available_spots"`
	EVChargingSpots int    `gorm:"not null" json:"ev_charging_spots"`

	Reservations []Reservation `gorm:"foreignKey:ParkingZoneID" json:"reservations,omitempty"`
}