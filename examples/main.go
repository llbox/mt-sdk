package main

import (
	"context"
	"fmt"
	"log/slog"
	mt "github.com/llbox/mt-sdk/pkg"
	"os"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	c := mt.New(mt.WithApiKey(""))
	ctx := context.Background()

	// 搜索：链式构造查询条件
	req := mt.NewSearch(mt.SearchModeAdult).
		Page(1, 20).
		Visible(1).
		Free().
		Build()
	rt, err := c.Torrent.Search(ctx, req)
	if err != nil {
		slog.Error("search fail", "err", err)
		return
	}

	first := rt.Data[1]
	slog.Info("search success",
		"first.id", first.ID,
		"first.discount", first.Status.Discount,
		"first.name", first.Name,
		"first.size", mt.FormatSizeStr(first.Size))

	// 详情
	detail, err := c.Torrent.Detail(ctx, first.ID)
	if err != nil {
		slog.Error("detail fail", "err", err)
		return
	}
	slog.Info("detail success", "desc", detail.Descr)

	// 下载
	url, err := c.Torrent.GenDlToken(ctx, first.ID)
	if err != nil {
		slog.ErrorContext(ctx, "gen token fail", "err", err)
		return
	}
	slog.Info("gen token success", "url", url)

	data, filename, err := c.Torrent.Download(ctx, url)
	if err != nil {
		slog.ErrorContext(ctx, "download fail", "err", err)
		return
	}

	// 保存 .torrent 文件，服务器未提供文件名时用种子名兜底
	if filename == "" {
		filename = fmt.Sprintf("%s.torrent", first.Name)
	}
	path, err := mt.SaveFile(".", filename, data)
	if err != nil {
		slog.ErrorContext(ctx, "save file fail", "err", err)
		return
	}
	slog.InfoContext(ctx, "save file success", "path", path, "size", mt.FormatSize(int64(len(data))))
}
