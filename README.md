# mt-sdk

M-Team API 的 Go SDK，基于 [resty](https://github.com/go-resty/resty) 实现。

## 功能

- **种子服务**（`Torrent`）：搜索、详情、生成下载令牌、下载 .torrent 文件
- **枚举服务**（`EnumService`）：制作组、国家/地区、来源、视频/音频编码、清晰度、媒介、分类、语言等列表
- **DMM 服务**（`Dmm`）：按番号查询 DMM 作品信息
- **工具函数**：文件保存（自动建目录、清洗非法字符）、文件大小格式化

## 安装

```bash
go get github.com/llbox/mt-sdk
```

## 快速开始

```go
package main

import (
	"context"
	"log/slog"

	mt "github.com/llbox/mt-sdk/pkg"
)

func main() {
	c := mt.New(mt.WithApiKey("你的 API Key"))
	ctx := context.Background()

	// 搜索种子：链式构造查询条件
	rt, err := c.Torrent.Search(ctx, mt.NewSearch(mt.SearchModeAdult).
		Page(1, 20).
		Visible(1).
		Free().
		Build())
	if err != nil {
		panic(err)
	}
	first := rt.Data[0]

	// 详情
	detail, err := c.Torrent.Detail(ctx, first.ID)

	// 生成下载链接
	url, err := c.Torrent.GenDlToken(ctx, first.ID)

	// 下载 .torrent 文件，返回内容和服务器建议的文件名
	data, filename, err := c.Torrent.Download(ctx, url)

	// 保存到文件（自动创建目录、清洗文件名非法字符）
	path, err := mt.SaveFile(".", filename, data)
	slog.Info("saved", "path", path, "size", mt.FormatSize(int64(len(data))))
}
```

完整可运行示例见 [examples/main.go](examples/main.go)：

```bash
go run ./examples
```

## 客户端配置

```go
c := mt.New(
	mt.WithApiKey("你的 API Key"),     // API Key，通过 X-Api-Key 请求头发送
	mt.WithBaseUrl("https://api.m-team.cc"), // 自定义 API 地址
	mt.WithTimeout(30*time.Second),   // 请求超时
	mt.WithRetryCount(3),             // 失败重试次数
)
```

## API 概览

### Torrent

| 方法 | 说明 |
| --- | --- |
| `Search(ctx, req)` | 搜索种子，返回分页列表 |
| `Detail(ctx, id)` | 种子详情 |
| `GenDlToken(ctx, id)` | 生成下载令牌，返回临时下载链接 |
| `Download(ctx, url)` | 下载 .torrent 文件，返回文件内容和服务器建议的文件名（自动跟随 302 跳转，从 `Content-Disposition` 解析文件名） |

### 搜索构造器（SearchBuilder）

`NewSearch(mode)` 链式构造 `TorrentSearchReq`，所有过滤参数都有对应的类型化枚举，IDE 可自动补全：

```go
req := mt.NewSearch(mt.SearchModeAdult).
	Page(1, 20).                                   // 页码、每页条数（默认 1, 20）
	Visible(1).                                    // 仅可见
	Free().                                        // 仅 free 促销，等价于 Discount(mt.DiscountFree)
	Categories(mt.CategoryGroupAdult...).          // 分类，支持分组常量整体传入
	VideoCodecs(mt.VideoCodecH265).                // 视频编码
	AudioCodecs(mt.AudioCodecTrueHDAtmos).         // 音频编码
	Standards(mt.Standard4K).                      // 分辨率
	Countries(mt.CountryJapan).                    // 国家/地区
	Teams(mt.TeamMTeam).                           // 制作组
	ChineseSubtitle().                             // 仅中字，等价于 Labels(mt.LabelsNewChineseSubtitle)
	UploadDateRange("2026-07-01 00:00:00", "").    // 上传时间范围
	Build()
```

类型化枚举：`SearchMode` / `Discount` / `Label` / `Category` / `VideoCodec` / `AudioCodec` / `Standard` / `Country` / `Lang` / `Source` / `Medium` / `Team`，传错类型会编译报错。也可以直接构造 `TorrentSearchReq` 结构体，效果相同。

### EnumService

`TeamList` / `SourceList` / `CountryList` / `VideoCodecList` / `AudioCodecList` / `StandardList` / `CategoryList` / `MediumList` / `Langs`，均为 `(ctx)` 调用，返回对应枚举列表。

### Dmm

| 方法 | 说明 |
| --- | --- |
| `DmmInfo(ctx, dmmCode)` | 按番号查询 DMM 作品信息 |

### 工具函数

| 函数 | 说明 |
| --- | --- |
| `SaveFile(dir, filename, data)` | 保存文件，自动创建目录并清洗 Windows 非法字符，返回完整路径 |
| `SanitizeFilename(name)` | 替换文件名中 `\/:*?"<>\|` 等非法字符 |
| `FormatSize(size int64)` | 字节数转可读格式，如 `6403750869` → `"5.96 GB"` |
| `FormatSizeStr(size string)` | 同上，接收字符串形式的字节数 |
| `ParseTimeMillis(s)` | 解析接口返回的 UTC+8 时间字符串（如 `discountEndTime`: `"2026-07-27 00:32:15"`）为毫秒时间戳，带 error 返回 |
| `TimeMillis(s)` | 同上，解析失败返回 0 |

## 日志

SDK 内部通过 resty 钩子统一记录每个请求的方法、URL、状态码、耗时和响应大小，使用 `slog` 默认 handler。调用方通过 `slog.SetDefault` 自定义输出格式和级别：

```go
slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelInfo,
})))
```

## 目录结构

```
├── pkg/            # SDK 本体
│   ├── client.go   # 客户端与通用响应结构
│   ├── options.go  # 客户端配置项
│   ├── torrent.go  # 种子服务
│   ├── search.go   # 搜索构造器
│   ├── enums.go    # 枚举服务
│   ├── dmm.go      # DMM 服务
│   ├── utils.go    # 文件保存 / 大小格式化 / 时间转换
│   └── const_*.go  # 类型化枚举常量
└── examples/       # 可运行示例
```

## 错误处理

API 业务错误（`code != "0"`）和 HTTP 状态错误统一返回 `*APIError`，可用 `errors.As` 提取：

```go
var apiErr *mt.APIError
if errors.As(err, &apiErr) {
	// apiErr.Code == "-1" 表示 HTTP 状态错误，其余为业务错误码
	slog.Error("request fail", "code", apiErr.Code, "msg", apiErr.Msg)
}
```
