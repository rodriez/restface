package restface

import "net/http"

type ApiError struct {
	StatusCode int      `json:"status"`
	Message    string   `json:"message"`
	Errors     []string `json:"cause,omitempty"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func (e *ApiError) Code() int {
	return e.StatusCode
}

func (e *ApiError) Cause() []string {
	return e.Errors
}

//BadRequest - Return a bad_request error
func BadRequest(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusBadRequest,
		Message:    http.StatusText(http.StatusBadRequest),
		Errors:     err,
	}
}

//Forbidden - Return a forbidden error
func Forbidden(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusForbidden,
		Message:    http.StatusText(http.StatusForbidden),
		Errors:     err,
	}
}

//NotFound - Return a not_found error
func NotFound(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotFound,
		Message:    http.StatusText(http.StatusNotFound),
		Errors:     err,
	}
}

//PaymentRequired - Return a payment_required error
func PaymentRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusPaymentRequired,
		Message:    http.StatusText(http.StatusPaymentRequired),
		Errors:     err,
	}
}

//MethodNotAllowed - Return a method_not_allowed error
func MethodNotAllowed(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    http.StatusText(http.StatusMethodNotAllowed),
		Errors:     err,
	}
}

//NotAcceptable - Return a not_acceptable error
func NotAcceptable(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotAcceptable,
		Message:    http.StatusText(http.StatusNotAcceptable),
		Errors:     err,
	}
}

//ProxyAuthRequired - Return a proxy_auth_required error
func ProxyAuthRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusProxyAuthRequired,
		Message:    http.StatusText(http.StatusProxyAuthRequired),
		Errors:     err,
	}
}

//RequestTimeout - Return a request_time_out error
func RequestTimeout(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestTimeout,
		Message:    http.StatusText(http.StatusRequestTimeout),
		Errors:     err,
	}
}

//Conflict - Return a conflict error
func Conflict(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusConflict,
		Message:    http.StatusText(http.StatusConflict),
		Errors:     err,
	}
}

//Gone - Return a gone error
func Gone(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusGone,
		Message:    http.StatusText(http.StatusGone),
		Errors:     err,
	}
}

//LengthRequired - Return a length_required error
func LengthRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusLengthRequired,
		Message:    http.StatusText(http.StatusLengthRequired),
		Errors:     err,
	}
}

//PreconditionFailed - Return a precondition_failed error
func PreconditionFailed(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusPreconditionFailed,
		Message:    http.StatusText(http.StatusPreconditionFailed),
		Errors:     err,
	}
}

//RequestEntityTooLarge - Return a request_entity_too_large error
func RequestEntityTooLarge(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestEntityTooLarge,
		Message:    http.StatusText(http.StatusRequestEntityTooLarge),
		Errors:     err,
	}
}

//RequestURITooLong - Return a request_uri_too_long error
func RequestURITooLong(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestURITooLong,
		Message:    http.StatusText(http.StatusRequestURITooLong),
		Errors:     err,
	}
}

//UnsupportedMediaType - Return a unsupported_media_type error
func UnsupportedMediaType(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnsupportedMediaType,
		Message:    http.StatusText(http.StatusUnsupportedMediaType),
		Errors:     err,
	}
}

//RequestedRangeNotSatisfiable - Return a requested_range_not_satisfiable error
func RequestedRangeNotSatisfiable(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestedRangeNotSatisfiable,
		Message:    http.StatusText(http.StatusRequestedRangeNotSatisfiable),
		Errors:     err,
	}
}

//ExpectationFailed - Return a expectation_failed error
func ExpectationFailed(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusExpectationFailed,
		Message:    http.StatusText(http.StatusExpectationFailed),
		Errors:     err,
	}
}

//Teapot - Return a teapot error
func Teapot(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusTeapot,
		Message:    http.StatusText(http.StatusTeapot),
		Errors:     err,
	}
}

//MisdirectedRequest - Return a misdirected_request error
func MisdirectedRequest(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusMisdirectedRequest,
		Message:    http.StatusText(http.StatusMisdirectedRequest),
		Errors:     err,
	}
}

//UnprocessableEntity - Return a unprocessable_entity error
func UnprocessableEntity(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    http.StatusText(http.StatusUnprocessableEntity),
		Errors:     err,
	}
}

//Locked - Return a locked error
func Locked(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusLocked,
		Message:    http.StatusText(http.StatusLocked),
		Errors:     err,
	}
}

//FailedDependency - Return a failed_dependency error
func FailedDependency(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusFailedDependency,
		Message:    http.StatusText(http.StatusFailedDependency),
		Errors:     err,
	}
}

//TooEarly - Return a too_early error
func TooEarly(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusTooEarly,
		Message:    http.StatusText(http.StatusTooEarly),
		Errors:     err,
	}
}

//UpgradeRequired - Return a upgrade_required error
func UpgradeRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUpgradeRequired,
		Message:    http.StatusText(http.StatusUpgradeRequired),
		Errors:     err,
	}
}

//PreconditionRequired - Return a precondition_required error
func PreconditionRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusPreconditionRequired,
		Message:    http.StatusText(http.StatusPreconditionRequired),
		Errors:     err,
	}
}

//TooManyRequests - Return a too_many_requests error
func TooManyRequests(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusTooManyRequests,
		Message:    http.StatusText(http.StatusTooManyRequests),
		Errors:     err,
	}
}

//RequestHeaderFieldsTooLarge - Return a XX error
func RequestHeaderFieldsTooLarge(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestHeaderFieldsTooLarge,
		Message:    http.StatusText(http.StatusRequestHeaderFieldsTooLarge),
		Errors:     err,
	}
}

//UnavailableForLegalReasons - Return a unavailable_for_legal_reasons error
func UnavailableForLegalReasons(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnavailableForLegalReasons,
		Message:    http.StatusText(http.StatusUnavailableForLegalReasons),
		Errors:     err,
	}
}

//Unauthorized - Return a unauthorized error
func Unauthorized(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnauthorized,
		Message:    http.StatusText(http.StatusUnauthorized),
		Errors:     err,
	}
}

//InternalError - Return a internal_server error
func InternalError(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusInternalServerError,
		Message:    http.StatusText(http.StatusInternalServerError),
		Errors:     err,
	}
}

//NotImplemented - Return a not_implemented error
func NotImplemented(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotImplemented,
		Message:    http.StatusText(http.StatusNotImplemented),
		Errors:     err,
	}
}

//BadGateway - Return a bad_gateway error
func BadGateway(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusBadGateway,
		Message:    http.StatusText(http.StatusBadGateway),
		Errors:     err,
	}
}

//ServiceUnavailable - Return a service_unavailable error
func ServiceUnavailable(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusServiceUnavailable,
		Message:    http.StatusText(http.StatusServiceUnavailable),
		Errors:     err,
	}
}

//GatewayTimeout - Return a gateway_timeout error
func GatewayTimeout(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusGatewayTimeout,
		Message:    http.StatusText(http.StatusGatewayTimeout),
		Errors:     err,
	}
}

//HTTPVersionNotSupported - Return a http_version_not_supported error
func HTTPVersionNotSupported(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusHTTPVersionNotSupported,
		Message:    http.StatusText(http.StatusHTTPVersionNotSupported),
		Errors:     err,
	}
}

//VariantAlsoNegotiates - Return a variant_also_negotiates error
func VariantAlsoNegotiates(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusVariantAlsoNegotiates,
		Message:    http.StatusText(http.StatusVariantAlsoNegotiates),
		Errors:     err,
	}
}

//InsufficientStorage - Return a insufficient_storage error
func InsufficientStorage(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusInsufficientStorage,
		Message:    http.StatusText(http.StatusInsufficientStorage),
		Errors:     err,
	}
}

//LoopDetected - Return a loop_detected error
func LoopDetected(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusLoopDetected,
		Message:    http.StatusText(http.StatusLoopDetected),
		Errors:     err,
	}
}

//NotExtended - Return a not_extended error
func NotExtended(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotExtended,
		Message:    http.StatusText(http.StatusNotExtended),
		Errors:     err,
	}
}

//NetworkAuthenticationRequired - Return a network_authentication_required error
func NetworkAuthenticationRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNetworkAuthenticationRequired,
		Message:    http.StatusText(http.StatusNetworkAuthenticationRequired),
		Errors:     err,
	}
}
