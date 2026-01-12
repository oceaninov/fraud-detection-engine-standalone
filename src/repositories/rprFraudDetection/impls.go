package rprFraudDetection

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint repository function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */
		ReadRowsBlacklistDTTOT(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistDTTOT, error)
		ReadRowsBlacklistMerchant(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistMerchant, error)
		ReadRowsBlacklistSender(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistSender, error)
		ReadRowsBlacklistReceiver(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistReceiver, error)
		ReadRowsRules(ctx context.Context, findBy map[string]interface{}) (*[]EntityRule, error)
		ReadRowsTransactionsByPhoneAndDateRange(ctx context.Context, phone string, s, e time.Time) (*[]EntityTransaction, error)
		ReadRowsTransactionsBySenderAndReceiverId(ctx context.Context, senderId, receiverId string) (*[]EntityTransaction, error)
		ReadRowsBlacklistDTTOTByPerformerName(ctx context.Context, performerName string) (*[]EntityBlacklistDTTOT, error)
		ReadRowsBlacklistDTTOTByPerformerNameV2(ctx context.Context, performerName string) (*[]EntityBlacklistDTTOT, error)
		ReadRowsBlacklistSenderByPerformerName(ctx context.Context, performerName, userId, transactionType string) (*[]EntityBlacklistSender, error)
		ReadRowsBlacklistReceiverByPerformerName(ctx context.Context, performerName, userId, transactionType string) (*[]EntityBlacklistReceiver, error)
		ReadRowsBlacklistMerchantByPerformerNameAndUserId(ctx context.Context, performerName, userId string) (*[]EntityBlacklistMerchant, error)
		ReadRowsKeywordWithWords(ctx context.Context, words map[string]map[string]bool) (*[]EntityKeyword, *[]EntityKeyword, error)
		ReadRowsKeyword(ctx context.Context) (*[]EntityKeyword, error)

		WriteRowLog(ctx context.Context, entity EntityLog) (*EntityLog, *string, error)
		WriteRowTransaction(ctx context.Context, entity EntityTransaction) (*EntityTransaction, *string, error)
		WriteRowBlacklistHistory(ctx context.Context, entity EntityBlacklistHistory) (*EntityBlacklistHistory, *string, error)
		WriteRowBlacklistSender(ctx context.Context, entity EntityBlacklistSender) (*EntityBlacklistSender, *string, error)

		UpdateRowTransaction(ctx context.Context, entity EntityTransaction) (*EntityTransaction, *string, error)
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

