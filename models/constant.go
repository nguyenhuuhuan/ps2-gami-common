package models

import "time"

// List Const Define
const (
	DateFormat                          string        = "2006-01-02T15:04:05.000Z"
	MySQLDateFormat                     string        = "2006-01-02 15:04:05"
	LIMIT                               int           = 50
	ASC                                 string        = "asc"
	DESC                                string        = "desc"
	DuplicateCode                       uint16        = 1062
	CacheExpiredDefault                 time.Duration = 60 * 60 * 24 * 7 * time.Second // 1 weeks
	CacheExpiredOneDay                  time.Duration = 24 * 60 * 60 * time.Second     // 1 hour
	CacheExpired20Minute                time.Duration = 20 * 60 * time.Second          // 1 hour
	AndroidPlatform                     string        = "android"
	IOSPlatform                         string        = "ios"
	AllForcedUpdateCacheKey             string        = "all_forced_update_cache_key"
	NumOfSecondsRemindBeforeEndCampaign int64         = 60 * 60 * 24 * 2 // Seconds
	NumOfSecondsRemindBefore            int           = 60 * 60 * 3      // Seconds
	EnvProd                             string        = "PRODUCTION"
)

// Redis Key
const (
	RedisEventRemoveKey = "event:redis:remove"
)

// List Mystery Box config key
const (
	BackgroundImage       string = "background_image"
	LeaderBoardBackground string = "leader_board_background"
	HistoryBackground     string = "history_background"
	OpenGiftImage         string = "open_gift_image"
	Title                 string = "title"
	Guide                 string = "guide"
	SubTitleLeaderBoard   string = "sub_title_leader_board"
	TitleHidden           string = "title_hidden"
	EnableGemLeaderBoard  string = "enable_gem_leader_board"
	LightStatusBar        string = "light_status_bar"
	TapOpenGift           string = "tap_open_gift"
	BoxImageType          string = "box_image_type"
	BoxImages             string = "box_images"
	TotalBox              string = "total_box"
	MaximumTurn           string = "maximum_turn"
	RuleLink              string = "rule_link"
	RuleImages            string = "rule_images"
	GemIcon               string = "gem_icon"
	TurnIcon              string = "turn_icon"
	NoTurnPopup           string = "no_turn_popup"
	OutOfTurnPopup        string = "out_of_turn_popup"
	NewTurnPopup          string = "new_turn_popup"
	LeaderBoardSize       string = "leader_board_size"
	TurnTab               string = "turn_tab"
	QuestRewardName       string = "reward_name"
	QuestRewardIcon       string = "reward_icon"
	QuestBanner           string = "quest_banner"
	HaveTurnNoti          string = "have_turn_noti"
	FullTurnNoti          string = "full_turn_noti"
	TitleNoti             string = "title"
	MessageNoti           string = "message"
)
