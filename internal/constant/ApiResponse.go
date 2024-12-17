package constant

import "time"

type ApiResponse struct {
	Txid      string      `json:"txid"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	TimeStamp time.Time   `json:"timestamp"`
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{Txid: GenTxID(), TimeStamp: time.Now()}
}

func (r *ApiResponse) BadReqResp(err error) *ApiResponse {

	r.Code = STATUS_BAD_REQUEST_CODE
	r.Message = STATUS_BAD_REQUEST_MSG
	r.Data = err.Error()

	return r
}

func (r *ApiResponse) OkResp() *ApiResponse {

	r.Code = STATUS_SUCCESS_CODE
	r.Message = STATUS_SUCCESS_MSG
	r.Data = "OK"

	return r
}

func (r *ApiResponse) OkRespWithData(data interface{}) *ApiResponse {
	r.Code = STATUS_SUCCESS_CODE
	r.Message = STATUS_SUCCESS_MSG
	r.Data = data
	return r
}

func (r *ApiResponse) FailResp(err error) *ApiResponse {

	r.Code = STATUS_FAIL_CODE
	r.Message = STATUS_FAIL_MSG
	r.Data = err.Error()

	return r
}

func (r *ApiResponse) InternalDbErrorResp(err error) *ApiResponse {

	r.Code = STATUS_INTERNAL_DB_ERROR_CODE
	r.Message = STATUS_INTERNAL_DB_ERROR_MSG
	r.Data = err.Error()

	return r
}

func (r *ApiResponse) NotFoundErrorResp() *ApiResponse {

	r.Code = STATUS_NOT_FOUND_CODE
	r.Message = STATUS_NOT_FOUND_MSG
	r.Data = "This request is not found"

	return r
}

func (r *ApiResponse) UserExistResp() *ApiResponse {

	r.Code = STATUS_USER_EXIST_CODE
	r.Message = STATUS_USER_EXIST_MSG
	r.Data = "존재하는 아이디입니다."

	return r
}

func (r *ApiResponse) UserNotExistResp() *ApiResponse {
	r.Code = STATUS_USER_NOT_EXIST_CODE
	r.Message = STATUS_USER_NOT_EXIST_MSG
	r.Data = "존재하지 않는 아이디입니다."

	return r
}

func (r *ApiResponse) RequireIdPasswordResp() *ApiResponse {
	r.Code = STATUS_REQUIRE_ID_PASSWORD_CODE
	r.Message = STATUS_REQUIRE_ID_PASSWORD_MSG
	r.Data = "아이디와 패스워드를 모두 입력해주세요."

	return r
}

func (r *ApiResponse) LoginSuccessResp(userName string) *ApiResponse {
	r.Code = STATUS_SUCCESS_CODE
	r.Message = STATUS_SUCCESS_MSG
	r.Data = userName + "님 환영합니다."

	return r
}

func (r *ApiResponse) LoginFailResp() *ApiResponse {
	r.Code = STATUS_FAIL_CODE
	r.Message = STATUS_FAIL_MSG
	r.Data = "아이디 혹은 패스워드 오류입니다."

	return r
}

func (r *ApiResponse) BoardNotExistResp() *ApiResponse {
	r.Code = STATUS_BOARD_NOT_EXIST_CODE
	r.Message = STATUS_BOARD_NOT_EXIST_MSG
	r.Data = "존재하지 않는 게시판입니다."

	return r
}
