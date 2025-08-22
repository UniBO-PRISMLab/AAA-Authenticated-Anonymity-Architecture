package identity

import (
	"github.com/UniBO-PRISMLab/nip/models"
	"github.com/UniBO-PRISMLab/nip/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type IdentityController struct {
	authenticatedRoute   *gin.RouterGroup
	internalRoute        *gin.RouterGroup
	unauthenticatedRoute *gin.RouterGroup
	logger               *zerolog.Logger
	configuration        models.Configuration
	identityService      *Service
}

func InjectRoutes(
	routes models.DefaultRoutes,
	configuration models.Configuration,
	identityService *Service,
) {
	controller := newIdentityController(
		routes,
		configuration,
		identityService,
	)

	controller.injectUnAuthenticatedRoutes()
}

func newIdentityController(routes models.DefaultRoutes,
	configuration models.Configuration,
	identityService *Service,
) *IdentityController {
	controllerLogger := utils.InitServiceAdvancedLogger("IdentityController")

	return &IdentityController{
		unauthenticatedRoute: routes.UnauthenticatedRoute,
		authenticatedRoute:   routes.AuthenticatedRoute,
		internalRoute:        routes.InternalRoute,
		configuration:        configuration,
		logger:               controllerLogger,
		identityService:      identityService,
	}
}

func (c *IdentityController) injectUnAuthenticatedRoutes() {
	v1 := c.unauthenticatedRoute.Group("v1")

	{
		v1.POST(
			"identity/pid",
			c.getPID(),
		)

	}
}

// getPID godoc
//
//	@Tags			identity
//	@Schemes		http
//	@Router			/v1/identity/pid [post]
//	@Summary		PID issuance request
//	@Description	Allows a user to submit a request to their National Identity Provider (NIP) with personal data (e.g., ID card, passport details) and a Public Key (PK) to initiate the Public Identity Data (PID) issuance. The NIP verifies the user's real identity and issues a PID.
//	@Accept			json
//	@Produce		json
//	@Param			models.PIDRequestModel	body		models.PIDRequestModel		true	"PID Request Model"
//	@Success		200						{object}	models.PIDResponseModel		"The PID"
//	@Failure		400						{object}	models.ErrorResponseModel	"Bad request"
//	@Failure		500						{object}	models.ErrorResponseModel	"An error occurred"
func (c *IdentityController) getPID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var PID *models.PIDResponseModel
		var err error

		PID, err = c.identityService.GetPID(ctx)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error during PID issuance")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, PID)
	}
}
