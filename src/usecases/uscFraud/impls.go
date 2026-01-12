package uscFraud

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"gitlab.com/fds22/detection-sys/pkg/environments"
	"gitlab.com/fds22/detection-sys/src/repositories/rprTransactionType"

	guuid "github.com/google/uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/src/repositories/rprFraudDetection"
	"go.uber.org/zap"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint business function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */
		Detect(ctx context.Context, request *RequestDetect) (*ResponseDetect, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		rprFraudDetection  rprFraudDetection.Blueprint
		rprTransactionType rprTransactionType.Blueprint
		env                *environments.Envs
		log                *zap.SugaredLogger // log logging instance
	}
)

func NewBlueprint(
	rprFraudDetection rprFraudDetection.Blueprint,
	rprTransactionType rprTransactionType.Blueprint,
	env *environments.Envs,
	log *zap.SugaredLogger,
) Blueprint {
	const fName = "usecases.uscFraud.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.rprFraudDetection = rprFraudDetection
	bp.rprTransactionType = rprTransactionType
	bp.env = env
	bp.log = log
	return bp
}

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */

func containsRuleData(rules []basicObject.RuleData, ruleId string) bool {
	for _, rule := range rules {
		if rule.RuleId == ruleId {
			return true
		}
	}
	return false
}

func containsRuleEntity(rules []rprFraudDetection.EntityRule, ruleId string) bool {
	for _, rule := range rules {
		if rule.Id == ruleId {
			return true
		}
	}
	return false
}

