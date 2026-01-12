package uscBlacklistManagement

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/csvProcessor"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistMerchant"
	"strconv"
	"strings"
	"time"
)

type (
	ResponseUploadBlacklistMerchant struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
		FileId          string `json:"fileId"`
	}
	RequestUploadBlacklistMerchant struct{}
)

type (
	ResponseRejectBlacklistMerchant struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestRejectBlacklistMerchant struct {
		FileId      string `param:"fileId"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"creator"`
	}
)

type (
	ResponseApproveBlacklistMerchant struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestApproveBlacklistMerchant struct {
		FileId      string `param:"fileId"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"creator"`
	}
)

type (
	ResponseRetrieveBlacklistMerchantApprovalData struct {
		ID           string `json:"ID"`
		FileLink     string `json:"FileLink"`
		FileName     string `json:"FileName"`
		FilePath     string `json:"FilePath"`
		ApprovalType string `json:"ApprovalType"`
		Note         string `json:"Note"`
		Active       string `json:"Active"`
		Status       string `json:"Status"`
		CreatedBy    string `json:"CreatedBy"`
		CreatedAt    string `json:"CreatedAt"`
		UpdatedBy    string `json:"UpdatedBy"`
		UpdatedAt    string `json:"UpdatedAt"`
		ApprovedBy   string `json:"ApprovedBy"`
		ApprovedAt   string `json:"ApprovedAt"`
	}
	ResponseRetrieveBlacklistMerchantApproval struct {
		ResponseCode    bool                                            `json:"success"`
		ResponseMessage string                                          `json:"messages"`
		Data            []ResponseRetrieveBlacklistMerchantApprovalData `json:"data"`
		Meta            basicObject.Meta                                `json:"meta"`
	}
	RequestRetrieveBlacklistMerchantApproval struct {
		StatusApproval int `json:"statusApproval"`
		Page           int `json:"page"`
		Limit          int `json:"limit"`
	}
)

type (
	RequestRetrieveBlacklistMerchantApproved struct{}
)

type (
	RequestRetrieveBlacklistMerchantTemplate struct{}
)

type (
	ResponseRetrieveActiveMerchantData struct {
		Id           string `json:"id"`
		CreatedAt    string `json:"createdAt"`
		CreatedBy    string `json:"createdBy"`
		FileLink     string `json:"fileLink"`
		FileName     string `json:"fileName"`
		FilePath     string `json:"filePath"`
		Active       bool   `json:"active"`
		ApprovalType string `json:"approvalType"`
		ApprovalBy   string `json:"approvalBy"`
		ApprovedAt   string `json:"approvedAt"`
		RejectNote   string `json:"rejectNote"`
	}
	ResponseRetrieveActiveMerchant struct {
		ResponseCode    bool                                 `json:"success"`
		ResponseMessage string                               `json:"messages"`
		Data            []ResponseRetrieveActiveMerchantData `json:"blacklistApproval"`
		Meta            basicObject.Meta                     `json:"meta"`
	}
	RequestRetrieveActiveMerchant struct {
		Page   int    `query:"page"`
		Limit  int    `query:"limit"`
		Search string `query:"search"`
	}
)

type (
	ResponseRetrieveActiveMerchantContentData struct {
		Id           string `json:"id"`
		NMID         string `json:"nmid"`
		MerchantName string `json:"name"`
		Datasource   string `json:"datasource"`
		FileId       string `json:"fileId"`
		FileLink     string `json:"fileLink"`
		CreatedBy    string `json:"createdBy"`
		CreatedAt    string `json:"createdAt"`
		UpdatedBy    string `json:"updatedBy"`
		UpdatedAt    string `json:"updatedAt"`
		ApprovedBy   string `json:"approvedBy"`
		ApprovedAt   string `json:"approvedAt"`
	}
	ResponseRetrieveActiveMerchantContent struct {
		ResponseCode    bool                                        `json:"success"`
		ResponseMessage string                                      `json:"messages"`
		Data            []ResponseRetrieveActiveMerchantContentData `json:"blacklistMerchant"`
		Meta            basicObject.Meta                            `json:"meta"`
	}
	RequestRetrieveActiveMerchantContent struct {
		Page   int    `query:"page"`
		Limit  int    `query:"limit"`
		Search string `query:"search"`
	}
)

