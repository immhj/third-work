package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type warehouse struct {
	Num   string
	Space int
}

type cloth struct {
	Num   string
	Size  string
	Price int
	Tp    string
}
type supplier struct {
	Num  string
	Name string
}
type supplier_env struct {
	Cnum   string
	Snum   int
	Qulify string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/infomation?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db = db.AutoMigrate(&warehouse{})
	u1 := warehouse{"A", 10}
	u2 := warehouse{"B", 20}
	db.Create(&u1)
	db.Create(&u2)
	db.AutoMigrate(&cloth{})
	u3 := cloth{"A", "S", 90, "A"}
	u4 := cloth{"B", "M", 120, "B"}
	db.Create(&u3)
	db.Create(&u4)
	db.AutoMigrate(&supplier{})
	u5 := supplier{"1", "A"}
	db.Create(&u5)
	db.AutoMigrate(&supplier_env{})
	u6 := supplier_env{"A", 1, "不合格"}
	db.Create(&u6) //写入数据*/

	//查询

	var s_max int
	var warehouses []warehouse
	db.Find(&warehouses)
	s_max = warehouses[0].Space
	for i := 0; i < len(warehouses); i++ {
		if warehouses[i].Space > s_max {
			s_max = warehouses[i].Space
		}
	}

	var cloths []cloth
	db.Find(&cloths)
	for i := 0; i < len(cloths); i++ {
		if cloths[i].Size == "S" && cloths[i].Price < 100 {
			fmt.Println(cloths[i].Num, cloths[i].Price, cloths[i].Size, cloths[i].Tp)

		}
	}
	var s int
	for i := 0; i < len(cloths); i++ {
		if cloths[i].Size == "S" {
			s += 1
		}
	}
	fmt.Printf("A类服装的总量为%d", s)
	for i := 0; i < len(cloths); i++ {
		if cloths[i].Num == "A" {
			fmt.Println("服装的编码以A开头的服装有：")
			fmt.Println(cloths[i].Num, cloths[i].Price, cloths[i].Size, cloths[i].Tp)
		}
	}
	var ps []supplier_env
	db.Find(&ps)
	for i := 0; i < len(ps); i++ {
		if ps[i].Qulify == "不合格" {
			fmt.Println("服装质量不合格的供应商信息如下：")
			fmt.Println(ps[i].Cnum, ps[i].Qulify, ps[i].Snum)
		}
	}
	//更新
	for i := 0; i < len(cloths); i++ {
		if cloths[i].Size == "S" {
			cloths[i].Price = cloths[i].Price * 11 / 10
		}
	}
	//删除

	for i := 0; i < len(ps); i++ {
		if ps[i].Qulify == "不合格" {
			db.Delete(&ps[i])
		}
	}

}
