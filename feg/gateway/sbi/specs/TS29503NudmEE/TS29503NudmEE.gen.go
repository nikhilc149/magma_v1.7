// Package TS29503NudmEE provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29503NudmEE

import (
	"encoding/json"
	"fmt"

	externalRef0 "magma/feg/gateway/sbi/specs/TS29571CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// AssociationType defines model for AssociationType.
type AssociationType interface{}

// ChangeOfSupiPeiAssociationReport defines model for ChangeOfSupiPeiAssociationReport.
type ChangeOfSupiPeiAssociationReport struct {
	NewPei externalRef0.Pei `json:"newPei"`
}

// CreatedEeSubscription defines model for CreatedEeSubscription.
type CreatedEeSubscription struct {
	EeSubscription EeSubscription         `json:"eeSubscription"`
	EventReports   *[]MonitoringReport    `json:"eventReports,omitempty"`
	NumberOfUes    *externalRef0.Uinteger `json:"numberOfUes,omitempty"`
}

// EeSubscription defines model for EeSubscription.
type EeSubscription struct {
	CallbackReference externalRef0.Uri `json:"callbackReference"`

	// A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations
	MonitoringConfigurations EeSubscription_MonitoringConfigurations `json:"monitoringConfigurations"`
	ReportingOptions         *ReportingOptions                       `json:"reportingOptions,omitempty"`
	SubscriptionId           *string                                 `json:"subscriptionId,omitempty"`
	SupportedFeatures        *externalRef0.SupportedFeatures         `json:"supportedFeatures,omitempty"`
}

// A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations
type EeSubscription_MonitoringConfigurations struct {
	AdditionalProperties map[string]MonitoringConfiguration `json:"-"`
}

// EventType defines model for EventType.
type EventType interface{}

// LocationAccuracy defines model for LocationAccuracy.
type LocationAccuracy interface{}

// LocationReportingConfiguration defines model for LocationReportingConfiguration.
type LocationReportingConfiguration struct {
	Accuracy        *LocationAccuracy `json:"accuracy,omitempty"`
	CurrentLocation bool              `json:"currentLocation"`
	OneTime         *bool             `json:"oneTime,omitempty"`
}

// MaxNumOfReports defines model for MaxNumOfReports.
type MaxNumOfReports int

// MonitoringConfiguration defines model for MonitoringConfiguration.
type MonitoringConfiguration struct {
	AssociationType                *AssociationType                `json:"associationType,omitempty"`
	EventType                      EventType                       `json:"eventType"`
	ImmediateFlag                  *bool                           `json:"immediateFlag,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration `json:"locationReportingConfiguration,omitempty"`
}

// MonitoringReport defines model for MonitoringReport.
type MonitoringReport struct {
	EventType   EventType             `json:"eventType"`
	Gpsi        *externalRef0.Gpsi    `json:"gpsi,omitempty"`
	ReferenceId ReferenceId           `json:"referenceId"`
	Report      *Report               `json:"report,omitempty"`
	TimeStamp   externalRef0.DateTime `json:"timeStamp"`
}

// ReferenceId defines model for ReferenceId.
type ReferenceId int

// Report defines model for Report.
type Report interface{}

// ReportingOptions defines model for ReportingOptions.
type ReportingOptions struct {
	Expiry          *externalRef0.DateTime `json:"expiry,omitempty"`
	MaxNumOfReports *MaxNumOfReports       `json:"maxNumOfReports,omitempty"`
}

// RoamingStatusReport defines model for RoamingStatusReport.
type RoamingStatusReport struct {
	NewServingPlmn externalRef0.PlmnId `json:"newServingPlmn"`
	Roaming        bool                `json:"roaming"`
}

// CreateEeSubscriptionJSONBody defines parameters for CreateEeSubscription.
type CreateEeSubscriptionJSONBody EeSubscription

// CreateEeSubscriptionJSONRequestBody defines body for CreateEeSubscription for application/json ContentType.
type CreateEeSubscriptionJSONRequestBody CreateEeSubscriptionJSONBody

// Getter for additional properties for EeSubscription_MonitoringConfigurations. Returns the specified
// element and whether it was found
func (a EeSubscription_MonitoringConfigurations) Get(fieldName string) (value MonitoringConfiguration, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for EeSubscription_MonitoringConfigurations
func (a *EeSubscription_MonitoringConfigurations) Set(fieldName string, value MonitoringConfiguration) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]MonitoringConfiguration)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for EeSubscription_MonitoringConfigurations to handle AdditionalProperties
func (a *EeSubscription_MonitoringConfigurations) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]MonitoringConfiguration)
		for fieldName, fieldBuf := range object {
			var fieldVal MonitoringConfiguration
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for EeSubscription_MonitoringConfigurations to handle AdditionalProperties
func (a EeSubscription_MonitoringConfigurations) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
