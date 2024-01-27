package avatar

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 防盗链的域名，不带http或者https
var allowedDomain string = os.Getenv("domain")

func AvaterHandler(w http.ResponseWriter, r *http.Request) {
	// 加入防盗链，防止被盗刷
	referer := r.Header.Get("Referer")
	// 判定是否为http或者https开头
	if !strings.HasPrefix(referer, fmt.Sprint("http://", allowedDomain)) && !strings.HasPrefix(referer, fmt.Sprint("https://", allowedDomain)) {
		http.Error(w, "403 Forbidden", http.StatusForbidden)
		return
	}
	query := r.URL.Query()
	// 头像默认尺寸，没有就取80，加快速度，因为cravatar默认就是80
	s := query.Get("s")
	if s == "" {
		s = "80"
	}

	d := query.Get("d")
	if d == "" {
		d = "wavatar" // 卡通头像，我喜欢这个
	}

	gravatar := fmt.Sprintf("https://0.gravatar.com%s?s=%s&d=%s", r.URL.Path, s, d)
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
