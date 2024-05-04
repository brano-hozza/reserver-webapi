package reserver

import (
	"net/http"

	"github.com/brano-hozza/reserver-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
)

// CreateExamination - Create new examination
func (this *implExaminationReservationAPI) CreateExamination(ctx *gin.Context) {
	value, exists := ctx.Get("examination_service")
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

	db, ok := value.(db_service.DbService[Examination])
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

	examination := Examination{}
	err := ctx.BindJSON(&examination)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Error parsing request",
				"error":   err.Error(),
			})
		return
	}

	err = db.CreateDocument(ctx, examination.Id, &examination)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Error creating examination",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, examination)
}

// DeleteExamination - Delete examination
func (this *implExaminationReservationAPI) DeleteExamination(ctx *gin.Context) {
	value, exists := ctx.Get("examination_service")
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

	db, ok := value.(db_service.DbService[Examination])
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
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Error deleting examination",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// GetExaminationById - Provides the examination by id
func (this *implExaminationReservationAPI) GetExaminationById(ctx *gin.Context) {
	value, exists := ctx.Get("examination_service")
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

	db, ok := value.(db_service.DbService[Examination])
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
	examination, err := db.FindDocument(ctx, id)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Examination not found",
					"error":   err.Error(),
				},
			)
		default:
			ctx.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to get examination from database",
					"error":   err.Error(),
				},
			)
		}
		return
	}

	ctx.JSON(http.StatusOK, examination)
}

// GetExaminations - Provides the list of current examinations
func (this *implExaminationReservationAPI) GetExaminations(ctx *gin.Context) {
	value, exists := ctx.Get("examination_service")
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

	db, ok := value.(db_service.DbService[Examination])
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

	examinations, err := db.FindAllDocuments(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Error fetching examinations",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, examinations)
}

// UpdateExamination - Update examination
func (this *implExaminationReservationAPI) UpdateExamination(ctx *gin.Context) {
	value, exists := ctx.Get("examination_service")
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

	db, ok := value.(db_service.DbService[Examination])
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

	examination := Examination{}
	err := ctx.BindJSON(&examination)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Error parsing request",
				"error":   err.Error(),
			})
		return
	}

	err = db.UpdateDocument(ctx, examination.Id, &examination)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Error updating examination",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, examination)
}
