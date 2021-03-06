// Package TS29122ReportingNetworkStatus provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29122ReportingNetworkStatus

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29122CommonData"
	externalRef1 "magma/feg/gateway/sbi/specs/TS29571CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// Possible values are - HIGH: The congestion status is high. - MEDIUM: The congestion status is medium. - LOW: The congestion status is low.
type CongestionType interface{}

// Unsigned integer with valid values between 0 and 31. The value 0 indicates that there is no congestion. The value 1 is the lowest congestion level and value 31 is the highest congestion level.
type CongestionValue int

// NetworkStatusReportingNotification defines model for NetworkStatusReportingNotification.
type NetworkStatusReportingNotification struct {
	// Possible values are - HIGH: The congestion status is high. - MEDIUM: The congestion status is medium. - LOW: The congestion status is low.
	NsiType *CongestionType `json:"nsiType,omitempty"`

	// Unsigned integer with valid values between 0 and 31. The value 0 indicates that there is no congestion. The value 1 is the lowest congestion level and value 31 is the highest congestion level.
	NsiValue *CongestionValue `json:"nsiValue,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Subscription externalRef0.Link `json:"subscription"`
}

// NetworkStatusReportingSubscription defines model for NetworkStatusReportingSubscription.
type NetworkStatusReportingSubscription struct {
	LocationArea externalRef0.LocationArea `json:"locationArea"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination externalRef0.Link `json:"notificationDestination"`

	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in subclause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification *bool `json:"requestTestNotification,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self              *externalRef0.Link              `json:"self,omitempty"`
	SupportedFeatures *externalRef1.SupportedFeatures `json:"supportedFeatures,omitempty"`

	// Identifies a list of congestion level(s) with abstracted value that the SCS/AS requests to be informed of when reached.
	ThresholdTypes *[]CongestionType `json:"thresholdTypes,omitempty"`

	// Identifies a list of congestion level(s) with exact value that the SCS/AS requests to be informed of when reached.
	ThresholdValues *[]CongestionValue `json:"thresholdValues,omitempty"`

	// string with format "date-time" as defined in OpenAPI.
	TimeDuration       *externalRef0.DateTime           `json:"timeDuration,omitempty"`
	WebsockNotifConfig *externalRef0.WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
}

// PostScsAsIdSubscriptionsJSONBody defines parameters for PostScsAsIdSubscriptions.
type PostScsAsIdSubscriptionsJSONBody NetworkStatusReportingSubscription

// PutScsAsIdSubscriptionsSubscriptionIdJSONBody defines parameters for PutScsAsIdSubscriptionsSubscriptionId.
type PutScsAsIdSubscriptionsSubscriptionIdJSONBody NetworkStatusReportingSubscription

// PostScsAsIdSubscriptionsJSONRequestBody defines body for PostScsAsIdSubscriptions for application/json ContentType.
type PostScsAsIdSubscriptionsJSONRequestBody PostScsAsIdSubscriptionsJSONBody

// PutScsAsIdSubscriptionsSubscriptionIdJSONRequestBody defines body for PutScsAsIdSubscriptionsSubscriptionId for application/json ContentType.
type PutScsAsIdSubscriptionsSubscriptionIdJSONRequestBody PutScsAsIdSubscriptionsSubscriptionIdJSONBody
