package rprRuleDetection

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

		WriteRowRuleDetection(ctx context.Context, entity EntityRule) (*EntityRule, *string, error)
		DeleteRowRuleDetection(ctx context.Context, id string) error
		UpdateRowRuleDetection(ctx context.Context, entity EntityRule) (*EntityRule, *string, error)

		WriteRowRuleDetectionApproval(ctx context.Context, entity EntityRuleApproval) (*EntityRuleApproval, *string, error)
		DeleteRowRuleDetectionApproval(ctx context.Context, id string) error
		UpdateRowRuleDetectionApproval(ctx context.Context, entity EntityRuleApproval) (*EntityRuleApproval, *string, error)

		WriteApprovalForRuleDetectionApproval(ctx context.Context, id, note, actionTaker string) error
		WriteRejectForRuleDetectionApproval(ctx context.Context, id, note, actionTaker string) error

		ReadRowRuleDetection(ctx context.Context, findBy map[string]interface{}) (*EntityRule, error)
		ReadRowRuleDetectionApproved(ctx context.Context, findBy map[string]interface{}) (*EntityRule, error)
		ReadRowRuleDetectionApproval(ctx context.Context, findBy map[string]interface{}) (*EntityRuleApproval, error)
		ReadRowsRuleDetectionApproval(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityRuleApproval, *basicObject.Meta, error)
		ReadRowsRuleDetectionApproved(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityRule, *basicObject.Meta, error)
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
	const fName = "repositories.rprRuleDetection.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.log = log
	bp.orm = orm
	return bp
}

func (b *blueprint) WriteRowRuleDetection(ctx context.Context, entity EntityRule) (*EntityRule, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.WriteRowRuleDetection"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowRuleDetection */
	model := new(EntityRule)
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

func (b *blueprint) UpdateRowRuleDetection(ctx context.Context, entity EntityRule) (*EntityRule, *string, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.UpdateRowRuleDetection"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UpdateRowRuleDetection */
	model := new(EntityRule)
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

func (b *blueprint) DeleteRowRuleDetection(ctx context.Context, id string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.DeleteRowRuleDetection"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_DeleteRowRuleDetection */
	model := new(EntityRule)
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

func (b *blueprint) WriteRowRuleDetectionApproval(ctx context.Context, entity EntityRuleApproval) (*EntityRuleApproval, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.WriteRowRuleDetectionApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowRuleDetectionApproval */
	model := new(EntityRuleApproval)
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

func (b *blueprint) UpdateRowRuleDetectionApproval(ctx context.Context, entity EntityRuleApproval) (*EntityRuleApproval, *string, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.UpdateRowRuleDetectionApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UpdateRowRuleDetectionApproval */
	model := new(EntityRuleApproval)
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

func (b *blueprint) DeleteRowRuleDetectionApproval(ctx context.Context, id string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.DeleteRowRuleDetectionApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_DeleteRowRuleDetectionApproval */
	model := new(EntityRuleApproval)
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

func (b *blueprint) ReadRowsRuleDetectionApproval(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityRuleApproval, *basicObject.Meta, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.ReadRowsRuleDetectionApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsRuleDetectionApproval */
	model := new(EntityRuleApproval)
	var entities []EntityRuleApproval

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
		query = query.Where("rule_name LIKE ?", searchPattern)
	}

	// Count total records for pagination
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Fetch the records with pagination
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

	// returning result
	return &entities, meta, nil
}

func (b *blueprint) ReadRowsRuleDetectionApproved(ctx context.Context, findBy map[string]interface{}, page, limit int, search string) (*[]EntityRule, *basicObject.Meta, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.ReadRowsRuleDetectionApproved"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsRuleDetectionApproved */
	model := new(EntityRule)
	var entities []EntityRule

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
		query = query.Where("rule_name LIKE ?", searchPattern)
	}

	// Count total records for pagination
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Fetch the records with pagination
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

	// returning result
	return &entities, meta, nil
}

func (b *blueprint) ReadRowRuleDetectionApproved(ctx context.Context, findBy map[string]interface{}) (*EntityRule, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.ReadRowRuleDetectionApproved"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowRuleDetectionApproved */
	model := new(EntityRule)
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

func (b *blueprint) ReadRowRuleDetectionApproval(ctx context.Context, findBy map[string]interface{}) (*EntityRuleApproval, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.ReadRowRuleDetectionApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowRuleDetectionApproval */
	model := new(EntityRuleApproval)
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

func (b *blueprint) ReadRowRuleDetection(ctx context.Context, findBy map[string]interface{}) (*EntityRule, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprRuleDetection.ReadRowRuleDetection"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowRuleDetection */
	model := new(EntityRule)
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

func (b *blueprint) WriteApprovalForRuleDetectionApproval(ctx context.Context, id, note, actionTaker string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.WriteApprovalForRuleDetectionApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRejectForRuleDetectionApproval */
	model := new(EntityRuleApproval)
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

func (b *blueprint) WriteRejectForRuleDetectionApproval(ctx context.Context, id, note, actionTaker string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.WriteRejectForRuleDetectionApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRejectForRuleDetectionReject */
	model := new(EntityRuleApproval)
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
