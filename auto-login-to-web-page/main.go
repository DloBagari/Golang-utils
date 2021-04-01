package auto_login_to_web_page

import (
	"net/http"
	"time"
)

type LogsDownloader struct {
	LastDownloadTimeWeb   time.Time
	LastDownloadTimeEmail time.Time
	UserName              string
	Password              string
	LogsUrl               string
	*http.Client
}

func (d *LogsDownloader) Do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(d.UserName, d.Password)
	return http.DefaultClient.Do(req)
}

// use LogsDownloader.Do(req)
