package outer

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func Download(url, baseDir string) (string, error) {

	err := os.MkdirAll(baseDir, 0700)
	if err != nil && !os.IsExist(err) {
		return "", err
	}

	imageName := strings.Split(path.Base(url), "?")[0]
	imageLocalPath := path.Join(baseDir, imageName)
	out, err := os.Create(imageLocalPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return imageLocalPath, err
}
