// Code generated by go-swagger; DO NOT EDIT.

package subscribers

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

// NewDeleteLTENetworkIDSubscribersSubscriberIDParams creates a new DeleteLTENetworkIDSubscribersSubscriberIDParams object
// with the default values initialized.
func NewDeleteLTENetworkIDSubscribersSubscriberIDParams() *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	var ()
	return &DeleteLTENetworkIDSubscribersSubscriberIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLTENetworkIDSubscribersSubscriberIDParamsWithTimeout creates a new DeleteLTENetworkIDSubscribersSubscriberIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLTENetworkIDSubscribersSubscriberIDParamsWithTimeout(timeout time.Duration) *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	var ()
	return &DeleteLTENetworkIDSubscribersSubscriberIDParams{

		timeout: timeout,
	}
}

// NewDeleteLTENetworkIDSubscribersSubscriberIDParamsWithContext creates a new DeleteLTENetworkIDSubscribersSubscriberIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLTENetworkIDSubscribersSubscriberIDParamsWithContext(ctx context.Context) *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	var ()
	return &DeleteLTENetworkIDSubscribersSubscriberIDParams{

		Context: ctx,
	}
}

// NewDeleteLTENetworkIDSubscribersSubscriberIDParamsWithHTTPClient creates a new DeleteLTENetworkIDSubscribersSubscriberIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLTENetworkIDSubscribersSubscriberIDParamsWithHTTPClient(client *http.Client) *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	var ()
	return &DeleteLTENetworkIDSubscribersSubscriberIDParams{
		HTTPClient: client,
	}
}

/*DeleteLTENetworkIDSubscribersSubscriberIDParams contains all the parameters to send to the API endpoint
for the delete LTE network ID subscribers subscriber ID operation typically these are written to a http.Request
*/
type DeleteLTENetworkIDSubscribersSubscriberIDParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*SubscriberID
	  Subscriber ID

	*/
	SubscriberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) WithTimeout(timeout time.Duration) *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) WithContext(ctx context.Context) *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) WithHTTPClient(client *http.Client) *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) WithNetworkID(networkID string) *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSubscriberID adds the subscriberID to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) WithSubscriberID(subscriberID string) *DeleteLTENetworkIDSubscribersSubscriberIDParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the delete LTE network ID subscribers subscriber ID params
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLTENetworkIDSubscribersSubscriberIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param subscriber_id
	if err := r.SetPathParam("subscriber_id", o.SubscriberID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}