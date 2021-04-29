package responses

import (
	"github.com/SemmiDev/go-backend/commons/statuscode"
	"net/http"
)

var RESPONSE_OK = ResponseBase{
	Code:    http.StatusOK,
	Message: statuscode.OK.String(),
}

var RESPONSE_FORBIDDEN = ResponseBase{
	Code:    http.StatusForbidden,
	Message: statuscode.NoAccess.String(),
}
