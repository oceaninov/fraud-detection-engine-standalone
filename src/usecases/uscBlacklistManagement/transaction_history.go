package uscBlacklistManagement

import (
	"context"
	"fmt"

	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/pkg/humanTime"
)

type (
	ResponseRetrieveTransactionHistoryData struct {
		Id                string `json:"id"`
		TransactionId     string `json:"transactionId"`
		TransactionTypeId string `json:"transactionTypeId"`
		Rules             string `json:"rules"`
		Title             string `json:"title"`
		Channel           string `json:"channel"`
		BodyReq           string `json:"bodyReq"`
		FlagId            string `json:"flagId"`
		Flag              string `json:"flag"`
		CreatedAt         string `json:"createdAt"`
		UserId            string `json:"userId"`
		Amount            string `json:"amount"`
	}
	ResponseRetrieveTransactionHistory struct {
		ResponseCode    bool                                     `json:"success"`
		ResponseMessage string                                   `json:"messages"`
		Data            []ResponseRetrieveTransactionHistoryData `json:"transactionReport"`
		Meta            basicObject.Meta                         `json:"meta"`
	}
	RequestRetrieveTransactionHistory struct {
		Page      int    `query:"page"`
		Limit     int    `query:"limit"`
		StartDate string `query:"start_date"`
		EndDate   string `query:"end_date"`
		Search    string `query:"transaction_type"`
	}
)

func (b *blueprint) TransactionHistories(ctx context.Context, request *RequestRetrieveTransactionHistory) (*ResponseRetrieveTransactionHistory, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT016"
	const fName = "usecases.uscBlacklistManagement.TransactionHistory"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_BlacklistHistories */
	dataResult := []ResponseRetrieveTransactionHistoryData{}
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
	data, meta, err := b.rprTransaction.ReadRowsTransactionHistory(ctx, map[string]interface{}{}, request.Page, request.Limit, search, startDate, endDate)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if data != nil {
		for _, d := range *data {
			flag := ""
			row, _ := b.rprTransaction.ReadRowFlag(ctx, map[string]interface{}{"id": d.FlagId})
			if row != nil {
				flag = row.Title
			}
			result := ResponseRetrieveTransactionHistoryData{
				Id:                d.Id,
				TransactionId:     d.TransactionId,
				TransactionTypeId: d.TransactionType,
				Rules:             d.Rules,
				Title:             d.Title,
				Channel:           d.Channel,
				BodyReq:           *d.BodyReq,
				FlagId:            d.FlagId,
				Flag:              flag,
				CreatedAt:         d.CreatedAt,
				UserId:            d.UserId,
				Amount:            d.Amount,
			}
			createdAt, _ := humanTime.FormatDate(d.CreatedAt)
			result.CreatedAt = createdAt
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseRetrieveTransactionHistory)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}
