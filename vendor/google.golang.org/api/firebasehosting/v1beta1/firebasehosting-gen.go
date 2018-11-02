// Package firebasehosting provides access to the Firebase Hosting API.
//
// See https://firebase.google.com/docs/hosting/
//
// Usage example:
//
//   import "google.golang.org/api/firebasehosting/v1beta1"
//   ...
//   firebasehostingService, err := firebasehosting.New(oauthHttpClient)
package firebasehosting // import "google.golang.org/api/firebasehosting/v1beta1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "firebasehosting:v1beta1"
const apiName = "firebasehosting"
const apiVersion = "v1beta1"
const basePath = "https://firebasehosting.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"

	// View and administer all your Firebase data and settings
	FirebaseScope = "https://www.googleapis.com/auth/firebase"

	// View all your Firebase data and settings
	FirebaseReadonlyScope = "https://www.googleapis.com/auth/firebase.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Sites = NewSitesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Sites *SitesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewSitesService(s *Service) *SitesService {
	rs := &SitesService{s: s}
	rs.Domains = NewSitesDomainsService(s)
	rs.Releases = NewSitesReleasesService(s)
	rs.Versions = NewSitesVersionsService(s)
	return rs
}

type SitesService struct {
	s *Service

	Domains *SitesDomainsService

	Releases *SitesReleasesService

	Versions *SitesVersionsService
}

func NewSitesDomainsService(s *Service) *SitesDomainsService {
	rs := &SitesDomainsService{s: s}
	return rs
}

type SitesDomainsService struct {
	s *Service
}

func NewSitesReleasesService(s *Service) *SitesReleasesService {
	rs := &SitesReleasesService{s: s}
	return rs
}

type SitesReleasesService struct {
	s *Service
}

func NewSitesVersionsService(s *Service) *SitesVersionsService {
	rs := &SitesVersionsService{s: s}
	rs.Files = NewSitesVersionsFilesService(s)
	return rs
}

type SitesVersionsService struct {
	s *Service

	Files *SitesVersionsFilesService
}

func NewSitesVersionsFilesService(s *Service) *SitesVersionsFilesService {
	rs := &SitesVersionsFilesService{s: s}
	return rs
}

type SitesVersionsFilesService struct {
	s *Service
}

