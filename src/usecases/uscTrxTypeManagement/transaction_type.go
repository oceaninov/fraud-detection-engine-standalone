package uscTrxTypeManagement

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/src/repositories/rprTransactionType"
	"net/http"
	"time"
)

type (
	ResponseTransactionTypeCreate struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestTransactionTypeCreate struct {
		Name string `json:"name"`
	}
)

type (
	ResponseTransactionTypeUpdate struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestTransactionTypeUpdate struct {
		Id   string `param:"id"`
		Name string `json:"name"`
	}
)

type (
	ResponseTransactionTypeDelete struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestTransactionTypeDelete struct {
		Id string `param:"id"`
	}
)

type (
	ResponseTransactionTypeGetApprovedData struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
	ResponseTransactionTypeGetApproved struct {
		ResponseCode    bool                                     `json:"success"`
		ResponseMessage string                                   `json:"messages"`
		Data            []ResponseTransactionTypeGetApprovedData `json:"transactionTypes"`
	}
	RequestTransactionTypeGetApproved struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}
)

func (b *blueprint) TransactionTypeCreate(ctx context.Context, request *RequestTransactionTypeCreate) (*ResponseTransactionTypeCreate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT001"
	const fName = "usecases.uscTransactionTypeManagement.TransactionTypeCreate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_TransactionTypeCreate */
	currentData, err := b.rprTransactionType.ReadRowTransactionType(ctx, map[string]interface{}{
		"name": request.Name,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentData != nil {
		responseMessage := basicObject.ResponseError{
			Message:    "transaction type already exist",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "transaction type already exist")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	currentTime := time.Now().Format(basicObject.DateAndTime)
	newEntity := new(rprTransactionType.EntityTransactionType)
	newEntity.Id = guuid.NewString()
	newEntity.Name = request.Name
	newEntity.ApprovedBy = email
	newEntity.ApprovedAt = &currentTime
	newEntity.CreatedAt = currentTime
	newEntity.UpdatedAt = currentTime
	_, _, err = b.rprTransactionType.WriteRowTransactionType(ctx, *newEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseTransactionTypeCreate)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Id = newEntity.Id
	return response, nil
}

func (b *blueprint) TransactionTypeUpdate(ctx context.Context, request *RequestTransactionTypeUpdate) (*ResponseTransactionTypeUpdate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT002"
	const fName = "usecases.uscTransactionTypeManagement.TransactionTypeUpdate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_TransactionTypeUpdate */
	// read current approval data
	currentApprovalEntity, err := b.rprTransactionType.ReadRowTransactionType(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	used, err := b.rprTransactionType.TransactionTypeIsUsed(ctx, currentApprovalEntity.Name)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	if !used {
		currentTime := time.Now().Format(basicObject.DateAndTime)
		existingEntity := new(rprTransactionType.EntityTransactionType)
		existingEntity.Id = currentApprovalEntity.Id
		existingEntity.Name = request.Name
		existingEntity.ApprovedBy = email
		existingEntity.ApprovedAt = &currentTime
		existingEntity.CreatedAt = currentTime
		existingEntity.UpdatedAt = currentTime
		_, _, err = b.rprTransactionType.UpdateRowTransactionType(ctx, *existingEntity)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}

		// response
		response := new(ResponseTransactionTypeUpdate)
		response.ResponseCode = basicObject.SuccessfullyTrue
		response.ResponseMessage = "TransactionType has been updated"
		response.Id = existingEntity.Id
		return response, nil
	} else {
		// response
		responseMessage := basicObject.ResponseError{
			Message:    "transaction type used by user cannot Update",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "transaction type used by user cannot Update")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}
}

func (b *blueprint) TransactionTypeDelete(ctx context.Context, request *RequestTransactionTypeDelete) (*ResponseTransactionTypeDelete, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT003"
	const fName = "usecases.uscTransactionTypeManagement.TransactionTypeDelete"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_TransactionTypeDelete */
	// read current approval data
	currentTransactionType, err := b.rprTransactionType.ReadRowTransactionType(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentTransactionType == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	used, err := b.rprTransactionType.TransactionTypeIsUsed(ctx, currentTransactionType.Name)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	if !used {
		err = b.rprTransactionType.DeleteRowTransactionType(ctx, currentTransactionType.Id)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		// response
		response := new(ResponseTransactionTypeDelete)
		response.ResponseCode = basicObject.SuccessfullyTrue
		response.ResponseMessage = "TransactionType has been deleted"
		response.Id = currentTransactionType.Id
		return response, nil
	} else {
		// response
		responseMessage := basicObject.ResponseError{
			Message:    "transaction type used by user cannot Delete",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "transaction type used by user cannot Delete")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}
}

func (b *blueprint) TransactionTypeGetApproved(ctx context.Context, request *RequestTransactionTypeGetApproved) (*ResponseTransactionTypeGetApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT007"
	const fName = "usecases.uscTransactionTypeManagement.TransactionTypeGetApproved"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_TransactionTypeGetApproved */
	var dataResult []ResponseTransactionTypeGetApprovedData
	data, err := b.rprTransactionType.ReadRowsTransactionTypeApproved(ctx, map[string]interface{}{}, request.Page, request.Limit, false)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseTransactionTypeGetApprovedData{
				Id:        d.Id,
				Name:      d.Name,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseTransactionTypeGetApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	return response, nil
}

func (b *blueprint) TransactionTypeGetApprovedAll(ctx context.Context, request *RequestTransactionTypeGetApproved) (*ResponseTransactionTypeGetApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT007"
	const fName = "usecases.uscTransactionTypeManagement.TransactionTypeGetApprovedAll"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_TransactionTypeGetApproved */
	var dataResult []ResponseTransactionTypeGetApprovedData
	data, err := b.rprTransactionType.ReadRowsTransactionTypeApproved(ctx, map[string]interface{}{}, 0, 0, true)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseTransactionTypeGetApprovedData{
				Id:        d.Id,
				Name:      d.Name,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseTransactionTypeGetApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	return response, nil
}
