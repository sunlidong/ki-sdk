package api

import (
	r_action "ki-sdk/router/action"
	g "github.com/gin-gonic/gin"
)

func InitRouter()(*g.Engine){
	return	r_action.InitRouter()
}