package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
    "strings"
)

func setup_db() (db *sql.DB) {
    //db, err := sql.Open("postgres", "user=ec2-user password=test dbname=twilio sslmode=disable")	
    db, err := sql.Open("postgres", "postgres://ec2-user:test@localhost:8887/twilio?sslmode=disable")
	panic_if(err)
	return	
}

func query_db(db *sql.DB, query string) (rows *sql.Rows){
	rows, err := db.Query(query)
	panic_if(err)
	return
}

func panic_if(err error) {
	if err != nil {
		panic(err)
	}
}

func print_rows(rows *sql.Rows) {
	panic_if(rows.Err())
	for rows.Next() {
		/*
		var id int
		var str string
		if err := rows.Scan(&id, &str); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d is '%s'\n", id, str)
		*/
		//var id, ts int
		//var to, from, account_sid, sms_sid, sms_status, message_sid, message_status, smsstatusfinal, messagestatusfinal string
		var count int
		//err := rows.Scan(&id, &ts, &to, &from, &account_sid, &sms_sid, &sms_status, &message_sid, &message_status, &smsstatusfinal, &messagestatusfinal)
		err := rows.Scan(&count)
		
		panic_if(err)
		fmt.Printf("%d\n", count); 
	}
	
	foo := "SmsSid"
	fmt.Println(strings.ToLower(foo))
	var post_data = map[string][]string{} //make(map[string][]string)
	fmt.Println(post_data)
	
}

func main() {
	db := setup_db()
	//rows := query_db(db, "SELECT * FROM sms.test")
	rows := query_db(db, "SELECT count(*) FROM sms.data where smssid = 'SMa38991df30f44f1bb50a6c2989553f39'")
	columns, _ := rows.Columns()
	fmt.Println(columns)
	print_rows(rows)
}
