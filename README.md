# gRPC-CpuTemp

This Project has two grpc services respectively CpuTempService and OpenWeaterService. CpuTempService has one method which named as 'GetCpuTemperature' this method triggered by clint which imlement grpc proto file.

todo:...




# OpenWeatherService generate gRPC codes
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc_Cpu/protobuf/cpu_temp.proto

# CpuTempService generate gRPC codes
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc_OpenWeather/protobuf/open_weather.proto

# run server
$ go run greeter_server/main.go

# run client
$ go run greeter_client/main.go