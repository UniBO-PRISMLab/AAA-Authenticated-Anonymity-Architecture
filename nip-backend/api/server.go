package api

import (
	"fmt"
	"net/http"

	"github.com/UniBO-PRISMLab/nip-backend/api/auth"
	"github.com/UniBO-PRISMLab/nip-backend/api/identity"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine          *gin.Engine
	configuration   models.Configuration
	identityService *identity.Service
	authService     *auth.Service
}

func NewServer(
	configuration models.Configuration,
	identityService *identity.Service,
	authService *auth.Service,
) *Server {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowHeaders:     configuration.CORS.AllowHeaders,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowOrigins:     configuration.CORS.AllowOrigins,
		AllowCredentials: true,
	}))

	server := &Server{
		engine:          engine,
		configuration:   configuration,
		identityService: identityService,
		authService:     authService,
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	unauthenticatedRoute := s.engine.Group("/")
	// internalRoute := s.engine.Group("/")

	unauthenticatedRoute.GET("/", s.healthRoute())
	unauthenticatedRoute.GET("/health", s.healthRoute())

	routes := models.DefaultRoutes{
		UnauthenticatedRoute: unauthenticatedRoute,
		// AuthenticatedRoute:   authenticatedRoute,
		// InternalRoute:        internalRoute,
	}

	identity.InjectRoutes(routes, s.configuration, s.identityService)
	auth.InjectRoutes(routes, s.configuration, s.authService)

	if s.configuration.Environment == models.Development {
		log.Info().Msgf("Enabled swagger on http://%s:%d/swagger/index.html", s.configuration.HTTPHost, s.configuration.HTTPPort)
		s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (s *Server) healthRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	}
}

func (s *Server) Listen() error {
	address := fmt.Sprintf("%s:%d", s.configuration.HTTPHost, s.configuration.HTTPPort)

	log.Info().Msgf("Listening on %s", address)
	return s.engine.Run(address)
}
