package mt

import (
	"context"
	"log/slog"
	"mime"
)

type TorrentService struct {
	c *MTClient
}

func NewTorrentService(c *MTClient) *TorrentService {
	return &TorrentService{c}
}

const (
	searchPath     = "/api/torrent/search"
	detailPath     = "/api/torrent/detail"
	genDlTokenPath = "/api/torrent/genDlToken"
	dmmInfoPath    = "/api/dmm/dmmInfo"
)

type TorrentSearchReq struct {
	Mode            SearchMode   `json:"mode"`
	PageNumber      int          `json:"pageNumber"`
	PageSize        int          `json:"pageSize"`
	AudioCodes      []AudioCodec `json:"audioCodes,omitempty"`
	VideoCodes      []VideoCodec `json:"videoCodes,omitempty"`
	Categories      []Category   `json:"categories,omitempty"`
	Countries       []Country    `json:"countries,omitempty"`
	Discount        Discount     `json:"discount,omitempty"`
	LabelsNew       []Label      `json:"labelsNew,omitempty"`
	Standards       []Standard   `json:"standards,omitempty"`
	Teams           []Team       `json:"teams,omitempty"`
	UploadDateStart string       `json:"uploadDateStart,omitempty"`
	UploadDateEnd   string       `json:"uploadDateEnd,omitempty"`
	Visible         int          `json:"visible,omitempty"`
}
type TorrentSearchResp struct {
	PageNumber string         `json:"pageNumber"`
	PageSize   string         `json:"pageSize"`
	Total      string         `json:"total"`
	TotalPages string         `json:"totalPages"`
	Data       []*TorrentItem `json:"data"`
}

// TorrentItem
// name 文件名
// status.discount、status.discountEndTime判断促销
// status.seeders 做种人数
// status.leechers下载人数
// status.createdDate 创建时间
// size 文件大小("6403750869"=5.96GB)
// imageList 列表图片
type TorrentItem struct {
	ID                 string         `json:"id"`
	CreatedDate        string         `json:"createdDate"`
	LastModifiedDate   string         `json:"lastModifiedDate"`
	Name               string         `json:"name"`
	SmallDescr         string         `json:"smallDescr"`
	Imdb               string         `json:"imdb"`
	ImdbRating         *string        `json:"imdbRating"`
	Douban             string         `json:"douban"`
	DoubanRating       *string        `json:"doubanRating"`
	DmmCode            string         `json:"dmmCode"`
	Author             *string        `json:"author"`
	Category           string         `json:"category"`
	Source             *string        `json:"source"`
	Medium             *string        `json:"medium"`
	Standard           string         `json:"standard"`
	VideoCodec         string         `json:"videoCodec"`
	AudioCodec         string         `json:"audioCodec"`
	Team               *string        `json:"team"`
	Processing         *string        `json:"processing"`
	Countries          []string       `json:"countries"`
	Numfiles           string         `json:"numfiles"`
	Size               string         `json:"size"`
	Labels             string         `json:"labels"`
	LabelsNew          []string       `json:"labelsNew"`
	MsUp               string         `json:"msUp"`
	Anonymous          bool           `json:"anonymous"`
	InfoHash           *string        `json:"infoHash"`
	HasChineseSubtitle bool           `json:"hasChineseSubtitle"`
	Status             *TorrentStatus `json:"status"`
	DmmInfo            *DmmInfo       `json:"dmmInfo"`
	EditedBy           *string        `json:"editedBy"`
	EditDate           *string        `json:"editDate"`
	Collection         bool           `json:"collection"`
	CollectionStatus   any            `json:"collectionStatus"`
	InRss              bool           `json:"inRss"`
	CanVote            bool           `json:"canVote"`
	ImageList          []string       `json:"imageList"`
	ResetBox           *string        `json:"resetBox"`
}

type TorrentStatus struct {
	ID               string          `json:"id"`
	CreatedDate      string          `json:"createdDate"`
	LastModifiedDate string          `json:"lastModifiedDate"`
	PickType         string          `json:"pickType"`
	ToppingLevel     string          `json:"toppingLevel"`
	ToppingEndTime   *string         `json:"toppingEndTime"`
	Discount         string          `json:"discount"`
	DiscountEndTime  *string         `json:"discountEndTime"`
	TimesCompleted   string          `json:"timesCompleted"`
	Comments         string          `json:"comments"`
	LastAction       *string         `json:"lastAction"`
	LastSeederAction *string         `json:"lastSeederAction"`
	Views            string          `json:"views"`
	Hits             string          `json:"hits"`
	Support          string          `json:"support"`
	Oppose           string          `json:"oppose"`
	Status           string          `json:"status"`
	Seeders          string          `json:"seeders"`
	Leechers         string          `json:"leechers"`
	Banned           bool            `json:"banned"`
	Visible          bool            `json:"visible"`
	PromotionRule    any             `json:"promotionRule"`
	MallSingleFree   *MallSingleFree `json:"mallSingleFree"`
}

