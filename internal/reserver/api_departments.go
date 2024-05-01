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

type DepartmentsAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // GetDepartments - Provide the list of all departments
   GetDepartments(ctx *gin.Context)

    // GetDoctors - Provide list of all doctors
   GetDoctors(ctx *gin.Context)

 }

// partial implementation of DepartmentsAPI - all functions must be implemented in add on files
type implDepartmentsAPI struct {

}

func newDepartmentsAPI() DepartmentsAPI {
  return &implDepartmentsAPI{}
}

func (this *implDepartmentsAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodGet, "/departments", this.GetDepartments)
  routerGroup.Handle( http.MethodGet, "/doctors", this.GetDoctors)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // GetDepartments - Provide the list of all departments
// func (this *implDepartmentsAPI) GetDepartments(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetDoctors - Provide list of all doctors
// func (this *implDepartmentsAPI) GetDoctors(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
