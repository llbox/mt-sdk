package mt

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Windows 文件名中不允许的字符
var invalidFilenameChars = strings.NewReplacer(
	`\`, "_", `/`, "_", `:`, "_", `*`, "_", `?`, "_",
	`"`, "_", `<`, "_", `>`, "_", `|`, "_",
)

// SanitizeFilename 将文件名中 Windows 不允许的字符替换为下划线
func SanitizeFilename(name string) string {
	return invalidFilenameChars.Replace(name)
}

// SaveFile 保存数据到 dir/filename，目录不存在时自动创建，
// 文件名中的非法字符会被清洗，返回实际写入的完整路径
func SaveFile(dir, filename string, data []byte) (string, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	path := filepath.Join(dir, SanitizeFilename(filename))
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return "", err
	}
	return path, nil
}

// FormatSize 字节数转可读字符串，如 6403750869 -> "5.96 GB"
func FormatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// FormatSizeStr 解析字符串形式的字节数并格式化（如 TorrentItem.Size），
// 解析失败时原样返回
func FormatSizeStr(size string) string {
	n, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return size
	}
	return FormatSize(n)
}

// TimeLayout 接口返回的时间字符串格式（UTC+8），如 discountEndTime: "2026-07-27 00:32:15"
const TimeLayout = "2006-01-02 15:04:05"

// 接口时间字符串使用的时区（UTC+8）
var timeLocation = time.FixedZone("UTC+8", 8*60*60)

// ParseTimeMillis 解析接口返回的 UTC+8 时间字符串，返回毫秒时间戳
func ParseTimeMillis(s string) (int64, error) {
	t, err := time.ParseInLocation(TimeLayout, s, timeLocation)
	if err != nil {
		return 0, err
	}
	return t.UnixMilli(), nil
}

// TimeMillis 同 ParseTimeMillis，解析失败时返回 0
func TimeMillis(s string) int64 {
	ms, _ := ParseTimeMillis(s)
	return ms
}
