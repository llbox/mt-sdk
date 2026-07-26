package pkg

import (
	"context"
)

type EnumService struct {
	c *MTClient
}

type TypeEnum int

const (
	EnumTeam TypeEnum = iota + 1
	EnumCountry
	EnumSource
	EnumVideoCode
	EnumAudioCode
	EnumStandard
	EnumMedium
	EnumLanguage
	EnumCategory
)

const (
	teamPath      = "/api/torrent/teamList"
	countryPath   = "/api/system/countryList"
	sourcePath    = "/api/torrent/sourceList"
	videoCodePath = "/api/torrent/videoCodecList"
	audioCodePath = "/api/torrent/audioCodecList"
	standardPath  = "/api/torrent/standardList"
	categoryPath  = "/api/torrent/categoryList"
	mediumPath    = "/api/torrent/mediumList"
	langsPath     = "/api/system/langs"
)

func NewEnumService(c *MTClient) *EnumService {
	return &EnumService{c: c}
}

type BaseEnumItem struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Order            string `json:"order"`
	CreatedDate      string `json:"createdDate"`
	LaseModifiedDate string `json:"laseModifiedDate"`
}

type MultiName struct {
	NameChs string `json:"nameChs"`
	NameCht string `json:"nameCht"`
	NameEng string `json:"nameEng"`
}

type TeamItem struct {
	BaseEnumItem
	Leader    string   `json:"leader"`
	Members   []string `json:"members"`
	FreeOffer bool     `json:"freeOffer"`
}

type SourceItem struct {
	BaseEnumItem
	MultiName
}

type MediumItem struct {
	BaseEnumItem
	MultiName
}

type CountryItem struct {
	BaseEnumItem
	Pic string `json:"pic"`
}

type StandardItem struct {
	BaseEnumItem
}

type LangItem struct {
	BaseEnumItem
	FlagPic  string `json:"flagpic"`
	LangName string `json:"langName"`
	LangTag  any    `json:"langTag"`
	SiteLang bool   `json:"siteLang"`
	SubLang  bool   `json:"subLang"`
}

type VideoCodeItem struct {
	BaseEnumItem
}

type AudioCodeItem struct {
	BaseEnumItem
}

type CategoryItem struct {
	Adult     []string           `json:"adult"`
	Movie     []string           `json:"movie"`
	Music     []string           `json:"music"`
	TVShow    []string           `json:"tvshow"`
	Waterfall []string           `json:"waterfall"`
	List      []CategoryListItem `json:"list"`
}

type CategoryListItem struct {
	MultiName
	Id               string  `json:"id"`
	Parent           *string `json:"parent"`
	Image            string  `json:"image"`
	Order            string  `json:"order"`
	CreatedDate      string  `json:"createdDate"`
	LaseModifiedDate string  `json:"laseModifiedDate"`
}

func (es *EnumService) TeamList(ctx context.Context) ([]*TeamItem, error) {
	var rt Result[[]*TeamItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(teamPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

func (es *EnumService) SourceList(ctx context.Context) ([]*SourceItem, error) {
	var rt Result[[]*SourceItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(sourcePath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

func (es *EnumService) CountryList(ctx context.Context) ([]*CountryItem, error) {
	var rt Result[[]*CountryItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(countryPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

func (es *EnumService) VideoCodecList(ctx context.Context) ([]*VideoCodeItem, error) {
	var rt Result[[]*VideoCodeItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(videoCodePath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

func (es *EnumService) AudioCodecList(ctx context.Context) ([]*AudioCodeItem, error) {
	var rt Result[[]*AudioCodeItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(audioCodePath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

func (es *EnumService) StandardList(ctx context.Context) ([]*StandardItem, error) {
	var rt Result[[]*StandardItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(standardPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

func (es *EnumService) CategoryList(ctx context.Context) (*CategoryItem, error) {
	var rt Result[*CategoryItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(categoryPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

func (es *EnumService) MediumList(ctx context.Context) ([]*MediumItem, error) {
	var rt Result[[]*MediumItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(mediumPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}

func (es *EnumService) Langs(ctx context.Context) ([]*LangItem, error) {
	var rt Result[[]*LangItem]
	if resp, err := es.c.http.R().SetContext(ctx).SetResult(&rt).Post(langsPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}

	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}
