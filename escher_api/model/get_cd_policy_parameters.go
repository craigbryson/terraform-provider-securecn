// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCdPolicyParams creates a new GetCdPolicyParams object
// with the default values initialized.
func NewGetCdPolicyParams() *GetCdPolicyParams {

	return &GetCdPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCdPolicyParamsWithTimeout creates a new GetCdPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCdPolicyParamsWithTimeout(timeout time.Duration) *GetCdPolicyParams {

	return &GetCdPolicyParams{

		timeout: timeout,
	}
}

// NewGetCdPolicyParamsWithContext creates a new GetCdPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCdPolicyParamsWithContext(ctx context.Context) *GetCdPolicyParams {

	return &GetCdPolicyParams{

		Context: ctx,
	}
}

// NewGetCdPolicyParamsWithHTTPClient creates a new GetCdPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCdPolicyParamsWithHTTPClient(client *http.Client) *GetCdPolicyParams {

	return &GetCdPolicyParams{
		HTTPClient: client,
	}
}

/*GetCdPolicyParams contains all the parameters to send to the API endpoint
for the get cd policy operation typically these are written to a http.Request
*/
type GetCdPolicyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cd policy params
func (o *GetCdPolicyParams) WithTimeout(timeout time.Duration) *GetCdPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cd policy params
func (o *GetCdPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cd policy params
func (o *GetCdPolicyParams) WithContext(ctx context.Context) *GetCdPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cd policy params
func (o *GetCdPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cd policy params
func (o *GetCdPolicyParams) WithHTTPClient(client *http.Client) *GetCdPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cd policy params
func (o *GetCdPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCdPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}