package reserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDepartments - Provide the list of all departments
func (this *implDepartmentsAPI) GetDepartments(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetDoctors - Provide list of all doctors
func (this *implDepartmentsAPI) GetDoctors(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}
