package dto

type CreateParkingZoneRequest struct {
	Name            string `json:"name"`
	Location        string `json:"location"`
	TotalSpots      int    `json:"total_spots"`
	AvailableSpots  int    `json:"available_spots"`
	EVChargingSpots int    `json:"ev_charging_spots"`
}