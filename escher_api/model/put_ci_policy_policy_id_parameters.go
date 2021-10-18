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

// NewPutCiPolicyPolicyIDParams creates a new PutCiPolicyPolicyIDParams object
// with the default values initialized.
func NewPutCiPolicyPolicyIDParams() *PutCiPolicyPolicyIDParams {
	var ()
	return &PutCiPolicyPolicyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCiPolicyPolicyIDParamsWithTimeout creates a new PutCiPolicyPolicyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCiPolicyPolicyIDParamsWithTimeout(timeout time.Duration) *PutCiPolicyPolicyIDParams {
	var ()
	return &PutCiPolicyPolicyIDParams{

		timeout: timeout,
	}
}

// NewPutCiPolicyPolicyIDParamsWithContext creates a new PutCiPolicyPolicyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCiPolicyPolicyIDParamsWithContext(ctx context.Context) *PutCiPolicyPolicyIDParams {
	var ()
	return &PutCiPolicyPolicyIDParams{

		Context: ctx,
	}
}

// NewPutCiPolicyPolicyIDParamsWithHTTPClient creates a new PutCiPolicyPolicyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCiPolicyPolicyIDParamsWithHTTPClient(client *http.Client) *PutCiPolicyPolicyIDParams {
	var ()
	return &PutCiPolicyPolicyIDParams{
		HTTPClient: client,
	}
}

/*PutCiPolicyPolicyIDParams contains all the parameters to send to the API endpoint
for the put ci policy policy ID operation typically these are written to a http.Request
*/
type PutCiPolicyPolicyIDParams struct {

	/*Body*/
	Body *CiPolicy
	/*PolicyID*/
	PolicyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) WithTimeout(timeout time.Duration) *PutCiPolicyPolicyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) WithContext(ctx context.Context) *PutCiPolicyPolicyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) WithHTTPClient(client *http.Client) *PutCiPolicyPolicyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) WithBody(body *CiPolicy) *PutCiPolicyPolicyIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) SetBody(body *CiPolicy) {
	o.Body = body
}

// WithPolicyID adds the policyID to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) WithPolicyID(policyID strfmt.UUID) *PutCiPolicyPolicyIDParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the put ci policy policy ID params
func (o *PutCiPolicyPolicyIDParams) SetPolicyID(policyID strfmt.UUID) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCiPolicyPolicyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param policyId
	if err := r.SetPathParam("policyId", o.PolicyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}