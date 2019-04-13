package main

import(
	"flag"
	"kidb/src/memdb"
	"kidb/src/server"
	"fmt"
)

func main(){
	var port = flag.String("Port", "8099", "Listening Port")
	db := memdb.NewDB()
	server := server.NewServer(db)
	fmt.Println("dbServer starting on localhost:"+*port)
	server.Start(*port)
}
