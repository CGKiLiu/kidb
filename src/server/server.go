package server

import (
	"kidb/src/memdb"
	"kidb/src/datatype"
	"kidb/src/rpcdatatype"
	"net/rpc"
	"net"
	"fmt"
	"net/http"
)


type Server struct{
	db *memdb.DB
}

func (s *Server) Put(args *rpcdatatype.Store, reply *rpcdatatype.NullResult) error{
	keySlice := datatype.NewSlice(args.Key)
	valueSlice := datatype.NewSlice(args.Value)
	s.db.Put(keySlice, valueSlice)
	*reply = 0
	return nil
}

func (s *Server) Get(args *rpcdatatype.Load, reply *rpcdatatype.ValueResult) error{
	keySlice := datatype.NewSlice(args.Key)
	valueSlice := s.db.Get(keySlice)

	if valueSlice==nil{
		reply = nil
	}else{
		reply = &rpcdatatype.ValueResult{
			Value : valueSlice.Data(),
		}
	}
	return nil
}

func NewServer(ndb *memdb.DB) *Server{
	server := &Server{
		db:ndb,
	}
	return server
}

func (s *Server) Start(port string){
	rpc.Register(s)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", port)
	if err != nil{
		fmt.Println("fatal")
	}
	http.Serve(l, nil)
}

/*
func main(){
	var port = flag.String("Port", "8080", "Listening Port")
	db := memdb.NewDB()
	server := NewServer(db)
	fmt.Println("dbServer starting on localhost:"+*port)
	server.Start(*port)
}
*/