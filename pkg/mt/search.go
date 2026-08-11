package mt

// SearchBuilder 链式构造 TorrentSearchReq，简化多参数搜索
type SearchBuilder struct {
	req TorrentSearchReq
}

// NewSearch 创建搜索构造器，默认第 1 页、每页 20 条
func NewSearch(mode SearchMode) *SearchBuilder {
	return &SearchBuilder{req: TorrentSearchReq{
		Mode:       mode,
		PageNumber: 1,
		PageSize:   20,
	}}
}

// Page 设置页码和每页条数
func (b *SearchBuilder) Page(number, size int) *SearchBuilder {
	b.req.PageNumber = number
	b.req.PageSize = size
	return b
}

// Visible 设置可见性过滤（1 = 仅可见）
func (b *SearchBuilder) Visible(v int) *SearchBuilder {
	b.req.Visible = v
	return b
}

// Discount 设置促销过滤
func (b *SearchBuilder) Discount(d Discount) *SearchBuilder {
	b.req.Discount = d
	return b
}

// Free 仅搜索 free 促销的种子
func (b *SearchBuilder) Free() *SearchBuilder {
	return b.Discount(DiscountFree)
}

// Categories 追加分类过滤，可配合 CategoryGroupAdult 等分组使用：
// Categories(mt.CategoryGroupAdult...)
func (b *SearchBuilder) Categories(c ...Category) *SearchBuilder {
	b.req.Categories = append(b.req.Categories, c...)
	return b
}

// VideoCodecs 追加视频编码过滤
func (b *SearchBuilder) VideoCodecs(c ...VideoCodec) *SearchBuilder {
	b.req.VideoCodes = append(b.req.VideoCodes, c...)
	return b
}

// AudioCodecs 追加音频编码过滤
func (b *SearchBuilder) AudioCodecs(c ...AudioCodec) *SearchBuilder {
	b.req.AudioCodes = append(b.req.AudioCodes, c...)
	return b
}

// Standards 追加分辨率过滤
func (b *SearchBuilder) Standards(s ...Standard) *SearchBuilder {
	b.req.Standards = append(b.req.Standards, s...)
	return b
}

// Countries 追加国家/地区过滤
func (b *SearchBuilder) Countries(c ...Country) *SearchBuilder {
	b.req.Countries = append(b.req.Countries, c...)
	return b
}

// Teams 追加制作组过滤
func (b *SearchBuilder) Teams(t ...Team) *SearchBuilder {
	b.req.Teams = append(b.req.Teams, t...)
	return b
}

// Labels 追加标记过滤
func (b *SearchBuilder) Labels(l ...Label) *SearchBuilder {
	b.req.LabelsNew = append(b.req.LabelsNew, l...)
	return b
}

// ChineseSubtitle 仅搜索带中文字幕的种子
func (b *SearchBuilder) ChineseSubtitle() *SearchBuilder {
	return b.Labels(LabelsNewChineseSubtitle)
}

// UploadDateRange 设置上传时间范围，格式 "2006-01-02 15:04:05"
func (b *SearchBuilder) UploadDateRange(start, end string) *SearchBuilder {
	b.req.UploadDateStart = start
	b.req.UploadDateEnd = end
	return b
}

// Build 返回构造好的搜索请求
func (b *SearchBuilder) Build() TorrentSearchReq {
	return b.req
}
