package main

import (
  // "googlemaps.github.io/maps"

  "database/sql"
  "fmt"
  "log"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  //first parameter is driver name of the database. This string registers itself
  //with datbaase/sql and is conventionally the same as the package name
  //second argument is driver specific syntax that tells the driver how to access
  //the underlying datastore. In this, we are connecting to the employeedb
  //database in our local database
  //format: username:password@tcp(ip_address:port)/database_schema

  //note that this does NOT establish any connection to the database, nor does it
  //validate any parameters. Instead, it prepares the database Go abstraction
  db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/album")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Connection Established")
  }
  pingErr := db.Ping()
  if pingErr == nil {
    fmt.Println("Database is available.")
  } else {
    fmt.Println(pingErr)
  }

  var (
    artist string
  )
  rows, err := db.Query("select artist from album;")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&artist) //read column into variable
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(artist)
  }
  err = rows.Err()
  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()

  // reader := bufio.NewReader(os.Stdin)
  // fmt.Print("Enter text: ")
  // str, err := reader.ReadString('\n')
  // if err != nil {
  //  fmt.Println(str)
  // } else {
  //  fmt.Println("Dumbass you entered some invalid shit")
  // }

}


