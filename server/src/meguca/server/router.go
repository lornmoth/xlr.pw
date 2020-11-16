package server

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"meguca/assets"
	"meguca/auth"
	"meguca/config"
	"meguca/db"
	"meguca/imager"
	"meguca/util"
	"meguca/websockets"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/go-playground/log"
	"github.com/gorilla/handlers"
)

var (
	// Address is the listening address of the HTTP web server
	address string

	// Defines if HTTPS should be used for listening for incoming connections.
	// Requires sslCert and sslKey to be set.
	ssl bool

	// Path to SSL certificate
	sslCert string

	// Path to SSL key
	sslKey string

	// Defines, if all traffic should be piped through a gzip compression
	// -decompression handler
	enableGzip bool

	isTest bool

	healthCheckMsg = []byte("God's in His heaven, all's right with the world")
)

// Used for overriding during tests
var webRoot = "www"

func startWebServer() (err error) {
	r := createRouter()

	var w bytes.Buffer
	w.WriteString("listening on http")
	if ssl {
		w.WriteByte('s')
	}
	fmt.Fprintf(&w, "://%s", address)
	log.Info(w.String())

	if ssl {
		err = http.ListenAndServeTLS(address, sslCert, sslKey, r)
	} else {
		err = http.ListenAndServe(address, r)
	}
	if err != nil {
		return util.WrapError("error starting web server", err)
	}
	return
}

func handlePanic(w http.ResponseWriter, r *http.Request, err interface{}) {
	http.Error(w, fmt.Sprintf("500 %s", err), 500)
	ip, ipErr := auth.GetIP(r)
	if ipErr != nil {
		ip = "invalid IP"
	}
	log.Errorf("server: %s: %#v\n%s\n", ip, err, debug.Stack())
}

// Create the monolithic router for routing HTTP requests. Separated into own
// function for easier testability.
func createRouter() http.Handler {
	r := httptreemux.NewContextMux()
	r.NotFoundHandler = func(w http.ResponseWriter, _ *http.Request) {
		text404(w)
	}
	r.PanicHandler = handlePanic

	r.GET("/robots.txt", serveRobotsTXT)

	api := r.NewGroup("/api")
	api.GET("/health-check", healthCheck)
	assets := r.NewGroup("/assets")
	txts := r.NewGroup("/txt")
	if config.ImagerMode != config.NoImager {
		// All upload images
		api.POST("/upload", imager.NewImageUpload)
		api.POST("/upload-hash", imager.UploadImageHash)
		api.POST("/create-thread", createThread)
		api.POST("/create-reply", createReply)

		assets.GET("/images/*path", serveImages)

		// Captcha API
		captcha := api.NewGroup("/captcha")
		captcha.GET("/:board", serveNewCaptcha)
		captcha.POST("/:board", authenticateCaptcha)
		captcha.GET("/confirmation", renderCaptchaConfirmation)
	}
	if config.ImagerMode != config.ImagerOnly {
		// HTML
		r.GET("/", redirectToDefault)
		r.GET("/:board/", func(w http.ResponseWriter, r *http.Request) {
			boardHTML(w, r, extractParam(r, "board"), false)
		})
		r.GET("/:board/catalog", func(w http.ResponseWriter, r *http.Request) {
			boardHTML(w, r, extractParam(r, "board"), true)
		})
		// Needs override, because it conflicts with crossRedirect
		r.GET("/all/catalog", func(w http.ResponseWriter, r *http.Request) {
			// Artificially set board to "all"
			boardHTML(w, r, "all", true)
		})
		r.GET("/:board/:thread", threadHTML)
		r.GET("/all/:id", crossRedirect)

		html := r.NewGroup("/html")
		html.GET("/board-navigation", boardNavigation)
		html.GET("/owned-boards/:userID", ownedBoardSelection)
		html.GET("/create-board", boardCreationForm)
		html.GET("/change-password", changePasswordForm)
		html.POST("/configure-board/:board", boardConfigurationForm)
		html.POST("/configure-server", serverConfigurationForm)
		html.GET("/assign-staff/:board", staffAssignmentForm)
		html.GET("/set-banners", bannerSettingForm)
		html.GET("/set-loading", loadingAnimationForm)
		html.GET("/bans/:board", banList)
		html.GET("/mod-log/:board", modLog)
		html.GET("/report/:id", reportForm)
		html.GET("/reports/:board", reportList)

		// JSON API
		json := r.NewGroup("/json")
		boards := json.NewGroup("/boards")
		boards.GET("/:board/", func(w http.ResponseWriter, r *http.Request) {
			boardJSON(w, r, false)
		})
		boards.GET("/:board/catalog", func(w http.ResponseWriter,
			r *http.Request,
		) {
			boardJSON(w, r, true)
		})
		boards.GET("/:board/:thread", threadJSON)
		json.GET("/post/:post", servePost)
		json.GET("/config", serveConfigs)
		json.GET("/extensions", serveExtensionMap)
		json.GET("/board-config/:board", serveBoardConfigs)
		json.GET("/board-list", serveBoardList)
		json.GET("/ip-count", serveIPCount)
		json.POST("/thread-updates", serveThreadUpdates)

		// Internal API
		api.GET("/socket", func(w http.ResponseWriter, r *http.Request) {
			err := websockets.Handler(w, r)
			if err != nil {
				httpError(w, r, err)
			}
		})
		api.GET("/youtube-data/:id", youTubeData)
		api.GET("/bitchute-title/:id", bitChuteTitle)
		api.POST("/register", register)
		api.POST("/login", login)
		api.POST("/logout", logout)
		api.POST("/logout-all", logoutAll)
		api.POST("/change-password", changePassword)
		api.POST("/board-config/:board", servePrivateBoardConfigs)
		api.POST("/configure-board/:board", configureBoard)
		api.POST("/config", servePrivateServerConfigs)
		api.POST("/configure-server", configureServer)
		api.POST("/create-board", createBoard)
		api.POST("/delete-board", deleteBoard)
		api.POST("/delete-post", deletePost)
		api.POST("/delete-image", deleteImage)
		api.POST("/spoiler-image", modSpoilerImage)
		api.POST("/ban", ban)
		api.POST("/notification", sendNotification)
		api.POST("/assign-staff", assignStaff)
		api.POST("/same-IP/:id", getSameIPPosts)
		api.POST("/sticky", setThreadSticky)
		api.POST("/lock-thread", setThreadLock)
		api.POST("/unban/:board", unban)
		api.POST("/set-banners", setBanners)
		api.POST("/set-loading", setLoadingAnimation)
		api.POST("/report", report)
		api.POST("/purge-post", purgePost)

		redir := api.NewGroup("/redirect")
		redir.POST("/by-ip", redirectByIP)
		redir.POST("/by-thread", redirectByThread)

		// Assets
		assets.GET("/banners/:board/:id", serveBanner)
		assets.GET("/loading/:board", serveLoadingAnimation)
		assets.GET("/*path", serveAssets)

		// text assets
		// serves /assets/txt
		txts.GET("/*path", serveTextFile)

	}

	h := http.Handler(r)
	if enableGzip {
		h = handlers.CompressHandlerLevel(h, gzip.DefaultCompression)
	}

	return h
}

