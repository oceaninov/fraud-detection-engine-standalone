package rprKeyword

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

		WriteRowKeyword(ctx context.Context, entity EntityKeyword) (*EntityKeyword, *string, error)
		DeleteRowKeyword(ctx context.Context, id string) error
		UpdateRowKeyword(ctx context.Context, entity EntityKeyword) (*EntityKeyword, *string, error)

		WriteRowKeywordApproval(ctx context.Context, entity EntityKeywordApproval) (*EntityKeywordApproval, *string, error)
		DeleteRowKeywordApproval(ctx context.Context, id string) error
		UpdateRowKeywordApproval(ctx context.Context, entity EntityKeywordApproval) (*EntityKeywordApproval, *string, error)

		WriteApprovalForKeywordApproval(ctx context.Context, id, note, actionTaker string) error
		WriteRejectForKeywordApproval(ctx context.Context, id, note, actionTaker string) error

		ReadRowKeyword(ctx context.Context, findBy map[string]interface{}) (*EntityKeyword, error)
		ReadRowKeywordApproval(ctx context.Context, findBy map[string]interface{}) (*EntityKeywordApproval, error)
		ReadRowsKeywordApproval(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityKeywordApproval, *basicObject.Meta, error)
		ReadRowsKeywordApproved(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityKeyword, *basicObject.Meta, error)
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
	const fName = "repositories.rprKeyword.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.log = log
	bp.orm = orm
	return bp
}

func (b *blueprint) WriteRowKeyword(ctx context.Context, entity EntityKeyword) (*EntityKeyword, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.WriteRowKeyword"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowKeyword */
	model := new(EntityKeyword)
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
	return &entity, &entity.Id, nil
}

func (b *blueprint) UpdateRowKeyword(ctx context.Context, entity EntityKeyword) (*EntityKeyword, *string, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.UpdateRowKeyword"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UpdateRowKeyword */
	model := new(EntityKeyword)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		// Attempting to update the record
		if err := tx.Model(&model).Where("id = ?", entity.Id).Updates(&entity).Error; err != nil {
			_ = tx.Rollback()
			return err
		}
		return nil
	}); err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Returning the updated result
	return &entity, &entity.Id, nil
}

func (b *blueprint) DeleteRowKeyword(ctx context.Context, id string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.DeleteRowKeyword"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_DeleteRowKeyword */
	model := new(EntityKeyword)
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

func (b *blueprint) WriteRowKeywordApproval(ctx context.Context, entity EntityKeywordApproval) (*EntityKeywordApproval, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.WriteRowKeywordApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowKeywordApproval */
	model := new(EntityKeywordApproval)
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
	return &entity, &entity.Id, nil
}

func (b *blueprint) UpdateRowKeywordApproval(ctx context.Context, entity EntityKeywordApproval) (*EntityKeywordApproval, *string, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.UpdateRowKeywordApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UpdateRowKeywordApproval */
	model := new(EntityKeywordApproval)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		// Attempting to update the record
		if err := tx.Model(&model).Where("id = ?", entity.Id).Updates(&entity).Error; err != nil {
			_ = tx.Rollback()
			return err
		}
		return nil
	}); err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Returning the updated result
	return &entity, &entity.Id, nil
}

func (b *blueprint) DeleteRowKeywordApproval(ctx context.Context, id string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.DeleteRowKeywordApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_DeleteRowKeywordApproval */
	model := new(EntityKeywordApproval)
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

func (b *blueprint) ReadRowsKeywordApproval(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityKeywordApproval, *basicObject.Meta, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.ReadRowsKeywordApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsKeywordApproval */
	model := new(EntityKeywordApproval)
	var entities []EntityKeywordApproval

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

	// Apply filters from findBy map
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}

	// Apply search filter with WHERE LIKE
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("keyword LIKE ?", searchPattern)
	}

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

func (b *blueprint) ReadRowsKeywordApproved(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityKeyword, *basicObject.Meta, error) {
	// Constructing standard module name with request ID as its prefix
	const fName = "repositories.rprKeyword.ReadRowsKeywordApproved"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsKeywordApproved */
	model := new(EntityKeyword)
	var entities []EntityKeyword

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

	// Apply filters from findBy map
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}

	// Apply search filter with WHERE LIKE
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("keyword LIKE ?", searchPattern)
	}

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

func (b *blueprint) ReadRowKeywordApproval(ctx context.Context, findBy map[string]interface{}) (*EntityKeywordApproval, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.ReadRowKeywordApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowKeywordApproval */
	model := new(EntityKeywordApproval)
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

func (b *blueprint) ReadRowKeyword(ctx context.Context, findBy map[string]interface{}) (*EntityKeyword, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprKeyword.ReadRowKeyword"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowKeyword */
	model := new(EntityKeyword)
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

func (b *blueprint) WriteApprovalForKeywordApproval(ctx context.Context, id, note, actionTaker string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.WriteApprovalForKeywordApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRejectForKeywordApproval */
	model := new(EntityKeywordApproval)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		// Attempting to update the record
		if err := tx.Model(&model).Where("id = ?", id).Updates(map[string]interface{}{
			"note":        note,
			"approved_by": actionTaker,
			"approved_at": time.Now(),
			"status":      basicObject.ApprovalApproved,
		}).Error; err != nil {
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

func (b *blueprint) WriteRejectForKeywordApproval(ctx context.Context, id, note, actionTaker string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.WriteRejectForKeywordApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRejectForKeywordReject */
	model := new(EntityKeywordApproval)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		// Attempting to update the record
		if err := tx.Model(&model).Where("id = ?", id).Updates(map[string]interface{}{
			"note":        note,
			"approved_by": actionTaker,
			"approved_at": time.Now(),
			"status":      basicObject.ApprovalRejected,
		}).Error; err != nil {
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
