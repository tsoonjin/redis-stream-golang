package handler

import (
	"net/http"
    "github.com/tsoonjin/redis-stream-golang/service/redis"
    "fmt"
)

type CommandHandler struct {
    RedisClient *service.Redis
}

func (h *CommandHandler) Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Are you ready to rumble ?")
}
