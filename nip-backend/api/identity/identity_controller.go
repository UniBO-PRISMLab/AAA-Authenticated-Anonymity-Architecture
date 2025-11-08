package identity

import (
	"errors"

	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/UniBO-PRISMLab/nip-backend/utils"
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
//	@Description	Public Identity Data (PID) is an anonymous token which can be released only by the NIP and it identifies the user without explicitly sharing information. PID is derived from user's public key and a random nonce, the server calculates `HMAC_SHA256(SHA256(PK) || nonce, SK)` where `HMAC_SHA256` is a keyed-hash message authentication code (HMAC) using SHA256 as the hash function PK is the public key and SK is a secret key stored on the server. The resulting PID is a 32 byte string encoded using base64. The API accepts 2048 bit RSA PKCS#8 keys encoded in base64.
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
