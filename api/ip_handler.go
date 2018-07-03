package api

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/v2af/ipify/models"
)

func GetIP(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	var ip string

	if ip_bytes := net.ParseIP(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]); ip_bytes != nil {
		ip = ip_bytes.String()
	} else {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	if format, ok := r.Form["format"]; ok && len(format) > 0 {
		jsonStr, _ := json.Marshal(models.IPAddress{ip})

		switch format[0] {
		case "json":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(jsonStr))
			return
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, ip)
}
