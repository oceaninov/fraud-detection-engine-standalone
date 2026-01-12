package rprBlacklistDTTOT

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

		WriteBulkRowBlacklistDTTOT(ctx context.Context, entities []EntityBlacklistDTTOTApproval) error
		WriteBulkRowBlacklistDTTOTApproval(ctx context.Context, fileId, createdBy string, entity []csvProcessor.DTTOTCSVFormat) error
		WriteRowBlacklistDTTOTFile(ctx context.Context, entity EntityBlacklistDTTOTFile) (*EntityBlacklistDTTOTFile, *string, error)
		WriteRejectForBlacklistDTTOTApproval(ctx context.Context, id, note, actionTaker string) error
		WriteApprovalForBlacklistDTTOTApproval(ctx context.Context, id, note, actionTaker string) error

		ReadRowBlacklistDTTOTFile(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistDTTOTFile, error)
		ReadRowsBlacklistDTTOTFile(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistDTTOTFile, *basicObject.Meta, error)
		ReadRowsBlacklistDTTOTApproval(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistDTTOTApproval, error)
		ReadRowsBlacklistDTTOTApproved(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistDTTOT, error)
		ReadRowsBlacklistDTTOTApprovedQuery(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistDTTOT, *basicObject.Meta, error)
		ReadRowBlacklistDTTOTApprovedSingle(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistDTTOT, error)
		ReadRowBlacklistDTTOTFileSingle(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistDTTOTFile, error)
		ReadRowsBlacklistDTTOTApprovalQuery(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistDTTOTApproval, *basicObject.Meta, error)

		RemoveBlacklistDTTOTApprovalData(ctx context.Context, fileId string) error
		RemoveBlacklistDTTOTData(ctx context.Context, fileId string) error
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