type (
	ResponseRetrieveActiveMerchantContentById struct {
		Id           string `json:"id"`
		NMID         string `json:"nmid"`
		MerchantName string `json:"name"`
		Datasource   string `json:"datasource"`
		FileId       string `json:"fileId"`
		FileLink     string `json:"fileLink"`
		CreatedBy    string `json:"createdBy"`
		CreatedAt    string `json:"createdAt"`
		UpdatedBy    string `json:"updatedBy"`
		UpdatedAt    string `json:"updatedAt"`
		ApprovedBy   string `json:"approvedBy"`
		ApprovedAt   string `json:"approvedAt"`
	}
	RequestRetrieveActiveMerchantContentById struct {
		Id string `param:"id"`
	}
)

type (
	RequestMerchantDownloadContentByFileId struct {
		Id string `param:"id"`
	}
)

type (
	ResponseBlacklistMerchantFileById struct {
		ID           string `json:"id"`
		FileLink     string `json:"fileLink"`
		FileName     string `json:"fileName"`
		FilePath     string `json:"filePath"`
		ApprovalType string `json:"approvalType"`
		Note         string `json:"rejectNote"`
		Status       string `json:"status"`
		Active       string `json:"active"`
		CreatedBy    string `json:"createdBy"`
		CreatedAt    string `json:"createdAt"`
		UpdatedBy    string `json:"updatedBy"`
		UpdatedAt    string `json:"updatedAt"`
		ApprovedBy   string `json:"approvedBy"`
		ApprovedAt   string `json:"approvedAt"`
	}
	RequestBlacklistMerchantFileById struct {
		Id string `param:"id"`
	}
)

type (
	ResponseBlacklistMerchantFileContentData struct {
		Id           string `json:"id"`
		NMID         string `json:"nmid"`
		MerchantName string `json:"name"`
		Datasource   string `json:"datasource"`
		FileId       string `json:"fileId"`
		FileLink     string `json:"fileLink"`
		CreatedBy    string `json:"createdBy"`
		CreatedAt    string `json:"createdAt"`
		UpdatedBy    string `json:"updatedBy"`
		UpdatedAt    string `json:"updatedAt"`
		ApprovedBy   string `json:"approvedBy"`
		ApprovedAt   string `json:"approvedAt"`
	}
	ResponseBlacklistMerchantFileContent struct {
		ResponseCode    bool                                       `json:"success"`
		ResponseMessage string                                     `json:"messages"`
		Data            []ResponseBlacklistMerchantFileContentData `json:"blacklistMerchant"`
		Meta            basicObject.Meta                           `json:"meta"`
	}
	RequestBlacklistMerchantFileContent struct {
		Id    string `param:"id"`
		Page  int    `query:"page"`
		Limit int    `query:"limit"`
	}
)

