package main

import (
    "io/ioutil"
    "log"
)

func main() {
    if files, err := ioutil.ReadDir("/home/rif/Documents/prog/go/src/cmd"); err != nil {
        log.Print(err)            
    } else {
        for _, f := range files {ioutil.ReadAll(os.Open
            log.Print(f.Name())            
        }
    }
}
