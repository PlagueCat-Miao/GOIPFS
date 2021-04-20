package MySql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

	"testing"
)

func init() {
	log.Println("Gorm连接数据库MySql 测试")
}

//由于grom是使用的orm映射，所以需要定义要操作的表的model，在go中需要定义一个struct， struct的名字就是对应数据库中的表名
//targetdb string

//对应的字段名的查找会先按照tag里面的名称去里面查找，如果没有定义标签则按照struct定义的字段查找，查找的时候struct字段中的大写会被转义成“ ”,例“GroupZone”会去查找表中的group_zone字段
//https://www.jb51.net/article/151051.htm
type Cloud struct {
	Dhash string `gorm:"column:Dhash;primary_key"` //一定要声明主键
	//1。 首先gorm是不知道主键、外键等tag的，很多操作可能导致gorm语法级别的err
	//删除Delete、修改Model.Update都是仅参考输入的主键，（剩下的靠Where()）
	//如果声明struct时未声明 `gorm:"primary_key"` 那即使次字段是数据库的主键
	// gorm 也会认为没有主键，而导致全选（相当于Model($xx{}).update）
	Status   int //此处必须大写，但是没写tag,所以gorm会转义为小写
	Ip       string
	Capacity int
	Remain   int
	//gorm.Model //这个在删除的时候用
}

func TestGorm_CDUR(t *testing.T) {
	t.Log("Hellcat: gorm-MySql测试")
	t.Log("测试内容：MySql连通性")

	db, err := gorm.Open("mysql", "root:12345678@tcp(localhost:3306)/ipfs_netserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	db.SingularTable(true) //让grom转义struct名字的时候不用加上s//https://www.jb51.net/article/151051.htm
	fmt.Println("单数表模式指，首字母大写的struct 不会")
	fmt.Println("结构体Cloud是否连接上MySql表cloud ", db.HasTable(Cloud{}))
	//db=db.Table("cloud")
	// 自动迁移模式
	//gorm可以支持自动迁移,也就是自动的表结构迁移,只会创建表,补充缺少的列,缺少的索引。
	//db.AutoMigrate(&Cloud{})

	// 读取Read
	var cadd, cadd2 Cloud
	err = db.Where("Dhash = ?", "QmGorm_新建喵").Find(&cadd).Error

	//db.First(&cadd2, 1) // 获取第一条数据,可以接在任何DB后(如where().First)，也可以自己写一些条件
	db.First(&cadd2, "Dhash = ?", "QmGorm_新建喵") // 获取第一条数据,可以接在任何DB后(如where().First)，也可以自己写一些条件
	// 查询code为l1212的product

	log.Println(cadd2)

	// 更新Update - 更新数据库
	err = db.Model(&cadd2).Updates(map[string]interface{}{"ip": "127.0.0.1:3306", "Capacity": 500}).Error
	// UPDATE Cloud SET ip='127.0.0.1:3306', Capacity=500 WHERE Dhash = kkk;(这是因为主键)
	err = db.Model(&Cloud{}).Update("remain", 140).Error
	// UPDATE Cloud SET remain=140 (这是因为主键)
	// 你要是想改主键 ，就删掉重新建立
	db.Delete(&cadd2)
	cadd2.Dhash = "Qm一定要改回主键喵"
	db.Save(&cadd2)
	if err != nil {
		log.Println("更新失败： ", err)
		fmt.Println("这种失败会导致 go的struct数据，改变而数据库未改变,一定要注意err的返回啊")

	} else {
		log.Println("Update success\n", ">>>>>> A\n", cadd, "\n======\n", cadd2, "\n>>>>>> B")
		fmt.Println("即使数据库成update功,go的struct旧副本没有更新（A:cadd B:cadd2）,一定要注意用最新的副本")
	}

	// 创建 Create
	db.Create(&Cloud{

		Dhash:    "QmGorm_新建喵",
		Status:   3,
		Ip:       "127.0.0.1",
		Capacity: 400,
		Remain:   300})
	db.Save(&Cloud{

		Dhash:    "QmGorm_+w覆盖喵",
		Status:   3,
		Ip:       "127.0.0.2",
		Capacity: 400,
		Remain:   300})

	//删除 Delete  - 删除product
	//db.Delete(&cadd2)
	t.Log("Hellcat: MySql 测试结束")
}
