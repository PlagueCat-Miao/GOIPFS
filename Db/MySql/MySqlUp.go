// +build ignore

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"net/http"


	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func init() {
	fmt.Println("主MySql节点上线");
}

//不大写，不public
type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required" `
	Password string `form:"password" json:"password" xml:"password" binding:"required" `
	Targetdb string `form:"targetdb" json:"targetdb" xml:"targetdb" binding:"required"`
	Dhash string `form:"Dhash" json:"Dhash" xml:"Dhash" binding:"required"`
	Status string `form:"status" json:"status" xml:"status" binding:"required"`
	Ip string `form:"ip" json:"ip" xml:"ip" binding:"required"`
	Capacity string `form:"capacity" json:"capacity" xml:"capacity" binding:"required"`
	Remain string `form:"remain" json:"remain" xml:"remain" binding:"required"`

}


func test1(c *gin.Context) {

	var json4 Login
	if err := c.ShouldBindJSON(&json4); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err })
		return
	}
	ClientIP:= c.ClientIP()

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"targetdb" : json4.Targetdb,
		"Dhash" : json4.Dhash,
		"status" : json4.Status,
		"ip" : json4.Ip,
		"ClientIP" : ClientIP,
		"capacity" : json4.Capacity,
		"remain" : json4.Remain,
		"success":true,
		})


}


func main() {
		router := gin.Default()

		// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
		//可以获取/:xxxx处的路径名称，并跳转至/xxxx
		router.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name) //以字符串的形式回复
		})

		// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
		// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	    //可以获取/:xxxx/*yyyy处的路径名称，并跳转至/xxxx
		router.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
		router.POST("/do/test1",test1)

		router.Run(":8080")


}


func Add_MYSQL() {
	conn, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/ipfs_netserver")
	if nil != err {
		fmt.Println("connect db error: ", err)
	}
	targetdb :="cloud"
	Dhash := "QmABC"
	status := "3"
	ip := "127.0.0.1"
	capacity :="400"
	remain :="300" ;

	var SET_VAL string
	SET_VAL="('"+ Dhash +"','"+status+"','"+ip+"','"+capacity+"','"+remain+"') "
	fmt.Println(targetdb +" 插入: "+ SET_VAL);

	_ , err =  conn.Query("INSERT INTO `ipfs_netserver`.`"+ targetdb +"` VALUES " +SET_VAL)
	if nil != err{
		fmt.Println("query db insert error: ", err.Error())
		return
	}
	err =conn.Close();
	if nil != err{
		fmt.Println("close db error: ", err.Error())
		return
	}
}
