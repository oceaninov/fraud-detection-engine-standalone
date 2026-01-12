package uscKeywordManagement

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/src/repositories/rprKeyword"
	"net/http"
	"strconv"
	"time"
)

type (
	ResponseKeywordCreate struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestKeywordCreate struct {
		//Channel string `json:"channel"`
		Keyword string `json:"keyword"`
		Action  string `json:"action"`
	}
)

type (
	ResponseKeywordUpdate struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestKeywordUpdate struct {
		Id string `json:"id"`
		//Channel string `json:"channel"`
		Keyword string `json:"keyword"`
		Action  string `json:"action"`
	}
)

type (
	ResponseKeywordDelete struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestKeywordDelete struct {
		Id string `param:"id"`
	}
)

type (
	ResponseKeywordApprove struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestKeywordApprove struct {
		Id          string `param:"id"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"actionTaker"`
	}
)

type (
	ResponseKeywordReject struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestKeywordReject struct {
		Id          string `param:"id"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"actionTaker"`
	}
)

type (
	ResponseKeywordGetApprovalData struct {
		Id           string `json:"id"`
		KeywordId    string `json:"keywordId"`
		Keyword      string `json:"keyword"`
		Action       string `json:"action"`
		Channel      string `json:"channel"`
		ApprovalType string `json:"approvalType"`
		Note         string `json:"note"`
		Status       string `json:"status"`
		CreatedBy    string `json:"createdBy"`
		CreatedAt    string `json:"createdAt"`
		UpdatedBy    string `json:"updatedBy"`
		UpdatedAt    string `json:"updatedAt"`
		ApprovedBy   string `json:"approvedBy"`
		ApprovedAt   string `json:"approvedAt"`
	}
	ResponseKeywordGetApproval struct {
		ResponseCode    bool                             `json:"success"`
		ResponseMessage string                           `json:"messages"`
		Data            []ResponseKeywordGetApprovalData `json:"keyword"`
		Meta            basicObject.Meta                 `json:"meta"`
	}
	RequestKeywordGetApproval struct {
		StatusApproval int    `query:"approval_content"`
		Page           int    `query:"page"`
		Limit          int    `query:"limit"`
		Search         string `query:"search"`
	}
)

