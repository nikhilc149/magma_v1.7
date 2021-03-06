// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetTenantsTenantIDMetricsQueryRangeReader is a Reader for the GetTenantsTenantIDMetricsQueryRange structure.
type GetTenantsTenantIDMetricsQueryRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsTenantIDMetricsQueryRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsTenantIDMetricsQueryRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantsTenantIDMetricsQueryRangeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantsTenantIDMetricsQueryRangeOK creates a GetTenantsTenantIDMetricsQueryRangeOK with default headers values
func NewGetTenantsTenantIDMetricsQueryRangeOK() *GetTenantsTenantIDMetricsQueryRangeOK {
	return &GetTenantsTenantIDMetricsQueryRangeOK{}
}

/*GetTenantsTenantIDMetricsQueryRangeOK handles this case with default header values.

List of PromQL metrics results
*/
type GetTenantsTenantIDMetricsQueryRangeOK struct {
	Payload *models.PromqlReturnObject
}

func (o *GetTenantsTenantIDMetricsQueryRangeOK) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}/metrics/query_range][%d] getTenantsTenantIdMetricsQueryRangeOK  %+v", 200, o.Payload)
}

func (o *GetTenantsTenantIDMetricsQueryRangeOK) GetPayload() *models.PromqlReturnObject {
	return o.Payload
}

func (o *GetTenantsTenantIDMetricsQueryRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromqlReturnObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantsTenantIDMetricsQueryRangeDefault creates a GetTenantsTenantIDMetricsQueryRangeDefault with default headers values
func NewGetTenantsTenantIDMetricsQueryRangeDefault(code int) *GetTenantsTenantIDMetricsQueryRangeDefault {
	return &GetTenantsTenantIDMetricsQueryRangeDefault{
		_statusCode: code,
	}
}

/*GetTenantsTenantIDMetricsQueryRangeDefault handles this case with default header values.

Unexpected Error
*/
type GetTenantsTenantIDMetricsQueryRangeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tenants tenant ID metrics query range default response
func (o *GetTenantsTenantIDMetricsQueryRangeDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantsTenantIDMetricsQueryRangeDefault) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}/metrics/query_range][%d] GetTenantsTenantIDMetricsQueryRange default  %+v", o._statusCode, o.Payload)
}

func (o *GetTenantsTenantIDMetricsQueryRangeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTenantsTenantIDMetricsQueryRangeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
