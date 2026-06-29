package service

import (
	"errors"

	"spotsync-api/dto"
	"spotsync-api/models"
	"spotsync-api/repository"
)

type ReservationService interface {
	Create(userID uint, req dto.CreateReservationRequest) error
	GetMyReservations(userID uint) ([]models.Reservation, error)
	Cancel(userID uint, reservationID uint) error
	GetAll() ([]models.Reservation, error)
}

type reservationService struct {
	reservationRepo repository.ReservationRepository
	parkingRepo     repository.ParkingZoneRepository
}

func NewReservationService(
	reservationRepo repository.ReservationRepository,
	parkingRepo repository.ParkingZoneRepository,
) ReservationService {

	return &reservationService{
		reservationRepo: reservationRepo,
		parkingRepo:     parkingRepo,
	}
}

func (s *reservationService) Create(userID uint, req dto.CreateReservationRequest) error {

	// Check Parking Zone
	zone, err := s.parkingRepo.FindByID(req.ParkingZoneID)
	if err != nil {
		return errors.New("parking zone not found")
	}

	// Check Available Spot
	if zone.AvailableSpots <= 0 {
		return errors.New("parking zone is full")
	}

	// Create Reservation
	reservation := models.Reservation{
		UserID:        userID,
		ParkingZoneID: req.ParkingZoneID,
		VehicleNumber: req.VehicleNumber,
		IsEV:          req.IsEV,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		Status:        "pending",
	}

	// Save Reservation
	if err := s.reservationRepo.Create(&reservation); err != nil {
		return err
	}

	// Decrease Available Spots
	zone.AvailableSpots--

	// Update Parking Zone
	if err := s.parkingRepo.Update(zone); err != nil {
		return err
	}

	return nil
}

func (s *reservationService) GetMyReservations(userID uint) ([]models.Reservation, error) {
	return s.reservationRepo.GetByUserID(userID)
}

func (s *reservationService) Cancel(userID uint, reservationID uint) error {

	// Find Reservation
	reservation, err := s.reservationRepo.FindByID(reservationID)
	if err != nil {
		return errors.New("reservation not found")
	}

	// Check Reservation Owner
	if reservation.UserID != userID {
		return errors.New("access denied")
	}

	// Already Cancelled
	if reservation.Status == "cancelled" {
		return errors.New("reservation already cancelled")
	}

	// Completed Reservation Cannot Be Cancelled
	if reservation.Status == "completed" {
		return errors.New("completed reservation cannot be cancelled")
	}

	// Find Parking Zone
	zone, err := s.parkingRepo.FindByID(reservation.ParkingZoneID)
	if err != nil {
		return errors.New("parking zone not found")
	}

	// Update Reservation Status
	reservation.Status = "cancelled"

	if err := s.reservationRepo.Update(reservation); err != nil {
		return err
	}

	// Release Parking Spot
	zone.AvailableSpots++

	if err := s.parkingRepo.Update(zone); err != nil {
		return err
	}

	return nil
}

func (s *reservationService) GetAll() ([]models.Reservation, error) {
	return s.reservationRepo.GetAll()
}
