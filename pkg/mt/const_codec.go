package mt

// VideoCodec 视频编码
type VideoCodec string

// AudioCodec 音频编码
type AudioCodec string

// Standard 分辨率
type Standard string

// 视频编码
const (
	VideoCodecH264  VideoCodec = "1"  // H.264(x264/AVC)
	VideoCodecVC1   VideoCodec = "2"  // VC-1
	VideoCodecXvid  VideoCodec = "3"  // Xvid
	VideoCodecMPEG2 VideoCodec = "4"  // MPEG-2
	VideoCodecH265  VideoCodec = "16" // H.265(x265/HEVC)
	VideoCodecAV1   VideoCodec = "19" // AV1
	VideoCodecVP89  VideoCodec = "21" // VP8/9
	VideoCodecAVS   VideoCodec = "22" // AVS
)

// 音频编码
const (
	AudioCodecFLAC        AudioCodec = "1"  // FLAC
	AudioCodecAPE         AudioCodec = "2"  // APE
	AudioCodecDTS         AudioCodec = "3"  // DTS
	AudioCodecMP23        AudioCodec = "4"  // MP2/3
	AudioCodecOGG         AudioCodec = "5"  // OGG
	AudioCodecAAC         AudioCodec = "6"  // AAC
	AudioCodecOther       AudioCodec = "7"  // Other
	AudioCodecAC3         AudioCodec = "8"  // AC3(DD)
	AudioCodecTrueHD      AudioCodec = "9"  // TrueHD
	AudioCodecTrueHDAtmos AudioCodec = "10" // TrueHD Atmos
	AudioCodecDTSHDMA     AudioCodec = "11" // DTS-HD MA
	AudioCodecEAC3        AudioCodec = "12" // E-AC3(DDP)
	AudioCodecEAC3Atmos   AudioCodec = "13" // E-AC3 Atoms(DDP Atoms)
	AudioCodecLPCM        AudioCodec = "14" // LPCM/PCM
	AudioCodecWAV         AudioCodec = "15" // WAV
)

// 分辨率
const (
	StandardP1080 Standard = "1" // 1080p
	StandardI1080 Standard = "2" // 1080i
	StandardP720  Standard = "3" // 720p
	StandardSD    Standard = "5" // SD
	Standard4K    Standard = "6" // 4K
	Standard8K    Standard = "7" // 8K
)

// VideoCodecName 返回视频编码的名称
var VideoCodecName = map[VideoCodec]string{
	VideoCodecH264:  "H.264(x264/AVC)",
	VideoCodecVC1:   "VC-1",
	VideoCodecXvid:  "Xvid",
	VideoCodecMPEG2: "MPEG-2",
	VideoCodecH265:  "H.265(x265/HEVC)",
	VideoCodecAV1:   "AV1",
	VideoCodecVP89:  "VP8/9",
	VideoCodecAVS:   "AVS",
}

// AudioCodecName 返回音频编码的名称
var AudioCodecName = map[AudioCodec]string{
	AudioCodecFLAC:        "FLAC",
	AudioCodecAPE:         "APE",
	AudioCodecDTS:         "DTS",
	AudioCodecMP23:        "MP2/3",
	AudioCodecOGG:         "OGG",
	AudioCodecAAC:         "AAC",
	AudioCodecOther:       "Other",
	AudioCodecAC3:         "AC3(DD)",
	AudioCodecTrueHD:      "TrueHD",
	AudioCodecTrueHDAtmos: "TrueHD Atmos",
	AudioCodecDTSHDMA:     "DTS-HD MA",
	AudioCodecEAC3:        "E-AC3(DDP)",
	AudioCodecEAC3Atmos:   "E-AC3 Atoms(DDP Atoms)",
	AudioCodecLPCM:        "LPCM/PCM",
	AudioCodecWAV:         "WAV",
}

// StandardName 返回分辨率的名称
var StandardName = map[Standard]string{
	StandardP1080: "1080p",
	StandardI1080: "1080i",
	StandardP720:  "720p",
	StandardSD:    "SD",
	Standard4K:    "4K",
	Standard8K:    "8K",
}
