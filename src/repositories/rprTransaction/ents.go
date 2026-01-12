package rprTransaction

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"strings"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityTransaction struct {
	Id              string  `gorm:"column:id"`
	TransactionId   string  `gorm:"column:transaction_id"`
	TransactionType string  `gorm:"column:transaction_type"`
	Rules           string  `gorm:"column:rules"`
	Title           string  `gorm:"column:title"`
	Channel         string  `gorm:"column:channel"`
	BodyReq         *string `gorm:"column:body_req"`
	FlagId          string  `gorm:"column:flag_id"`
	CreatedAt       string  `gorm:"column:created_at"`
	UserId          string  `gorm:"column:user_id"`
	Amount          string  `gorm:"column:amount"`
	DestinationId   string  `gorm:"column:destination_id"`
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
		fmt.Sprintf("%s.destination_id", ent.TableName()),
	}, ", ")
}

type EntityFlag struct {
	Id    string `gorm:"column:id"`
	Title string `gorm:"column:title"`
}

func (ent *EntityFlag) TableName() string {
	return "flag"
}

func (ent *EntityFlag) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.title", ent.TableName()),
	}, ", ")
}
