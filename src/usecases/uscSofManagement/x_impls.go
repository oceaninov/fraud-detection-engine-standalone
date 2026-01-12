package uscSofManagement

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"context"
	"gitlab.com/fds22/detection-sys/src/repositories/rprSofMan"
	"go.uber.org/zap"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint business function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		SOFCreate(ctx context.Context, request *RequestSOFCreate) (*ResponseSOFCreate, error)
		SOFUpdate(ctx context.Context, request *RequestSOFUpdate) (*ResponseSOFUpdate, error)
		SOFDelete(ctx context.Context, request *RequestSOFDelete) (*ResponseSOFDelete, error)
		SOFGetApproved(ctx context.Context, request *RequestSOFGetApproved) (*ResponseSOFGetApproved, error)
		SOFGetApprovedAll(ctx context.Context, request *RequestSOFGetApproved) (*ResponseSOFGetApproved, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		rprSofMan rprSofMan.Blueprint
		log       *zap.SugaredLogger // log logging instance
	}
)

func NewBlueprint(
	/* [CODE GENERATOR] FUNC_PARAM */
	rprSofMan rprSofMan.Blueprint,
	log *zap.SugaredLogger,
) Blueprint {
	const fName = "usecases.uscSofManagement.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.rprSofMan = rprSofMan
	bp.log = log
	return bp
}

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */
