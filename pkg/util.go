package pkg

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func wrapUrl(host, api string) string {
	return fmt.Sprintf(
		"%s/%s",
		strings.TrimRight(host, "/"),
		strings.TrimLeft(api, "/"),
	)
}

func debugHttpResp(resp *http.Response, title string) error {
	if log.Debug().Enabled() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return errors.WithStack(err)
		}
		log.Debug().Msgf("%s resp:\n===\n%s\n===", title, dump)
	}
	return nil
}
