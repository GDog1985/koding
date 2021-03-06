package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJAccountIsEmailVerifiedIDParams creates a new PostRemoteAPIJAccountIsEmailVerifiedIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountIsEmailVerifiedIDParams() *PostRemoteAPIJAccountIsEmailVerifiedIDParams {
	var ()
	return &PostRemoteAPIJAccountIsEmailVerifiedIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountIsEmailVerifiedIDParamsWithTimeout creates a new PostRemoteAPIJAccountIsEmailVerifiedIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountIsEmailVerifiedIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountIsEmailVerifiedIDParams {
	var ()
	return &PostRemoteAPIJAccountIsEmailVerifiedIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountIsEmailVerifiedIDParamsWithContext creates a new PostRemoteAPIJAccountIsEmailVerifiedIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountIsEmailVerifiedIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountIsEmailVerifiedIDParams {
	var ()
	return &PostRemoteAPIJAccountIsEmailVerifiedIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountIsEmailVerifiedIDParams contains all the parameters to send to the API endpoint
for the post remote API j account is email verified ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountIsEmailVerifiedIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j account is email verified ID params
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountIsEmailVerifiedIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account is email verified ID params
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account is email verified ID params
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountIsEmailVerifiedIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account is email verified ID params
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j account is email verified ID params
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDParams) WithID(id string) *PostRemoteAPIJAccountIsEmailVerifiedIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account is email verified ID params
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
