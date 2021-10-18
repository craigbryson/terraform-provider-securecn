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

// NewPostCdDeploymentRuleParams creates a new PostCdDeploymentRuleParams object
// with the default values initialized.
func NewPostCdDeploymentRuleParams() *PostCdDeploymentRuleParams {
	var ()
	return &PostCdDeploymentRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCdDeploymentRuleParamsWithTimeout creates a new PostCdDeploymentRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCdDeploymentRuleParamsWithTimeout(timeout time.Duration) *PostCdDeploymentRuleParams {
	var ()
	return &PostCdDeploymentRuleParams{

		timeout: timeout,
	}
}

// NewPostCdDeploymentRuleParamsWithContext creates a new PostCdDeploymentRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCdDeploymentRuleParamsWithContext(ctx context.Context) *PostCdDeploymentRuleParams {
	var ()
	return &PostCdDeploymentRuleParams{

		Context: ctx,
	}
}

// NewPostCdDeploymentRuleParamsWithHTTPClient creates a new PostCdDeploymentRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCdDeploymentRuleParamsWithHTTPClient(client *http.Client) *PostCdDeploymentRuleParams {
	var ()
	return &PostCdDeploymentRuleParams{
		HTTPClient: client,
	}
}

/*PostCdDeploymentRuleParams contains all the parameters to send to the API endpoint
for the post cd deployment rule operation typically these are written to a http.Request
*/
type PostCdDeploymentRuleParams struct {

	/*Body*/
	Body *CdAppRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cd deployment rule params
func (o *PostCdDeploymentRuleParams) WithTimeout(timeout time.Duration) *PostCdDeploymentRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cd deployment rule params
func (o *PostCdDeploymentRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cd deployment rule params
func (o *PostCdDeploymentRuleParams) WithContext(ctx context.Context) *PostCdDeploymentRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cd deployment rule params
func (o *PostCdDeploymentRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cd deployment rule params
func (o *PostCdDeploymentRuleParams) WithHTTPClient(client *http.Client) *PostCdDeploymentRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cd deployment rule params
func (o *PostCdDeploymentRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post cd deployment rule params
func (o *PostCdDeploymentRuleParams) WithBody(body *CdAppRule) *PostCdDeploymentRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post cd deployment rule params
func (o *PostCdDeploymentRuleParams) SetBody(body *CdAppRule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCdDeploymentRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}