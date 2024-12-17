package httpgin

import (
	"net/http"

	"backend/internal/app"

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
	s.app.Use(cors.Default())
	api := s.app.Group("/api/v1")
	AppRouter(api, a)
	return s
}

func (s *Server) Listen() error {
	return s.app.Run(s.port)
}

func (s *Server) Handler() http.Handler {
	return s.app
}
