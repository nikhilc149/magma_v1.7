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

// PutLTENetworkIDGatewaysGatewayIDCellularRanReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDCellularRan structure.
type PutLTENetworkIDGatewaysGatewayIDCellularRanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDCellularRanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularRanNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularRanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularRanNoContent creates a PutLTENetworkIDGatewaysGatewayIDCellularRanNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularRanNoContent() *PutLTENetworkIDGatewaysGatewayIDCellularRanNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDCellularRanNoContent{}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularRanNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDCellularRanNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularRanNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/ran][%d] putLteNetworkIdGatewaysGatewayIdCellularRanNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularRanNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularRanDefault creates a PutLTENetworkIDGatewaysGatewayIDCellularRanDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularRanDefault(code int) *PutLTENetworkIDGatewaysGatewayIDCellularRanDefault {
	return &PutLTENetworkIDGatewaysGatewayIDCellularRanDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularRanDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDCellularRanDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID cellular ran default response
func (o *PutLTENetworkIDGatewaysGatewayIDCellularRanDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularRanDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/ran][%d] PutLTENetworkIDGatewaysGatewayIDCellularRan default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularRanDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularRanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
