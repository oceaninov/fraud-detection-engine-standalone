package uscChannelManagement

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/src/repositories/rprChannel"
	"net/http"
	"strings"
	"time"
)

type (
	ResponseChannelCreate struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestChannelCreate struct {
		Name string `json:"name"`
	}
)

type (
	ResponseChannelUpdate struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestChannelUpdate struct {
		Id   string `param:"id"`
		Name string `json:"name"`
	}
)

type (
	ResponseChannelDelete struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestChannelDelete struct {
		Id string `param:"id"`
	}
)

type (
	ResponseChannelGetApprovedData struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
	ResponseChannelGetApproved struct {
		ResponseCode    bool                             `json:"success"`
		ResponseMessage string                           `json:"messages"`
		Data            []ResponseChannelGetApprovedData `json:"channel"`
	}
	RequestChannelGetApproved struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}
)

func (b *blueprint) ChannelCreate(ctx context.Context, request *RequestChannelCreate) (*ResponseChannelCreate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT001"
	const fName = "usecases.uscChannelManagement.ChannelCreate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ChannelCreate */
	currentData, err := b.rprChannel.ReadRowChannel(ctx, map[string]interface{}{
		"channel_name": strings.ToLower(request.Name),
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
	newEntity := new(rprChannel.EntityChannel)
	newEntity.Id = guuid.NewString()
	newEntity.ChannelName = strings.ToLower(request.Name)
	newEntity.ApprovedBy = email
	newEntity.ApprovedAt = currentTime
	newEntity.CreatedAt = currentTime
	newEntity.UpdatedAt = currentTime
	_, _, err = b.rprChannel.WriteRowChannel(ctx, *newEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseChannelCreate)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Id = newEntity.Id
	return response, nil
}

func (b *blueprint) ChannelUpdate(ctx context.Context, request *RequestChannelUpdate) (*ResponseChannelUpdate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT002"
	const fName = "usecases.uscChannelManagement.ChannelUpdate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ChannelUpdate */
	// read current approval data
	currentApprovalEntity, err := b.rprChannel.ReadRowChannel(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	used, err := b.rprChannel.ChannelIsUsed(ctx, currentApprovalEntity.ChannelName)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	if !used {
		currentTime := time.Now().Format(basicObject.DateAndTime)
		existingEntity := new(rprChannel.EntityChannel)
		existingEntity.Id = currentApprovalEntity.Id
		existingEntity.ChannelName = strings.ToLower(request.Name)
		existingEntity.ApprovedBy = email
		existingEntity.ApprovedAt = currentTime
		existingEntity.CreatedAt = currentTime
		existingEntity.UpdatedAt = currentTime
		_, _, err = b.rprChannel.UpdateRowChannel(ctx, *existingEntity)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}

		// response
		response := new(ResponseChannelUpdate)
		response.ResponseCode = basicObject.SuccessfullyTrue
		response.ResponseMessage = "Channel has been updated"
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

func (b *blueprint) ChannelDelete(ctx context.Context, request *RequestChannelDelete) (*ResponseChannelDelete, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT003"
	const fName = "usecases.uscChannelManagement.ChannelDelete"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ChannelDelete */
	// read current approval data
	currentChannel, err := b.rprChannel.ReadRowChannel(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentChannel == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	used, err := b.rprChannel.ChannelIsUsed(ctx, currentChannel.ChannelName)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	if !used {
		err = b.rprChannel.DeleteRowChannel(ctx, currentChannel.Id)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		// response
		response := new(ResponseChannelDelete)
		response.ResponseCode = basicObject.SuccessfullyTrue
		response.ResponseMessage = "Channel has been deleted"
		response.Id = currentChannel.Id
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

func (b *blueprint) ChannelGetApproved(ctx context.Context, request *RequestChannelGetApproved) (*ResponseChannelGetApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT007"
	const fName = "usecases.uscChannelManagement.ChannelGetApproved"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ChannelGetApproved */
	var dataResult []ResponseChannelGetApprovedData
	data, err := b.rprChannel.ReadRowsChannelApproved(ctx, map[string]interface{}{}, request.Page, request.Limit, false)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseChannelGetApprovedData{
				Id:        d.Id,
				Name:      d.ChannelName,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseChannelGetApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	return response, nil
}

func (b *blueprint) ChannelGetApprovedAll(ctx context.Context, request *RequestChannelGetApproved) (*ResponseChannelGetApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT007"
	const fName = "usecases.uscChannelManagement.ChannelGetApprovedAll"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ChannelGetApproved */
	var dataResult []ResponseChannelGetApprovedData
	data, err := b.rprChannel.ReadRowsChannelApproved(ctx, map[string]interface{}{}, 0, 0, true)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseChannelGetApprovedData{
				Id:        d.Id,
				Name:      d.ChannelName,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseChannelGetApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	return response, nil
}
