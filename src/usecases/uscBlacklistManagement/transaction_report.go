package uscBlacklistManagement

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.com/fds22/detection-sys/pkg/csvProcessor"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
)

type (
	transactionType struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	interval struct {
		Value int    `json:"value"`
		Type  string `json:"type"`
	}
	actions struct {
		IsReport bool `json:"is_report"`
		IsBlock  bool `json:"is_block"`
	}
	rule struct {
		Id              string `json:"Id"`
		RuleName        string `json:"RuleName"`
		Types           string `json:"Types"`
		TransactionType string `json:"TransactionType"`
		Interval        string `json:"Interval"`
		Amount          int    `json:"Amount"`
		Actions         string `json:"Actions"`
		CreatedAt       string `json:"CreatedAt"`
		UpdatedAt       string `json:"UpdatedAt"`
		Status          int    `json:"Status"`
		TimeRangeType   string `json:"TimeRangeType"`
		StartTimeRange  string `json:"StartTimeRange"`
		EndTimeRange    string `json:"EndTimeRange"`
		Sofs            string `json:"Sofs"`
	}
)

type (
	RequestRetrieveTransactionReporting struct {
		Page      int    `query:"page"`
		Limit     int    `query:"limit"`
		StartDate string `query:"date_from"`
		EndDate   string `query:"date_to"`
		Search    string `query:"transaction_type"`
	}
)

func (b *blueprint) TransactionReporting(ctx context.Context, request *RequestRetrieveTransactionReporting) ([]byte, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "MGT017"
	const fName = "usecases.uscBlacklistManagement.TransactionReporting"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_TransactionReporting */
	var dataResult []csvProcessor.TransactionCSVFormat
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
	data, _, err := b.rprTransaction.ReadRowsTransactionHistoryWithoutPagination(ctx, map[string]interface{}{}, search, startDate, endDate)
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
			var result csvProcessor.TransactionCSVFormat
			result.Id = d.Id
			result.TransactionId = d.TransactionId
			result.TransactionType = d.TransactionType
			result.Title = d.Title
			result.Channel = d.Channel
			result.Flag = flag
			result.UserId = d.UserId
			result.Amount = d.Amount
			result.DestinationId = d.DestinationId
			result.CreatedAt = d.CreatedAt
			if d.BodyReq == nil {
				result.BodyReq = "N/A"
			}
			if *d.BodyReq == "" || *d.BodyReq == "null" {
				result.BodyReq = "N/A"
			} else {
				result.BodyReq = *d.BodyReq
			}
			var ruleData []rule
			err := json.Unmarshal([]byte(d.Rules), &ruleData)
			if err != nil {
				result.RuleName = "N/A"
				result.RuleAmount = "N/A"
				result.RuleInterval = "N/A"
				result.RuleTransactionType = "N/A"
				result.RuleAction = "N/A"
			} else {
				var t transactionType
				var a actions
				var i interval

				if ruleData == nil {
					result.RuleName = "N/A"
					result.RuleAmount = "N/A"
					result.RuleInterval = "N/A"
					result.RuleTransactionType = "N/A"
					result.RuleAction = "N/A"
				}

				if len(ruleData) > 0 {
					errT := json.Unmarshal([]byte(ruleData[0].TransactionType), &t)
					if errT != nil {
						result.RuleTransactionType = "N/A"
					} else {
						result.RuleTransactionType = t.Type
					}

					errI := json.Unmarshal([]byte(ruleData[0].Interval), &i)
					if errI != nil {
						result.RuleInterval = "N/A"
					} else {
						result.RuleInterval = i.Type
					}

					errA := json.Unmarshal([]byte(ruleData[0].Actions), &a)
					if errA != nil {
						result.RuleAction = "N/A"
					} else {
						result.RuleAction = fmt.Sprintf("reported: %v, blocked: %v", a.IsReport, a.IsBlock)
					}

					result.RuleAmount = strconv.Itoa(ruleData[0].Amount)
					result.RuleName = ruleData[0].RuleName
				}
			}
			dataResult = append(dataResult, result)
		}
	}
	csvFileContent, err := csvProcessor.ConvertTransactionToCSV(dataResult)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	return []byte(csvFileContent), nil
}
