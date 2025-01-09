package httpserver

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const reverseProxyForwardedByHeader = "X-Forwarded-For"

func ClientIp(req *http.Request) string {
	ipAddress := req.RemoteAddr
	fwdAddress := req.Header.Get(reverseProxyForwardedByHeader)
	if fwdAddress != "" {
		ipAddress = fwdAddress

		ips := strings.Split(fwdAddress, ", ")
		if len(ips) > 1 {
			ipAddress = ips[0]
		}
	}

	return ipAddress
}

func CloneRequest(r *http.Request) *http.Request {
	var bodyBytes []byte
	newRequest := *r.WithContext(r.Context())

	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	newRequest.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return &newRequest
}

func RequestValidationJsonSchemaPath(basePath string, schemaPath string) string {
	return fmt.Sprintf("%s/request/%s", strings.TrimSuffix(basePath, "/"), schemaPath)
}