// ActingUser: Contains metadata about the user who performed an action,
// such as creating
// a release or finalizing a version.
type ActingUser struct {
	// Email: The email address of the user when the user performed the
	// action.
	Email string `json:"email,omitempty"`

	// ImageUrl: A profile image URL for the user. May not be present if the
	// user has
	// changed their email address or deleted their account.
	ImageUrl string `json:"imageUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ActingUser) MarshalJSON() ([]byte, error) {
	type NoMethod ActingUser
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CertDnsChallenge: Represents a DNS certificate challenge.
type CertDnsChallenge struct {
	// DomainName: The domain name upon which the DNS challenge must be
	// satisfied.
	DomainName string `json:"domainName,omitempty"`

	// Token: The value that must be present as a TXT record on the domain
	// name to
	// satisfy the challenge.
	Token string `json:"token,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DomainName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DomainName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CertDnsChallenge) MarshalJSON() ([]byte, error) {
	type NoMethod CertDnsChallenge
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CertHttpChallenge: Represents an HTTP certificate challenge.
type CertHttpChallenge struct {
	// Path: The URL path on which to serve the specified token to satisfy
	// the
	// certificate challenge.
	Path string `json:"path,omitempty"`

	// Token: The token to serve at the specified URL path to satisfy the
	// certificate
	// challenge.
	Token string `json:"token,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Path") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Path") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CertHttpChallenge) MarshalJSON() ([]byte, error) {
	type NoMethod CertHttpChallenge
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Domain: The intended behavior and status information of a domain.
type Domain struct {
	// DomainName: Required. The domain name of the association.
	DomainName string `json:"domainName,omitempty"`

	// DomainRedirect: If set, the domain should redirect with the provided
	// parameters.
	DomainRedirect *DomainRedirect `json:"domainRedirect,omitempty"`

	// Provisioning: Output only. Information about the provisioning of
	// certificates and the
	// health of the DNS resolution for the domain.
	Provisioning *DomainProvisioning `json:"provisioning,omitempty"`

	// Site: Required. The site name of the association.
	Site string `json:"site,omitempty"`

	// Status: Output only. Additional status of the domain association.
	//
	// Possible values:
	//   "DOMAIN_STATUS_UNSPECIFIED" - Unspecified domain association
	// status.
	//   "DOMAIN_CHANGE_PENDING" - An operation is in progress on the domain
	// association and no further
	// operations can be performed until it is complete.
	//   "DOMAIN_ACTIVE" - The domain association is active and no
	// additional action is required.
	//   "DOMAIN_VERIFICATION_REQUIRED" - The domain was previously verified
	// in the legacy system. User must
	// reverify the domain through the ownership service.
	//   "DOMAIN_VERIFICATION_LOST" - The domain verification has been lost
	// and the domain is in the grace period
	// before being removed from the Firebase Hosting site.
	Status string `json:"status,omitempty"`

	// UpdateTime: Output only. The time at which the domain was last
	// updated.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DomainName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DomainName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Domain) MarshalJSON() ([]byte, error) {
	type NoMethod Domain
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DomainProvisioning: The current certificate provisioning status
// information for a domain.
type DomainProvisioning struct {
	// CertChallengeDiscoveredTxt: The TXT records (for the certificate
	// challenge) that were found at the last
	// DNS fetch.
	CertChallengeDiscoveredTxt []string `json:"certChallengeDiscoveredTxt,omitempty"`

	// CertChallengeDns: The DNS challenge for generating a certificate.
	CertChallengeDns *CertDnsChallenge `json:"certChallengeDns,omitempty"`

	// CertChallengeHttp: The HTTP challenge for generating a certificate.
	CertChallengeHttp *CertHttpChallenge `json:"certChallengeHttp,omitempty"`

	// CertStatus: The certificate provisioning status; updated when
	// Firebase Hosting
	// provisions an SSL certificate for the domain.
	//
	// Possible values:
	//   "CERT_STATUS_UNSPECIFIED" - Unspecified certificate provisioning
	// status.
	//   "CERT_PENDING" - Waiting for certificate challenge to be created.
	//   "CERT_MISSING" - Waiting for certificate challenge to be met.
	//   "CERT_PROCESSING" - Certificate challenge met; attempting to
	// acquire/propagate certificate.
	//   "CERT_PROPAGATING" - Certificate obtained; propagating to the CDN.
	//   "CERT_ACTIVE" - Certificate provisioned and deployed across the
	// CDN.
	//   "CERT_ERROR" - Certificate provisioning failed in a non-recoverable
	// manner.
	CertStatus string `json:"certStatus,omitempty"`

	// DiscoveredIps: The IPs found at the last DNS fetch.
	DiscoveredIps []string `json:"discoveredIps,omitempty"`

	// DnsFetchTime: The time at which the last DNS fetch occurred.
	DnsFetchTime string `json:"dnsFetchTime,omitempty"`

	// DnsStatus: The DNS record match status as of the last DNS fetch.
	//
	// Possible values:
	//   "DNS_STATUS_UNSPECIFIED" - Unspecified DNS status.
	//   "DNS_PENDING" - No DNS records have been specified for this domain
	// yet.
	//   "DNS_MISSING" - None of the required DNS records have been detected
	// on the domain.
	//   "DNS_PARTIAL_MATCH" - Some of the required DNS records were
	// detected, but not all of them. No
	// extra (non-required) DNS records were detected.
	//   "DNS_MATCH" - All required DNS records were detected. No extra
	// (non-required) DNS records
	// were detected.
	//   "DNS_EXTRANEOUS_MATCH" - The domain has at least one of the
	// required DNS records, and it has at
	// least one extra (non-required) DNS record.
	DnsStatus string `json:"dnsStatus,omitempty"`

	// ExpectedIps: The list of IPs to which the domain is expected to
	// resolve.
	ExpectedIps []string `json:"expectedIps,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "CertChallengeDiscoveredTxt") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "CertChallengeDiscoveredTxt") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DomainProvisioning) MarshalJSON() ([]byte, error) {
	type NoMethod DomainProvisioning
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DomainRedirect: Defines the behavior of a domain-level redirect.
// Domain redirects preserve
// the path of the redirect but replace the requested domain with the
// one
// specified in the redirect configuration.
type DomainRedirect struct {
	// DomainName: Required. The domain name to redirect to.
	DomainName string `json:"domainName,omitempty"`

	// Type: Required. The redirect status code.
	//
	// Possible values:
	//   "REDIRECT_TYPE_UNSPECIFIED" - The default redirect type; should not
	// be intentionlly used.
	//   "MOVED_PERMANENTLY" - The redirect will respond with an HTTP status
	// code of
	// `301 Moved Permanently`.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DomainName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DomainName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DomainRedirect) MarshalJSON() ([]byte, error) {
	type NoMethod DomainRedirect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Header: A `Header` defines custom headers to add to a response should
// the request
// URL path match the pattern.
type Header struct {
	// Glob: Required. The user-supplied
	// [glob pattern](/docs/hosting/full-config#section-glob) to match
	// against
	// the request URL path.
	Glob string `json:"glob,omitempty"`

	// Headers: Required. The additional headers to add to the response.
	Headers map[string]string `json:"headers,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Glob") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Glob") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Header) MarshalJSON() ([]byte, error) {
	type NoMethod Header
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListDomainsResponse struct {
	// Domains: The list of domains, if any exist.
	Domains []*Domain `json:"domains,omitempty"`

	// NextPageToken: The pagination token, if more results exist.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Domains") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Domains") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListDomainsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListDomainsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListReleasesResponse struct {
	// NextPageToken: If there are additional releases remaining beyond the
	// ones in this
	// response, then supply this token in the
	// next
	// [`list`](../sites.versions.files/list) call to continue with the next
	// set
	// of releases.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Releases: The list of hashes of files that still need to be uploaded,
	// if any exist.
	Releases []*Release `json:"releases,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListReleasesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListReleasesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListVersionFilesResponse struct {
	// Files: The list path/hashes in the specified version.
	Files []*VersionFile `json:"files,omitempty"`

	// NextPageToken: The pagination token, if more results exist.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Files") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Files") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListVersionFilesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListVersionFilesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PopulateVersionFilesRequest struct {
	// Files: A set of file paths to the hashes corresponding to assets that
	// should be
	// added to the version. Note that a file path to an empty hash will
	// remove
	// the path from the version. Calculate a hash by Gzipping the file
	// then
	// taking the SHA256 hash of the newly compressed file.
	Files map[string]string `json:"files,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Files") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Files") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PopulateVersionFilesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod PopulateVersionFilesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PopulateVersionFilesResponse struct {
	// UploadRequiredHashes: The content hashes of the specified files that
	// need to be uploaded to the
	// specified endpoint.
	UploadRequiredHashes []string `json:"uploadRequiredHashes,omitempty"`

	// UploadUrl: The URL to which the files should be uploaded, in the
	// format:
	// <br>"https://upload-firebasehosting.googleapis.com/upload/site
	// s/<var>site-name</var>/versions/<var>versionID</var>/files".
	// <br>Perfo
	// rm a multipart `POST` of the Gzipped file contents to the URL
	// using a forward slash and the hash of the file appended to the end.
	UploadUrl string `json:"uploadUrl,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "UploadRequiredHashes") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "UploadRequiredHashes") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PopulateVersionFilesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod PopulateVersionFilesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Redirect: A `Redirect` represents the configuration for returning an
// HTTP redirect
// response given a matching request URL path.
type Redirect struct {
	// Glob: Required. The user-supplied
	// [glob pattern](/docs/hosting/full-config#section-glob) to match
	// against
	// the request URL path.
	Glob string `json:"glob,omitempty"`

	// Location: Required. The value to put in the HTTP location header of
	// the response.
	// <br>The location can contain capture group values from the pattern
	// using a
	// ":" prefix to identify the segment and an optional "*" to capture
	// the
	// rest of the URL.
	// For example:
	// <code>"glob": "/:capture*",
	// <br>"statusCode": 301,
	// <br>"location": "https://example.com/foo/:capture"</code>
	Location string `json:"location,omitempty"`

	// StatusCode: Required. The status HTTP code to return in the response.
	// It must be a
	// valid 3xx status code.
	StatusCode int64 `json:"statusCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Glob") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Glob") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Redirect) MarshalJSON() ([]byte, error) {
	type NoMethod Redirect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Release: A `Release` is a particular
// [collection of configurations and files](sites.versions)
// that is set to be public at a particular time.
type Release struct {
	// Message: The deploy description when the release was created. The
	// value can be up to
	// 512&nbsp;characters.
	Message string `json:"message,omitempty"`

	// Name: Output only. The unique identifier for the release, in the
	// format:
	// <code>sites/<var>site-name</var>/releases/<var>releaseID</var>
	// </code>
	// This name is provided in the response body when you call
	// the
	// [`CreateRelease`](sites.releases/create) endpoint.
	Name string `json:"name,omitempty"`

	// ReleaseTime: Output only. The time at which the version is set to be
	// public.
	ReleaseTime string `json:"releaseTime,omitempty"`

	// ReleaseUser: Output only. Identifies the user who created the
	// release.
	ReleaseUser *ActingUser `json:"releaseUser,omitempty"`

	// Type: Explains the reason for the release.
	// <br>Specify a value for this field only when creating a
	// `SITE_DISABLE`
	// type release.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - An unspecified type. Indicates that a version
	// was released.
	// <br>This is the default value when no other `type` is
	// explicitly
	// specified.
	//   "DEPLOY" - A version was uploaded to Firebase Hosting and released.
	//   "ROLLBACK" - The release points back to a previously deployed
	// version.
	//   "SITE_DISABLE" - The release prevents the site from serving
	// content. Firebase Hosting acts
	// as if the site never existed.
	Type string `json:"type,omitempty"`

	// Version: Output only.  The configuration and content that was
	// released.
	Version *Version `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Message") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Message") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Release) MarshalJSON() ([]byte, error) {
	type NoMethod Release
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Rewrite: A `Rewrite` represents an internal content rewrite on the
// version. If the
// pattern matches, the request will be handled as if it were to
// the
// destination path specified in the configuration.
type Rewrite struct {
	// DynamicLinks: The request will be forwarded to Firebase Dynamic
	// Links.
	DynamicLinks bool `json:"dynamicLinks,omitempty"`

	// Function: The function to proxy requests to. Must match the exported
	// function
	// name exactly.
	Function string `json:"function,omitempty"`

	// Glob: Required. The user-supplied
	// [glob pattern](/docs/hosting/full-config#section-glob) to match
	// against
	// the request URL path.
	Glob string `json:"glob,omitempty"`

	// Path: The URL path to rewrite the request to.
	Path string `json:"path,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DynamicLinks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DynamicLinks") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Rewrite) MarshalJSON() ([]byte, error) {
	type NoMethod Rewrite
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ServingConfig: The configuration for how incoming requests to a site
// should be routed and
// processed before serving content. The patterns are matched and
// applied
// according to a specific
// [priority
// order](/docs/hosting/url-redirects-rewrites#section-priorities).
type ServingConfig struct {
	// AppAssociation: How to handle well known App Association files.
	//
	// Possible values:
	//   "AUTO" - The app association files will be automattically created
	// from the apps
	// that exist in the Firebase project.
	//   "NONE" - No special handling of the app association files will
	// occur, these paths
	// will result in a 404 unless caught with a Rewrite.
	AppAssociation string `json:"appAssociation,omitempty"`

	// CleanUrls: Defines whether to drop the file extension from uploaded
	// files.
	CleanUrls bool `json:"cleanUrls,omitempty"`

	// Headers: A list of custom response headers that are added to the
	// content if the
	// request URL path matches the glob.
	Headers []*Header `json:"headers,omitempty"`

	// Redirects: A list of globs that will cause the response to redirect
	// to another
	// location.
	Redirects []*Redirect `json:"redirects,omitempty"`

	// Rewrites: A list of rewrites that will act as if the service were
	// given the
	// destination URL.
	Rewrites []*Rewrite `json:"rewrites,omitempty"`

	// TrailingSlashBehavior: Defines how to handle a trailing slash in the
	// URL path.
	//
	// Possible values:
	//   "TRAILING_SLASH_BEHAVIOR_UNSPECIFIED" - No behavior is
	// specified.
	// <br>Files are served at their exact location only, and trailing
	// slashes
	// are only added to directory indexes.
	//   "ADD" - Trailing slashes are _added_ to directory indexes as well
	// as to any URL
	// path not ending in a file extension.
	//   "REMOVE" - Trailing slashes are _removed_ from directory indexes as
	// well as from any
	// URL path not ending in a file extension.
	TrailingSlashBehavior string `json:"trailingSlashBehavior,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppAssociation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppAssociation") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ServingConfig) MarshalJSON() ([]byte, error) {
	type NoMethod ServingConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Version: A `Version` is the collection of configuration and
// [static files](sites.versions.files) that determine how a site is
// displayed.
type Version struct {
	// Config: The configuration for the behavior of the site. This
	// configuration exists
	// in the [`firebase.json`](/docs/cli/#the_firebasejson_file) file.
	Config *ServingConfig `json:"config,omitempty"`

	// CreateTime: Output only. The time at which the version was created.
	CreateTime string `json:"createTime,omitempty"`

	// CreateUser: Output only. Identifies the user who created the version.
	CreateUser *ActingUser `json:"createUser,omitempty"`

	// DeleteTime: Output only. The time at which the version was `DELETED`.
	DeleteTime string `json:"deleteTime,omitempty"`

	// DeleteUser: Output only. Identifies the user who `DELETED` the
	// version.
	DeleteUser *ActingUser `json:"deleteUser,omitempty"`

	// FileCount: Output only. The total number of files associated with the
	// version.
	// <br>This value is calculated after a version is `FINALIZED`.
	FileCount int64 `json:"fileCount,omitempty,string"`

	// FinalizeTime: Output only. The time at which the version was
	// `FINALIZED`.
	FinalizeTime string `json:"finalizeTime,omitempty"`

	// FinalizeUser: Output only. Identifies the user who `FINALIZED` the
	// version.
	FinalizeUser *ActingUser `json:"finalizeUser,omitempty"`

	// Labels: The labels used for extra metadata and/or filtering.
	Labels map[string]string `json:"labels,omitempty"`

	// Name: The unique identifier for a version, in the
	// format:
	// <code>sites/<var>site-name</var>/versions/<var>versionID</var>
	// </code>
	// This name is provided in the response body when you call
	// the
	// [`CreateVersion`](../sites.versions/create) endpoint.
	Name string `json:"name,omitempty"`

	// Status: The deploy status of a version.
	// <br>
	// <br>For a successful deploy, call
	// the
	// [`CreateVersion`](sites.versions/create) endpoint to make a new
	// version
	// (`CREATED` status),
	// [upload all desired files](sites.versions/populateFiles) to the
	// version,
	// then [update](sites.versions/patch) the version to the `FINALIZED`
	// status.
	// <br>
	// <br>Note that if you leave the version in the `CREATED` state for
	// more
	// than 12&nbsp;hours, the system will automatically mark the version
	// as
	// `ABANDONED`.
	// <br>
	// <br>You can also change the status of a version to `DELETED` by
	// calling the
	// [`DeleteVersion`](sites.versions/delete) endpoint.
	//
	// Possible values:
	//   "VERSION_STATUS_UNSPECIFIED" - The default status; should not be
	// intentionally used.
	//   "CREATED" - The version has been created, and content is currently
	// being added to the
	// version.
	//   "FINALIZED" - All content has been added to the version, and the
	// version can no longer be
	// changed.
	//   "DELETED" - The version has been deleted.
	//   "ABANDONED" - The version was not updated to `FINALIZED` within
	// 12&nbsp;hours and was
	// automatically deleted.
	//   "EXPIRED" - The version has fallen out of the site-configured
	// retention window and its
	// associated files in GCS have been/been scheduled for deletion.
	Status string `json:"status,omitempty"`

	// VersionBytes: Output only. The total stored bytesize of the
	// version.
	// <br>This value is calculated after a version is `FINALIZED`.
	VersionBytes int64 `json:"versionBytes,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Config") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Config") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Version) MarshalJSON() ([]byte, error) {
	type NoMethod Version
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VersionFile: A static content file that is part of a version.
type VersionFile struct {
	// Hash: The SHA256 content hash of the file.
	Hash string `json:"hash,omitempty"`

	// Path: The URI at which the file's content should display.
	Path string `json:"path,omitempty"`

	// Status: Output only. The current status of a particular file in the
	// specified
	// version.
	// <br>The value will be either `pending upload` or `uploaded`.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - The default status; should not be
	// intentionally used.
	//   "EXPECTED" - The file has been included in the version and is
	// expected to be uploaded
	// in the near future.
	//   "ACTIVE" - The file has already been uploaded to Firebase Hosting.
	Status string `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Hash") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Hash") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VersionFile) MarshalJSON() ([]byte, error) {
	type NoMethod VersionFile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "firebasehosting.sites.domains.create":

type SitesDomainsCreateCall struct {
	s          *Service
	parent     string
	domain     *Domain
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a domain mapping on the specified site.
func (r *SitesDomainsService) Create(parent string, domain *Domain) *SitesDomainsCreateCall {
	c := &SitesDomainsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.domain = domain
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesDomainsCreateCall) Fields(s ...googleapi.Field) *SitesDomainsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesDomainsCreateCall) Context(ctx context.Context) *SitesDomainsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesDomainsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesDomainsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.domain)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/domains")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.domains.create" call.
// Exactly one of *Domain or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Domain.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SitesDomainsCreateCall) Do(opts ...googleapi.CallOption) (*Domain, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Domain{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a domain mapping on the specified site.",
	//   "flatPath": "v1beta1/sites/{sitesId}/domains",
	//   "httpMethod": "POST",
	//   "id": "firebasehosting.sites.domains.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The parent to create the domain association for, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e\u003c/code\u003e",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/domains",
	//   "request": {
	//     "$ref": "Domain"
	//   },
	//   "response": {
	//     "$ref": "Domain"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/firebase"
	//   ]
	// }

}

// method id "firebasehosting.sites.domains.delete":

type SitesDomainsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the existing domain mapping on the specified site.
func (r *SitesDomainsService) Delete(name string) *SitesDomainsDeleteCall {
	c := &SitesDomainsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesDomainsDeleteCall) Fields(s ...googleapi.Field) *SitesDomainsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesDomainsDeleteCall) Context(ctx context.Context) *SitesDomainsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesDomainsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesDomainsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.domains.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SitesDomainsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the existing domain mapping on the specified site.",
	//   "flatPath": "v1beta1/sites/{sitesId}/domains/{domainsId}",
	//   "httpMethod": "DELETE",
	//   "id": "firebasehosting.sites.domains.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the domain association to delete.",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+/domains/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/firebase"
	//   ]
	// }

}

