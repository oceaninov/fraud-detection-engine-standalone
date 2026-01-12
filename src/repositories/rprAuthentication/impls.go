package rprAuthentication

import (
	"context"
	"errors"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint repository function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		WriteUser(ctx context.Context, entity *EntityUser) (*EntityUser, *string, error)
		ReadRowUser(ctx context.Context, findBy map[string]interface{}) (*EntityUser, error)
		DeleteRowUser(ctx context.Context, id string) error
		ReadRowRoles(ctx context.Context) (*[]EntityRole, error)
		ReadRowUsers(ctx context.Context, page, limit int, fullName, email, gender string, roleId int) (*[]EntityUser, *basicObject.Meta, error)

		ReadRowUserByEmail(ctx context.Context, email string) (*EntityUser, error)
		WriteRowResetPassword(ctx context.Context, userID, token string) (int64, error)
		ReadRowResetPasswordByTokenLatest(ctx context.Context, token string) (*EntityResetPassword, error)
		UpdateResetPassword(ctx context.Context, resetPasswordID int64) error
		UpdateUserPasswordByID(ctx context.Context, userID string, hashedPassword string) error
		ReadResetPasswordByTokenLatest(ctx context.Context, token string) (*EntityResetPassword, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		log *zap.SugaredLogger // log logging instance
		orm *gorm.DB           // orm database gorm orm instance
	}
)

func NewBlueprint(
	/* [CODE GENERATOR] FUNC_PARAM */
	log *zap.SugaredLogger,
	orm *gorm.DB,
) Blueprint {
	const fName = "repositories.rprBlacklistDTTOT.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.log = log
	bp.orm = orm
	return bp
}

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */

func (b *blueprint) ReadRowUsers(ctx context.Context, page, limit int, fullName, email, gender string, roleId int) (*[]EntityUser, *basicObject.Meta, error) {
	// Constructing standard module name with request ID as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowUsers"
	requestId, _ := ctx.Value(defaultHeaders.XRequestId).(string) // Handle missing request ID gracefully
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowUsers */
	model := new(EntityUser)
	var entities []EntityUser

	// Calculate offset based on page and limit
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // Default limit if not specified
	}
	offset := (page - 1) * limit

	// Build the query with pagination
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())

	// Count total records for pagination
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Fetch the records with pagination and order by created_at descending
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Apply filters only if the values are not empty or zero
	if fullName != "" {
		query = query.Where("full_name LIKE ?", "%"+fullName+"%") // Partial match support
	}
	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%") // Partial match support
	}
	if gender != "" {
		query = query.Where("gender = ?", gender)
	}
	if roleId > 0 { // Ensure roleId is valid before filtering
		query = query.Where("role_id = ?", roleId)
	}

	// Fetch filtered results, ordered by `created_at DESC`
	if err := query.Order("created_at DESC").Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Calculate pagination metadata
	meta := &basicObject.Meta{
		Count:         strconv.Itoa(int(totalCount)),
		CurrentPage:   strconv.Itoa(page),
		LastPage:      strconv.Itoa(int((totalCount + int64(limit) - 1) / int64(limit))),
		RecordPerPage: strconv.Itoa(limit),
	}

	// Returning result
	return &entities, meta, nil
}

func (b *blueprint) ReadRowRoles(ctx context.Context) (*[]EntityRole, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowRoles"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowRoles */
	model := new(EntityRole)
	var entities []EntityRole
	// Build the query with pagination
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if err := query.Order("created_at DESC").Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}
	// returning result
	return &entities, nil
}

func (b *blueprint) DeleteRowUser(ctx context.Context, id string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.DeleteRowUser"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_DeleteRowUser */
	model := new(EntityUser)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		// Attempting to update the record
		if err := tx.Model(&model).Where("id = ?", id).Delete(model).Error; err != nil {
			_ = tx.Rollback()
			return err
		}
		return nil
	}); err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil
	}

	// Returning the updated result
	return nil
}

func (b *blueprint) WriteUser(ctx context.Context, entity *EntityUser) (*EntityUser, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteUser"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteUser */
	model := new(EntityUser)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&entity).Error; err != nil {
			_ = tx.Rollback()
			return err
		}
		return nil
	}); err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// returning result
	return entity, &entity.Id, nil
}

func (b *blueprint) ReadRowUser(ctx context.Context, findBy map[string]interface{}) (*EntityUser, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowUser"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowUser */
	model := new(EntityUser)
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}
	if err := query.First(model).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			logging.Errorw(fName, "reason", err.Error())
			return nil, nil
		}
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return model, nil
}

func (b *blueprint) ReadRowUserByEmail(ctx context.Context, email string) (*EntityUser, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprUser.ReadRowUserByEmail"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowUserByEmail */
	model := new(EntityUser)

	// Build the query
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns()).Where("email = ?", email)

	// Execute query
	if err := query.First(model).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// Returning result
	return model, nil
}

func (b *blueprint) WriteRowResetPassword(ctx context.Context, userID, token string) (int64, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprUser.WriteRowResetPassword"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowResetPassword */
	now := time.Now()
	expired := now.Add(time.Hour * 1)

	// Prepare reset password entity
	model := &EntityResetPassword{
		UserId:    userID,
		Token:     token,
		ExpiredAt: expired,
		CreatedAt: now,
	}

	// Execute insert query
	if err := b.orm.Table(model.TableName()).Create(model).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return 0, err
	}

	// Returning inserted ID
	return int64(model.Id), nil
}

func (b *blueprint) ReadRowResetPasswordByTokenLatest(ctx context.Context, token string) (*EntityResetPassword, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprUser.ReadRowResetPasswordByTokenLatest"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowResetPasswordByTokenLatest */
	var entity EntityResetPassword

	// Fetch latest reset password entry by token
	if err := b.orm.Table(entity.TableName()).
		Where("token = ?", token).
		Order("created_at DESC").
		Limit(1).
		Find(&entity).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// Returning result
	return &entity, nil
}

func (b *blueprint) UpdateResetPassword(ctx context.Context, resetPasswordID int64) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprUser.UpdateResetPassword"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UpdateResetPassword */
	// Set expiration time to 1 hour ago
	expiredAt := time.Now().Add(time.Hour * -1)

	// Execute update query
	if err := b.orm.Table("reset_password").
		Where("id = ?", resetPasswordID).
		Update("expired_at", expiredAt).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return err
	}

	// Returning success
	return nil
}

func (b *blueprint) UpdateUserPasswordByID(ctx context.Context, userID string, hashedPassword string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprUser.UpdateUserPasswordByID"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UpdateUserPasswordByID */
	// Update user password and set updated_at timestamp
	updateData := map[string]interface{}{
		"password":   hashedPassword,
		"updated_at": time.Now(),
	}

	// Execute update query
	if err := b.orm.Table("users").
		Where("id = ?", userID).
		Updates(updateData).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return err
	}

	// Returning success
	return nil
}

func (b *blueprint) ReadResetPasswordByTokenLatest(ctx context.Context, token string) (*EntityResetPassword, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprResetPassword.ReadResetPasswordByTokenLatest"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadResetPasswordByTokenLatest */
	model := new(EntityResetPassword)

	// Execute query to fetch latest reset password entry by token
	if err := b.orm.Table(model.TableName()).
		Where("token = ?", token).
		Order("created_at DESC").
		Limit(1).
		Find(model).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// Returning result
	return model, nil
}
