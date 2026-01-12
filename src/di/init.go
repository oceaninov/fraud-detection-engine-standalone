package di

/* [CODE GENERATOR] IMPORT_PKG_RPR */
import (
	"gitlab.com/fds22/detection-sys/src/repositories/rprAuthentication"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistDTTOT"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistHistory"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistMerchant"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistReceiver"
	"gitlab.com/fds22/detection-sys/src/repositories/rprBlacklistSender"
	"gitlab.com/fds22/detection-sys/src/repositories/rprChannel"
	"gitlab.com/fds22/detection-sys/src/repositories/rprFraudDetection"
	"gitlab.com/fds22/detection-sys/src/repositories/rprKeyword"
	"gitlab.com/fds22/detection-sys/src/repositories/rprRuleDetection"
	"gitlab.com/fds22/detection-sys/src/repositories/rprSof"
	"gitlab.com/fds22/detection-sys/src/repositories/rprSofMan"
	"gitlab.com/fds22/detection-sys/src/repositories/rprTransaction"
	"gitlab.com/fds22/detection-sys/src/repositories/rprTransactionType"
	"gitlab.com/fds22/detection-sys/src/usecases/uscAuthentication"
	"gitlab.com/fds22/detection-sys/src/usecases/uscBlacklistManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscChannelManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscKeywordManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscRuleManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscSofManagement"
	"gitlab.com/fds22/detection-sys/src/usecases/uscTrxTypeManagement"
)

/* [CODE GENERATOR] IMPORT_PKG_USC */
import "gitlab.com/fds22/detection-sys/src/usecases/uscFraud"

import (
	"gitlab.com/fds22/detection-sys/pkg/customValidator"
	"gitlab.com/fds22/detection-sys/pkg/environments"
	"gitlab.com/fds22/detection-sys/pkg/logger"
	"gitlab.com/fds22/detection-sys/src/endpoints"
	"gitlab.com/fds22/detection-sys/src/instances"
	"go.uber.org/dig"
)

// Container master containers for all dependencies injected
// this global variable will be accessed from main function
// and will provide needed instances across functionalities
var Container = dig.New()

// init default initialization function from golang
func init() {
	var err error
	// Injecting needed dependencies across functionalities
	err = Container.Provide(environments.NewEnvs)
	err = Container.Provide(logger.NewZapLogger)
	err = Container.Provide(customValidator.NewCustomValidator)
	err = Container.Provide(instances.NewOrm)

	// Outsiders
	/* [CODE GENERATOR] PROVIDE_OUTSIDERS */

	// Repositories
	/* [CODE GENERATOR] PROVIDE_REPOSITORIES */
	err = Container.Provide(rprBlacklistDTTOT.NewBlueprint)
	err = Container.Provide(rprBlacklistMerchant.NewBlueprint)
	err = Container.Provide(rprBlacklistReceiver.NewBlueprint)
	err = Container.Provide(rprBlacklistSender.NewBlueprint)
	err = Container.Provide(rprFraudDetection.NewBlueprint)
	err = Container.Provide(rprRuleDetection.NewBlueprint)
	err = Container.Provide(rprKeyword.NewBlueprint)
	err = Container.Provide(rprTransactionType.NewBlueprint)
	err = Container.Provide(rprAuthentication.NewBlueprint)
	err = Container.Provide(rprSof.NewBlueprint)
	err = Container.Provide(rprBlacklistHistory.NewBlueprint)
	err = Container.Provide(rprTransaction.NewBlueprint)
	err = Container.Provide(rprSofMan.NewBlueprint)
	err = Container.Provide(rprChannel.NewBlueprint)

	// Usecases
	/* [CODE GENERATOR] PROVIDE_USECASES */
	err = Container.Provide(uscFraud.NewBlueprint)
	err = Container.Provide(uscBlacklistManagement.NewBlueprint)
	err = Container.Provide(uscRuleManagement.NewBlueprint)
	err = Container.Provide(uscKeywordManagement.NewBlueprint)
	err = Container.Provide(uscTrxTypeManagement.NewBlueprint)
	err = Container.Provide(uscAuthentication.NewBlueprint)
	err = Container.Provide(uscSofManagement.NewBlueprint)
	err = Container.Provide(uscChannelManagement.NewBlueprint)

	// Endpoint
	err = Container.Provide(endpoints.NewEndpoint)

	if err != nil {
		panic(err)
	}
}
