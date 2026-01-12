package uscSofManagement

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/src/repositories/rprSofMan"
	"net/http"
	"time"
)

type (
	ResponseSOFCreate struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestSOFCreate struct {
		Name string `json:"name"`
	}
)

type (
	ResponseSOFUpdate struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestSOFUpdate struct {
		Id   string `param:"id"`
		Name string `json:"name"`
	}
)

type (
	ResponseSOFDelete struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestSOFDelete struct {
		Id string `param:"id"`
	}
)

type (
	ResponseSOFGetApprovedData struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
	ResponseSOFGetApproved struct {
		ResponseCode    bool                         `json:"success"`
		ResponseMessage string                       `json:"messages"`
		Data            []ResponseSOFGetApprovedData `json:"sofs"`
	}
	RequestSOFGetApproved struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}
)

func (b *blueprint) SOFCreate(ctx context.Context, request *RequestSOFCreate) (*ResponseSOFCreate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT001"
	const fName = "usecases.uscSOFManagement.SOFCreate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_SOFCreate */
	currentData, err := b.rprSofMan.ReadRowSOF(ctx, map[string]interface{}{
		"sof_name": request.Name,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentData != nil {
		responseMessage := basicObject.ResponseError{
			Message:    "source of fund already exist",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "source of fund already exist")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	currentTime := time.Now().Format(basicObject.DateAndTime)
	newEntity := new(rprSofMan.EntitySOF)
	newEntity.Id = guuid.NewString()
	newEntity.SofName = request.Name
	newEntity.ApprovedBy = email
	newEntity.ApprovedAt = currentTime
	newEntity.CreatedAt = currentTime
	newEntity.UpdatedAt = currentTime
	_, _, err = b.rprSofMan.WriteRowSOF(ctx, *newEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseSOFCreate)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Id = newEntity.Id
	return response, nil
}

func (b *blueprint) SOFUpdate(ctx context.Context, request *RequestSOFUpdate) (*ResponseSOFUpdate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT002"
	const fName = "usecases.uscSOFManagement.SOFUpdate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_SOFUpdate */
	// read current approval data
	currentApprovalEntity, err := b.rprSofMan.ReadRowSOF(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	used, err := b.rprSofMan.SOFIsUsed(ctx, currentApprovalEntity.Id)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	if !used {
		currentTime := time.Now().Format(basicObject.DateAndTime)
		existingEntity := new(rprSofMan.EntitySOF)
		existingEntity.Id = currentApprovalEntity.Id
		existingEntity.SofName = request.Name
		existingEntity.ApprovedBy = email
		existingEntity.ApprovedAt = currentTime
		existingEntity.CreatedAt = currentTime
		existingEntity.UpdatedAt = currentTime
		_, _, err = b.rprSofMan.UpdateRowSOF(ctx, *existingEntity)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}

		// response
		response := new(ResponseSOFUpdate)
		response.ResponseCode = basicObject.SuccessfullyTrue
		response.ResponseMessage = "Source of Fund has been updated"
		response.Id = existingEntity.Id
		return response, nil
	} else {
		// response
		responseMessage := basicObject.ResponseError{
			Message:    "source of fund used by user cannot Update",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "source of fund used by user cannot Update")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}
}

func (b *blueprint) SOFDelete(ctx context.Context, request *RequestSOFDelete) (*ResponseSOFDelete, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT003"
	const fName = "usecases.uscSOFManagement.SOFDelete"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_SOFDelete */
	// read current approval data
	currentSOF, err := b.rprSofMan.ReadRowSOF(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentSOF == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	used, err := b.rprSofMan.SOFIsUsed(ctx, currentSOF.Id)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	if !used {
		err = b.rprSofMan.DeleteRowSOF(ctx, currentSOF.Id)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		// response
		response := new(ResponseSOFDelete)
		response.ResponseCode = basicObject.SuccessfullyTrue
		response.ResponseMessage = "Source of Fund has been deleted"
		response.Id = currentSOF.Id
		return response, nil
	} else {
		// response
		responseMessage := basicObject.ResponseError{
			Message:    "source of fund used by user cannot Delete",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "source of fund used by user cannot Delete")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}
}

func (b *blueprint) SOFGetApproved(ctx context.Context, request *RequestSOFGetApproved) (*ResponseSOFGetApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT007"
	const fName = "usecases.uscSOFManagement.SOFGetApproved"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_SOFGetApproved */
	var dataResult []ResponseSOFGetApprovedData
	data, err := b.rprSofMan.ReadRowsSOFApproved(ctx, map[string]interface{}{}, request.Page, request.Limit, false)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseSOFGetApprovedData{
				Id:        d.Id,
				Name:      d.SofName,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.CreatedAt,
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseSOFGetApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	return response, nil
}

func (b *blueprint) SOFGetApprovedAll(ctx context.Context, request *RequestSOFGetApproved) (*ResponseSOFGetApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT007"
	const fName = "usecases.uscSOFManagement.SOFGetApprovedAll"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_SOFGetApproved */
	var dataResult []ResponseSOFGetApprovedData
	data, err := b.rprSofMan.ReadRowsSOFApproved(ctx, map[string]interface{}{}, 0, 0, true)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseSOFGetApprovedData{
				Id:        d.Id,
				Name:      d.SofName,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.CreatedAt,
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseSOFGetApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	return response, nil
}
