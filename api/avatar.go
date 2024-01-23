package avatar

import (
	"fmt"
	"io"
	"net/http"
)

func AvaterHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "你好，世界")
}
