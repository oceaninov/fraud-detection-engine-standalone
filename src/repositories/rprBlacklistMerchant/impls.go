package rprBlacklistMerchant

import (
	"context"
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/csvProcessor"
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

		WriteBulkRowBlacklistMerchant(ctx context.Context, entities []EntityBlacklistMerchantApproval) error
		WriteBulkRowBlacklistMerchantApproval(ctx context.Context, fileId, createdBy string, entity []csvProcessor.MerchantCSVFormat) error
		WriteRowBlacklistMerchantFile(ctx context.Context, entity EntityBlacklistMerchantFile) (*EntityBlacklistMerchantFile, *string, error)

		WriteRejectForBlacklistMerchantApproval(ctx context.Context, id, note, actionTaker string) error
		WriteApprovalForBlacklistMerchantApproval(ctx context.Context, id, note, actionTaker string) error

		ReadRowBlacklistMerchantFile(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistMerchantFile, error)
		ReadRowsBlacklistMerchantFile(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistMerchantFile, *basicObject.Meta, error)
		ReadRowsBlacklistMerchantApproval(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistMerchantApproval, error)
		ReadRowsBlacklistMerchantApproved(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistMerchant, error)
		ReadRowsBlacklistMerchantApprovedQuery(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistMerchant, *basicObject.Meta, error)
		ReadRowBlacklistMerchantApprovedSingle(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistMerchant, error)
		ReadRowBlacklistMerchantFileSingle(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistMerchantFile, error)
		ReadRowsBlacklistMerchantApprovalQuery(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistMerchantApproval, *basicObject.Meta, error)

		RemoveBlacklistMerchantApprovalData(ctx context.Context, fileId string) error
		RemoveBlacklistMerchantData(ctx context.Context, fileId string) error
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
	const fName = "repositories.rprBlacklistMerchant.NewBlueprint"
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

func (b *blueprint) WriteRejectForBlacklistMerchantApproval(ctx context.Context, id, note, actionTaker string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.WriteRejectForBlacklistMerchantApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRejectForBlacklistMerchantApproval */
	model := new(EntityBlacklistMerchantFile)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		// Attempting to update the record
		if err := tx.Model(&model).Where("id = ?", id).Updates(map[string]interface{}{
			"note":        note,
			"approved_by": actionTaker,
			"approved_at": time.Now(),
			"status":      basicObject.ApprovalRejected,
			"active":      "1",
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

func (b *blueprint) WriteApprovalForBlacklistMerchantApproval(ctx context.Context, id, note, actionTaker string) error {
	const fName = "repositories.rprBlacklistSender.WriteApprovalForBlacklistMerchantApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	model := new(EntityBlacklistMerchantFile)
	err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model).Where("id = ?", id).Updates(map[string]interface{}{
			"note":        note,
			"approved_by": actionTaker,
			"approved_at": time.Now(),
			"status":      basicObject.ApprovalApproved,
			"active":      "1",
		}).Error
		if err != nil {
			return fmt.Errorf("update failed: %w", err)
		}
		return nil
	})
	if err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return err
	}
	return nil
}

func (b *blueprint) ReadRowBlacklistMerchantFile(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistMerchantFile, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowBlacklistMerchantFile"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowBlacklistMerchantFile */
	model := new(EntityBlacklistMerchantFile)
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

func (b *blueprint) ReadRowsBlacklistMerchantFile(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistMerchantFile, *basicObject.Meta, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistMerchantFile"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistMerchantFile */
	model := new(EntityBlacklistMerchantFile)
	var entities []EntityBlacklistMerchantFile

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

func (b *blueprint) WriteBulkRowBlacklistMerchant(ctx context.Context, entities []EntityBlacklistMerchantApproval) error {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteBulkRowBlacklistMerchantApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteBulkRowBlacklistMerchantApproval */
	model := new(EntityBlacklistMerchant)
	for _, entityData := range entities {
		if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
			ct := time.Now().Format(basicObject.DateAndTime)
			if err := tx.Create(&EntityBlacklistMerchant{
				Id:           entityData.BlacklistMerchantId,
				NMID:         entityData.NMID,
				MerchantName: entityData.MerchantName,
				Datasource:   entityData.Datasource,
				FileId:       entityData.FileId,
				FileLink:     entityData.FileLink,
				CreatedAt:    ct,
				CreatedBy:    entityData.CreatedBy,
				UpdatedAt:    ct,
				UpdatedBy:    entityData.UpdatedBy,
				ApprovedAt:   &ct,
				ApprovedBy:   entityData.ApprovedBy,
			}).Error; err != nil {
				_ = tx.Rollback()
				return err
			}
			return nil
		}); err != nil {
			logging.Errorw(fName, "reason", err.Error())
			return err
		}
	}

	// returning result
	return nil
}

