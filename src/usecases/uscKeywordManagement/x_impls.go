package uscKeywordManagement

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"context"
	"gitlab.com/fds22/detection-sys/src/repositories/rprKeyword"
	"go.uber.org/zap"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint business function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		KeywordCreate(ctx context.Context, request *RequestKeywordCreate) (*ResponseKeywordCreate, error)
		KeywordUpdate(ctx context.Context, request *RequestKeywordUpdate) (*ResponseKeywordUpdate, error)
		KeywordDelete(ctx context.Context, request *RequestKeywordDelete) (*ResponseKeywordDelete, error)
		KeywordApprove(ctx context.Context, request *RequestKeywordApprove) (*ResponseKeywordApprove, error)
		KeywordReject(ctx context.Context, request *RequestKeywordReject) (*ResponseKeywordReject, error)
		KeywordGetApprovalDetail(ctx context.Context, request *RequestKeywordGetApprovalDetail) (*ResponseKeywordGetApprovalDetail, error)
		KeywordGetApproval(ctx context.Context, request *RequestKeywordGetApproval) (*ResponseKeywordGetApproval, error)
		KeywordGetApproved(ctx context.Context, request *RequestKeywordGetApproved) (*ResponseKeywordGetApproved, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		rprKeyword rprKeyword.Blueprint
		log        *zap.SugaredLogger // log logging instance
	}
)

func NewBlueprint(
	/* [CODE GENERATOR] FUNC_PARAM */
	rprKeyword rprKeyword.Blueprint,
	log *zap.SugaredLogger,
) Blueprint {
	const fName = "usecases.rprKeyword.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.rprKeyword = rprKeyword
	bp.log = log
	return bp
}

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */
