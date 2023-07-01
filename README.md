# example service  

An example service. Init new service and contribute to make it better. Thanks :D 

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
│   ├── constants       # Static variables 
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

## Change something 
```shell
Go to go.mod file change **example-service** to **projectname-service**
Changing all .go files **"example-service/** to **"projectname-service/**
```

## ENV
```shell
cp .example.env .env  
```

## Pull all Dependencies 
```shell
export GOPRIVATE="gitlab.marathon.edu.vn" 
or 
export GOPROXY="https://proxy.golang.org,direct"
go mod tidy 
```

## BUILD
```shell
go build -o service main.go
```

## RUN
```shell
./service serve 
```
## Swagger 
* https://github.com/swaggo/gin-swagger
```shell
swag init
```

## Features 
```shell 
General function for get all
type Query struct {
	Q            string           //search string
	Select       []string         //select fields
	SearchFields []string         //search fields
	Filters      []*Filter        //filters
	Preloads     []string         //preloads 
	Joins        []*Join          //joins 
	Pagination   *Pagination      //pagination 
	Sort         *Sort            //sort 
	HaveCount    bool             //having count total results.
}
Building query for BFF, API Gateway get data in only one function instead of many functions like: 
- GetList
- GetListWithPagination
- GetListWithPaginationAndSort
- GetListWithPaginationAndSortAndFilter
- ...   
```

* Run your app, and browse to http://localhost:8080/swagger/index.html