func (b *blueprint) Detect(ctx context.Context, request *RequestDetect) (*ResponseDetect, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "DTC001"
	const fName = "usecases.uscFraud.Detect"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_Detect */
	response := new(ResponseDetect)
	var rulesEntityEvaluation []rprFraudDetection.EntityRule
	var evaluationData basicObject.EvaluationData
	const successResponseCode = "00"
	const successResponseDescription = "Request successful"

	// Detection Eligibility flow

	// Retrieve Transaction History
	// Parse transaction time from request into time.DateTime layout

	parsedTime, err := time.Parse(time.DateTime, request.TransactionTime)
	if err != nil {
		errMsg := fmt.Errorf("invalid transaction time")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapBusinessError(errMsg)
	}

	// Convert amount to float
	requestAmount, err := strconv.ParseFloat(request.Amount.Value, 64)
	if err != nil {
		errMsg := fmt.Errorf("invalid amount value")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// Accessing components and extract it into variables
	year := parsedTime.Year()
	month := parsedTime.Month()
	day := parsedTime.Day()
	hour := parsedTime.Hour()
	minute := parsedTime.Minute()
	second := parsedTime.Second()
	now := time.Date(year, month, day, hour, minute, second, 0, time.Local)

	// Adding default transaction id if partner reference number not filled
	var partnerReferenceNo string
	if len(request.PartnerReferenceNumber) == 0 {
		partnerReferenceNo = guuid.NewString()
	} else {
		partnerReferenceNo = request.PartnerReferenceNumber
	}

	// Insert transaction history data
	transactionId := guuid.NewString()
	newInsertedTrxData, newTrxDataId, err := b.rprFraudDetection.WriteRowTransaction(
		ctx, rprFraudDetection.EntityTransaction{
			TransactionId:   partnerReferenceNo,
			Rules:           "",
			Title:           request.TransactionType,
			FlagId:          basicObject.WarningFlagId,
			CreatedAt:       now.Format(basicObject.DateAndTime),
			Amount:          request.Amount.Value,
			Id:              transactionId,
			UserId:          request.BenefactorIdentityNumber,
			DestinationId:   request.BeneficiaryIdentityNumber,
			BodyReq:         "",
			Channel:         request.Channel,
			TransactionType: request.TransactionType,
		},
	)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if newInsertedTrxData == nil && newTrxDataId == nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", "failed to new insert trx data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// Retrieve Blacklist DTTOT
	// Checking benefactor name
	benefactorResult, err := b.rprFraudDetection.ReadRowsBlacklistDTTOTByPerformerName(ctx,
		strings.ToUpper(request.BenefactorName))
	if err != nil && err.Error() != errors.New("not found").Error() {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if benefactorResult != nil {
		for _, benefactor := range *benefactorResult {
			var dttotData basicObject.BlacklistDTTOTData
			dttotData.Name = benefactor.Name
			dttotData.FileId = benefactor.FileId
			dttotData.FileLink = benefactor.FileLink
			dttotData.Datasource = benefactor.Datasource
			dttotData.NIK = benefactor.NIK
			dttotData.Performer = basicObject.Benefactor
			evaluationData.BlacklistDTTOT = append(evaluationData.BlacklistDTTOT, dttotData)
		}
	}

	// Checking beneficiary name
	beneficiaryResult, err := b.rprFraudDetection.ReadRowsBlacklistDTTOTByPerformerName(ctx,
		strings.ToUpper(request.BeneficiaryName))
	if err != nil && err.Error() != errors.New("not found").Error() {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if beneficiaryResult != nil {
		for _, beneficiary := range *beneficiaryResult {
			var dttotData basicObject.BlacklistDTTOTData
			dttotData.Name = beneficiary.Name
			dttotData.FileId = beneficiary.FileId
			dttotData.FileLink = beneficiary.FileLink
			dttotData.Datasource = beneficiary.Datasource
			dttotData.NIK = beneficiary.NIK
			dttotData.Performer = basicObject.Beneficiary
			evaluationData.BlacklistDTTOT = append(evaluationData.BlacklistDTTOT, dttotData)
		}
	}

	// Retrieve Blacklist Receiver
	receiverResult, err := b.rprFraudDetection.ReadRowsBlacklistReceiverByPerformerName(ctx,
		strings.ToUpper(request.BeneficiaryName), request.BeneficiaryIdentityNumber, request.TransactionType)
	if err != nil && err.Error() != errors.New("not found").Error() {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if receiverResult != nil {
		for _, receiver := range *receiverResult {
			var receiverData basicObject.BlacklistReceiverData
			receiverData.Name = receiver.BeneficiaryName
			receiverData.Identity = receiver.PhoneNumber
			evaluationData.BlacklistReceiver = append(evaluationData.BlacklistReceiver, receiverData)
		}
	}

	// Retrieve Blacklist Sender
	senderResult, err := b.rprFraudDetection.ReadRowsBlacklistSenderByPerformerName(ctx,
		strings.ToUpper(request.BenefactorName), request.BenefactorIdentityNumber, request.TransactionType)
	if err != nil && err.Error() != errors.New("not found").Error() {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if senderResult != nil {
		for _, sender := range *senderResult {
			var senderData basicObject.BlacklistSenderData
			senderData.Name = sender.BeneficiaryName
			senderData.Identity = sender.PhoneNumber
			evaluationData.BlacklistSender = append(evaluationData.BlacklistSender, senderData)
		}
	}

	// Retrieve Blacklist Merchant
	merchantResult, err := b.rprFraudDetection.ReadRowsBlacklistMerchantByPerformerNameAndUserId(ctx,
		strings.ToUpper(request.BeneficiaryName), request.BeneficiaryIdentityNumber)
	if err != nil && err.Error() != errors.New("not found").Error() {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if merchantResult != nil {
		for _, merchant := range *merchantResult {
			var merchantData basicObject.BlacklistMerchantData
			merchantData.Name = merchant.MerchantName
			merchantData.NMID = merchant.NMID
			merchantData.FileLink = merchant.FileLink
			merchantData.FileId = merchant.FileId
			merchantData.Datasource = merchant.Datasource
			evaluationData.BlacklistMerchant = append(evaluationData.BlacklistMerchant, merchantData)
		}
	}

	// Retrieve Transaction Rule
	ct := now
	var shortedRules []rprFraudDetection.EntityRule
	rules, err := b.rprFraudDetection.ReadRowsRules(ctx, map[string]interface{}{
		"type": request.TransactionType,
		"sofs": request.Sof,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if rules == nil {
		// condition there is no rule
		shortedRules = []rprFraudDetection.EntityRule{}
	} else {
		if len(*rules) <= 0 {
			// condition there is no rule
			shortedRules = []rprFraudDetection.EntityRule{}
		} else {
			// condition there is rules for that
			shortedRules = *rules
		}
	}

	// sorting the rules
	sort.Slice(shortedRules, func(i, j int) bool {
		var trxTypeI basicObject.TypeAndValue2
		var intervalI basicObject.TypeAndValue1
		_ = json.Unmarshal([]byte(shortedRules[i].TransactionType), &trxTypeI)
		_ = json.Unmarshal([]byte(shortedRules[i].Interval), &intervalI)

		var trxTypeJ basicObject.TypeAndValue2
		var intervalJ basicObject.TypeAndValue1
		_ = json.Unmarshal([]byte(shortedRules[j].TransactionType), &trxTypeJ)
		_ = json.Unmarshal([]byte(shortedRules[j].Interval), &intervalJ)

		//countCondition := trxTypeI.Value < trxTypeJ.Value
		amountCondition := shortedRules[i].Amount < shortedRules[j].Amount
		return amountCondition
	})
	logging.Errorw(fName, "rules_data", shortedRules)
	for _, rule := range shortedRules {
		ruleTimeValidity := false
		switch {
		case rule.TimeRangeType != basicObject.NoneEnum &&
			rule.StartTimeRange != basicObject.NoneEnum &&
			rule.EndTimeRange != basicObject.NoneEnum:
			logging.Errorw(fName, "rule_time_setting", "available")
			if rule.TimeRangeType == basicObject.TimeOnlyEnum {
				ruleSTR, _ := time.ParseInLocation(basicObject.TimeOnly, rule.StartTimeRange, time.Local)
				ruleETR, _ := time.ParseInLocation(basicObject.TimeOnly, rule.EndTimeRange, time.Local)
				startTime := time.Date(ct.Year(), ct.Month(), ct.Day(),
					ruleSTR.Hour(), ruleSTR.Minute(), 0, 0, time.Local)
				endTime := time.Date(ct.Year(), ct.Month(), ct.Day(),
					ruleETR.Hour(), ruleETR.Minute(), 0, 0, time.Local)
				logging.Errorw(fName, "rule_time_setting", "date_and_time_only", "start_raw", ruleSTR.String(), "end_raw", ruleETR.String())
				logging.Errorw(fName, "rule_time_setting", "date_and_time_only", "current", ct.String())
				valid := !ct.Before(startTime) && !ct.After(endTime)
				if valid {
					logging.Errorw(fName, "rule_time_setting", "validated request within the the time setting time-only")
					ruleTimeValidity = true
				}
			}

			if rule.TimeRangeType == basicObject.DateOnlyEnum {
				ruleSTR, _ := time.ParseInLocation(basicObject.DateOnly, rule.StartTimeRange, time.Local)
				ruleETR, _ := time.ParseInLocation(basicObject.DateOnly, rule.EndTimeRange, time.Local)
				startTime := time.Date(ct.Year(), ruleSTR.Month(),
					ruleSTR.Day(), ct.Hour(), ct.Minute(), 0, 0, time.Local)
				endTime := time.Date(ct.Year(), ruleETR.Month(),
					ruleETR.Day(), ct.Hour(), ct.Minute(), 0, 0, time.Local)
				logging.Errorw(fName, "rule_time_setting", "date_and_time_only", "start_raw", ruleSTR.String(), "end_raw", ruleETR.String())
				logging.Errorw(fName, "rule_time_setting", "date_and_time_only", "current", ct.String())
				valid := !ct.Before(startTime) && !ct.After(endTime)
				if valid {
					logging.Errorw(fName, "rule_time_setting", "validated request within the the time setting date-only")
					ruleTimeValidity = true
				}
			}

			if rule.TimeRangeType == basicObject.DateAndTimeEnum {
				ruleSTR, _ := time.ParseInLocation(basicObject.DateAndTime, rule.StartTimeRange, time.Local)
				ruleETR, _ := time.ParseInLocation(basicObject.DateAndTime, rule.EndTimeRange, time.Local)
				logging.Errorw(fName, "rule_time_setting", "date_and_time_only", "start_raw", ruleSTR.String(), "end_raw", ruleETR.String())
				logging.Errorw(fName, "rule_time_setting", "date_and_time_only", "current", ct.String())
				valid := !ct.Before(ruleSTR) && !ct.After(ruleETR)
				if valid {
					logging.Errorw(fName, "rule_time_setting", "validated request within the the time setting date-and-time")
					ruleTimeValidity = true
				}
			}
		case rule.TimeRangeType == basicObject.NoneEnum &&
			rule.StartTimeRange == basicObject.NoneEnum &&
			rule.EndTimeRange == basicObject.NoneEnum:
			logging.Errorw(fName, "rule_time_setting", "unavailable")
			logging.Errorw(fName, "rule_time_setting", "none")
			ruleTimeValidity = true
		default:
			errMsg := fmt.Errorf("invalid rule time range setting")
			logging.Errorw(fName, "reason", errMsg)
			return nil, errWrap.WrapRepositoryError(errMsg)
		}

		if ruleTimeValidity {
			var trxType basicObject.TypeAndValue2
			var interval basicObject.TypeAndValue1
			var action basicObject.Actions

			err := json.Unmarshal([]byte(rule.TransactionType), &trxType)
			if err != nil {
				errMsg := fmt.Errorf("malformed transaction type")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			err = json.Unmarshal([]byte(rule.Interval), &interval)
			if err != nil {
				errMsg := fmt.Errorf("malformed interval")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			err = json.Unmarshal([]byte(rule.Actions), &action)
			if err != nil {
				errMsg := fmt.Errorf("malformed action")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			// Detection by interval time
			var intervalViolation bool
			var durationString string
			var intervalDuration time.Duration
			var errIntervalDurationParse error
			switch interval.Type {
			case "second":
				durationString = fmt.Sprintf("%ds", interval.Value)
				intervalDuration, errIntervalDurationParse = time.ParseDuration(durationString)
			case "minute":
				durationString = fmt.Sprintf("%dm", interval.Value)
				intervalDuration, errIntervalDurationParse = time.ParseDuration(durationString)
			case "hour":
				durationString = fmt.Sprintf("%dh", interval.Value)
				intervalDuration, errIntervalDurationParse = time.ParseDuration(durationString)
			}
			if errIntervalDurationParse != nil {
				errMsg := fmt.Errorf("invalid interval data from datasouce")
				logging.Errorw(fName, "reason", errMsg)
				return nil, errWrap.WrapRepositoryError(errMsg)
			}

			// Retrieve transaction data by user id and start and end date
			thresholdIntervalDuration := intervalDuration + 1*time.Second
			endDate := parsedTime
			//endDate := parsedTime.Local()
			startDate := endDate.Add(-thresholdIntervalDuration)

			var totalAmountWithDateRange float64
			transactions, err := b.rprFraudDetection.
				ReadRowsTransactionsByPhoneAndDateRange(ctx, request.BenefactorIdentityNumber, startDate, endDate)
			if err != nil {
				errMsg := fmt.Errorf("transaction data not found")
				logging.Errorw(fName, "reason", errMsg.Error())
				return nil, errWrap.WrapRepositoryError(err)
			}
			if transactions != nil {
				for _, trx := range *transactions {
					feetFloat, _ := strconv.ParseFloat(strings.TrimSpace(trx.Amount), 64)
					logging.Infow(fName, "transaction_amount_transaction_with_date_range", feetFloat)
					totalAmountWithDateRange += feetFloat
				}
				logging.Infow(fName, "transaction_amount_transaction_with_date_range_total", totalAmountWithDateRange)
				if len(*transactions) > 0 {
					latestTransaction := (*transactions)[0]
					lastTransactionTime, err := time.ParseInLocation(basicObject.DatabaseDateAndTime,
						latestTransaction.CreatedAt, time.Local)
					if err != nil {
						errMsg := fmt.Errorf("failed to parse")
						logging.Errorw(fName, "reason", err.Error())
						return nil, errWrap.WrapRepositoryError(errMsg)
					}
					intervalViolation = now.Sub(lastTransactionTime) < intervalDuration
					logging.Errorw(fName, "interval_violation", intervalViolation)
				}
			}

			//if totalAmountWithDateRange <= rule.Amount {
			if trxType.Type == basicObject.Multiple {
				if intervalViolation {
					if totalAmountWithDateRange > rule.Amount {
						var ruleData basicObject.RuleData
						ruleData.RuleId = rule.Id
						ruleData.IsRejected = action.IsBlock
						ruleData.IsReported = action.IsReport
						ruleData.RuleName = rule.RuleName
						ruleData.RuleType = rule.Types
						ruleData.Violation = basicObject.Transaction
						ruleData.TransactionType = basicObject.Multiple
						if !containsRuleData(evaluationData.Rules, ruleData.RuleId) {
							evaluationData.Rules = append(evaluationData.Rules, ruleData)
						}
						if !containsRuleEntity(rulesEntityEvaluation, rule.Id) {
							rulesEntityEvaluation = append(rulesEntityEvaluation, rule)
						}
						logging.Errorw(fName, "transaction_type", basicObject.Multiple)
					}
				}
			}

			if trxType.Type == basicObject.Multiple {
				if intervalViolation {
					trxTypeValue, _ := strconv.Atoi(trxType.Value)
					if transactions != nil {
						if len(*transactions) > trxTypeValue {
							if totalAmountWithDateRange >= rule.Amount && totalAmountWithDateRange != rule.Amount {
								var ruleData basicObject.RuleData
								ruleData.RuleId = rule.Id
								ruleData.IsRejected = action.IsBlock
								ruleData.IsReported = action.IsReport
								ruleData.RuleName = rule.RuleName
								ruleData.RuleType = rule.Types
								ruleData.Violation = basicObject.Transaction
								ruleData.TransactionType = basicObject.Multiple
								if !containsRuleData(evaluationData.Rules, ruleData.RuleId) {
									evaluationData.Rules = append(evaluationData.Rules, ruleData)
								}
								if !containsRuleEntity(rulesEntityEvaluation, rule.Id) {
									rulesEntityEvaluation = append(rulesEntityEvaluation, rule)
								}
								logging.Errorw(fName, "transaction_type", basicObject.Multiple)
							} else {
								var ruleData basicObject.RuleData
								ruleData.RuleId = rule.Id
								ruleData.IsRejected = action.IsBlock
								ruleData.IsReported = action.IsReport
								ruleData.RuleName = rule.RuleName
								ruleData.RuleType = rule.Types
								ruleData.Violation = basicObject.Transaction
								ruleData.TransactionType = basicObject.Multiple
								if !containsRuleData(evaluationData.Rules, ruleData.RuleId) {
									evaluationData.Rules = append(evaluationData.Rules, ruleData)
								}
								if !containsRuleEntity(rulesEntityEvaluation, rule.Id) {
									rulesEntityEvaluation = append(rulesEntityEvaluation, rule)
								}
								logging.Errorw(fName, "transaction_type", basicObject.Multiple)
							}
						}
					}
				}
			}

			if trxType.Type == basicObject.Single {
				if requestAmount >= rule.Amount && requestAmount != rule.Amount {
					var ruleData basicObject.RuleData
					ruleData.RuleId = rule.Id
					ruleData.IsRejected = action.IsBlock
					ruleData.IsReported = action.IsReport
					ruleData.RuleName = rule.RuleName
					ruleData.RuleType = rule.Types
					ruleData.Violation = basicObject.Transaction
					ruleData.TransactionType = basicObject.Single
					if !containsRuleData(evaluationData.Rules, ruleData.RuleId) {
						evaluationData.Rules = append(evaluationData.Rules, ruleData)
					}
					if !containsRuleEntity(rulesEntityEvaluation, rule.Id) {
						rulesEntityEvaluation = append(rulesEntityEvaluation, rule)
					}
					logging.Errorw(fName, "transaction_type", basicObject.Single)
				}
			}
			//}
		}
	}

	// Keyword retrieval
	keywords, err := b.rprFraudDetection.ReadRowsKeyword(ctx)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if keywords != nil {
		for mkIdx, mk := range *keywords {
			words := strings.Split(mk.Keyword, " ")
			logging.Infow(fName, "matched_id", mkIdx, "matched_keyword", mk)
			switch mk.Action {
			case basicObject.KeywordActionBlockWithReport:
				for _, word := range words {
					if strings.Contains(strings.ToLower(request.BeneficiaryName), strings.ToLower(word)) {
						evaluationData.Keywords = append(evaluationData.Keywords, basicObject.KeywordData{
							KeywordId:       mk.Id,
							KeywordViolated: mk.Keyword,
							Allowed:         false,
							Report:          true,
						})
					}
					if strings.Contains(strings.ToLower(request.BenefactorName), strings.ToLower(word)) {
						evaluationData.Keywords = append(evaluationData.Keywords, basicObject.KeywordData{
							KeywordId:       mk.Id,
							KeywordViolated: mk.Keyword,
							Allowed:         false,
							Report:          true,
						})
					}
				}
			case basicObject.KeywordActionAllowWithReport:
				for _, word := range words {
					if strings.Contains(strings.ToLower(request.BeneficiaryName), strings.ToLower(word)) {
						evaluationData.Keywords = append(evaluationData.Keywords, basicObject.KeywordData{
							KeywordId:       mk.Id,
							KeywordViolated: mk.Keyword,
							Allowed:         true,
							Report:          true,
						})
					}
					if strings.Contains(strings.ToLower(request.BenefactorName), strings.ToLower(word)) {
						evaluationData.Keywords = append(evaluationData.Keywords, basicObject.KeywordData{
							KeywordId:       mk.Id,
							KeywordViolated: mk.Keyword,
							Allowed:         true,
							Report:          true,
						})
					}
				}
			default:
				errMsg := fmt.Errorf("invalid keyword rule data from datasouce")
				logging.Errorw(fName, "reason", errMsg.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}
		}
	}

	// Inserting log data
	additionalInfoBytes, _ := json.Marshal(request.AdditionalInfo)
	insertedLogData, logDataId, err := b.rprFraudDetection.WriteRowLog(ctx, rprFraudDetection.EntityLog{
		Id:              guuid.NewString(),
		UserId:          request.BenefactorIdentityNumber,
		DestinationId:   request.BeneficiaryIdentityNumber,
		Amount:          requestAmount,
		StartDate:       now.Format(basicObject.DateAndTime),
		EndDate:         time.Now().Format(basicObject.DateAndTime),
		BodyReq:         string(additionalInfoBytes),
		Channel:         request.Channel,
		TransactionType: request.TransactionType,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if insertedLogData == nil && logDataId == nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", "failed to insert log data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// Re-evaluate detection journey
	var ruleViolationMustReject []basicObject.RuleData
	var ruleViolationMustReport []basicObject.RuleData
	for _, ruleViolation := range evaluationData.Rules {
		if ruleViolation.IsRejected {
			ruleViolationMustReject = append(ruleViolationMustReject, ruleViolation)
		}
		if ruleViolation.IsReported {
			ruleViolationMustReport = append(ruleViolationMustReport, ruleViolation)
		}
	}

	// Re-evaluate detection keyword
	var keywordRejected []basicObject.KeywordData
	var keywordReported []basicObject.KeywordData
	for _, kd := range evaluationData.Keywords {
		if kd.Report {
			keywordReported = append(keywordReported, kd)
		}
		if !kd.Allowed {
			keywordRejected = append(keywordRejected, kd)
		}
	}

	// Report and Reject decision
	hasDTTOTRecord := len(evaluationData.BlacklistDTTOT) != 0
	hasSenderRecord := len(evaluationData.BlacklistSender) != 0
	hasReceiverRecord := len(evaluationData.BlacklistReceiver) != 0
	hasMerchantRecord := len(evaluationData.BlacklistMerchant) != 0
	rejected := (len(ruleViolationMustReject) > 0 || len(keywordRejected) > 0) ||
		(hasDTTOTRecord || hasMerchantRecord || hasReceiverRecord || hasSenderRecord)
	reported := len(ruleViolationMustReport) > 0 || len(keywordReported) > 0
	var transactionSevereLevelCategory = basicObject.AllowedFlagId

	defaultAct := strings.TrimSpace(strings.ToLower(b.env.SettingsFDSGovDetectionDefaultAct))

	transactionTypesData, err := b.rprTransactionType.ReadRowsTransactionTypeApproved(ctx, map[string]interface{}{}, 0, 0, true)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if transactionTypesData == nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", "failed to read transaction types data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	var transactionType []string
	for _, trxTypeData := range *transactionTypesData {
		transactionType = append(transactionType, trxTypeData.Name)
	}

	// Rejected and Reported
	if rejected {
		if defaultAct == basicObject.GovActReject {
			response.ResponseObject.Conclusion = basicObject.Rejected
			transactionSevereLevelCategory = basicObject.RejectedFlagId
			if senderResult != nil && len(*senderResult) == 0 {
				insertedSenderData, trxDataId, err := b.rprFraudDetection.WriteRowBlacklistSender(
					ctx, rprFraudDetection.EntityBlacklistSender{
						Id:               guuid.NewString(),
						PhoneNumber:      request.BenefactorIdentityNumber,
						BeneficiaryName:  request.BenefactorName,
						CreatedAt:        now.Format(basicObject.DateAndTime),
						CreatedBy:        request.Channel,
						UpdatedAt:        now.Format(basicObject.DateAndTime),
						UpdatedBy:        request.Channel,
						Status:           1,
						TransactionTypes: strings.Join(transactionType, ","),
					},
				)
				if err != nil {
					errMsg := fmt.Errorf("internal server error")
					logging.Errorw(fName, "reason", err.Error())
					return nil, errWrap.WrapRepositoryError(errMsg)
				}
				if insertedSenderData == nil && trxDataId == nil {
					errMsg := fmt.Errorf("internal server error")
					logging.Errorw(fName, "reason", "failed to insert new blacklist sender data")
					return nil, errWrap.WrapRepositoryError(errMsg)
				}
			}
		}
		if defaultAct == basicObject.GovActAllow {
			if hasDTTOTRecord {
				response.ResponseObject.Conclusion = basicObject.Allowed
				transactionSevereLevelCategory = basicObject.WarningFlagId
			} else {
				response.ResponseObject.Conclusion = basicObject.Rejected
				transactionSevereLevelCategory = basicObject.RejectedFlagId
				insertedSenderData, trxDataId, err := b.rprFraudDetection.WriteRowBlacklistSender(
					ctx, rprFraudDetection.EntityBlacklistSender{
						Id:               guuid.NewString(),
						PhoneNumber:      request.BenefactorIdentityNumber,
						BeneficiaryName:  request.BenefactorName,
						CreatedAt:        now.Format(basicObject.DateAndTime),
						CreatedBy:        request.Channel,
						UpdatedAt:        now.Format(basicObject.DateAndTime),
						UpdatedBy:        request.Channel,
						Status:           1,
						TransactionTypes: strings.Join(transactionType, ","),
					},
				)
				if err != nil {
					errMsg := fmt.Errorf("internal server error")
					logging.Errorw(fName, "reason", err.Error())
					return nil, errWrap.WrapRepositoryError(errMsg)
				}
				if insertedSenderData == nil && trxDataId == nil {
					errMsg := fmt.Errorf("internal server error")
					logging.Errorw(fName, "reason", "failed to insert new blacklist sender data")
					return nil, errWrap.WrapRepositoryError(errMsg)
				}
			}
			response.ResponseObject.Reported = basicObject.WithReport
			//if senderResult != nil && len(*senderResult) == 0 {
			//	insertedSenderData, trxDataId, err := b.rprFraudDetection.WriteRowBlacklistSender(
			//		ctx, rprFraudDetection.EntityBlacklistSender{
			//			Id:               guuid.NewString(),
			//			PhoneNumber:      request.BenefactorIdentityNumber,
			//			BeneficiaryName:  request.BenefactorName,
			//			CreatedAt:        now.Format(basicObject.DateAndTime),
			//			CreatedBy:        request.Channel,
			//			UpdatedAt:        now.Format(basicObject.DateAndTime),
			//			UpdatedBy:        request.Channel,
			//			Status:           1,
			//			TransactionTypes: request.TransactionType,
			//		},
			//	)
			//	if err != nil {
			//		errMsg := fmt.Errorf("internal server error")
			//		logging.Errorw(fName, "reason", err.Error())
			//		return nil, errWrap.WrapRepositoryError(errMsg)
			//	}
			//	if insertedSenderData == nil && trxDataId == nil {
			//		errMsg := fmt.Errorf("internal server error")
			//		logging.Errorw(fName, "reason", "failed to insert new blacklist sender data")
			//		return nil, errWrap.WrapRepositoryError(errMsg)
			//	}
			//}
		}
	} else {
		response.ResponseObject.Conclusion = basicObject.Allowed
	}
	if reported {
		// Inserting trx data
		response.ResponseObject.Reported = basicObject.WithReport
		if !rejected {
			transactionSevereLevelCategory = basicObject.WarningFlagId
		}
		if senderResult != nil && len(*senderResult) == 0 {
			insertedReportData, trxDataId, err := b.rprFraudDetection.WriteRowBlacklistHistory(
				ctx, rprFraudDetection.EntityBlacklistHistory{
					Id:               guuid.NewString(),
					CreatedAt:        now.Format(basicObject.DateAndTime),
					Event:            basicObject.Detection,
					CreatedBy:        request.Channel,
					PhoneNumber:      request.BenefactorIdentityNumber,
					BeneficiaryName:  request.BenefactorName,
					TransactionTypes: request.TransactionType,
				},
			)
			if err != nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", err.Error())
				return nil, errWrap.WrapRepositoryError(errMsg)
			}
			if insertedReportData == nil && trxDataId == nil {
				errMsg := fmt.Errorf("internal server error")
				logging.Errorw(fName, "reason", "failed to insert report data")
				return nil, errWrap.WrapRepositoryError(errMsg)
			}
		}
	} else {
		if defaultAct == basicObject.GovActReject {
			response.ResponseObject.Reported = basicObject.WithoutReport
		}
		if defaultAct == basicObject.GovActAllow {
			if hasDTTOTRecord {
				response.ResponseObject.Reported = basicObject.WithReport
			} else {
				response.ResponseObject.Reported = basicObject.WithoutReport
			}
		}
	}

	// Update transaction history data
	ruleBytes, _ := json.Marshal(rulesEntityEvaluation)
	insertedTrxData, trxDataId, err := b.rprFraudDetection.UpdateRowTransaction(
		ctx, rprFraudDetection.EntityTransaction{
			TransactionId:   partnerReferenceNo,
			Rules:           string(ruleBytes),
			Title:           request.TransactionType,
			FlagId:          transactionSevereLevelCategory,
			CreatedAt:       now.Format(basicObject.DateAndTime),
			Amount:          strconv.FormatFloat(requestAmount, 'f', 2, 64),
			Id:              transactionId,
			UserId:          request.BenefactorIdentityNumber,
			DestinationId:   request.BeneficiaryIdentityNumber,
			BodyReq:         string(additionalInfoBytes),
			Channel:         request.Channel,
			TransactionType: request.TransactionType,
		},
	)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if insertedTrxData == nil && trxDataId == nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", "failed to insert trx data")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response code and description extraction
	responseCode := successResponseCode
	responseDescription := successResponseDescription

	// returning result data
	response.RequestId = requestId
	response.ResponseCode = responseCode
	response.ResponseDescription = responseDescription
	response.ResponseObject.EvaluationData = evaluationData
	response.TransactionId = partnerReferenceNo
	if insertedLogData != nil {
		response.LogId = insertedLogData.Id
	}
	return response, nil
}