type (
	ResponseKeywordGetApprovalDetail struct {
		Id           string `json:"id"`
		Keyword      string `json:"keyword"`
		Action       string `json:"action"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		CreatedBy    string `json:"createdBy"`
		ApprovalType string `json:"approvalType"`
		ApprovedAt   string `json:"approvedAt"`
		ApprovalBy   string `json:"approvalBy"`
		RejectNote   string `json:"rejectNote"`
		Status       string `json:"status"`
	}
	RequestKeywordGetApprovalDetail struct {
		ApprovalContent bool   `query:"approval_content"`
		Id              string `param:"id"`
	}
)

type (
	ResponseKeywordGetApprovedData struct {
		Id         string `json:"id"`
		Keyword    string `json:"keyword"`
		Action     string `json:"action"`
		Channel    string `json:"channel"`
		CreatedBy  string `json:"createdBy"`
		CreatedAt  string `json:"createdAt"`
		UpdatedBy  string `json:"updatedBy"`
		UpdatedAt  string `json:"updatedAt"`
		ApprovedBy string `json:"approvedBy"`
		ApprovedAt string `json:"approvedAt"`
	}
	ResponseKeywordGetApproved struct {
		ResponseCode    bool                             `json:"success"`
		ResponseMessage string                           `json:"messages"`
		Data            []ResponseKeywordGetApprovedData `json:"keyword"`
		Meta            basicObject.Meta                 `json:"meta"`
	}
	RequestKeywordGetApproved struct {
		Page   int    `query:"page"`
		Limit  int    `query:"limit"`
		Search string `query:"search"`
	}
)

func (b *blueprint) KeywordCreate(ctx context.Context, request *RequestKeywordCreate) (*ResponseKeywordCreate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT001"
	const fName = "usecases.uscKeywordManagement.KeywordCreate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_KeywordCreate */
	// read current approval data
	currentApprovalEntity, err := b.rprKeyword.ReadRowKeywordApproval(ctx, map[string]interface{}{
		"keyword":       request.Keyword,
		"approval_type": basicObject.CreateApprovalType,
		"status":        1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity != nil {
		responseMessage := basicObject.ResponseError{
			Message:    "keyword already exist",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "keyword already exist")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	dataEntity := new(rprKeyword.EntityKeywordApproval)
	dataEntity.Id = guuid.NewString()
	dataEntity.KeywordId = guuid.NewString()
	dataEntity.Keyword = request.Keyword
	dataEntity.Action = request.Action
	//dataEntity.Channel = request.Channel
	dataEntity.CreatedBy = email
	dataEntity.UpdatedBy = email
	dataEntity.ApprovalType = basicObject.CreateApprovalType
	dataEntity.ApprovedBy = basicObject.BlankString
	dataEntity.ApprovedAt = nil
	dataEntity.Note = basicObject.BlankString
	dataEntity.Status = basicObject.ApprovalPending
	dataEntity.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	entityRuleApproval, entityRuleId, err := b.rprKeyword.WriteRowKeywordApproval(ctx, *dataEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if entityRuleId == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if entityRuleApproval == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseKeywordCreate)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success create approval for create a keyword"
	response.Id = dataEntity.Id
	return response, nil
}

func (b *blueprint) KeywordUpdate(ctx context.Context, request *RequestKeywordUpdate) (*ResponseKeywordUpdate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT002"
	const fName = "usecases.uscKeywordManagement.KeywordUpdate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_KeywordUpdate */
	// read current approval data
	currentApprovalEntity, err := b.rprKeyword.ReadRowKeywordApproval(ctx, map[string]interface{}{
		"keyword_id":    request.Id,
		"approval_type": basicObject.UpdateApprovalType,
		"status":        1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity != nil {
		errMsg := fmt.Errorf("already exist on approval data")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	currentKeyword, err := b.rprKeyword.ReadRowKeyword(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentKeyword == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	dataEntity := new(rprKeyword.EntityKeywordApproval)
	dataEntity.Id = guuid.NewString()
	dataEntity.KeywordId = currentKeyword.Id
	dataEntity.Keyword = request.Keyword
	dataEntity.Action = request.Action
	//dataEntity.Channel = request.Channel
	dataEntity.ApprovalType = basicObject.UpdateApprovalType
	dataEntity.Note = basicObject.BlankString
	dataEntity.Status = basicObject.ApprovalPending
	dataEntity.ApprovedBy = basicObject.BlankString
	dataEntity.ApprovedAt = nil
	dataEntity.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.CreatedBy = email
	dataEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.UpdatedBy = email
	entityRuleApproval, entityRuleId, err := b.rprKeyword.WriteRowKeywordApproval(ctx, *dataEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if entityRuleId == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if entityRuleApproval == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseKeywordUpdate)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Id = dataEntity.Id
	return response, nil
}

func (b *blueprint) KeywordDelete(ctx context.Context, request *RequestKeywordDelete) (*ResponseKeywordDelete, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT003"
	const fName = "usecases.uscKeywordManagement.KeywordDelete"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_KeywordDelete */
	// read current approval data
	currentApprovalEntity, err := b.rprKeyword.ReadRowKeywordApproval(ctx, map[string]interface{}{
		"keyword_id":    request.Id,
		"approval_type": basicObject.DeleteApprovalType,
		"status":        1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity != nil {
		responseMessage := basicObject.ResponseError{
			Message:    "already submitted",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "already submitted")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	currentKeyword, err := b.rprKeyword.ReadRowKeyword(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentKeyword == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	dataEntity := new(rprKeyword.EntityKeywordApproval)
	dataEntity.Id = guuid.NewString()
	dataEntity.KeywordId = currentKeyword.Id
	dataEntity.Keyword = currentKeyword.Keyword
	dataEntity.Action = currentKeyword.Action
	dataEntity.Channel = currentKeyword.Channel
	dataEntity.CreatedBy = email
	dataEntity.UpdatedBy = email
	dataEntity.ApprovalType = basicObject.DeleteApprovalType
	dataEntity.ApprovedBy = basicObject.BlankString
	dataEntity.ApprovedAt = nil
	dataEntity.Note = basicObject.BlankString
	dataEntity.Status = basicObject.ApprovalPending
	dataEntity.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	entityRuleApproval, entityRuleId, err := b.rprKeyword.WriteRowKeywordApproval(ctx, *dataEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if entityRuleId == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if entityRuleApproval == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseKeywordDelete)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success delete approval for create a keyword"
	response.Id = dataEntity.Id
	return response, nil
}

func (b *blueprint) KeywordApprove(ctx context.Context, request *RequestKeywordApprove) (*ResponseKeywordApprove, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT004"
	const fName = "usecases.uscKeywordManagement.KeywordApprove"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_KeywordApprove */
	// read current approval data
	currentApprovalEntity, err := b.rprKeyword.ReadRowKeywordApproval(ctx, map[string]interface{}{
		"id":     request.Id,
		"status": 1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity == nil {
		errMsg := fmt.Errorf("data not found on approval data")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	switch currentApprovalEntity.ApprovalType {
	case basicObject.CreateApprovalType:
		newEntity := new(rprKeyword.EntityKeyword)
		newEntity.Id = currentApprovalEntity.KeywordId
		newEntity.Keyword = currentApprovalEntity.Keyword
		newEntity.Action = currentApprovalEntity.Action
		newEntity.Channel = currentApprovalEntity.Channel
		newEntity.ApprovedAt = currentApprovalEntity.ApprovedAt
		newEntity.CreatedBy = currentApprovalEntity.CreatedBy
		newEntity.UpdatedBy = currentApprovalEntity.UpdatedBy
		newEntity.CreatedAt = time.Now().Format(basicObject.DateAndTime)
		newEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
		newEntity.ApprovedBy = email
		approvedAt := time.Now().Format(basicObject.DateAndTime)
		newEntity.ApprovedAt = &approvedAt
		_, _, err := b.rprKeyword.WriteRowKeyword(ctx, *newEntity)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	case basicObject.UpdateApprovalType:
		existingEntity := new(rprKeyword.EntityKeyword)
		existingEntity.Id = currentApprovalEntity.KeywordId
		existingEntity.Keyword = currentApprovalEntity.Keyword
		existingEntity.Action = currentApprovalEntity.Action
		existingEntity.Channel = currentApprovalEntity.Channel
		existingEntity.CreatedBy = email
		existingEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
		existingEntity.ApprovedBy = email
		existingEntity.UpdatedBy = email
		approvedAt := time.Now().Format(basicObject.DateAndTime)
		existingEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
		existingEntity.ApprovedAt = &approvedAt
		_, _, err := b.rprKeyword.UpdateRowKeyword(ctx, *existingEntity)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	case basicObject.DeleteApprovalType:
		err := b.rprKeyword.DeleteRowKeyword(ctx, currentApprovalEntity.KeywordId)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	}

	err = b.rprKeyword.WriteApprovalForKeywordApproval(ctx, request.Id, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseKeywordApprove)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success approve keyword data"
	return response, nil
}

func (b *blueprint) KeywordReject(ctx context.Context, request *RequestKeywordReject) (*ResponseKeywordReject, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT005"
	const fName = "usecases.uscKeywordManagement.KeywordReject"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_KeywordReject */
	// read current approval data
	currentApprovalEntity, err := b.rprKeyword.ReadRowKeywordApproval(ctx, map[string]interface{}{
		"id":     request.Id,
		"status": 1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity == nil {
		errMsg := fmt.Errorf("data not found on approval data")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	err = b.rprKeyword.WriteRejectForKeywordApproval(ctx, request.Id, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseKeywordReject)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success reject keyword data"
	return response, nil
}

func (b *blueprint) KeywordGetApproval(ctx context.Context, request *RequestKeywordGetApproval) (*ResponseKeywordGetApproval, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT006"
	const fName = "usecases.uscKeywordManagement.KeywordGetApproval"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_KeywordGetApproval */
	dataResult := []ResponseKeywordGetApprovalData{}
	queryData := map[string]interface{}{}
	if request.StatusApproval != 0 {
		queryData["status"] = request.StatusApproval
	}
	search := ""
	if request.Search != "" && request.Search != "230" {
		search = request.Search
	}
	data, meta, err := b.rprKeyword.ReadRowsKeywordApproval(ctx, queryData, request.Page, request.Limit, search)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseKeywordGetApprovalData{
				Id:           d.Id,
				KeywordId:    d.KeywordId,
				Keyword:      d.Keyword,
				Action:       d.Action,
				Channel:      d.Channel,
				ApprovalType: d.ApprovalType,
				Note:         d.Note,
				Status:       d.Status,
				ApprovedBy:   d.ApprovedBy,
				CreatedBy:    d.CreatedBy,
				CreatedAt:    d.CreatedAt,
				UpdatedBy:    d.UpdatedBy,
				UpdatedAt:    d.UpdatedAt,
			}
			status, _ := strconv.Atoi(d.Status)
			result.Status = strconv.Itoa(status - 1)
			if d.ApprovedAt != nil {
				result.ApprovedAt = *d.ApprovedAt
			} else {
				result.ApprovedAt = basicObject.BlankString
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseKeywordGetApproval)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) KeywordGetApproved(ctx context.Context, request *RequestKeywordGetApproved) (*ResponseKeywordGetApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT007"
	const fName = "usecases.uscKeywordManagement.KeywordGetApproved"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_KeywordGetApproved */
	dataResult := []ResponseKeywordGetApprovedData{}
	search := ""
	if request.Search != "" && request.Search != "230" {
		search = request.Search
	}
	data, meta, err := b.rprKeyword.ReadRowsKeywordApproved(ctx, map[string]interface{}{}, request.Page, request.Limit, search)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseKeywordGetApprovedData{
				Id:         d.Id,
				Keyword:    d.Keyword,
				Action:     d.Action,
				Channel:    d.Channel,
				ApprovedBy: d.ApprovedBy,
				CreatedBy:  d.CreatedBy,
				CreatedAt:  d.CreatedAt,
				UpdatedBy:  d.UpdatedBy,
				UpdatedAt:  d.UpdatedAt,
			}
			if d.ApprovedAt != nil {
				result.ApprovedAt = *d.ApprovedAt
			} else {
				result.ApprovedAt = basicObject.BlankString
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseKeywordGetApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) KeywordGetApprovalDetail(ctx context.Context, request *RequestKeywordGetApprovalDetail) (*ResponseKeywordGetApprovalDetail, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "KDT008"
	const fName = "usecases.uscKeywordManagement.KeywordGetApprovalDetail"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_KeywordGetApprovalDetail */
	if request.ApprovalContent {
		// Retrieve
		currentApprovalEntity, err := b.rprKeyword.ReadRowKeywordApproval(ctx, map[string]interface{}{
			"id": request.Id,
		})
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		if currentApprovalEntity == nil {
			errMsg := fmt.Errorf("data not found")
			logging.Errorw(fName, "reason", errMsg.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		// response
		response := new(ResponseKeywordGetApprovalDetail)
		status, _ := strconv.Atoi(currentApprovalEntity.Status)
		response.Status = strconv.Itoa(status - 1)
		response.Id = currentApprovalEntity.Id
		response.Keyword = currentApprovalEntity.Keyword
		response.Action = currentApprovalEntity.Action
		response.CreatedAt = currentApprovalEntity.CreatedAt
		response.UpdatedAt = currentApprovalEntity.UpdatedAt
		response.CreatedBy = currentApprovalEntity.CreatedBy
		response.ApprovalType = currentApprovalEntity.ApprovalType
		response.ApprovalBy = currentApprovalEntity.ApprovedBy
		response.RejectNote = currentApprovalEntity.Note
		if currentApprovalEntity.ApprovedAt != nil {
			response.ApprovedAt = *currentApprovalEntity.ApprovedAt
		}
		return response, nil
	} else {
		currentKeyword, err := b.rprKeyword.ReadRowKeyword(ctx, map[string]interface{}{
			"id": request.Id,
		})
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		if currentKeyword == nil {
			errMsg := fmt.Errorf("data not found")
			logging.Errorw(fName, "reason", errMsg.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		// response
		response := new(ResponseKeywordGetApprovalDetail)
		response.Id = currentKeyword.Id
		response.Keyword = currentKeyword.Keyword
		response.Action = currentKeyword.Action
		response.CreatedAt = currentKeyword.CreatedAt
		response.UpdatedAt = currentKeyword.UpdatedAt
		response.CreatedBy = currentKeyword.CreatedBy
		response.ApprovalType = ""
		response.ApprovalBy = currentKeyword.ApprovedBy
		response.RejectNote = ""
		if currentKeyword.ApprovedAt != nil {
			response.ApprovedAt = *currentKeyword.ApprovedAt
		}
		return response, nil
	}
}
