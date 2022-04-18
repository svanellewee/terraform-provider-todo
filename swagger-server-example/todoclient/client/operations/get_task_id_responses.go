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

// GetTaskIDReader is a Reader for the GetTaskID structure.
type GetTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTaskIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTaskIDOK creates a GetTaskIDOK with default headers values
func NewGetTaskIDOK() *GetTaskIDOK {
	return &GetTaskIDOK{}
}

/* GetTaskIDOK describes a response with status code 200, with default header values.

got one task
*/
type GetTaskIDOK struct {
	Payload *models.Task
}

func (o *GetTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /task/{id}][%d] getTaskIdOK  %+v", 200, o.Payload)
}
func (o *GetTaskIDOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GetTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskIDDefault creates a GetTaskIDDefault with default headers values
func NewGetTaskIDDefault(code int) *GetTaskIDDefault {
	return &GetTaskIDDefault{
		_statusCode: code,
	}
}

/* GetTaskIDDefault describes a response with status code -1, with default header values.

error occurred getting one task
*/
type GetTaskIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get task ID default response
func (o *GetTaskIDDefault) Code() int {
	return o._statusCode
}

func (o *GetTaskIDDefault) Error() string {
	return fmt.Sprintf("[GET /task/{id}][%d] GetTaskID default  %+v", o._statusCode, o.Payload)
}
func (o *GetTaskIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTaskIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}