package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Info struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func getProjetoKorp(c *gin.Context) {
	horario := Info{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.TimeOnly),
	}
	c.IndentedJSON(http.StatusOK, horario)
}

func main() {
	router := gin.Default()
	router.GET("/projeto-korp", getProjetoKorp)

	router.Run("0.0.0.0:8080")
}
