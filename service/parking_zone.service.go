package service

import (
	"errors"
	"spotsync-api/dto"
	"spotsync-api/models"
	"spotsync-api/repository"
)

type ParkingZoneService interface {
	Create(req dto.CreateParkingZoneRequest) error
	GetAll() ([]models.ParkingZone, error)
	GetByID(id uint) (*models.ParkingZone, error)
}

type parkingZoneService struct {
	repo repository.ParkingZoneRepository
}

func NewParkingZoneService(repo repository.ParkingZoneRepository) ParkingZoneService {
	return &parkingZoneService{
		repo: repo,
	}
}

func (s *parkingZoneService) Create(req dto.CreateParkingZoneRequest) error {

	// Validation
	if req.TotalSpots <= 0 {
		return errors.New("total spots must be greater than 0")
	}

	if req.EVChargingSpots > req.TotalSpots {
		return errors.New("EV charging spots cannot exceed total spots")
	}

	// DTO → Model
	zone := models.ParkingZone{
		Name:            req.Name,
		Location:        req.Location,
		TotalSpots:      req.TotalSpots,
		AvailableSpots:  req.TotalSpots, // শুরুতে সব Spot Available
		EVChargingSpots: req.EVChargingSpots,
	}

	// Save
	return s.repo.Create(&zone)
}

/* GET PARKING ZONE */
func (s *parkingZoneService) GetAll() ([]models.ParkingZone, error) {
	return s.repo.GetAll()
}


/* GET SINGLE PARING ZONE BY "ID" */
func (s *parkingZoneService) GetByID(id uint) (*models.ParkingZone, error) {

	zone, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return zone, nil
}