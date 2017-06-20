package main

import (
	"log"
	"net/http"
	"fmt"
)

type SyncUser struct {
	Id			int64		
	Name		string		`xorm:"unique"`
	Contact		string		
}

var engine *xorm.Engine

func postgresEngine() (*xorm.Engine, error){
	return xorm.NewEngine("postgres", "dbname=golang user=golang password=123123 sslmode=disable")
}

type engineFunc func() (*xorm.Engine, error)

func sync(engine *xorm.Engine) error {
	return engine.Sync(&SyncUser{}, &SyncUser{})
}

func main() {
	var router = NewRouter()

	//engine, err = xorm.NewEngine("postgres", "dbname=golang sslmode=disable")

	engines := []engineFunc{postgresEngine}

	for _, enginefunc := range engines {
		Orm, err := enginefunc()
		fmt.Println("--------", Orm.DriverName(), "----------")
		if err != nil {
			fmt.Println(err)
			return
		}
		Orm.ShowSQL(true)
		err = sync(Orm)
		if err != nil {
			fmt.Println(err)
		}

		_, err = Orm.Where("id > 0").Delete(&SyncUser{})
		if err != nil {
			fmt.Println(err)
		}

		user := &SyncUser{
			Name: "BLALBA",
			Contact: "AKDAD",
		}
		_, err = Orm.Insert(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		isexist, err := Orm.IsTableExist("Users")
		if err != nil {
			fmt.Println(err)
			return
		}
		if !isexist {
			fmt.Println("sync_user2 is not exist")
			return
		}
	}

	log.Fatal(http.ListenAndServe(":8081", router))
}