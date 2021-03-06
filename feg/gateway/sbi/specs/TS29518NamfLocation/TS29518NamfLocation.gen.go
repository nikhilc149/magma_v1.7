// Package TS29518NamfLocation provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29518NamfLocation

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29571CommonData"
	externalRef1 "magma/feg/gateway/sbi/specs/TS29572NlmfLocation"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// LocationEvent defines model for LocationEvent.
type LocationEvent interface{}

// LocationType defines model for LocationType.
type LocationType interface{}

// NotifiedPosInfo defines model for NotifiedPosInfo.
type NotifiedPosInfo struct {
	AgeOfLocationEstimate   *externalRef1.AgeOfLocationEstimate           `json:"ageOfLocationEstimate,omitempty"`
	Altitude                *externalRef1.Altitude                        `json:"altitude,omitempty"`
	BarometricPressure      *externalRef1.BarometricPressure              `json:"barometricPressure,omitempty"`
	CivicAddress            *externalRef1.CivicAddress                    `json:"civicAddress,omitempty"`
	Ecgi                    *externalRef0.Ecgi                            `json:"ecgi,omitempty"`
	GnssPositioningDataList *[]externalRef1.GnssPositioningMethodAndUsage `json:"gnssPositioningDataList,omitempty"`
	Gpsi                    *externalRef0.Gpsi                            `json:"gpsi,omitempty"`
	LocationEstimate        *externalRef1.GeographicArea                  `json:"locationEstimate,omitempty"`
	LocationEvent           LocationEvent                                 `json:"locationEvent"`
	Ncgi                    *externalRef0.Ncgi                            `json:"ncgi,omitempty"`
	Pei                     *externalRef0.Pei                             `json:"pei,omitempty"`
	PositioningDataList     *[]externalRef1.PositioningMethodAndUsage     `json:"positioningDataList,omitempty"`
	ServingNode             *externalRef0.NfInstanceId                    `json:"servingNode,omitempty"`
	Supi                    *externalRef0.Supi                            `json:"supi,omitempty"`
	VelocityEstimate        *externalRef1.VelocityEstimate                `json:"velocityEstimate,omitempty"`
}

// ProvideLocInfo defines model for ProvideLocInfo.
type ProvideLocInfo struct {
	CurrentLoc        *bool                               `json:"currentLoc,omitempty"`
	GeoInfo           *externalRef1.GeographicArea        `json:"geoInfo,omitempty"`
	Location          *externalRef0.UserLocation          `json:"location,omitempty"`
	LocationAge       *externalRef1.AgeOfLocationEstimate `json:"locationAge,omitempty"`
	RatType           *externalRef0.RatType               `json:"ratType,omitempty"`
	SupportedFeatures *externalRef0.SupportedFeatures     `json:"supportedFeatures,omitempty"`
	Timezone          *externalRef0.TimeZone              `json:"timezone,omitempty"`
}

// ProvidePosInfo defines model for ProvidePosInfo.
type ProvidePosInfo struct {
	AccuracyFulfilmentIndicator *externalRef1.AccuracyFulfilmentIndicator     `json:"accuracyFulfilmentIndicator,omitempty"`
	AgeOfLocationEstimate       *externalRef1.AgeOfLocationEstimate           `json:"ageOfLocationEstimate,omitempty"`
	Altitude                    *externalRef1.Altitude                        `json:"altitude,omitempty"`
	BarometricPressure          *externalRef1.BarometricPressure              `json:"barometricPressure,omitempty"`
	CivicAddress                *externalRef1.CivicAddress                    `json:"civicAddress,omitempty"`
	Ecgi                        *externalRef0.Ecgi                            `json:"ecgi,omitempty"`
	GnssPositioningDataList     *[]externalRef1.GnssPositioningMethodAndUsage `json:"gnssPositioningDataList,omitempty"`
	LocationEstimate            *externalRef1.GeographicArea                  `json:"locationEstimate,omitempty"`
	Ncgi                        *externalRef0.Ncgi                            `json:"ncgi,omitempty"`
	PositioningDataList         *[]externalRef1.PositioningMethodAndUsage     `json:"positioningDataList,omitempty"`
	SupportedFeatures           *externalRef0.SupportedFeatures               `json:"supportedFeatures,omitempty"`
	TargetServingNode           *externalRef0.NfInstanceId                    `json:"targetServingNode,omitempty"`
	VelocityEstimate            *externalRef1.VelocityEstimate                `json:"velocityEstimate,omitempty"`
}

// RequestLocInfo defines model for RequestLocInfo.
type RequestLocInfo struct {
	OldGuami          *externalRef0.Guami             `json:"oldGuami,omitempty"`
	Req5gsLoc         *bool                           `json:"req5gsLoc,omitempty"`
	ReqCurrentLoc     *bool                           `json:"reqCurrentLoc,omitempty"`
	ReqRatType        *bool                           `json:"reqRatType,omitempty"`
	ReqTimeZone       *bool                           `json:"reqTimeZone,omitempty"`
	SupportedFeatures *externalRef0.SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// RequestPosInfo defines model for RequestPosInfo.
type RequestPosInfo struct {
	AdditionalLcsSuppGADShapes *[]externalRef1.SupportedGADShapes `json:"additionalLcsSuppGADShapes,omitempty"`
	Gpsi                       *externalRef0.Gpsi                 `json:"gpsi,omitempty"`
	LcsClientType              externalRef1.ExternalClientType    `json:"lcsClientType"`
	LcsLocation                LocationType                       `json:"lcsLocation"`
	LcsQoS                     *externalRef1.LocationQoS          `json:"lcsQoS,omitempty"`
	LcsSupportedGADShapes      *externalRef1.SupportedGADShapes   `json:"lcsSupportedGADShapes,omitempty"`
	LocationNotificationUri    *externalRef0.Uri                  `json:"locationNotificationUri,omitempty"`
	OldGuami                   *externalRef0.Guami                `json:"oldGuami,omitempty"`
	Priority                   *externalRef1.LcsPriority          `json:"priority,omitempty"`
	Supi                       *externalRef0.Supi                 `json:"supi,omitempty"`
	SupportedFeatures          *externalRef0.SupportedFeatures    `json:"supportedFeatures,omitempty"`
	VelocityRequested          *externalRef1.VelocityRequested    `json:"velocityRequested,omitempty"`
}

// ProvideLocationInfoJSONBody defines parameters for ProvideLocationInfo.
type ProvideLocationInfoJSONBody RequestLocInfo

// ProvidePositioningInfoJSONBody defines parameters for ProvidePositioningInfo.
type ProvidePositioningInfoJSONBody RequestPosInfo

// ProvideLocationInfoJSONRequestBody defines body for ProvideLocationInfo for application/json ContentType.
type ProvideLocationInfoJSONRequestBody ProvideLocationInfoJSONBody

// ProvidePositioningInfoJSONRequestBody defines body for ProvidePositioningInfo for application/json ContentType.
type ProvidePositioningInfoJSONRequestBody ProvidePositioningInfoJSONBody
