package main

import (	
	"log"
	"os/exec"	
)

const (
	TRANSFORMER = "avconv" 
)

func main(){
	path, err := exec.LookPath(TRANSFORMER)
	if err != nil {
		log.Fatal(TRANSFORMER + " not installed!")
	}
	log.Printf(TRANSFORMER + " is available at %s\n", path)	
	err = exec.Command(TRANSFORMER, "-i", "files/Agape_E_Raul_Sfant-D7yrvG4y5f8.aac", "-c:a", "libmp3lame", "-b:a", "96k", "-q:a", "9", "-y", "files/test.mp3").Run() 			
	if err != nil {
	    log.Print(err)
	}		
}
