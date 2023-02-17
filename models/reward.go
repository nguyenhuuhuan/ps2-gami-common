package models

import (
	"encoding/json"
	"time"
)

// RewardStatus is a string type.
type RewardStatus string

// Reward Define
const (
	RewardStatusActive   RewardStatus = "ACTIVE"
	RewardStatusInactive RewardStatus = "INACTIVE"
)

// String returns MerchantStatus as a string.
func (cs RewardStatus) String() string {
	return string(cs)
}

// RewardType define type
type RewardType string

// Reward Define
const (
	RewardTypeGem                   RewardType = "GEM"
	RewardTypeScore                 RewardType = "SCORE"
	RewardTypeWish                  RewardType = "WISH"
	RewardTypeCombo                 RewardType = "COMBO"
	RewardTypeLottery               RewardType = "LOTTERY"
	RewardTypeVoucher               RewardType = "VOUCHER"
	RewardTypePhysical              RewardType = "PHYSICAL"
	RewardTypeCashBack              RewardType = "CASHBACK"
	RewardTypeTCBCashBack           RewardType = "TCB_CASHBACK"
	RewardTypeLoyaltyPoint          RewardType = "LOYALTY_POINT"
	RewardTypeMiddleReward          RewardType = "MIDDLE_REWARD"
	RewardTypeUnlockVoucher         RewardType = "UNLOCK_VOUCHER"
	RewardTypeVoucherOrdering       RewardType = "VOUCHER_ORDERING"
	RewardTypeLoyaltyPointByProgram RewardType = "LOYALTY_POINT_BY_PROGRAM"
	RewardTypeStamp                 RewardType = "STAMP"
	RewardTypeOneFarm               RewardType = "ONE_FARM"
	RewardTypeCashVoucher           RewardType = "CASH_VOUCHER"
)

// LoyaltyPoint define type
const (
	LoyaltyPointFixed      = "FIXED"
	LoyaltyPointPercentage = "PERCENTAGE"
)

// CashBack define cash back reward type.
const (
	CashBackFixed      = "FIXED"
	CashBackPercentage = "PERCENTAGE"
)

// LoyaltyPointByProgram define type
const (
	LoyaltyPointByProgramFixed      = "FIXED"
	LoyaltyPointByProgramPercentage = "PERCENTAGE"
)

// CashVoucher define type
const (
	RedeemTypeFixed      = "FIXED"
	RedeemTypePercentage = "PERCENTAGE"
)

// String returns MerchantStatus as a string.
func (cs RewardType) String() string {
	return string(cs)
}

// Reward Model
type Reward struct {
	Base
	Code        string              `gorm:"<-:create" json:"code"`
	Type        RewardType          `json:"type"`
	Name        string              `json:"name"`
	Icon        string              `json:"icon"`
	Maker       string              `json:"maker"`
	Status      RewardStatus        `json:"status"`
	Checker     string              `json:"checker"`
	Content     string              `json:"content"`
	Deeplink    string              `json:"deeplink"`
	Description string              `json:"description"`
	Config      *JSON               `json:"config,omitempty"`
	Voucher     *Voucher            `json:"voucher,omitempty"`
	Combo       []*ComboItem        `json:"combo,omitempty"`
	RewardPool  *RewardPool         `json:"reward_pool,omitempty"`
	Wish        map[string]WishItem `gorm:"-"`
	TagID       int64               `json:"tag_id"`
	TenantID    int64               `json:"tenant_id"`
}

// Voucher model
type Voucher struct {
	BaseWithoutDeletedAt
	Code     string
	RewardID int64
	Image    string
	Reward   Reward `gorm:"foreignKey:RewardID"`
}

// TableName voucher
func (Voucher) TableName() string {
	return "voucher"
}

// Wish model
type Wish struct {
	ID       int64 `gorm:"primary_key" json:"id"`
	RewardID int64
	Reward   *Reward `json:"reward,omitempty"`
	Title    string
	Content  string
	Icon     string
}

// ComboItem model
type ComboItem struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	RewardID  int64
	Reward    *Reward `json:"reward,omitempty"`
	ItemID    int64
	UnitValue int
	CreatedAt time.Time
}

type WishData struct {
	Items []*WishItem `json:"items"`
}

type WishItem struct {
	Name        string `json:"name"`
	UniqueKey   string `json:"unique_key"`
	Description string `json:"description"`
}

func (r *Reward) ParseWishConfig() {
	if r.Config != nil {
		var wishData *WishData
		if err := json.Unmarshal(r.Config.MarshalJSON(), &wishData); err != nil {
			_ = json.Unmarshal(r.Config.MarshalJSON(), &wishData.Items)
		}

		r.Wish = make(map[string]WishItem, len(wishData.Items))
		for _, item := range wishData.Items {
			r.Wish[item.UniqueKey] = *item
		}
	}
}