// method id "firebasehosting.sites.domains.get":

type SitesDomainsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a domain mapping on the specified site.
func (r *SitesDomainsService) Get(name string) *SitesDomainsGetCall {
	c := &SitesDomainsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesDomainsGetCall) Fields(s ...googleapi.Field) *SitesDomainsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SitesDomainsGetCall) IfNoneMatch(entityTag string) *SitesDomainsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesDomainsGetCall) Context(ctx context.Context) *SitesDomainsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesDomainsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesDomainsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.domains.get" call.
// Exactly one of *Domain or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Domain.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SitesDomainsGetCall) Do(opts ...googleapi.CallOption) (*Domain, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Domain{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a domain mapping on the specified site.",
	//   "flatPath": "v1beta1/sites/{sitesId}/domains/{domainsId}",
	//   "httpMethod": "GET",
	//   "id": "firebasehosting.sites.domains.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the domain configuration to get.",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+/domains/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Domain"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/firebase",
	//     "https://www.googleapis.com/auth/firebase.readonly"
	//   ]
	// }

}

// method id "firebasehosting.sites.domains.list":

type SitesDomainsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the domains for the specified site.
func (r *SitesDomainsService) List(parent string) *SitesDomainsListCall {
	c := &SitesDomainsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The page size to
// return. Defaults to 50.
func (c *SitesDomainsListCall) PageSize(pageSize int64) *SitesDomainsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token from a previous request, if provided.
func (c *SitesDomainsListCall) PageToken(pageToken string) *SitesDomainsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesDomainsListCall) Fields(s ...googleapi.Field) *SitesDomainsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SitesDomainsListCall) IfNoneMatch(entityTag string) *SitesDomainsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesDomainsListCall) Context(ctx context.Context) *SitesDomainsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesDomainsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesDomainsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/domains")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.domains.list" call.
// Exactly one of *ListDomainsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListDomainsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SitesDomainsListCall) Do(opts ...googleapi.CallOption) (*ListDomainsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListDomainsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the domains for the specified site.",
	//   "flatPath": "v1beta1/sites/{sitesId}/domains",
	//   "httpMethod": "GET",
	//   "id": "firebasehosting.sites.domains.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The page size to return. Defaults to 50.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token from a previous request, if provided.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The parent for which to list domains, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e\u003c/code\u003e",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/domains",
	//   "response": {
	//     "$ref": "ListDomainsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/firebase",
	//     "https://www.googleapis.com/auth/firebase.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SitesDomainsListCall) Pages(ctx context.Context, f func(*ListDomainsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "firebasehosting.sites.domains.update":

