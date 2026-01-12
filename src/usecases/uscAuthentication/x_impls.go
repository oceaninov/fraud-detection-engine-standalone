package uscAuthentication

import (
	"context"
	"gitlab.com/fds22/detection-sys/pkg/environments"
	"gitlab.com/fds22/detection-sys/pkg/otp"
	"gitlab.com/fds22/detection-sys/src/repositories/rprAuthentication"
	"go.uber.org/zap"
)

/* [CODE GENERATOR] INITIALIZE_CODE */
type (
	// Blueprint business function interface design
	Blueprint interface {
		/* [CODE GENERATOR] FUNC_BLUEPRINT_CNTR */

		UserRegister(ctx context.Context, request *RequestUserRegister) (*ResponseUserRegister, error)
		UserLogin(ctx context.Context, request *RequestLoginRegister) (*ResponseLoginRegister, error)
		UserBanned(ctx context.Context, request *RequestBannedUser) (*ResponseBannedUser, error)
		GetAvailableRoles(ctx context.Context, request *RequestGetRoles) (*ResponseGetRoles, error)
		GetRegisteredUsers(ctx context.Context, request *RequestRegisteredUsers) (*ResponseRegisteredUsers, error)
		ResetPassword(ctx context.Context, request *RequestResetPassword) (*ResponseResetPassword, error)
		ResetPasswordSubmit(ctx context.Context, request *RequestResetPasswordSubmit) (*ResponseResetPasswordSubmit, error)
		ResetPasswordValidate(ctx context.Context, request *RequestResetPasswordValidate) (*ResponseResetPasswordValidate, error)
		ChangePassword(ctx context.Context, request *RequestChangePassword) (*ResponseChangePassword, error)
	}
	// blueprint constructor parameters
	blueprint struct {
		/* [CODE GENERATOR] ATTR_CREATION */
		rprAuthentication rprAuthentication.Blueprint
		env               *environments.Envs
		log               *zap.SugaredLogger // log logging instance
		otpSender         otp.OTP
	}
)

/* [CODE GENERATOR] FUNC_BLUEPRINT_IMPL */

func NewBlueprint(
	/* [CODE GENERATOR] FUNC_PARAM */
	rprAuthentication rprAuthentication.Blueprint,
	env *environments.Envs,
	log *zap.SugaredLogger,
) Blueprint {
	const fName = "usecases.uscAuthentication.NewBlueprint"
	log.Infow(fName, "reason", "instance initialization started")
	defer log.Infow(fName, "reason", "instance initialization ended")

	// initialize
	bp := new(blueprint)
	/* [CODE GENERATOR] ATTR_ASSIGN */
	bp.rprAuthentication = rprAuthentication
	bp.env = env
	bp.log = log
	bp.otpSender = otp.NewOTP(otp.Config{
		SMTPHost:     env.SMTPHost,
		SMTPPort:     env.SMTPPort,
		SMTPAuthPass: env.SMTPAuthPass,
		SMTPAuth:     env.SMTPAuth,
		SMTPSender:   env.SMTPSender,
	})
	return bp
}
