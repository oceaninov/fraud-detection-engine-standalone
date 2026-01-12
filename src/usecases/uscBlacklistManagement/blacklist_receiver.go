package uscBlacklistManagement

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/pkg/humanTime"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistReceiver"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	ResponseCreateBlacklistReceiver struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestCreateBlacklistReceiver struct {
		PhoneNumber     string `json:"phone_number"`
		BeneficiaryName string `json:"beneficiary_name"`
		TransactionType []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"transaction_types"`
	}
)

type (
	ResponseDeleteBlacklistReceiver struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestDeleteBlacklistReceiver struct {
		Id string `param:"id"`
	}
)

type (
	ResponseRejectBlacklistReceiver struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestRejectBlacklistReceiver struct {
		Id          string `param:"id"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"actionTaker"`
	}
)

type (
	ResponseApproveBlacklistReceiver struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestApproveBlacklistReceiver struct {
		Id          string `param:"id"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"actionTaker"`
	}
)

type (
	ResponseRetrieveBlacklistReceiverApprovalData struct {
		Id               string    `json:"id"`
		PhoneNumber      string    `json:"phoneNumber"`
		ApprovalType     string    `json:"approvalType"`
		Status           string    `json:"status"`
		Note             string    `json:"rejectNote"`
		Event            string    `json:"event"`
		BlacklistId      string    `json:"blacklistId"`
		BeneficiaryName  string    `json:"beneficiaryName"`
		TransactionTypes []TrxType `json:"transactionTypes"`
		CreatedBy        string    `json:"createdBy"`
		CreatedAt        string    `json:"createdAt"`
		UpdatedBy        string    `json:"updatedBy"`
		UpdatedAt        string    `json:"updatedAt"`
		ApprovedBy       string    `json:"approvedBy"`
		ApprovedAt       string    `json:"approvedAt"`
	}
	ResponseRetrieveBlacklistReceiverApproval struct {
		ResponseCode    bool                                            `json:"success"`
		ResponseMessage string                                          `json:"messages"`
		Data            []ResponseRetrieveBlacklistReceiverApprovalData `json:"blacklistApproval"`
		Meta            basicObject.Meta                                `json:"meta"`
	}
	RequestRetrieveBlacklistReceiverApproval struct {
		StatusApproval int `query:"statusApproval"`
		Page           int `query:"page"`
		Limit          int `query:"limit"`
	}
)

type (
	ResponseRetrieveBlacklistReceiverApprovedData struct {
		Id               string    `json:"id"`
		PhoneNumber      string    `json:"phoneNumber"`
		CreatedAt        string    `json:"createdAt"`
		CreatedBy        string    `json:"createdBy"`
		UpdatedAt        string    `json:"updatedAt"`
		UpdatedBy        string    `json:"updatedBy"`
		Status           string    `json:"status"`
		BeneficiaryName  string    `json:"beneficiaryName"`
		TransactionTypes []TrxType `json:"transactionTypes"`
	}
	ResponseRetrieveBlacklistReceiverApproved struct {
		ResponseCode    bool                                            `json:"success"`
		ResponseMessage string                                          `json:"messages"`
		Data            []ResponseRetrieveBlacklistReceiverApprovedData `json:"blacklist"`
		Meta            basicObject.Meta                                `json:"meta"`
	}
	RequestRetrieveBlacklistReceiverApproved struct {
		Page        int    `query:"page"`
		Limit       int    `query:"limit"`
		PhoneNumber string `query:"phone_number"`
	}
)

type (
	ResponseRetrieveBlacklistReceiverApprovedById struct {
		Id               string    `json:"id"`
		PhoneNumber      string    `json:"phoneNumber"`
		CreatedAt        string    `json:"createdAt"`
		CreatedBy        string    `json:"createdBy"`
		UpdatedAt        string    `json:"updatedAt"`
		UpdatedBy        string    `json:"updatedBy"`
		Status           string    `json:"status"`
		BeneficiaryName  string    `json:"beneficiaryName"`
		TransactionTypes []TrxType `json:"transactionTypes"`
	}
	ResponseRetrieveBlacklistReceiverApprovalById struct {
		Id               string    `json:"id"`
		PhoneNumber      string    `json:"phoneNumber"`
		ApprovalType     string    `json:"approvalType"`
		Status           string    `json:"status"`
		Note             string    `json:"rejectNote"`
		Event            string    `json:"event"`
		BlacklistId      string    `json:"blacklistId"`
		BeneficiaryName  string    `json:"beneficiaryName"`
		TransactionTypes []TrxType `json:"transactionTypes"`
		CreatedBy        string    `json:"createdBy"`
		CreatedAt        string    `json:"createdAt"`
		UpdatedBy        string    `json:"updatedBy"`
		UpdatedAt        string    `json:"updatedAt"`
		ApprovedBy       string    `json:"approvedBy"`
		ApprovedAt       string    `json:"approvedAt"`
	}
	RequestRetrieveBlacklistReceiverById struct {
		Id string `param:"id"`
	}
)

func (b *blueprint) BlacklistReceiverApprovalGetById(ctx context.Context, request *RequestRetrieveBlacklistReceiverById) (*ResponseRetrieveBlacklistReceiverApprovalById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT001"
	const fName = "usecases.uscBlacklistManagement.BlacklistReceiverApprovalGetById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReceiverApprovalGetById */
	// read current data
	currentEntity, err := b.rprBlacklistReceiver.ReadRowBlacklistReceiverApproval(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentEntity == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", "data not found")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// retrieve trx type
	trxTypeData := []TrxType{}
	trxTypeIds := strings.Split(currentEntity.TransactionTypes, ",")
	for _, si := range trxTypeIds {
		trxTypeDetail, _ := b.rprTransactionType.ReadRowTransactionType(ctx, map[string]interface{}{
			"name": si,
		})
		if trxTypeDetail != nil {
			trxTypeData = append(trxTypeData, TrxType{
				Id:        trxTypeDetail.Id,
				Name:      trxTypeDetail.Name,
				CreatedAt: trxTypeDetail.CreatedAt,
				UpdatedAt: trxTypeDetail.UpdatedAt,
			})
		}
	}

	// response
	response := new(ResponseRetrieveBlacklistReceiverApprovalById)
	response.Id = currentEntity.Id
	response.PhoneNumber = currentEntity.PhoneNumber
	response.ApprovalType = currentEntity.ApprovalType
	response.Status = currentEntity.Status
	response.Note = currentEntity.Note
	response.Event = currentEntity.Event
	response.BlacklistId = currentEntity.BlacklistId
	response.BeneficiaryName = currentEntity.BeneficiaryName
	response.TransactionTypes = trxTypeData
	response.CreatedBy = currentEntity.CreatedBy
	response.CreatedAt = currentEntity.CreatedAt
	response.UpdatedBy = currentEntity.UpdatedBy
	response.UpdatedAt = currentEntity.UpdatedAt
	response.ApprovedBy = currentEntity.ApprovedBy
	if currentEntity.ApprovedAt != nil {
		response.ApprovedAt = *currentEntity.ApprovedAt
	} else {
		response.ApprovedAt = basicObject.BlankString
	}
	status, _ := strconv.Atoi(currentEntity.Status)
	response.Status = strconv.Itoa(status - 1)
	return response, nil
}

func (b *blueprint) BlacklistReceiverApprovedGetById(ctx context.Context, request *RequestRetrieveBlacklistReceiverById) (*ResponseRetrieveBlacklistReceiverApprovedById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT001"
	const fName = "usecases.uscBlacklistManagement.BlacklistReceiverApprovedGetById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReceiverApprovedGetById */
	// read current data
	currentEntity, err := b.rprBlacklistReceiver.ReadRowBlacklistReceiver(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentEntity == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", "data not found")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// retrieve trx type
	trxTypeData := []TrxType{}
	trxTypeIds := strings.Split(currentEntity.TransactionTypes, ",")
	for _, si := range trxTypeIds {
		trxTypeDetail, _ := b.rprTransactionType.ReadRowTransactionType(ctx, map[string]interface{}{
			"name": si,
		})
		if trxTypeDetail != nil {
			trxTypeData = append(trxTypeData, TrxType{
				Id:        trxTypeDetail.Id,
				Name:      trxTypeDetail.Name,
				CreatedAt: trxTypeDetail.CreatedAt,
				UpdatedAt: trxTypeDetail.UpdatedAt,
			})
		}
	}

	// response
	response := new(ResponseRetrieveBlacklistReceiverApprovedById)
	response.Id = currentEntity.Id
	response.PhoneNumber = currentEntity.PhoneNumber
	response.CreatedAt = currentEntity.CreatedAt
	response.CreatedBy = currentEntity.CreatedBy
	response.UpdatedAt = currentEntity.UpdatedAt
	response.UpdatedBy = currentEntity.UpdatedBy
	response.Status = currentEntity.Status
	response.BeneficiaryName = currentEntity.BeneficiaryName
	response.TransactionTypes = trxTypeData
	status, _ := strconv.Atoi(currentEntity.Status)
	response.Status = strconv.Itoa(status - 1)
	return response, nil
}

func (b *blueprint) BlacklistReceiverCreate(ctx context.Context, request *RequestCreateBlacklistReceiver) (*ResponseCreateBlacklistReceiver, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT001"
	const fName = "usecases.uscBlacklistManagement.BlacklistReceiverCreate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReceiverCreate */
	// read current data
	currentEntity, err := b.rprBlacklistReceiver.ReadRowBlacklistReceiver(ctx, map[string]interface{}{
		"phone_number": request.PhoneNumber,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentEntity != nil {
		responseMessage := basicObject.ResponseError{
			Message:    "blacklist already exist",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "already exist on approved data")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}
	// read current approval data
	currentApprovalEntity, err := b.rprBlacklistReceiver.ReadRowBlacklistReceiverApproval(ctx, map[string]interface{}{
		"phone_number": request.PhoneNumber,
		//"beneficiary_name": request.BeneficiaryName,
		"status": 1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity != nil {
		responseMessage := basicObject.ResponseError{
			Message:    "blacklist already exist",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "already exist on approval data")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	// create approval data
	newEntityApproval := new(rprBlacklistReceiver.EntityBlacklistReceiverApproval)
	newEntityApproval.Id = guuid.NewString()
	newEntityApproval.PhoneNumber = request.PhoneNumber
	newEntityApproval.ApprovalType = basicObject.CreateApprovalType
	newEntityApproval.Status = basicObject.ApprovalPending
	newEntityApproval.Note = basicObject.BlankString
	newEntityApproval.Event = basicObject.InsertEvent
	newEntityApproval.BlacklistId = guuid.NewString()
	newEntityApproval.BeneficiaryName = request.BeneficiaryName
	newEntityApproval.ApprovedBy = basicObject.BlankString
	newEntityApproval.ApprovedAt = nil
	newEntityApproval.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	newEntityApproval.CreatedBy = email
	newEntityApproval.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	newEntityApproval.UpdatedBy = email
	var transactionTypesId []string
	for _, tt := range request.TransactionType {
		transactionTypesId = append(transactionTypesId, tt.Name)
	}
	newEntityApproval.TransactionTypes = strings.Join(transactionTypesId, ",")
	insertedEntity, insertedId, err := b.rprBlacklistReceiver.WriteRowBlacklistReceiverApproval(ctx, *newEntityApproval)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if insertedId == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if insertedEntity == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseCreateBlacklistReceiver)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Blacklist has been requested please call approval admin"
	response.Id = newEntityApproval.Id
	return response, nil
}

func (b *blueprint) BlacklistReceiverDelete(ctx context.Context, request *RequestDeleteBlacklistReceiver) (*ResponseDeleteBlacklistReceiver, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT002"
	const fName = "usecases.uscBlacklistManagement.BlacklistReceiverDelete"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReceiverDelete */
	// read current data
	currentEntity, err := b.rprBlacklistReceiver.ReadRowBlacklistReceiver(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentEntity == nil {
		errMsg := fmt.Errorf("data not found on approved data")
		logging.Errorw(fName, "reason", "data not found on approved data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// read current approval data
	currentApprovalEntity, err := b.rprBlacklistReceiver.ReadRowBlacklistReceiverApproval(ctx, map[string]interface{}{
		"blacklist_id": request.Id,
		"status":       basicObject.ApprovalPending,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity != nil {
		errMsg := fmt.Errorf("already exist on approval data")
		logging.Errorw(fName, "reason", "already exist on approval data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// create approval data
	newEntityApproval := new(rprBlacklistReceiver.EntityBlacklistReceiverApproval)
	newEntityApproval.Id = guuid.NewString()
	newEntityApproval.PhoneNumber = currentEntity.PhoneNumber
	newEntityApproval.ApprovalType = basicObject.DeleteApprovalType
	newEntityApproval.Status = basicObject.ApprovalPending
	newEntityApproval.Note = basicObject.BlankString
	newEntityApproval.Event = basicObject.RemoveEvent
	newEntityApproval.BlacklistId = currentEntity.Id
	newEntityApproval.BeneficiaryName = currentEntity.BeneficiaryName
	newEntityApproval.ApprovedBy = basicObject.BlankString
	newEntityApproval.ApprovedAt = nil
	newEntityApproval.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	newEntityApproval.CreatedBy = email
	newEntityApproval.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	newEntityApproval.UpdatedBy = email
	transactionTypesId := strings.Split(currentEntity.TransactionTypes, ",")
	newEntityApproval.TransactionTypes = strings.Join(transactionTypesId, ",")
	insertedEntity, insertedId, err := b.rprBlacklistReceiver.WriteRowBlacklistReceiverApproval(ctx, *newEntityApproval)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if insertedId == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if insertedEntity == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseDeleteBlacklistReceiver)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Blacklist has been requested to delete please call approval admin"
	response.Id = newEntityApproval.Id
	return response, nil
}

func (b *blueprint) BlacklistReceiverReject(ctx context.Context, request *RequestRejectBlacklistReceiver) (*ResponseRejectBlacklistReceiver, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT003"
	const fName = "usecases.uscBlacklistManagement.BlacklistReceiverReject"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReceiverReject */
	err := b.rprBlacklistReceiver.WriteRejectForBlacklistApproval(ctx, request.Id, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseRejectBlacklistReceiver)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Reject Blacklist"
	return response, nil
}

func (b *blueprint) BlacklistReceiverApprove(ctx context.Context, request *RequestApproveBlacklistReceiver) (*ResponseApproveBlacklistReceiver, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT004"
	const fName = "usecases.uscBlacklistManagement.BlacklistReceiverApprove"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReceiverApprove */
	// read current approval data
	currentApprovalEntity, err := b.rprBlacklistReceiver.ReadRowBlacklistReceiverApproval(ctx, map[string]interface{}{
		"id":     request.Id,
		"status": basicObject.ApprovalPending,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity == nil {
		errMsg := fmt.Errorf("data not found on approval data")
		logging.Errorw(fName, "reason", "data not found on approval data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	err = b.rprBlacklistReceiver.WriteApprovalForBlacklistApproval(ctx, request.Id, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	switch currentApprovalEntity.Event {
	case basicObject.InsertEvent:
		currentTime := time.Now().Format(basicObject.DateAndTime)
		newEntityData := new(rprBlacklistReceiver.EntityBlacklistReceiver)
		newEntityData.Id = currentApprovalEntity.BlacklistId
		newEntityData.PhoneNumber = currentApprovalEntity.PhoneNumber
		newEntityData.CreatedAt = currentTime
		newEntityData.CreatedBy = email
		newEntityData.UpdatedAt = currentTime
		newEntityData.UpdatedBy = email
		newEntityData.ApprovedAt = &currentTime
		newEntityData.ApprovedBy = email
		newEntityData.Status = currentApprovalEntity.Status
		newEntityData.BeneficiaryName = currentApprovalEntity.BeneficiaryName
		newEntityData.TransactionTypes = currentApprovalEntity.TransactionTypes
		_, _, err := b.rprBlacklistReceiver.WriteRowBlacklistReceiver(ctx, *newEntityData)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	case basicObject.RemoveEvent:
		err := b.rprBlacklistReceiver.RemoveRowBlacklistReceiver(ctx, currentApprovalEntity.BlacklistId)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	}

	// response
	response := new(ResponseApproveBlacklistReceiver)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Approve Blacklist"
	return response, nil
}

func (b *blueprint) BlacklistReceiverApprovalGet(ctx context.Context, request *RequestRetrieveBlacklistReceiverApproval) (*ResponseRetrieveBlacklistReceiverApproval, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT005"
	const fName = "usecases.uscBlacklistManagement.BlacklistReceiverApprovalGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReceiverApprovalGet */
	dataResult := []ResponseRetrieveBlacklistReceiverApprovalData{}
	queryData := map[string]interface{}{}
	//if request.StatusApproval != 0 {
	//	queryData["status"] = request.StatusApproval
	//}
	data, meta, err := b.rprBlacklistReceiver.ReadRowsBlacklistReceiverApproval(ctx, queryData, request.Page, request.Limit, "")
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseRetrieveBlacklistReceiverApprovalData{
				Id:              d.Id,
				PhoneNumber:     d.PhoneNumber,
				ApprovalType:    d.ApprovalType,
				Note:            d.Note,
				Event:           d.Event,
				BlacklistId:     d.BlacklistId,
				BeneficiaryName: d.BeneficiaryName,
				ApprovedBy:      d.ApprovedBy,
				CreatedBy:       d.CreatedBy,
				UpdatedBy:       d.UpdatedBy,
			}

			trxTypeData := []TrxType{}
			trxTypeIds := strings.Split(d.TransactionTypes, ",")
			for _, si := range trxTypeIds {
				trxTypeDetail, _ := b.rprTransactionType.ReadRowTransactionType(ctx, map[string]interface{}{
					"name": si,
				})
				if trxTypeDetail != nil {
					trxTypeData = append(trxTypeData, TrxType{
						Id:        trxTypeDetail.Id,
						Name:      trxTypeDetail.Name,
						CreatedAt: trxTypeDetail.CreatedAt,
						UpdatedAt: trxTypeDetail.UpdatedAt,
					})
				}
			}

			//createdAt, _ := humanTime.FormatDate(d.CreatedAt)
			//updatedAt, _ := humanTime.FormatDate(d.UpdatedAt)
			result.UpdatedAt = d.UpdatedAt
			result.CreatedAt = d.CreatedAt

			status, _ := strconv.Atoi(d.Status)
			result.Status = strconv.Itoa(status - 1)
			if d.ApprovedAt != nil {
				result.ApprovedAt = *d.ApprovedAt
			} else {
				result.ApprovedAt = basicObject.BlankString
			}
			result.TransactionTypes = trxTypeData
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseRetrieveBlacklistReceiverApproval)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistReceiverApprovedGet(ctx context.Context, request *RequestRetrieveBlacklistReceiverApproved) (*ResponseRetrieveBlacklistReceiverApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT006"
	const fName = "usecases.uscBlacklistManagement.BlacklistReceiverApprovedGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReceiverApprovedGet */
	dataResult := []ResponseRetrieveBlacklistReceiverApprovedData{}
	queryData := map[string]interface{}{}
	search := ""
	if request.PhoneNumber != "" && request.PhoneNumber != "230" {
		search = request.PhoneNumber
	}
	data, meta, err := b.rprBlacklistReceiver.ReadRowsBlacklistReceiver(ctx, queryData, request.Page, request.Limit, search)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dataSender := ResponseRetrieveBlacklistReceiverApprovedData{
				Id:              d.Id,
				PhoneNumber:     d.PhoneNumber,
				CreatedBy:       d.CreatedBy,
				UpdatedBy:       d.UpdatedBy,
				Status:          d.Status,
				BeneficiaryName: d.BeneficiaryName,
			}

			createdAt, _ := humanTime.FormatDate(d.CreatedAt)
			updatedAt, _ := humanTime.FormatDate(d.UpdatedAt)
			dataSender.UpdatedAt = updatedAt
			dataSender.CreatedAt = createdAt

			var trxTypeData []TrxType
			trxTypeIds := strings.Split(d.TransactionTypes, ",")
			for _, si := range trxTypeIds {
				trxTypeDetail, _ := b.rprTransactionType.ReadRowTransactionType(ctx, map[string]interface{}{
					"name": si,
				})
				if trxTypeDetail != nil {
					trxTypeData = append(trxTypeData, TrxType{
						Id:        trxTypeDetail.Id,
						Name:      trxTypeDetail.Name,
						CreatedAt: trxTypeDetail.CreatedAt,
						UpdatedAt: trxTypeDetail.UpdatedAt,
					})
				}
			}
			dataSender.TransactionTypes = trxTypeData
			dataResult = append(dataResult, dataSender)
		}
	}

	// response
	response := new(ResponseRetrieveBlacklistReceiverApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}
