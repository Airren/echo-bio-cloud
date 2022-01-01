package model

import (
	"time"
)

type OrderType string

const (
	TypeBugFix     OrderType = "bug_fix"
	TypeNewFeature OrderType = "new_feature"
	TypeImprove    OrderType = "improve"
)

type OrderPriority string

const (
	PriorityZero  OrderPriority = "P0"
	PriorityOne   OrderPriority = "P1"
	PriorityTwo   OrderPriority = "P2"
	PriorityThree OrderPriority = "P3"
	PriorityFour  OrderPriority = "P4"
)

type OrderStatus string

const (
	Done       OrderStatus = "done"
	Fixed      OrderStatus = "fixed"
	Processing OrderStatus = "processing"
	Pending    OrderStatus = "pending"
	Pause      OrderStatus = "pause"
	Blocked    OrderStatus = "blocked"
)

type Order struct {
	RecordMeta

	Name           string       `json:"name"           gorm:"type:varchar(255);not null;"`
	PeriodId       int64        `json:"period_id"`
	Type           OrderType     `json:"type"           gorm:"type:varchar(64);not null"`
	Priority       OrderPriority `json:"priority"       gorm:"type:varchar(64);not null"`
	Claimer        string       `json:"claimer"        gorm:"type:varchar(64);not null"`
	Operator       string       `json:"operator"       gorm:"type:varchar(64)"`
	TimeCost       float64      `json:"time_cost"`
	DeadLine       time.Time    `json:"dead_line" `
	Status         OrderStatus   `json:"status"          gorm:"type:varchar(64);not null"`
	BackendStatus  bool         `json:"backend_status"  gorm:"not null"`
	FrontendStatus bool         `json:"frontend_status" gorm:"not null"`
	CheckStatus    bool         `json:"check_status"    gorm:"not null"`
	Comment        string       `json:"comment"         gorm:"type:text;not null"`
	Description    string       `json:"description"     gorm:"type:text;not null"`
	Image          string       `json:"image"           gorm:"type:varchar(64)"`
}
