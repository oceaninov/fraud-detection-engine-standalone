package uscAuthentication

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/hashicorp/go-uuid"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"gitlab.com/fds22/detection-sys/pkg/hashing"
	"gitlab.com/fds22/detection-sys/pkg/humanTime"
	"gitlab.com/fds22/detection-sys/src/repositories/rprAuthentication"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type (
	ResponseUserRegister struct {
		ResponseCode    bool   `json:"success"`
		ResponseMessage string `json:"messages"`
	}
	RequestUserRegister struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     int    `json:"role_id"`
		FullName string `json:"full_name"`
		Gender   string `json:"gender"`
	}
)

type (
	ResponseLoginRegister struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		AccessToken     string `json:"accessToken"`
		RefreshToken    string `json:"refreshToken"`
		User            struct {
			Id        string `json:"id"`
			Email     string `json:"email"`
			RoleId    string `json:"roleId"`
			AvatarUrl string `json:"avatarUrl"`
			FullName  string `json:"fullName"`
			Gender    string `json:"gender"`
			CreatedAt string `json:"createdAt"`
			UpdatedAt string `json:"updatedAt"`
		} `json:"user"`
	}

	RequestLoginRegister struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type (
	ResponseBannedUser struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
	}
	RequestBannedUser struct {
		Id string `param:"id"`
	}
)

