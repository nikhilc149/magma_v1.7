// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDReader is a Reader for the PostFegLTENetworkIDSubscriberConfigRuleNamesRuleID structure.
type PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated creates a PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated with default headers values
func NewPostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated() *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated {
	return &PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated{}
}

/*PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated handles this case with default header values.

Success
*/
type PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated struct {
}

func (o *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated) Error() string {
	return fmt.Sprintf("[POST /feg_lte/{network_id}/subscriber_config/rule_names/{rule_id}][%d] postFegLteNetworkIdSubscriberConfigRuleNamesRuleIdCreated ", 201)
}

func (o *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault creates a PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault with default headers values
func NewPostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault(code int) *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault {
	return &PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault{
		_statusCode: code,
	}
}

/*PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault handles this case with default header values.

Unexpected Error
*/
type PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post feg LTE network ID subscriber config rule names rule ID default response
func (o *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault) Error() string {
	return fmt.Sprintf("[POST /feg_lte/{network_id}/subscriber_config/rule_names/{rule_id}][%d] PostFegLTENetworkIDSubscriberConfigRuleNamesRuleID default  %+v", o._statusCode, o.Payload)
}

func (o *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
