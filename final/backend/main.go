package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	//json tag to serialize json body
	Name string `json:"name"`
}

func main() {
	engine := gin.New()
	engine.POST("/test", func(context *gin.Context) {
		body := Body{}
		//using blindJson metho to serialize body with struct
		if err := context.BindJSON(&body); err != nil {
			fmt.Println(err.Error())
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		context.JSON(http.StatusAccepted, &body)
	})
	engine.Run(":3000")
}
