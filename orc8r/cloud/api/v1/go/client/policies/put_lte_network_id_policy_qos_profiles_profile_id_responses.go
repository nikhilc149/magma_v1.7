// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDPolicyQosProfilesProfileIDReader is a Reader for the PutLTENetworkIDPolicyQosProfilesProfileID structure.
type PutLTENetworkIDPolicyQosProfilesProfileIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDPolicyQosProfilesProfileIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDPolicyQosProfilesProfileIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDPolicyQosProfilesProfileIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDPolicyQosProfilesProfileIDNoContent creates a PutLTENetworkIDPolicyQosProfilesProfileIDNoContent with default headers values
func NewPutLTENetworkIDPolicyQosProfilesProfileIDNoContent() *PutLTENetworkIDPolicyQosProfilesProfileIDNoContent {
	return &PutLTENetworkIDPolicyQosProfilesProfileIDNoContent{}
}

/*PutLTENetworkIDPolicyQosProfilesProfileIDNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDPolicyQosProfilesProfileIDNoContent struct {
}

func (o *PutLTENetworkIDPolicyQosProfilesProfileIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/policy_qos_profiles/{profile_id}][%d] putLteNetworkIdPolicyQosProfilesProfileIdNoContent ", 204)
}

func (o *PutLTENetworkIDPolicyQosProfilesProfileIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDPolicyQosProfilesProfileIDDefault creates a PutLTENetworkIDPolicyQosProfilesProfileIDDefault with default headers values
func NewPutLTENetworkIDPolicyQosProfilesProfileIDDefault(code int) *PutLTENetworkIDPolicyQosProfilesProfileIDDefault {
	return &PutLTENetworkIDPolicyQosProfilesProfileIDDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDPolicyQosProfilesProfileIDDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDPolicyQosProfilesProfileIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID policy qos profiles profile ID default response
func (o *PutLTENetworkIDPolicyQosProfilesProfileIDDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDPolicyQosProfilesProfileIDDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/policy_qos_profiles/{profile_id}][%d] PutLTENetworkIDPolicyQosProfilesProfileID default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDPolicyQosProfilesProfileIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDPolicyQosProfilesProfileIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
