package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
	"bufio"
	"os"
)

func main() {
	db , err := sql.Open("mysql","root:@/golang")
	if err != nil {
		log.Panic(err)
	}
	error := db.Ping()
	if error != nil {
		log.Panic(error)
	}
	defer db.Close()
	stmt , er :=  db.Prepare("INSERT INTO harsh VALUES (?,?)")
	if er != nil {
		log.Panic(er)
	}
	defer stmt.Close()
	fmt.Print("Enter Name : ")
	inp , _ := bufio.NewReader(os.Stdin).ReadBytes('\n')
	name :=  string(inp[0:len(inp)-1])
	fmt.Print("Enter Phone : ")	
	inp , _  = bufio.NewReader(os.Stdin).ReadBytes('\n')
	phone :=  string(inp[0:len(inp)-1])
	_ , e := stmt.Exec(name,phone)
	if e != nil {
		log.Panic(e)
	}


}

//use Exec command for query other than select
// use Query command for select query
// Result is returned when Exec is used
// *Rows is returned when Query is used
// RowsAffected method of Result interface is used to know the no of rowsaffected
// Next() method of Rows is same as it works for resultset in java
// Scan method of Rows interface is used to read values
// data source name is passed as second argument of Open method 
//"root:@/golang" is a dsn( data source name)
//"root:@/golang" can be also written as "root:@tcp(localhost:3306)/golang"
// because tcp(localhost:3306) are default so it doesn't matter whether we write it or not
//if username and password is different then you can write "username:password@/golang" 
//if you have  protocol except tcp or want to connect other than localhost then write as
//"username:password@protocol(ipaddress:portnumber)/databasename"
// don't get confused by golang in the dsn("/golang") -> it is name of my database used
//in place of ipaddress you can also write domain name