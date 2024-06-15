package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/mati-fp/go-hexagonal/adapters/db"
	"github.com/mati-fp/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	createTable := `CREATE TABLE products (
		"id" string PRIMARY KEY,
		"name" string,
		"price" FLOAT,
		"status" string
	);`

	stmt, err := db.Prepare(createTable)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values ("abc", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer DB.Close()

	productDb := db.NewProductDb(DB)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer DB.Close()

	productDb := db.NewProductDb(DB)
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25.0

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, "Product Test", productResult.GetName())
	require.Equal(t, 25.0, productResult.GetPrice())
	require.Equal(t, "disabled", productResult.GetStatus())

	fmt.Println(productResult.GetStatus())

	product.Status = application.ENABLED
	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, application.ENABLED, productResult.GetStatus())

}
