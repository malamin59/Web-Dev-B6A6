package repository

import (
	"spotsync-api/models"

	"gorm.io/gorm"
)

type ReservationRepository interface {
	Create(reservation *models.Reservation) error
	FindByID(id uint) (*models.Reservation, error)
	Update(reservation *models.Reservation) error
	GetByUserID(userID uint) ([]models.Reservation, error)
	GetAll() ([]models.Reservation, error)
}


type reservationRepository struct {
	db *gorm.DB
}  

 func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationRepository{
		db: db,
	}
}

func (r *reservationRepository) Create(reservation *models.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *reservationRepository) GetByUserID(userID uint) ([]models.Reservation, error) {

	var reservations []models.Reservation

	err := r.db.
		Preload("ParkingZone").
		Where("user_id = ?", userID).
		Find(&reservations).Error

	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *reservationRepository) FindByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation

	err := r.db.First(&reservation, id).Error
	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *reservationRepository) Update(reservation *models.Reservation) error {
	return r.db.Save(reservation).Error
}

func (r *reservationRepository) GetAll() ([]models.Reservation, error) {

	var reservations []models.Reservation

	err := r.db.
		Preload("User").
		Preload("ParkingZone").
		Find(&reservations).Error

	if err != nil {
		return nil, err
	}

	return reservations, nil
}