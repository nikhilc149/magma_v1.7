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

// GetLTENetworkIDPolicyQosProfilesProfileIDReader is a Reader for the GetLTENetworkIDPolicyQosProfilesProfileID structure.
type GetLTENetworkIDPolicyQosProfilesProfileIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDPolicyQosProfilesProfileIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDPolicyQosProfilesProfileIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDPolicyQosProfilesProfileIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDPolicyQosProfilesProfileIDOK creates a GetLTENetworkIDPolicyQosProfilesProfileIDOK with default headers values
func NewGetLTENetworkIDPolicyQosProfilesProfileIDOK() *GetLTENetworkIDPolicyQosProfilesProfileIDOK {
	return &GetLTENetworkIDPolicyQosProfilesProfileIDOK{}
}

/*GetLTENetworkIDPolicyQosProfilesProfileIDOK handles this case with default header values.

Policy QoS profie
*/
type GetLTENetworkIDPolicyQosProfilesProfileIDOK struct {
	Payload *models.PolicyQosProfile
}

func (o *GetLTENetworkIDPolicyQosProfilesProfileIDOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/policy_qos_profiles/{profile_id}][%d] getLteNetworkIdPolicyQosProfilesProfileIdOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDPolicyQosProfilesProfileIDOK) GetPayload() *models.PolicyQosProfile {
	return o.Payload
}

func (o *GetLTENetworkIDPolicyQosProfilesProfileIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyQosProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDPolicyQosProfilesProfileIDDefault creates a GetLTENetworkIDPolicyQosProfilesProfileIDDefault with default headers values
func NewGetLTENetworkIDPolicyQosProfilesProfileIDDefault(code int) *GetLTENetworkIDPolicyQosProfilesProfileIDDefault {
	return &GetLTENetworkIDPolicyQosProfilesProfileIDDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDPolicyQosProfilesProfileIDDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDPolicyQosProfilesProfileIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID policy qos profiles profile ID default response
func (o *GetLTENetworkIDPolicyQosProfilesProfileIDDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDPolicyQosProfilesProfileIDDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/policy_qos_profiles/{profile_id}][%d] GetLTENetworkIDPolicyQosProfilesProfileID default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDPolicyQosProfilesProfileIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDPolicyQosProfilesProfileIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
