package uscBlacklistManagement

import (
	"context"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistDTTOT"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistHistory"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistMerchant"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistReceiver"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistSender"
	"gitlab.com/fds22/detection-sys/src/repositories/rprTransaction"
	"gitlab.com/fds22/detection-sys/src/repositories/rprTransactionType"
	"go.uber.org/zap"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint business function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		BlacklistSenderCreate(ctx context.Context, request *RequestCreateBlacklistSender) (*ResponseCreateBlacklistSender, error)
		BlacklistSenderDelete(ctx context.Context, request *RequestDeleteBlacklistSender) (*ResponseDeleteBlacklistSender, error)
		BlacklistSenderReject(ctx context.Context, request *RequestRejectBlacklistSender) (*ResponseRejectBlacklistSender, error)
		BlacklistSenderApprove(ctx context.Context, request *RequestApproveBlacklistSender) (*ResponseApproveBlacklistSender, error)
		BlacklistSenderApprovalGet(ctx context.Context, request *RequestRetrieveBlacklistSenderApproval) (*ResponseRetrieveBlacklistSenderApproval, error)
		BlacklistSenderApprovedGet(ctx context.Context, request *RequestRetrieveBlacklistSenderApproved) (*ResponseRetrieveBlacklistSenderApproved, error)
		BlacklistSenderApprovalGetById(ctx context.Context, request *RequestRetrieveBlacklistSenderById) (*ResponseRetrieveBlacklistSenderApprovalById, error)
		BlacklistSenderApprovedGetById(ctx context.Context, request *RequestRetrieveBlacklistSenderById) (*ResponseRetrieveBlacklistSenderApprovedById, error)

		BlacklistReceiverCreate(ctx context.Context, request *RequestCreateBlacklistReceiver) (*ResponseCreateBlacklistReceiver, error)
		BlacklistReceiverDelete(ctx context.Context, request *RequestDeleteBlacklistReceiver) (*ResponseDeleteBlacklistReceiver, error)
		BlacklistReceiverReject(ctx context.Context, request *RequestRejectBlacklistReceiver) (*ResponseRejectBlacklistReceiver, error)
		BlacklistReceiverApprove(ctx context.Context, request *RequestApproveBlacklistReceiver) (*ResponseApproveBlacklistReceiver, error)
		BlacklistReceiverApprovalGet(ctx context.Context, request *RequestRetrieveBlacklistReceiverApproval) (*ResponseRetrieveBlacklistReceiverApproval, error)
		BlacklistReceiverApprovedGet(ctx context.Context, request *RequestRetrieveBlacklistReceiverApproved) (*ResponseRetrieveBlacklistReceiverApproved, error)
		BlacklistReceiverApprovalGetById(ctx context.Context, request *RequestRetrieveBlacklistReceiverById) (*ResponseRetrieveBlacklistReceiverApprovalById, error)
		BlacklistReceiverApprovedGetById(ctx context.Context, request *RequestRetrieveBlacklistReceiverById) (*ResponseRetrieveBlacklistReceiverApprovedById, error)

		BlacklistDTTOTUpload(ctx context.Context, request *RequestUploadBlacklistDTTOT) (*ResponseUploadBlacklistDTTOT, error)
		BlacklistDTTOTReject(ctx context.Context, request *RequestRejectBlacklistDTTOT) (*ResponseRejectBlacklistDTTOT, error)
		BlacklistDTTOTApprove(ctx context.Context, request *RequestApproveBlacklistDTTOT) (*ResponseApproveBlacklistDTTOT, error)
		BlacklistDTTOTApprovalGet(ctx context.Context, request *RequestRetrieveBlacklistDTTOTApproval) (*ResponseRetrieveBlacklistDTTOTApproval, error)
		BlacklistDTTOTApprovedGet(ctx context.Context, request *RequestRetrieveBlacklistDTTOTApproved) ([]byte, error)
		BlacklistDTTOTApprovedTemplate(ctx context.Context, request *RequestRetrieveBlacklistDTTOTTemplate) ([]byte, error)
		BlacklistDTTOTActiveData(ctx context.Context, request *RequestRetrieveActiveDTTOT) (*ResponseRetrieveActiveDTTOT, error)
		BlacklistDTTOTActiveDataContent(ctx context.Context, request *RequestRetrieveActiveDTTOTContent) (*ResponseRetrieveActiveDTTOTContent, error)
		BlacklistDTTOTActiveDataContentById(ctx context.Context, request *RequestRetrieveActiveDTTOTContentById) (*ResponseRetrieveActiveDTTOTContentById, error)
		BlacklistDTTOTApprovedDownloadContentByFileId(ctx context.Context, request *RequestDTTOTDownloadContentByFileId) ([]byte, error)
		BlacklistDTTOTApprovalDownloadContentByFileId(ctx context.Context, request *RequestDTTOTDownloadContentByFileId) ([]byte, error)
		BlacklistDTTOTFileById(ctx context.Context, request *RequestBlacklistDTTOTFileById) (*ResponseBlacklistDTTOTFileById, error)
		BlacklistDTTOTFileContentById(ctx context.Context, request *RequestBlacklistDTTOTFileContent) (*ResponseBlacklistDTTOTFileContent, error)

		BlacklistMerchantUpload(ctx context.Context, request *RequestUploadBlacklistMerchant) (*ResponseUploadBlacklistMerchant, error)
		BlacklistMerchantReject(ctx context.Context, request *RequestRejectBlacklistMerchant) (*ResponseRejectBlacklistMerchant, error)
		BlacklistMerchantApprove(ctx context.Context, request *RequestApproveBlacklistMerchant) (*ResponseApproveBlacklistMerchant, error)
		BlacklistMerchantApprovalGet(ctx context.Context, request *RequestRetrieveBlacklistMerchantApproval) (*ResponseRetrieveBlacklistMerchantApproval, error)
		BlacklistMerchantApprovedGet(ctx context.Context, request *RequestRetrieveBlacklistMerchantApproved) ([]byte, error)
		BlacklistMerchantTemplate(ctx context.Context, request *RequestRetrieveBlacklistMerchantTemplate) ([]byte, error)
		BlacklistMerchantActiveData(ctx context.Context, request *RequestRetrieveActiveMerchant) (*ResponseRetrieveActiveMerchant, error)
		BlacklistMerchantActiveDataContent(ctx context.Context, request *RequestRetrieveActiveMerchantContent) (*ResponseRetrieveActiveMerchantContent, error)
		BlacklistMerchantActiveDataContentById(ctx context.Context, request *RequestRetrieveActiveMerchantContentById) (*ResponseRetrieveActiveMerchantContentById, error)
		BlacklistMerchantApprovedDownloadContentByFileId(ctx context.Context, request *RequestMerchantDownloadContentByFileId) ([]byte, error)
		BlacklistMerchantApprovalDownloadContentByFileId(ctx context.Context, request *RequestMerchantDownloadContentByFileId) ([]byte, error)
		BlacklistMerchantFileById(ctx context.Context, request *RequestBlacklistMerchantFileById) (*ResponseBlacklistMerchantFileById, error)
		BlacklistMerchantFileContentById(ctx context.Context, request *RequestBlacklistMerchantFileContent) (*ResponseBlacklistMerchantFileContent, error)

		BlacklistHistories(ctx context.Context, request *RequestRetrieveBlacklistHistory) (*ResponseRetrieveBlacklistHistory, error)
		BlacklistReporting(ctx context.Context, request *RequestRetrieveBlacklistReporting) ([]byte, error)
		TransactionHistories(ctx context.Context, request *RequestRetrieveTransactionHistory) (*ResponseRetrieveTransactionHistory, error)
		TransactionReporting(ctx context.Context, request *RequestRetrieveTransactionReporting) ([]byte, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		rprBlacklistReceiver rprBlacklistReceiver.Blueprint
		rprBlacklistSender   rprBlacklistSender.Blueprint
		rprBlacklistMerchant rprBlacklistMerchant.Blueprint
		rprBlacklistDTTOT    rprBlacklistDTTOT.Blueprint
		rprBlacklistHistory  rprBlacklistHistory.Blueprint
		rprTransactionType   rprTransactionType.Blueprint
		rprTransaction       rprTransaction.Blueprint
		log                  *zap.SugaredLogger // log logging instance
	}
)

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */

func NewBlueprint(
	/* [CODE GENERATOR] FUNC_PARAM */
	rprBlacklistReceiver rprBlacklistReceiver.Blueprint,
	rprBlacklistSender rprBlacklistSender.Blueprint,
	rprBlacklistMerchant rprBlacklistMerchant.Blueprint,
	rprBlacklistHistory rprBlacklistHistory.Blueprint,
	rprTransactionType rprTransactionType.Blueprint,
	rprBlacklistDTTOT rprBlacklistDTTOT.Blueprint,
	rprTransaction rprTransaction.Blueprint,
	log *zap.SugaredLogger,
) Blueprint {
	const fName = "usecases.uscBlacklistManagement.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.rprBlacklistReceiver = rprBlacklistReceiver
	bp.rprBlacklistSender = rprBlacklistSender
	bp.rprBlacklistMerchant = rprBlacklistMerchant
	bp.rprBlacklistDTTOT = rprBlacklistDTTOT
	bp.rprBlacklistHistory = rprBlacklistHistory
	bp.rprTransactionType = rprTransactionType
	bp.rprTransaction = rprTransaction
	bp.log = log
	return bp
}