func (b *blueprint) WriteBulkRowBlacklistMerchantApproval(ctx context.Context, fileId, createdBy string, entity []csvProcessor.MerchantCSVFormat) error {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteBulkRowBlacklistMerchantApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteBulkRowBlacklistMerchantApproval */
	model := new(EntityBlacklistMerchantApproval)
	for _, entityData := range entity {
		if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&EntityBlacklistMerchantApproval{
				ID:                  guuid.NewString(),
				NMID:                entityData.NMID,
				MerchantName:        entityData.Name,
				Datasource:          entityData.DataSource,
				FileId:              fileId,
				FileLink:            basicObject.BlankString,
				BlacklistMerchantId: guuid.NewString(),
				Note:                basicObject.BlankString,
				ApprovalType:        basicObject.CreateApprovalType,
				ApprovedBy:          basicObject.BlankString,
				ApprovedAt:          nil,
				CreatedBy:           createdBy,
				CreatedAt:           time.Now().Format(basicObject.DateAndTime),
				UpdatedBy:           createdBy,
				UpdatedAt:           time.Now().Format(basicObject.DateAndTime),
			}).Error; err != nil {
				_ = tx.Rollback()
				return err
			}
			return nil
		}); err != nil {
			logging.Errorw(fName, "reason", err.Error())
			return err
		}
	}

	// returning result
	return nil
}

func (b *blueprint) WriteRowBlacklistMerchantFile(ctx context.Context, entity EntityBlacklistMerchantFile) (*EntityBlacklistMerchantFile, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteRowBlacklistMerchantFile"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowBlacklistMerchantFile */
	model := new(EntityBlacklistMerchantFile)
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
	return &entity, &entity.ID, nil
}

func (b *blueprint) ReadRowsBlacklistMerchantApproval(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistMerchantApproval, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistMerchantApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistMerchantApproval */
	model := new(EntityBlacklistMerchantApproval)
	var entities []EntityBlacklistMerchantApproval
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}
	if err := query.Find(&entities).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			logging.Errorw(fName, "reason", err.Error())
			return nil, nil
		}
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return &entities, nil
}

func (b *blueprint) RemoveBlacklistMerchantApprovalData(ctx context.Context, fileId string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.RemoveBlacklistMerchantApprovalData"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RemoveBlacklistMerchantApprovalData */
	model := new(EntityBlacklistMerchantApproval)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		// Attempting to update the record
		if err := tx.Model(&model).Where("file_id = ?", fileId).Delete(model).Error; err != nil {
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

func (b *blueprint) RemoveBlacklistMerchantData(ctx context.Context, fileId string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.RemoveBlacklistMerchantData"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RemoveBlacklistMerchantData */
	model := new(EntityBlacklistMerchant)
	if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
		// Attempting to update the record
		if err := tx.Model(&model).Where("file_id != ?", fileId).Delete(model).Error; err != nil {
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

func (b *blueprint) ReadRowsBlacklistMerchantApproved(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistMerchant, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistMerchantApproved"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistMerchantApproved */
	model := new(EntityBlacklistMerchant)
	var entities []EntityBlacklistMerchant

	// Build the query with pagination
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}
	if err := query.Order("created_at DESC").Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsBlacklistMerchantApprovedQuery(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistMerchant, *basicObject.Meta, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistMerchantApproved"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistMerchantApproved */
	model := new(EntityBlacklistMerchant)
	var entities []EntityBlacklistMerchant

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

func (b *blueprint) ReadRowBlacklistMerchantApprovedSingle(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistMerchant, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowBlacklistMerchantApprovedSingle"
	requestId, _ := ctx.Value(defaultHeaders.XRequestId).(string) // Handle missing request ID gracefully
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowBlacklistMerchantApprovedSingle */
	model := new(EntityBlacklistMerchant)

	// Build the query with filters
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())

	// Apply filters from findBy map
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}

	// Fetch a single record ordered by `created_at DESC`
	if err := query.Order("created_at DESC").First(model).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// Return the found record
	return model, nil
}

func (b *blueprint) ReadRowBlacklistMerchantFileSingle(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistMerchantFile, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowBlacklistMerchantFileSingle"
	requestId, _ := ctx.Value(defaultHeaders.XRequestId).(string) // Handle missing request ID gracefully
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowBlacklistMerchantFileSingle */
	model := new(EntityBlacklistMerchantFile)

	// Build the query with filters
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())

	// Apply filters from findBy map
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}

	// Fetch a single record ordered by `created_at DESC`
	if err := query.Order("created_at DESC").First(model).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// Return the found record
	return model, nil
}

func (b *blueprint) ReadRowsBlacklistMerchantApprovalQuery(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistMerchantApproval, *basicObject.Meta, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistMerchantApprovalQuery"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistMerchantApprovalQuery */
	model := new(EntityBlacklistMerchantApproval)
	var entities []EntityBlacklistMerchantApproval

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
