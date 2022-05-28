package main

import (
	"fmt"
	"github.com/GenericP3rson/TigerGo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	conn := TigerGo.TigerGraphConnection{
		Token:     token(),
		Host:      "https://SUBDOMAIN.i.tgcloud.io",
		GraphName: "GRAPHNAME",
		Username:  "tigergraph",
		Password:  password(),
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/echo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": conn.Echo(),
		})
	})

	r.POST("/join", func(c *gin.Context) {

		usr := c.Query("user")

		fmt.Println(usr)

		fmt.Println(conn.UpsertVertex("User", usr, map[string]string{"user": usr}))

	})

	r.POST("/add-lang", func(c *gin.Context) {

		usr := c.Query("user")
		lang := c.Query("lang")
		exp := c.Query("experience")

		fmt.Println(usr)

		fmt.Println(conn.UpsertVertex("User", usr, map[string]string{"user": usr}))
		fmt.Println(conn.UpsertVertex("Language", lang, map[string]string{"language": lang}))
		fmt.Println(conn.UpsertEdge("User", usr, "USER_LANGUAGE", "Language", lang, map[string]string{"experience": exp}))

	})

	r.POST("/add-interest", func(c *gin.Context) {

		usr := c.Query("user")
		interest := c.Query("interest")

		fmt.Println(usr)

		fmt.Println(conn.UpsertVertex("User", usr, map[string]string{"user": usr}))
		fmt.Println(conn.UpsertVertex("Interest", interest, map[string]string{"interest": interest}))
		fmt.Println(conn.UpsertEdge("User", usr, "USER_INTEREST", "Interest", interest, nil))

	})

	r.GET("/suggest-teammates", func(c *gin.Context) {

		usr := c.Query("user")

		c.JSON(http.StatusOK, gin.H{
			"results": conn.RunInstalledQuery("suggestTeammates", map[string]string{"username": usr}),
		})

	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
