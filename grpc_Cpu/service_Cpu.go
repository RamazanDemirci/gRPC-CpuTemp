package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"

	pb "github.com/RamazanDemirci/gRPC-CpuTemp/grpc_Cpu/protobuf"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 8882, "The server port")
)

type cpuTempServer struct {
	pb.UnimplementedCpuTempServer
}

// GetFeature returns the feature at the given point.
func (s *cpuTempServer) GetCpuTemperature(ctx context.Context, empty *pb.Empty) (*pb.Units, error) {
	cmd := exec.Command("sensors")
	out, err := cmd.CombinedOutput()
	lines := strings.Split(string(out), "\n")

	units := pb.Units{}

	for _, line := range lines {
		regex := *regexp.MustCompile(`Core (\d):.*\+([0-9]*[.])?[0-9]+°C  \(high = \+([0-9]*[.])?[0-9]+°C, crit = \+([0-9]*[.])?[0-9]+°C\)`)
		res := regex.FindAllStringSubmatch(line, -1)

		for i := range res {
			unit := res[i][1]
			current, err := strconv.ParseFloat(res[i][2], 64)
			if err != nil {
				return nil, err
			}
			high, err := strconv.ParseFloat(res[i][3], 64)
			if err != nil {
				return nil, err
			}
			critical, err := strconv.ParseFloat(res[i][4], 64)
			if err != nil {
				return nil, err
			}

			temp := pb.Temperature{
				Unit:     unit,
				Current:  current,
				High:     high,
				Critical: critical,
			}
			units.Temp = append(units.Temp, &temp)
		}
	}

	return &units, err
}

func newServer() *cpuTempServer {
	s := &cpuTempServer{}
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
	pb.RegisterCpuTempServer(grpcServer, newServer())
	log.Printf("server listening at %v", lis.Addr())
	grpcServer.Serve(lis)

	os.Exit(0)
}
