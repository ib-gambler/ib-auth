package pkg

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func ValidateSSO() error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http.DefaultClient.Transport = tr

	httpResp, err := http.Get(wrapUrl(Config.Host, "/v1/api/sso/validate"))
	if err != nil {
		return errors.WithStack(err)
	}
	defer httpResp.Body.Close()

	if log.Debug().Enabled() {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err != nil {
			return errors.WithStack(err)
		}
		log.Debug().Msgf("ValidateSSO resp %s", dump)
	}

	return nil
}
