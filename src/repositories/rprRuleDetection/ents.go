package rprRuleDetection

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityRule struct {
	Id              string  `gorm:"column:id"`
	RuleName        string  `gorm:"column:rule_name"`
	Types           string  `gorm:"column:type"`
	TransactionType string  `gorm:"column:transaction_type"`
	Interval        string  `gorm:"column:interval"`
	Amount          string  `gorm:"column:amount"`
	Actions         string  `gorm:"column:actions"`
	Status          string  `gorm:"column:status"`
	TimeRangeType   string  `gorm:"column:time_range_type"`
	StartTimeRange  string  `gorm:"column:start_time_range"`
	EndTimeRange    string  `gorm:"column:end_time_range"`
	Sofs            string  `gorm:"column:sofs"`
	Channel         string  `gorm:"column:channel"`
	CreatedBy       string  `gorm:"column:created_by"`
	CreatedAt       string  `gorm:"column:created_at"`
	UpdatedBy       string  `gorm:"column:updated_by"`
	UpdatedAt       string  `gorm:"column:updated_at"`
	ApprovedBy      string  `gorm:"column:approved_by"`
	ApprovedAt      *string `gorm:"column:approved_at"`
}

func (ent *EntityRule) TableName() string {
	return "rules"
}

func (ent *EntityRule) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.rule_name", ent.TableName()),
		fmt.Sprintf("%s.type", ent.TableName()),
		fmt.Sprintf("%s.transaction_type", ent.TableName()),
		fmt.Sprintf("%s.interval", ent.TableName()),
		fmt.Sprintf("%s.amount", ent.TableName()),
		fmt.Sprintf("%s.actions", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.time_range_type", ent.TableName()),
		fmt.Sprintf("%s.start_time_range", ent.TableName()),
		fmt.Sprintf("%s.end_time_range", ent.TableName()),
		fmt.Sprintf("%s.sofs", ent.TableName()),
		fmt.Sprintf("%s.channel", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityRule) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityRuleApproval struct {
	Id              string  `gorm:"column:id"`
	RuleName        string  `gorm:"column:rule_name"`
	Types           string  `gorm:"column:type"`
	TransactionType string  `gorm:"column:transaction_type"`
	Interval        string  `gorm:"column:interval"`
	Amount          float64 `gorm:"column:amount"`
	Actions         string  `gorm:"column:actions"`
	Status          string  `gorm:"column:status"`
	TimeRangeType   string  `gorm:"column:time_range_type"`
	StartTimeRange  string  `gorm:"column:start_time_range"`
	EndTimeRange    string  `gorm:"column:end_time_range"`
	Note            string  `gorm:"column:note"`
	Sofs            string  `gorm:"column:sofs"`
	Channel         string  `gorm:"column:channel"`
	RuleId          string  `gorm:"column:rule_id"`
	ApprovalType    string  `gorm:"column:approval_type"`
	CreatedBy       string  `gorm:"column:created_by"`
	CreatedAt       string  `gorm:"column:created_at"`
	UpdatedBy       string  `gorm:"column:updated_by"`
	UpdatedAt       string  `gorm:"column:updated_at"`
	ApprovedBy      string  `gorm:"column:approved_by"`
	ApprovedAt      *string `gorm:"column:approved_at"`
}

func (ent *EntityRuleApproval) TableName() string {
	return "rules_approval"
}

func (ent *EntityRuleApproval) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.rule_name", ent.TableName()),
		fmt.Sprintf("%s.type", ent.TableName()),
		fmt.Sprintf("%s.transaction_type", ent.TableName()),
		fmt.Sprintf("%s.interval", ent.TableName()),
		fmt.Sprintf("%s.amount", ent.TableName()),
		fmt.Sprintf("%s.actions", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.time_range_type", ent.TableName()),
		fmt.Sprintf("%s.start_time_range", ent.TableName()),
		fmt.Sprintf("%s.end_time_range", ent.TableName()),
		fmt.Sprintf("%s.note", ent.TableName()),
		fmt.Sprintf("%s.sofs", ent.TableName()),
		fmt.Sprintf("%s.channel", ent.TableName()),
		fmt.Sprintf("%s.rule_id", ent.TableName()),
		fmt.Sprintf("%s.approval_type", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityRuleApproval) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