type SitesDomainsUpdateCall struct {
	s          *Service
	name       string
	domain     *Domain
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates the specified domain mapping, creating the mapping as
// if it does
// not exist.
func (r *SitesDomainsService) Update(name string, domain *Domain) *SitesDomainsUpdateCall {
	c := &SitesDomainsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.domain = domain
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesDomainsUpdateCall) Fields(s ...googleapi.Field) *SitesDomainsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesDomainsUpdateCall) Context(ctx context.Context) *SitesDomainsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesDomainsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesDomainsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.domain)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.domains.update" call.
// Exactly one of *Domain or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Domain.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SitesDomainsUpdateCall) Do(opts ...googleapi.CallOption) (*Domain, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Domain{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified domain mapping, creating the mapping as if it does\nnot exist.",
	//   "flatPath": "v1beta1/sites/{sitesId}/domains/{domainsId}",
	//   "httpMethod": "PUT",
	//   "id": "firebasehosting.sites.domains.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the domain association to update or create, if an\nassociation doesn't already exist.",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+/domains/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "Domain"
	//   },
	//   "response": {
	//     "$ref": "Domain"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/firebase"
	//   ]
	// }

}

// method id "firebasehosting.sites.releases.create":

type SitesReleasesCreateCall struct {
	s          *Service
	parent     string
	release    *Release
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new release which makes the content of the
// specified version
// actively display on the site.
func (r *SitesReleasesService) Create(parent string, release *Release) *SitesReleasesCreateCall {
	c := &SitesReleasesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.release = release
	return c
}

// VersionName sets the optional parameter "versionName": The unique
// identifier for a version, in the
// format:
// <code>/sites/<var>site-name</var>/versions/<var>versionID</var
// ></code>
// The <var>site-name</var> in this version identifier must match
// the
// <var>site-name</var> in the `parent` parameter.
// <br>
// <br>This query parameter must be empty if the `type` field in
// the
// request body is `SITE_DISABLE`.
func (c *SitesReleasesCreateCall) VersionName(versionName string) *SitesReleasesCreateCall {
	c.urlParams_.Set("versionName", versionName)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesReleasesCreateCall) Fields(s ...googleapi.Field) *SitesReleasesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesReleasesCreateCall) Context(ctx context.Context) *SitesReleasesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesReleasesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesReleasesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.release)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/releases")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.releases.create" call.
