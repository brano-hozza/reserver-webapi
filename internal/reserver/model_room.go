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

type Room struct {

	// Unique identifier of the room
	Id string `json:"id"`

	// Room number
	RoomNumber string `json:"roomNumber"`
}
