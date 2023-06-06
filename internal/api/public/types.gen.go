// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package public

import (
	"time"
)

const (
	X_rh_identityScopes = "x_rh_identity.Scopes"
)

// Defines values for CreateDomainDomainType.
const (
	CreateDomainDomainTypeRhelIdm CreateDomainDomainType = "rhel-idm"
)

// Defines values for DomainDomainType.
const (
	DomainDomainTypeRhelIdm DomainDomainType = "rhel-idm"
)

// Defines values for ListDomainsDataDomainType.
const (
	ListDomainsDataDomainTypeRhelIdm ListDomainsDataDomainType = "rhel-idm"
)

// Defines values for RegisterDomainDomainType.
const (
	RegisterDomainDomainTypeRhelIdm RegisterDomainDomainType = "rhel-idm"
)

// Defines values for UpdateDomainDomainType.
const (
	RhelIdm UpdateDomainDomainType = "rhel-idm"
)

// CheckHosts Define the input data for the /check_host action.
//
// This action is launched from the ipa server to check the host that is being
// auto-enrolled.
type CheckHosts struct {
	// DomainName The domain name where to enroll the host vm.
	DomainName string `json:"domain_name"`

	// DomainType Indicate the type of domain. Actually only ipa is supported.
	DomainType string `json:"domain_type"`

	// InventoryId The id of the host vm into the insight host inventory.
	InventoryId string `json:"inventory_id"`

	// SubscriptionManagerId The subscription manager id for auto-enroll the host vm.
	SubscriptionManagerId string `json:"subscription_manager_id"`
}

// CreateDomain A domain resource
type CreateDomain struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// Description Human readable description for this domain.
	Description string `json:"description"`

	// DomainType Type of this domain. Currently only rhel-idm is supported.
	DomainType CreateDomainDomainType `json:"domain_type"`

	// Title the title for the entry
	Title string `json:"title"`
}

// CreateDomainDomainType Type of this domain. Currently only rhel-idm is supported.
type CreateDomainDomainType string

// Domain A domain resource
type Domain struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// Description Human readable description abou the domain.
	Description string `json:"description"`

	// DomainName Domain name
	DomainName string `json:"domain_name"`

	// DomainType Type of this domain. Currently only rhel-idm is supported.
	DomainType DomainDomainType `json:"domain_type"`

	// DomainUuid Internal id for this domain
	DomainUuid string `json:"domain_uuid"`

	// RhelIdm Options for ipa domains
	RhelIdm *DomainIpa `json:"rhel-idm,omitempty"`

	// Title Title to describe the domain.
	Title string `json:"title"`
}

// DomainDomainType Type of this domain. Currently only rhel-idm is supported.
type DomainDomainType string

// DomainIpa Options for ipa domains
type DomainIpa struct {
	// CaCerts A base64 representation of all the list of chain of certificates, including the server ca.
	CaCerts []DomainIpaCert `json:"ca_certs"`

	// RealmDomains List of realm associated to the IPA domain.
	RealmDomains []string `json:"realm_domains"`

	// RealmName The kerberos realm name associated to this IPA domain.
	RealmName string `json:"realm_name"`

	// Servers List of auto-enrollment enabled servers for this domain.
	Servers []DomainIpaServer `json:"servers"`
}

// DomainIpaCert Represent a certificate item in the cacerts list for the Ipa domain type.
type DomainIpaCert struct {
	Issuer         string    `json:"issuer"`
	Nickname       string    `json:"nickname"`
	NotValidAfter  time.Time `json:"not_valid_after"`
	NotValidBefore time.Time `json:"not_valid_before"`
	Pem            string    `json:"pem"`
	SerialNumber   string    `json:"serial_number"`
	Subject        string    `json:"subject"`
}

// DomainIpaServer Server schema for an entry into the Ipa domain type.
type DomainIpaServer struct {
	CaServer            bool   `json:"ca_server"`
	Fqdn                string `json:"fqdn"`
	HccEnrollmentServer bool   `json:"hcc_enrollment_server"`
	HccUpdateServer     bool   `json:"hcc_update_server"`

	// Location A location identifier (lower-case DNS label).
	Location              string `json:"location"`
	PkinitServer          bool   `json:"pkinit_server"`
	SubscriptionManagerId string `json:"subscription_manager_id"`
}

