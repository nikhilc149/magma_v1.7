// Package TS29531NnssfNSSelection provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29531NnssfNSSelection

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29510NnrfNFManagement"
	externalRef1 "magma/feg/gateway/sbi/specs/TS29571CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// AllowedNssai defines model for AllowedNssai.
type AllowedNssai struct {
	AccessType        externalRef1.AccessType `json:"accessType"`
	AllowedSnssaiList []AllowedSnssai         `json:"allowedSnssaiList"`
}

// AllowedSnssai defines model for AllowedSnssai.
type AllowedSnssai struct {
	AllowedSnssai      externalRef1.Snssai  `json:"allowedSnssai"`
	MappedHomeSnssai   *externalRef1.Snssai `json:"mappedHomeSnssai,omitempty"`
	NsiInformationList *[]NsiInformation    `json:"nsiInformationList,omitempty"`
}

// AuthorizedNetworkSliceInfo defines model for AuthorizedNetworkSliceInfo.
type AuthorizedNetworkSliceInfo struct {
	AllowedNssaiList        *[]AllowedNssai                 `json:"allowedNssaiList,omitempty"`
	CandidateAmfList        *[]externalRef1.NfInstanceId    `json:"candidateAmfList,omitempty"`
	ConfiguredNssai         *[]ConfiguredSnssai             `json:"configuredNssai,omitempty"`
	NrfAmfSet               *externalRef1.Uri               `json:"nrfAmfSet,omitempty"`
	NrfAmfSetAccessTokenUri *externalRef1.Uri               `json:"nrfAmfSetAccessTokenUri,omitempty"`
	NrfAmfSetNfMgtUri       *externalRef1.Uri               `json:"nrfAmfSetNfMgtUri,omitempty"`
	NsiInformation          *NsiInformation                 `json:"nsiInformation,omitempty"`
	RejectedNssaiInPlmn     *[]externalRef1.Snssai          `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa       *[]externalRef1.Snssai          `json:"rejectedNssaiInTa,omitempty"`
	SupportedFeatures       *externalRef1.SupportedFeatures `json:"supportedFeatures,omitempty"`
	TargetAmfSet            *string                         `json:"targetAmfSet,omitempty"`
}

// ConfiguredSnssai defines model for ConfiguredSnssai.
type ConfiguredSnssai struct {
	ConfiguredSnssai externalRef1.Snssai  `json:"configuredSnssai"`
	MappedHomeSnssai *externalRef1.Snssai `json:"mappedHomeSnssai,omitempty"`
}

// MappingOfSnssai defines model for MappingOfSnssai.
type MappingOfSnssai struct {
	HomeSnssai    externalRef1.Snssai `json:"homeSnssai"`
	ServingSnssai externalRef1.Snssai `json:"servingSnssai"`
}

// NsiId defines model for NsiId.
type NsiId string

// NsiInformation defines model for NsiInformation.
type NsiInformation struct {
	NrfAccessTokenUri *externalRef1.Uri `json:"nrfAccessTokenUri,omitempty"`
	NrfId             externalRef1.Uri  `json:"nrfId"`
	NrfNfMgtUri       *externalRef1.Uri `json:"nrfNfMgtUri,omitempty"`
	NsiId             *NsiId            `json:"nsiId,omitempty"`
}

// RoamingIndication defines model for RoamingIndication.
type RoamingIndication interface{}

// SliceInfoForPDUSession defines model for SliceInfoForPDUSession.
type SliceInfoForPDUSession struct {
	HomeSnssai        *externalRef1.Snssai `json:"homeSnssai,omitempty"`
	RoamingIndication RoamingIndication    `json:"roamingIndication"`
	SNssai            externalRef1.Snssai  `json:"sNssai"`
}

// SliceInfoForRegistration defines model for SliceInfoForRegistration.
type SliceInfoForRegistration struct {
	AllowedNssaiCurrentAccess  *AllowedNssai          `json:"allowedNssaiCurrentAccess,omitempty"`
	AllowedNssaiOtherAccess    *AllowedNssai          `json:"allowedNssaiOtherAccess,omitempty"`
	DefaultConfiguredSnssaiInd *bool                  `json:"defaultConfiguredSnssaiInd,omitempty"`
	MappingOfNssai             *[]MappingOfSnssai     `json:"mappingOfNssai,omitempty"`
	RequestMapping             *bool                  `json:"requestMapping,omitempty"`
	RequestedNssai             *[]externalRef1.Snssai `json:"requestedNssai,omitempty"`
	SNssaiForMapping           *[]externalRef1.Snssai `json:"sNssaiForMapping,omitempty"`
	SubscribedNssai            *[]SubscribedSnssai    `json:"subscribedNssai,omitempty"`
}

// SliceInfoForUEConfigurationUpdate defines model for SliceInfoForUEConfigurationUpdate.
type SliceInfoForUEConfigurationUpdate struct {
	AllowedNssaiCurrentAccess  *AllowedNssai          `json:"allowedNssaiCurrentAccess,omitempty"`
	AllowedNssaiOtherAccess    *AllowedNssai          `json:"allowedNssaiOtherAccess,omitempty"`
	DefaultConfiguredSnssaiInd *bool                  `json:"defaultConfiguredSnssaiInd,omitempty"`
	MappingOfNssai             *[]MappingOfSnssai     `json:"mappingOfNssai,omitempty"`
	RequestedNssai             *[]externalRef1.Snssai `json:"requestedNssai,omitempty"`
	SubscribedNssai            *[]SubscribedSnssai    `json:"subscribedNssai,omitempty"`
}

// SubscribedSnssai defines model for SubscribedSnssai.
type SubscribedSnssai struct {
	DefaultIndication *bool               `json:"defaultIndication,omitempty"`
	SubscribedSnssai  externalRef1.Snssai `json:"subscribedSnssai"`
}

// NSSelectionGetParams defines parameters for NSSelectionGet.
type NSSelectionGetParams struct {
	// NF type of the NF service consumer
	NfType externalRef0.NFType `json:"nf-type"`

	// NF Instance ID of the NF service consumer
	NfId externalRef1.NfInstanceId `json:"nf-id"`

	// Requested network slice information during Registration procedure
	SliceInfoRequestForRegistration *SliceInfoForRegistration `json:"slice-info-request-for-registration,omitempty"`

	// Requested network slice information during PDU session establishment procedure
	SliceInfoRequestForPduSession *SliceInfoForPDUSession `json:"slice-info-request-for-pdu-session,omitempty"`

	// Requested network slice information during UE confuguration update procedure
	SliceInfoRequestForUeCu *SliceInfoForUEConfigurationUpdate `json:"slice-info-request-for-ue-cu,omitempty"`

	// PLMN ID of the HPLMN
	HomePlmnId *externalRef1.PlmnId `json:"home-plmn-id,omitempty"`

	// TAI of the UE
	Tai *externalRef1.Tai `json:"tai,omitempty"`

	// Features required to be supported by the NFs in the target slice instance
	SupportedFeatures *externalRef1.SupportedFeatures `json:"supported-features,omitempty"`
}
