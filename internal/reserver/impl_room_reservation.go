package reserver

import (
	"net/http"

	"github.com/brano-hozza/reserver-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateReservation - Create new reservation
func (this *implRoomReservationAPI) CreateReservation(ctx *gin.Context) {
	value, exists := ctx.Get("reservation_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[RoomReservation])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	reservation := RoomReservation{}
	err := ctx.BindJSON(&reservation)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	if reservation.Id == "@new" {
		reservation.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, reservation.Id, &reservation)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			reservation,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Reservation already exists",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create reservation in database",
				"error":   err.Error(),
			},
		)
	}
}

// DeleteReservation - Delete reservation
func (this *implRoomReservationAPI) DeleteReservation(ctx *gin.Context) {
	value, exists := ctx.Get("reservation_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[RoomReservation])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	id := ctx.Param("id")
	err := db.DeleteDocument(ctx, id)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Reservation not found",
					"error":   err.Error(),
				},
			)
		default:
			ctx.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to delete reservation from database",
					"error":   err.Error(),
				},
			)
		}
		return
	}

	ctx.JSON(
		http.StatusNoContent,
		nil,
	)
}

// GetReservationById - Provides the reservation by id
func (this *implRoomReservationAPI) GetReservationById(ctx *gin.Context) {
	value, exists := ctx.Get("reservation_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[RoomReservation])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	id := ctx.Param("id")
	reservation, err := db.FindDocument(ctx, id)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Reservation not found",
					"error":   err.Error(),
				},
			)
		default:
			ctx.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to get reservation from database",
					"error":   err.Error(),
				},
			)
		}
		return
	}

	ctx.JSON(
		http.StatusOK,
		reservation,
	)
}

// GetReservations - Provides the list of current reservations
func (this *implRoomReservationAPI) GetReservations(ctx *gin.Context) {
	value, exists := ctx.Get("reservation_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[RoomReservation])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	reservations, err := db.FindAllDocuments(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to get reservations from database",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		reservations,
	)

}

// GetRooms - Provides the list of all rooms
func (this *implRoomReservationAPI) GetRooms(ctx *gin.Context) {
	val, exists := ctx.Get("room_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := val.(db_service.DbService[Room])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	rooms, err := db.FindAllDocuments(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to get rooms from database",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		rooms,
	)
}

// UpdateReservation - Update reservation
func (this *implRoomReservationAPI) UpdateReservation(ctx *gin.Context) {
	value, exists := ctx.Get("reservation_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[RoomReservation])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	reservation := RoomReservation{}
	err := ctx.BindJSON(&reservation)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	err = db.UpdateDocument(ctx, reservation.Id, &reservation)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Reservation not found",
					"error":   err.Error(),
				},
			)
		default:
			ctx.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to update reservation in database",
					"error":   err.Error(),
				},
			)
		}
		return
	}

	ctx.JSON(
		http.StatusOK,
		reservation,
	)
}
