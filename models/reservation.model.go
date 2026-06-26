package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`

	ParkingZoneID uint        `gorm:"not null" json:"parking_zone_id"`
	ParkingZone   ParkingZone `gorm:"foreignKey:ParkingZoneID" json:"parking_zone"`

	VehicleNumber  string    `gorm:"size:20;not null" json:"vehicle_number"`
	IsEV           bool      `gorm:"default:false" json:"is_ev"`
	StartTime      time.Time `gorm:"not null" json:"start_time"`
	EndTime        time.Time `gorm:"not null" json:"end_time"`
	Status         string    `gorm:"size:20;default:pending" json:"status"`
}