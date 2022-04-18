// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/svanellewee/todoclient/models"
)

// UpdateOneReader is a Reader for the UpdateOne structure.
type UpdateOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateOneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateOneOK creates a UpdateOneOK with default headers values
func NewUpdateOneOK() *UpdateOneOK {
	return &UpdateOneOK{}
}

/* UpdateOneOK describes a response with status code 200, with default header values.

updated task
*/
type UpdateOneOK struct {
	Payload *models.Task
}

func (o *UpdateOneOK) Error() string {
	return fmt.Sprintf("[PUT /task/{id}][%d] updateOneOK  %+v", 200, o.Payload)
}
func (o *UpdateOneOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *UpdateOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOneDefault creates a UpdateOneDefault with default headers values
func NewUpdateOneDefault(code int) *UpdateOneDefault {
	return &UpdateOneDefault{
		_statusCode: code,
	}
}

/* UpdateOneDefault describes a response with status code -1, with default header values.

could not update task
*/
type UpdateOneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update one default response
func (o *UpdateOneDefault) Code() int {
	return o._statusCode
}

func (o *UpdateOneDefault) Error() string {
	return fmt.Sprintf("[PUT /task/{id}][%d] updateOne default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateOneDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}