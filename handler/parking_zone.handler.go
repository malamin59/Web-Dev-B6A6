package handler

import (
	"net/http"
	"strconv"

	"spotsync-api/dto"
	"spotsync-api/service"

	"github.com/labstack/echo/v4"
)

type ParkingZoneHandler struct {
	parkingService service.ParkingZoneService
}

func NewParkingZoneHandler(parkingService service.ParkingZoneService) *ParkingZoneHandler {
	return &ParkingZoneHandler{
		parkingService: parkingService,
	}
}

func (h *ParkingZoneHandler) Create(c echo.Context) error {

	var req dto.CreateParkingZoneRequest

	// Parse JSON
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	// Call Service
	if err := h.parkingService.Create(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// Success Response
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Parking zone created successfully",
	})
}

func (h *ParkingZoneHandler) GetAll(c echo.Context) error {

	zones, err := h.parkingService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, zones)
}

func (h *ParkingZoneHandler) GetByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parking zone id",
		})
	}

	zone, err := h.parkingService.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "parking zone not found",
		})
	}

	return c.JSON(http.StatusOK, zone)
}