func (b *blueprint) WriteRejectForBlacklistDTTOTApproval(ctx context.Context, id, note, actionTaker string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.WriteRejectForBlacklistDTTOTApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRejectForBlacklistDTTOTApproval */
	model := new(EntityBlacklistDTTOTFile)
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

func (b *blueprint) WriteApprovalForBlacklistDTTOTApproval(ctx context.Context, id, note, actionTaker string) error {
	const fName = "repositories.rprBlacklistSender.WriteApprovalForBlacklistDTTOTApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	model := new(EntityBlacklistDTTOTFile)
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

func (b *blueprint) ReadRowBlacklistDTTOTFile(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistDTTOTFile, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowBlacklistDTTOTFile"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowBlacklistDTTOTFile */
	model := new(EntityBlacklistDTTOTFile)
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

func (b *blueprint) ReadRowsBlacklistDTTOTFile(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistDTTOTFile, *basicObject.Meta, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistDTTOTFile"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistDTTOTFile */
	model := new(EntityBlacklistDTTOTFile)
	var entities []EntityBlacklistDTTOTFile

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

func (b *blueprint) WriteBulkRowBlacklistDTTOT(ctx context.Context, entities []EntityBlacklistDTTOTApproval) error {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteBulkRowBlacklistDTTOTApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteBulkRowBlacklistDTTOTApproval */
	model := new(EntityBlacklistDTTOT)
	for _, entityData := range entities {
		if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
			ct := time.Now().Format(basicObject.DateAndTime)
			if err := tx.Create(&EntityBlacklistDTTOT{
				Id:         entityData.BlacklistDTTOTId,
				PPATKId:    entityData.PPATKId,
				Name:       entityData.Name,
				NIK:        entityData.NIK,
				Bod:        entityData.Bod,
				Datasource: entityData.Datasource,
				FileId:     entityData.FileId,
				FileLink:   entityData.FileLink,
				CreatedAt:  ct,
				CreatedBy:  entityData.CreatedBy,
				UpdatedAt:  ct,
				UpdatedBy:  entityData.UpdatedBy,
				ApprovedAt: &ct,
				ApprovedBy: entityData.ApprovedBy,
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

func (b *blueprint) WriteBulkRowBlacklistDTTOTApproval(ctx context.Context, fileId, createdBy string, entity []csvProcessor.DTTOTCSVFormat) error {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteBulkRowBlacklistDTTOTApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteBulkRowBlacklistDTTOTApproval */
	model := new(EntityBlacklistDTTOTApproval)
	for _, entityData := range entity {
		if err := b.orm.Table(model.TableName()).Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&EntityBlacklistDTTOTApproval{
				ID:               guuid.NewString(),
				NIK:              entityData.NIK,
				PPATKId:          entityData.PPATKID,
				Name:             entityData.Name,
				Bod:              entityData.DateOfBirth,
				Datasource:       entityData.DataSource,
				FileId:           fileId,
				FileLink:         basicObject.BlankString,
				BlacklistDTTOTId: guuid.NewString(),
				Note:             basicObject.BlankString,
				ApprovalType:     basicObject.CreateApprovalType,
				ApprovedBy:       basicObject.BlankString,
				ApprovedAt:       nil,
				CreatedBy:        createdBy,
				CreatedAt:        time.Now().Format(basicObject.DateAndTime),
				UpdatedBy:        createdBy,
				UpdatedAt:        time.Now().Format(basicObject.DateAndTime),
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

func (b *blueprint) WriteRowBlacklistDTTOTFile(ctx context.Context, entity EntityBlacklistDTTOTFile) (*EntityBlacklistDTTOTFile, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteRowBlacklistDTTOTFile"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowBlacklistDTTOTFile */
	model := new(EntityBlacklistDTTOTFile)
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

func (b *blueprint) ReadRowsBlacklistDTTOTApproval(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistDTTOTApproval, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistDTTOTApproval"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistDTTOTApproval */
	model := new(EntityBlacklistDTTOTApproval)
	var entities []EntityBlacklistDTTOTApproval
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

func (b *blueprint) RemoveBlacklistDTTOTApprovalData(ctx context.Context, fileId string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.RemoveBlacklistDTTOTApprovalData"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RemoveBlacklistDTTOTApprovalData */
	model := new(EntityBlacklistDTTOTApproval)
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

func (b *blueprint) RemoveBlacklistDTTOTData(ctx context.Context, fileId string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.RemoveBlacklistDTTOTData"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RemoveBlacklistDTTOTData */
	model := new(EntityBlacklistDTTOT)
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

func (b *blueprint) ReadRowsBlacklistDTTOTApproved(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistDTTOT, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistDTTOTApproved"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistDTTOTApproved */
	model := new(EntityBlacklistDTTOT)
	var entities []EntityBlacklistDTTOT

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

func (b *blueprint) ReadRowsBlacklistDTTOTApprovedQuery(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistDTTOT, *basicObject.Meta, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistDTTOTApprovedQuery"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistDTTOTApprovedQuery */
	model := new(EntityBlacklistDTTOT)
	var entities []EntityBlacklistDTTOT

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

func (b *blueprint) ReadRowBlacklistDTTOTApprovedSingle(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistDTTOT, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowBlacklistDTTOTApprovedSingle"
	requestId, _ := ctx.Value(defaultHeaders.XRequestId).(string) // Handle missing request ID gracefully
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowBlacklistDTTOTApprovedSingle */
	model := new(EntityBlacklistDTTOT)

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

func (b *blueprint) ReadRowBlacklistDTTOTFileSingle(ctx context.Context, findBy map[string]interface{}) (*EntityBlacklistDTTOTFile, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowBlacklistDTTOTFileSingle"
	requestId, _ := ctx.Value(defaultHeaders.XRequestId).(string) // Handle missing request ID gracefully
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowBlacklistDTTOTFileSingle */
	model := new(EntityBlacklistDTTOTFile)

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

func (b *blueprint) ReadRowsBlacklistDTTOTApprovalQuery(ctx context.Context, findBy map[string]interface{}, page, limit int) (*[]EntityBlacklistDTTOTApproval, *basicObject.Meta, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistSender.ReadRowsBlacklistDTTOTApprovalQuery"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistDTTOTApprovalQuery */
	model := new(EntityBlacklistDTTOTApproval)
	var entities []EntityBlacklistDTTOTApproval

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
