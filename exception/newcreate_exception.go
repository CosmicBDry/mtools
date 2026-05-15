package exception

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewApiException(namespace string, code int, reason string, format string, a ...interface{}) ApiException {

	e := new(exception)

	e.Namespace = namespace

	//code=-1为了排除零值
	if code == 0 {
		code = -1

	}

	if code/100 >= 1 && code/100 <= 5 {
		e.Httpcode = code
	} else {
		e.Httpcode = InternalServerError
	}

	e.Reason = reason

	e.Message = fmt.Sprintf(format, a...)

	return e

}

func NewApiExceptionFromString(err_msg string) ApiException {

	e := &exception{}

	if !strings.HasPrefix(err_msg, "{") {

		e.Message = err_msg
		e.ErrorCode = InternalServerError
		e.Httpcode = InternalServerError

	}

	err := json.Unmarshal([]byte(err_msg), e)

	if err != nil {
		e.Message = err_msg
		e.ErrorCode = InternalServerError
		e.Httpcode = InternalServerError
	}

	return e
}

func NewApiExceptionFromError(err error) ApiException {

	return NewApiExceptionFromString(err.Error())

}

func NewUnauthorized(format string, a ...interface{}) ApiException {

	return NewApiException("", Unauthorized, codeReason(Unauthorized), format, a...)

}

func NewBadRequest(format string, a ...interface{}) ApiException {

	return NewApiException("", BadRequest, codeReason(BadRequest), format, a...)

}

func NewInternalServerError(format string, a ...interface{}) ApiException {

	return NewApiException("", InternalServerError, codeReason(InternalServerError), format, a...)

}

func NewForbidden(format string, a ...interface{}) ApiException {

	return NewApiException("", Forbidden, codeReason(Forbidden), format, a...)

}

func NewNotFound(format string, a ...interface{}) ApiException {

	return NewApiException("", NotFound, codeReason(NotFound), format, a...)

}

func NewConflict(format string, a ...interface{}) ApiException {

	return NewApiException("", Conflict, codeReason(Conflict), format, a...)

}

func NewPasswordExired(format string, a ...interface{}) ApiException {

	return NewApiException("", PasswordExired, codeReason(PasswordExired), format, a...)

}

func NewVerifyCodeRequired(format string, a ...interface{}) ApiException {

	return NewApiException("", VerifyCodeRequired, codeReason(VerifyCodeRequired), format, a...)

}

func NewRefreshTokenIllegal(format string, a ...interface{}) ApiException {

	return NewApiException("", RefreshTokenIllegal, codeReason(RefreshTokenIllegal), format, a...)

}

func NewAccessTokenIllegal(format string, a ...interface{}) ApiException {

	return NewApiException("", AccessTokenIllegal, codeReason(AccessTokenIllegal), format, a...)

}

func NewRefreshTokenExpired(format string, a ...interface{}) ApiException {

	return NewApiException("", RefreshTokenExpired, codeReason(RefreshTokenExpired), format, a...)

}

func NewAccessTokenExpired(format string, a ...interface{}) ApiException {

	return NewApiException("", AccessTokenExpired, codeReason(AccessTokenExpired), format, a...)

}

func NewSessionTerminated(format string, a ...interface{}) ApiException {

	return NewApiException("", SessionTerminated, codeReason(SessionTerminated), format, a...)

}

func NewOtherClientsLoggedIn(format string, a ...interface{}) ApiException {

	return NewApiException("", OtherClientsLoggedIn, codeReason(OtherClientsLoggedIn), format, a...)

}

func NewOtherIPLoggedIn(format string, a ...interface{}) ApiException {

	return NewApiException("", OtherIPLoggedIn, codeReason(OtherIPLoggedIn), format, a...)

}

func NewOtherPlaceLoggedIn(format string, a ...interface{}) ApiException {

	return NewApiException("", OtherPlaceLoggedIn, codeReason(OtherPlaceLoggedIn), format, a...)

}

func IsNotFoundError(err error) bool {

	if err == nil {
		return false
	}

	v, ok := err.(ApiException)

	if ok {
		return v.GetHttpCode() == NotFound
	}

	return false

}

func IsConflict(err error) bool {

	if err == nil {
		return false
	}

	v, ok := err.(ApiException)

	if !ok {
		return false
	}
	return v.GetHttpCode() == Conflict

}
