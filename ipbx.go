package main

import (
	"log"
	//"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "api.cloudipbx.com:2810")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := map[string]string{"Name": "cgrates.com"}
	reply := map[string]string{}
	err = client.Call("Apier.GetDomain", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("Reply: %+v", reply)
}
