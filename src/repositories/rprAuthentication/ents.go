package rprAuthentication

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/entityExtractor"
	"strings"
	"time"
)

/* [CODE GENERATOR] INITIALIZE_CODE */

type EntityUser struct {
	Id        string `gorm:"column:id"`
	RoleId    int    `gorm:"column:role_id"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	FullName  string `gorm:"column:full_name"`
	AvatarUrl string `gorm:"column:avatar_url"`
	Gender    string `gorm:"column:gender"`
	Activated string `gorm:"column:activated"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
	//CreatedBy  string `gorm:"column:created_by"`
	//UpdatedBy  string `gorm:"column:updated_by"`
	//ApprovedAt string `gorm:"column:approved_at"`
	//ApprovedBy string `gorm:"column:approved_by"`
}

func (ent *EntityUser) TableName() string {
	return "users"
}

func (ent *EntityUser) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.role_id", ent.TableName()),
		fmt.Sprintf("%s.email", ent.TableName()),
		fmt.Sprintf("%s.password", ent.TableName()),
		fmt.Sprintf("%s.full_name", ent.TableName()),
		fmt.Sprintf("%s.avatar_url", ent.TableName()),
		fmt.Sprintf("%s.gender", ent.TableName()),
		fmt.Sprintf("%s.activated", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		//fmt.Sprintf("%s.created_by", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		//fmt.Sprintf("%s.updated_by", ent.TableName()),
		//fmt.Sprintf("%s.approved_at", ent.TableName()),
		//fmt.Sprintf("%s.approved_by", ent.TableName()),
	}, ", ")
}

func (ent *EntityUser) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityRole struct {
	Id          string `gorm:"column:id"`
	Title       string `gorm:"column:title"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedAt   string `gorm:"column:updated_at"`
	Description string `gorm:"column:description"`
}

func (ent *EntityRole) TableName() string {
	return "roles"
}

func (ent *EntityRole) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.title", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
		fmt.Sprintf("%s.updated_at", ent.TableName()),
		fmt.Sprintf("%s.description", ent.TableName()),
	}, ", ")
}

func (ent *EntityRole) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}

type EntityResetPassword struct {
	Id        int       `gorm:"column:id"`
	UserId    string    `gorm:"column:user_id"`
	Token     string    `gorm:"column:token"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (ent *EntityResetPassword) TableName() string {
	return "reset_password"
}

func (ent *EntityResetPassword) DefaultColumns() string {
	return strings.Join([]string{
		fmt.Sprintf("%s.id", ent.TableName()),
		fmt.Sprintf("%s.user_id", ent.TableName()),
		fmt.Sprintf("%s.token", ent.TableName()),
		fmt.Sprintf("%s.expired_at", ent.TableName()),
		fmt.Sprintf("%s.created_at", ent.TableName()),
	}, ", ")
}

func (ent *EntityResetPassword) ToMap() map[string]interface{} {
	return entityExtractor.ConvertEntityToMap(ent)
}
