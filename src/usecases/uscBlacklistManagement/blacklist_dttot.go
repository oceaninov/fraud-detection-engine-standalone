package uscBlacklistManagement

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/csvProcessor"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistDTTOT"
	"strconv"
	"strings"
	"time"
)

type (
	ResponseUploadBlacklistDTTOT struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
		FileId          string `json:"fileId"`
	}
	RequestUploadBlacklistDTTOT struct{}
)

type (
	ResponseRejectBlacklistDTTOT struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestRejectBlacklistDTTOT struct {
		FileId      string `param:"fileId"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"creator"`
	}
)

type (
	RequestApproveBlacklistDTTOT struct {
		FileId      string `param:"fileId"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"creator"`
	}
	ResponseApproveBlacklistDTTOT struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
)

type (
	ResponseRetrieveBlacklistDTTOTApprovalData struct {
		ID           string `json:"ID"`
		FileLink     string `json:"FileLink"`
		FileName     string `json:"FileName"`
		FilePath     string `json:"FilePath"`
		ApprovalType string `json:"ApprovalType"`
		Note         string `json:"Note"`
		Status       string `json:"Status"`
		Active       string `json:"Active"`
		CreatedBy    string `json:"CreatedBy"`
		CreatedAt    string `json:"CreatedAt"`
		UpdatedBy    string `json:"UpdatedBy"`
		UpdatedAt    string `json:"UpdatedAt"`
		ApprovedBy   string `json:"ApprovedBy"`
		ApprovedAt   string `json:"ApprovedAt"`
	}
	ResponseRetrieveBlacklistDTTOTApproval struct {
		ResponseCode    string                                       `json:"responseCode"`
		ResponseMessage string                                       `json:"responseMessage"`
		Data            []ResponseRetrieveBlacklistDTTOTApprovalData `json:"data"`
		Meta            basicObject.Meta                             `json:"meta"`
	}
	RequestRetrieveBlacklistDTTOTApproval struct {
		StatusApproval int `json:"statusApproval"`
		Page           int `json:"page"`
		Limit          int `json:"limit"`
	}
)

type (
	RequestRetrieveBlacklistDTTOTApproved struct{}
)

type (
	RequestRetrieveBlacklistDTTOTTemplate struct{}
)

