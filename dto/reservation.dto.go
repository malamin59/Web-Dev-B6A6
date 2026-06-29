package dto

import "time"

type CreateReservationRequest struct {
	ParkingZoneID uint      `json:"parking_zone_id"`
	VehicleNumber string    `json:"vehicle_number"`
	IsEV          bool      `json:"is_ev"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
}