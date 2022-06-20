# mapi 

A public api (BFF) serving as an aggregator and coordinator of requests.

## Getting started

### Project structure
```bash 
example
├── database            # db migrations
├── docs                # API docs swagger 
├── dto                 # Data transfer object 
├── cmd                 # Application commands.
├── pkg                 # 3rd party lib wrappers 
│   ├── configs         # Handle config via env vars and .env, yaml files.
│   ├── loggers         # Handle logging levels: Trace, Debug, Info, Warning, Error, Fatal and Panic..
│   ├── gormclient      # GORM MySQL Driver
├── internal            # Private application and library code
│   ├── http            # HTTP transport layer
│   ├──── handler       # API handlers & biz logic 
│   ├──── middleware    # Middleware handler pre processing of the request. 
│   ├──── server        # Http server with route 
│   ├── helper          # Helper function for application
│   ├── model           # Database model entity. Gorm models,
│   ├── usecase         # Handle business logic (optional),
│   └── repository      # CRUD repository implementation
└── ...
```

### Technical 
* Handler http server with Gin 
* Handle config with viper
* Handle commands with cobra 
* Handle logs with logrus 
* Handle mySQL with gorm 
* Goose: a database migration tool 

## ENV
```shell
cp .example.env .env  
```
## BUILD
```shell
go build -o service main.go
```

## RUN
```shell
./service serve 
```

