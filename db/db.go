package db
 
import (
    "database/sql"
    "fmt"
	  _ "github.com/lib/pq"
	  "api"
	  "log"
  
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "Postgres"
    dbname   = "postgres"
)

 
func Save( p api.Activity) {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
   fmt.Println(" P", p.Key)
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()   
	var  row int
	 err1 := db.QueryRow("SELECT count(*) FROM public.activity where key = $1 ", p.Key).Scan(&row)
	   if err1 != nil {
	log.Fatal(err1)
}
	   fmt.Println("rows " ,row)
	   

	   if row == 0 {
   insertDynStmt := `insert into public."activity"("key", "activity") values($1, $2)`
    _, e := db.Exec(insertDynStmt, p.Key, p.Activity)
	
    CheckError(e)
	}else {
	fmt.Println("Not a unique activity")
	}
	
}
 
func CheckError(err error) {
    if err != nil {
       log.Fatal(err)
    }
}