package uscFraud

/* [CODE GENERATOR] INITIALIZE_CODE */

type ResponseDetect struct {
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
	RequestId           string `json:"requestId"`
	LogId               string `json:"logId"`
	TransactionId       string `json:"transactionId"`
	ResponseObject      struct {
		Conclusion     string      `json:"conclusion"`
		Reported       string      `json:"reported"`
		EvaluationData interface{} `json:"evaluationData"`
	} `json:"responseObject"`
}

type RequestDetect struct {
	BenefactorIdentityNumber  string `json:"benefactorIdentityNumber" validate:"required"`
	BenefactorName            string `json:"benefactorName" validate:"required"`
	BeneficiaryIdentityNumber string `json:"beneficiaryIdentityNumber" validate:"required"`
	BeneficiaryName           string `json:"beneficiaryName" validate:"required"`
	PartnerReferenceNumber    string `json:"partnerReferenceNumber" validate:"required"`
	TransactionType           string `json:"transactionType" validate:"required"`
	TransactionTime           string `json:"transactionTime" validate:"required"`
	Channel                   string `json:"channel" validate:"required"`
	Sof                       string `json:"sof" validate:"required"`
	Amount                    struct {
		Value    string `json:"value" validate:"required"`
		Currency string `json:"currency" validate:"required"`
	} `json:"amount" validate:"required"`
	AdditionalInfo interface{} `json:"additionalInfo,omitempty"`
}
