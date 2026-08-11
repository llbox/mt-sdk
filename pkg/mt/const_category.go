package mt

// Category 种子分类
type Category string

// 分类
const (
	CategoryMovie               Category = "100" // 电影 Movie
	CategoryTVSeries            Category = "105" // 影剧/综艺 TV Series
	CategoryMusic               Category = "110" // Music
	CategoryAVCensored          Category = "115" // AV(有码) AV(有碼)
	CategoryAVUncensored        Category = "120" // AV(无码) AV(無碼)
	CategoryMovieSD             Category = "401" // 电影/SD Movie/SD
	CategoryTVSeriesHD          Category = "402" // 影剧/综艺/HD TV Series/HD
	CategoryTVSeriesSD          Category = "403" // 影剧/综艺/SD TV Series/SD
	CategoryRecord              Category = "404" // 纪录 Record
	CategoryAnime               Category = "405" // 动画 Anime
	CategoryMV                  Category = "406" // 演唱 MV
	CategorySports              Category = "407" // 运动 Sports
	CategoryMiscOther           Category = "409" // Misc(其他) Misc(Other)
	CategoryAVHDCensored        Category = "410" // AV(有码)/HD Censored AV(有碼)/HD Censored
	CategoryHGame               Category = "411" // H-游戏 H-Game
	CategoryHAnime              Category = "412" // H-动漫 H-Anime
	CategoryHComic              Category = "413" // H-漫画 H-Comic
	CategoryMovieHD             Category = "419" // 电影/HD Movie/HD
	CategoryMovieDVDiSo         Category = "420" // 电影/DVDiSo Movie/DVDiSo
	CategoryMovieBluRay         Category = "421" // 电影/BluRay Movie/BluRay
	CategorySoftware            Category = "422" // 软件 Software
	CategoryPCGame              Category = "423" // PC游戏 PCGame
	CategoryAVSDCensored        Category = "424" // AV(有码)/SD Censored AV(有碼)/SD Censored
	CategoryIVVideoCollection   Category = "425" // IV(写真影集) IV/Video Collection
	CategoryAVDVDiSoUncensored  Category = "426" // AV(无码)/DVDiSo Uncensored AV(無碼)/DVDiSo Uncensored
	CategoryEBook               Category = "427" // 電子書 E-Book
	CategoryAVHDUncensored      Category = "429" // AV(无码)/HD Uncensored AV(無碼)/HD Uncensored
	CategoryAVSDUncensored      Category = "430" // AV(无码)/SD Uncensored AV(無碼)/SD Uncensored
	CategoryAVBluRayCensored    Category = "431" // AV(有码)/Blu-Ray Censored AV(有碼)/Blu-Ray Censored
	CategoryAVBluRayUncensored  Category = "432" // AV(无码)/Blu-Ray Uncensored AV(無碼)/Blu-Ray Uncensored
	CategoryIVPictureCollection Category = "433" // IV(写真图集) IV/Picture Collection
	CategoryMusicLossless       Category = "434" // Music(无损) Music(Lossless)
	CategoryTVSeriesDVDiSo      Category = "435" // 影剧/综艺/DVDiSo TV Series/DVDiSo
	CategoryAV0Day              Category = "436" // AV(网站)/0Day AV(網站)/0Day
	CategoryAVDVDiSoCensored    Category = "437" // AV(有码)/DVDiSo Censored AV(有碼)/DVDiSo Censored
	CategoryTVSeriesBluRay      Category = "438" // 影剧/综艺/BluRay TV Series/BluRay
	CategoryMovieRemux          Category = "439" // 电影/Remux Movie/Remux
	CategoryAVGayHD             Category = "440" // AV(Gay)/HD
	CategoryAudioBook           Category = "442" // 有聲書 AuiBook
	CategoryDocumentary         Category = "444" // 紀錄 BBC
	CategoryIV                  Category = "445" // IV
	CategoryHACG                Category = "446" // H-ACG
	CategoryGame                Category = "447" // 遊戲
	CategoryTvGame              Category = "448" // TV遊戲 TvGame
	CategoryAnimeRoot           Category = "449" // 動漫 Anime
	CategoryMisc                Category = "450" // 其他
	CategoryEducation           Category = "451" // 教育影片
	CategoryAnimeBluRay         Category = "453" // 动画/BluRay Anime/BluRay
)

// 各搜索 mode 对应的分类 id 分组
var (
	CategoryGroupAdult     = []Category{CategoryAVHDCensored, CategoryAVHDUncensored, CategoryAVSDCensored, CategoryAVSDUncensored, CategoryAVDVDiSoUncensored, CategoryAVDVDiSoCensored, CategoryAVBluRayCensored, CategoryAVBluRayUncensored, CategoryAV0Day, CategoryIVVideoCollection, CategoryIVPictureCollection, CategoryHGame, CategoryHAnime, CategoryHComic, CategoryAVGayHD}
	CategoryGroupMovie     = []Category{CategoryMovieSD, CategoryMovieHD, CategoryMovieDVDiSo, CategoryMovieBluRay, CategoryMovieRemux}
	CategoryGroupMusic     = []Category{CategoryMV, CategoryMusicLossless}
	CategoryGroupTVShow    = []Category{CategoryTVSeriesSD, CategoryTVSeriesHD, CategoryTVSeriesDVDiSo, CategoryTVSeriesBluRay}
	CategoryGroupWaterfall = []Category{CategoryAVHDCensored, CategoryMovieSD, CategoryMovieHD, CategoryMovieDVDiSo, CategoryMovieBluRay, CategoryMovieRemux, CategoryTVSeriesHD, CategoryTVSeriesSD, CategoryTVSeriesDVDiSo, CategoryTVSeriesBluRay, CategoryMusicLossless, CategoryAVSDCensored, CategoryAVBluRayCensored, CategoryAVDVDiSoCensored, CategoryAVDVDiSoUncensored, CategoryAVHDUncensored, CategoryAVSDUncensored, CategoryAVBluRayUncensored, CategoryAV0Day, CategoryAVGayHD, CategoryRecord, CategoryAnime, CategoryMV, CategorySports, CategoryMiscOther, CategoryHGame, CategoryHAnime, CategoryHComic, CategorySoftware, CategoryPCGame, CategoryIVVideoCollection, CategoryEBook, CategoryIVPictureCollection, CategoryAudioBook, CategoryTvGame}
)
