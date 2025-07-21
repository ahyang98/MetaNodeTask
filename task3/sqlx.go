package task3

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Employees struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func query(db *sqlx.DB) {
	var employees []Employees
	err := db.Select(&employees, "select * from employees where department = '技术部'")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", employees)
	var topEmployees []Employees
	err = db.Select(&topEmployees, `
    SELECT * FROM employees 
    WHERE salary = (SELECT MAX(salary) FROM employees)
`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", topEmployees)
}

func initDb() *sqlx.DB {
	dsn := "root:ahyang@tcp(192.168.252.128:30306)/blog?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func TestEmployee() {
	db := initDb()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			println(err)
		}
	}(db)
	query(db)
}

type Book struct {
	ID     int             `db:"id"`
	Title  string          `db:"title"`
	Author string          `db:"author"`
	Price  decimal.Decimal `db:"price"`
}

func queryBook(db *sqlx.DB) {
	var books []Book
	err := db.Select(&books, "select * from books where price > 50")
	if err != nil {
		println(err)
		return
	}
	fmt.Printf("%+v\n", books)
}

func TestBook() {
	db := initDb()
	defer db.Close()
	queryBook(db)
}
