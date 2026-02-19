package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type number struct {
	Number1 int `form:number1,default=10`
	Number2 int `form:number2`
	Number3 int `form:number3`
}
type calc struct {
	Number1  int    `json:number1`
	Number2  int    `json:number2`
	Operator string `json:operator`
}
type result struct {
	Result int `json:result`
}

func helloworld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "helloworld",
	})
}
func readnum(c *gin.Context) {
	var req number
	if err := c.ShouldBind(&req); err != nil { // infers binder by Content-Type
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}
func calculator(c *gin.Context) {
	var req calc
	var res result
	if err1 := c.ShouldBindJSON(&req); err1 != nil { // infers binder by Content-Type
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	fmt.Println(req.Number1)
	fmt.Println(req.Number2)
	switch req.Operator {
	case "+":
		res.Result = req.Number1 + req.Number2
	case "-":
		res.Result = req.Number1 - req.Number2
	case "*":
		res.Result = req.Number1 * req.Number2
	case "/":
		if req.Number2 == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Result": "Zero division error"})
			return
		} else {
			res.Result = req.Number1 / req.Number2
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"Result": "invalid operation"})
		return
	}
	fmt.Println(res.Result)
	c.JSON(http.StatusOK, gin.H{"Result": res.Result})
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	r.GET("/hello", helloworld)
	r.POST("/readnum", readnum)
	r.POST("/calculator", calculator)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()

}
