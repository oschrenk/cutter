package internal

import (
	"time"
)

type Cookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	Expires  time.Time `json:"expires"`
	MaxAge   int       `json:"max_age"`
	Secure   bool      `json:"secure"`
	HttpOnly bool      `json:"http_only"`
	Creation time.Time `json:"creation"`
}
