// Code generated by go-swagger; DO NOT EDIT.

// Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
// Copyright (C) 2020 Panther Labs Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

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

	"github.com/panther-labs/panther/api/gateway/remediation/models"
)

// NewRemediateResourceAsyncParams creates a new RemediateResourceAsyncParams object
// with the default values initialized.
func NewRemediateResourceAsyncParams() *RemediateResourceAsyncParams {
	var ()
	return &RemediateResourceAsyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemediateResourceAsyncParamsWithTimeout creates a new RemediateResourceAsyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemediateResourceAsyncParamsWithTimeout(timeout time.Duration) *RemediateResourceAsyncParams {
	var ()
	return &RemediateResourceAsyncParams{

		timeout: timeout,
	}
}

// NewRemediateResourceAsyncParamsWithContext creates a new RemediateResourceAsyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemediateResourceAsyncParamsWithContext(ctx context.Context) *RemediateResourceAsyncParams {
	var ()
	return &RemediateResourceAsyncParams{

		Context: ctx,
	}
}

// NewRemediateResourceAsyncParamsWithHTTPClient creates a new RemediateResourceAsyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemediateResourceAsyncParamsWithHTTPClient(client *http.Client) *RemediateResourceAsyncParams {
	var ()
	return &RemediateResourceAsyncParams{
		HTTPClient: client,
	}
}

/*RemediateResourceAsyncParams contains all the parameters to send to the API endpoint
for the remediate resource async operation typically these are written to a http.Request
*/
type RemediateResourceAsyncParams struct {

	/*Body*/
	Body *models.RemediateResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remediate resource async params
func (o *RemediateResourceAsyncParams) WithTimeout(timeout time.Duration) *RemediateResourceAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remediate resource async params
func (o *RemediateResourceAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remediate resource async params
func (o *RemediateResourceAsyncParams) WithContext(ctx context.Context) *RemediateResourceAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remediate resource async params
func (o *RemediateResourceAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remediate resource async params
func (o *RemediateResourceAsyncParams) WithHTTPClient(client *http.Client) *RemediateResourceAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remediate resource async params
func (o *RemediateResourceAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remediate resource async params
func (o *RemediateResourceAsyncParams) WithBody(body *models.RemediateResource) *RemediateResourceAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remediate resource async params
func (o *RemediateResourceAsyncParams) SetBody(body *models.RemediateResource) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemediateResourceAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
