package main

import ("fmt"; "time"; "flag")

func main() {
	flag.Parse()
	var slice []error
	slice = append(slice, nil)
	slice = append(slice, nil)
	slice = append(slice, nil)
	fmt.Println(len(slice))
	d := time.Date(2012, time.February, 1, 14, 30, 1, 0, time.UTC)
	unixNano := d.UnixNano()
	d1 := time.Unix(0, unixNano)
	fmt.Println(d.Equal(d1))
	layout := "2006-01-02T15:04:05Z07:00"
	magic := d.Format(layout)	
	d1,_ = time.Parse(layout, magic)	
	//fmt.Println(d == d1)
	a:="mama are mere"
	fmt.Println(a[:4])
	duration,_ := time.ParseDuration(fmt.Sprintf("%ds", 10))
	fmt.Println(duration)
	fmt.Println(len(flag.Args()), flag.NArg())
}