// Redirects to / requests to /all/ board
func redirectToDefault(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/all/", 301)
}

// Generate a robots.txt with only select boards preventing indexing
func serveRobotsTXT(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	// Would be pointles without the /all/ board disallowed.
	// Also, this board can be huge. Don't want bots needlessly crawling it.
	buf.WriteString("User-agent: *\nDisallow: /all/\n")
	for _, c := range config.GetAllBoardConfigs() {
		if c.DisableRobots {
			fmt.Fprintf(&buf, "Disallow: /%s/\n", c.ID)
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	buf.WriteTo(w)
}

// Redirect the client to the appropriate board through a cross-board redirect
func crossRedirect(w http.ResponseWriter, r *http.Request) {
	idStr := extractParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		text404(w)
		return
	}

	board, op, err := db.GetPostParenthood(id)
	if err != nil {
		httpError(w, r, err)
		return
	}
	url := r.URL
	url.Path = fmt.Sprintf("/%s/%d", board, op)
	if url.Query().Get("last") != "" {
		url.Fragment = "bottom"
	} else {
		url.Fragment = "p" + idStr
	}
	http.Redirect(w, r, url.String(), 301)
}

// Health check to ensure server is still online
func healthCheck(w http.ResponseWriter, r *http.Request) {
	if config.ImagerMode != config.NoImager {
		// Ensure thumbnailing queue is not blocked
		err := func() (err error) {
			const name = "invalid_upload.psd"
			buf, err := assets.Asset(name)
			if err != nil {
				return
			}

			var body bytes.Buffer
			wr := multipart.NewWriter(&body)
			f, err := wr.CreateFormFile("image", name)
			if err != nil {
				return
			}
			_, err = f.Write(buf)
			if err != nil {
				return
			}

			req := httptest.NewRequest("POST", "/api/upload", &body)
			req.Header.Set("Content-Length", strconv.Itoa(len(buf)))
			req.Header.Set("Content-Type", wr.FormDataContentType())
			req.Header.Set("Authorization", "Bearer "+config.Get().Salt)

			ch := make(chan error)
			go func() {
				rec := httptest.NewRecorder()
				imager.NewImageUpload(rec, req)
				if rec.Code == 400 {
					ch <- nil
				} else {
					ch <- fmt.Errorf(
						"invalid healthcheck upload: code=%d body=`%s`",
						rec.Code, rec.Body.String())
				}
			}()

			timer := time.NewTimer(time.Second * 10)
			defer timer.Stop()
			select {
			case <-timer.C:
				return fmt.Errorf("healthcheck upload timeout")
			case err = <-ch:
				return
			}
		}()
		if err != nil {
			httpError(w, r, err)
			return
		}
	}
	w.Write(healthCheckMsg)
}
