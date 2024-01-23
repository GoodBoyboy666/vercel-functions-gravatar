package avatar

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const allowedDomain string = "https://blog.del.pub"

func AvaterHandler(w http.ResponseWriter, r *http.Request) {
	// 加入防盗链
	referer := r.Header.Get("Referer")
	if false && !strings.HasPrefix(referer, allowedDomain) {
		http.Error(w, "403 Forbidden", http.StatusForbidden)
		return
	}
	query := r.URL.Query()
	gravatar := fmt.Sprintf("https://0.gravatar.com%s?s=%s&d=%s", r.URL.Path, s, query.Get("d", "wavatar"), query.Get("s", 100))
	resp, err := http.Get(gravatar)
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	// 没有G级头像
	if resp.StatusCode == http.StatusNotFound {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	io.Copy(w, resp.Body)
}
