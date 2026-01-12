package rprFraudDetection

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityBlacklistDTTOT struct {
	Id         string `gorm:"column:id"`
	PPATKId    string `gorm:"column:ppatk_id"`
	Name       string `gorm:"column:name"`
	Bod        string `gorm:"column:bod"`
	Datasource string `gorm:"column:datasource"`
	FileId     string `gorm:"column:file_id"`
	FileLink   string `gorm:"column:file_link"`
	CreatedBy  string `gorm:"column:created_by"`
	CreatedAt  string `gorm:"column:created_at"`
	UpdatedAt  string `gorm:"column:updated_at"`
	ApprovedAt string `gorm:"column:approved_at"`
	NIK        string `gorm:"column:nik"`
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
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
		fmt.Sprintf("%s.nik", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistDTTOT) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistMerchant struct {
	Id           string `gorm:"column:id"`
	NMID         string `gorm:"column:nmid"`
	MerchantName string `gorm:"column:merchant_name"`
	Datasource   string `gorm:"column:datasource"`
	FileId       string `gorm:"column:file_id"`
	FileLink     string `gorm:"column:file_link"`
	CreatedBy    string `gorm:"column:created_by"`
	CreatedAt    string `gorm:"column:created_at"`
	UpdatedAt    string `gorm:"column:updated_at"`
	ApprovedAt   string `gorm:"column:approved_at"`
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
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistMerchant) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistReceiver struct {
	Id               string `gorm:"column:id"`
	PhoneNumber      string `gorm:"column:phone_number"`
	CreatedAt        string `gorm:"column:created_at"`
	CreatedBy        string `gorm:"column:created_by"`
	UpdatedAt        string `gorm:"column:updated_at"`
	UpdatedBy        string `gorm:"column:updated_by"`
	Status           int    `gorm:"column:status"`
	BeneficiaryName  string `gorm:"column:beneficiary_name"`
	TransactionTypes string `gorm:"column:transaction_types"`
}

func (ent *EntityBlacklistReceiver) TableName() string {
	return "black_list_receiver"
}

func (ent *EntityBlacklistReceiver) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.phone_number", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.beneficiary_name", ent.TableName()),
		fmt.Sprintf("%s.transaction_types", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistReceiver) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistSender struct {
	Id               string `gorm:"column:id"`
	PhoneNumber      string `gorm:"column:phone_number"`
	CreatedAt        string `gorm:"column:created_at"`
	CreatedBy        string `gorm:"column:created_by"`
	UpdatedAt        string `gorm:"column:updated_at"`
	UpdatedBy        string `gorm:"column:updated_by"`
	Status           int    `gorm:"column:status"`
	BeneficiaryName  string `gorm:"column:beneficiary_name"`
	TransactionTypes string `gorm:"column:transaction_types"`
}

func (ent *EntityBlacklistSender) TableName() string {
	return "black_list_sender"
}

func (ent *EntityBlacklistSender) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.phone_number", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.updated_by", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.beneficiary_name", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistSender) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityRule struct {
	Id              string  `gorm:"column:id"`
	RuleName        string  `gorm:"column:rule_name"`
	Types           string  `gorm:"column:type"`
	TransactionType string  `gorm:"column:transaction_type"`
	Interval        string  `gorm:"column:interval"`
	Amount          float64 `gorm:"column:amount"`
	Actions         string  `gorm:"column:actions"`
	CreatedAt       string  `gorm:"column:created_at"`
	UpdatedAt       string  `gorm:"column:updated_at"`
	Status          int     `gorm:"column:status"`
	TimeRangeType   string  `gorm:"column:time_range_type"`
	StartTimeRange  string  `gorm:"column:start_time_range"`
	EndTimeRange    string  `gorm:"column:end_time_range"`
	Sofs            string  `gorm:"column:sofs"`
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
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.status", ent.TableName()),
		fmt.Sprintf("%s.time_range_type", ent.TableName()),
		fmt.Sprintf("%s.start_time_range", ent.TableName()),
		fmt.Sprintf("%s.end_time_range", ent.TableName()),
		fmt.Sprintf("%s.sofs", ent.TableName()),
	}, ", ")
}

