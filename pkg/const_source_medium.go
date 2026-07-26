package pkg

// Source 来源
type Source string

// Medium 媒介
type Medium string

// 来源
const (
	SourceBluRay Source = "1"  // Bluray
	SourceDVD    Source = "3"  // DVD
	SourceRemux  Source = "4"  // Remux
	SourceHDTV   Source = "5"  // HDTV/TV
	SourceOther  Source = "6"  // Other
	SourceWebDL  Source = "8"  // Web-DL
	SourceCD     Source = "10" // CD
)

// 媒介
const (
	MediumBluRay Medium = "1"  // Blu-ray
	MediumHDDVD  Medium = "2"  // HD DVD
	MediumRemux  Medium = "3"  // Remux
	MediumMiniBD Medium = "4"  // MiniBD
	MediumHDTV   Medium = "5"  // HDTV
	MediumDVDR   Medium = "6"  // DVDR
	MediumEncode Medium = "7"  // Encode
	MediumCD     Medium = "8"  // CD
	MediumTrack  Medium = "9"  // Track
	MediumWebDL  Medium = "10" // Web-DL
)