func (b *blueprint) ReadRowsBlacklistDTTOT(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistDTTOT, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistDTTOT"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistDTTOT */
	model := new(EntityBlacklistDTTOT)
	var entities []EntityBlacklistDTTOT
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}
	if err := query.Limit(10).Offset(1).Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsBlacklistMerchant(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistMerchant, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistMerchant"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistMerchant */
	model := new(EntityBlacklistMerchant)
	var entities []EntityBlacklistMerchant
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}
	if err := query.Limit(10).Offset(1).Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsBlacklistSender(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistSender, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistSender"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistSender */
	model := new(EntityBlacklistSender)
	var entities []EntityBlacklistSender
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}
	if err := query.Limit(10).Offset(1).Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsBlacklistReceiver(ctx context.Context, findBy map[string]interface{}) (*[]EntityBlacklistReceiver, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistReceiver"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistReceiver */
	model := new(EntityBlacklistReceiver)
	var entities []EntityBlacklistReceiver
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if len(findBy) > 0 {
		query = query.Where(findBy)
	}
	if err := query.Limit(10).Offset(1).Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsRules(ctx context.Context, findBy map[string]interface{}) (*[]EntityRule, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsRules"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsRules */
	model := new(EntityRule)
	var entities []EntityRule
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if len(findBy) > 0 {
		sofs, sofsOk := findBy["sofs"]
		types, typesOk := findBy["type"]
		if sofsOk && typesOk {
			sofsData := fmt.Sprintf("%%%s%%", sofs)
			query = query.Where("rules.sofs LIKE ? AND rules.type = ? AND rules.status = 1", sofsData, types)
		}
	}
	if err := query.Order("created_at DESC").Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}
	logging.Infow(fName, "length_rules", len(entities))

	// returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsTransactionsByPhoneAndDateRange(ctx context.Context, phone string, s, e time.Time) (*[]EntityTransaction, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsTransactionsByPhoneAndDateRange"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsTransactionsByPhoneAndDateRange */
	model := new(EntityTransaction)
	var entities []EntityTransaction
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	query = query.Where("transaction.user_id = ? AND transaction.created_at BETWEEN ? AND ?", phone, s, e)
	if err := query.Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsTransactionsBySenderAndReceiverId(ctx context.Context, senderId, receiverId string) (*[]EntityTransaction, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsTransactionsBySenderAndReceiverId"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsTransactionsBySenderAndReceiverId */
	model := new(EntityTransaction)
	var entities []EntityTransaction
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	query = query.Where("transaction.user_id = ? and transaction.destination_id = ?", senderId, receiverId)
	if err := query.Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsBlacklistDTTOTByPerformerName(ctx context.Context, performerName string) (*[]EntityBlacklistDTTOT, error) {
	// Constructing standard module name with request ID as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistDTTOTByPerformerName"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistDTTOTByPerformerName */
	model := new(EntityBlacklistDTTOT)
	var entities []EntityBlacklistDTTOT
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	query = query.Where("UPPER(black_list_dttot.name) LIKE ?", "%"+strings.ToUpper(performerName)+"%")

	if err := query.Find(&entities).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// Returning result
	return &entities, nil
}

func (b *blueprint) ReadRowsBlacklistDTTOTByPerformerNameV2(ctx context.Context, performerName string) (*[]EntityBlacklistDTTOT, error) {
	// Constructing standard module name with request ID as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistDTTOTByPerformerNameV2"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistDTTOTByPerformerNameV2 */
	model := new(EntityBlacklistDTTOT)
	var entities []EntityBlacklistDTTOT
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())

	if err := query.Find(&entities).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	var newScannedNameDTTO []EntityBlacklistDTTOT
	seen := make(map[string]bool)

	for _, entity := range entities {
		nameSplit := strings.Split(entity.Name, " ")
		performerNameSplit := strings.Split(performerName, " ")

		for _, existingName := range nameSplit {
			for _, performerPart := range performerNameSplit {
				if strings.Contains(performerPart, existingName) {
					if !seen[entity.Name] {
						newScannedNameDTTO = append(newScannedNameDTTO, entity)
						seen[entity.Name] = true
					}
					break
				}
			}
		}
	}

	// Returning result
	return &newScannedNameDTTO, nil
}

func (b *blueprint) ReadRowsBlacklistSenderByPerformerName(ctx context.Context, performerName, userId, transactionType string) (*[]EntityBlacklistSender, error) {
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistSenderByPerformerName"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	model := new(EntityBlacklistSender)
	var entities []EntityBlacklistSender
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())

	const q = `
		black_list_sender.phone_number = ?
		AND black_list_sender.status = 1
		AND (
			black_list_sender.transaction_types = ?
			OR black_list_sender.transaction_types LIKE ?
			OR black_list_sender.transaction_types LIKE ?
			OR black_list_sender.transaction_types LIKE ?
		)
	`
	query = query.Where(q,
		userId,
		transactionType,
		transactionType+",%",
		"%,"+transactionType+",%",
		"%,"+transactionType,
	)

	if err := query.Find(&entities).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	return &entities, nil
}

func (b *blueprint) ReadRowsBlacklistReceiverByPerformerName(ctx context.Context, performerName, userId, transactionType string) (*[]EntityBlacklistReceiver, error) {
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistReceiverByPerformerName"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	model := new(EntityBlacklistReceiver)
	var entities []EntityBlacklistReceiver
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())

	const q = `
		black_list_receiver.phone_number = ?
		AND black_list_receiver.status = 1
		AND (
			black_list_receiver.transaction_types = ?
			OR black_list_receiver.transaction_types LIKE ?
			OR black_list_receiver.transaction_types LIKE ?
			OR black_list_receiver.transaction_types LIKE ?
		)
	`
	query = query.Where(q,
		userId,
		transactionType,
		transactionType+",%",
		"%,"+transactionType+",%",
		"%,"+transactionType,
	)

	if err := query.Find(&entities).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	return &entities, nil
}

