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
			c.issueSAC(),
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
//	@Description	Allows a user to request a Public Authentication Code (PAC). It accepts a message `PID` and `SIGN(PID, SK)`, namely the PID signed with the private key of the public-private key pair used to obtain the PID.  When logging onto the public service, the user can show the PAC, and the service has only to query the NIP to verify that the code is associated with an authenticated user.
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

		var req models.PACRequestModel

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
			if errors.Is(err, models.ErrorPIDSignatureVerification) {
				ctx.JSON(400, models.ErrorResponseModelWithMsg(400, err.Error()))
				return
			}
			c.logger.Error().Err(err).Msg("Error during PAC issuance")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, PAC)
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

// issueSAC godoc
//
//	@Tags			auth
//	@Schemes		http
//	@Router			/v1/auth/sac [post]
//	@Summary		SAC issuance request
//	@Description	The SAC (Secret Authentication Code) is a one-time code used to authenticate the user as an anonymous user. It accepts `ENC(SID, SK)`, i.e., the SID signed with the private key associated with the public key saved on the blockchain at the moment of seed phrase creation and used in the record where the SID was stored. The NIP retrieves the SID record from the blockchain and checks that it was actually signed by that user via the PK saved in the record. This certifies that the user is the true owner of that SID.
//	@Accept			json
//	@Produce		json
//	@Param			models.SACRequestModel	body		models.SACRequestModel		true	"SAC Request Model"
//	@Success		200						{object}	models.SACResponseModel		"The Secret Authentication Code"
//	@Failure		400						{object}	models.ErrorResponseModel	"Bad request"
//	@Failure		500						{object}	models.ErrorResponseModel	"An error occurred"
func (c *AuthController) issueSAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var resp *models.SACResponseModel
		var err error

		resp, err = c.authService.IssueSAC(ctx)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error during SAC issuance")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, resp)
	}
}
