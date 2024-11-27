package endpoints

import (
	"github.com/VitalyCone/kuznecov_coins_api/internal/app/store"
	"github.com/gin-gonic/gin"
)

type Endpoints struct {
	store *store.Store
}

func NewEndpoints(s *store.Store) *Endpoints {
	return &Endpoints{
		store: s,
	}
}

func (ep *Endpoints) Ping(g *gin.Context){
	g.JSON(200, "PING")
}