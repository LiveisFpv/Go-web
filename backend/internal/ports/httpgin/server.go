package httpgin

import (
	"net/http"

	"backend/internal/app"
	"backend/internal/crypt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
	app  *gin.Engine
}

func NewHTTPServer(port string, a *app.App) Server {
	gin.SetMode(gin.ReleaseMode)
	s := Server{port: port, app: gin.Default()}
	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Explicitly allow frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	open := s.app.Group("/")
	OpenRouter(open, a)
	api := open.Group("api/v1")
	api.Use(crypt.AuthMiddleware())
	AppRouter(api, a)
	return s
}

func (s *Server) Listen() error {
	return s.app.Run(s.port)
}

func (s *Server) Handler() http.Handler {
	return s.app
}
