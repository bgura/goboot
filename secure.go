package goboot

import (
    "net/http"
)

func SecureHeaders(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      // Disallow IE and Chrome from sniffing response away from actual type
      w.Header().Add("X-Content-Type-Options", "nosniff")

      // Add XSS Protection
      w.Header().Add("X-XSS-Protection", "1; mode=blockFilter")
      // Disable render in an IFrame
      //w.Header().Add("X-Frame-Options", "DENY")
    })
}