type (
	ResponseRetrieveActiveDTTOTData struct {
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
	ResponseRetrieveActiveDTTOT struct {
		ResponseCode    bool                              `json:"success"`
		ResponseMessage string                            `json:"messages"`
		Data            []ResponseRetrieveActiveDTTOTData `json:"blacklistApproval"`
		Meta            basicObject.Meta                  `json:"meta"`
	}
	RequestRetrieveActiveDTTOT struct {
		Page   int    `query:"page"`
		Limit  int    `query:"limit"`
		Search string `query:"search"`
	}
)

type (
	ResponseRetrieveActiveDTTOTContentData struct {
		Id         string `json:"id"`
		PPATKId    string `json:"ppatkId"`
		Name       string `json:"name"`
		Bod        string `json:"bod"`
		Datasource string `json:"datasource"`
		FileId     string `json:"fileId"`
		FileLink   string `json:"fileLink"`
		NIK        string `json:"nik"`
		CreatedBy  string `json:"createdBy"`
		CreatedAt  string `json:"createdAt"`
		UpdatedBy  string `json:"updatedBy"`
		UpdatedAt  string `json:"updatedAt"`
		ApprovedBy string `json:"approvedBy"`
		ApprovedAt string `json:"approvedAt"`
	}
	ResponseRetrieveActiveDTTOTContent struct {
		ResponseCode    bool                                     `json:"success"`
		ResponseMessage string                                   `json:"messages"`
		Data            []ResponseRetrieveActiveDTTOTContentData `json:"blacklistDttot"`
		Meta            basicObject.Meta                         `json:"meta"`
	}
	RequestRetrieveActiveDTTOTContent struct {
		Page   int    `query:"page"`
		Limit  int    `query:"limit"`
		Search string `query:"search"`
	}
)

type (
	ResponseRetrieveActiveDTTOTContentById struct {
		Id         string `json:"id"`
		PPATKId    string `json:"ppatkId"`
		Name       string `json:"name"`
		Bod        string `json:"bod"`
		Datasource string `json:"datasource"`
		FileId     string `json:"fileId"`
		FileLink   string `json:"fileLink"`
		NIK        string `json:"nik"`
		CreatedBy  string `json:"createdBy"`
		CreatedAt  string `json:"createdAt"`
		UpdatedBy  string `json:"updatedBy"`
		UpdatedAt  string `json:"updatedAt"`
		ApprovedBy string `json:"approvedBy"`
		ApprovedAt string `json:"approvedAt"`
	}
	RequestRetrieveActiveDTTOTContentById struct {
		Id string `param:"id"`
	}
)

type (
	RequestDTTOTDownloadContentByFileId struct {
		Id string `param:"id"`
	}
)

type (
	ResponseBlacklistDTTOTFileById struct {
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
	RequestBlacklistDTTOTFileById struct {
		Id string `param:"id"`
	}
)

type (
	ResponseBlacklistDTTOTFileContentData struct {
		Id         string `json:"id"`
		PPATKId    string `json:"ppatkId"`
		Name       string `json:"name"`
		Bod        string `json:"bod"`
		Datasource string `json:"datasource"`
		FileId     string `json:"fileId"`
		FileLink   string `json:"fileLink"`
		NIK        string `json:"nik"`
		CreatedBy  string `json:"createdBy"`
		CreatedAt  string `json:"createdAt"`
		UpdatedBy  string `json:"updatedBy"`
		UpdatedAt  string `json:"updatedAt"`
		ApprovedBy string `json:"approvedBy"`
		ApprovedAt string `json:"approvedAt"`
	}
	ResponseBlacklistDTTOTFileContent struct {
		ResponseCode    bool                                    `json:"success"`
		ResponseMessage string                                  `json:"messages"`
		Data            []ResponseBlacklistDTTOTFileContentData `json:"blacklistDttot"`
		Meta            basicObject.Meta                        `json:"meta"`
	}
	RequestBlacklistDTTOTFileContent struct {
		Id    string `param:"id"`
		Page  int    `query:"page"`
		Limit int    `query:"limit"`
	}
)

func (b *blueprint) BlacklistDTTOTUpload(ctx context.Context, request *RequestUploadBlacklistDTTOT) (*ResponseUploadBlacklistDTTOT, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT013"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTUpload"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTUpload */
	fileBytes, ok := ctx.Value("X-File-Bytes").([]byte)
	if !ok {
		errMsg := fmt.Errorf("no file attached")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	resultData, err := csvProcessor.ReadCSVForDTTOT(fileBytes)
	if err != nil {
		errMsg := fmt.Errorf("csv file processor failed")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// inserting data into database
	fileId := guuid.NewString()
	err = b.rprBlacklistDTTOT.WriteBulkRowBlacklistDTTOTApproval(ctx, fileId, email, resultData)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// inserting file data
	fileName := strings.Replace(fileId, "-", "", -1)
	fileEntity := new(rprBlacklistDTTOT.EntityBlacklistDTTOTFile)
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
	fileData, fileDataId, err := b.rprBlacklistDTTOT.WriteRowBlacklistDTTOTFile(ctx, *fileEntity)
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
	response := new(ResponseUploadBlacklistDTTOT)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.FileId = fileData.ID
	return response, nil
}

func (b *blueprint) BlacklistDTTOTReject(ctx context.Context, request *RequestRejectBlacklistDTTOT) (*ResponseRejectBlacklistDTTOT, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT014"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTReject"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTReject */
	currentApprovalEntity, err := b.rprBlacklistDTTOT.ReadRowBlacklistDTTOTFile(ctx, map[string]interface{}{
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

	//err = b.rprBlacklistDTTOT.RemoveBlacklistDTTOTApprovalData(ctx, currentApprovalEntity.ID)
	//if err != nil {
	//	errMsg := fmt.Errorf("internal server error")
	//	logging.Errorw(fName, "reason", err.Error())
	//	return nil, errWrap.WrapRepositoryError(errMsg)
	//}

	err = b.rprBlacklistDTTOT.WriteRejectForBlacklistDTTOTApproval(ctx, request.FileId, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseRejectBlacklistDTTOT)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Reject Blacklist DTTOT"
	return response, nil
}

func (b *blueprint) BlacklistDTTOTApprove(ctx context.Context, request *RequestApproveBlacklistDTTOT) (*ResponseApproveBlacklistDTTOT, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT015"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTApprove"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTApprovalGet */
	currentApprovalEntity, err := b.rprBlacklistDTTOT.ReadRowBlacklistDTTOTFile(ctx, map[string]interface{}{
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
		data, err := b.rprBlacklistDTTOT.ReadRowsBlacklistDTTOTApproval(ctx, map[string]interface{}{
			"file_id": currentApprovalEntity.ID,
		})
		logging.Errorw(fName, "DATA", data)

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
		err = b.rprBlacklistDTTOT.WriteBulkRowBlacklistDTTOT(ctx, *data)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		//err = b.rprBlacklistDTTOT.RemoveBlacklistDTTOTApprovalData(ctx, currentApprovalEntity.ID)
		//if err != nil {
		//	errMsg := fmt.Errorf("internal server error")
		//	logging.Errorw(fName, "reason", err.Error())
		//	return nil, errWrap.WrapRepositoryError(errMsg)
		//}
		err = b.rprBlacklistDTTOT.RemoveBlacklistDTTOTData(ctx, currentApprovalEntity.ID)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
		err = b.rprBlacklistDTTOT.WriteApprovalForBlacklistDTTOTApproval(ctx, currentApprovalEntity.ID, request.Note, email)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", errMsg.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	}

	// response
	response := new(ResponseApproveBlacklistDTTOT)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Approve Blacklist DTTOT"
	return response, nil
}

func (b *blueprint) BlacklistDTTOTApprovalGet(ctx context.Context, request *RequestRetrieveBlacklistDTTOTApproval) (*ResponseRetrieveBlacklistDTTOTApproval, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT016"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTApprovalGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTApprovalGet */
	dataResult := []ResponseRetrieveBlacklistDTTOTApprovalData{}
	data, meta, err := b.rprBlacklistDTTOT.ReadRowsBlacklistDTTOTFile(ctx, map[string]interface{}{
		"status": request.StatusApproval,
	}, request.Page, request.Limit)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseRetrieveBlacklistDTTOTApprovalData{
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
	response := new(ResponseRetrieveBlacklistDTTOTApproval)
	response.ResponseCode = basicObject.SuccessfullyCode
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistDTTOTApprovedGet(ctx context.Context, request *RequestRetrieveBlacklistDTTOTApproved) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTApprovedGet"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTApprovedGet */
	var dataResult []csvProcessor.DTTOTCSVFormat
	data, err := b.rprBlacklistDTTOT.ReadRowsBlacklistDTTOTApproved(ctx, map[string]interface{}{})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dataResult = append(dataResult, csvProcessor.DTTOTCSVFormat{
				PPATKID:     d.PPATKId,
				Name:        d.Name,
				NIK:         d.NIK,
				DateOfBirth: d.Bod,
				DataSource:  d.Datasource,
			})
		}
	}
	csvFileContent, err := csvProcessor.ConvertDTTOTToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}

func (b *blueprint) BlacklistDTTOTApprovedTemplate(ctx context.Context, request *RequestRetrieveBlacklistDTTOTTemplate) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTApprovedTemplate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTApprovedTemplate */
	var dataResult []csvProcessor.DTTOTCSVFormat
	csvFileContent, err := csvProcessor.ConvertDTTOTToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}

func (b *blueprint) BlacklistDTTOTActiveData(ctx context.Context, request *RequestRetrieveActiveDTTOT) (*ResponseRetrieveActiveDTTOT, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTActiveData"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTActiveData */
	dataResult := []ResponseRetrieveActiveDTTOTData{}
	queryData := map[string]interface{}{}
	data, meta, err := b.rprBlacklistDTTOT.ReadRowsBlacklistDTTOTFile(ctx, queryData, request.Page, request.Limit)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			active, _ := strconv.ParseBool(d.Active)
			dn := ResponseRetrieveActiveDTTOTData{
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
	response := new(ResponseRetrieveActiveDTTOT)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistDTTOTActiveDataContent(ctx context.Context, request *RequestRetrieveActiveDTTOTContent) (*ResponseRetrieveActiveDTTOTContent, error) {
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
	dataResult := []ResponseRetrieveActiveDTTOTContentData{}
	queryData := map[string]interface{}{}
	data, meta, err := b.rprBlacklistDTTOT.ReadRowsBlacklistDTTOTApprovedQuery(ctx, queryData, request.Page, request.Limit)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dn := ResponseRetrieveActiveDTTOTContentData{
				Id:         d.Id,
				PPATKId:    d.PPATKId,
				Name:       d.Name,
				Bod:        d.Bod,
				Datasource: d.Datasource,
				FileId:     d.FileId,
				FileLink:   d.FileLink,
				NIK:        d.NIK,
				CreatedBy:  d.CreatedBy,
				CreatedAt:  d.CreatedAt,
				UpdatedBy:  d.UpdatedBy,
				UpdatedAt:  d.UpdatedAt,
				ApprovedBy: d.ApprovedBy,
			}
			if d.ApprovedAt != nil {
				dn.ApprovedAt = *d.ApprovedAt
			}
			dataResult = append(dataResult, dn)
		}
	}

	// response
	response := new(ResponseRetrieveActiveDTTOTContent)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistDTTOTActiveDataContentById(ctx context.Context, request *RequestRetrieveActiveDTTOTContentById) (*ResponseRetrieveActiveDTTOTContentById, error) {
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
	queryData := map[string]interface{}{}
	queryData["id"] = request.Id
	data, err := b.rprBlacklistDTTOT.ReadRowBlacklistDTTOTApprovedSingle(ctx, queryData)
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
	response := new(ResponseRetrieveActiveDTTOTContentById)
	response.Id = data.Id
	response.PPATKId = data.PPATKId
	response.Name = data.Name
	response.Bod = data.Bod
	response.Datasource = data.Datasource
	response.FileId = data.FileId
	response.FileLink = data.FileLink
	response.NIK = data.NIK
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

func (b *blueprint) BlacklistDTTOTApprovedDownloadContentByFileId(ctx context.Context, request *RequestDTTOTDownloadContentByFileId) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTApprovedDownloadContentByFileId"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTApprovedDownloadContentByFileId */
	var dataResult []csvProcessor.DTTOTCSVFormat
	data, err := b.rprBlacklistDTTOT.ReadRowsBlacklistDTTOTApproved(ctx, map[string]interface{}{
		"file_id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dataResult = append(dataResult, csvProcessor.DTTOTCSVFormat{
				PPATKID:     d.PPATKId,
				Name:        d.Name,
				NIK:         d.NIK,
				DateOfBirth: d.Bod,
				DataSource:  d.Datasource,
			})
		}
	}
	csvFileContent, err := csvProcessor.ConvertDTTOTToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}

func (b *blueprint) BlacklistDTTOTApprovalDownloadContentByFileId(ctx context.Context, request *RequestDTTOTDownloadContentByFileId) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTApprovalDownloadContentByFileId"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTApprovalDownloadContentByFileId */
	var dataResult []csvProcessor.DTTOTCSVFormat
	data, err := b.rprBlacklistDTTOT.ReadRowsBlacklistDTTOTApproval(ctx, map[string]interface{}{
		"file_id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dataResult = append(dataResult, csvProcessor.DTTOTCSVFormat{
				PPATKID:     d.PPATKId,
				Name:        d.Name,
				NIK:         d.NIK,
				DateOfBirth: d.Bod,
				DataSource:  d.Datasource,
			})
		}
	}
	csvFileContent, err := csvProcessor.ConvertDTTOTToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}

func (b *blueprint) BlacklistDTTOTFileById(ctx context.Context, request *RequestBlacklistDTTOTFileById) (*ResponseBlacklistDTTOTFileById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.BlacklistDTTOTFileById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistDTTOTFileById */
	queryData := map[string]interface{}{}
	queryData["id"] = request.Id
	data, err := b.rprBlacklistDTTOT.ReadRowBlacklistDTTOTFileSingle(ctx, queryData)
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
	response := new(ResponseBlacklistDTTOTFileById)
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

func (b *blueprint) BlacklistDTTOTFileContentById(ctx context.Context, request *RequestBlacklistDTTOTFileContent) (*ResponseBlacklistDTTOTFileContent, error) {
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
	dataResult := []ResponseBlacklistDTTOTFileContentData{}
	queryData := map[string]interface{}{
		"file_id": request.Id,
	}
	data, meta, err := b.rprBlacklistDTTOT.ReadRowsBlacklistDTTOTApprovalQuery(ctx, queryData, request.Page, request.Limit)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			dn := ResponseBlacklistDTTOTFileContentData{
				Id:         d.BlacklistDTTOTId,
				PPATKId:    d.PPATKId,
				Name:       d.Name,
				Bod:        d.Bod,
				Datasource: d.Datasource,
				FileId:     d.FileId,
				FileLink:   d.FileLink,
				NIK:        d.NIK,
				CreatedBy:  d.CreatedBy,
				CreatedAt:  d.CreatedAt,
				UpdatedBy:  d.UpdatedBy,
				UpdatedAt:  d.UpdatedAt,
				ApprovedBy: d.ApprovedBy,
			}
			if d.ApprovedAt != nil {
				dn.ApprovedAt = *d.ApprovedAt
			}
			dataResult = append(dataResult, dn)
		}
	}

	// response
	response := new(ResponseBlacklistDTTOTFileContent)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}
