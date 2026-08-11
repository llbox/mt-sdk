package mt

import (
	"encoding/json"
	"testing"
)

func TestSearchBuilder(t *testing.T) {
	req := NewSearch(SearchModeAdult).
		Page(2, 50).
		Visible(1).
		Free().
		Categories(CategoryAVHDCensored).
		Categories(CategoryGroupMovie...).
		VideoCodecs(VideoCodecH265, VideoCodecAV1).
		ChineseSubtitle().
		UploadDateRange("2026-07-01 00:00:00", "2026-07-27 00:00:00").
		Build()

	if req.Mode != SearchModeAdult || req.PageNumber != 2 || req.PageSize != 50 || req.Visible != 1 {
		t.Errorf("basic fields wrong: %+v", req)
	}
	if req.Discount != DiscountFree {
		t.Errorf("discount = %q, want FREE", req.Discount)
	}
	if len(req.Categories) != 1+len(CategoryGroupMovie) {
		t.Errorf("categories len = %d, want %d", len(req.Categories), 1+len(CategoryGroupMovie))
	}
	if req.LabelsNew[0] != LabelsNewChineseSubtitle {
		t.Errorf("label = %q, want 中字", req.LabelsNew[0])
	}

	body, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(body, &m); err != nil {
		t.Fatal(err)
	}
	if m["mode"] != "adult" {
		t.Errorf("mode json = %v", m["mode"])
	}
	vcs := m["videoCodes"].([]any)
	if vcs[0] != "16" || vcs[1] != "19" {
		t.Errorf("videoCodes json = %v", vcs)
	}
}