func (b *blueprint) ReadRowsBlacklistMerchantByPerformerNameAndUserId(ctx context.Context, performerName, userId string) (*[]EntityBlacklistMerchant, error) {
	// Constructing standard module name with request ID as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsBlacklistMerchantByPerformerNameAndNMID"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistMerchantByPerformerNameAndNMID */
	model := new(EntityBlacklistMerchant)
	var entities []EntityBlacklistMerchant
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	const q = "UPPER(black_list_merchant.merchant_name) LIKE ? and black_list_merchant.nmid = ?"
	query = query.Where(q, "%"+strings.ToUpper(performerName)+"%", userId)

	if err := query.Find(&entities).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		logging.Errorw(fName, "reason", err.Error())
		return nil, err
	}

	// Returning result
	return &entities, nil
}

func (b *blueprint) WriteRowLog(ctx context.Context, entity EntityLog) (*EntityLog, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteRowLog"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowLog */
	model := new(EntityLog)
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

func (b *blueprint) WriteRowTransaction(ctx context.Context, entity EntityTransaction) (*EntityTransaction, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteRowTransaction"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowTransaction */
	model := new(EntityTransaction)
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

func (b *blueprint) WriteRowBlacklistHistory(ctx context.Context, entity EntityBlacklistHistory) (*EntityBlacklistHistory, *string, error) {
	const fName = "repositories.rprFraudDetection.WriteRowBlacklistHistory"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	model := new(EntityBlacklistHistory)
	// Create new record only if not exists
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

	return &entity, &entity.Id, nil
}

func (b *blueprint) WriteRowBlacklistSender(ctx context.Context, entity EntityBlacklistSender) (*EntityBlacklistSender, *string, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.WriteRowBlacklistSender"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_WriteRowBlacklistSender */
	// Check if the record already exists (you can change the condition based on your logic)
	var existing EntityBlacklistSender
	err := b.orm.Table(existing.TableName()).
		Where("phone_number = ? AND beneficiary_name = ?",
			entity.PhoneNumber, entity.BeneficiaryName).
		First(&existing).Error
	if err == nil {
		logging.Infow(fName, "reason", "record already exists, skipping insert")
		return &existing, &existing.Id, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logging.Errorw(fName, "reason", "failed to query existing record", "error", err.Error())
		return nil, nil, err
	}

	model := new(EntityBlacklistSender)
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

func (b *blueprint) ReadRowsKeywordWithWords(ctx context.Context, words map[string]map[string]bool) (*[]EntityKeyword, *[]EntityKeyword, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsKeywordWithWords"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsKeywordWithWords */
	model := new(EntityKeyword)
	var entitiesMatched []EntityKeyword
	var entitiesLiked []EntityKeyword

	for word, detectionRule := range words {
		isLike := detectionRule["like"]
		isMatch := detectionRule["match"]
		if isLike {
			var entities []EntityKeyword
			const isLikeQuery = "UPPER(keyword.keyword) LIKE ?"
			query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
			query = query.Where(isLikeQuery, "%"+strings.ToUpper(word)+"%")
			if err := query.Find(&entities).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					logging.Errorw(fName, "reason", err.Error())
					return nil, nil, err
				}
			}
			entitiesLiked = append(entitiesLiked, entities...)
		}
		if isMatch {
			var entities []EntityKeyword
			const isMatchQuery = "UPPER(keyword.keyword) = ?"
			query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
			query = query.Where(isMatchQuery, strings.ToUpper(word)) // Note: Removed '%' for exact match
			if err := query.Find(&entities).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					logging.Errorw(fName, "reason", err.Error())
					return nil, nil, err
				}
			}
			entitiesMatched = append(entitiesMatched, entities...)
		}
	}

	// Returning result
	return &entitiesMatched, &entitiesLiked, nil
}

func (b *blueprint) ReadRowsKeyword(ctx context.Context) (*[]EntityKeyword, error) {
	// constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.ReadRowsKeyword"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsKeyword */
	model := new(EntityKeyword)
	var entities []EntityKeyword
	query := b.orm.Table(model.TableName()).Select(model.DefaultColumns())
	if err := query.Find(&entities).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logging.Errorw(fName, "reason", err.Error())
			return nil, err
		}
	}
	// Returning result
	return &entities, nil
}

func (b *blueprint) UpdateRowTransaction(ctx context.Context, entity EntityTransaction) (*EntityTransaction, *string, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprFraudDetection.UpdateRowTransaction"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UpdateRowTransaction */
	model := new(EntityTransaction)
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
