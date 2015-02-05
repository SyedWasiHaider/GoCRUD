package main

import "log"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main(){
  db, err := sql.Open("mysql", "YOURUSERNAME:YOURPASSWORD@/items_database")
  if err != nil {
    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
}
  defer db.Close()

  rows, err := db.Query("Select * from items");
  if err != nil {
    log.Fatal(err)
  }

  defer rows.Close()

  var id * int;
  var name, description, price * string;

  for rows.Next() {
    err := rows.Scan(&id, &name, &description, &price)
    if err != nil {
      log.Fatal(err)
    }
    log.Println(*id, *name, *description, *price)
  }


}