type MallSingleFree struct {
	Auction          string `json:"auction"`
	CreateDate       string `json:"createDate"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	FreeDay          string `json:"freeDay"`
	Id               string `json:"id"`
	isAdult          bool   `json:"isAdult"`
	LastModifiedDate string `json:"lastModifiedDate"`
	Points           string `json:"points"`
	Status           string `json:"status"`
	Torrent          string `json:"torrent"`
	UserId           string `json:"userId"`
}

type DmmInfo struct {
	CreatedDate      string   `json:"createdDate"`
	LastModifiedDate string   `json:"lastModifiedDate"`
	ID               string   `json:"id"`
	ProductNumber    string   `json:"productNumber"`
	Director         string   `json:"director"`
	Series           string   `json:"series"`
	Maker            string   `json:"maker"`
	Label            string   `json:"label"`
	KeywordList      []string `json:"keywordList"`
	ActressList      []string `json:"actressList"`
}

func (tc *TorrentService) Search(ctx context.Context, req TorrentSearchReq) (*TorrentSearchResp, error) {
	var rt Result[*TorrentSearchResp]
	if resp, err := tc.c.http.R().SetContext(ctx).SetBody(req).SetResult(&rt).Post(searchPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}
	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

// TorrentDetail
// descr 富文本展示
type TorrentDetail struct {
	TorrentItem
	OriginFileName string  `json:"originFileName"`
	Descr          string  `json:"descr"`
	Nfo            *string `json:"nfo"`
	Mediainfo      string  `json:"mediainfo"`
	Cids           any     `json:"cids"`
	Aids           any     `json:"aids"`
	Scope          string  `json:"scope"`
	ScopeTeams     any     `json:"scopeTeams"`
	Thanked        bool    `json:"thanked"`
	Rewarded       bool    `json:"rewarded"`
	AlbumList      any     `json:"albumList"`
}

func (tc *TorrentService) Detail(ctx context.Context, id string) (*TorrentDetail, error) {
	var rt Result[*TorrentDetail]
	if resp, err := tc.c.http.R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"id": id,
		}).
		SetResult(&rt).
		Post(detailPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}
	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

// GenDlToken 获取种子的下载链接
func (tc *TorrentService) GenDlToken(ctx context.Context, id string) (string, error) {
	var rt Result[string]
	if resp, err := tc.c.http.R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"id": id,
		}).
		SetResult(&rt).
		Post(genDlTokenPath); err != nil {
		return "", errWithResp(err, resp)
	} else if resp.IsError() {
		return "", NewHttpStatusError(resp.StatusCode())
	}
	if rt.IsError() {
		return "", NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

// Download 下载种子文件，返回 .torrent 文件内容和服务器建议的文件名
// （从 Content-Disposition 解析，服务器未提供时为空字符串）
func (tc *TorrentService) Download(ctx context.Context, url string) ([]byte, string, error) {
	resp, err := tc.c.http.R().SetContext(ctx).Get(url)
	if err != nil {
		return nil, "", errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, "", NewHttpStatusError(resp.StatusCode())
	}
	slog.InfoContext(ctx, "download success",
		"finalUrl", resp.RawResponse.Request.URL.String(),
		"status", resp.StatusCode(),
		"contentType", resp.Header().Get("Content-Type"),
		"contentDisposition", resp.Header().Get("Content-Disposition"),
		"size", resp.Size(),
	)
	return resp.Body(), parseFilename(resp.Header().Get("Content-Disposition")), nil
}

// parseFilename 从 Content-Disposition 头解析文件名，如：
// attachment; filename="[M-TEAM][1041294]START-364-UC.torrent"
func parseFilename(contentDisposition string) string {
	if contentDisposition == "" {
		return ""
	}
	if _, params, err := mime.ParseMediaType(contentDisposition); err == nil {
		return params["filename"]
	}
	return ""
}
