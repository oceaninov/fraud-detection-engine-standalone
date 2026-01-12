package rprChannel

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityChannel struct {
	Id            string `gorm:"column:id"`
	ChannelName   string `gorm:"column:channel_name"`
	ChannelStatus string `gorm:"column:channel_status"`
	CreatedAt     string `gorm:"column:created_at"`
	CreatedBy     string `gorm:"column:created_by"`
	UpdatedAt     string `gorm:"column:updated_at"`
	UpdatedBy     string `gorm:"column:updated_by"`
	ApprovedAt    string `gorm:"column:approved_at"`
	ApprovedBy    string `gorm:"column:approved_by"`
}

func (ent *EntityChannel) TableName() string {
	return "registered_channel"
}

func (ent *EntityChannel) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.channel_name", ent.TableName()),
		fmt.Sprintf("%s.channel_status", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
		fmt.Sprintf("%s.approved_by", ent.TableName()),
	}, ", ")
}

func (ent *EntityChannel) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
