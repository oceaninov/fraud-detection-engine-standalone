package rprBlacklistMerchant

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityBlacklistMerchant struct {
	Id           string  `gorm:"column:id"`
	NMID         string  `gorm:"column:nmid"`
	MerchantName string  `gorm:"column:merchant_name"`
	Datasource   string  `gorm:"column:datasource"`
	FileId       string  `gorm:"column:file_id"`
	FileLink     string  `gorm:"column:file_link"`
	CreatedBy    string  `gorm:"column:created_by"`
	CreatedAt    string  `gorm:"column:created_at"`
	UpdatedBy    string  `gorm:"column:updated_by"`
	UpdatedAt    string  `gorm:"column:updated_at"`
	ApprovedBy   string  `gorm:"column:approved_by"`
	ApprovedAt   *string `gorm:"column:approved_at"`
}

func (ent *EntityBlacklistMerchant) TableName() string {
	return "black_list_merchant"
}

func (ent *EntityBlacklistMerchant) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.nmid", ent.TableName()),
		fmt.Sprintf("%s.merchant_name", ent.TableName()),
		fmt.Sprintf("%s.datasource", ent.TableName()),
		fmt.Sprintf("%s.file_id", ent.TableName()),
		fmt.Sprintf("%s.file_link", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistMerchant) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistMerchantFile struct {
	ID           string  `gorm:"column:id"`
	FileLink     string  `gorm:"column:file_link"`
	FileName     string  `gorm:"column:file_name"`
	FilePath     string  `gorm:"column:file_path"`
	ApprovalType string  `gorm:"column:approval_type"`
	Note         string  `gorm:"column:note"`
	Active       string  `gorm:"column:active"`
	Status       string  `gorm:"column:status"`
	CreatedBy    string  `gorm:"column:created_by"`
	CreatedAt    string  `gorm:"column:created_at"`
	UpdatedBy    string  `gorm:"column:updated_by"`
	UpdatedAt    string  `gorm:"column:updated_at"`
	ApprovedBy   string  `gorm:"column:approved_by"`
	ApprovedAt   *string `gorm:"column:approved_at"`
}

func (ent *EntityBlacklistMerchantFile) TableName() string {
	return "black_list_merchant_file"
}

func (ent *EntityBlacklistMerchantFile) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.file_link", ent.TableName()),
		fmt.Sprintf("%s.file_name", ent.TableName()),
		fmt.Sprintf("%s.file_path", ent.TableName()),
		fmt.Sprintf("%s.approval_type", ent.TableName()),
		fmt.Sprintf("%s.note", ent.TableName()),
		fmt.Sprintf("%s.active", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistMerchantFile) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistMerchantApproval struct {
	ID                  string  `gorm:"column:id"`
	BlacklistMerchantId string  `gorm:"column:black_list_merchant_id"`
	NMID                string  `gorm:"column:nmid"`
	MerchantName        string  `gorm:"column:merchant_name"`
	Datasource          string  `gorm:"column:datasource"`
	FileId              string  `gorm:"column:file_id"`
	FileLink            string  `gorm:"column:file_link"`
	ApprovalType        string  `gorm:"column:approval_type"`
	Note                string  `gorm:"column:note"`
	CreatedBy           string  `gorm:"column:created_by"`
	CreatedAt           string  `gorm:"column:created_at"`
	UpdatedBy           string  `gorm:"column:updated_by"`
	UpdatedAt           string  `gorm:"column:updated_at"`
	ApprovedBy          string  `gorm:"column:approved_by"`
	ApprovedAt          *string `gorm:"column:approved_at"`
}

func (ent *EntityBlacklistMerchantApproval) TableName() string {
	return "black_list_merchant_approval"
}

func (ent *EntityBlacklistMerchantApproval) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.black_list_merchant_id", ent.TableName()),
		fmt.Sprintf("%s.nmid", ent.TableName()),
		fmt.Sprintf("%s.merchant_name", ent.TableName()),
		fmt.Sprintf("%s.datasource", ent.TableName()),
		fmt.Sprintf("%s.file_id", ent.TableName()),
		fmt.Sprintf("%s.file_link", ent.TableName()),
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

func (ent *EntityBlacklistMerchantApproval) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