// Exactly one of *Release or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Release.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SitesReleasesCreateCall) Do(opts ...googleapi.CallOption) (*Release, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Release{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new release which makes the content of the specified version\nactively display on the site.",
	//   "flatPath": "v1beta1/sites/{sitesId}/releases",
	//   "httpMethod": "POST",
	//   "id": "firebasehosting.sites.releases.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The site that the release belongs to, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e\u003c/code\u003e",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "versionName": {
	//       "description": "The unique identifier for a version, in the format:\n\u003ccode\u003e/sites/\u003cvar\u003esite-name\u003c/var\u003e/versions/\u003cvar\u003eversionID\u003c/var\u003e\u003c/code\u003e\nThe \u003cvar\u003esite-name\u003c/var\u003e in this version identifier must match the\n\u003cvar\u003esite-name\u003c/var\u003e in the `parent` parameter.\n\u003cbr\u003e\n\u003cbr\u003eThis query parameter must be empty if the `type` field in the\nrequest body is `SITE_DISABLE`.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/releases",
	//   "request": {
	//     "$ref": "Release"
	//   },
	//   "response": {
	//     "$ref": "Release"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/firebase"
	//   ]
	// }

}

// method id "firebasehosting.sites.releases.list":

type SitesReleasesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the releases that have been created on the specified
// site.
func (r *SitesReleasesService) List(parent string) *SitesReleasesListCall {
	c := &SitesReleasesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The page size to
// return. Defaults to 100.
func (c *SitesReleasesListCall) PageSize(pageSize int64) *SitesReleasesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token from a previous request, if provided.
func (c *SitesReleasesListCall) PageToken(pageToken string) *SitesReleasesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesReleasesListCall) Fields(s ...googleapi.Field) *SitesReleasesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SitesReleasesListCall) IfNoneMatch(entityTag string) *SitesReleasesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesReleasesListCall) Context(ctx context.Context) *SitesReleasesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesReleasesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesReleasesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/releases")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.releases.list" call.
// Exactly one of *ListReleasesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListReleasesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SitesReleasesListCall) Do(opts ...googleapi.CallOption) (*ListReleasesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListReleasesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the releases that have been created on the specified site.",
	//   "flatPath": "v1beta1/sites/{sitesId}/releases",
	//   "httpMethod": "GET",
	//   "id": "firebasehosting.sites.releases.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The page size to return. Defaults to 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token from a previous request, if provided.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The parent for which to list files, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e\u003c/code\u003e",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/releases",
	//   "response": {
	//     "$ref": "ListReleasesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/firebase",
	//     "https://www.googleapis.com/auth/firebase.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SitesReleasesListCall) Pages(ctx context.Context, f func(*ListReleasesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "firebasehosting.sites.versions.create":

