// Package TS29518NamfMT provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29518NamfMT

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29518NamfEventExposure"
	externalRef1 "magma/feg/gateway/sbi/specs/TS29571CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// EnableUeReachabilityReqData defines model for EnableUeReachabilityReqData.
type EnableUeReachabilityReqData struct {
	OldGuami          *externalRef1.Guami             `json:"oldGuami,omitempty"`
	Reachability      externalRef0.UeReachability     `json:"reachability"`
	SupportedFeatures *externalRef1.SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// EnableUeReachabilityRspData defines model for EnableUeReachabilityRspData.
type EnableUeReachabilityRspData struct {
	Reachability      externalRef0.UeReachability     `json:"reachability"`
	SupportedFeatures *externalRef1.SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// UeContextInfo defines model for UeContextInfo.
type UeContextInfo struct {
	AccessType        *externalRef1.AccessType        `json:"accessType,omitempty"`
	LastActTime       *externalRef1.DateTime          `json:"lastActTime,omitempty"`
	RatType           *externalRef1.RatType           `json:"ratType,omitempty"`
	SupportVoPS       *bool                           `json:"supportVoPS,omitempty"`
	SupportVoPSn3gpp  *bool                           `json:"supportVoPSn3gpp,omitempty"`
	SupportedFeatures *externalRef1.SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// UeContextInfoClass defines model for UeContextInfoClass.
type UeContextInfoClass interface{}

// ProvideDomainSelectionInfoParams defines parameters for ProvideDomainSelectionInfo.
type ProvideDomainSelectionInfoParams struct {
	// UE Context Information Class
	InfoClass *UeContextInfoClass `json:"info-class,omitempty"`

	// Supported Features
	SupportedFeatures *externalRef1.SupportedFeatures `json:"supported-features,omitempty"`

	// Old GUAMI
	OldGuami *externalRef1.Guami `json:"old-guami,omitempty"`
}

// EnableUeReachabilityJSONBody defines parameters for EnableUeReachability.
type EnableUeReachabilityJSONBody EnableUeReachabilityReqData

// EnableUeReachabilityJSONRequestBody defines body for EnableUeReachability for application/json ContentType.
type EnableUeReachabilityJSONRequestBody EnableUeReachabilityJSONBody
