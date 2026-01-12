package endpoints

/* [CODE GENERATOR] IMPORT_PKG_USC */
import (
	"github.com/labstack/echo/v4"
	"gitlab.com/fds22/detection-sys/pkg/customController"
	"gitlab.com/fds22/detection-sys/pkg/customEcho"
	"gitlab.com/fds22/detection-sys/pkg/customMiddleware"
	"gitlab.com/fds22/detection-sys/pkg/customValidator"
	"gitlab.com/fds22/detection-sys/pkg/environments"
	"gitlab.com/fds22/detection-sys/src/repositories/rprAuthentication"
	"gitlab.com/fds22/detection-sys/src/usecases/uscAuthentication"
	"gitlab.com/fds22/detection-sys/src/usecases/uscBlacklistManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscChannelManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscFraud"
	"gitlab.com/fds22/detection-sys/src/usecases/uscKeywordManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscRuleManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscSofManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscTrxTypeManagement"
	"go.uber.org/zap"
)

func NewEndpoint(
	env *environments.Envs,
	log *zap.SugaredLogger,
	vld *customValidator.CustomValidator,
	rprAuthentication rprAuthentication.Blueprint,
	/* [CODE GENERATOR] FUNC_PARAM */
	/* [CODE GENERATOR] FUNC_PARAM_uscFraudInjected */
	uscFraudInjected uscFraud.Blueprint,
	/* [CODE GENERATOR] FUNC_PARAM */
	/* [CODE GENERATOR] FUNC_PARAM_uscBlacklistManagementInjected */
	uscBlacklistManagementInjected uscBlacklistManagement.Blueprint,
	/* [CODE GENERATOR] FUNC_PARAM */
	/* [CODE GENERATOR] FUNC_PARAM_uscRuleManagementInjected */
	uscRuleManagementInjected uscRuleManagement.Blueprint,
	/* [CODE GENERATOR] FUNC_PARAM */
	/* [CODE GENERATOR] FUNC_PARAM_uscKeywordManagementInjected */
	uscKeywordManagementInjected uscKeywordManagement.Blueprint,
	/* [CODE GENERATOR] FUNC_PARAM */
	/* [CODE GENERATOR] FUNC_PARAM_uscTrxTypeManagementInjected */
	uscTrxTypeManagementInjected uscTrxTypeManagement.Blueprint,
	/* [CODE GENERATOR] FUNC_PARAM */
	/* [CODE GENERATOR] FUNC_PARAM_uscAuthenticationInjected */
	uscAuthenticationInjected uscAuthentication.Blueprint,
	/* [CODE GENERATOR] FUNC_PARAM */
	/* [CODE GENERATOR] FUNC_PARAM_uscSofManagementInjected */
	uscSofManagementInjected uscSofManagement.Blueprint,
	/* [CODE GENERATOR] FUNC_PARAM */
	/* [CODE GENERATOR] FUNC_PARAM_uscChannelManagementInjected */
	uscChannelManagementInjected uscChannelManagement.Blueprint,
) *echo.Echo {
	ech := customEcho.NewCustomEcho(log, vld)

	// master endpoints
	master := ech.Group("/api/v1")

	/* [CODE GENERATOR] ENDPOINT */
	/* [CODE GENERATOR] ENDPOINT_GROUP_everything */
	everythingEp := master.Group("/fds")

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_Detect */
	everythingEp.POST("/detect-eligibility", customController.RequestController(uscFraudInjected.Detect, uscFraud.RequestDetect{}), customMiddleware.AuthBasic(env))

	/* [CODE GENERATOR] ENDPOINT */
	/* [CODE GENERATOR] ENDPOINT_GROUP_management */
	managementEp := everythingEp.Group("/management")

	/* [CODE GENERATOR] ENDPOINT */
	/* [CODE GENERATOR] ENDPOINT_GROUP_authentication */
	authenticationEp := everythingEp.Group("/authentication")

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_Authentication */
	authenticationEp.POST("/register", customController.RequestController(uscAuthenticationInjected.UserRegister, uscAuthentication.RequestUserRegister{}))
	authenticationEp.POST("/login", customController.RequestController(uscAuthenticationInjected.UserLogin, uscAuthentication.RequestLoginRegister{}))
	authenticationEp.POST("/banned/:id", customController.RequestController(uscAuthenticationInjected.UserBanned, uscAuthentication.RequestBannedUser{}))
	authenticationEp.GET("/roles", customController.RequestController(uscAuthenticationInjected.GetAvailableRoles, uscAuthentication.RequestGetRoles{}))
	authenticationEp.GET("/registered/users", customController.RequestController(uscAuthenticationInjected.GetRegisteredUsers, uscAuthentication.RequestRegisteredUsers{}))
	authenticationEp.POST("/reset/password", customController.RequestController(uscAuthenticationInjected.ResetPassword, uscAuthentication.RequestResetPassword{}))
	authenticationEp.POST("/reset/password/submit", customController.RequestController(uscAuthenticationInjected.ResetPasswordSubmit, uscAuthentication.RequestResetPasswordSubmit{}))
	authenticationEp.GET("/reset/password/validate/:token", customController.RequestController(uscAuthenticationInjected.ResetPasswordValidate, uscAuthentication.RequestResetPasswordValidate{}))
	authenticationEp.POST("/change/password", customController.RequestController(uscAuthenticationInjected.ChangePassword, uscAuthentication.RequestChangePassword{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementBlacklistHistory */
	blacklistHistoryEp := managementEp.Group("/blacklist-history")
	blacklistHistoryEp.GET("/retrieve", customController.RequestController(uscBlacklistManagementInjected.BlacklistHistories, uscBlacklistManagement.RequestRetrieveBlacklistHistory{}), customMiddleware.AuthJWT(env, rprAuthentication))
	blacklistHistoryEp.GET("/download", customController.DownloadController(uscBlacklistManagementInjected.BlacklistReporting, uscBlacklistManagement.RequestRetrieveBlacklistReporting{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementTransactionHistory */
	transactionHistoryEp := managementEp.Group("/transaction-history")
	transactionHistoryEp.GET("/retrieve", customController.RequestController(uscBlacklistManagementInjected.TransactionHistories, uscBlacklistManagement.RequestRetrieveTransactionHistory{}), customMiddleware.AuthJWT(env, rprAuthentication))
	transactionHistoryEp.GET("/download", customController.DownloadController(uscBlacklistManagementInjected.TransactionReporting, uscBlacklistManagement.RequestRetrieveTransactionReporting{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementSender */
	senderEp := managementEp.Group("/sender")
	senderEp.POST("/create", customController.RequestController(uscBlacklistManagementInjected.BlacklistSenderCreate, uscBlacklistManagement.RequestCreateBlacklistSender{}), customMiddleware.AuthJWT(env, rprAuthentication))
	senderEp.DELETE("/delete/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistSenderDelete, uscBlacklistManagement.RequestDeleteBlacklistSender{}), customMiddleware.AuthJWT(env, rprAuthentication))
	senderEp.POST("/approve/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistSenderApprove, uscBlacklistManagement.RequestApproveBlacklistSender{}), customMiddleware.AuthJWT(env, rprAuthentication))
	senderEp.POST("/reject/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistSenderReject, uscBlacklistManagement.RequestRejectBlacklistSender{}), customMiddleware.AuthJWT(env, rprAuthentication))
	senderEp.GET("/retrieve-approval", customController.RequestController(uscBlacklistManagementInjected.BlacklistSenderApprovalGet, uscBlacklistManagement.RequestRetrieveBlacklistSenderApproval{}), customMiddleware.AuthJWT(env, rprAuthentication))
	senderEp.GET("/retrieve-approved", customController.RequestController(uscBlacklistManagementInjected.BlacklistSenderApprovedGet, uscBlacklistManagement.RequestRetrieveBlacklistSenderApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))
	senderEp.GET("/retrieve-approval/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistSenderApprovalGetById, uscBlacklistManagement.RequestRetrieveBlacklistSenderById{}), customMiddleware.AuthJWT(env, rprAuthentication))
	senderEp.GET("/retrieve-approved/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistSenderApprovedGetById, uscBlacklistManagement.RequestRetrieveBlacklistSenderById{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementReceiver */
	receiverEp := managementEp.Group("/receiver")
	receiverEp.POST("/create", customController.RequestController(uscBlacklistManagementInjected.BlacklistReceiverCreate, uscBlacklistManagement.RequestCreateBlacklistReceiver{}), customMiddleware.AuthJWT(env, rprAuthentication))
	receiverEp.DELETE("/delete/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistReceiverDelete, uscBlacklistManagement.RequestDeleteBlacklistReceiver{}), customMiddleware.AuthJWT(env, rprAuthentication))
	receiverEp.POST("/approve/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistReceiverApprove, uscBlacklistManagement.RequestApproveBlacklistReceiver{}), customMiddleware.AuthJWT(env, rprAuthentication))
	receiverEp.POST("/reject/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistReceiverReject, uscBlacklistManagement.RequestRejectBlacklistReceiver{}), customMiddleware.AuthJWT(env, rprAuthentication))
	receiverEp.GET("/retrieve-approval", customController.RequestController(uscBlacklistManagementInjected.BlacklistReceiverApprovalGet, uscBlacklistManagement.RequestRetrieveBlacklistReceiverApproval{}), customMiddleware.AuthJWT(env, rprAuthentication))
	receiverEp.GET("/retrieve-approved", customController.RequestController(uscBlacklistManagementInjected.BlacklistReceiverApprovedGet, uscBlacklistManagement.RequestRetrieveBlacklistReceiverApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))
	receiverEp.GET("/retrieve-approval/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistReceiverApprovalGetById, uscBlacklistManagement.RequestRetrieveBlacklistReceiverById{}), customMiddleware.AuthJWT(env, rprAuthentication))
	receiverEp.GET("/retrieve-approved/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistReceiverApprovedGetById, uscBlacklistManagement.RequestRetrieveBlacklistReceiverById{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementDTTOT */
	dttotEp := managementEp.Group("/dttot")
	dttotEp.POST("/upload", customController.UploadController(uscBlacklistManagementInjected.BlacklistDTTOTUpload, uscBlacklistManagement.RequestUploadBlacklistDTTOT{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.POST("/approve/:fileId", customController.RequestController(uscBlacklistManagementInjected.BlacklistDTTOTApprove, uscBlacklistManagement.RequestApproveBlacklistDTTOT{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.POST("/reject/:fileId", customController.RequestController(uscBlacklistManagementInjected.BlacklistDTTOTReject, uscBlacklistManagement.RequestRejectBlacklistDTTOT{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.POST("/retrieve-approval", customController.RequestController(uscBlacklistManagementInjected.BlacklistDTTOTApprovalGet, uscBlacklistManagement.RequestRetrieveBlacklistDTTOTApproval{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/retrieve-approved", customController.DownloadController(uscBlacklistManagementInjected.BlacklistDTTOTApprovedGet, uscBlacklistManagement.RequestRetrieveBlacklistDTTOTApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/template", customController.DownloadController(uscBlacklistManagementInjected.BlacklistDTTOTApprovedTemplate, uscBlacklistManagement.RequestRetrieveBlacklistDTTOTTemplate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/files", customController.RequestController(uscBlacklistManagementInjected.BlacklistDTTOTActiveData, uscBlacklistManagement.RequestRetrieveActiveDTTOT{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/contents", customController.RequestController(uscBlacklistManagementInjected.BlacklistDTTOTActiveDataContent, uscBlacklistManagement.RequestRetrieveActiveDTTOTContent{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/contents/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistDTTOTActiveDataContentById, uscBlacklistManagement.RequestRetrieveActiveDTTOTContentById{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/contents-approved/:id", customController.DownloadController(uscBlacklistManagementInjected.BlacklistDTTOTApprovedDownloadContentByFileId, uscBlacklistManagement.RequestDTTOTDownloadContentByFileId{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/contents-approval/:id", customController.DownloadController(uscBlacklistManagementInjected.BlacklistDTTOTApprovalDownloadContentByFileId, uscBlacklistManagement.RequestDTTOTDownloadContentByFileId{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/file/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistDTTOTFileById, uscBlacklistManagement.RequestBlacklistDTTOTFileById{}), customMiddleware.AuthJWT(env, rprAuthentication))
	dttotEp.GET("/file-content/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistDTTOTFileContentById, uscBlacklistManagement.RequestBlacklistDTTOTFileContent{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementMerchant */
	merchantEp := managementEp.Group("/merchant")
	merchantEp.POST("/upload", customController.UploadController(uscBlacklistManagementInjected.BlacklistMerchantUpload, uscBlacklistManagement.RequestUploadBlacklistMerchant{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.POST("/approve/:fileId", customController.RequestController(uscBlacklistManagementInjected.BlacklistMerchantApprove, uscBlacklistManagement.RequestApproveBlacklistMerchant{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.POST("/reject/:fileId", customController.RequestController(uscBlacklistManagementInjected.BlacklistMerchantReject, uscBlacklistManagement.RequestRejectBlacklistMerchant{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.POST("/retrieve-approval", customController.RequestController(uscBlacklistManagementInjected.BlacklistMerchantApprovalGet, uscBlacklistManagement.RequestRetrieveBlacklistMerchantApproval{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/retrieve-approved", customController.DownloadController(uscBlacklistManagementInjected.BlacklistMerchantApprovedGet, uscBlacklistManagement.RequestRetrieveBlacklistMerchantApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/template", customController.DownloadController(uscBlacklistManagementInjected.BlacklistMerchantTemplate, uscBlacklistManagement.RequestRetrieveBlacklistMerchantTemplate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/files", customController.RequestController(uscBlacklistManagementInjected.BlacklistMerchantActiveData, uscBlacklistManagement.RequestRetrieveActiveMerchant{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/contents", customController.RequestController(uscBlacklistManagementInjected.BlacklistMerchantActiveDataContent, uscBlacklistManagement.RequestRetrieveActiveMerchantContent{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/contents/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistMerchantActiveDataContentById, uscBlacklistManagement.RequestRetrieveActiveMerchantContentById{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/contents-approved/:id", customController.DownloadController(uscBlacklistManagementInjected.BlacklistMerchantApprovedDownloadContentByFileId, uscBlacklistManagement.RequestMerchantDownloadContentByFileId{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/contents-approval/:id", customController.DownloadController(uscBlacklistManagementInjected.BlacklistMerchantApprovalDownloadContentByFileId, uscBlacklistManagement.RequestMerchantDownloadContentByFileId{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/file/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistMerchantFileById, uscBlacklistManagement.RequestBlacklistMerchantFileById{}), customMiddleware.AuthJWT(env, rprAuthentication))
	merchantEp.GET("/file-content/:id", customController.RequestController(uscBlacklistManagementInjected.BlacklistMerchantFileContentById, uscBlacklistManagement.RequestBlacklistMerchantFileContent{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementRule */
	ruleEp := managementEp.Group("/rule")
	ruleEp.POST("/create", customController.RequestController(uscRuleManagementInjected.RuleDetectionCreate, uscRuleManagement.RequestRuleDetectionCreate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.DELETE("/delete/:id", customController.RequestController(uscRuleManagementInjected.RuleDetectionDelete, uscRuleManagement.RequestRuleDetectionDelete{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.PUT("/update/:id", customController.RequestController(uscRuleManagementInjected.RuleDetectionUpdate, uscRuleManagement.RequestRuleDetectionUpdate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.POST("/approve/:id", customController.RequestController(uscRuleManagementInjected.RuleDetectionApprove, uscRuleManagement.RequestRuleDetectionApprove{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.POST("/reject/:id", customController.RequestController(uscRuleManagementInjected.RuleDetectionReject, uscRuleManagement.RequestRuleDetectionReject{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.GET("/retrieve-approval", customController.RequestController(uscRuleManagementInjected.RuleDetectionGetApproval, uscRuleManagement.RequestRuleDetectionGetApproval{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.GET("/retrieve-approved", customController.RequestController(uscRuleManagementInjected.RuleDetectionGetApproved, uscRuleManagement.RequestRuleDetectionGetApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.GET("/retrieve-approval/:id", customController.RequestController(uscRuleManagementInjected.RuleDetectionGetApprovalById, uscRuleManagement.RequestRuleApprovalById{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.GET("/retrieve-approved/:id", customController.RequestController(uscRuleManagementInjected.RuleDetectionGetApprovedById, uscRuleManagement.RequestRuleApprovedById{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.POST("/enabled/:id", customController.RequestController(uscRuleManagementInjected.RuleDetectionEnable, uscRuleManagement.RequestRuleEnabled{}), customMiddleware.AuthJWT(env, rprAuthentication))
	ruleEp.POST("/disabled/:id", customController.RequestController(uscRuleManagementInjected.RuleDetectionDisable, uscRuleManagement.RequestRuleDisabled{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementKeyword */
	keywordEp := managementEp.Group("/keyword")
	keywordEp.POST("/create", customController.RequestController(uscKeywordManagementInjected.KeywordCreate, uscKeywordManagement.RequestKeywordCreate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	keywordEp.POST("/delete/:id", customController.RequestController(uscKeywordManagementInjected.KeywordDelete, uscKeywordManagement.RequestKeywordDelete{}), customMiddleware.AuthJWT(env, rprAuthentication))
	keywordEp.PUT("/update", customController.RequestController(uscKeywordManagementInjected.KeywordUpdate, uscKeywordManagement.RequestKeywordUpdate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	keywordEp.POST("/approve/:id", customController.RequestController(uscKeywordManagementInjected.KeywordApprove, uscKeywordManagement.RequestKeywordApprove{}), customMiddleware.AuthJWT(env, rprAuthentication))
	keywordEp.POST("/reject/:id", customController.RequestController(uscKeywordManagementInjected.KeywordReject, uscKeywordManagement.RequestKeywordReject{}), customMiddleware.AuthJWT(env, rprAuthentication))
	keywordEp.GET("/retrieve-detail/:id", customController.RequestController(uscKeywordManagementInjected.KeywordGetApprovalDetail, uscKeywordManagement.RequestKeywordGetApprovalDetail{}), customMiddleware.AuthJWT(env, rprAuthentication))
	keywordEp.GET("/retrieve-approval", customController.RequestController(uscKeywordManagementInjected.KeywordGetApproval, uscKeywordManagement.RequestKeywordGetApproval{}), customMiddleware.AuthJWT(env, rprAuthentication))
	keywordEp.GET("/retrieve-approved", customController.RequestController(uscKeywordManagementInjected.KeywordGetApproved, uscKeywordManagement.RequestKeywordGetApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementTrxType */
	trxTypeEp := managementEp.Group("/transaction-type")
	trxTypeEp.POST("/create", customController.RequestController(uscTrxTypeManagementInjected.TransactionTypeCreate, uscTrxTypeManagement.RequestTransactionTypeCreate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	trxTypeEp.DELETE("/delete/:id", customController.RequestController(uscTrxTypeManagementInjected.TransactionTypeDelete, uscTrxTypeManagement.RequestTransactionTypeDelete{}), customMiddleware.AuthJWT(env, rprAuthentication))
	trxTypeEp.PUT("/update/:id", customController.RequestController(uscTrxTypeManagementInjected.TransactionTypeUpdate, uscTrxTypeManagement.RequestTransactionTypeUpdate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	trxTypeEp.GET("/retrieve", customController.RequestController(uscTrxTypeManagementInjected.TransactionTypeGetApproved, uscTrxTypeManagement.RequestTransactionTypeGetApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))
	trxTypeEp.GET("/retrieve/all", customController.RequestController(uscTrxTypeManagementInjected.TransactionTypeGetApprovedAll, uscTrxTypeManagement.RequestTransactionTypeGetApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementSof */
	sofEp := managementEp.Group("/sof")
	sofEp.POST("/create", customController.RequestController(uscSofManagementInjected.SOFCreate, uscSofManagement.RequestSOFCreate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	sofEp.DELETE("/delete/:id", customController.RequestController(uscSofManagementInjected.SOFDelete, uscSofManagement.RequestSOFDelete{}), customMiddleware.AuthJWT(env, rprAuthentication))
	sofEp.PUT("/update/:id", customController.RequestController(uscSofManagementInjected.SOFUpdate, uscSofManagement.RequestSOFUpdate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	sofEp.GET("/retrieve", customController.RequestController(uscSofManagementInjected.SOFGetApproved, uscSofManagement.RequestSOFGetApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))
	sofEp.GET("/retrieve/all", customController.RequestController(uscSofManagementInjected.SOFGetApprovedAll, uscSofManagement.RequestSOFGetApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))

	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT */
	/* [CODE GENERATOR] ENDPOINT_ASSIGNMENT_everything_ManagementChannel */
	channelEp := managementEp.Group("/channel")
	channelEp.POST("/create", customController.RequestController(uscChannelManagementInjected.ChannelCreate, uscChannelManagement.RequestChannelCreate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	channelEp.DELETE("/delete/:id", customController.RequestController(uscChannelManagementInjected.ChannelDelete, uscChannelManagement.RequestChannelDelete{}), customMiddleware.AuthJWT(env, rprAuthentication))
	channelEp.PUT("/update/:id", customController.RequestController(uscChannelManagementInjected.ChannelUpdate, uscChannelManagement.RequestChannelUpdate{}), customMiddleware.AuthJWT(env, rprAuthentication))
	channelEp.GET("/retrieve", customController.RequestController(uscChannelManagementInjected.ChannelGetApproved, uscChannelManagement.RequestChannelGetApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))
	channelEp.GET("/retrieve/all", customController.RequestController(uscChannelManagementInjected.ChannelGetApprovedAll, uscChannelManagement.RequestChannelGetApproved{}), customMiddleware.AuthJWT(env, rprAuthentication))

	return ech
}
