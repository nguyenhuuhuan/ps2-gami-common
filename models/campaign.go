package models

import (
	"time"
)

// CampaignStatus is a string type.
type CampaignStatus string

// CampaignRuleRequirement is a string type
type CampaignRuleRequirement string

// String returns CampaignStatus as a string.
func (cs CampaignStatus) String() string {
	return string(cs)
}

// String returns CampaignRuleRequirement as a string.
func (cs CampaignRuleRequirement) String() string {
	return string(cs)
}

// CampaignType is a string type.
type CampaignType string

// String returns MerchantStatus as a string.
func (cs CampaignType) String() string {
	return string(cs)
}

// CampaignRequestStatus is a string type.
type CampaignRequestStatus string

// String returns MerchantStatus as a string.
func (cs CampaignRequestStatus) String() string {
	return string(cs)
}

// App type
const (
	AppVinID   = "VINID"
	AppVinShop = "VINSHOP"
)

// Campaign status enum definition
const (
	CampaignStatusActive   CampaignStatus = "ACTIVE"
	CampaignStatusEnd      CampaignStatus = "END"
	CampaignStatusSuspend  CampaignStatus = "SUSPEND"
	CampaignStatusInactive CampaignStatus = "INACTIVE"
)

// Campaign Rule Requirement Type enum definition
const (
	CampaignPassOneRule CampaignRuleRequirement = "PASS_ONE_RULE"
	CampaignPassAllRule CampaignRuleRequirement = "PASS_ALL_RULE"
)

// Campaign status enum definition
const (
	CampaignTypeLottery                 CampaignType = "LOTTERY"
	CampaignTypeRegularQuest            CampaignType = "QUEST"
	CampaignTypeMysteryBox              CampaignType = "MYSTERY_BOX"
	CampaignTypeLuckySpin               CampaignType = "LUCKY_SPIN"
	CampaignTypeCollection              CampaignType = "COLLECTION"
	CampaignTypeQuestReferral           CampaignType = "QUEST_REFERRAL"
	CampaignTypeReferral                CampaignType = "REFERRAL"
	CampaignTypeReferralVinhomes        CampaignType = "VINHOMES"
	CampaignTypeGiveaway                CampaignType = "GIVEAWAY"
	CampaignTypeFlamingo                CampaignType = "FLAMINGO"
	CampaignTypeItemSet                 CampaignType = "ITEM_SET"
	CampaignTypeIndirectPromotion       CampaignType = "INDIRECT_PROMOTION"
	CampaignTypeDirectPromotion         CampaignType = "DIRECT_PROMOTION"
	CampaignTypeChallenge               CampaignType = "CHALLENGE"
	CampaignTypeScratchCard             CampaignType = "SCRATCH_CARD"
	CampaignTypeDailyCheckInProgression CampaignType = "DAILY_CHECK_IN_PROGRESSION"
	CampaignTypeChallengeReferral       CampaignType = "CHALLENGE_REFERRAL"
	CampaignTypeCollectionV2            CampaignType = "COLLECTION_V2"
	CampaignLeaderboard                 CampaignType = "LEADERBOARD"
	CampaignQuiz                        CampaignType = "QUIZ"
	CampaignTypeOneFarm                 CampaignType = "ONE_FARM"
	CampaignTypeMatching                CampaignType = "MATCHING"
	CampaignTypeSlotMachine             CampaignType = "SLOT_MACHINE"
)

// Campaign models
type Campaign struct {
	Base
	TenantID        int64                   `json:"tenant_id"`
	Name            string                  `json:"name"`
	Code            string                  `gorm:"<-:create" json:"code"`
	Description     *string                 `json:"description,omitempty"`
	Content         string                  `json:"content"`
	Banner          string                  `json:"banner"`
	DateStart       time.Time               `json:"date_start"`
	DateEnd         time.Time               `json:"date_end"`
	Status          CampaignStatus          `json:"status"`
	Type            CampaignType            `gorm:"<-:create" json:"type"`
	Maker           string                  `json:"maker"`
	Checker         string                  `json:"checker"`
	Rule            []*Rule                 `json:"rule,omitempty"`
	Pool            []*Pool                 `json:"pool,omitempty"`
	RuleRequirement CampaignRuleRequirement `json:"rule_requirement"`
	BlackWhiteList  *BlackWhiteList         `json:"black_white_list,omitempty"`
	Config          *JSON                   `sql:"type:json" json:"config,omitempty"`
	AppType         string                  `json:"app_type"`
	Purpose         string                  `json:"purpose"`
	IOCode          string                  `json:"io_code"`
}

// DateStartStr function to convert to string when get detail
func (m *Campaign) DateStartStr() string {
	return m.DateStart.Format(DateFormat)
}

// DateEndStr function to convert to string when get detail
func (m *Campaign) DateEndStr() string {
	return m.DateEnd.Format(DateFormat)
}

// DateEndTimestamp Convert Date Closed to timestamp
func (m Campaign) DateEndTimestamp() int64 {
	return m.DateEnd.Unix()
}
