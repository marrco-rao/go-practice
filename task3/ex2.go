package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

// Use a simple singleton
var (
	db      *sqlx.DB
	initErr error
	once    sync.Once
)

func Init(dsn string) error {
	once.Do(func() {
		db, initErr = sqlx.Connect("mysql", dsn)
	})
	// test connection
	//err = db.Ping()
	if initErr == nil {
		fmt.Println("Database connected successfully")
	}
	return initErr
}

func GetDB() (*sqlx.DB, error) {
	if db == nil {
		return db, errors.New("db is nil")
	}
	return db, nil
}

/*
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
	编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

type Employee struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int64  `db:"salary"`
}

func QueryEmployees(department string) ([]Employee, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	employees := []Employee{}
	e := db.Select(&employees, "select * from employees where department = ?", department)
	if e != nil {
		fmt.Println("Query error: ", e)
	}
	return employees, nil
}
func FindTopSalaryEmployee() (Employee, error) {
	db, err := GetDB()
	if err != nil {
		return Employee{}, err
	}
	employee := Employee{}
	e := db.Get(&employee, "select * from employees order by salary desc limit 1")
	if e != nil {
		return Employee{}, e
	}
	return employee, err
}

// test employee
func TestEmployee() {
	employees, err := QueryEmployees("技术部")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("employees:")
	fmt.Println(employees)

	//
	employee, err := FindTopSalaryEmployee()
	if err != nil {
		return
	}
	fmt.Println("the employee with top salary:")
	fmt.Println(employee)
}

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
	定义一个 Book 结构体，包含与 books 表对应的字段。
	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Books struct {
	ID     sql.NullInt64   `db:"id"`
	Title  sql.NullString  `db:"title"`
	Author sql.NullString  `db:"author"`
	Price  sql.NullFloat64 `db:"price"`
}

func QueryBooks() ([]Books, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	books := []Books{}
	e := db.Select(&books, "select * from books")
	if e != nil {
		return nil, e
	}
	return books, nil
}
func QueryBooksByPrice(price float64) ([]Books, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	books := []Books{}
	e := db.Select(&books, "select * from books where price > ?", price)
	if e != nil {
		return nil, e
	}
	return books, nil
}

func TestBooks() {
	books, err := QueryBooksByPrice(50)
	if err != nil {
		return
	}
	for _, book := range books {
		fmt.Println(book)
	}
}

func main1() {
	//TestBase()
	dsn := "root:Ab123456@(127.0.0.1:3306)/test1"
	err2 := Init(dsn)
	if err2 != nil {
		return
	}

	//CheckSingleton()
	//TestEmployee()
	TestBooks()
}

// check singleton
func CheckSingleton() {
	db1, err := GetDB()
	if err != nil {
		fmt.Println(err)
	}
	db2, err := GetDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db2 == db1)
}

// Check the db connection
func TestBase() {
	db, err := sqlx.Connect("mysql", "root:Ab123456@(127.0.0.1:3306)/test1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	employees := []Employee{}
	err1 := db.Select(&employees, "select * from employees")
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	log.Println("employees:")
	log.Println(employees)

}