func (b *blueprint) BlacklistMerchantUpload(ctx context.Context, request *RequestUploadBlacklistMerchant) (*ResponseUploadBlacklistMerchant, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT013"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantUpload"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantUpload */
	fileBytes, ok := ctx.Value("X-File-Bytes").([]byte)
	if !ok {
		errMsg := fmt.Errorf("no file attached")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	resultData, err := csvProcessor.ReadCSVForMerchant(fileBytes)
	if err != nil {
		errMsg := fmt.Errorf("csv file processor failed")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// inserting data into database
	fileId := guuid.NewString()
	err = b.rprBlacklistMerchant.WriteBulkRowBlacklistMerchantApproval(ctx, fileId, email, resultData)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// inserting file data
	fileName := strings.Replace(fileId, "-", "", -1)
	fileEntity := new(rprBlacklistMerchant.EntityBlacklistMerchantFile)
	fileEntity.ID = fileId
	fileEntity.FileLink = fmt.Sprintf("https://fds.repository/files/%s.csv", fileName)
	fileEntity.FileName = fmt.Sprintf("%s.csv", fileName)
	fileEntity.FilePath = fmt.Sprintf("/files/%s.csv", fileName)
	fileEntity.ApprovalType = basicObject.CreateApprovalType
	fileEntity.Note = basicObject.BlankString
	fileEntity.Active = "0"
	fileEntity.Status = basicObject.ApprovalPending
	fileEntity.ApprovedBy = basicObject.BlankString
	fileEntity.ApprovedAt = nil
	fileEntity.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	fileEntity.CreatedBy = email
	fileEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	fileEntity.UpdatedBy = email
	fileData, fileDataId, err := b.rprBlacklistMerchant.WriteRowBlacklistMerchantFile(ctx, *fileEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if fileDataId == nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if fileData == nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseUploadBlacklistMerchant)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.FileId = fileData.ID
	return response, nil
}

func (b *blueprint) BlacklistMerchantReject(ctx context.Context, request *RequestRejectBlacklistMerchant) (*ResponseRejectBlacklistMerchant, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT014"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantReject"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantReject */
	currentApprovalEntity, err := b.rprBlacklistMerchant.ReadRowBlacklistMerchantFile(ctx, map[string]interface{}{
		"id":     request.FileId,
		"status": 1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity == nil {
		errMsg := fmt.Errorf("not found on approval data")
		logging.Errorw(fName, "reason", "not found on approval data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	//err = b.rprBlacklistMerchant.RemoveBlacklistMerchantApprovalData(ctx, currentApprovalEntity.ID)
	//if err != nil {
	//	errMsg := fmt.Errorf("internal server error")
	//	logging.Errorw(fName, "reason", err.Error())
	//	return nil, errWrap.WrapRepositoryError(errMsg)
	//}

	err = b.rprBlacklistMerchant.WriteRejectForBlacklistMerchantApproval(ctx, request.FileId, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseRejectBlacklistMerchant)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Reject Blacklist Merchant"
	return response, nil
}

func (b *blueprint) BlacklistMerchantApprove(ctx context.Context, request *RequestApproveBlacklistMerchant) (*ResponseApproveBlacklistMerchant, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT015"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantApprove"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantApprovalGet */
	currentApprovalEntity, err := b.rprBlacklistMerchant.ReadRowBlacklistMerchantFile(ctx, map[string]interface{}{
		"id":     request.FileId,
		"status": 1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity == nil {
		errMsg := fmt.Errorf("not found on approval data")
		logging.Errorw(fName, "reason", "not found on approval data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	switch currentApprovalEntity.ApprovalType {
	case basicObject.CreateApprovalType:
		data, err := b.rprBlacklistMerchant.ReadRowsBlacklistMerchantApproval(ctx, map[string]interface{}{
			"file_id": currentApprovalEntity.ID,
		})
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		if data == nil {
			errMsg := fmt.Errorf("data not found")
			logging.Errorw(fName, "reason", errMsg.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		err = b.rprBlacklistMerchant.WriteBulkRowBlacklistMerchant(ctx, *data)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		//err = b.rprBlacklistMerchant.RemoveBlacklistMerchantApprovalData(ctx, currentApprovalEntity.ID)
		//if err != nil {
		//	errMsg := fmt.Errorf("internal server error")
		//	logging.Errorw(fName, "reason", err.Error())
		//	return nil, errWrap.WrapRepositoryError(errMsg)
		//}
		err = b.rprBlacklistMerchant.RemoveBlacklistMerchantData(ctx, currentApprovalEntity.ID)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		err = b.rprBlacklistMerchant.WriteApprovalForBlacklistMerchantApproval(ctx, currentApprovalEntity.ID, request.Note, email)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", errMsg.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	}

	// response
	response := new(ResponseApproveBlacklistMerchant)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Approve Blacklist Merchant"
	return response, nil
}

func (b *blueprint) BlacklistMerchantApprovalGet(ctx context.Context, request *RequestRetrieveBlacklistMerchantApproval) (*ResponseRetrieveBlacklistMerchantApproval, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT016"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantApprovalGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantApprovalGet */
	dataResult := []ResponseRetrieveBlacklistMerchantApprovalData{}
	data, meta, err := b.rprBlacklistMerchant.ReadRowsBlacklistMerchantFile(ctx, map[string]interface{}{
		"status": request.StatusApproval,
	}, request.Page, request.Limit)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseRetrieveBlacklistMerchantApprovalData{
				ID:           d.ID,
				FileLink:     d.FileLink,
				FileName:     d.FileName,
				FilePath:     d.FilePath,
				ApprovalType: d.ApprovalType,
				Note:         d.Note,
				Status:       d.Status,
				Active:       d.Active,
				ApprovedBy:   d.ApprovedBy,
				CreatedAt:    d.CreatedAt,
				CreatedBy:    d.CreatedBy,
				UpdatedAt:    d.UpdatedAt,
				UpdatedBy:    d.UpdatedBy,
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
	response := new(ResponseRetrieveBlacklistMerchantApproval)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistMerchantApprovedGet(ctx context.Context, request *RequestRetrieveBlacklistMerchantApproved) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantApprovedGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantApprovedGet */
	var dataResult []csvProcessor.MerchantCSVFormat
	data, err := b.rprBlacklistMerchant.ReadRowsBlacklistMerchantApproved(ctx, map[string]interface{}{})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dataResult = append(dataResult, csvProcessor.MerchantCSVFormat{
				NMID:       d.NMID,
				Name:       d.MerchantName,
				DataSource: d.Datasource,
			})
		}
	}
	csvFileContent, err := csvProcessor.ConvertMerchantToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}

func (b *blueprint) BlacklistMerchantTemplate(ctx context.Context, request *RequestRetrieveBlacklistMerchantTemplate) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantTemplate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantTemplate */
	var dataResult []csvProcessor.MerchantCSVFormat
	csvFileContent, err := csvProcessor.ConvertMerchantToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}

func (b *blueprint) BlacklistMerchantActiveData(ctx context.Context, request *RequestRetrieveActiveMerchant) (*ResponseRetrieveActiveMerchant, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantApprovedGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantApprovedGet */
	dataResult := []ResponseRetrieveActiveMerchantData{}
	queryData := map[string]interface{}{}
	data, meta, err := b.rprBlacklistMerchant.ReadRowsBlacklistMerchantFile(ctx, queryData, request.Page, request.Limit)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			active, _ := strconv.ParseBool(d.Active)
			dn := ResponseRetrieveActiveMerchantData{
				Id:         d.ID,
				CreatedAt:  d.CreatedAt,
				CreatedBy:  d.CreatedBy,
				FileLink:   d.FileLink,
				FileName:   d.FileName,
				FilePath:   d.FilePath,
				Active:     active,
				ApprovalBy: d.ApprovedBy,
				RejectNote: d.Note,
			}
			if d.ApprovedAt != nil {
				dn.ApprovedAt = *d.ApprovedAt
			}
			switch d.Status {
			case "3":
				dn.ApprovalType = "Rejected"
			case "2":
				dn.ApprovalType = "Approved"
			default:
				dn.ApprovalType = "Create"
			}
			dataResult = append(dataResult, dn)
		}
	}

	// response
	response := new(ResponseRetrieveActiveMerchant)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistMerchantActiveDataContent(ctx context.Context, request *RequestRetrieveActiveMerchantContent) (*ResponseRetrieveActiveMerchantContent, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTActiveDataContent"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTActiveDataContent */
	dataResult := []ResponseRetrieveActiveMerchantContentData{}
	queryData := map[string]interface{}{}
	data, meta, err := b.rprBlacklistMerchant.ReadRowsBlacklistMerchantApprovedQuery(ctx, queryData, request.Page, request.Limit)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dn := ResponseRetrieveActiveMerchantContentData{
				Id:           d.Id,
				NMID:         d.NMID,
				MerchantName: d.MerchantName,
				Datasource:   d.Datasource,
				FileId:       d.FileId,
				FileLink:     d.FileLink,
				CreatedBy:    d.CreatedBy,
				CreatedAt:    d.CreatedAt,
				UpdatedBy:    d.UpdatedBy,
				UpdatedAt:    d.UpdatedAt,
				ApprovedBy:   d.ApprovedBy,
			}
			if d.ApprovedAt != nil {
				dn.ApprovedAt = *d.ApprovedAt
			}
			dataResult = append(dataResult, dn)
		}
	}

	// response
	response := new(ResponseRetrieveActiveMerchantContent)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistMerchantActiveDataContentById(ctx context.Context, request *RequestRetrieveActiveMerchantContentById) (*ResponseRetrieveActiveMerchantContentById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantActiveDataContentById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantActiveDataContentById */
	queryData := map[string]interface{}{}
	queryData["id"] = request.Id
	data, err := b.rprBlacklistMerchant.ReadRowBlacklistMerchantApprovedSingle(ctx, queryData)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", "data not found")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	response := new(ResponseRetrieveActiveMerchantContentById)
	response.Id = data.Id
	response.NMID = data.NMID
	response.MerchantName = data.MerchantName
	response.Datasource = data.Datasource
	response.FileId = data.FileId
	response.FileLink = data.FileLink
	response.CreatedBy = data.CreatedBy
	response.CreatedAt = data.CreatedAt
	response.UpdatedBy = data.UpdatedBy
	response.UpdatedAt = data.UpdatedAt
	response.ApprovedBy = data.ApprovedBy
	if data.ApprovedAt != nil {
		response.ApprovedAt = *data.ApprovedAt
	}
	return response, nil
}

func (b *blueprint) BlacklistMerchantApprovedDownloadContentByFileId(ctx context.Context, request *RequestMerchantDownloadContentByFileId) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantApprovedDownloadContentByFileId"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantApprovedDownloadContentByFileId */
	var dataResult []csvProcessor.MerchantCSVFormat
	data, err := b.rprBlacklistMerchant.ReadRowsBlacklistMerchantApproved(ctx, map[string]interface{}{
		"file_id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dataResult = append(dataResult, csvProcessor.MerchantCSVFormat{
				NMID:       d.NMID,
				Name:       d.MerchantName,
				DataSource: d.Datasource,
			})
		}
	}
	csvFileContent, err := csvProcessor.ConvertMerchantToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}

func (b *blueprint) BlacklistMerchantApprovalDownloadContentByFileId(ctx context.Context, request *RequestMerchantDownloadContentByFileId) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantApprovalDownloadContentByFileId"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantApprovalDownloadContentByFileId */
	var dataResult []csvProcessor.MerchantCSVFormat
	data, err := b.rprBlacklistMerchant.ReadRowsBlacklistMerchantApproval(ctx, map[string]interface{}{
		"file_id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dataResult = append(dataResult, csvProcessor.MerchantCSVFormat{
				NMID:       d.NMID,
				Name:       d.MerchantName,
				DataSource: d.Datasource,
			})
		}
	}
	csvFileContent, err := csvProcessor.ConvertMerchantToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}

func (b *blueprint) BlacklistMerchantFileById(ctx context.Context, request *RequestBlacklistMerchantFileById) (*ResponseBlacklistMerchantFileById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistMerchantFileById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistMerchantFileById */
	queryData := map[string]interface{}{}
	queryData["id"] = request.Id
	data, err := b.rprBlacklistMerchant.ReadRowBlacklistMerchantFileSingle(ctx, queryData)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", "data not found")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseBlacklistMerchantFileById)
	response.ID = data.ID
	response.FileLink = data.FileLink
	response.FileName = data.FileName
	response.FilePath = data.FilePath
	response.ApprovalType = data.ApprovalType
	response.Note = data.Note
	response.Status = data.Status
	response.Active = data.Active
	response.CreatedBy = data.CreatedBy
	response.CreatedAt = data.CreatedAt
	response.UpdatedBy = data.UpdatedBy
	response.UpdatedAt = data.UpdatedAt
	response.ApprovedBy = data.ApprovedBy
	if data.ApprovedAt != nil {
		response.ApprovedAt = *data.ApprovedAt
	}
	status, _ := strconv.Atoi(data.Status)
	response.Status = strconv.Itoa(status - 1)
	switch status {
	case 3:
		response.ApprovalType = "Rejected"
	case 2:
		response.ApprovalType = "Approved"
	case 1:
		response.ApprovalType = "Create"
	}
	return response, nil
}

func (b *blueprint) BlacklistMerchantFileContentById(ctx context.Context, request *RequestBlacklistMerchantFileContent) (*ResponseBlacklistMerchantFileContent, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTFileContentById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTFileContentById */
	dataResult := []ResponseBlacklistMerchantFileContentData{}
	queryData := map[string]interface{}{
		"file_id": request.Id,
	}
	data, meta, err := b.rprBlacklistMerchant.ReadRowsBlacklistMerchantApprovalQuery(ctx, queryData, request.Page, request.Limit)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dn := ResponseBlacklistMerchantFileContentData{
				Id:           d.BlacklistMerchantId,
				NMID:         d.NMID,
				MerchantName: d.MerchantName,
				Datasource:   d.Datasource,
				FileId:       d.FileId,
				FileLink:     d.FileLink,
				CreatedBy:    d.CreatedBy,
				CreatedAt:    d.CreatedAt,
				UpdatedBy:    d.UpdatedBy,
				UpdatedAt:    d.UpdatedAt,
				ApprovedBy:   d.ApprovedBy,
			}
			if d.ApprovedAt != nil {
				dn.ApprovedAt = *d.ApprovedAt
			}
			dataResult = append(dataResult, dn)
		}
	}

	// response
	response := new(ResponseBlacklistMerchantFileContent)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}