func (ent *EntityRule) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityTransaction struct {
	Id              string `gorm:"column:id"`
	TransactionId   string `gorm:"column:transaction_id"`
	TransactionType string `gorm:"column:transaction_type"`
	Rules           string `gorm:"column:rules"`
	Title           string `gorm:"column:title"`
	Channel         string `gorm:"column:channel"`
	BodyReq         string `gorm:"column:body_req"`
	FlagId          string `gorm:"column:flag_id"`
	CreatedAt       string `gorm:"column:created_at"`
	UserId          string `gorm:"column:user_id"`
	Amount          string `gorm:"column:amount"`
	DestinationId   string `gorm:"column:destination_id"`
}

func (ent *EntityTransaction) TableName() string {
	return "transaction"
}

func (ent *EntityTransaction) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.transaction_id", ent.TableName()),
		fmt.Sprintf("%s.transaction_type", ent.TableName()),
		fmt.Sprintf("%s.rules", ent.TableName()),
		fmt.Sprintf("%s.title", ent.TableName()),
		fmt.Sprintf("%s.channel", ent.TableName()),
		fmt.Sprintf("%s.body_req", ent.TableName()),
		fmt.Sprintf("%s.flag_id", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.user_id", ent.TableName()),
		fmt.Sprintf("%s.amount", ent.TableName()),
	}, ", ")
}

func (ent *EntityTransaction) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityLog struct {
	Id              string  `gorm:"column:id"`
	UserId          string  `gorm:"column:user_id"`
	Amount          float64 `gorm:"column:amount"`
	StartDate       string  `gorm:"column:start_date"`
	BodyReq         string  `gorm:"column:body_req"`
	Channel         string  `gorm:"column:channel"`
	TransactionType string  `gorm:"column:transaction_type"`
	EndDate         string  `gorm:"column:end_date"`
	DestinationId   string  `gorm:"column:destination_id"`
}

func (ent *EntityLog) TableName() string {
	return "log"
}

func (ent *EntityLog) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.user_id", ent.TableName()),
		fmt.Sprintf("%s.amount", ent.TableName()),
		fmt.Sprintf("%s.start_date", ent.TableName()),
		fmt.Sprintf("%s.body_req", ent.TableName()),
		fmt.Sprintf("%s.channel", ent.TableName()),
		fmt.Sprintf("%s.transaction_type", ent.TableName()),
		fmt.Sprintf("%s.end_date", ent.TableName()),
		fmt.Sprintf("%s.destination_id", ent.TableName()),
	}, ", ")
}

func (ent *EntityLog) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityBlacklistHistory struct {
	Id               string `gorm:"column:id"`
	PhoneNumber      string `gorm:"column:phone_number"`
	CreatedAt        string `gorm:"column:created_at"`
	Event            string `gorm:"column:event"`
	CreatedBy        string `gorm:"column:created_by"`
	BeneficiaryName  string `gorm:"column:beneficiary_name"`
	TransactionTypes string `gorm:"column:transaction_types"`
}

func (ent *EntityBlacklistHistory) TableName() string {
	return "black_list_history"
}

func (ent *EntityBlacklistHistory) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.phone_number", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.event", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.beneficiary_name", ent.TableName()),
		fmt.Sprintf("%s.transaction_types", ent.TableName()),
	}, ", ")
}

func (ent *EntityBlacklistHistory) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityKeyword struct {
	Id         string `gorm:"column:id"`
	Keyword    string `gorm:"column:keyword"`
	Action     string `gorm:"column:action"`
	ApprovedAt string `gorm:"column:approved_at"`
	CreatedBy  string `gorm:"column:created_by"`
	UpdatedAt  string `gorm:"column:updated_at"`
}

func (ent *EntityKeyword) TableName() string {
	return "keyword"
}

func (ent *EntityKeyword) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.keyword", ent.TableName()),
		fmt.Sprintf("%s.action", ent.TableName()),
		fmt.Sprintf("%s.approved_at", ent.TableName()),
		fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityKeyword) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
