package main

import ("log";"os";"encoding/csv")

func main(){
	fp, err := os.Open("Actions.csv")
	if err != nil {
		log.Printf("Could not open actions file: %v", err)
		return
	}
	defer fp.Close()
	csvReader := csv.NewReader(fp)	
	for record, err := csvReader.Read(); err == nil; record, err = csvReader.Read() {
		log.Print(record)
	}	
}
