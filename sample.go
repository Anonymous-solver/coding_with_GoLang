package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// type User struct{
//     Name string `json: "title"`
// }

func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:amijanina@tcp(127.0.0.1:3306)/user")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    fmt.Println("Successfully connected MYSQL database!")

    //******************//


    // //post method
    // insert, err := db.Query("INSERT INTO books VALUES('Story', 500, 'Hindi','Porimol')")

    // if err != nil{
    //     panic(err)
    // }

    // defer insert.Close()

    // fmt.Println("Inserted successful!")

    //**********************//

    // //get method

    // results, err := db.Query("SELECT title FROM books")
    // if err != nil{
    //     panic(err)
    // }
    // for results.Next(){
    //     var user User
    //     err = results.Scan(&user.Name)
    //     if err != nil{
    //         panic(err)
    //     }
    //     fmt.Println(user.Name)
    // }

    //*********************//

    //update method

    // results, err := db.Query("UPDATE books SET author = 'Shiraz' WHERE title = 'Book title'")
    // if err != nil{
    //     panic(err)
    // }

    //  fmt.Println("Updated successful!")

    //  defer results.Close()

    //**************************//

    //delete method

    // results, err := db.Query("DELETE FROM books WHERE title = 'Book title'")
    // if err != nil{
    //     panic(err)
    // }

    //  fmt.Println("Deleted successful!")

    //  defer results.Close()

}