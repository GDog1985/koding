package j_group

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

// NewPostRemoteAPIJGroupKickMemberIDParams creates a new PostRemoteAPIJGroupKickMemberIDParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupKickMemberIDParams() *PostRemoteAPIJGroupKickMemberIDParams {
	var ()
	return &PostRemoteAPIJGroupKickMemberIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupKickMemberIDParamsWithTimeout creates a new PostRemoteAPIJGroupKickMemberIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupKickMemberIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupKickMemberIDParams {
	var ()
	return &PostRemoteAPIJGroupKickMemberIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupKickMemberIDParamsWithContext creates a new PostRemoteAPIJGroupKickMemberIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupKickMemberIDParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupKickMemberIDParams {
	var ()
	return &PostRemoteAPIJGroupKickMemberIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupKickMemberIDParams contains all the parameters to send to the API endpoint
for the post remote API j group kick member ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupKickMemberIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j group kick member ID params
func (o *PostRemoteAPIJGroupKickMemberIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupKickMemberIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group kick member ID params
func (o *PostRemoteAPIJGroupKickMemberIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group kick member ID params
func (o *PostRemoteAPIJGroupKickMemberIDParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupKickMemberIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group kick member ID params
func (o *PostRemoteAPIJGroupKickMemberIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j group kick member ID params
func (o *PostRemoteAPIJGroupKickMemberIDParams) WithID(id string) *PostRemoteAPIJGroupKickMemberIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j group kick member ID params
func (o *PostRemoteAPIJGroupKickMemberIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupKickMemberIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
