package uscRuleManagement

import (
	"context"
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/src/repositories/rprRuleDetection"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SOF struct {
	Id        string `json:"id"`
	SofName   string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type (
	ResponseRuleDetectionCreate struct {
		ResponseCode    bool   `json:"code"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestRuleDetectionCreate struct {
		RuleName        string `json:"ruleName"`
		Type            string `json:"type"`
		TransactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"transactionType"`
		Interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"interval"`
		Amount int `json:"amount"`
		Action struct {
			IsReport bool `json:"is_report"`
			IsBlock  bool `json:"is_block"`
		} `json:"action"`
		Channel        string   `json:"channel"`
		TimeRangeType  string   `json:"time_range_type"`
		StartTimeRange string   `json:"start_time_range"`
		EndTimeRange   string   `json:"end_time_range"`
		Sofs           []string `json:"sofs"`
	}
)

type (
	ResponseRuleDetectionUpdate struct {
		ResponseCode    bool   `json:"code"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestRuleDetectionUpdate struct {
		Id              string `param:"id"`
		RuleName        string `json:"ruleName"`
		Type            string `json:"type"`
		TransactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"transactionType"`
		Interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"interval"`
		Amount int `json:"amount"`
		Action struct {
			IsReport bool `json:"is_report"`
			IsBlock  bool `json:"is_block"`
		} `json:"action"`
		Channel        string   `json:"channel"`
		TimeRangeType  string   `json:"time_range_type"`
		StartTimeRange string   `json:"start_time_range"`
		EndTimeRange   string   `json:"end_time_range"`
		Sofs           []string `json:"sofs"`
	}
)

type (
	ResponseRuleDisabled struct {
		ResponseCode    bool   `json:"code"`
		ResponseMessage string `json:"message"`
		Id              string `json:"id"`
	}
	RequestRuleDisabled struct {
		Id string `param:"id"`
	}
)

type (
	ResponseRuleEnabled struct {
		ResponseCode    bool   `json:"code"`
		ResponseMessage string `json:"message"`
		Id              string `json:"id"`
	}
	RequestRuleEnabled struct {
		Id string `param:"id"`
	}
)

type (
	ResponseRuleDetectionDelete struct {
		ResponseCode    bool   `json:"code"`
		ResponseMessage string `json:"messages"`
		Id              string `json:"id"`
	}
	RequestRuleDetectionDelete struct {
		Id string `param:"id"`
	}
)

type (
	ResponseRuleDetectionApprove struct {
		ResponseCode    bool   `json:"code"`
		ResponseMessage string `json:"messages"`
	}
	RequestRuleDetectionApprove struct {
		Id          string `param:"id"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"actionTaker"`
	}
)

type (
	ResponseRuleDetectionReject struct {
		ResponseCode    bool   `json:"code"`
		ResponseMessage string `json:"messages"`
	}
	RequestRuleDetectionReject struct {
		Id          string `param:"id"`
		Note        string `json:"reject_note"`
		ActionTaker string `json:"actionTaker"`
	}
)

type (
	ResponseRuleDetectionGetApprovalData struct {
		Id              string `json:"id"`
		RuleName        string `json:"ruleName"`
		Types           string `json:"type"`
		TransactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"transactionType"`
		Interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"interval"`
		Action struct {
			IsReport bool `json:"isReport"`
			IsBlock  bool `json:"isBlock"`
		} `json:"action"`
		Note           string `json:"rejectNote"`
		Amount         string `json:"amount"`
		Status         string `json:"status"`
		TimeRangeType  string `json:"timeRangeType"`
		StartTimeRange string `json:"startTimeRange"`
		EndTimeRange   string `json:"endTimeRange"`
		Channel        string `json:"channel"`
		RuleId         string `json:"ruleId"`
		ApprovalType   string `json:"approvalType"`
		CreatedBy      string `json:"createdBy"`
		CreatedAt      string `json:"createdAt"`
		UpdatedBy      string `json:"updatedBy"`
		UpdatedAt      string `json:"updatedAt"`
		ApprovedBy     string `json:"approvedBy"`
		ApprovedAt     string `json:"approvedAt"`
		Sofs           []SOF  `json:"sofs"`
	}
	ResponseRuleDetectionGetApproval struct {
		ResponseCode    bool                                   `json:"code"`
		ResponseMessage string                                 `json:"messages"`
		Data            []ResponseRuleDetectionGetApprovalData `json:"ruleApproval"`
		Meta            basicObject.Meta                       `json:"meta"`
	}
	RequestRuleDetectionGetApproval struct {
		StatusApproval int `query:"statusApproval"`
		Page           int `query:"page"`
		Limit          int `query:"limit"`
	}
)

type (
	ResponseRuleDetectionGetApprovedData struct {
		Id              string `json:"id"`
		RuleName        string `json:"ruleName"`
		TransactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"transactionType"`
		Interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"interval"`
		Action struct {
			IsReport bool `json:"isReport"`
			IsBlock  bool `json:"isBlock"`
		} `json:"action"`
		Types          string `json:"type"`
		Amount         string `json:"amount"`
		Status         string `json:"status"`
		TimeRangeType  string `json:"timeRangeType"`
		StartTimeRange string `json:"startTimeRange"`
		EndTimeRange   string `json:"endTimeRange"`
		Channel        string `json:"channel"`
		CreatedBy      string `json:"createdBy"`
		CreatedAt      string `json:"createdAt"`
		UpdatedBy      string `json:"updatedBy"`
		UpdatedAt      string `json:"updatedAt"`
		ApprovedBy     string `json:"approvedBy"`
		ApprovedAt     string `json:"approvedAt"`
		Sofs           []SOF  `json:"sofs"`
	}
	ResponseRuleDetectionGetApproved struct {
		ResponseCode    bool                                   `json:"code"`
		ResponseMessage string                                 `json:"messages"`
		Data            []ResponseRuleDetectionGetApprovedData `json:"rules"`
		Meta            basicObject.Meta                       `json:"meta"`
	}
	RequestRuleDetectionGetApproved struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}
)

type (
	ResponseRuleApprovalById struct {
		Id              string `json:"id"`
		RuleName        string `json:"ruleName"`
		Types           string `json:"type"`
		TransactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"transactionType"`
		Interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"interval"`
		Action struct {
			IsReport bool `json:"isReport"`
			IsBlock  bool `json:"isBlock"`
		} `json:"action"`
		Note           string `json:"rejectNote"`
		Amount         string `json:"amount"`
		Status         string `json:"status"`
		TimeRangeType  string `json:"timeRangeType"`
		StartTimeRange string `json:"startTimeRange"`
		EndTimeRange   string `json:"endTimeRange"`
		Channel        string `json:"channel"`
		RuleId         string `json:"ruleId"`
		ApprovalType   string `json:"approvalType"`
		CreatedBy      string `json:"createdBy"`
		CreatedAt      string `json:"createdAt"`
		UpdatedBy      string `json:"updatedBy"`
		UpdatedAt      string `json:"updatedAt"`
		ApprovedBy     string `json:"approvedBy"`
		ApprovedAt     string `json:"approvedAt"`
		Sofs           []SOF  `json:"sofs"`
	}
	RequestRuleApprovalById struct {
		Id string `param:"id"`
	}
)

type (
	ResponseRuleApprovedById struct {
		Id              string `json:"id"`
		RuleName        string `json:"ruleName"`
		Types           string `json:"type"`
		TransactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"transactionType"`
		Interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"interval"`
		Action struct {
			IsReport bool `json:"isReport"`
			IsBlock  bool `json:"isBlock"`
		} `json:"action"`
		Amount         string `json:"amount"`
		Status         string `json:"status"`
		TimeRangeType  string `json:"timeRangeType"`
		StartTimeRange string `json:"startTimeRange"`
		EndTimeRange   string `json:"endTimeRange"`
		Channel        string `json:"channel"`
		RuleId         string `json:"ruleId"`
		ApprovalType   string `json:"approvalType"`
		CreatedBy      string `json:"createdBy"`
		CreatedAt      string `json:"createdAt"`
		UpdatedBy      string `json:"updatedBy"`
		UpdatedAt      string `json:"updatedAt"`
		ApprovedBy     string `json:"approvedBy"`
		ApprovedAt     string `json:"approvedAt"`
		Sofs           []SOF  `json:"sofs"`
	}
	RequestRuleApprovedById struct {
		Id string `param:"id"`
	}
)

func (b *blueprint) activityRuleTime(timeRangeType, startTimeRange, endTimeRange string) error {
	const layoutDateTime = basicObject.DateAndTime
	const layoutTimeOnly = basicObject.TimeOnly
	const layoutDateOnly = basicObject.DateOnly
	if timeRangeType != "NONE" {
		switch timeRangeType {
		case "DNT":
			str, err := time.Parse(layoutDateTime, startTimeRange)
			if err != nil {
				return fmt.Errorf("%s %s", "start format should be like this ", layoutDateTime)
			}
			etr, err := time.Parse(layoutDateTime, endTimeRange)
			if err != nil {
				return fmt.Errorf("%s %s", "end format should be like this ", layoutDateTime)
			}
			if !etr.After(str) {
				return fmt.Errorf("%s", "end should be greater than start ")
			}
			startTimeRange = str.Format(layoutDateTime)
			endTimeRange = etr.Format(layoutDateTime)
		case "TO":
			str, err := time.Parse(layoutTimeOnly, startTimeRange)
			if err != nil {
				return fmt.Errorf("%s %s", "start format should be like this ", layoutTimeOnly)
			}
			etr, err := time.Parse(layoutTimeOnly, endTimeRange)
			if err != nil {
				return fmt.Errorf("%s %s", "end format should be like this ", layoutTimeOnly)
			}
			if !etr.After(str) {
				return fmt.Errorf("%s", "end should be greater than start ")
			}
			startTimeRange = str.Format(layoutTimeOnly)
			endTimeRange = etr.Format(layoutTimeOnly)
		case "DO":
			str, err := time.Parse(layoutDateOnly, startTimeRange)
			if err != nil {
				return fmt.Errorf("%s %s", "start format should be like this ", layoutDateOnly)
			}
			etr, err := time.Parse(layoutDateOnly, endTimeRange)
			if err != nil {
				return fmt.Errorf("%s %s", "end format should be like this ", layoutDateOnly)
			}
			if !etr.After(str) {
				return fmt.Errorf("%s", "end should be greater than start ")
			}
			startTimeRange = str.Format(layoutDateOnly)
			endTimeRange = etr.Format(layoutDateOnly)
		default:
			return fmt.Errorf("%s %s", "time range type", "invalid time range type")
		}
	}
	return nil
}

func (b *blueprint) RuleDetectionCreate(ctx context.Context, request *RequestRuleDetectionCreate) (*ResponseRuleDetectionCreate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT002"
	const fName = "usecases.uscRuleManagement.RuleDetectionCreate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionCreate */
	// read current approval data
	currentEntity, err := b.rprRuleDetection.ReadRowRuleDetection(ctx, map[string]interface{}{
		"rule_name": request.RuleName,
		"type":      request.Type,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentEntity != nil {
		responseMessage := basicObject.ResponseError{
			Message:    "rule name already exists",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "rule name already exists")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	currentApprovalEntity, err := b.rprRuleDetection.ReadRowRuleDetectionApproval(ctx, map[string]interface{}{
		"rule_name":     request.RuleName,
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
			Message:    "rule name already exists",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "rule name already exists")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	dataEntity := new(rprRuleDetection.EntityRuleApproval)
	amount, err := strconv.ParseFloat(strconv.Itoa(request.Amount), 64)
	if err != nil {
		errMsg := fmt.Errorf("invalid amount value")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	transactionTypeJson, err := json.Marshal(request.TransactionType)
	if err != nil {
		errMsg := fmt.Errorf("invalid data request")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	intervalJson, err := json.Marshal(request.Interval)
	if err != nil {
		errMsg := fmt.Errorf("invalid data request")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	actionJson, err := json.Marshal(request.Action)
	if err != nil {
		errMsg := fmt.Errorf("invalid data request")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	err = b.activityRuleTime(request.TimeRangeType, request.StartTimeRange, request.EndTimeRange)
	if err != nil {
		errMsg := fmt.Errorf("invalid time range rule activity")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	dataEntity.Id = guuid.NewString()
	dataEntity.RuleName = request.RuleName
	dataEntity.Types = request.Type
	dataEntity.TransactionType = string(transactionTypeJson)
	dataEntity.Interval = string(intervalJson)
	dataEntity.Actions = string(actionJson)
	dataEntity.Amount = amount
	dataEntity.Status = basicObject.ApprovalPending
	dataEntity.TimeRangeType = request.TimeRangeType
	dataEntity.StartTimeRange = request.StartTimeRange
	dataEntity.EndTimeRange = request.EndTimeRange
	dataEntity.Sofs = strings.Join(request.Sofs, ",")
	dataEntity.Channel = request.Channel
	dataEntity.RuleId = guuid.NewString()
	dataEntity.ApprovalType = basicObject.CreateApprovalType
	dataEntity.ApprovedBy = basicObject.BlankString
	dataEntity.ApprovedAt = nil
	dataEntity.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.CreatedBy = email
	dataEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.UpdatedBy = email
	entityRuleApproval, entityRuleId, err := b.rprRuleDetection.WriteRowRuleDetectionApproval(ctx, *dataEntity)
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
	response := new(ResponseRuleDetectionCreate)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "new Rules has been add please call checker admin to approve"
	response.Id = dataEntity.Id
	return response, nil
}

func (b *blueprint) RuleDetectionUpdate(ctx context.Context, request *RequestRuleDetectionUpdate) (*ResponseRuleDetectionUpdate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT002"
	const fName = "usecases.uscRuleManagement.RuleDetectionUpdate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionUpdate */
	// read current approval data
	currentApprovalEntity, err := b.rprRuleDetection.ReadRowRuleDetectionApproval(ctx, map[string]interface{}{
		"rule_id":       request.Id,
		"approval_type": basicObject.UpdateApprovalType,
		"status":        1,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentApprovalEntity != nil {
		responseMessage := basicObject.ResponseError{
			Message:    "already exist on approval data",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "already exist on approval data")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	currentRule, err := b.rprRuleDetection.ReadRowRuleDetection(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentRule == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	dataEntity := new(rprRuleDetection.EntityRuleApproval)
	amount, err := strconv.ParseFloat(strconv.Itoa(request.Amount), 64)
	if err != nil {
		errMsg := fmt.Errorf("invalid amount value")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	transactionTypeJson, err := json.Marshal(request.TransactionType)
	if err != nil {
		errMsg := fmt.Errorf("invalid data request")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	intervalJson, err := json.Marshal(request.Interval)
	if err != nil {
		errMsg := fmt.Errorf("invalid data request")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	actionJson, err := json.Marshal(request.Action)
	if err != nil {
		errMsg := fmt.Errorf("invalid data request")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	err = b.activityRuleTime(request.TimeRangeType, request.StartTimeRange, request.EndTimeRange)
	if err != nil {
		errMsg := fmt.Errorf("invalid time range rule activity")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	dataEntity.Id = guuid.NewString()
	dataEntity.RuleName = request.RuleName
	dataEntity.Types = request.Type
	dataEntity.TransactionType = string(transactionTypeJson)
	dataEntity.Interval = string(intervalJson)
	dataEntity.Actions = string(actionJson)
	dataEntity.Amount = amount
	dataEntity.Status = basicObject.ApprovalPending
	dataEntity.TimeRangeType = request.TimeRangeType
	dataEntity.StartTimeRange = request.StartTimeRange
	dataEntity.EndTimeRange = request.EndTimeRange
	dataEntity.Sofs = strings.Join(request.Sofs, ",")
	dataEntity.Channel = request.Channel
	dataEntity.RuleId = currentRule.Id
	dataEntity.ApprovalType = basicObject.UpdateApprovalType
	dataEntity.ApprovedBy = basicObject.BlankString
	dataEntity.ApprovedAt = nil
	dataEntity.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.CreatedBy = email
	dataEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.UpdatedBy = email
	entityRuleApproval, entityRuleId, err := b.rprRuleDetection.WriteRowRuleDetectionApproval(ctx, *dataEntity)
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
	response := new(ResponseRuleDetectionUpdate)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Update rules request has been send to approval"
	response.Id = dataEntity.Id
	return response, nil
}

func (b *blueprint) RuleDetectionDelete(ctx context.Context, request *RequestRuleDetectionDelete) (*ResponseRuleDetectionDelete, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT003"
	const fName = "usecases.uscRuleManagement.RuleDetectionDelete"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionDelete */
	// read current approval data
	currentApprovalEntity, err := b.rprRuleDetection.ReadRowRuleDetectionApproval(ctx, map[string]interface{}{
		"rule_id":       request.Id,
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
			Message:    "already exist on approval data",
			StatusCode: http.StatusBadRequest,
			Reasons:    []string{},
		}
		logging.Errorw(fName, "reason", "already exist on approval data")
		return nil, errWrap.WrapErrorFromResponse(responseMessage)
	}

	currentRule, err := b.rprRuleDetection.ReadRowRuleDetection(ctx, map[string]interface{}{
		"id": request.Id,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentRule == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	amount, err := strconv.ParseFloat(currentRule.Amount, 64)
	if err != nil {
		errMsg := fmt.Errorf("invalid amount value")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	dataEntity := new(rprRuleDetection.EntityRuleApproval)
	dataEntity.Id = guuid.NewString()
	dataEntity.RuleName = currentRule.RuleName
	dataEntity.Types = currentRule.Types
	dataEntity.TransactionType = currentRule.TransactionType
	dataEntity.Interval = currentRule.Interval
	dataEntity.Actions = currentRule.Actions
	dataEntity.Amount = amount
	dataEntity.Status = basicObject.ApprovalPending
	dataEntity.TimeRangeType = currentRule.TimeRangeType
	dataEntity.StartTimeRange = currentRule.StartTimeRange
	dataEntity.EndTimeRange = currentRule.EndTimeRange
	dataEntity.Sofs = currentRule.Sofs
	dataEntity.Channel = currentRule.Channel
	dataEntity.RuleId = currentRule.Id
	dataEntity.ApprovalType = basicObject.DeleteApprovalType
	dataEntity.ApprovedBy = basicObject.BlankString
	dataEntity.ApprovedAt = nil
	dataEntity.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.CreatedBy = email
	dataEntity.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	dataEntity.UpdatedBy = email
	entityRuleApproval, entityRuleId, err := b.rprRuleDetection.WriteRowRuleDetectionApproval(ctx, *dataEntity)
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
	response := new(ResponseRuleDetectionDelete)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Delete rules request has been send to approval"
	response.Id = dataEntity.Id
	return response, nil
}

func (b *blueprint) RuleDetectionApprove(ctx context.Context, request *RequestRuleDetectionApprove) (*ResponseRuleDetectionApprove, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT004"
	const fName = "usecases.uscRuleManagement.RuleDetectionApprove"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionApprove */
	// read current approval data
	currentApprovalEntity, err := b.rprRuleDetection.ReadRowRuleDetectionApproval(ctx, map[string]interface{}{
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
		currentTime := time.Now().Format(basicObject.DateAndTime)
		newEntity := new(rprRuleDetection.EntityRule)
		newEntity.Id = currentApprovalEntity.RuleId
		newEntity.RuleName = currentApprovalEntity.RuleName
		newEntity.Types = currentApprovalEntity.Types
		newEntity.TransactionType = currentApprovalEntity.TransactionType
		newEntity.Interval = currentApprovalEntity.Interval
		newEntity.Amount = fmt.Sprintf("%.2f", currentApprovalEntity.Amount)
		newEntity.Actions = currentApprovalEntity.Actions
		newEntity.Status = currentApprovalEntity.Status
		newEntity.TimeRangeType = currentApprovalEntity.TimeRangeType
		newEntity.StartTimeRange = currentApprovalEntity.StartTimeRange
		newEntity.EndTimeRange = currentApprovalEntity.EndTimeRange
		newEntity.Sofs = currentApprovalEntity.Sofs
		newEntity.Channel = currentApprovalEntity.Channel
		newEntity.CreatedAt = currentTime
		newEntity.UpdatedAt = currentTime
		newEntity.ApprovedAt = &currentTime
		newEntity.CreatedBy = email
		newEntity.UpdatedBy = email
		newEntity.ApprovedBy = email
		_, _, err := b.rprRuleDetection.WriteRowRuleDetection(ctx, *newEntity)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	case basicObject.UpdateApprovalType:
		currentTime := time.Now().Format(basicObject.DateAndTime)
		existingEntity := new(rprRuleDetection.EntityRule)
		existingEntity.Id = currentApprovalEntity.RuleId
		existingEntity.RuleName = currentApprovalEntity.RuleName
		existingEntity.Types = currentApprovalEntity.Types
		existingEntity.TransactionType = currentApprovalEntity.TransactionType
		existingEntity.Interval = currentApprovalEntity.Interval
		existingEntity.Amount = fmt.Sprintf("%.2f", currentApprovalEntity.Amount)
		existingEntity.Actions = currentApprovalEntity.Actions
		existingEntity.Status = currentApprovalEntity.Status
		existingEntity.TimeRangeType = currentApprovalEntity.TimeRangeType
		existingEntity.StartTimeRange = currentApprovalEntity.StartTimeRange
		existingEntity.EndTimeRange = currentApprovalEntity.EndTimeRange
		existingEntity.Sofs = currentApprovalEntity.Sofs
		existingEntity.Channel = currentApprovalEntity.Channel
		existingEntity.CreatedAt = currentTime
		existingEntity.UpdatedAt = currentTime
		existingEntity.ApprovedAt = &currentTime
		existingEntity.CreatedBy = email
		existingEntity.UpdatedBy = email
		existingEntity.ApprovedBy = email
		_, _, err := b.rprRuleDetection.UpdateRowRuleDetection(ctx, *existingEntity)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	case basicObject.DeleteApprovalType:
		err := b.rprRuleDetection.DeleteRowRuleDetection(ctx, currentApprovalEntity.RuleId)
		if err != nil {
			errMsg := fmt.Errorf("internal server error")
			logging.Errorw(fName, "reason", err.Error())
			return nil, errWrap.WrapRepositoryError(errMsg)
		}
	}

	err = b.rprRuleDetection.WriteApprovalForRuleDetectionApproval(ctx, request.Id, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseRuleDetectionApprove)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Approve Rules"
	return response, nil
}

func (b *blueprint) RuleDetectionReject(ctx context.Context, request *RequestRuleDetectionReject) (*ResponseRuleDetectionReject, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT005"
	const fName = "usecases.uscRuleManagement.RuleDetectionReject"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionReject */
	// read current approval data
	currentApprovalEntity, err := b.rprRuleDetection.ReadRowRuleDetectionApproval(ctx, map[string]interface{}{
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
	err = b.rprRuleDetection.WriteRejectForRuleDetectionApproval(ctx, request.Id, request.Note, email)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseRuleDetectionReject)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Success Reject Rules"
	return response, nil
}

func (b *blueprint) RuleDetectionGetApproval(ctx context.Context, request *RequestRuleDetectionGetApproval) (*ResponseRuleDetectionGetApproval, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT006"
	const fName = "usecases.uscRuleManagement.RuleDetectionGetApproval"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionGetApproval */
	dataResult := []ResponseRuleDetectionGetApprovalData{}
	queryData := map[string]interface{}{}
	if request.StatusApproval != 0 {
		queryData["status"] = request.StatusApproval
	}
	data, meta, err := b.rprRuleDetection.ReadRowsRuleDetectionApproval(ctx, queryData, request.Page, request.Limit, "")
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	type (
		transactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		}
		interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		}
		actions struct {
			IsReport bool `json:"is_report"`
			IsBlock  bool `json:"is_block"`
		}
	)

	if data != nil {
		for _, d := range *data {

			var trxType transactionType
			err = json.Unmarshal([]byte(d.TransactionType), &trxType)
			if err != nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			var itrv interval
			err = json.Unmarshal([]byte(d.Interval), &itrv)
			if err != nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			var act actions
			err = json.Unmarshal([]byte(d.Actions), &act)
			if err != nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			var sofData []SOF
			sofIds := strings.Split(d.Sofs, ",")
			for _, si := range sofIds {
				sofDetail, _ := b.rprSof.ReadRowSOF(ctx, map[string]interface{}{
					"id": si,
				})
				if sofDetail != nil {
					sofData = append(sofData, SOF{
						Id:        sofDetail.Id,
						SofName:   sofDetail.SofName,
						Status:    sofDetail.SofStatus,
						CreatedAt: sofDetail.CreatedAt,
						UpdatedAt: sofDetail.UpdatedAt,
					})
				}
			}

			result := ResponseRuleDetectionGetApprovalData{
				Id:             d.Id,
				RuleName:       d.RuleName,
				Types:          d.Types,
				Amount:         fmt.Sprintf("%.2f", d.Amount),
				Status:         d.Status,
				TimeRangeType:  d.TimeRangeType,
				StartTimeRange: d.StartTimeRange,
				EndTimeRange:   d.EndTimeRange,
				Note:           d.Note,
				Sofs:           sofData,
				Channel:        d.Channel,
				RuleId:         d.RuleId,
				ApprovalType:   d.ApprovalType,
				ApprovedBy:     d.ApprovedBy,
				UpdatedAt:      d.UpdatedAt,
				UpdatedBy:      d.UpdatedBy,
				CreatedAt:      d.CreatedAt,
				CreatedBy:      d.CreatedBy,
			}
			status, _ := strconv.Atoi(d.Status)
			result.Status = strconv.Itoa(status - 1)
			result.TransactionType.Value = trxType.Value
			result.TransactionType.Type = trxType.Type
			result.Interval.Value = itrv.Value
			result.Interval.Type = itrv.Type
			result.Action.IsBlock = act.IsBlock
			result.Action.IsReport = act.IsReport
			if d.ApprovedAt != nil {
				result.ApprovedAt = *d.ApprovedAt
			} else {
				result.ApprovedAt = basicObject.BlankString
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseRuleDetectionGetApproval)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) RuleDetectionGetApproved(ctx context.Context, request *RequestRuleDetectionGetApproved) (*ResponseRuleDetectionGetApproved, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT007"
	const fName = "usecases.uscRuleManagement.RuleDetectionGetApproved"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionGetApproved */
	dataResult := []ResponseRuleDetectionGetApprovedData{}
	data, meta, err := b.rprRuleDetection.ReadRowsRuleDetectionApproved(ctx, map[string]interface{}{}, request.Page, request.Limit, "")
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	type (
		transactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		}
		interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		}
		actions struct {
			IsReport bool `json:"is_report"`
			IsBlock  bool `json:"is_block"`
		}
	)

	if data != nil {
		for _, d := range *data {
			amount, err := strconv.ParseFloat(strings.TrimSpace(d.Amount), 64)
			if err != nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			var trxType transactionType
			err = json.Unmarshal([]byte(d.TransactionType), &trxType)
			if err != nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			var itrv interval
			err = json.Unmarshal([]byte(d.Interval), &itrv)
			if err != nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			var act actions
			err = json.Unmarshal([]byte(d.Actions), &act)
			if err != nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			var sofData []SOF
			sofIds := strings.Split(d.Sofs, ",")
			for _, si := range sofIds {
				sofDetail, _ := b.rprSof.ReadRowSOF(ctx, map[string]interface{}{
					"id": si,
				})
				if sofDetail != nil {
					sofData = append(sofData, SOF{
						Id:        sofDetail.Id,
						SofName:   sofDetail.SofName,
						Status:    sofDetail.SofStatus,
						CreatedAt: sofDetail.CreatedAt,
						UpdatedAt: sofDetail.UpdatedAt,
					})
				}
			}

			result := ResponseRuleDetectionGetApprovedData{
				Id:             d.Id,
				RuleName:       d.RuleName,
				Types:          d.Types,
				Amount:         fmt.Sprintf("%.2f", amount),
				Status:         d.Status,
				TimeRangeType:  d.TimeRangeType,
				StartTimeRange: d.StartTimeRange,
				EndTimeRange:   d.EndTimeRange,
				Sofs:           sofData,
				Channel:        d.Channel,
				ApprovedBy:     d.ApprovedBy,
				UpdatedAt:      d.UpdatedAt,
				UpdatedBy:      d.UpdatedBy,
				CreatedAt:      d.CreatedAt,
				CreatedBy:      d.CreatedBy,
			}
			//status, _ := strconv.Atoi(d.Status)
			//result.Status = strconv.Itoa(status - 1)
			result.TransactionType.Value = trxType.Value
			result.TransactionType.Type = trxType.Type
			result.Interval.Value = itrv.Value
			result.Interval.Type = itrv.Type
			result.Action.IsBlock = act.IsBlock
			result.Action.IsReport = act.IsReport
			if d.ApprovedAt != nil {
				result.ApprovedAt = *d.ApprovedAt
			} else {
				result.ApprovedAt = basicObject.BlankString
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseRuleDetectionGetApproved)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = basicObject.Successfully
	response.Data = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) RuleDetectionGetApprovalById(ctx context.Context, request *RequestRuleApprovalById) (*ResponseRuleApprovalById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT007"
	const fName = "usecases.uscRuleManagement.RuleDetectionGetApprovalById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RRuleDetectionGetApprovalById */
	// read current approval data
	currentApprovalEntity, err := b.rprRuleDetection.ReadRowRuleDetectionApproval(ctx, map[string]interface{}{
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

	type (
		transactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		}
		interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		}
		actions struct {
			IsReport bool `json:"is_report"`
			IsBlock  bool `json:"is_block"`
		}
	)

	var trxType transactionType
	err = json.Unmarshal([]byte(currentApprovalEntity.TransactionType), &trxType)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	var itrv interval
	err = json.Unmarshal([]byte(currentApprovalEntity.Interval), &itrv)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	var act actions
	err = json.Unmarshal([]byte(currentApprovalEntity.Actions), &act)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	var sofData []SOF
	sofIds := strings.Split(currentApprovalEntity.Sofs, ",")
	for _, si := range sofIds {
		sofDetail, _ := b.rprSof.ReadRowSOF(ctx, map[string]interface{}{
			"id": si,
		})
		if sofDetail != nil {
			sofData = append(sofData, SOF{
				Id:        sofDetail.Id,
				SofName:   sofDetail.SofName,
				Status:    sofDetail.SofStatus,
				CreatedAt: sofDetail.CreatedAt,
				UpdatedAt: sofDetail.UpdatedAt,
			})
		}
	}

	response := ResponseRuleApprovalById{
		Id:              currentApprovalEntity.Id,
		RuleName:        currentApprovalEntity.RuleName,
		Types:           currentApprovalEntity.Types,
		TransactionType: trxType,
		Interval:        itrv,
		Action: struct {
			IsReport bool `json:"isReport"`
			IsBlock  bool `json:"isBlock"`
		}{
			IsReport: act.IsReport,
			IsBlock:  act.IsBlock,
		},
		Amount:         fmt.Sprintf("%.2f", currentApprovalEntity.Amount),
		Status:         currentApprovalEntity.Status,
		TimeRangeType:  currentApprovalEntity.TimeRangeType,
		StartTimeRange: currentApprovalEntity.StartTimeRange,
		EndTimeRange:   currentApprovalEntity.EndTimeRange,
		Sofs:           sofData,
		Channel:        currentApprovalEntity.Channel,
		RuleId:         currentApprovalEntity.RuleId,
		ApprovalType:   currentApprovalEntity.ApprovalType,
		ApprovedBy:     currentApprovalEntity.ApprovedBy,
		UpdatedAt:      currentApprovalEntity.UpdatedAt,
		UpdatedBy:      currentApprovalEntity.UpdatedBy,
		CreatedAt:      currentApprovalEntity.CreatedAt,
		CreatedBy:      currentApprovalEntity.CreatedBy,
		Note:           currentApprovalEntity.Note,
	}
	status, _ := strconv.Atoi(currentApprovalEntity.Status)
	response.Status = strconv.Itoa(status - 1)
	if currentApprovalEntity.ApprovedAt != nil {
		response.ApprovedAt = *currentApprovalEntity.ApprovedAt
	} else {
		response.ApprovedAt = basicObject.BlankString
	}
	return &response, nil
}

func (b *blueprint) RuleDetectionGetApprovedById(ctx context.Context, request *RequestRuleApprovedById) (*ResponseRuleApprovedById, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT007"
	const fName = "usecases.uscRuleManagement.RuleDetectionGetApprovedById"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionGetApprovedById */
	// read current approval data
	currentApprovalEntity, err := b.rprRuleDetection.ReadRowRuleDetectionApproved(ctx, map[string]interface{}{
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

	type (
		transactionType struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		}
		interval struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		}
		actions struct {
			IsReport bool `json:"is_report"`
			IsBlock  bool `json:"is_block"`
		}
	)

	amount, err := strconv.ParseFloat(strings.TrimSpace(currentApprovalEntity.Amount), 64)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	var trxType transactionType
	err = json.Unmarshal([]byte(currentApprovalEntity.TransactionType), &trxType)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	var itrv interval
	err = json.Unmarshal([]byte(currentApprovalEntity.Interval), &itrv)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	var act actions
	err = json.Unmarshal([]byte(currentApprovalEntity.Actions), &act)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	var sofData []SOF
	sofIds := strings.Split(currentApprovalEntity.Sofs, ",")
	for _, si := range sofIds {
		sofDetail, _ := b.rprSof.ReadRowSOF(ctx, map[string]interface{}{
			"id": si,
		})
		if sofDetail != nil {
			sofData = append(sofData, SOF{
				Id:        sofDetail.Id,
				SofName:   sofDetail.SofName,
				Status:    sofDetail.SofStatus,
				CreatedAt: sofDetail.CreatedAt,
				UpdatedAt: sofDetail.UpdatedAt,
			})
		}
	}

	response := ResponseRuleApprovedById{
		Id:              currentApprovalEntity.Id,
		RuleName:        currentApprovalEntity.RuleName,
		Types:           currentApprovalEntity.Types,
		TransactionType: trxType,
		Interval:        itrv,
		Action: struct {
			IsReport bool `json:"isReport"`
			IsBlock  bool `json:"isBlock"`
		}{
			IsReport: act.IsReport,
			IsBlock:  act.IsBlock,
		},
		Amount:         fmt.Sprintf("%.2f", amount),
		Status:         currentApprovalEntity.Status,
		TimeRangeType:  currentApprovalEntity.TimeRangeType,
		StartTimeRange: currentApprovalEntity.StartTimeRange,
		EndTimeRange:   currentApprovalEntity.EndTimeRange,
		Sofs:           sofData,
		Channel:        currentApprovalEntity.Channel,
		ApprovedBy:     currentApprovalEntity.ApprovedBy,
		UpdatedAt:      currentApprovalEntity.UpdatedAt,
		UpdatedBy:      currentApprovalEntity.UpdatedBy,
		CreatedAt:      currentApprovalEntity.CreatedAt,
		CreatedBy:      currentApprovalEntity.CreatedBy,
	}
	//status, _ := strconv.Atoi(currentApprovalEntity.Status)
	//response.Status = strconv.Itoa(status - 1)
	if currentApprovalEntity.ApprovedAt != nil {
		response.ApprovedAt = *currentApprovalEntity.ApprovedAt
	} else {
		response.ApprovedAt = basicObject.BlankString
	}
	return &response, nil
}

func (b *blueprint) RuleDetectionEnable(ctx context.Context, request *RequestRuleEnabled) (*ResponseRuleEnabled, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT002"
	const fName = "usecases.uscRuleManagement.RuleDetectionEnable"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionEnable */
	// read current approval data
	currentEntity, err := b.rprRuleDetection.ReadRowRuleDetection(ctx, map[string]interface{}{"id": request.Id})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentEntity == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	currentTime := time.Now().Format(basicObject.DateAndTime)
	existingEntity := new(rprRuleDetection.EntityRule)
	existingEntity.Id = currentEntity.Id
	existingEntity.Status = "1"
	existingEntity.CreatedAt = currentTime
	existingEntity.UpdatedAt = currentTime
	existingEntity.ApprovedAt = &currentTime
	existingEntity.CreatedBy = email
	existingEntity.UpdatedBy = email
	existingEntity.ApprovedBy = email
	_, _, err = b.rprRuleDetection.UpdateRowRuleDetection(ctx, *existingEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseRuleEnabled)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Rules has been enabled"
	response.Id = existingEntity.Id
	return response, nil
}

func (b *blueprint) RuleDetectionDisable(ctx context.Context, request *RequestRuleDisabled) (*ResponseRuleDisabled, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "RDT002"
	const fName = "usecases.uscRuleManagement.RuleDetectionDisable"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_RuleDetectionDisable */
	// read current approval data
	currentEntity, err := b.rprRuleDetection.ReadRowRuleDetection(ctx, map[string]interface{}{"id": request.Id})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if currentEntity == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	currentTime := time.Now().Format(basicObject.DateAndTime)
	existingEntity := new(rprRuleDetection.EntityRule)
	existingEntity.Id = currentEntity.Id
	existingEntity.Status = "0"
	existingEntity.CreatedAt = currentTime
	existingEntity.UpdatedAt = currentTime
	existingEntity.ApprovedAt = &currentTime
	existingEntity.CreatedBy = email
	existingEntity.UpdatedBy = email
	existingEntity.ApprovedBy = email
	_, _, err = b.rprRuleDetection.UpdateRowRuleDetection(ctx, *existingEntity)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseRuleDisabled)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "Rules has been disabled"
	response.Id = existingEntity.Id
	return response, nil
}