// Error General error schema
type Error struct {
	// Details A detailed explanation specific to this occurrence of the problem, e.g. a traceback.
	Details string `json:"details"`

	// Id A unique identifier for this particular occurrence of the problem.
	Id string `json:"id"`

	// Status The HTTP status code applicable to this problem, expressed as a string value. This SHOULD be provided.
	Status *int `json:"status,omitempty"`

	// Title A human-readable short explanation specific to this occurrence of the problem. This field’s value can be localized.
	Title string `json:"title"`
}

// Errors General error response returned by the hmsidm API
type Errors struct {
	// Errors Error objects provide additional information about problems encountered while performing an operation.
	Errors *[]Error `json:"errors,omitempty"`
}

// HostConf Represent the request payload for the /hostconf/:fqdn endpoint.
type HostConf struct {
	// Fqdn The full qualified domain name of the host vm that is being enrolled.
	Fqdn *string `json:"fqdn,omitempty"`

	// SubscriptionManagerId The id for the subscription manager.
	SubscriptionManagerId *string `json:"subscription_manager_id,omitempty"`
}

// HostConfResponseSchema The response for the action to retrieve the host vm information when
// it is being enrolled. This action is taken from the host vm.
type HostConfResponseSchema struct {
	DomainName *string `json:"domain_name,omitempty"`
	DomainType *string `json:"domain_type,omitempty"`
	Ipa        *struct {
		CaCert     *string   `json:"ca_cert,omitempty"`
		RealmName  *string   `json:"realm_name,omitempty"`
		ServerList *[]string `json:"server_list,omitempty"`
	} `json:"ipa,omitempty"`
}

// ListDomainsData The data listed for the domains.
type ListDomainsData struct {
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// Description Human-readable description of the domain entry.
	Description string                    `json:"description"`
	DomainName  string                    `json:"domain_name"`
	DomainType  ListDomainsDataDomainType `json:"domain_type"`
	DomainUuid  string                    `json:"domain_uuid"`

	// Title Human-friendly title for the domain entry.
	Title string `json:"title"`
}

// ListDomainsDataDomainType defines model for ListDomainsData.DomainType.
type ListDomainsDataDomainType string

// ListDomainsResponseSchema Represent a paginated result for a list of domains
type ListDomainsResponseSchema struct {
	// Data The content for this page.
	Data []ListDomainsData `json:"data"`

	// Links Represent the navigation links for the data paginated.
	Links PaginationLinks `json:"links"`

	// Meta Metadata for the paginated responses.
	Meta PaginationMeta `json:"meta"`
}

// PaginationLinks Represent the navigation links for the data paginated.
type PaginationLinks struct {
	// First Reference to the first page of the request.
	First *string `json:"first,omitempty"`

	// Last Reference to the last page of the request.
	Last *string `json:"last,omitempty"`

	// Next Reference to the next page of the request.
	Next *string `json:"next,omitempty"`

	// Previous Reference to the previous page of the request.
	Previous *string `json:"previous,omitempty"`
}

// PaginationMeta Metadata for the paginated responses.
type PaginationMeta struct {
	// Count total records in the collection.
	Count int64 `json:"count"`

	// Limit Number of items per page.
	Limit int `json:"limit"`

	// Offset Initial record of the page.
	Offset int `json:"offset"`
}

// RegisterDomain defines model for RegisterDomain.
type RegisterDomain struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// Description Human readable description abou the domain.
	Description string `json:"description"`

	// DomainName Domain name
	DomainName string `json:"domain_name"`

	// DomainType Type of this domain. Currently only rhel-idm is supported.
	DomainType RegisterDomainDomainType `json:"domain_type"`

	// DomainUuid Internal id for this domain
	DomainUuid string `json:"domain_uuid"`

	// RhelIdm Options for ipa domains
	RhelIdm *DomainIpa `json:"rhel-idm,omitempty"`

	// Title Title to describe the domain.
	Title string `json:"title"`
}

// RegisterDomainDomainType Type of this domain. Currently only rhel-idm is supported.
type RegisterDomainDomainType string

