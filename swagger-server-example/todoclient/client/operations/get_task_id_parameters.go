// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetTaskIDParams creates a new GetTaskIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTaskIDParams() *GetTaskIDParams {
	return &GetTaskIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskIDParamsWithTimeout creates a new GetTaskIDParams object
// with the ability to set a timeout on a request.
func NewGetTaskIDParamsWithTimeout(timeout time.Duration) *GetTaskIDParams {
	return &GetTaskIDParams{
		timeout: timeout,
	}
}

// NewGetTaskIDParamsWithContext creates a new GetTaskIDParams object
// with the ability to set a context for a request.
func NewGetTaskIDParamsWithContext(ctx context.Context) *GetTaskIDParams {
	return &GetTaskIDParams{
		Context: ctx,
	}
}

// NewGetTaskIDParamsWithHTTPClient creates a new GetTaskIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTaskIDParamsWithHTTPClient(client *http.Client) *GetTaskIDParams {
	return &GetTaskIDParams{
		HTTPClient: client,
	}
}

/* GetTaskIDParams contains all the parameters to send to the API endpoint
   for the get task ID operation.

   Typically these are written to a http.Request.
*/
type GetTaskIDParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get task ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskIDParams) WithDefaults() *GetTaskIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get task ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get task ID params
func (o *GetTaskIDParams) WithTimeout(timeout time.Duration) *GetTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task ID params
func (o *GetTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task ID params
func (o *GetTaskIDParams) WithContext(ctx context.Context) *GetTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task ID params
func (o *GetTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task ID params
func (o *GetTaskIDParams) WithHTTPClient(client *http.Client) *GetTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task ID params
func (o *GetTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get task ID params
func (o *GetTaskIDParams) WithID(id int64) *GetTaskIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get task ID params
func (o *GetTaskIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
