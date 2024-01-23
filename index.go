package handler

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "欢迎使用我的头像加速！")
}
