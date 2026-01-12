package rprSofMan

import (
	"context"
	"errors"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint repository function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		SOFIsUsed(ctx context.Context, transactionTypeName string) (bool, error)
		WriteRowSOF(ctx context.Context, entity EntitySOF) (*EntitySOF, *string, error)
		DeleteRowSOF(ctx context.Context, id string) error
		UpdateRowSOF(ctx context.Context, entity EntitySOF) (*EntitySOF, *string, error)

		ReadRowSOF(ctx context.Context, findBy map[string]interface{}) (*EntitySOF, error)
		ReadRowsSOFApproved(ctx context.Context, findBy map[string]interface{}, page, limit int, allData bool) (*[]EntitySOF, error)
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
	const fName = "repositories.rprSOF.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.log = log
	bp.orm = orm
	return bp
}

func (b *blueprint) SOFIsUsed(ctx context.Context, transactionTypeName string) (bool, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprSOF.SOFIsUsed"
	requestId, _ := ctx.Value(defaultHeaders.XRequestId).(string) // Handle missing request ID gracefully
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_SOFIsUsed */

	// Define the model and check if transactionTypeName exists
	var count int64
	err := b.orm.Table("rules").Where("sofs LIKE ?", "%"+transactionTypeName+"%").Count(&count).Error
	if err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return false, err
	}

	// If count > 0, it means the transactionTypeName is in use
	return count > 0, nil
}

func (b *blueprint) WriteRowSOF(ctx context.Context, entity EntitySOF) (*EntitySOF, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprSOF.WriteRowSOF"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowSOF */
	model := new(EntitySOF)
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

func (b *blueprint) UpdateRowSOF(ctx context.Context, entity EntitySOF) (*EntitySOF, *string, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprSOF.UpdateRowSOF"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UpdateRowSOF */
	model := new(EntitySOF)
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

func (b *blueprint) DeleteRowSOF(ctx context.Context, id string) error {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprSOF.DeleteRowSOF"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_DeleteRowSOF */
	model := new(EntitySOF)
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

func (b *blueprint) ReadRowsSOFApproved(ctx context.Context, findBy map[string]interface{}, page, limit int, allData bool) (*[]EntitySOF, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprSOF.ReadRowsSOFApproved"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsSOFApproved */
	model := new(EntitySOF)
	var entities []EntitySOF

	if allData {
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
	} else {
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
		if len(findBy) > 0 {
			query = query.Where(findBy)
		}
		if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&entities).Error; err != nil {
			logging.Errorw(fName, "reason", err.Error())
			return nil, err
		}

		// returning result
		return &entities, nil
	}
}

func (b *blueprint) ReadRowSOF(ctx context.Context, findBy map[string]interface{}) (*EntitySOF, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprSOF.ReadRowSOF"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowSOF */
	model := new(EntitySOF)
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
