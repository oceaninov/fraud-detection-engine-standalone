package uscChannelManagement

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"context"
	"gitlab.com/fds22/detection-sys/src/repositories/rprChannel"
	"go.uber.org/zap"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint business function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		ChannelCreate(ctx context.Context, request *RequestChannelCreate) (*ResponseChannelCreate, error)
		ChannelUpdate(ctx context.Context, request *RequestChannelUpdate) (*ResponseChannelUpdate, error)
		ChannelDelete(ctx context.Context, request *RequestChannelDelete) (*ResponseChannelDelete, error)
		ChannelGetApproved(ctx context.Context, request *RequestChannelGetApproved) (*ResponseChannelGetApproved, error)
		ChannelGetApprovedAll(ctx context.Context, request *RequestChannelGetApproved) (*ResponseChannelGetApproved, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		rprChannel rprChannel.Blueprint
		log        *zap.SugaredLogger // log logging instance
	}
)

func NewBlueprint(
	/* [CODE GENERATOR] FUNC_PARAM */
	rprChannel rprChannel.Blueprint,
	log *zap.SugaredLogger,
) Blueprint {
	const fName = "usecases.rprChannel.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.rprChannel = rprChannel
	bp.log = log
	return bp
}

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */
