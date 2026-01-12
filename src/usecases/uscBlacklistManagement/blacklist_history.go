package uscBlacklistManagement

import (
	"context"
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/csvProcessor"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/pkg/humanTime"
	"strings"
)

type (
	ResponseRetrieveBlacklistHistoryData struct {
		Id               string    `json:"id"`
		PhoneNumber      string    `json:"phoneNumber"`
		BeneficiaryName  string    `json:"beneficiaryName"`
		TransactionTypes []TrxType `json:"transactionTypes"`
		Event            string    `json:"event"`
		CreatedAt        string    `json:"createdAt"`
		CreatedBy        string    `json:"createdBy"`
	}
	ResponseRetrieveBlacklistHistory struct {
		ResponseCode    bool                                   `json:"success"`
		ResponseMessage string                                 `json:"messages"`
		Data            []ResponseRetrieveBlacklistHistoryData `json:"blackListHistorys"`
		Meta            basicObject.Meta                       `json:"meta"`
	}
	RequestRetrieveBlacklistHistory struct {
		Page      int    `query:"page"`
		Limit     int    `query:"limit"`
		StartDate string `query:"start_date"`
		EndDate   string `query:"end_date"`
		Search    string `query:"phone_number"`
	}
)

type (
	RequestRetrieveBlacklistReporting struct {
		Page      int    `query:"page"`
		Limit     int    `query:"limit"`
		StartDate string `query:"date_from"`
		EndDate   string `query:"date_to"`
		Search    string `query:"phone_number"`
	}
)

func (b *blueprint) BlacklistHistories(ctx context.Context, request *RequestRetrieveBlacklistHistory) (*ResponseRetrieveBlacklistHistory, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT016"
	const fName = "usecases.uscBlacklistManagement.BlacklistHistories"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistHistories */
	dataResult := []ResponseRetrieveBlacklistHistoryData{}
	startDate := ""
	if request.StartDate != "" && request.StartDate != "230" {
		startDate = request.StartDate
	}
	endDate := ""
	if request.EndDate != "" && request.EndDate != "230" {
		endDate = request.EndDate
	}
	search := ""
	if request.Search != "" && request.Search != "230" {
		search = request.Search
	}
	data, meta, err := b.rprBlacklistHistory.ReadRowsBlacklistHistory(ctx, map[string]interface{}{}, request.Page, request.Limit, search, startDate, endDate)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := ResponseRetrieveBlacklistHistoryData{
				Id:              d.Id,
				PhoneNumber:     d.PhoneNumber,
				BeneficiaryName: d.BeneficiaryName,
				Event:           d.Event,
				CreatedAt:       d.CreatedAt,
				CreatedBy:       d.CreatedBy,
			}

			createdAt, _ := humanTime.FormatDate(d.CreatedAt)
			result.CreatedAt = createdAt

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
			result.TransactionTypes = trxTypeData
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseRetrieveBlacklistHistory)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) BlacklistReporting(ctx context.Context, request *RequestRetrieveBlacklistReporting) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT016"
	const fName = "usecases.uscBlacklistManagement.BlacklistReporting"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistReporting */
	var dataResult []csvProcessor.BlacklistCSVFormat
	startDate := ""
	if request.StartDate != "" && request.StartDate != "230" {
		startDate = request.StartDate
	}
	endDate := ""
	if request.EndDate != "" && request.EndDate != "230" {
		endDate = request.EndDate
	}
	search := ""
	if request.Search != "" && request.Search != "230" {
		search = request.Search
	}
	data, _, err := b.rprBlacklistHistory.ReadRowsBlacklistHistoryWithoutPagination(ctx, map[string]interface{}{}, search, startDate, endDate)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			result := csvProcessor.BlacklistCSVFormat{
				Id:              d.Id,
				PhoneNumber:     d.PhoneNumber,
				BeneficiaryName: d.BeneficiaryName,
				Event:           d.Event,
				CreatedAt:       d.CreatedAt,
				CreatedBy:       d.CreatedBy,
			}
			createdAt, _ := humanTime.FormatDate(d.CreatedAt)
			result.CreatedAt = createdAt
			result.TransactionTypes = d.TransactionTypes
			dataResult = append(dataResult, result)
		}
	}
	csvFileContent, err := csvProcessor.ConvertBlacklistToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}
