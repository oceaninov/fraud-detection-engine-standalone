package uscTrxTypeManagement

import (
	/* [CODE GENERATOR] IMPORT_PKG */
	"context"
	"gitlab.com/fds22/detection-sys/src/repositories/rprTransactionType"
	"go.uber.org/zap"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint business function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		TransactionTypeCreate(ctx context.Context, request *RequestTransactionTypeCreate) (*ResponseTransactionTypeCreate, error)
		TransactionTypeUpdate(ctx context.Context, request *RequestTransactionTypeUpdate) (*ResponseTransactionTypeUpdate, error)
		TransactionTypeDelete(ctx context.Context, request *RequestTransactionTypeDelete) (*ResponseTransactionTypeDelete, error)
		TransactionTypeGetApproved(ctx context.Context, request *RequestTransactionTypeGetApproved) (*ResponseTransactionTypeGetApproved, error)
		TransactionTypeGetApprovedAll(ctx context.Context, request *RequestTransactionTypeGetApproved) (*ResponseTransactionTypeGetApproved, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		rprTransactionType rprTransactionType.Blueprint
		log                *zap.SugaredLogger // log logging instance
	}
)

func NewBlueprint(
	/* [CODE GENERATOR] FUNC_PARAM */
	rprTransactionType rprTransactionType.Blueprint,
	log *zap.SugaredLogger,
) Blueprint {
	const fName = "usecases.rprTransactionType.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.rprTransactionType = rprTransactionType
	bp.log = log
	return bp
}

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */
