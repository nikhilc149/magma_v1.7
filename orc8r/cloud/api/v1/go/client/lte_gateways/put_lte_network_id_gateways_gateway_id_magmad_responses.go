// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDGatewaysGatewayIDMagmadReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDMagmad structure.
type PutLTENetworkIDGatewaysGatewayIDMagmadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDMagmadNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDMagmadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDMagmadNoContent creates a PutLTENetworkIDGatewaysGatewayIDMagmadNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDMagmadNoContent() *PutLTENetworkIDGatewaysGatewayIDMagmadNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDMagmadNoContent{}
}

/*PutLTENetworkIDGatewaysGatewayIDMagmadNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDMagmadNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDMagmadNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/magmad][%d] putLteNetworkIdGatewaysGatewayIdMagmadNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDMagmadNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDMagmadDefault creates a PutLTENetworkIDGatewaysGatewayIDMagmadDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDMagmadDefault(code int) *PutLTENetworkIDGatewaysGatewayIDMagmadDefault {
	return &PutLTENetworkIDGatewaysGatewayIDMagmadDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDMagmadDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDMagmadDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID magmad default response
func (o *PutLTENetworkIDGatewaysGatewayIDMagmadDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDMagmadDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/magmad][%d] PutLTENetworkIDGatewaysGatewayIDMagmad default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDGatewaysGatewayIDMagmadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDMagmadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