type SitesVersionsCreateCall struct {
	s          *Service
	parent     string
	version    *Version
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new version for a site.
func (r *SitesVersionsService) Create(parent string, version *Version) *SitesVersionsCreateCall {
	c := &SitesVersionsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.version = version
	return c
}

// SizeBytes sets the optional parameter "sizeBytes": The self-reported
// size of the version. This value is used for a pre-emptive
// quota check for legacy version uploads.
func (c *SitesVersionsCreateCall) SizeBytes(sizeBytes int64) *SitesVersionsCreateCall {
	c.urlParams_.Set("sizeBytes", fmt.Sprint(sizeBytes))
	return c
}

// VersionId sets the optional parameter "versionId": A unique id for
// the new version. This is only specified for legacy version
// creations.
func (c *SitesVersionsCreateCall) VersionId(versionId string) *SitesVersionsCreateCall {
	c.urlParams_.Set("versionId", versionId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesVersionsCreateCall) Fields(s ...googleapi.Field) *SitesVersionsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesVersionsCreateCall) Context(ctx context.Context) *SitesVersionsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesVersionsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesVersionsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.version)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/versions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.versions.create" call.
// Exactly one of *Version or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Version.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SitesVersionsCreateCall) Do(opts ...googleapi.CallOption) (*Version, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Version{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new version for a site.",
	//   "flatPath": "v1beta1/sites/{sitesId}/versions",
	//   "httpMethod": "POST",
	//   "id": "firebasehosting.sites.versions.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The parent to create the version for, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e\u003c/code\u003e",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sizeBytes": {
	//       "description": "The self-reported size of the version. This value is used for a pre-emptive\nquota check for legacy version uploads.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "versionId": {
	//       "description": "A unique id for the new version. This is only specified for legacy version\ncreations.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/versions",
	//   "request": {
	//     "$ref": "Version"
	//   },
	//   "response": {
	//     "$ref": "Version"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/firebase"
	//   ]
	// }

}

