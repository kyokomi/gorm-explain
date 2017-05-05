package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysqlを使う
	"github.com/jinzhu/gorm"
	explain "github.com/kyokomi/gorm-explain"
)

type options struct {
	user     string
	password string
	host     string
	port     int
	dbName   string
	location string
}

func buildDataSourceName(opts options) string {
	const format = "%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=%s"
	return fmt.Sprintf(format, opts.user, opts.password, opts.host, opts.port, opts.dbName, opts.location)
}

// Product product table struct
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	opts := options{
		user:     "test-user",
		password: "test-user",
		host:     "127.0.0.1",
		port:     3306,
		dbName:   "test_db",
		location: "UTC",
	}
	db, err := gorm.Open("mysql", buildDataSourceName(opts))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db = db.Debug() // query trace
	db.Callback().Query().Register("explain", explain.Callback)

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}
