package MySql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func init() {
	fmt.Println("test Testmysql_test.go（本文件）时，init()（此函数）执行")

}

func TestIpfsAPI(t *testing.T) {
	t.Log("Hellcat: MySql测试")
	t.Log("测试内容：MySql连通性")
	conn, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/ipfs_netserver")
	if nil != err {
		fmt.Println("connect db error: ", err)
	}
	rows, err := conn.Query("SELECT * FROM ipfs_netserver.ownership;")
	if nil != err {
		fmt.Println("query db error: ", err.Error())
		return
	}
	t.Log("Hellcat: map 不能为空 需要make")
	var temp_store map[int]string
	temp_store = make(map[int]string)
	fmt.Println("也不知道这个rows打印了个啥？")
	fmt.Println(rows)

	for i := 0; rows.Next(); i++ {
		var Dhash, Fhash string
		//Scan将值存入变量中，注意一行变量有几列就要有几个变量接住
		err = rows.Scan(&Dhash, &Fhash)
		if err != nil {
			panic("注意Panic:" + err.Error())
		}
		t.Log("Hellcat: map 插入")
		//存入temp_store(map)
		temp_store[i] = Dhash
		i++
		temp_store[i] = Fhash

	}
	t.Log("Hellcat: map 遍历")
	//获取
	for k, v := range temp_store {
		fmt.Println(k, "----", v)
	}

	t.Log("Hellcat: MySql 测试结束")
}
