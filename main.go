package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"crud/psql"
	"crud/api"
)
func main()  {
	fmt.Println("server run")

	gin.SetMode(gin.ReleaseMode)

	// db.AutoMigrate(&api.People{})
	psql.ConnectDB()
	api.RegisterRouter()
}

// https://www.youtube.com/watch?v=alcanyAiv14&list=PLYp_Kd32Xvcq2Qyxu8pnJ0ZGAz9ewBZ3H&index=8