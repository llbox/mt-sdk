package pkg

import "testing"

func TestParseTimeMillis(t *testing.T) {
	// 2026-07-27 00:32:15 UTC+8 = 2026-07-26 16:32:15 UTC
	ms, err := ParseTimeMillis("2026-07-27 00:32:15")
	if err != nil {
		t.Fatal(err)
	}
	const want int64 = 1785083535000
	if ms != want {
		t.Errorf("got %d, want %d", ms, want)
	}
}

func TestTimeMillisInvalid(t *testing.T) {
	if got := TimeMillis("not a time"); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}

func TestFormatSize(t *testing.T) {
	cases := map[int64]string{
		500:           "500 B",
		1048576:       "1.00 MB",
		6403750869:    "5.96 GB",
		2199023255552: "2.00 TB",
	}
	for in, want := range cases {
		if got := FormatSize(in); got != want {
			t.Errorf("FormatSize(%d) = %s, want %s", in, got, want)
		}
	}
}

func TestSanitizeFilename(t *testing.T) {
	got := SanitizeFilename(`a/b\c:d*e?f"g<h>i|j.torrent`)
	want := "a_b_c_d_e_f_g_h_i_j.torrent"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
