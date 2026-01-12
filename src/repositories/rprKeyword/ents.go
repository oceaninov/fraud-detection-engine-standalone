package rprKeyword

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityKeyword struct {
	Id         string  `gorm:"column:id"`
	Keyword    string  `gorm:"column:keyword"`
	Action     string  `gorm:"column:action"`
	Channel    string  `gorm:"column:channel"`
	CreatedBy  string  `gorm:"column:created_by"`
	CreatedAt  string  `gorm:"column:created_at"`
	UpdatedBy  string  `gorm:"column:updated_by"`
	UpdatedAt  string  `gorm:"column:updated_at"`
	ApprovedBy string  `gorm:"column:approved_by"`
	ApprovedAt *string `gorm:"column:approved_at"`
}

func (ent *EntityKeyword) TableName() string {
	return "keyword"
}

func (ent *EntityKeyword) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.keyword", ent.TableName()),
		fmt.Sprintf("%s.action", ent.TableName()),
		fmt.Sprintf("%s.channel", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityKeyword) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityKeywordApproval struct {
	Id           string  `gorm:"column:id"`
	KeywordId    string  `gorm:"column:keyword_id"`
	Keyword      string  `gorm:"column:keyword"`
	Action       string  `gorm:"column:action"`
	Channel      string  `gorm:"column:channel"`
	ApprovalType string  `gorm:"column:approval_type"`
	Note         string  `gorm:"column:note"`
	Status       string  `gorm:"column:status"`
	CreatedBy    string  `gorm:"column:created_by"`
	CreatedAt    string  `gorm:"column:created_at"`
	UpdatedBy    string  `gorm:"column:updated_by"`
	UpdatedAt    string  `gorm:"column:updated_at"`
	ApprovedBy   string  `gorm:"column:approved_by"`
	ApprovedAt   *string `gorm:"column:approved_at"`
}

func (ent *EntityKeywordApproval) TableName() string {
	return "keyword_approval"
}

func (ent *EntityKeywordApproval) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.keyword_id", ent.TableName()),
		fmt.Sprintf("%s.keyword", ent.TableName()),
		fmt.Sprintf("%s.action", ent.TableName()),
		fmt.Sprintf("%s.channel", ent.TableName()),
		fmt.Sprintf("%s.approval_type", ent.TableName()),
		fmt.Sprintf("%s.note", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityKeywordApproval) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
