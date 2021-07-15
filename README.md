# gRPC-CpuTemp

This Project has two grpc services respectively CpuTempService and OpenWeaterService. 

CpuTempService has one function which named as 'GetCpuTemperature' that triggered by client which implement grpc proto file. you can find client package clients/Cpu/client_Cpu.go

CpuOpenWheatherService has one function which named as 'GetTemperatureByLocation' that triggered by client which implement grpc proto file. you can find client package clients/Cpu/client_OpenWeather.go

# Folder Struct
 + clients
    + Cpu
        - client_Cpu.go
            - client implementation of Cpu grpc service
    + OpenWeather
        - client_OpenWeather.go
            - client implementation of OpenWeather grpc service
    + grpc_Cpu
        + protobuf
            - cpu_temp.proto
                - protobuf file which define the cpu grpc service
        - service_Cpu.go
            - Cpu service implementation and entry point
    + grpc_OpenWeather
        + protobuf
            - open_weather.proto
                - protobuf file which define the openWeather grpc service
        - service_OpenWeather.go
            - OpenWeather service implementation and entry point
    - main
        - InfluxDb Program Entry point


# grpc_Cpu : 

I installed linux machine lm-sensor service you can install following command

    $ sudo apt install lm-sensors

after install lm-sensor you can easy type following command displaying cpu temp information

    $ sensors

service type the sensors command. fetches output string. parses it. obtains the beneficial data. make this data to information and serves to clients. that's all.

# grpc_OpenWeather :
I use Openweather API to get the city tempreture information. OpenWeather API serve many other information such as humudity, wind, visibilty..etc. i interest only current, low and high tempreture information. This service wait any request from client and than it check its state if its temperature information is up-to-date it then serve this information and doesn't make OpenWeather api call. 

otherwise it request over api fetch the up-to-date information, update its state and serve this up-to-date information to client. why do i avoid to make api call each time. because i dont know what is the client periods to make request this service in addition could be many client at the same time, if i call many times to public and free Api the Openweather could block to my service which i don't desire such things.

# influxDB Service :
This actually main package of this repository. is responsible of fetch information from each services and format the information and write it down to influxdb.

why i use influxdb
- i have timeseries data, influxdb is timeseries nosql database.
- i want to display my data with elagant ui, thanksfully the influx data already have chronograf based displaying web interface which is easy to use and very useful.
- i work with hardware temprature time series information and influx already has telegraf service which collect all system information and save the influxdb.
- i can easily make chronograf dashboard with telefrag and my grpc services information to monitor the system.

all of above can develope and scale to large scale network of computer or IoT devices to monitor whole systems. 

i hope this code help somebody.

# OpenWeatherService generate gRPC codes
    $ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc_Cpu/protobuf/cpu_temp.proto

# CpuTempService generate gRPC codes
    $ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc_OpenWeather/protobuf/open_weather.proto

# run server
    $ go run ..service_dir/service_file_name.go

# run client
    $ go run ..client_dir/client_file_name.go