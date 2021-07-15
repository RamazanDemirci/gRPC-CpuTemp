package CpuClient

import (
	"context"
	"log"
	"time"

	"flag"

	pb "github.com/RamazanDemirci/gRPC-CpuTemp/grpc_Cpu/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
)

var (
	tls                = flag.Bool("_tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("_ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("_server_addr", "localhost:8882", "The server address in the format of host:port")
	serverHostOverride = flag.String("_server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

type Unit struct {
	Unit     string
	Current  float64
	High     float64
	Critical float64
}

func GetCpuTemp() ([]interface{}, error) { // cpu_temp grpc service client
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
	client := pb.NewCpuTempClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	units, err := client.GetCpuTemperature(ctx, &pb.Empty{})

	res := make([]interface{}, 0)
	if err != nil {
		return res, err
	}

	for _, temp := range units.Temp {
		res = append(res, Unit{
			Unit:     temp.Unit,
			Current:  temp.Current,
			High:     temp.High,
			Critical: temp.Critical,
		})
	}
	return res, err
}
