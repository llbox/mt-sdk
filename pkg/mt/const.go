package mt

// Discount 促销类型
type Discount string

// Label 种子标记
type Label string

// SearchMode 搜索模式
type SearchMode string

// 促销
const (
	DiscountFree      Discount = "FREE"       // free
	DiscountPercent50 Discount = "PERCENT_50" // 50%
	DiscountPercent70 Discount = "PERCENT_70" // 30%
	DiscountNormal    Discount = "NORMAL"
)

// 标记
const (
	LabelsNew4K              Label = "4k"
	LabelsNew8K              Label = "8k"
	LabelsNewHDR             Label = "hdr"
	LabelsNewHDR10           Label = "hdr10"
	LabelsNewHDR10Plus       Label = "hdr10+"
	LabelsNewHLG             Label = "hlg"
	LabelsNewDoVi            Label = "DoVi"  // 杜比视界
	LabelsNewHDRVi           Label = "HDRVi" // HDR Vivid
	LabelsNewChineseSubtitle Label = "中字"
	LabelsNewChineseDub      Label = "中配"
	LabelsNewAIPoJie         Label = "AI破解"
)

const (
	SearchModeNormal SearchMode = "normal"
	SearchModeAdult  SearchMode = "adult"
	SearchModeTVShow SearchMode = "tvshow"
	SearchModeMovie  SearchMode = "movie"
)

// DiscountName 返回促销类型的名称
var DiscountName = map[Discount]string{
	DiscountFree:      "free",
	DiscountPercent50: "50%",
	DiscountPercent70: "30%",
	DiscountNormal:    "DiscountNormal",
}

// LabelName 返回种子标记的名称
var LabelName = map[Label]string{
	LabelsNew4K:              "4k",
	LabelsNew8K:              "8k",
	LabelsNewHDR:             "hdr",
	LabelsNewHDR10:           "hdr10",
	LabelsNewHDR10Plus:       "hdr10+",
	LabelsNewHLG:             "hlg",
	LabelsNewDoVi:            "杜比视界",
	LabelsNewHDRVi:           "HDR Vivid",
	LabelsNewChineseSubtitle: "中字",
	LabelsNewChineseDub:      "中配",
	LabelsNewAIPoJie:         "AI破解",
}

// SearchModeName 返回搜索模式的名称
var SearchModeName = map[SearchMode]string{
	SearchModeNormal: "normal",
	SearchModeAdult:  "adult",
	SearchModeTVShow: "tvshow",
	SearchModeMovie:  "movie",
}
