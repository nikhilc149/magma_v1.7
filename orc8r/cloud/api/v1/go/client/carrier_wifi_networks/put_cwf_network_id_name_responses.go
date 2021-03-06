// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutCwfNetworkIDNameReader is a Reader for the PutCwfNetworkIDName structure.
type PutCwfNetworkIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDNameNoContent creates a PutCwfNetworkIDNameNoContent with default headers values
func NewPutCwfNetworkIDNameNoContent() *PutCwfNetworkIDNameNoContent {
	return &PutCwfNetworkIDNameNoContent{}
}

/*PutCwfNetworkIDNameNoContent handles this case with default header values.

Success
*/
type PutCwfNetworkIDNameNoContent struct {
}

func (o *PutCwfNetworkIDNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/name][%d] putCwfNetworkIdNameNoContent ", 204)
}

func (o *PutCwfNetworkIDNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDNameDefault creates a PutCwfNetworkIDNameDefault with default headers values
func NewPutCwfNetworkIDNameDefault(code int) *PutCwfNetworkIDNameDefault {
	return &PutCwfNetworkIDNameDefault{
		_statusCode: code,
	}
}

/*PutCwfNetworkIDNameDefault handles this case with default header values.

Unexpected Error
*/
type PutCwfNetworkIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID name default response
func (o *PutCwfNetworkIDNameDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDNameDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/name][%d] PutCwfNetworkIDName default  %+v", o._statusCode, o.Payload)
}

func (o *PutCwfNetworkIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
