package imager

import (
	"fmt"
	"io/ioutil"
	"meguca/common"
	"meguca/config"
	"meguca/imager/assets"
	"meguca/test"
	"os"
	"path/filepath"
	"testing"

	"github.com/bakape/thumbnailer"
)

func TestImageProcessing(t *testing.T) {
	config.Set(config.Configs{
		MaxWidth:    2000,
		MaxHeight:   2000,
		JPEGQuality: 80,
	})

	cases := [...]struct {
		ext   string
		dims  [4]uint16
		isPNG bool
	}{
		{"jpg", assets.StdDims["jpeg"], false},
		{"png", assets.StdDims["png"], true},
		{"gif", assets.StdDims["gif"], true},
	}

	for i := range cases {
		c := cases[i]
		t.Run(c.ext, func(t *testing.T) {
			t.Parallel()

			var img common.ImageCommon
			thumb, err := processFile(
				test.ReadSample(t, "sample."+c.ext),
				&img,
				thumbnailer.Options{
					ThumbDims: thumbnailer.Dims{
						Width:  150,
						Height: 150,
					},
					JPEGQuality: 90,
				},
			)
			if err != nil {
				t.Fatal(err)
			}

			assertThumbnail(t, thumb)
			assertDims(t, img.Dims, c.dims)

			thumbType := common.JPEG
			if c.isPNG {
				thumbType = common.PNG
			}
			assertFileType(t, img.ThumbType, thumbType)

			var thumbExt string
			if img.ThumbType == common.PNG {
				thumbExt = "png"
			} else {
				thumbExt = "jpg"
			}
			t.Logf(`dims: %dx%d`, img.Dims[2], img.Dims[3])
			writeSample(t, fmt.Sprintf("thumb_%s.%s", c.ext, thumbExt), thumb)
		})
	}
}

func writeSample(t *testing.T, name string, buf []byte) {
	t.Helper()

	path := filepath.Join("testdata", name)

	// Remove previous file, if any
	_, err := os.Stat(path)
	switch {
	case os.IsExist(err):
		if err := os.Remove(path); err != nil {
			t.Fatal(err)
		}
	case os.IsNotExist(err):
	case err == nil:
	default:
		t.Fatal(err)
	}

	err = ioutil.WriteFile(path, buf, 0600)
	if err != nil {
		t.Fatal(err)
	}
}
