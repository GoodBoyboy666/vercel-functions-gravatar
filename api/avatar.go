package avatar

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func AvaterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md5 := vars["md5"]
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, md5)
}
