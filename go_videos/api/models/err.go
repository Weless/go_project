package models

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpCode int
	Error    Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpCode: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrResponse{HttpCode: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
	ErrorDBError                = ErrResponse{HttpCode: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrResponse{HttpCode: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
