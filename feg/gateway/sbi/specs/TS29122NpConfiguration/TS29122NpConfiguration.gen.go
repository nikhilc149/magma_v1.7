// Package TS29122NpConfiguration provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29122NpConfiguration

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29122CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// ConfigurationNotification defines model for ConfigurationNotification.
type ConfigurationNotification struct {
	// The grouping configuration result notification provided by the SCEF.
	ConfigResults *[]externalRef0.ConfigResult `json:"configResults,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Configuration externalRef0.Link `json:"configuration"`
}

// NpConfiguration defines model for NpConfiguration.
type NpConfiguration interface{}

// NpConfigurationPatch defines model for NpConfigurationPatch.
type NpConfigurationPatch struct {
	// Unsigned integer identifying a period of time in units of seconds with "nullable=true" property.
	GroupReportGuardTime *externalRef0.DurationSecRm `json:"groupReportGuardTime"`

	// Unsigned integer identifying a period of time in units of seconds with "nullable=true" property.
	MaximumLatency *externalRef0.DurationSecRm `json:"maximumLatency"`

	// Unsigned integer identifying a period of time in units of seconds with "nullable=true" property.
	MaximumResponseTime *externalRef0.DurationSecRm `json:"maximumResponseTime"`

	// This parameter may be included to identify the number of packets that the serving gateway shall buffer in case that the UE is not reachable.
	SuggestedNumberOfDlPackets *int `json:"suggestedNumberOfDlPackets"`
}

// PostScsAsIdConfigurationsJSONBody defines parameters for PostScsAsIdConfigurations.
type PostScsAsIdConfigurationsJSONBody NpConfiguration

// PutScsAsIdConfigurationsConfigurationIdJSONBody defines parameters for PutScsAsIdConfigurationsConfigurationId.
type PutScsAsIdConfigurationsConfigurationIdJSONBody NpConfiguration

// PostScsAsIdConfigurationsJSONRequestBody defines body for PostScsAsIdConfigurations for application/json ContentType.
type PostScsAsIdConfigurationsJSONRequestBody PostScsAsIdConfigurationsJSONBody

// PutScsAsIdConfigurationsConfigurationIdJSONRequestBody defines body for PutScsAsIdConfigurationsConfigurationId for application/json ContentType.
type PutScsAsIdConfigurationsConfigurationIdJSONRequestBody PutScsAsIdConfigurationsConfigurationIdJSONBody
