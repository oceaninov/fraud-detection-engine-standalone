package rprBlacklistDTTOT

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityBlacklistDTTOT struct {
	Id         string  `gorm:"column:id"`
	PPATKId    string  `gorm:"column:ppatk_id"`
	Name       string  `gorm:"column:name"`
	Bod        string  `gorm:"column:bod"`
	Datasource string  `gorm:"column:datasource"`
	FileId     string  `gorm:"column:file_id"`
	FileLink   string  `gorm:"column:file_link"`
	NIK        string  `gorm:"column:nik"`
	CreatedBy  string  `gorm:"column:created_by"`
	CreatedAt  string  `gorm:"column:created_at"`
	UpdatedBy  string  `gorm:"column:updated_by"`
	UpdatedAt  string  `gorm:"column:updated_at"`
	ApprovedBy string  `gorm:"column:approved_by"`
	ApprovedAt *string `gorm:"column:approved_at"`
}

func (ent *EntityBlacklistDTTOT) TableName() string {
	return "black_list_dttot"
}

func (ent *EntityBlacklistDTTOT) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.ppatk_id", ent.TableName()),
		fmt.Sprintf("%s.name", ent.TableName()),
		fmt.Sprintf("%s.bod", ent.TableName()),
		fmt.Sprintf("%s.datasource", ent.TableName()),
		fmt.Sprintf("%s.file_id", ent.TableName()),
		fmt.Sprintf("%s.file_link", ent.TableName()),
		fmt.Sprintf("%s.nik", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistDTTOT) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistDTTOTFile struct {
	ID           string  `gorm:"column:id"`
	FileLink     string  `gorm:"column:file_link"`
	FileName     string  `gorm:"column:file_name"`
	FilePath     string  `gorm:"column:file_path"`
	ApprovalType string  `gorm:"column:approval_type"`
	Note         string  `gorm:"column:note"`
	Status       string  `gorm:"column:status"`
	Active       string  `gorm:"column:active"`
	CreatedBy    string  `gorm:"column:created_by"`
	CreatedAt    string  `gorm:"column:created_at"`
	UpdatedBy    string  `gorm:"column:updated_by"`
	UpdatedAt    string  `gorm:"column:updated_at"`
	ApprovedBy   string  `gorm:"column:approved_by"`
	ApprovedAt   *string `gorm:"column:approved_at"`
}

func (ent *EntityBlacklistDTTOTFile) TableName() string {
	return "black_list_dttot_file"
}

func (ent *EntityBlacklistDTTOTFile) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.file_link", ent.TableName()),
		fmt.Sprintf("%s.file_name", ent.TableName()),
		fmt.Sprintf("%s.file_path", ent.TableName()),
		fmt.Sprintf("%s.note", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.active", ent.TableName()),
		fmt.Sprintf("%s.approval_type", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistDTTOTFile) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistDTTOTApproval struct {
	ID               string  `gorm:"column:id"`
	BlacklistDTTOTId string  `gorm:"column:black_list_dttot_id"`
	PPATKId          string  `gorm:"column:ppatk_id"`
	Name             string  `gorm:"column:name"`
	NIK              string  `gorm:"column:nik"`
	Bod              string  `gorm:"column:bod"`
	Datasource       string  `gorm:"column:datasource"`
	FileId           string  `gorm:"column:file_id"`
	FileLink         string  `gorm:"column:file_link"`
	ApprovalType     string  `gorm:"column:approval_type"`
	Note             string  `gorm:"column:note"`
	CreatedBy        string  `gorm:"column:created_by"`
	CreatedAt        string  `gorm:"column:created_at"`
	UpdatedBy        string  `gorm:"column:updated_by"`
	UpdatedAt        string  `gorm:"column:updated_at"`
	ApprovedBy       string  `gorm:"column:approved_by"`
	ApprovedAt       *string `gorm:"column:approved_at"`
}

func (ent *EntityBlacklistDTTOTApproval) TableName() string {
	return "black_list_dttot_approval"
}

func (ent *EntityBlacklistDTTOTApproval) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.black_list_dttot_id", ent.TableName()),
		fmt.Sprintf("%s.ppatk_id", ent.TableName()),
		fmt.Sprintf("%s.name", ent.TableName()),
		fmt.Sprintf("%s.nik", ent.TableName()),
		fmt.Sprintf("%s.bod", ent.TableName()),
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

func (ent *EntityBlacklistDTTOTApproval) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
