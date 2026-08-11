package mt

import "context"

type DmmService struct {
	c *MTClient
}

func NewDmmService(c *MTClient) *DmmService {
	return &DmmService{c}
}

type DmmInfoDetail struct {
	Url           string            `json:"url"`
	ProductNumber string            `json:"productNumber"`
	Title         string            `json:"title"`
	Pic1          string            `json:"pic1"`
	Pic2          string            `json:"pic2"`
	BaseInfo      map[string]string `json:"baseInfo"`
	ActressList   []string          `json:"actressList"`
	KeywordList   []string          `json:"keywordList"`
	TextData      string            `json:"textData"`
	PicData       string            `json:"picData"`
	ReleaseDate   string            `json:"releaseDate"`
	Duration      string            `json:"duration"`
	Director      string            `json:"director"`
	Series        string            `json:"series"`
	Maker         string            `json:"maker"`
	Label         string            `json:"label"`
	Score         string            `json:"score"`
}

func (ds *DmmService) DmmInfo(ctx context.Context, dmmCode string) (*DmmInfoDetail, error) {
	var rt Result[*DmmInfoDetail]
	if resp, err := ds.c.http.R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"dmmCode": dmmCode,
		}).
		SetResult(&rt).
		Post(dmmInfoPath); err != nil {
		return nil, errWithResp(err, resp)
	} else if resp.IsError() {
		return nil, NewHttpStatusError(resp.StatusCode())
	}
	if rt.IsError() {
		return nil, NewAPIError(rt.Code, rt.Message)
	}
	return rt.Data, nil
}
