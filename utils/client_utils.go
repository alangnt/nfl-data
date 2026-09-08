package utils

import (
	"net/http"
	"time"
)

var sportradarClient = &http.Client{
	Timeout: 10 * time.Second,
}
