package rprBlacklistHistory

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityBlacklistHistory struct {
	Id               string `gorm:"column:id"`
	PhoneNumber      string `gorm:"column:phone_number"`
	Event            string `gorm:"column:event"`
	BeneficiaryName  string `gorm:"column:beneficiary_name"`
	CreatedAt        string `gorm:"column:created_at"`
	CreatedBy        string `gorm:"column:created_by"`
	UpdatedAt        string `gorm:"column:updated_at"`
	UpdatedBy        string `gorm:"column:updated_by"`
	ApprovedAt       string `gorm:"column:approved_at"`
	ApprovedBy       string `gorm:"column:approved_by"`
	TransactionTypes string `gorm:"column:transaction_types"`
}

func (ent *EntityBlacklistHistory) TableName() string {
	return "black_list_history"
}

func (ent *EntityBlacklistHistory) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.phone_number", ent.TableName()),
		fmt.Sprintf("%s.event", ent.TableName()),
		fmt.Sprintf("%s.beneficiary_name", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.transaction_types", ent.TableName()),
	}, ", ")
}
