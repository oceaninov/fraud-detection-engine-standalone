package uscBlacklistManagement

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/pkg/humanTime"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistSender"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	ResponseCreateBlacklistSender struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestCreateBlacklistSender struct {
		PhoneNumber     string `json:"phone_number"`
		BeneficiaryName string `json:"beneficiary_name"`
		TransactionType []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"transaction_types"`
	}
)

type (
	ResponseDeleteBlacklistSender struct {
		ResponseCode    bool   `json:"statCode"`
		ResponseMessage string `json:"statMsg"`
		Id              string `json:"id"`
	}
	RequestDeleteBlacklistSender struct {
		Id string `param:"id"`
	}
)

type (
	ResponseRejectBlacklistSender struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestRejectBlacklistSender struct {
		Id          string `param:"id"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"actionTaker"`
	}
)

type (
	ResponseApproveBlacklistSender struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestApproveBlacklistSender struct {
		Id          string `param:"id"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"actionTaker"`
	}
)

type (
	ResponseRetrieveBlacklistSenderApprovalData struct {
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
	ResponseRetrieveBlacklistSenderApproval struct {
		ResponseCode    bool                                          `json:"success"`
		ResponseMessage string                                        `json:"messages"`
		Data            []ResponseRetrieveBlacklistSenderApprovalData `json:"blacklistApproval"`
		Meta            basicObject.Meta                              `json:"meta"`
	}
	RequestRetrieveBlacklistSenderApproval struct {
		StatusApproval int `query:"statusApproval"`
		Page           int `query:"page"`
		Limit          int `query:"limit"`
	}
)

type TrxType struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type (
	ResponseRetrieveBlacklistSenderApprovedData struct {
		Id               string    `json:"id"`
		PhoneNumber      string    `json:"phoneNumber"`
		CreatedAt        string    `json:"createdAt"`
		CreatedBy        string    `json:"createdBy"`
		UpdatedAt        string    `json:"updatedAt"`
		UpdatedBy        string    `json:"updatedBy"`
		ApprovedBy       string    `json:"approvedBy"`
		ApprovedAt       string    `json:"approvedAt"`
		Status           string    `json:"status"`
		BeneficiaryName  string    `json:"beneficiaryName"`
		TransactionTypes []TrxType `json:"transactionTypes"`
	}
	ResponseRetrieveBlacklistSenderApproved struct {
		ResponseCode    bool                                          `json:"success"`
		ResponseMessage string                                        `json:"messages"`
		Data            []ResponseRetrieveBlacklistSenderApprovedData `json:"blacklist"`
		Meta            basicObject.Meta                              `json:"meta"`
	}
	RequestRetrieveBlacklistSenderApproved struct {
		Page        int    `query:"page"`
		Limit       int    `query:"limit"`
		PhoneNumber string `query:"phone_number"`
	}
)

type (
	ResponseRetrieveBlacklistSenderApprovedById struct {
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
	ResponseRetrieveBlacklistSenderApprovalById struct {
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
	RequestRetrieveBlacklistSenderById struct {
		Id string `param:"id"`
	}
)

func (b *blueprint) BlacklistSenderApprovalGetById(ctx context.Context, request *RequestRetrieveBlacklistSenderById) (*ResponseRetrieveBlacklistSenderApprovalById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT001"
	const fName = "usecases.uscBlacklistManagement.BlacklistSenderApprovalGetById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistSenderApprovalGetById */
	// read current data
	currentEntity, err := b.rprBlacklistSender.ReadRowBlacklistSenderApproval(ctx, map[string]interface{}{
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
	response := new(ResponseRetrieveBlacklistSenderApprovalById)
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

func (b *blueprint) BlacklistSenderApprovedGetById(ctx context.Context, request *RequestRetrieveBlacklistSenderById) (*ResponseRetrieveBlacklistSenderApprovedById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT001"
	const fName = "usecases.uscBlacklistManagement.BlacklistSenderApprovedGetById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistSenderApprovedGetById */
	// read current data
	currentEntity, err := b.rprBlacklistSender.ReadRowBlacklistSender(ctx, map[string]interface{}{
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
	response := new(ResponseRetrieveBlacklistSenderApprovedById)
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

func (b *blueprint) BlacklistSenderCreate(ctx context.Context, request *RequestCreateBlacklistSender) (*ResponseCreateBlacklistSender, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT001"
	const fName = "usecases.uscBlacklistManagement.BlacklistSenderCreate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistSenderCreate */
	// read current data
	currentEntity, err := b.rprBlacklistSender.ReadRowBlacklistSender(ctx, map[string]interface{}{
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
	currentApprovalEntity, err := b.rprBlacklistSender.ReadRowBlacklistSenderApproval(ctx, map[string]interface{}{
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
	newEntityApproval := new(rprBlacklistSender.EntityBlacklistSenderApproval)
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
	insertedEntity, insertedId, err := b.rprBlacklistSender.WriteRowBlacklistSenderApproval(ctx, *newEntityApproval)
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
	response := new(ResponseCreateBlacklistSender)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Blacklist has been requested to create please call approval admin"
	response.Id = newEntityApproval.Id
	return response, nil
}

func (b *blueprint) BlacklistSenderDelete(ctx context.Context, request *RequestDeleteBlacklistSender) (*ResponseDeleteBlacklistSender, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT002"
	const fName = "usecases.uscBlacklistManagement.BlacklistSenderDelete"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistSenderDelete */
	// read current data
	currentEntity, err := b.rprBlacklistSender.ReadRowBlacklistSender(ctx, map[string]interface{}{
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
	currentApprovalEntity, err := b.rprBlacklistSender.ReadRowBlacklistSenderApproval(ctx, map[string]interface{}{
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
	newEntityApproval := new(rprBlacklistSender.EntityBlacklistSenderApproval)
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
	insertedEntity, insertedId, err := b.rprBlacklistSender.WriteRowBlacklistSenderApproval(ctx, *newEntityApproval)
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
	response := new(ResponseDeleteBlacklistSender)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Blacklist has been requested to delete please call approval admin"
	response.Id = newEntityApproval.Id
	return response, nil
}

func (b *blueprint) BlacklistSenderReject(ctx context.Context, request *RequestRejectBlacklistSender) (*ResponseRejectBlacklistSender, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT003"
	const fName = "usecases.uscBlacklistManagement.BlacklistSenderReject"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistSenderReject */
	err := b.rprBlacklistSender.WriteRejectForBlacklistApproval(ctx, request.Id, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseRejectBlacklistSender)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Reject Blacklist"
	return response, nil
}

func (b *blueprint) BlacklistSenderApprove(ctx context.Context, request *RequestApproveBlacklistSender) (*ResponseApproveBlacklistSender, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT004"
	const fName = "usecases.uscBlacklistManagement.BlacklistSenderApprove"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistSenderApprove */
	// read current approval data
	currentApprovalEntity, err := b.rprBlacklistSender.ReadRowBlacklistSenderApproval(ctx, map[string]interface{}{
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

	err = b.rprBlacklistSender.WriteApprovalForBlacklistApproval(ctx, request.Id, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	switch currentApprovalEntity.Event {
	case basicObject.InsertEvent:
		currentTime := time.Now().Format(basicObject.DateAndTime)
		newEntityData := new(rprBlacklistSender.EntityBlacklistSender)
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
		_, _, err := b.rprBlacklistSender.WriteRowBlacklistSender(ctx, *newEntityData)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	case basicObject.RemoveEvent:
		err := b.rprBlacklistSender.RemoveRowBlacklistSender(ctx, currentApprovalEntity.BlacklistId)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	}

	// response
	response := new(ResponseApproveBlacklistSender)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Approve Blacklist"
	return response, nil
}

func (b *blueprint) BlacklistSenderApprovalGet(ctx context.Context, request *RequestRetrieveBlacklistSenderApproval) (*ResponseRetrieveBlacklistSenderApproval, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT005"
	const fName = "usecases.uscBlacklistManagement.BlacklistSenderApprovalGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistSenderApprovalGet */
	dataResult := []ResponseRetrieveBlacklistSenderApprovalData{}
	queryData := map[string]interface{}{}
	//if request.StatusApproval != 0 {
	//	queryData["status"] = request.StatusApproval
	//}
	data, meta, err := b.rprBlacklistSender.ReadRowsBlacklistSenderApproval(ctx, queryData, request.Page, request.Limit, "")
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseRetrieveBlacklistSenderApprovalData{
				Id:              d.Id,
				PhoneNumber:     d.PhoneNumber,
				ApprovalType:    d.ApprovalType,
				Note:            d.Note,
				Event:           d.Event,
				BlacklistId:     d.BlacklistId,
				BeneficiaryName: d.BeneficiaryName,
				ApprovedBy:      d.ApprovedBy,
				CreatedAt:       d.CreatedAt,
				CreatedBy:       d.CreatedBy,
				UpdatedAt:       d.UpdatedAt,
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
	response := new(ResponseRetrieveBlacklistSenderApproval)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistSenderApprovedGet(ctx context.Context, request *RequestRetrieveBlacklistSenderApproved) (*ResponseRetrieveBlacklistSenderApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT006"
	const fName = "usecases.uscBlacklistManagement.BlacklistSenderApprovedGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistSenderApprovedGet */
	dataResult := []ResponseRetrieveBlacklistSenderApprovedData{}
	queryData := map[string]interface{}{}
	search := ""
	if request.PhoneNumber != "" && request.PhoneNumber != "230" {
		search = request.PhoneNumber
	}
	data, meta, err := b.rprBlacklistSender.ReadRowsBlacklistSender(ctx, queryData, request.Page, request.Limit, search)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dataSender := ResponseRetrieveBlacklistSenderApprovedData{
				Id:              d.Id,
				PhoneNumber:     d.PhoneNumber,
				CreatedBy:       d.CreatedBy,
				UpdatedBy:       d.UpdatedBy,
				ApprovedBy:      d.ApprovedBy,
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
			if d.ApprovedAt != nil {
				dataSender.ApprovedAt = *d.ApprovedAt
			}
			dataResult = append(dataResult, dataSender)
		}
	}

	// response
	response := new(ResponseRetrieveBlacklistSenderApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}
