// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/configuration-service/models"
)

// DeleteProjectProjectNameStageStageNameServiceServiceNameNoContentCode is the HTTP code returned for type DeleteProjectProjectNameStageStageNameServiceServiceNameNoContent
const DeleteProjectProjectNameStageStageNameServiceServiceNameNoContentCode int = 204

/*DeleteProjectProjectNameStageStageNameServiceServiceNameNoContent Success. Service has been deleted. Response does not have a body.

swagger:response deleteProjectProjectNameStageStageNameServiceServiceNameNoContent
*/
type DeleteProjectProjectNameStageStageNameServiceServiceNameNoContent struct {
}

// NewDeleteProjectProjectNameStageStageNameServiceServiceNameNoContent creates DeleteProjectProjectNameStageStageNameServiceServiceNameNoContent with default headers values
func NewDeleteProjectProjectNameStageStageNameServiceServiceNameNoContent() *DeleteProjectProjectNameStageStageNameServiceServiceNameNoContent {

	return &DeleteProjectProjectNameStageStageNameServiceServiceNameNoContent{}
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequestCode is the HTTP code returned for type DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest
const DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequestCode int = 400

/*DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest Failed. Service could not be deleted.

swagger:response deleteProjectProjectNameStageStageNameServiceServiceNameBadRequest
*/
type DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest creates DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest with default headers values
func NewDeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest() *DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest {

	return &DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest{}
}

// WithPayload adds the payload to the delete project project name stage stage name service service name bad request response
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest) WithPayload(payload *models.Error) *DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project project name stage stage name service service name bad request response
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteProjectProjectNameStageStageNameServiceServiceNameDefault Error

swagger:response deleteProjectProjectNameStageStageNameServiceServiceNameDefault
*/
type DeleteProjectProjectNameStageStageNameServiceServiceNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectProjectNameStageStageNameServiceServiceNameDefault creates DeleteProjectProjectNameStageStageNameServiceServiceNameDefault with default headers values
func NewDeleteProjectProjectNameStageStageNameServiceServiceNameDefault(code int) *DeleteProjectProjectNameStageStageNameServiceServiceNameDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteProjectProjectNameStageStageNameServiceServiceNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete project project name stage stage name service service name default response
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameDefault) WithStatusCode(code int) *DeleteProjectProjectNameStageStageNameServiceServiceNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete project project name stage stage name service service name default response
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete project project name stage stage name service service name default response
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameDefault) WithPayload(payload *models.Error) *DeleteProjectProjectNameStageStageNameServiceServiceNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project project name stage stage name service service name default response
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameStageStageNameServiceServiceNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
