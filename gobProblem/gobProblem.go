package main

import ("encoding/gob"; "bytes"; "log")

type Test struct {
	A string
	B int
}

func main() {
	var stream bytes.Buffer
	gob.NewEncoder(&stream).Encode(Test{"one", 1})
	var result Test
	dec := gob.NewDecoder(&stream)
	err := dec.Decode(&result)
	log.Print(result, err)
	stream.Reset()
	var result2 Test
	gob.NewEncoder(&stream).Encode(Test{"two", 2})
	err = dec.Decode(&result2)
	log.Print(result2, err)
}