// method id "firebasehosting.sites.versions.delete":

type SitesVersionsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified version.
func (r *SitesVersionsService) Delete(name string) *SitesVersionsDeleteCall {
	c := &SitesVersionsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesVersionsDeleteCall) Fields(s ...googleapi.Field) *SitesVersionsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesVersionsDeleteCall) Context(ctx context.Context) *SitesVersionsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesVersionsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesVersionsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.versions.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SitesVersionsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified version.",
	//   "flatPath": "v1beta1/sites/{sitesId}/versions/{versionsId}",
	//   "httpMethod": "DELETE",
	//   "id": "firebasehosting.sites.versions.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the version to be deleted, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e/versions/\u003cvar\u003eversionID\u003c/var\u003e\u003c/code\u003e",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+/versions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/firebase"
	//   ]
	// }

}

// method id "firebasehosting.sites.versions.patch":

type SitesVersionsPatchCall struct {
	s          *Service
	nameid     string
	version    *Version
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates the specified metadata for a version. Note that this
// method will
// fail with `FAILED_PRECONDITION` in the event of an invalid
// state
// transition. The only valid transition for a version is currently from
// a
// `CREATED` status to a `FINALIZED` status.
// Use [`DeleteVersion`](../sites.versions/delete) to set the status of
// a
// version to `DELETED`.
func (r *SitesVersionsService) Patch(nameid string, version *Version) *SitesVersionsPatchCall {
	c := &SitesVersionsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.version = version
	return c
}

// UpdateMask sets the optional parameter "updateMask": A set of field
// names from your [version](../sites.versions) that you want
// to update.
// <br>A field will be overwritten if, and only if, it's in the
// mask.
// <br>If a mask is not provided then a default mask of
// only
// [`status`](../sites.versions#Version.FIELDS.status) will be used.
func (c *SitesVersionsPatchCall) UpdateMask(updateMask string) *SitesVersionsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesVersionsPatchCall) Fields(s ...googleapi.Field) *SitesVersionsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesVersionsPatchCall) Context(ctx context.Context) *SitesVersionsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesVersionsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesVersionsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.version)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.versions.patch" call.
// Exactly one of *Version or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Version.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SitesVersionsPatchCall) Do(opts ...googleapi.CallOption) (*Version, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Version{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified metadata for a version. Note that this method will\nfail with `FAILED_PRECONDITION` in the event of an invalid state\ntransition. The only valid transition for a version is currently from a\n`CREATED` status to a `FINALIZED` status.\nUse [`DeleteVersion`](../sites.versions/delete) to set the status of a\nversion to `DELETED`.",
	//   "flatPath": "v1beta1/sites/{sitesId}/versions/{versionsId}",
	//   "httpMethod": "PATCH",
	//   "id": "firebasehosting.sites.versions.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique identifier for a version, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e/versions/\u003cvar\u003eversionID\u003c/var\u003e\u003c/code\u003e\nThis name is provided in the response body when you call the\n[`CreateVersion`](../sites.versions/create) endpoint.",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+/versions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "A set of field names from your [version](../sites.versions) that you want\nto update.\n\u003cbr\u003eA field will be overwritten if, and only if, it's in the mask.\n\u003cbr\u003eIf a mask is not provided then a default mask of only\n[`status`](../sites.versions#Version.FIELDS.status) will be used.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "Version"
	//   },
	//   "response": {
	//     "$ref": "Version"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/firebase"
	//   ]
	// }

}

