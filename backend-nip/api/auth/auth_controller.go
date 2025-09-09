package auth

import (
	"errors"

	"github.com/UniBO-PRISMLab/nip/models"
	"github.com/UniBO-PRISMLab/nip/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type AuthController struct {
	authenticatedRoute   *gin.RouterGroup
	internalRoute        *gin.RouterGroup
	unauthenticatedRoute *gin.RouterGroup
	logger               *zerolog.Logger
	configuration        models.Configuration
	authService          *Service
}

func InjectRoutes(
	routes models.DefaultRoutes,
	configuration models.Configuration,
	authService *Service,
) {
	controller := newAuthController(
		routes,
		configuration,
		authService,
	)

	controller.injectUnAuthenticatedRoutes()
}

func newAuthController(routes models.DefaultRoutes,
	configuration models.Configuration,
	authService *Service,
) *AuthController {
	controllerLogger := utils.InitServiceAdvancedLogger("AuthController")

	return &AuthController{
		unauthenticatedRoute: routes.UnauthenticatedRoute,
		authenticatedRoute:   routes.AuthenticatedRoute,
		internalRoute:        routes.InternalRoute,
		configuration:        configuration,
		logger:               controllerLogger,
		authService:          authService,
	}
}

func (c *AuthController) injectUnAuthenticatedRoutes() {
	v1 := c.unauthenticatedRoute.Group("v1")

	{
		v1.POST(
			"auth/pac",
			c.issuePAC(),
		)

		v1.POST(
			"auth/sac",
			c.getSAC(),
		)

		v1.POST(
			"auth/verify-pac",
			c.verifyPAC(),
		)
	}
}

// issuePAC godoc
//
//	@Tags			auth
//	@Schemes		http
//	@Router			/v1/auth/pac [post]
//	@Summary		PAC issuance request
//	@Description	Allows a user to request a PAC (a one-time code) for services requiring an authenticated public identity. The user provide a payload namely `SIGN(PID, SK)` (PID signed with the Secret Key used to obtained the PID). PAC is temporarily stored locally.
//	@Accept			json
//	@Produce		json
//	@Param			models.PACRequestModel	body		models.PACRequestModel		true	"PAC Request Model"
//	@Success		200						{object}	models.PACResponseModel		"The Public Authentication Code"
//	@Failure		400						{object}	models.ErrorResponseModel	"Bad request"
//	@Failure		404						{object}	models.ErrorResponseModel	"Not found"
//	@Failure		500						{object}	models.ErrorResponseModel	"An error occurred"
func (c *AuthController) issuePAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var PAC *models.PACResponseModel
		var err error

		req := models.PACRequestModel{}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			c.logger.Error().Err(err).Msg("Error during PAC issuance")
			ctx.JSON(400, models.ErrorBadRequestResponseModel)
			return
		}

		PAC, err = c.authService.IssuePAC(ctx, &req)
		if err != nil {
			if errors.Is(err, models.ErrorUserWithPIDNotFound) {
				ctx.JSON(404, models.ErrorResponseModelWithMsg(404, err.Error()))
				return
			}
			c.logger.Error().Err(err).Msg("Error during PAC issuance")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, PAC)
	}
}

// getSAC godoc
//
//	@Tags			auth
//	@Schemes		http
//	@Router			/v1/auth/sac [post]
//	@Summary		SAC issuance request
//	@Description	Allows a user to request a SAC (one-time code) for services accepting an authenticated anonymous identity.        The user queries the NIP by sending a message containing `ENC(SID, SK)` (SID signed with the Secret Key associated with the PK used for SID storage at seed generation time). The NIP verifies ownership and issues the SAC storing the mapping SAC:SID.
//	@Accept			json
//	@Produce		json
//	@Param			models.SACRequestModel	body		models.SACRequestModel		true	"SAC Request Model"
//	@Success		200						{object}	models.SACResponseModel		"The Secret Authentication Code"
//	@Failure		400						{object}	models.ErrorResponseModel	"Bad request"
//	@Failure		500						{object}	models.ErrorResponseModel	"An error occurred"
func (c *AuthController) getSAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var SAC *models.SACResponseModel
		var err error

		SAC, err = c.authService.GetSAC(ctx)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error during SAC issuance")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, SAC)
	}
}

// verifyPAC godoc
//
//	@Tags			auth
//	@Schemes		http
//	@Router			/v1/auth/verify-pac [post]
//	@Summary		PAC verification request
//	@Description	Allows a user to verify a PAC (Public Authentication Code) for services accepting an authenticated anonymous identity.
//	@Accept			json
//	@Produce		json
//	@Param			models.PACVerificationRequestModel	body		models.PACVerificationRequestModel	true	"PAC Verification Request Model"
//	@Success		200									{object}	models.PACVerificationResponseModel	"The Public Authentication Code Verification Response"
//	@Failure		400									{object}	models.ErrorResponseModel			"Bad request"
//	@Failure		404									{object}	models.ErrorResponseModel			"Not found"
//	@Failure		500									{object}	models.ErrorResponseModel			"An error occurred"
func (c *AuthController) verifyPAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.PACVerificationRequestModel
		var resp *models.PACVerificationResponseModel
		var err error

		if err := ctx.ShouldBindJSON(&req); err != nil {
			c.logger.Error().Err(err).Msg("Error during PAC verification")
			ctx.JSON(400, models.ErrorBadRequestResponseModel)
			return
		}

		if resp, err = c.authService.VerifyPAC(ctx, &req); err != nil {
			if errors.Is(err, models.ErrorUserWithPIDNotFound) {
				ctx.JSON(404, models.ErrorResponseModelWithMsg(404, err.Error()))
				return
			}

			if errors.Is(err, models.ErrorPACNotValid) {
				ctx.JSON(404, models.ErrorResponseModelWithMsg(404, err.Error()))
				return
			}

			c.logger.Error().Err(err).Msg("Error during PAC verification")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, resp)
	}
}
