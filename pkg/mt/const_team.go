package mt

// Team 制作组
type Team string

// 制作组
const (
	TeamBMDru       Team = "6"  // BMDru
	TeamPack        Team = "8"  // Pack
	TeamMTeam       Team = "9"  // MTeam
	TeamCNHK        Team = "19" // CNHK
	TeamTnP         Team = "23" // TnP
	TeamCatEDU      Team = "25" // CatEDU
	TeamARiC        Team = "26" // ARiC
	Team7ACG        Team = "30" // 7³ACG
	TeamJKCT        Team = "31" // JKCT
	TeamG00DB0Y     Team = "35" // G00DB0Y
	TeamD0          Team = "36" // D0
	TeamHBO         Team = "40" // HBO
	TeamREE         Team = "41" // REE
	TeamTPTV        Team = "43" // TPTV
	TeamMWeb        Team = "44" // MWeb
	TeamCTRL        Team = "45" // CTRL
	TeamZTR         Team = "48" // ZTR
	Team126811      Team = "49" // 126811
	TeamDST         Team = "57" // DST
	TeamStarfallWeb Team = "59" // StarfallWeb
	TeamRRS         Team = "61" // RRS
	TeamLijiangTv   Team = "62" // lijiang-tv
	TeamZZH         Team = "63" // ZZH
	TeamDStudio     Team = "64" // DStudio
	TeamAisha       Team = "65" // Aisha
)

// TeamName 返回制作组的名称
var TeamName = map[Team]string{
	TeamBMDru:       "BMDru",
	TeamPack:        "Pack",
	TeamMTeam:       "MTeam",
	TeamCNHK:        "CNHK",
	TeamTnP:         "TnP",
	TeamCatEDU:      "CatEDU",
	TeamARiC:        "ARiC",
	Team7ACG:        "7³ACG",
	TeamJKCT:        "JKCT",
	TeamG00DB0Y:     "G00DB0Y",
	TeamD0:          "D0",
	TeamHBO:         "HBO",
	TeamREE:         "REE",
	TeamTPTV:        "TPTV",
	TeamMWeb:        "MWeb",
	TeamCTRL:        "CTRL",
	TeamZTR:         "ZTR",
	Team126811:      "126811",
	TeamDST:         "DST",
	TeamStarfallWeb: "StarfallWeb",
	TeamRRS:         "RRS",
	TeamLijiangTv:   "lijiang-tv",
	TeamZZH:         "ZZH",
	TeamDStudio:     "DStudio",
	TeamAisha:       "Aisha",
}
