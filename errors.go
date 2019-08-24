package rfc7807

// Global domain errors from Google API.
// Some duplicate namings have been merged into one or renaming avoid of confilits, see NOTE.
// See: https://developers.google.com/youtube/v3/docs/core_errors
const (
	// MOVED_PERMANENTLY (301)
	MovedPermanently = "movedPermanently"

	// SEE_OTHER (303)
	SeeOther              = "seeOther"
	MediaDownloadRedirect = "mediaDownloadRedirect"

	// NOT_MODIFIED (304)
	NotModified = "notModified"

	// TEMPORARY_REDIRECT (307)
	TemporaryRedirect = "temporaryRedirect"

	// BAD_REQUEST (400)
	BadRequest                  = "badRequest"
	BadBinaryDomainRequest      = "badBinaryDomainRequest"
	BadContent                  = "badContent"
	BadLockedDomainRequest      = "badLockedDomainRequest"
	CorsRequestWithXOrigin      = "corsRequestWithXOrigin"
	EndpointConstraintMismatch  = "endpointConstraintMismatch"
	Invalid                     = "invalid"
	InvalidAltValue             = "invalidAltValue"
	InvalidHeader               = "invalidHeader"
	InvalidParameter            = "invalidParameter"
	InvalidQuery                = "invalidQuery"
	KeyExpired                  = "keyExpired"
	KeyInvalid                  = "keyInvalid"
	LockedDomainCreationFailure = "lockedDomainCreationFailure"
	NotDownload                 = "notDownload"
	NotUpload                   = "notUpload"
	ParseError                  = "parseError"
	Required                    = "required"
	TooManyParts                = "tooManyParts"
	UnknownAPI                  = "unknownApi"
	UnsupportedMediaProtocol    = "unsupportedMediaProtocol"
	UnsupportedOutputFormat     = "unsupportedOutputFormat"
	WrongURLForUpload           = "wrongUrlForUpload"

	// UNAUTHORIZED (401)
	Unauthorized        = "unauthorized"
	AuthError           = "authError"
	Expired             = "expired"
	LockedDomainExpired = "lockedDomainExpired"
	LoginRequired       = "loginRequired"
	// Required            = "required" // Note: Conflict with 400 required, so rename above.

	// PAYMENT_REQUIRED (402)
	DailyLimitExceeded402 = "dailyLimitExceeded402"
	QuotaExceeded402      = "quotaExceeded402"
	User402               = "user402"

	// FORBIDDEN (403)
	Forbidden                        = "forbidden"
	AccessNotConfigured              = "accessNotConfigured"
	AccountDeleted                   = "accountDeleted"
	AccountDisabled                  = "accountDisabled"
	AccountUnverified                = "accountUnverified"
	ConcurrentLimitExceeded          = "concurrentLimitExceeded"
	DailyLimitExceeded               = "dailyLimitExceeded"
	DailyLimitExceededUnreg          = "dailyLimitExceededUnreg"
	DownloadServiceForbidden         = "downloadServiceForbidden"
	InsufficientAudience             = "insufficientAudience"
	InsufficientAuthorizedParty      = "insufficientAuthorizedParty"
	InsufficientPermissions          = "insufficientPermissions"
	LimitExceeded                    = "limitExceeded"
	LockedDomainForbidden            = "lockedDomainForbidden"
	QuotaExceeded                    = "quotaExceeded"
	RateLimitExceeded                = "rateLimitExceeded"
	RateLimitExceededUnreg           = "rateLimitExceededUnreg"
	ResponseTooLarge                 = "responseTooLarge"
	ServingLimitExceeded             = "servingLimitExceeded"
	SSLRequired                      = "sslRequired"
	UnknownAuth                      = "unknownAuth"
	UserRateLimitExceeded            = "userRateLimitExceeded"
	UserRateLimitExceededUnreg       = "userRateLimitExceededUnreg"
	VariableTermExpiredDailyExceeded = "variableTermExpiredDailyExceeded"
	VariableTermLimitExceeded        = "variableTermLimitExceeded"
	// AccessNotConfigured             = "accessNotConfigured" // Note: no needs three, just one meaning.
	// AccessNotConfigured             = "accessNotConfigured" // Note: same above.
	// DailyLimitExceeded              = "dailyLimitExceeded" // Note: no needs too, just one meaning.

	// NOT_FOUND (404)
	NotFound            = "notFound"
	UnsupportedProtocol = "unsupportedProtocol"
	// NotFound            = "notFound" // Note: no need two notfound, same meaning.

	// METHOD_NOT_ALLOWED (405)
	HTTPMethodNotAllowed = "httpMethodNotAllowed"

	// CONFLICT (409)
	Conflict  = "conflict"
	Duplicate = "duplicate"

	// GONE (410)
	Deleted = "deleted"

	// PRECONDITION_FAILED (412)
	ConditionNotMet = "conditionNotMet"

	// REQUEST_ENTITY_TOO_LARGE (413)
	BackendRequestTooLarge = "backendRequestTooLarge"
	BatchSizeTooLarge      = "batchSizeTooLarge"
	UploadTooLarge         = "uploadTooLarge"

	// REQUESTED_RANGE_NOT_SATISFIABLE (416)
	RequestedRangeNotSatisfiable = "requestedRangeNotSatisfiable"

	// EXPECTATION_FAILED (417)
	ExpectationFailed = "expectationFailed"

	// PRECONDITION_REQUIRED (428)
	PreconditionRequired = "preconditionRequired"

	// TOO_MANY_REQUESTS (429)
	TooManyRequests = "tooManyRequests"
	// RateLimitExceeded = "rateLimitExceeded" // Conflict, use above.

	// INTERNAL_SERVER_ERROR (500)
	InternalError = "internalError"

	// NOT_IMPLEMENTED (501)
	NotImplemented    = "notImplemented"
	UnsupportedMethod = "unsupportedMethod"

	// SERVICE_UNAVAILABLE (503)
	BackendError        = "backendError"
	BackendNotConnected = "backendNotConnected"
	NotReady            = "notReady"
)

