package rprSof

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntitySOF struct {
	Id         string `gorm:"column:id"`
	SofName    string `gorm:"column:sof_name"`
	SofStatus  string `gorm:"column:sof_status"`
	CreatedAt  string `gorm:"column:created_at"`
	CreatedBy  string `gorm:"column:created_by"`
	UpdatedAt  string `gorm:"column:updated_at"`
	UpdatedBy  string `gorm:"column:updated_by"`
	ApprovedAt string `gorm:"column:approved_at"`
	ApprovedBy string `gorm:"column:approved_by"`
}

func (ent *EntitySOF) TableName() string {
	return "sofs"
}

func (ent *EntitySOF) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.sof_name", ent.TableName()),
		fmt.Sprintf("%s.sof_status", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
	}, ", ")
}
