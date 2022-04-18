// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/svanellewee/todoserver/models"
)

// PostTaskCreatedCode is the HTTP code returned for type PostTaskCreated
const PostTaskCreatedCode int = 201

/*PostTaskCreated created one

swagger:response postTaskCreated
*/
type PostTaskCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Task `json:"body,omitempty"`
}

// NewPostTaskCreated creates PostTaskCreated with default headers values
func NewPostTaskCreated() *PostTaskCreated {

	return &PostTaskCreated{}
}

// WithPayload adds the payload to the post task created response
func (o *PostTaskCreated) WithPayload(payload *models.Task) *PostTaskCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post task created response
func (o *PostTaskCreated) SetPayload(payload *models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTaskCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostTaskDefault could not create

swagger:response postTaskDefault
*/
type PostTaskDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTaskDefault creates PostTaskDefault with default headers values
func NewPostTaskDefault(code int) *PostTaskDefault {
	if code <= 0 {
		code = 500
	}

	return &PostTaskDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post task default response
func (o *PostTaskDefault) WithStatusCode(code int) *PostTaskDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post task default response
func (o *PostTaskDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post task default response
func (o *PostTaskDefault) WithPayload(payload *models.Error) *PostTaskDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post task default response
func (o *PostTaskDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTaskDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
