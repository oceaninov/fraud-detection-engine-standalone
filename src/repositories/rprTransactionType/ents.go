package rprTransactionType

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityTransactionType struct {
	Id         string  `gorm:"column:id"`
	Name       string  `gorm:"column:name"`
	CreatedBy  string  `gorm:"column:created_by"`
	CreatedAt  string  `gorm:"column:created_at"`
	UpdatedBy  string  `gorm:"column:updated_by"`
	UpdatedAt  string  `gorm:"column:updated_at"`
	ApprovedBy string  `gorm:"column:approved_by"`
	ApprovedAt *string `gorm:"column:approved_at"`
}

func (ent *EntityTransactionType) TableName() string {
	return "transaction_type"
}

func (ent *EntityTransactionType) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.name", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityTransactionType) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityTransactionTypeApproval struct {
	Id                string  `gorm:"column:id"`
	TransactionTypeId string  `gorm:"column:transaction_type_id"`
	Name              string  `gorm:"column:name"`
	Status            string  `gorm:"column:status"`
	ApprovalType      string  `gorm:"column:approval_type"`
	Note              string  `gorm:"column:note"`
	CreatedBy         string  `gorm:"column:created_by"`
	CreatedAt         string  `gorm:"column:created_at"`
	UpdatedBy         string  `gorm:"column:updated_by"`
	UpdatedAt         string  `gorm:"column:updated_at"`
	ApprovedBy        string  `gorm:"column:approved_by"`
	ApprovedAt        *string `gorm:"column:approved_at"`
}

func (ent *EntityTransactionTypeApproval) TableName() string {
	return "transaction_type_approval"
}

func (ent *EntityTransactionTypeApproval) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.transaction_type_id", ent.TableName()),
		fmt.Sprintf("%s.name", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.approval_type", ent.TableName()),
		fmt.Sprintf("%s.note", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityTransactionTypeApproval) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
