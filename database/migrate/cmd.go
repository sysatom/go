package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./database/migrate")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	source := viper.GetStringMapString("source")
	target := viper.GetStringMapString("target")

	sourceDB := sqlx.MustConnect("mysql", source["url"])
	err = sourceDB.Ping()
	if err != nil {
		panic(err)
	}

	targetDB := sqlx.MustConnect("mysql", target["url"])
	err = targetDB.Ping()
	if err != nil {
		panic(err)
	}

	tables := viper.GetStringMap("tables")
	for table, fields := range tables {
		fmt.Println(table)
		fmt.Println(fields.([]interface{}))

		row := sourceDB.QueryRowx(fmt.Sprintf("select * from %s limit 1", table))

		columns, _ := row.Columns()
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}
		//var desc interface{}
		err := row.Scan(values...)
		if err != nil {
			fmt.Println(err)
		}

		m := make(map[string]interface{})
		for i, col := range columns {
			val := values[i].(*interface{})
			m[col] = val
		}
		fmt.Println(m)

	}
}
