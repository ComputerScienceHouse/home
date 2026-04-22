package server

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/ComputerScienceHouse/home/api"
	cshauth "github.com/computersciencehouse/csh-auth"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var (
	// During the build process, the frontend is built from the `/web` dir
	// and then is copied to `/server/web` where this "go:embed web" directive
	// will create an embedded file system which is served by the HTTP server
	//go:embed web
	webFS embed.FS
)

func handleStatic(c *gin.Context) {
	fsys, err := fs.Sub(webFS, "web") // Effectively 'cd web'
	if err != nil {
		// Explicitly Fatal because there is something wrong with the build process
		log.Fatal().Err(err).Msg("Couldn't create web filesystem (Did the build process fail?)")
	}

	// Cut trailing and leading slashes so that
	// /my/cool/path/ -> my/cool/path which is now a valid path
	path := strings.Trim(c.Request.URL.Path, "/")

	if !strings.HasPrefix(path, "api") {
		file, err := fsys.Open(path)
		if err != nil {
			// If there is no file at this path, serve the root dir
			if errors.Is(err, fs.ErrNotExist) || path == "" {
				c.FileFromFS("/", http.FS(fsys))
			} else {
				log.Error().Err(err).Msg("unexpected error serving web")
			}
			return
		}
		// Close the file
		if err := file.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close file")
		}
		c.FileFromFS(path, http.FS(fsys))
		return
	}
	c.AbortWithStatus(http.StatusNotFound)
}

func Serve() error {

	address := os.Getenv("SERVER_URI")

	auth := cshauth.CSHAuth{}
	auth.Init(os.Getenv("OIDC_ID"),
		os.Getenv("OIDC_SECRET"),
		os.Getenv("JWT_SECRET_KEY"),
		os.Getenv("STATE"),
		address,
		address+"/auth/callback",
		address+"/auth/login",
		[]string{"profile", "email", "groups"})

	// Create gin router
	router := gin.New()

	// Setup compression
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Setup logging
	router.Use(logger.SetLogger())

	// Handle API routes
	apiServer := api.NewAPIServer()

	router.GET("/auth/login", auth.AuthRequest)
	router.GET("/auth/callback", auth.AuthCallback)
	router.GET("/auth/logout", auth.AuthLogout)

	api.RegisterHandlersWithOptions(router, apiServer, api.GinServerOptions{
		Middlewares: []api.MiddlewareFunc{
			api.MiddlewareFunc(auth.AuthMiddleware()),
		}})

	// If we haven't explicitly defined a route, default to serving the frontend
	router.NoRoute(handleStatic)

	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Info().Msgf("Serving at %s", s.Addr)
	return s.ListenAndServe()
}
