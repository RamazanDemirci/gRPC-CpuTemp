package main

import (
	"context"
	"flag"
	"math/rand"
	"time"

	CpuClient "github.com/RamazanDemirci/gRPC-CpuTemp/clients/Cpu"
	OpenWeatherClient "github.com/RamazanDemirci/gRPC-CpuTemp/clients/OpenWeather"
	influxdb2 "github.com/influxdata/influxdb-client-go"
)

var (
	Server = flag.String("influx_db_server", "http://localhost:8086", "Desc")
	Token  = flag.String("influx_db_token", "XEUPVtWjitcpkFmqNnLNnrQGFjNAycA2-dlvVb94AmqrmcLlooxWtjYVR2MTIlTT5AiCNjw0uxwA3tESzPTpuA==", "Desc")
	Org    = flag.String("influx_db_org", "memoirem", "Desc")
	Bucket = flag.String("influx_db_bucket", "northern", "Desc")
)

func job() {
	measurement := "stat"
	client := influxdb2.NewClient(*Server, *Token)
	defer client.Close()
	writeAPI := client.WriteAPIBlocking(*Org, *Bucket)
	units, _ := CpuClient.GetCpuTemp()
	if units == nil {
		return
	}
	weather, _ := OpenWeatherClient.GetTempByLocation("Samsun")
	if weather == nil {
		return
	}
	w := weather.(OpenWeatherClient.Weather)

	for _, e := range units {
		unit := e.(CpuClient.Unit)
		tag := map[string]string{"unit": unit.Unit}
		field_json := map[string]interface{}{
			"Current":  unit.Current,
			"High":     unit.High,
			"Critical": unit.Critical,
			"Location": w.Location,
			"Env_Temp": w.Loc_Current,
			"Env_Min":  w.Loc_Min,
			"Env_Max":  w.Loc_Max,
		}
		p := influxdb2.NewPoint(measurement, tag, field_json, time.Now())
		writeAPI.WritePoint(context.Background(), p)
	}
}

func main() {
	flag.Parse()

	for {

		go func() {
			job()
		}()

		// wait random number of milliseconds
		Nsecs := rand.Intn(5000)
		time.Sleep(time.Millisecond * time.Duration(Nsecs))
	}
}
