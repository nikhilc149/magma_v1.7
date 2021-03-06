// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

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

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutLTENetworkIDGatewaysGatewayIDMagmadParams creates a new PutLTENetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized.
func NewPutLTENetworkIDGatewaysGatewayIDMagmadParams() *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDMagmadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDMagmadParamsWithTimeout creates a new PutLTENetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDGatewaysGatewayIDMagmadParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDMagmadParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDMagmadParamsWithContext creates a new PutLTENetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDGatewaysGatewayIDMagmadParamsWithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDMagmadParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDMagmadParamsWithHTTPClient creates a new PutLTENetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDGatewaysGatewayIDMagmadParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDMagmadParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDMagmadParams contains all the parameters to send to the API endpoint
for the put LTE network ID gateways gateway ID magmad operation typically these are written to a http.Request
*/
type PutLTENetworkIDGatewaysGatewayIDMagmadParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*Magmad
	  New magmad configuration

	*/
	Magmad *models.MagmadGatewayConfigs
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) WithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) WithGatewayID(gatewayID string) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithMagmad adds the magmad to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) WithMagmad(magmad *models.MagmadGatewayConfigs) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	o.SetMagmad(magmad)
	return o
}

// SetMagmad adds the magmad to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) SetMagmad(magmad *models.MagmadGatewayConfigs) {
	o.Magmad = magmad
}

// WithNetworkID adds the networkID to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) WithNetworkID(networkID string) *PutLTENetworkIDGatewaysGatewayIDMagmadParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID gateways gateway ID magmad params
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	if o.Magmad != nil {
		if err := r.SetBodyParam(o.Magmad); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
