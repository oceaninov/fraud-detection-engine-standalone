package basicObject

const (
	InsertEvent        = "insert"
	RemoveEvent        = "remove"
	CreateApprovalType = "Create"
	DeleteApprovalType = "Delete"
	UpdateApprovalType = "Update"
)

const (
	Benefactor  = "SENDER"
	Beneficiary = "RECEIVER"
)

const (
	Rejected      = "REJECTED"
	Allowed       = "ALLOWED"
	WithoutReport = "WITHOUT_REPORT"
	WithReport    = "WITH_REPORT"
)

const (
	Multiple    = "multiple"
	Single      = "single"
	Minute      = "minute"
	Second      = "second"
	Interval    = "interval"
	Transaction = "transaction"
)

const (
	DateOnly            = "2006-01-02"
	TimeOnly            = "15:04:05"
	DateAndTime         = "2006-01-02 15:04:05"
	DatabaseDateAndTime = "2006-01-02T15:04:05Z"
)

const (
	DateOnlyEnum    = "DO"
	TimeOnlyEnum    = "TO"
	DateAndTimeEnum = "DNT"
	NoneEnum        = "NONE"
)

const (
	AllowedFlagId  = "e2077d0f-aaf5-4520-95d7-e4b097d2f3a5"
	WarningFlagId  = "aedcac7f-064b-43d7-b866-47a2875eec63"
	RejectedFlagId = "b53ce142-b2fd-4634-a964-801dd7bfaf3b"
)

const (
	Detection = "detection"
)

const (
	KeywordActionAllowWithReport = "allow-with-report"
	KeywordActionBlockWithReport = "block-with-report"
)

const (
	GovActAllow  = "allow"
	GovActReject = "reject"
)

const (
	BlankString       = ""
	Successfully      = "Successfully"
	SuccessfullyCode  = "00"
	SuccessfullyTrue  = true
	SuccessfullyFalse = false
)

const (
	ApprovalPending  = "1"
	ApprovalRejected = "3"
	ApprovalApproved = "2"
)
const (
	RoleChecker = "3"
	RoleMaker   = "2"
)
