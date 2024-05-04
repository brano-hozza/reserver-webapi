/*
 * Reserver Api
 *
 * Room and ambulance reservation management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xhozza@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package reserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExaminationReservationAPI interface {

	// internal registration of api routes
	addRoutes(routerGroup *gin.RouterGroup)

	// CreateExamination - Create new examination
	CreateExamination(ctx *gin.Context)

	// DeleteExamination - Delete examination
	DeleteExamination(ctx *gin.Context)

	// GetExaminationById - Provides the examination by id
	GetExaminationById(ctx *gin.Context)

	// GetExaminations - Provides the list of current examinations
	GetExaminations(ctx *gin.Context)

	// UpdateExamination - Update examination
	UpdateExamination(ctx *gin.Context)
}

// partial implementation of ExaminationReservationAPI - all functions must be implemented in add on files
type implExaminationReservationAPI struct {
}

func newExaminationReservationAPI() ExaminationReservationAPI {
	return &implExaminationReservationAPI{}
}

func (this *implExaminationReservationAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodPost, "/examination", this.CreateExamination)
	routerGroup.Handle(http.MethodDelete, "/examination/:id", this.DeleteExamination)
	routerGroup.Handle(http.MethodGet, "/examination/:id", this.GetExaminationById)
	routerGroup.Handle(http.MethodGet, "/examination", this.GetExaminations)
	routerGroup.Handle(http.MethodPut, "/examination/:id", this.UpdateExamination)
}