var mapping = map[string]struct {
	status int
	title  string
}{
	MovedPermanently: {301, "This request and future requests for the same operation have to be sent to the URL specified in the Location header of this response instead of to the URL to which this request was sent."},

	SeeOther:              {303, "Your request was processed successfully. To obtain your response, send a GET request to the URL specified in the Location header."},
	MediaDownloadRedirect: {303, "Your request was processed successfully. To obtain your response, send a GET request to the URL specified in the Location header."},

	NotModified: {304, "The condition set for an If-None-Match header was not met. This response indicates that the requested document has not been modified and that a cached response should be retrieved. Check the value of the If-None-Match HTTP request header."},

	TemporaryRedirect: {307, "To have your request processed, resend it to the URL specified in the Location header of this response."},

	BadRequest:                  {400, "The API request is invalid or improperly formed. Consequently, the API server could not understand the request."},
	BadBinaryDomainRequest:      {400, "The binary domain request is invalid."},
	BadContent:                  {400, "The content type of the request data or the content type of a part of a multipart request is not supported."},
	BadLockedDomainRequest:      {400, "The locked domain request is invalid."},
	CorsRequestWithXOrigin:      {400, "The CORS request contains an XD3 X-Origin header, which is indicative of a bad CORS request."},
	EndpointConstraintMismatch:  {400, "The request failed because it did not match the specified API. Check the value of the URL path to make sure it is correct."},
	Invalid:                     {400, "The request failed because it contained an invalid value. The value could be a parameter value, a header value, or a property value."},
	InvalidAltValue:             {400, "The alt parameter value specifies an unknown output format."},
	InvalidHeader:               {400, "The request failed because it contained an invalid header."},
	InvalidParameter:            {400, "The request failed because it contained an invalid parameter or parameter value. Review the API documentation to determine which parameters are valid for your request."},
	InvalidQuery:                {400, "The request is invalid. Check the API documentation to determine what parameters are supported for the request and to see if the request contains an invalid combination of parameters or an invalid parameter value. Check the value of the q request parameter."},
	KeyExpired:                  {400, "The API key provided in the request expired, which means the API server is unable to check the quota limit for the application making the request."},
	KeyInvalid:                  {400, "The API key provided in the request is invalid, which means the API server is unable to check the quota limit for the application making the request."},
	LockedDomainCreationFailure: {400, "The OAuth token was received in the query string, which this API forbids for response formats other than JSON or XML. If possible, try sending the OAuth token in the Authorization header instead."},
	NotDownload:                 {400, "Only media downloads requests can be sent to /download/* URL paths. Resend the request to the same path, but without the /download prefix."},
	NotUpload:                   {400, "The request failed because it is not an upload request, and only upload requests can be sent to /upload/* URIs. Try resending the request to the same path, but without the /upload prefix."},
	ParseError:                  {400, "The API server cannot parse the request body."},
	Required:                    {400, "The API request is missing required information. The required information could be a parameter or resource property."},
	TooManyParts:                {400, "The multipart request failed because it contains too many parts"},
	UnknownAPI:                  {400, "The API that the request is calling is not recognized."},
	UnsupportedMediaProtocol:    {400, "The client is using an unsupported media protocol."},
	UnsupportedOutputFormat:     {400, "The alt parameter value specifies an output format that is not supported for this service. Check the value of the alt request parameter."},
	WrongURLForUpload:           {400, "The request is an upload request, but it failed because it was not sent to the proper URI. Upload requests must be sent to URIs that contain the /upload/* prefix. Try resending the request to the same path, but with the /upload prefix."},

	Unauthorized:        {401, "The user is not authorized to make the request."},
	AuthError:           {401, "The authorization credentials provided for the request are invalid. Check the value of the Authorization HTTP request header."},
	Expired:             {401, "Session Expired. Check the value of the Authorization HTTP request header."},
	LockedDomainExpired: {401, "The request failed because a previously valid locked domain has expired."},
	LoginRequired:       {401, "The user must be logged in to make this API request. Check the value of the Authorization HTTP request header."},

	DailyLimitExceeded402: {402, "A daily budget limit set by the developer has been reached."},
	QuotaExceeded402:      {402, "The requested operation requires more resources than the quota allows. Payment is required to complete the operation."},
	User402:               {402, "The requested operation requires some kind of payment from the authenticated user."},

	Forbidden:                        {403, "The requested operation is forbidden and cannot be completed."},
	AccessNotConfigured:              {403, "Your project is not configured to access this API. Either blocked due to abuse or marked for deletion."},
	AccountDeleted:                   {403, "The user account associated with the request's authorization credentials has been deleted. Check the value of the Authorization HTTP request header."},
	AccountDisabled:                  {403, "The user account associated with the request's authorization credentials has been disabled. Check the value of the Authorization HTTP request header."},
	AccountUnverified:                {403, "The email address for the user making the request has not been verified. Check the value of the Authorization HTTP request header."},
	ConcurrentLimitExceeded:          {403, "The request failed because a concurrent usage limit has been reached."},
	DailyLimitExceeded:               {403, "A daily quota limit for the API has been reached."},
	DailyLimitExceededUnreg:          {403, "The request failed because a daily limit for unauthenticated API use has been hit. Continued use of the API requires signup."},
	DownloadServiceForbidden:         {403, "The API does not support a download service."},
	InsufficientAudience:             {403, "The request cannot be completed for this audience."},
	InsufficientAuthorizedParty:      {403, "The request cannot be completed for this application."},
	InsufficientPermissions:          {403, "The authenticated user does not have sufficient permissions to execute this request."},
	LimitExceeded:                    {403, "The request cannot be completed due to access or rate limitations."},
	LockedDomainForbidden:            {403, "This API does not support locked domains."},
	QuotaExceeded:                    {403, "The requested operation requires more resources than the quota allows."},
	RateLimitExceeded:                {403, "Too many requests have been sent within a given time span."},
	RateLimitExceededUnreg:           {403, "A rate limit has been exceeded and you must register your application to be able to continue calling the API. Please sign up."},
	ResponseTooLarge:                 {403, "The requested resource is too large to return."},
	ServingLimitExceeded:             {403, "The overall rate limit specified for the API has already been reached."},
	SSLRequired:                      {403, "SSL is required to perform this operation."},
	UnknownAuth:                      {403, "The API server does not recognize the authorization scheme used for the request. Check the value of the Authorization HTTP request header."},
	UserRateLimitExceeded:            {403, "The request failed because a per-user rate limit has been reached."},
	UserRateLimitExceededUnreg:       {403, "The request failed because a per-user rate limit has been reached, and the client developer was not identified in the request."},
	VariableTermExpiredDailyExceeded: {403, "The request failed because a variable term quota expired and a daily limit was reached."},
	VariableTermLimitExceeded:        {403, "The request failed because a variable term quota limit was reached."},
	// AccessNotConfigured:             {403, "The project has been blocked due to abuse."},
	// AccessNotConfigured:             {403, "The project has been marked for deletion."},
	// DailyLimitExceeded:              {403, "The daily quota limit has been reached, and the project has been blocked due to abuse."},

	NotFound:            {404, "The requested operation failed because a resource associated with the request could not be found."},
	UnsupportedProtocol: {404, "The protocol used in the request is not supported."},
	// NotFound:           {404, "A resource associated with the request could not be found. If you have not used this API in the last two weeks, please re-deploy the App Engine app and try calling it again."},

	HTTPMethodNotAllowed: {405, "The HTTP method associated with the request is not supported."},

	Conflict:  {409, "The API request cannot be completed because the requested operation would conflict with an existing item."},
	Duplicate: {409, "The requested operation failed because it tried to create a resource that already exists."},

	Deleted: {410, "The request failed because the resource associated with the request has been deleted."},

	ConditionNotMet: {412, "The condition set in the request's If-Match or If-None-Match HTTP request header was not met."},

	BackendRequestTooLarge: {413, "The request is too large."},
	BatchSizeTooLarge:      {413, "The batch request contains too many elements."},
	UploadTooLarge:         {413, "The request failed because the data sent in the request is too large."},

	RequestedRangeNotSatisfiable: {416, "The request specified a range that cannot be satisfied."},

	ExpectationFailed: {417, "A client expectation cannot be met by the server."},

	PreconditionRequired: {428, "The request requires a precondition that is not provided. For this request to succeed, you need to provide either an If-Match or If-None-Match header with the request."},

	TooManyRequests: {429, "Too many requests have been sent within a given time span."},

	InternalError: {500, "The request failed due to an internal error."},

	NotImplemented:    {501, "The requested operation has not been implemented."},
	UnsupportedMethod: {501, "The request failed because it is trying to execute an unknown method or operation."},

	BackendError:        {503, "A backend error occurred."},
	BackendNotConnected: {503, "The request failed due to a connection error."},
	NotReady:            {503, "The API server is not ready to accept requests."},
}

// ExplainError descripts the typo along side with an HTTP status code.
// Last return value false when no predefined explaination found.
func ExplainError(typo string) (int, string, bool) {
	v, ok := mapping[typo]
	return v.status, v.title, ok
}
