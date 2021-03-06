// Code generated by go-swagger; DO NOT EDIT.

package subscribers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// DeleteLTENetworkIDMsisdnsMsisdnReader is a Reader for the DeleteLTENetworkIDMsisdnsMsisdn structure.
type DeleteLTENetworkIDMsisdnsMsisdnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLTENetworkIDMsisdnsMsisdnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLTENetworkIDMsisdnsMsisdnNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteLTENetworkIDMsisdnsMsisdnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLTENetworkIDMsisdnsMsisdnNoContent creates a DeleteLTENetworkIDMsisdnsMsisdnNoContent with default headers values
func NewDeleteLTENetworkIDMsisdnsMsisdnNoContent() *DeleteLTENetworkIDMsisdnsMsisdnNoContent {
	return &DeleteLTENetworkIDMsisdnsMsisdnNoContent{}
}

/*DeleteLTENetworkIDMsisdnsMsisdnNoContent handles this case with default header values.

Success
*/
type DeleteLTENetworkIDMsisdnsMsisdnNoContent struct {
}

func (o *DeleteLTENetworkIDMsisdnsMsisdnNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/msisdns/{msisdn}][%d] deleteLteNetworkIdMsisdnsMsisdnNoContent ", 204)
}

func (o *DeleteLTENetworkIDMsisdnsMsisdnNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLTENetworkIDMsisdnsMsisdnDefault creates a DeleteLTENetworkIDMsisdnsMsisdnDefault with default headers values
func NewDeleteLTENetworkIDMsisdnsMsisdnDefault(code int) *DeleteLTENetworkIDMsisdnsMsisdnDefault {
	return &DeleteLTENetworkIDMsisdnsMsisdnDefault{
		_statusCode: code,
	}
}

/*DeleteLTENetworkIDMsisdnsMsisdnDefault handles this case with default header values.

Unexpected Error
*/
type DeleteLTENetworkIDMsisdnsMsisdnDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete LTE network ID msisdns msisdn default response
func (o *DeleteLTENetworkIDMsisdnsMsisdnDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLTENetworkIDMsisdnsMsisdnDefault) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/msisdns/{msisdn}][%d] DeleteLTENetworkIDMsisdnsMsisdn default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLTENetworkIDMsisdnsMsisdnDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLTENetworkIDMsisdnsMsisdnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
