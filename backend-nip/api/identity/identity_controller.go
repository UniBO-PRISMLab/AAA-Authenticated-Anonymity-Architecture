package identity

import (
	"errors"

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
			c.issuePID,
		)
	}
}

// issuePID godoc
//
//	@Tags			identity
//	@Schemes		http
//	@Router			/v1/identity/pid [post]
//	@Summary		PID issuance request
//	@Description	The NIP issues a Public Identity Data (PID), i.e., an anonymous token that identifies the user without explicitly sharing information, saves it on its local database, and shares it with the user. The real identity of the user is carried by the PID alone, since it has been officially issued by an NIP that verified the identity of the user. Additionally, the NIP will be the only one able to connect a PID to the real information of the respective user.
//	@Accept			json
//	@Produce		json
//	@Param			models.PIDRequestModel	body		models.PIDRequestModel		true	"PID Request Model"
//	@Success		200						{object}	models.PIDResponseModel		"The PID"
//	@Failure		400						{object}	models.ErrorResponseModel	"Bad request"
//	@Failure		500						{object}	models.ErrorResponseModel	"An error occurred"
func (c *IdentityController) issuePID(ctx *gin.Context) {
	var PID *models.PIDResponseModel
	var err error

	req := models.PIDRequestModel{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error().Err(err).Msg("Error during PID issuance")
		ctx.JSON(400, models.ErrorBadRequestResponseModel)
		return
	}

	PID, err = c.identityService.IssuePID(ctx, &req)
	if err != nil {
		if errors.Is(err, models.ErrorInvalidPublicKeyHeader) ||
			errors.Is(err, models.ErrorInvalidPublicKey) ||
			errors.Is(err, models.ErrorPublicKeyDecoding) ||
			errors.Is(err, models.ErrorPKAlreadyAssociated) {
			c.logger.Error().Err(err).Msg("Error during PID issuance")
			ctx.JSON(400, models.ErrorResponseModelWithMsg(400, err.Error()))
			return
		}

		c.logger.Error().Err(err).Msg("Error during PID issuance")
		ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
		return
	}

	ctx.JSON(200, PID)
}
