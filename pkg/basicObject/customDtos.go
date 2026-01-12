package basicObject

type TypeAndValue1 struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type TypeAndValue2 struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Actions struct {
	IsReport bool `json:"is_report"`
	IsBlock  bool `json:"is_block"`
}

type RuleData struct {
	RuleId          string `json:"ruleId"`
	IsRejected      bool   `json:"isRejected"`
	IsReported      bool   `json:"isReported"`
	TransactionType string `json:"transactionType"`
	RuleType        string `json:"ruleType"`
	RuleName        string `json:"ruleName"`
	Violation       string `json:"violation"`
}

type BlacklistDTTOTData struct {
	FileId     string `json:"fileId"`
	FileLink   string `json:"fileLink"`
	Datasource string `json:"datasource"`
	Name       string `json:"name"`
	NIK        string `json:"nik"`
	Performer  string `json:"performer"`
}

type BlacklistMerchantData struct {
	Name       string `json:"name"`
	NMID       string `json:"nmid"`
	FileId     string `json:"fileId"`
	FileLink   string `json:"fileLink"`
	Datasource string `json:"datasource"`
}

type BlacklistReceiverData struct {
	Name     string `json:"name"`
	Identity string `json:"identity"`
}

type BlacklistSenderData struct {
	Name     string `json:"name"`
	Identity string `json:"identity"`
}

type KeywordData struct {
	KeywordId       string `json:"keywordID"`
	KeywordViolated string `json:"keywordViolated"`
	Allowed         bool   `json:"allowed"`
	Report          bool   `json:"report"`
}

type EvaluationData struct {
	Rules             []RuleData              `json:"rules"`
	Keywords          []KeywordData           `json:"keywords"`
	BlacklistDTTOT    []BlacklistDTTOTData    `json:"blacklistDTTOT"`
	BlacklistMerchant []BlacklistMerchantData `json:"blacklistMerchant"`
	BlacklistReceiver []BlacklistReceiverData `json:"blacklistReceiver"`
	BlacklistSender   []BlacklistSenderData   `json:"blacklistSender"`
}
