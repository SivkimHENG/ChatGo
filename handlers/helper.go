package handlers;

import (
    "net/http"
    "log"
    "golang.org/x/exp/slog"

)

type HTTPHandler func (w http.ResponseWriter, r * http.Request) error;



func Make (h HTTPHandler) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        if err := h(w,r); err != nil {
            slog.Error("HTTP Handler error", "err", err, "path", r.URL.Path);

        }

    }

}





