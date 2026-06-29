package repository

import (
	"spotsync-api/models"

	"gorm.io/gorm"
)

type ParkingZoneRepository interface {
	Create(zone *models.ParkingZone) error
}

type parkingZoneRepository struct {
	db *gorm.DB
}

func NewParkingZoneRepository(db *gorm.DB) ParkingZoneRepository {
	return &parkingZoneRepository{
		db: db,
	}

}
func (r *parkingZoneRepository) Create(zone *models.ParkingZone) error {
	return r.db.Create(zone).Error
}
