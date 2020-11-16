package server

import (
	"bytes"
	"fmt"
	"meguca/assets"
	"meguca/auth"
	"meguca/common"
	"meguca/db"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bakape/thumbnailer"
)

var (
	// Set of headers for serving images (and other uploaded files)
	imageHeaders = map[string]string{
		// max-age set to 350 days. Some caches and browsers ignore max-age, if
		// it is a year or greater, so keep it a little below.
		"Cache-Control":               "max-age=30240000, public, immutable",
		"Access-Control-Allow-Origin": "*",
	}

	// For overriding during tests
	imageWebRoot = "images"
)

// More performant handler for serving image assets. These are immutable
// (except deletion), so we can also set separate caching policies for them.
func serveImages(w http.ResponseWriter, r *http.Request) {
	path := extractParam(r, "path")
	file, err := os.Open(cleanJoin(imageWebRoot, path))
	if err != nil {
		text404(w)
		return
	}
	defer file.Close()

	head := w.Header()
	for key, val := range imageHeaders {
		head.Set(key, val)
	}

	http.ServeContent(w, r, path, time.Time{}, file)
}

func cleanJoin(a, b string) string {
	return filepath.Clean(filepath.Join(a, b))
}

// Server static assets
func serveAssets(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.RequestURI, "worker.js") {
		w.Header().Set("Service-Worker-Allowed", "/")
	}
	serveFile(w, r, cleanJoin(webRoot, extractParam(r, "path")))
}

func serveTextFile(w http.ResponseWriter, r *http.Request) {
	base := cleanJoin(webRoot, "txt")
	serveFile(w, r, cleanJoin(base, extractParam(r, "path")))
}

func serveFile(w http.ResponseWriter, r *http.Request, path string) {
	file, err := os.Open(path)
	if err != nil {
		text404(w)
		return
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		httpError(w, r, err)
		return
	}
	if stats.IsDir() {
		text404(w)
		return
	}
	modTime := stats.ModTime()
	etag := strconv.FormatInt(modTime.Unix(), 10)

	head := w.Header()
	head.Set("Cache-Control", "no-cache")
	head.Set("ETag", etag)
	http.ServeContent(w, r, path, modTime, file)
}

// Set the banners of a board
func setBanners(w http.ResponseWriter, r *http.Request) {
	board, ok := parseAssetForm(w, r, common.MaxNumBanners)
	if !ok {
		return
	}

	var (
		opts = thumbnailer.Options{
			MaxSourceDims: thumbnailer.Dims{
				Width:  300,
				Height: 100,
			},
			ThumbDims: thumbnailer.Dims{
				Width:  300,
				Height: 100,
			},
			AcceptedMimeTypes: map[string]bool{
				"image/jpeg": true,
				"image/png":  true,
				"image/gif":  true,
				"video/webm": true,
			},
		}
		banners = make([]assets.File, 0, common.MaxNumBanners)
		files   = r.MultipartForm.File["banners"]
	)
	defer func() {
		for _, f := range banners {
			if f.Data != nil {
				thumbnailer.ReturnBuffer(f.Data)
			}
		}
	}()

	for i := 0; i < common.MaxNumBanners && i < len(files); i++ {
		h := files[i]
		file, err := h.Open()
		if err != nil {
			sendFileError(w, h, err.Error())
			return
		}
		out, ok := readAssetFile(w, r, file, h, opts)
		if !ok {
			return
		}
		banners = append(banners, out)
	}

	if err := db.SetBanners(board, banners); err != nil {
		httpError(w, r, err)
	}
}

// Parse form for uploading file assets for a board.
// maxSize specifies maximum number of common.MaxAssetSize to accept.
// If ok == false, caller should return.
func parseAssetForm(w http.ResponseWriter, r *http.Request, maxSize uint) (
	board string, ok bool,
) {
	r.Body = http.MaxBytesReader(
		w,
		r.Body,
		int64(maxSize)*common.MaxAssetSize,
	)
	err := r.ParseMultipartForm(0)
	if err != nil {
		httpError(w, r, common.StatusError{err, 400})
		return
	}

	board = r.Form.Get("board")
	_, ok = canPerform(w, r, board, auth.BoardOwner, true)
	return
}

