package utils

import "net/http"

// GetClientIp returns the header forwarder by Trafik
func GetClientIp(r *http.Request) string {
	return r.Header.Get("X-Real-Ip")
}
