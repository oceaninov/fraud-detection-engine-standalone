package rprBlacklistHistory

import (
	"context"
	"fmt"
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
		ReadRowsBlacklistHistory(ctx context.Context, findBy map[string]interface{}, page, limit int, search string, start, end string) (*[]EntityBlacklistHistory, *basicObject.Meta, error)
		ReadRowsBlacklistHistoryWithoutPagination(ctx context.Context, findBy map[string]interface{}, search string, start, end string) (*[]EntityBlacklistHistory, *basicObject.Meta, error)
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
	const fName = "repositories.rprSof.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.log = log
	bp.orm = orm
	return bp
}

func (b *blueprint) ReadRowsBlacklistHistory(ctx context.Context, findBy map[string]interface{}, page, limit int, search string, start, end string) (*[]EntityBlacklistHistory, *basicObject.Meta, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.ReadRowsBlacklistHistory"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistHistory */
	model := new(EntityBlacklistHistory)
	var entities []EntityBlacklistHistory

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

	// Apply date range filter
	if start != "" && end != "" {
		// Ensure proper time format (YYYY-MM-DD)
		startTime, err1 := time.Parse("2006-01-02", start)
		endTime, err2 := time.Parse("2006-01-02", end)
		if err1 != nil || err2 != nil {
			logging.Errorw(fName, "reason", "invalid date format")
			return nil, nil, fmt.Errorf("invalid date format, use YYYY-MM-DD")
		}

		// Ensure end date includes the whole day
		endTime = endTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

		query = query.Where("created_at BETWEEN ? AND ?", startTime, endTime)
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

func (b *blueprint) ReadRowsBlacklistHistoryWithoutPagination(ctx context.Context, findBy map[string]interface{}, search string, start, end string) (*[]EntityBlacklistHistory, *basicObject.Meta, error) {
	// Constructing standard module name with request id as its prefix
	const fName = "repositories.rprBlacklistReceiver.ReadRowsBlacklistHistoryWithoutPagination"
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId)
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ReadRowsBlacklistHistoryWithoutPagination */
	model := new(EntityBlacklistHistory)
	var entities []EntityBlacklistHistory

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

	// Apply date range filter
	if start != "" && end != "" {
		// Ensure proper time format (YYYY-MM-DD)
		startTime, err1 := time.Parse("2006-01-02", start)
		endTime, err2 := time.Parse("2006-01-02", end)
		if err1 != nil || err2 != nil {
			logging.Errorw(fName, "reason", "invalid date format")
			return nil, nil, fmt.Errorf("invalid date format, use YYYY-MM-DD")
		}

		// Ensure end date includes the whole day
		endTime = endTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

		query = query.Where("created_at BETWEEN ? AND ?", startTime, endTime)
	}

	// Count total records for pagination
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Fetch the records with pagination and order by created_at descending
	if err := query.Order("created_at DESC").Find(&entities).Error; err != nil {
		logging.Errorw(fName, "reason", err.Error())
		return nil, nil, err
	}

	// Calculate pagination metadata
	meta := &basicObject.Meta{
		Count:         strconv.Itoa(int(totalCount)),
		CurrentPage:   "0",
		LastPage:      "0",
		RecordPerPage: "0",
	}

	// Returning result
	return &entities, meta, nil
}