// method id "firebasehosting.sites.versions.populateFiles":

type SitesVersionsPopulateFilesCall struct {
	s                           *Service
	parent                      string
	populateversionfilesrequest *PopulateVersionFilesRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// PopulateFiles: Adds content files to a version.
func (r *SitesVersionsService) PopulateFiles(parent string, populateversionfilesrequest *PopulateVersionFilesRequest) *SitesVersionsPopulateFilesCall {
	c := &SitesVersionsPopulateFilesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.populateversionfilesrequest = populateversionfilesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesVersionsPopulateFilesCall) Fields(s ...googleapi.Field) *SitesVersionsPopulateFilesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesVersionsPopulateFilesCall) Context(ctx context.Context) *SitesVersionsPopulateFilesCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesVersionsPopulateFilesCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesVersionsPopulateFilesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.populateversionfilesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}:populateFiles")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.versions.populateFiles" call.
// Exactly one of *PopulateVersionFilesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *PopulateVersionFilesResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SitesVersionsPopulateFilesCall) Do(opts ...googleapi.CallOption) (*PopulateVersionFilesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PopulateVersionFilesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds content files to a version.",
	//   "flatPath": "v1beta1/sites/{sitesId}/versions/{versionsId}:populateFiles",
	//   "httpMethod": "POST",
	//   "id": "firebasehosting.sites.versions.populateFiles",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The version to add files to, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e/versions/\u003cvar\u003eversionID\u003c/var\u003e\u003c/code\u003e",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+/versions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}:populateFiles",
	//   "request": {
	//     "$ref": "PopulateVersionFilesRequest"
	//   },
	//   "response": {
	//     "$ref": "PopulateVersionFilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/firebase"
	//   ]
	// }

}

// method id "firebasehosting.sites.versions.files.list":

type SitesVersionsFilesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the remaining files to be uploaded for the specified
// version.
func (r *SitesVersionsFilesService) List(parent string) *SitesVersionsFilesListCall {
	c := &SitesVersionsFilesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The page size to
// return. Defaults to 1000.
func (c *SitesVersionsFilesListCall) PageSize(pageSize int64) *SitesVersionsFilesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token from a previous request, if provided. This will be
// the
// encoded version of a
// firebase.hosting.proto.metadata.ListFilesPageToken.
func (c *SitesVersionsFilesListCall) PageToken(pageToken string) *SitesVersionsFilesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Status sets the optional parameter "status": The type of files in the
// version that should be listed.
//
// Possible values:
//   "STATUS_UNSPECIFIED"
//   "EXPECTED"
//   "ACTIVE"
func (c *SitesVersionsFilesListCall) Status(status string) *SitesVersionsFilesListCall {
	c.urlParams_.Set("status", status)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesVersionsFilesListCall) Fields(s ...googleapi.Field) *SitesVersionsFilesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SitesVersionsFilesListCall) IfNoneMatch(entityTag string) *SitesVersionsFilesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesVersionsFilesListCall) Context(ctx context.Context) *SitesVersionsFilesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesVersionsFilesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesVersionsFilesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/files")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasehosting.sites.versions.files.list" call.
// Exactly one of *ListVersionFilesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListVersionFilesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SitesVersionsFilesListCall) Do(opts ...googleapi.CallOption) (*ListVersionFilesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListVersionFilesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the remaining files to be uploaded for the specified version.",
	//   "flatPath": "v1beta1/sites/{sitesId}/versions/{versionsId}/files",
	//   "httpMethod": "GET",
	//   "id": "firebasehosting.sites.versions.files.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The page size to return. Defaults to 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token from a previous request, if provided. This will be the\nencoded version of a firebase.hosting.proto.metadata.ListFilesPageToken.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The parent to list files for, in the format:\n\u003ccode\u003esites/\u003cvar\u003esite-name\u003c/var\u003e/versions/\u003cvar\u003eversionID\u003c/var\u003e\u003c/code\u003e",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+/versions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "status": {
	//       "description": "The type of files in the version that should be listed.",
	//       "enum": [
	//         "STATUS_UNSPECIFIED",
	//         "EXPECTED",
	//         "ACTIVE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/files",
	//   "response": {
	//     "$ref": "ListVersionFilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/firebase",
	//     "https://www.googleapis.com/auth/firebase.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SitesVersionsFilesListCall) Pages(ctx context.Context, f func(*ListVersionFilesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
