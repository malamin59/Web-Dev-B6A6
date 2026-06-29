package repository

import (
	"spotsync-api/models"

	"gorm.io/gorm"
)

type ParkingZoneRepository interface {
	Create(zone *models.ParkingZone) error
	GetAll() ([]models.ParkingZone, error)
	FindByID(id uint) (*models.ParkingZone, error)
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


func (r *parkingZoneRepository) GetAll() ([]models.ParkingZone, error) {

	var zones []models.ParkingZone

	if err := r.db.Find(&zones).Error; err != nil {
		return nil, err
	}

	return zones, nil
}

func (r *parkingZoneRepository) FindByID(id uint) (*models.ParkingZone, error) {

	var zone models.ParkingZone

	if err := r.db.First(&zone, id).Error; err != nil {
		return nil, err
	}

	return &zone, nil
}