// Read a file from an asset submition form.
// If ok == false, caller should return.
// Call thumbnailer.ReturnBuffer() on out.Data to return the buffer to the
// memory pool.
func readAssetFile(
	w http.ResponseWriter,
	r *http.Request,
	f multipart.File,
	h *multipart.FileHeader,
	opts thumbnailer.Options,
) (
	out assets.File,
	ok bool,
) {
	defer f.Close()

	_buf := bytes.NewBuffer(thumbnailer.GetBuffer())
	_, err := _buf.ReadFrom(f)
	if err != nil {
		httpError(w, r, err)
		return
	}
	buf := _buf.Bytes()
	if len(buf) == 0 { // No file
		ok = true
		thumbnailer.ReturnBuffer(_buf.Bytes())
		return
	}
	if len(buf) > common.MaxAssetSize {
		sendFileError(w, h, "too large")
		return
	}

	src, thumb, err := thumbnailer.ProcessBuffer(buf, opts)
	defer func() {
		if thumb.Data != nil {
			thumbnailer.ReturnBuffer(thumb.Data)
		}
	}()
	switch {
	case err != nil:
		sendFileError(w, h, err.Error())
	case src.HasAudio:
		sendFileError(w, h, "has audio")
	default:
		ok = true
		out = assets.File{
			Data: buf,
			Mime: src.Mime,
		}
	}
	return
}

func setLoadingAnimation(w http.ResponseWriter, r *http.Request) {
	board, ok := parseAssetForm(w, r, 1)
	if !ok {
		return
	}

	var out assets.File
	file, h, err := r.FormFile("image")
	switch err {
	case nil:
		out, ok = readAssetFile(w, r, file, h, thumbnailer.Options{
			MaxSourceDims: thumbnailer.Dims{
				Width:  300,
				Height: 300,
			},
			ThumbDims: thumbnailer.Dims{
				Width:  300,
				Height: 300,
			},
			AcceptedMimeTypes: map[string]bool{
				"image/gif":  true,
				"video/webm": true,
			},
		})
		defer func() {
			if out.Data != nil {
				thumbnailer.ReturnBuffer(out.Data)
			}
		}()
		if !ok {
			return
		}
	case http.ErrMissingFile:
		err = nil
	default:
		httpError(w, r, common.StatusError{err, 400})
		return
	}

	if err := db.SetLoadingAnimation(board, out); err != nil {
		httpError(w, r, err)
	}
}

func sendFileError(w http.ResponseWriter, h *multipart.FileHeader, msg string) {
	http.Error(w, fmt.Sprintf("400 invalid file %s: %s", h.Filename, msg), 400)
}

// Serve board-specific image banner files
func serveBanner(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(extractParam(r, "id"))
	if err != nil {
		text404(w)
		return
	}

	f, ok := assets.Banners.Get(extractParam(r, "board"), id)
	if !ok {
		text404(w)
		return
	}
	serveAssetFromMemory(w, r, f)
}

func serveAssetFromMemory(
	w http.ResponseWriter,
	r *http.Request,
	f assets.File,
) {
	if checkClientEtag(w, r, f.Hash) {
		return
	}

	h := w.Header()
	h.Set("ETag", f.Hash)
	h.Set("Content-Type", f.Mime)
	h.Set("Content-Length", strconv.Itoa(len(f.Data)))
	w.Write(f.Data)
}

// Serve board-specific loading animation
func serveLoadingAnimation(w http.ResponseWriter, r *http.Request) {
	serveAssetFromMemory(w, r, assets.Loading.Get(extractParam(r, "board")))
}
