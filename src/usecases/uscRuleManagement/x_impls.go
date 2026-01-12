package uscRuleManagement

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"context"
	"gitlab.com/fds22/detection-sys/src/repositories/rprRuleDetection"
	"gitlab.com/fds22/detection-sys/src/repositories/rprSof"
	"go.uber.org/zap"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint business function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		RuleDetectionCreate(ctx context.Context, request *RequestRuleDetectionCreate) (*ResponseRuleDetectionCreate, error)
		RuleDetectionUpdate(ctx context.Context, request *RequestRuleDetectionUpdate) (*ResponseRuleDetectionUpdate, error)
		RuleDetectionDelete(ctx context.Context, request *RequestRuleDetectionDelete) (*ResponseRuleDetectionDelete, error)
		RuleDetectionApprove(ctx context.Context, request *RequestRuleDetectionApprove) (*ResponseRuleDetectionApprove, error)
		RuleDetectionReject(ctx context.Context, request *RequestRuleDetectionReject) (*ResponseRuleDetectionReject, error)
		RuleDetectionGetApproval(ctx context.Context, request *RequestRuleDetectionGetApproval) (*ResponseRuleDetectionGetApproval, error)
		RuleDetectionGetApproved(ctx context.Context, request *RequestRuleDetectionGetApproved) (*ResponseRuleDetectionGetApproved, error)
		RuleDetectionGetApprovalById(ctx context.Context, request *RequestRuleApprovalById) (*ResponseRuleApprovalById, error)
		RuleDetectionGetApprovedById(ctx context.Context, request *RequestRuleApprovedById) (*ResponseRuleApprovedById, error)
		RuleDetectionEnable(ctx context.Context, request *RequestRuleEnabled) (*ResponseRuleEnabled, error)
		RuleDetectionDisable(ctx context.Context, request *RequestRuleDisabled) (*ResponseRuleDisabled, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		rprRuleDetection rprRuleDetection.Blueprint
		rprSof           rprSof.Blueprint
		log              *zap.SugaredLogger // log logging instance
	}
)

func NewBlueprint(
	/* [CODE GENERATOR] FUNC_PARAM */
	rprRuleDetection rprRuleDetection.Blueprint,
	rprSof rprSof.Blueprint,
	log *zap.SugaredLogger,
) Blueprint {
	const fName = "usecases.uscRuleManagement.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.rprRuleDetection = rprRuleDetection
	bp.rprSof = rprSof
	bp.log = log
	return bp
}

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */
