package main

import (
	"net/rpc"
	"net/http"
	"math"
	"log"
	"flag"	
)

var (
	 = flag.String("server", "127.0.0.1:2000", "target host:port")
	listen = flag.String("listen", "127.0.0.1:1234", "target host:port")
)

type Sumer int

func (t *Sumer) Square(n float64, reply *float64) error {
	*reply = math.Sqrt(n)
	return nil
}

/*func stopSingnalHandler() {
	sig := <-signal.Incoming
	fmt.Printf("***Caught %s\n", sig)
	if usig,ok := sig.(os.UnixSignal); ok {
    	switch usig { 
    		case syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT:
    			fmt.Printf("***Caught signal %d, shutting down gracefully\n", sig)
    			os.Exit(1)
    	    default: 
    	    	os.Exit(1)
    	}    	
	}	 	
}*/

func registerToServer(){	
	client,err := rpc.DialHTTP("tcp", *server)
	if err != nil {
		log.Panic("Cannot register to server!")
	}
	var reply byte	
	log.Print("Registering to server ", *server)
	client.Call("RaterList.RegisterRater", *listen, &reply)	
	if err := client.Close(); err != nil {
		log.Panic("Could not close server registration!")
	}
}

func main() {
	flag.Parse()
	arith := new(Sumer)
	rpc.Register(arith)
	rpc.HandleHTTP()
	go registerToServer()
	//go stopSingnalHandler()
	http.ListenAndServe(*listen, nil)
}
