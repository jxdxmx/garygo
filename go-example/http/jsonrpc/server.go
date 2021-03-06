package main

import (
    "github.com/gorilla/mux"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
    "log"
    "net/http"
)

type Args struct {
    A, B int
}

type Arith int

type Result int

func(t *Arith) Multiply(r *http.Request, args *Args, result *Result) error{
    log.Printf("Multiply %d with %d\n",args.A,args.B)
    *result = Result(args.A * args.B)
    return nil
}

func main(){
    s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(),"application/json")
    s.RegisterCodec(json.NewCodec(),"application/json;charset=UTF-8")
    arith := new(Arith)
    s.RegisterService(arith,"")
    r := mux.NewRouter()
    r.Handle("/rpc",s)
    http.ListenAndServe(":1234",r)
}