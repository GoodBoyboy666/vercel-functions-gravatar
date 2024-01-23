package handler

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*") // 解决跨域问题
	w.Header().Set("content-type", "application/plain")
	io.WriteString(w, "欢迎使用我的头像加速！")
}
