// Package TS29522TrafficInfluence provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29522TrafficInfluence

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29122CommonData"
	externalRef1 "magma/feg/gateway/sbi/specs/TS29514NpcfPolicyAuthorization"
	externalRef2 "magma/feg/gateway/sbi/specs/TS29571CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// EventNotification defines model for EventNotification.
type EventNotification struct {
	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransId *string `json:"afTransId,omitempty"`

	// Possible values are - EARLY: Early notification of UP path reconfiguration. - EARLY_LATE: Early and late notification of UP path reconfiguration. This value shall only be present in the subscription to the DNAI change event. - LATE: Late notification of UP path reconfiguration.
	DnaiChgType        externalRef2.DnaiChangeType   `json:"dnaiChgType"`
	Gpsi               *externalRef2.Gpsi            `json:"gpsi,omitempty"`
	SourceDnai         *externalRef2.Dnai            `json:"sourceDnai,omitempty"`
	SourceTrafficRoute *externalRef2.RouteToLocation `json:"sourceTrafficRoute"`

	// string identifying a Ipv4 address formatted in the "dotted decimal" notation as defined in IETF RFC 1166.
	SrcUeIpv4Addr   *externalRef0.Ipv4Addr   `json:"srcUeIpv4Addr,omitempty"`
	SrcUeIpv6Prefix *externalRef2.Ipv6Prefix `json:"srcUeIpv6Prefix,omitempty"`

	// Possible values are - UP_PATH_CHANGE: The AF requests to be notified when the UP path changes for the PDU session.
	SubscribedEvent    SubscribedEvent               `json:"subscribedEvent"`
	TargetDnai         *externalRef2.Dnai            `json:"targetDnai,omitempty"`
	TargetTrafficRoute *externalRef2.RouteToLocation `json:"targetTrafficRoute"`

	// string identifying a Ipv4 address formatted in the "dotted decimal" notation as defined in IETF RFC 1166.
	TgtUeIpv4Addr   *externalRef0.Ipv4Addr   `json:"tgtUeIpv4Addr,omitempty"`
	TgtUeIpv6Prefix *externalRef2.Ipv6Prefix `json:"tgtUeIpv6Prefix,omitempty"`
	UeMac           *externalRef2.MacAddr48  `json:"ueMac,omitempty"`
}

// Possible values are - UP_PATH_CHANGE: The AF requests to be notified when the UP path changes for the PDU session.
type SubscribedEvent interface{}

// TrafficInfluSub defines model for TrafficInfluSub.
type TrafficInfluSub interface{}

// TrafficInfluSubPatch defines model for TrafficInfluSubPatch.
type TrafficInfluSubPatch struct {
	// Identifies whether an application can be relocated once a location of the application has been selected.
	AppReloInd *bool `json:"appReloInd"`

	// Identifies Ethernet packet filters.
	EthTrafficFilters *[]externalRef1.EthFlowDescription `json:"ethTrafficFilters,omitempty"`
	TempValidities    *[]externalRef1.TemporalValidity   `json:"tempValidities"`

	// Identifies IP packet filters.
	TrafficFilters *[]externalRef0.FlowInfo `json:"trafficFilters,omitempty"`

	// Identifies the N6 traffic routing requirement.
	TrafficRoutes *[]externalRef2.RouteToLocation `json:"trafficRoutes,omitempty"`

	// Identifies a geographic zone that the AF request applies only to the traffic of UE(s) located in this specific zone.
	ValidGeoZoneIds *[]string `json:"validGeoZoneIds"`
}

// PostAfIdSubscriptionsJSONBody defines parameters for PostAfIdSubscriptions.
type PostAfIdSubscriptionsJSONBody TrafficInfluSub

// PutAfIdSubscriptionsSubscriptionIdJSONBody defines parameters for PutAfIdSubscriptionsSubscriptionId.
type PutAfIdSubscriptionsSubscriptionIdJSONBody TrafficInfluSub

// PostAfIdSubscriptionsJSONRequestBody defines body for PostAfIdSubscriptions for application/json ContentType.
type PostAfIdSubscriptionsJSONRequestBody PostAfIdSubscriptionsJSONBody

// PutAfIdSubscriptionsSubscriptionIdJSONRequestBody defines body for PutAfIdSubscriptionsSubscriptionId for application/json ContentType.
type PutAfIdSubscriptionsSubscriptionIdJSONRequestBody PutAfIdSubscriptionsSubscriptionIdJSONBody