type (
	ResponseRegisteredUsersData struct {
		Id        string `json:"id"`
		RoleId    int    `json:"roleId"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		FullName  string `json:"fullName"`
		AvatarUrl string `json:"avatarUrl"`
		Gender    string `json:"gender"`
		Activated string `json:"activated"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
	ResponseRegisteredUsers struct {
		ResponseCode    string                        `json:"responseCode"`
		ResponseMessage string                        `json:"responseMessage"`
		Users           []ResponseRegisteredUsersData `json:"users"`
		Meta            basicObject.Meta              `json:"meta"`
	}
	RequestRegisteredUsers struct {
		Page     int    `query:"page"`
		Limit    int    `query:"limit"`
		RoleId   int    `query:"roleId"`
		Gender   string `query:"gender"`
		Email    string `query:"search"`
		FullName string `query:"fullName"`
	}
)

type (
	ResponseGetRolesData struct {
		Id          string `json:"id"`
		Title       string `json:"roleName"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		Description string `json:"description"`
	}
	RequestGetRoles  struct{}
	ResponseGetRoles struct {
		ResponseCode    string                 `json:"responseCode"`
		ResponseMessage string                 `json:"responseMessage"`
		Roles           []ResponseGetRolesData `json:"roles"`
	}
)

type (
	RequestResetPassword struct {
		Activation bool   `json:"activation"`
		RequestId  string `json:"request_id"`
		Id         string `json:"id"`
		Email      string `json:"email"`
	}
	ResponseResetPassword struct {
		Success  bool   `json:"success"`
		Messages string `json:"messages"`
	}
)

type (
	RequestResetPasswordSubmit struct {
		Token              string `json:"token"`
		CurrentPassword    string `json:"current_password"`
		NewPassword        string `json:"newPassword"`
		NewPasswordConfirm string `json:"newPasswordConfirm"`
	}
	ResponseResetPasswordSubmit struct {
		Success  bool   `json:"success"`
		Messages string `json:"messages"`
	}
)

type (
	RequestResetPasswordValidate struct {
		Token string `param:"token"`
	}
	ResponseResetPasswordValidate struct {
		Success  bool   `json:"success"`
		Messages string `json:"messages"`
	}
)

type (
	RequestChangePassword struct {
		Id                 string `json:"id"`
		CurrentPassword    string `json:"currentPassword"`
		NewPassword        string `json:"newPassword"`
		NewPasswordConfirm string `json:"newPasswordConfirm"`
	}
	ResponseChangePassword struct {
		Success  bool   `json:"success"`
		Messages string `json:"messages"`
	}
)

func (b *blueprint) ResetPassword(ctx context.Context, request *RequestResetPassword) (*ResponseResetPassword, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT001"
	const fName = "usecases.uscAuthentication.ResetPassword"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ResetPassword */
	tokenByte, err := uuid.GenerateRandomBytes(16)
	if err != nil {
		return nil, err
	}
	token := fmt.Sprintf("%x", tokenByte)

	user, err := b.rprAuthentication.ReadRowUser(ctx, map[string]interface{}{
		"email": request.Email,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if user == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// write reset password data
	_, err = b.rprAuthentication.WriteRowResetPassword(ctx, user.Id, token)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	resetLink := fmt.Sprintf("%s%s", b.env.ApplicationBO, "/auth/reset-password?email=%s&token=%s")
	link := fmt.Sprintf(resetLink, user.Email, token)
	logging.Infow(fName, "link_reset_email", link)

	err = b.otpSender.Email(user.Email, link)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	response := new(ResponseResetPassword)
	response.Success = true
	response.Messages = "reset password request has been sent"
	return response, nil
}

func (b *blueprint) ResetPasswordSubmit(ctx context.Context, request *RequestResetPasswordSubmit) (*ResponseResetPasswordSubmit, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT001"
	const fName = "usecases.uscAuthentication.ResetPasswordSubmit"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ResetPasswordSubmit */
	if request.Token == "" || request.NewPassword == "" || request.NewPasswordConfirm == "" {
		errMsg := fmt.Errorf("missing required parameters")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	resetPassword, err := b.rprAuthentication.ReadResetPasswordByTokenLatest(ctx, request.Token)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	expired := humanTime.StringTimestampToTime(resetPassword.ExpiredAt.Format(time.RFC3339))
	difSec := humanTime.DiffSecFromNow(expired)
	if difSec < 1 {
		errMsg := fmt.Errorf("token Invalid or Already Expired")
		logging.Errorw(fName, "reason", "token Invalid or Already Expired")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	if request.NewPassword != request.NewPasswordConfirm {
		errMsg := fmt.Errorf("new password must be same as confirm new password")
		logging.Errorw(fName, "reason", "new password must be same as confirm new password.")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	err = b.rprAuthentication.UpdateResetPassword(ctx, int64(resetPassword.Id))
	if err != nil {
		errMsg := fmt.Errorf("failed submit reset password")
		logging.Errorw(fName, "reason", "failed submit reset password")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	newHashed, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		errMsg := fmt.Errorf("failed submit reset password")
		logging.Errorw(fName, "reason", "failed submit reset password")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	err = b.rprAuthentication.UpdateUserPasswordByID(ctx, resetPassword.UserId, string(newHashed))
	if err != nil {
		errMsg := fmt.Errorf("failed submit reset password")
		logging.Errorw(fName, "reason", "failed submit reset password")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	response := new(ResponseResetPasswordSubmit)
	response.Success = true
	response.Messages = "Success Submit Reset Password"
	return response, nil
}

func (b *blueprint) ResetPasswordValidate(ctx context.Context, request *RequestResetPasswordValidate) (*ResponseResetPasswordValidate, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT001"
	const fName = "usecases.uscAuthentication.ResetPasswordValidate"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ResetPasswordValidate */
	resetPassword, err := b.rprAuthentication.ReadResetPasswordByTokenLatest(ctx, request.Token)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	expired := humanTime.StringTimestampToTime(resetPassword.ExpiredAt.Format(time.RFC3339))
	difSec := humanTime.DiffSecFromNow(expired)
	fmt.Println(difSec)
	if difSec < 1 {
		errMsg := fmt.Errorf("token Invalid or Already Expired")
		logging.Errorw(fName, "reason", "token Invalid or Already Expired")
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	response := new(ResponseResetPasswordValidate)
	response.Success = true
	response.Messages = "Token Valid"
	return response, nil
}

func (b *blueprint) ChangePassword(ctx context.Context, request *RequestChangePassword) (*ResponseChangePassword, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT001"
	const fName = "usecases.uscAuthentication.ChangePassword"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	email := ctx.Value(defaultHeaders.Email).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_ChangePassword */
	if request.CurrentPassword == "" || request.NewPassword == "" || request.NewPasswordConfirm == "" {
		errMsg := fmt.Errorf("missing required parameters")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	user, err := b.rprAuthentication.ReadRowUser(ctx, map[string]interface{}{
		"email": email,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if user == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.CurrentPassword))
	if err != nil {
		errMsg := fmt.Errorf("incorrect password")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	err = b.rprAuthentication.UpdateUserPasswordByID(ctx, user.Id, string(hashedPassword))
	if err != nil {
		errMsg := fmt.Errorf("unable to hash password")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	response := new(ResponseChangePassword)
	response.Success = true
	response.Messages = "Change Password Success"
	return response, nil
}

func (b *blueprint) GetRegisteredUsers(ctx context.Context, request *RequestRegisteredUsers) (*ResponseRegisteredUsers, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT001"
	const fName = "usecases.uscAuthentication.GetRegisteredUsers"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_GetRegisteredUsers */
	dataResult := []ResponseRegisteredUsersData{}
	roles, meta, err := b.rprAuthentication.ReadRowUsers(ctx, request.Page, request.Limit, request.FullName, request.Email, request.Gender, request.RoleId)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if roles != nil {
		for _, d := range *roles {
			result := ResponseRegisteredUsersData{
				Id:        d.Id,
				RoleId:    d.RoleId,
				Email:     d.Email,
				Password:  d.Password,
				FullName:  d.FullName,
				AvatarUrl: d.AvatarUrl,
				Gender:    d.Gender,
				Activated: d.Activated,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseRegisteredUsers)
	response.ResponseCode = basicObject.SuccessfullyCode
	response.ResponseMessage = basicObject.Successfully
	response.Users = dataResult
	response.Meta = *meta
	return response, nil
}

func (b *blueprint) UserBanned(ctx context.Context, request *RequestBannedUser) (*ResponseBannedUser, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT001"
	const fName = "usecases.uscAuthentication.UserBanned"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UserBanned */
	err := b.rprAuthentication.DeleteRowUser(ctx, request.Id)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseBannedUser)
	response.ResponseCode = basicObject.SuccessfullyCode
	response.ResponseMessage = basicObject.Successfully
	return response, nil
}

func (b *blueprint) GetAvailableRoles(ctx context.Context, request *RequestGetRoles) (*ResponseGetRoles, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT001"
	const fName = "usecases.uscAuthentication.GetAvailableRoles"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UserBanned */
	dataResult := []ResponseGetRolesData{}
	roles, err := b.rprAuthentication.ReadRowRoles(ctx)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if roles != nil {
		for _, d := range *roles {
			result := ResponseGetRolesData{
				Id:          d.Id,
				Title:       d.Title,
				CreatedAt:   d.CreatedAt,
				UpdatedAt:   d.UpdatedAt,
				Description: d.Description,
			}
			dataResult = append(dataResult, result)
		}
	}

	// response
	response := new(ResponseGetRoles)
	response.ResponseCode = basicObject.SuccessfullyCode
	response.ResponseMessage = basicObject.Successfully
	response.Roles = dataResult
	return response, nil
}

func (b *blueprint) UserRegister(ctx context.Context, request *RequestUserRegister) (*ResponseUserRegister, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT001"
	const fName = "usecases.uscAuthentication.UserRegister"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UserRegister */
	user, err := b.rprAuthentication.ReadRowUser(ctx, map[string]interface{}{
		"email": request.Email,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if user != nil {
		errMsg := fmt.Errorf("already registered")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	newEntityUser := new(rprAuthentication.EntityUser)
	newEntityUser.Id = guuid.NewString()
	newEntityUser.RoleId = request.Role
	newEntityUser.Email = request.Email
	newEntityUser.FullName = request.FullName
	newEntityUser.AvatarUrl = "https://icons.veryicon.com/png/o/miscellaneous/standard/avatar-15.png"
	newEntityUser.Gender = request.Gender
	newEntityUser.Activated = "0"
	newEntityUser.CreatedAt = time.Now().Format(basicObject.DateAndTime)
	//newEntityUser.CreatedBy = basicObject.BlankString
	newEntityUser.UpdatedAt = time.Now().Format(basicObject.DateAndTime)
	//newEntityUser.UpdatedBy = basicObject.BlankString
	//newEntityUser.ApprovedAt = time.Now().Format(basicObject.DateAndTime)
	//newEntityUser.ApprovedBy = basicObject.BlankString

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	newEntityUser.Password = string(hashedPassword)

	// Insert user
	insertedEntity, insertedId, err := b.rprAuthentication.WriteUser(ctx, newEntityUser)
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if insertedId == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if insertedEntity == nil {
		errMsg := fmt.Errorf("failed to write data")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// response
	response := new(ResponseUserRegister)
	response.ResponseCode = basicObject.SuccessfullyTrue
	response.ResponseMessage = "register request has been sent"
	return response, nil
}

func (b *blueprint) UserLogin(ctx context.Context, request *RequestLoginRegister) (*ResponseLoginRegister, error) {
	// constructing standard module name with request id as its prefix
	const fCode = "AUT002"
	const fName = "usecases.uscAuthentication.UserLogin"
	errWrap := errorWrapper.NewWrapper(fCode)
	requestId := ctx.Value(defaultHeaders.XRequestId).(string)
	logging := b.log.With("request_id", requestId, "request_body", request, "function_code", errWrap.FCode())
	logging.Infow(fName, "reason", "execution started")
	defer logging.Infow(fName, "reason", "execution ended")

	// business logic process will be defined below this comment
	/* [CODE GENERATOR] FUNC_BLUEPRINT_LGCS_UserLogin */
	user, err := b.rprAuthentication.ReadRowUser(ctx, map[string]interface{}{
		"email": request.Email,
	})
	if err != nil {
		errMsg := fmt.Errorf("internal server error")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}
	if user == nil {
		errMsg := fmt.Errorf("data not found")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// password compare
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		errMsg := fmt.Errorf("incorrect password")
		logging.Errorw(fName, "reason", errMsg.Error())
		return nil, errWrap.WrapRepositoryError(errMsg)
	}

	// parsing access token expired duration
	accessTokenExpiredDuration, err := time.ParseDuration(b.env.JWTAccessTokenDuration)
	if err != nil {
		errMsg := fmt.Errorf("mismatch configuration")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapBusinessError(errMsg)
	}

	// parsing refresh token expired duration
	refreshTokenExpiredDuration, err := time.ParseDuration(b.env.JWTRefreshTokenDuration)
	if err != nil {
		errMsg := fmt.Errorf("mismatch configuration")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapBusinessError(errMsg)
	}

	// generate jwt access token
	refreshToken, err := hashing.GenerateJWT(
		strconv.Itoa(user.RoleId),
		user.Email,
		user.Id,
		b.env.JWTSecret,
		refreshTokenExpiredDuration,
	)
	if err != nil {
		errMsg := fmt.Errorf("mismatch configuration")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapBusinessError(errMsg)
	}

	// generate jwt access token
	accessToken, err := hashing.GenerateJWT(
		strconv.Itoa(user.RoleId),
		user.Email,
		user.Id,
		b.env.JWTSecret,
		accessTokenExpiredDuration,
	)
	if err != nil {
		errMsg := fmt.Errorf("mismatch configuration")
		logging.Errorw(fName, "reason", err.Error())
		return nil, errWrap.WrapBusinessError(errMsg)
	}

	// response
	response := new(ResponseLoginRegister)
	response.ResponseCode = basicObject.SuccessfullyCode
	response.ResponseMessage = basicObject.Successfully
	response.RefreshToken = refreshToken
	response.AccessToken = accessToken
	response.User.Id = user.Id
	response.User.Email = user.Email
	response.User.RoleId = strconv.Itoa(user.RoleId)
	//response.User.RoleId = user.RoleId
	response.User.AvatarUrl = user.AvatarUrl
	response.User.FullName = user.FullName
	response.User.Gender = user.Gender
	response.User.CreatedAt = user.CreatedAt
	response.User.UpdatedAt = user.UpdatedAt

	return response, nil
}
