package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"
	"time"
)

type CallDescriptor struct {
	TOR                                int
	CstmId, Subject, DestinationPrefix string
	TimeStart, TimeEnd                 time.Time
	Amount                             float64
	ActivationPeriods                  []*ActivationPeriod
}

type ActivationPeriod struct {
	ActivationTime time.Time
	Intervals      []*Interval
}

type Interval struct {
	Month                                  time.Month
	MonthDay                               int
	WeekDays                               []time.Weekday
	StartTime, EndTime                     string // ##:##:## format
	Ponder, ConnectFee, Price, BillingUnit float64
}

var (
	t1 = time.Date(2012, time.February, 8, 22, 50, 0, 0, time.UTC)
	t2 = time.Date(2012, time.February, 8, 22, 51, 50, 0, time.UTC)
	i  = &Interval{Month: time.February,
		MonthDay:  1,
		WeekDays:  []time.Weekday{time.Wednesday, time.Thursday},
		StartTime: "14:30:00",
		EndTime:   "15:00:00"}
	ap      = &ActivationPeriod{ActivationTime: t1, Intervals: []*Interval{i}}
	cd      = &CallDescriptor{CstmId: "vdf", Subject: "minutosu", DestinationPrefix: "0723", TimeStart: t1, TimeEnd: t2, ActivationPeriods: []*ActivationPeriod{ap}}
	network bytes.Buffer
	encGob  = gob.NewEncoder(&network)
	decGob  = gob.NewDecoder(&network)
	encJson = json.NewEncoder(&network)
	decJson = json.NewDecoder(&network)
)

func (ap *ActivationPeriod) storeGob() (result []byte) {
	network.Reset()
	encGob.Encode(ap)
	return network.Bytes()
}

func (ap *ActivationPeriod) restoreGob(input []byte) {
	network.Reset()
	network.Write(input)
	decGob.Decode(ap)
}

func (ap *ActivationPeriod) storeJson() (result []byte) {
	network.Reset()
	encJson.Encode(ap)
	return network.Bytes()
}

func (ap *ActivationPeriod) restoreJson(input []byte) {
	network.Reset()
	network.Write(input)
	decJson.Decode(ap)
}

func (ap *ActivationPeriod) storeGobNew() (result []byte) {
	network.Reset()
	gob.NewEncoder(&network).Encode(ap)
	return network.Bytes()
}

func (ap *ActivationPeriod) restoreGobNew(input []byte) {
	network.Reset()
	network.Write(input)
	gob.NewDecoder(&network).Decode(ap)
}

func (ap *ActivationPeriod) storeJsonNew() (result []byte) {
	network.Reset()
	json.NewEncoder(&network).Encode(ap)
	return network.Bytes()
}

func (ap *ActivationPeriod) restoreJsonNew(input []byte) {
	network.Reset()
	network.Write(input)
	json.NewDecoder(&network).Decode(ap)
}

/***************************** benchmarks *****************************/

func BenchmarkJson(b *testing.B) {
	cd1 := CallDescriptor{}
	for i := 0; i < b.N; i++ {
		encJson.Encode(cd)
		decJson.Decode(&cd1)
	}
}

func BenchmarkGob(b *testing.B) {
	cd1 := CallDescriptor{}
	for i := 0; i < b.N; i++ {
		encGob.Encode(cd)
		decGob.Decode(&cd1)
	}
}

func BenchmarkFuncJson(b *testing.B) {
	ap1 := ActivationPeriod{}
	for i := 0; i < b.N; i++ {
		encJson.Encode(ap)
		decJson.Decode(&ap1)
	}
}

func BenchmarkFuncGob(b *testing.B) {
	ap1 := ActivationPeriod{}
	for i := 0; i < b.N; i++ {
		encGob.Encode(ap)
		decGob.Decode(&ap1)
	}
}

func BenchmarkStoreRestoreJson(b *testing.B) {
	ap1 := ActivationPeriod{}
	for i := 0; i < b.N; i++ {
		result := ap.storeJson()
		ap1.restoreJson(result)
	}
}

func BenchmarkStoreRestoreGob(b *testing.B) {
	b.StopTimer()
	ap1 := ActivationPeriod{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		result := ap.storeGob()
		ap1.restoreGob(result)
	}
}

func BenchmarkRestoreOnlyJson(b *testing.B) {
	b.StopTimer()
	ap1 := ActivationPeriod{}
	result := ap.storeJson()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ap1.restoreJson(result)
	}
}

func BenchmarkRestoreOnlyGob(b *testing.B) {
	b.StopTimer()
	ap1 := ActivationPeriod{}
	result := ap.storeGob()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ap1.restoreGob(result)
	}
}

func BenchmarkStoreRestoreJsonNew(b *testing.B) {
	ap1 := ActivationPeriod{}
	for i := 0; i < b.N; i++ {
		result := ap.storeJsonNew()
		ap1.restoreJsonNew(result)
	}
}

func BenchmarkStoreRestoreJsonMarshall(b *testing.B) {
	ap1 := ActivationPeriod{}
	for i := 0; i < b.N; i++ {
		result,_ := json.Marshal(ap1)
		json.Unmarshal(result, &ap1)
	}
}

func BenchmarkStoreRestoreGobNew(b *testing.B) {
	ap1 := ActivationPeriod{}
	for i := 0; i < b.N; i++ {
		result := ap.storeGobNew()
		ap1.restoreGobNew(result)
	}
}

func BenchmarkRestoreOnlyJsonNew(b *testing.B) {
	b.StopTimer()
	ap1 := ActivationPeriod{}
	result := ap.storeJsonNew()	
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ap1.restoreJsonNew(result)
	}
}

func BenchmarkRestoreOnlyGobNew(b *testing.B) {
	b.StopTimer()
	ap1 := ActivationPeriod{}
	result := ap.storeGobNew()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ap1.restoreGobNew(result)
	}
}
