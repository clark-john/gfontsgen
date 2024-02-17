package font

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/clark-john/gfontsgen/utils"
	"github.com/fatih/color"
)

const NormalPerm os.FileMode = 0666

type DownloadOptions struct {
	Url string
	Path string
	Family string
	Variant string
}

func DownloadFile(options DownloadOptions) bool {
	_url := options.Url
	_path := options.Path
	_family := options.Family
	_varnt := options.Variant

	_, err := os.Open(_path)
	if err != nil {
		os.Mkdir(_path, NormalPerm)
	}
	u, _ := url.Parse(_url)

	ext := path.Ext(u.Path)
	
	file := CreateFileName(_family, _varnt, ext)

	resp, _ := http.Get(_url)
	data, _ := io.ReadAll(resp.Body)

	err = os.WriteFile(path.Join(_path, file), data, NormalPerm)

	if err != nil {
		return false
	}

	color.HiGreen("%s successfully downloaded", file)
	return true
}

func CreateFileName(family string, vrnt string, ext string) string {
	var s strings.Builder
	s.WriteString(family + "-")

	isRegOrIt, _ := regexp.MatchString("^[ri]", vrnt)
	isItalic, _ := regexp.MatchString("italic$", vrnt)

	if isRegOrIt {
		s.WriteString(utils.Capitalize(vrnt))
	} else {
		if isItalic {		
			wght := vrnt[0:3]
			style := vrnt[3:]
			s.WriteString(wght + utils.Capitalize(style))
		} else {
			s.WriteString(vrnt + "Regular")
		}
	}

	s.WriteString(ext)

	return s.String()
}
