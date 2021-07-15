package OpenWeatherClient

import (
	"log"
	"time"

	"context"
	"flag"

	pb "github.com/RamazanDemirci/gRPC-CpuTemp/grpc_OpenWeather/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:8881", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

type Weather struct {
	Location    string
	Loc_Current float64
	Loc_Min     float64
	Loc_Max     float64
}

func GetTempByLocation(location string) (interface{}, error) { // open_weather grpc service client
	// Set up a connection to the server.
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOpenWeatherClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	weather, err := client.GetTemperatureByLocation(ctx, &pb.Location{Location: location})
	if err != nil {
		return nil, nil
	}
	res := Weather{
		Location:    weather.Location,
		Loc_Current: weather.Temp,
		Loc_Min:     weather.TempMin,
		Loc_Max:     weather.TempMax,
	}
	return res, err
	/*gTemp, err := client.GetTemperatureByLocation(ctx, &pb.Location{Location: location})
	if err != nil {
		log.Fatalf("%v.GetTemperatureByLocation(_) = _, %v: ", client, err)
	}
	log.Println(gTemp)
	json_data, _ := json.MarshalIndent(gTemp, "", "    ")
	fmt.Printf("temp_json : %s", json_data)*/
}
