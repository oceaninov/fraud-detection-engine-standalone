package rprBlacklistReceiver

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

		ReadRowsBlacklistReceiver(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityBlacklistReceiver, *basicObject.Meta, error)
		ReadRowBlacklistReceiver(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistReceiver, error)
		ReadRowsBlacklistReceiverApproval(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityBlacklistReceiverApproval, *basicObject.Meta, error)
		ReadRowBlacklistReceiverApproval(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistReceiverApproval, error)

		WriteRowBlacklistReceiverApproval(ctx context.Context, entity EntityBlacklistReceiverApproval) (*EntityBlacklistReceiverApproval, *string, error)
		WriteRowBlacklistReceiver(ctx context.Context, entity EntityBlacklistReceiver) (*EntityBlacklistReceiver, *string, error)
		WriteRejectForBlacklistApproval(ctx context.Context, id, note, actionTaker string) error
		WriteApprovalForBlacklistApproval(ctx context.Context, id, note, actionTaker string) error

		RemoveRowBlacklistReceiver(ctx context.Context, id string) error
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
	const fName = "repositories.rprFraudDetection.NewBlueprint"
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

func (b *blueprint) ReadRowsBlacklistReceiver(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityBlacklistReceiver, *basicObject.Meta, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.ReadRowsBlacklistReceiver"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistReceiver */
	model := new(EntityBlacklistReceiver)
	var entities []EntityBlacklistReceiver

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
		query = query.Where("phone_number LIKE ?", searchPattern)
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

func (b *blueprint) ReadRowBlacklistReceiver(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistReceiver, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.ReadRowBlacklistReceiver"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowBlacklistReceiver */
	model := new(EntityBlacklistReceiver)
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

func (b *blueprint) ReadRowsBlacklistReceiverApproval(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityBlacklistReceiverApproval, *basicObject.Meta, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.ReadRowsBlacklistReceiverApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistReceiverApproval */
	model := new(EntityBlacklistReceiverApproval)
	var entities []EntityBlacklistReceiverApproval

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
		query = query.Where("phone_number LIKE ?", searchPattern)
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

func (b *blueprint) ReadRowBlacklistReceiverApproval(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistReceiverApproval, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.ReadRowBlacklistReceiverApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowBlacklistReceiverApproval */
	model := new(EntityBlacklistReceiverApproval)
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

func (b *blueprint) WriteRowBlacklistReceiverApproval(ctx context.Context, entity EntityBlacklistReceiverApproval) (*EntityBlacklistReceiverApproval, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.WriteRowBlacklistReceiverApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowBlacklistReceiverApproval */
	model := new(EntityBlacklistReceiverApproval)
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

func (b *blueprint) WriteRowBlacklistReceiver(ctx context.Context, entity EntityBlacklistReceiver) (*EntityBlacklistReceiver, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.WriteRowBlacklistReceiver"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowBlacklistReceiver */
	model := new(EntityBlacklistReceiver)
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

func (b *blueprint) RemoveRowBlacklistReceiver(ctx context.Context, id string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.RemoveRowBlacklistReceiver"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RemoveRowBlacklistReceiver */
	model := new(EntityBlacklistReceiver)
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

func (b *blueprint) WriteRejectForBlacklistApproval(ctx context.Context, id, note, actionTaker string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.WriteRejectForBlacklistApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRejectForBlacklistApproval */
	model := new(EntityBlacklistReceiverApproval)
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

func (b *blueprint) WriteApprovalForBlacklistApproval(ctx context.Context, id, note, actionTaker string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.WriteApprovalForBlacklistApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteApprovalForBlacklistApproval */
	model := new(EntityBlacklistReceiverApproval)
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
