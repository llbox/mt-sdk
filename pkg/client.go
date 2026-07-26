package pkg

import (
	"encoding/json"
	"log/slog"

	"github.com/go-resty/resty/v2"
)

type MTClient struct {
	http        *resty.Client
	Torrent     *TorrentService
	EnumService *EnumService
	Dmm         *DmmService
}

func New(opts ...Option) *MTClient {
	c := resty.New().SetBaseURL("https://api.m-team.cc")

	// 统一记录请求日志，使用 slog 默认 handler（调用方可通过 slog.SetDefault 自定义）
	c.OnAfterResponse(func(_ *resty.Client, resp *resty.Response) error {
		slog.InfoContext(resp.Request.Context(), "http request",
			"method", resp.Request.Method,
			"url", resp.Request.URL,
			"status", resp.StatusCode(),
			"duration", resp.Time(),
			"size", resp.Size(),
		)
		return nil
	})
	c.OnError(func(req *resty.Request, err error) {
		slog.ErrorContext(req.Context(), "http request fail",
			"method", req.Method,
			"url", req.URL,
			"err", err,
		)
	})

	mc := &MTClient{
		http: c,
	}
	for _, opt := range opts {
		opt(mc)
	}
	mc.Torrent = NewTorrentService(mc)
	mc.EnumService = NewEnumService(mc)
	mc.Dmm = NewDmmService(mc)
	return mc
}

// FlexString 兼容 JSON 字符串和数字两种写法，统一按字符串处理
type FlexString string

func (f *FlexString) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '"' {
		return json.Unmarshal(b, (*string)(f))
	}
	*f = FlexString(b)
	return nil
}

type Result[T any] struct {
	Code    FlexString `json:"code"`
	Message string     `json:"message"`
	Data    T          `json:"data"`
}

func (r *Result[T]) IsError() bool {
	if r.Code != "0" || r.Message != "SUCCESS" {
		return true
	}
	return false
}
