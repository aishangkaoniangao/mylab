package main

import "fmt"
import "os"
import "github.com/ziutek/mymysql/mysql"
import "github.com/ziutek/mymysql/native"
//import "reflect"
import "strconv"

func main() {
    delete()
}

func insert() {
    //fields := []string{"id", "username", "password", "create_at", "update_at"}

    user := "root"
    pass := "123456"
    dbname := "blog"
    db := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)
    db.Connect()
    
    stmt, err := db.Prepare("insert into user (username, password) values(?, ?)")
    checkError(err)

    type Data struct {
        Username string
        Password string
    }

    data := new(Data)
    data.Username = "test1234"
    data.Password = "test5678"

    _, err = stmt.Run(data.Username, data.Password)
    checkError(err)

    err1 := db.Close()
    if err1 != nil {
        panic(err)
    }
}

func update() {
    //fields := []string{"id", "username", "password", "create_at", "update_at"}

    user := "root"
    pass := "123456"
    dbname := "blog"
    db := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)
    db.Connect()
    
    var id int = 1
    
    stmt, err := db.Prepare("update user set username=?, password=? where id=?")
    checkError(err)

    type Data struct {
        Id       int
        Username string
        Password string
    }

    data := new(Data)
    data.Id = id
    data.Username = "username"
    data.Password = "password"

    _, err = stmt.Run(data.Username, data.Password, data.Id)
    checkError(err)

    err1 := db.Close()
    if err1 != nil {
        panic(err)
    }
}

func delete() {
    //fields := []string{"id", "username", "password", "create_at", "update_at"}

    user := "root"
    pass := "123456"
    dbname := "blog"
    db := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)
    db.Connect()
    
    var id int = 2
    
    stmt, err := db.Prepare("delete from user where id=?")
    checkError(err)

    type Data struct {
        Id       int
        Username string
        Password string
    }

    data := new(Data)
    data.Id = id
    data.Username = "username"
    data.Password = "password"

    _, err = stmt.Run(data.Id)
    checkError(err)

    err1 := db.Close()
    if err1 != nil {
        panic(err)
    }
}

func getRow() {
    fields := []string{"id", "username", "password", "create_at", "update_at"}

    user := "root"
    pass := "123456"
    dbname := "blog"
    db := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)
    db.Connect()
    
    var id int = 1
    
    rows, res, err := db.Query("select * from user where id=%d", id)
    if err != nil {
        panic(err)
    }
    
    var fieldType map[string]byte = map[string]byte{}
    //var fieldIndex map[int]string = map[int]string{}
    for _, col1 := range res.Fields() {
        fieldType[col1.OrgName] = col1.Type
        //fieldIndex[]
    }
    
    fieldIndex := map[int]string{}
    for _, v := range fields {
        fmt.Println(v, res.Map(v))
        fieldIndex[res.Map(v)] = v
    }

    //还缺少一个index=>fieldName的转换结构

    for _, row := range rows {
        for k, col := range row {
            fieldName := fieldIndex[k]

            if (fieldType[fieldName] == native.MYSQL_TYPE_LONG) {
                var v int

                byteCol, ok := col.([]byte)
                if ok == false {
                    panic(ok)
                }

                v, _ = strconv.Atoi(string(byteCol))
                fmt.Println(fieldName, ":", v)
            }
            if (fieldType[fieldName] == native.MYSQL_TYPE_VAR_STRING) {
                var v string
                byteCol, ok := col.([]byte)
                if ok == false {
                    panic(ok)
                }
                
                v = string(byteCol)
                fmt.Println(fieldName, ":", v)
            }
            if (fieldType[fieldName] == native.MYSQL_TYPE_TIME) {
                fmt.Println(fieldName, col)
            }
            if (fieldType[fieldName] == native.MYSQL_TYPE_DATETIME) {
                fmt.Println(fieldName, col)
            }
            if (fieldType[fieldName] == native.MYSQL_TYPE_TIMESTAMP) {
                fmt.Println(fieldName, col)
            }            
        }
    }
    fmt.Println("fieldType:", fieldType)
    
    err1 := db.Close()
    if err1 != nil {
        panic(err)
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }   
}