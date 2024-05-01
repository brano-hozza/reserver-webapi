package reserver

import (
	"net/http"

	"github.com/brano-hozza/reserver-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
)

// GetDepartments - Provide the list of all departments
func (this *implDepartmentsAPI) GetDepartments(ctx *gin.Context) {
	value, exists := ctx.Get("department_service")
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

	db, ok := value.(db_service.DbService[Department])
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

	departments, err := db.FindAllDocuments(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Error fetching departments",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, departments)
}

// GetDoctors - Provide list of all doctors
func (this *implDepartmentsAPI) GetDoctors(ctx *gin.Context) {
	value, exists := ctx.Get("doctor_service")
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

	db, ok := value.(db_service.DbService[Doctor])
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

	doctors, err := db.FindAllDocuments(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Error fetching doctors",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, doctors)
}
