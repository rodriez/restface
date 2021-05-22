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

func BadRequest(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusBadRequest,
		Message:    http.StatusText(http.StatusBadRequest),
		Errors:     err,
	}
}

func Forbidden(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusForbidden,
		Message:    http.StatusText(http.StatusForbidden),
		Errors:     err,
	}
}

func NotFound(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotFound,
		Message:    http.StatusText(http.StatusNotFound),
		Errors:     err,
	}
}

func PaymentRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusPaymentRequired,
		Message:    http.StatusText(http.StatusPaymentRequired),
		Errors:     err,
	}
}

func MethodNotAllowed(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    http.StatusText(http.StatusMethodNotAllowed),
		Errors:     err,
	}
}

func NotAcceptable(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotAcceptable,
		Message:    http.StatusText(http.StatusNotAcceptable),
		Errors:     err,
	}
}

func ProxyAuthRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusProxyAuthRequired,
		Message:    http.StatusText(http.StatusProxyAuthRequired),
		Errors:     err,
	}
}

func RequestTimeout(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestTimeout,
		Message:    http.StatusText(http.StatusRequestTimeout),
		Errors:     err,
	}
}

func Conflict(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusConflict,
		Message:    http.StatusText(http.StatusConflict),
		Errors:     err,
	}
}

func Gone(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusGone,
		Message:    http.StatusText(http.StatusGone),
		Errors:     err,
	}
}

func LengthRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusLengthRequired,
		Message:    http.StatusText(http.StatusLengthRequired),
		Errors:     err,
	}
}

func PreconditionFailed(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusPreconditionFailed,
		Message:    http.StatusText(http.StatusPreconditionFailed),
		Errors:     err,
	}
}

func RequestEntityTooLarge(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestEntityTooLarge,
		Message:    http.StatusText(http.StatusRequestEntityTooLarge),
		Errors:     err,
	}
}

func RequestURITooLong(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestURITooLong,
		Message:    http.StatusText(http.StatusRequestURITooLong),
		Errors:     err,
	}
}

func UnsupportedMediaType(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnsupportedMediaType,
		Message:    http.StatusText(http.StatusUnsupportedMediaType),
		Errors:     err,
	}
}

func RequestedRangeNotSatisfiable(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestedRangeNotSatisfiable,
		Message:    http.StatusText(http.StatusRequestedRangeNotSatisfiable),
		Errors:     err,
	}
}

func ExpectationFailed(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusExpectationFailed,
		Message:    http.StatusText(http.StatusExpectationFailed),
		Errors:     err,
	}
}

func Teapot(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusTeapot,
		Message:    http.StatusText(http.StatusTeapot),
		Errors:     err,
	}
}

func MisdirectedRequest(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusMisdirectedRequest,
		Message:    http.StatusText(http.StatusMisdirectedRequest),
		Errors:     err,
	}
}

func UnprocessableEntity(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    http.StatusText(http.StatusUnprocessableEntity),
		Errors:     err,
	}
}

func Locked(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusLocked,
		Message:    http.StatusText(http.StatusLocked),
		Errors:     err,
	}
}

func FailedDependency(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusFailedDependency,
		Message:    http.StatusText(http.StatusFailedDependency),
		Errors:     err,
	}
}

func TooEarly(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusTooEarly,
		Message:    http.StatusText(http.StatusTooEarly),
		Errors:     err,
	}
}

func UpgradeRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUpgradeRequired,
		Message:    http.StatusText(http.StatusUpgradeRequired),
		Errors:     err,
	}
}

func PreconditionRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusPreconditionRequired,
		Message:    http.StatusText(http.StatusPreconditionRequired),
		Errors:     err,
	}
}

func TooManyRequests(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusTooManyRequests,
		Message:    http.StatusText(http.StatusTooManyRequests),
		Errors:     err,
	}
}

func RequestHeaderFieldsTooLarge(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestHeaderFieldsTooLarge,
		Message:    http.StatusText(http.StatusRequestHeaderFieldsTooLarge),
		Errors:     err,
	}
}

func UnavailableForLegalReasons(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnavailableForLegalReasons,
		Message:    http.StatusText(http.StatusUnavailableForLegalReasons),
		Errors:     err,
	}
}

func Unauthorized(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnauthorized,
		Message:    http.StatusText(http.StatusUnauthorized),
		Errors:     err,
	}
}

func InternalError(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusInternalServerError,
		Message:    http.StatusText(http.StatusInternalServerError),
		Errors:     err,
	}
}

func NotImplemented(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotImplemented,
		Message:    http.StatusText(http.StatusNotImplemented),
		Errors:     err,
	}
}

func BadGateway(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusBadGateway,
		Message:    http.StatusText(http.StatusBadGateway),
		Errors:     err,
	}
}

func ServiceUnavailable(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusServiceUnavailable,
		Message:    http.StatusText(http.StatusServiceUnavailable),
		Errors:     err,
	}
}

func GatewayTimeout(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusGatewayTimeout,
		Message:    http.StatusText(http.StatusGatewayTimeout),
		Errors:     err,
	}
}

func HTTPVersionNotSupported(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusHTTPVersionNotSupported,
		Message:    http.StatusText(http.StatusHTTPVersionNotSupported),
		Errors:     err,
	}
}

func VariantAlsoNegotiates(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusVariantAlsoNegotiates,
		Message:    http.StatusText(http.StatusVariantAlsoNegotiates),
		Errors:     err,
	}
}

func InsufficientStorage(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusInsufficientStorage,
		Message:    http.StatusText(http.StatusInsufficientStorage),
		Errors:     err,
	}
}

func LoopDetected(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusLoopDetected,
		Message:    http.StatusText(http.StatusLoopDetected),
		Errors:     err,
	}
}

func NotExtended(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotExtended,
		Message:    http.StatusText(http.StatusNotExtended),
		Errors:     err,
	}
}

func NetworkAuthenticationRequired(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNetworkAuthenticationRequired,
		Message:    http.StatusText(http.StatusNetworkAuthenticationRequired),
		Errors:     err,
	}
}
