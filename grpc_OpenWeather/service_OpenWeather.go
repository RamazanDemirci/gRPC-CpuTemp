package main

import (
	"context"
	"encoding/json"
	"flag"
	owm "github.com/briandowns/openweathermap" // "owm" for easier use
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"fmt"
	"net"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"

	pb "github.com/RamazanDemirci/gRPC-CpuTemp/grpc_OpenWeather/protobuf"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 8881, "The server port")
)

// URL is a constant that contains where to find the IP locale info
const URL = "http://ip-api.com/json"

// template used for output
const weatherTemplate = `Current weather for {{.Name}}:
    Conditions: {{range .Weather}} {{.Description}} {{end}}
    Now:         {{.Main.Temp}} {{.Unit}}
    High:        {{.Main.TempMax}} {{.Unit}}
    Low:         {{.Main.TempMin}} {{.Unit}}
`

const forecastTemplate = `Weather Forecast for {{.City.Name}}:
{{range .List}}Date & Time: {{.DtTxt}}
Conditions:  {{range .Weather}}{{.Main}} {{.Description}}{{end}}
Temp:        {{.Main.Temp}} 
High:        {{.Main.TempMax}} 
Low:         {{.Main.TempMin}}
{{end}}
`

// Pointers to hold the contents of the flag args.
var (
	whereFlag = flag.String("w", "", "Location to get weather.  If location has a space, wrap the location in double quotes.")
	unitFlag  = flag.String("u", "", "Unit of measure to display temps in")
	langFlag  = flag.String("l", "", "Language to display temps in")
	whenFlag  = flag.String("t", "current", "current | forecast")
)

var LastLocation = ""
var LastUpdateTime = int64(0)
var LastTempValue = &pb.Temperature{Temp: 0, TempMin: 0, TempMax: 0, TimeStamp: 0}

// Data will hold the result of the query to get the IP
// address of the caller.
type Data struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"org"`
	AS          string  `json:"as"`
	Message     string  `json:"message"`
	Query       string  `json:"query"`
}

// getLocation will get the location details for where this
// application has been run from.
func getLocation() (*Data, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	r := &Data{}
	if err = json.Unmarshal(result, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type openWeatherServer struct {
	pb.UnimplementedOpenWeatherServer
}

// GetFeature returns the feature at the given point.
func (s *openWeatherServer) GetTemperatureByLocation(ctx context.Context, location *pb.Location) (*pb.Temperature, error) {
	// No feature was found, return an unnamed feature
	if LastLocation == location.Location && LastUpdateTime > time.Now().Unix()-int64(300000) {
		return LastTempValue, nil
	}

	LastUpdateTime = time.Now().Unix()

	w, err := getCurrent(location.Location, "C", "EN")

	LastTempValue = &pb.Temperature{Temp: w.Main.Temp, TempMin: w.Main.TempMin, TempMax: w.Main.Temp, Location: location.Location, TimeStamp: LastUpdateTime}

	return LastTempValue, err
}

// getCurrent gets the current weather for the provided
// location in the units provided.
func getCurrent(location, units, lang string) (*owm.CurrentWeatherData, error) {
	apiKey := os.Getenv("OWM_API_KEY")
	w, err := owm.NewCurrent(units, lang, apiKey)
	if err != nil {
		return nil, err
	}
	w.CurrentByName(location)
	return w, nil
}

func newServer() *openWeatherServer {
	s := &openWeatherServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOpenWeatherServer(grpcServer, newServer())
	log.Printf("server listening at %v", lis.Addr())
	grpcServer.Serve(lis)

	os.Exit(0)
}
