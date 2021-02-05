package pkg

import (
	"fmt"
	"strings"
)

func wrapUrl(host, api string) string {
	return fmt.Sprintf(
		"%s/%s",
		strings.TrimRight(host, "/"),
		strings.TrimLeft(api, "/"),
	)
}
