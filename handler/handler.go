package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Halo dari gin",
	})
}

func RootHandlerJsonOnline(c *gin.Context) {

	url := "https://jsonplaceholder.typicode.com/posts"
	fmt.Println("URL:>", url)

	var query = []byte(`your query`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(query))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	c.JSON(200, string(body))
}

func PostHandler(c *gin.Context) {
	var json struct {
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		c.JSON(200, gin.H{"message": json.Message})
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