// UpdateDomain defines model for UpdateDomain.
type UpdateDomain struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// Description Human readable description abou the domain.
	Description string `json:"description"`

	// DomainName Domain name
	DomainName string `json:"domain_name"`

	// DomainType Type of this domain. Currently only rhel-idm is supported.
	DomainType UpdateDomainDomainType `json:"domain_type"`

	// DomainUuid Internal id for this domain
	DomainUuid string `json:"domain_uuid"`

	// RhelIdm Options for ipa domains
	RhelIdm *DomainIpa `json:"rhel-idm,omitempty"`

	// Title Title to describe the domain.
	Title string `json:"title"`
}

// UpdateDomainDomainType Type of this domain. Currently only rhel-idm is supported.
type UpdateDomainDomainType string

// CreateDomainResponse A domain resource
type CreateDomainResponse = Domain

// ErrorResponse General error response returned by the hmsidm API
type ErrorResponse = Errors

// HostConfResponse The response for the action to retrieve the host vm information when
// it is being enrolled. This action is taken from the host vm.
type HostConfResponse = HostConfResponseSchema

// ListDomainsResponse Represent a paginated result for a list of domains
type ListDomainsResponse = ListDomainsResponseSchema

// ReadDomainResponse A domain resource
type ReadDomainResponse = Domain

// RegisterDomainResponse A domain resource
type RegisterDomainResponse = Domain

// UpdateDomainResponse A domain resource
type UpdateDomainResponse = Domain

// CheckHostParams defines parameters for CheckHost.
type CheckHostParams struct {
	// XRhIdentity Identity header.
	XRhIdentity string `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *string `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// ListDomainsParams defines parameters for ListDomains.
type ListDomainsParams struct {
	// Offset pagination offset
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Number of items per page
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// XRhIdentity Identity header for the request
	XRhIdentity string `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *string `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// CreateDomainParams defines parameters for CreateDomain.
type CreateDomainParams struct {
	// XRhIdentity Identity header for the request
	XRhIdentity string `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *string `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// DeleteDomainParams defines parameters for DeleteDomain.
type DeleteDomainParams struct {
	// XRhIdentity Identity header for the request
	XRhIdentity string `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *string `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// ReadDomainParams defines parameters for ReadDomain.
type ReadDomainParams struct {
	// XRhIdentity Identity header for the request
	XRhIdentity string `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *string `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// RegisterDomainParams defines parameters for RegisterDomain.
type RegisterDomainParams struct {
	// XRhIdentity Identity header
	XRhIdentity string `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id
	XRhInsightsRequestId string `json:"X-Rh-Insights-Request-Id"`

	// XRhIdmRegistrationToken One time use token to register ipa information.
	XRhIdmRegistrationToken string `json:"X-Rh-Idm-Registration-Token"`

	// XRhIdmVersion ipa-hcc agent version
	XRhIdmVersion string `json:"X-Rh-Idm-Version"`
}

// UpdateDomainParams defines parameters for UpdateDomain.
type UpdateDomainParams struct {
	// XRhIdentity Identity header
	XRhIdentity string `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id
	XRhInsightsRequestId string `json:"X-Rh-Insights-Request-Id"`

	// XRhIdmVersion ipa-hcc agent version
	XRhIdmVersion string `json:"X-Rh-Idm-Version"`
}

// HostConfParams defines parameters for HostConf.
type HostConfParams struct {
	// XRhIdentity The identity header of the request.
	XRhIdentity string `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Unique request id for distributing tracing.
	XRhInsightsRequestId *string `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// CheckHostJSONRequestBody defines body for CheckHost for application/json ContentType.
type CheckHostJSONRequestBody = CheckHosts

// CreateDomainJSONRequestBody defines body for CreateDomain for application/json ContentType.
type CreateDomainJSONRequestBody = CreateDomain

// RegisterDomainJSONRequestBody defines body for RegisterDomain for application/json ContentType.
type RegisterDomainJSONRequestBody = RegisterDomain

// UpdateDomainJSONRequestBody defines body for UpdateDomain for application/json ContentType.
type UpdateDomainJSONRequestBody = UpdateDomain

// HostConfJSONRequestBody defines body for HostConf for application/json ContentType.
type HostConfJSONRequestBody = HostConf
