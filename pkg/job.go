package pkg

import (
	"crypto/tls"
	"net/http"

	"github.com/pkg/errors"
)

func init() {
	// enable insecure https
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http.DefaultClient.Transport = tr
}

func ValidateSSO() error {
	httpResp, err := http.Get(wrapUrl(Config.Host, "/v1/api/sso/validate"))
	if err != nil {
		return errors.WithStack(err)
	}
	defer httpResp.Body.Close()
	if err := debugHttpResp(httpResp, "validate sso"); err != nil {
		return err
	}
	if httpResp.StatusCode != http.StatusOK {
		return errors.Errorf("validate sso status: %v", httpResp.StatusCode)
	}
	return nil
}

func GetAuthenticationStatus() error {
	httpResp, err := http.Post(wrapUrl(Config.Host, "/v1/api/iserver/auth/status"), "", nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer httpResp.Body.Close()
	if err := debugHttpResp(httpResp, "get authentication status"); err != nil {
		return err
	}
	return nil
}
