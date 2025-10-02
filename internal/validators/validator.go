package validators

import (
	"net/url"
	"strings"
)

func IsValidUrl(userUrl string) bool {

	if !strings.HasPrefix(userUrl, "http://") && !strings.HasPrefix(userUrl, "https://") {
		userUrl = "https://" + userUrl
	}

	_, err := url.ParseRequestURI(userUrl)
	return err == nil
}
