// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/hongkailiu/test-go/swagger/swagger/models"
)

// GetUserUserIDOKCode is the HTTP code returned for type GetUserUserIDOK
const GetUserUserIDOKCode int = 200

/*GetUserUserIDOK A user object.

swagger:response getUserUserIdOK
*/
type GetUserUserIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserUserIDOK creates GetUserUserIDOK with default headers values
func NewGetUserUserIDOK() *GetUserUserIDOK {

	return &GetUserUserIDOK{}
}

// WithPayload adds the payload to the get user user Id o k response
func (o *GetUserUserIDOK) WithPayload(payload *models.User) *GetUserUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user user Id o k response
func (o *GetUserUserIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserUserIDBadRequestCode is the HTTP code returned for type GetUserUserIDBadRequest
const GetUserUserIDBadRequestCode int = 400

/*GetUserUserIDBadRequest The specified user ID is invalid (not a number).

swagger:response getUserUserIdBadRequest
*/
type GetUserUserIDBadRequest struct {
}

// NewGetUserUserIDBadRequest creates GetUserUserIDBadRequest with default headers values
func NewGetUserUserIDBadRequest() *GetUserUserIDBadRequest {

	return &GetUserUserIDBadRequest{}
}

// WriteResponse to the client
func (o *GetUserUserIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetUserUserIDNotFoundCode is the HTTP code returned for type GetUserUserIDNotFound
const GetUserUserIDNotFoundCode int = 404

/*GetUserUserIDNotFound A user with the specified ID was not found.

swagger:response getUserUserIdNotFound
*/
type GetUserUserIDNotFound struct {
}

// NewGetUserUserIDNotFound creates GetUserUserIDNotFound with default headers values
func NewGetUserUserIDNotFound() *GetUserUserIDNotFound {

	return &GetUserUserIDNotFound{}
}

// WriteResponse to the client
func (o *GetUserUserIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*GetUserUserIDDefault Unexpected error

swagger:response getUserUserIdDefault
*/
type GetUserUserIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserUserIDDefault creates GetUserUserIDDefault with default headers values
func NewGetUserUserIDDefault(code int) *GetUserUserIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserUserIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user user ID default response
func (o *GetUserUserIDDefault) WithStatusCode(code int) *GetUserUserIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user user ID default response
func (o *GetUserUserIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user user ID default response
func (o *GetUserUserIDDefault) WithPayload(payload *models.Error) *GetUserUserIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user user ID default response
func (o *GetUserUserIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserUserIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
