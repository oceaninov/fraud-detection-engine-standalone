package rprBlacklistSender

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityBlacklistSender struct {
	Id               string  `gorm:"column:id"`
	PhoneNumber      string  `gorm:"column:phone_number"`
	Status           string  `gorm:"column:status"`
	BeneficiaryName  string  `gorm:"column:beneficiary_name"`
	TransactionTypes string  `gorm:"column:transaction_types"`
	CreatedBy        string  `gorm:"column:created_by"`
	CreatedAt        string  `gorm:"column:created_at"`
	UpdatedBy        string  `gorm:"column:updated_by"`
	UpdatedAt        string  `gorm:"column:updated_at"`
	ApprovedBy       string  `gorm:"column:approved_by"`
	ApprovedAt       *string `gorm:"column:approved_at"`
}

func (ent *EntityBlacklistSender) TableName() string {
	return "black_list_sender"
}

func (ent *EntityBlacklistSender) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.phone_number", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.beneficiary_name", ent.TableName()),
		fmt.Sprintf("%s.transaction_types", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistSender) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistSenderApproval struct {
	Id               string  `gorm:"column:id"`
	PhoneNumber      string  `gorm:"column:phone_number"`
	ApprovalType     string  `gorm:"column:approval_type"`
	Status           string  `gorm:"column:status"`
	Note             string  `gorm:"column:note"`
	Event            string  `gorm:"column:event"`
	BlacklistId      string  `gorm:"column:blacklist_id"`
	BeneficiaryName  string  `gorm:"column:beneficiary_name"`
	TransactionTypes string  `gorm:"column:transaction_types"`
	CreatedBy        string  `gorm:"column:created_by"`
	CreatedAt        string  `gorm:"column:created_at"`
	UpdatedBy        string  `gorm:"column:updated_by"`
	UpdatedAt        string  `gorm:"column:updated_at"`
	ApprovedBy       string  `gorm:"column:approved_by"`
	ApprovedAt       *string `gorm:"column:approved_at"`
}

func (ent *EntityBlacklistSenderApproval) TableName() string {
	return "black_list_sender_approval"
}

func (ent *EntityBlacklistSenderApproval) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.phone_number", ent.TableName()),
		fmt.Sprintf("%s.approval_type", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.note", ent.TableName()),
		fmt.Sprintf("%s.event", ent.TableName()),
		fmt.Sprintf("%s.blacklist_id", ent.TableName()),
		fmt.Sprintf("%s.beneficiary_name", ent.TableName()),
		fmt.Sprintf("%s.transaction_types", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistSenderApproval) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
