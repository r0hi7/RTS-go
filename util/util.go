package util

import (
	"github.com/nishitm/RTS-go/config"
)

//Source interface is for implemententing methin GetSearchedTerm for every sources
type Source interface {
	GetSearchedTerm(config.Config)
}
