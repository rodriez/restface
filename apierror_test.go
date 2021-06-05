package restface_test

import (
	"net/http"
	"testing"

	"github.com/rodriez/restface"
)

func TestConstructors(t *testing.T) {
	tests := []struct {
		name          string
		err           *restface.ApiError
		expectedCode  int
		expectedError string
	}{
		{
			name:          "BadRequest",
			err:           restface.BadRequest(),
			expectedCode:  http.StatusBadRequest,
			expectedError: http.StatusText(http.StatusBadRequest),
		},
		{
			name:          "Forbidden",
			err:           restface.Forbidden(),
			expectedCode:  http.StatusForbidden,
			expectedError: http.StatusText(http.StatusForbidden),
		},
		{
			name:          "NotFound",
			err:           restface.NotFound(),
			expectedCode:  http.StatusNotFound,
			expectedError: http.StatusText(http.StatusNotFound),
		},
		{
			name:          "PaymentRequired",
			err:           restface.PaymentRequired(),
			expectedCode:  http.StatusPaymentRequired,
			expectedError: http.StatusText(http.StatusPaymentRequired),
		},
		{
			name:          "MethodNotAllowed",
			err:           restface.MethodNotAllowed(),
			expectedCode:  http.StatusMethodNotAllowed,
			expectedError: http.StatusText(http.StatusMethodNotAllowed),
		},
		{
			name:          "NotAcceptable",
			err:           restface.NotAcceptable(),
			expectedCode:  http.StatusNotAcceptable,
			expectedError: http.StatusText(http.StatusNotAcceptable),
		},
		{
			name:          "ProxyAuthRequired",
			err:           restface.ProxyAuthRequired(),
			expectedCode:  http.StatusProxyAuthRequired,
			expectedError: http.StatusText(http.StatusProxyAuthRequired),
		},
		{
			name:          "RequestTimeout",
			err:           restface.RequestTimeout(),
			expectedCode:  http.StatusRequestTimeout,
			expectedError: http.StatusText(http.StatusRequestTimeout),
		},
		{
			name:          "Conflict",
			err:           restface.Conflict(),
			expectedCode:  http.StatusConflict,
			expectedError: http.StatusText(http.StatusConflict),
		},
		{
			name:          "Gone",
			err:           restface.Gone(),
			expectedCode:  http.StatusGone,
			expectedError: http.StatusText(http.StatusGone),
		},
		{
			name:          "LengthRequired",
			err:           restface.LengthRequired(),
			expectedCode:  http.StatusLengthRequired,
			expectedError: http.StatusText(http.StatusLengthRequired),
		},
		{
			name:          "PreconditionFailed",
			err:           restface.PreconditionFailed(),
			expectedCode:  http.StatusPreconditionFailed,
			expectedError: http.StatusText(http.StatusPreconditionFailed),
		},
		{
			name:          "RequestEntityTooLarge",
			err:           restface.RequestEntityTooLarge(),
			expectedCode:  http.StatusRequestEntityTooLarge,
			expectedError: http.StatusText(http.StatusRequestEntityTooLarge),
		},
		{
			name:          "RequestURITooLong",
			err:           restface.RequestURITooLong(),
			expectedCode:  http.StatusRequestURITooLong,
			expectedError: http.StatusText(http.StatusRequestURITooLong),
		},
		{
			name:          "UnsupportedMediaType",
			err:           restface.UnsupportedMediaType(),
			expectedCode:  http.StatusUnsupportedMediaType,
			expectedError: http.StatusText(http.StatusUnsupportedMediaType),
		},
		{
			name:          "RequestedRangeNotSatisfiable",
			err:           restface.RequestedRangeNotSatisfiable(),
			expectedCode:  http.StatusRequestedRangeNotSatisfiable,
			expectedError: http.StatusText(http.StatusRequestedRangeNotSatisfiable),
		},
		{
			name:          "ExpectationFailed",
			err:           restface.ExpectationFailed(),
			expectedCode:  http.StatusExpectationFailed,
			expectedError: http.StatusText(http.StatusExpectationFailed),
		},
		{
			name:          "Teapot",
			err:           restface.Teapot(),
			expectedCode:  http.StatusTeapot,
			expectedError: http.StatusText(http.StatusTeapot),
		},
		{
			name:          "MisdirectedRequest",
			err:           restface.MisdirectedRequest(),
			expectedCode:  http.StatusMisdirectedRequest,
			expectedError: http.StatusText(http.StatusMisdirectedRequest),
		},
		{
			name:          "UnprocessableEntity",
			err:           restface.UnprocessableEntity(),
			expectedCode:  http.StatusUnprocessableEntity,
			expectedError: http.StatusText(http.StatusUnprocessableEntity),
		},
		{
			name:          "Locked",
			err:           restface.Locked(),
			expectedCode:  http.StatusLocked,
			expectedError: http.StatusText(http.StatusLocked),
		},
		{
			name:          "FailedDependency",
			err:           restface.FailedDependency(),
			expectedCode:  http.StatusFailedDependency,
			expectedError: http.StatusText(http.StatusFailedDependency),
		},
		{
			name:          "TooEarly",
			err:           restface.TooEarly(),
			expectedCode:  http.StatusTooEarly,
			expectedError: http.StatusText(http.StatusTooEarly),
		},
		{
			name:          "UpgradeRequired",
			err:           restface.UpgradeRequired(),
			expectedCode:  http.StatusUpgradeRequired,
			expectedError: http.StatusText(http.StatusUpgradeRequired),
		},
		{
			name:          "PreconditionRequired",
			err:           restface.PreconditionRequired(),
			expectedCode:  http.StatusPreconditionRequired,
			expectedError: http.StatusText(http.StatusPreconditionRequired),
		},
		{
			name:          "TooManyRequests",
			err:           restface.TooManyRequests(),
			expectedCode:  http.StatusTooManyRequests,
			expectedError: http.StatusText(http.StatusTooManyRequests),
		},
		{
			name:          "RequestHeaderFieldsTooLarge",
			err:           restface.RequestHeaderFieldsTooLarge(),
			expectedCode:  http.StatusRequestHeaderFieldsTooLarge,
			expectedError: http.StatusText(http.StatusRequestHeaderFieldsTooLarge),
		},
		{
			name:          "UnavailableForLegalReasons",
			err:           restface.UnavailableForLegalReasons(),
			expectedCode:  http.StatusUnavailableForLegalReasons,
			expectedError: http.StatusText(http.StatusUnavailableForLegalReasons),
		},
		{
			name:          "Unauthorized",
			err:           restface.Unauthorized(),
			expectedCode:  http.StatusUnauthorized,
			expectedError: http.StatusText(http.StatusUnauthorized),
		},
		{
			name:          "InternalError",
			err:           restface.InternalError(),
			expectedCode:  http.StatusInternalServerError,
			expectedError: http.StatusText(http.StatusInternalServerError),
		},
		{
			name:          "NotImplemented",
			err:           restface.NotImplemented(),
			expectedCode:  http.StatusNotImplemented,
			expectedError: http.StatusText(http.StatusNotImplemented),
		},
		{
			name:          "BadGateway",
			err:           restface.BadGateway(),
			expectedCode:  http.StatusBadGateway,
			expectedError: http.StatusText(http.StatusBadGateway),
		},
		{
			name:          "ServiceUnavailable",
			err:           restface.ServiceUnavailable(),
			expectedCode:  http.StatusServiceUnavailable,
			expectedError: http.StatusText(http.StatusServiceUnavailable),
		},
		{
			name:          "GatewayTimeout",
			err:           restface.GatewayTimeout(),
			expectedCode:  http.StatusGatewayTimeout,
			expectedError: http.StatusText(http.StatusGatewayTimeout),
		},
		{
			name:          "HTTPVersionNotSupported",
			err:           restface.HTTPVersionNotSupported(),
			expectedCode:  http.StatusHTTPVersionNotSupported,
			expectedError: http.StatusText(http.StatusHTTPVersionNotSupported),
		},
		{
			name:          "VariantAlsoNegotiates",
			err:           restface.VariantAlsoNegotiates(),
			expectedCode:  http.StatusVariantAlsoNegotiates,
			expectedError: http.StatusText(http.StatusVariantAlsoNegotiates),
		},
		{
			name:          "InsufficientStorage",
			err:           restface.InsufficientStorage(),
			expectedCode:  http.StatusInsufficientStorage,
			expectedError: http.StatusText(http.StatusInsufficientStorage),
		},
		{
			name:          "LoopDetected",
			err:           restface.LoopDetected(),
			expectedCode:  http.StatusLoopDetected,
			expectedError: http.StatusText(http.StatusLoopDetected),
		},
		{
			name:          "NotExtended",
			err:           restface.NotExtended(),
			expectedCode:  http.StatusNotExtended,
			expectedError: http.StatusText(http.StatusNotExtended),
		},
		{
			name:          "NetworkAuthenticationRequired",
			err:           restface.NetworkAuthenticationRequired(),
			expectedCode:  http.StatusNetworkAuthenticationRequired,
			expectedError: http.StatusText(http.StatusNetworkAuthenticationRequired),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Code() != tt.expectedCode || tt.expectedError != tt.err.Error() {
				t.Errorf("Expected code %d and error %s but get %d, %s", tt.expectedCode, tt.expectedError, tt.err.Code(), tt.err.Error())
			}
		})
	}
}
