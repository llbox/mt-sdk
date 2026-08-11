package mt

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type APIError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("mt: api error (code=%s): %s", e.Code, e.Msg)
}

func NewAPIError(code FlexString, msg string) *APIError {
	return &APIError{Code: string(code), Msg: msg}
}

func NewHttpStatusError(status int) *APIError {
	msg := fmt.Sprintf("http status is %d", status)
	return NewAPIError("-1", msg)
}

// errWithResp 把原始响应体附在错误信息后，方便排查解析失败等问题
func errWithResp(err error, resp *resty.Response) error {
	if resp == nil {
		return err
	}
	body := resp.String()
	if len(body) > 1024 {
		body = body[:1024] + "...(truncated)"
	}
	return fmt.Errorf("%w, response body: %s", err, body)
}
