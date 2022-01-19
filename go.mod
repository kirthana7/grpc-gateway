module github.com/infobloxopen/grpc-gateway

go 1.17

replace github.com/grpc-ecosystem/grpc-gateway => ./

require (
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/spec v0.17.1
	github.com/go-resty/resty v0.4.2-0.20180405024425-f8815663de1e
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/rogpeppe/fastuuid v0.0.0-20150106093220-6724a57986af
	golang.org/x/net v0.0.0-20181005035420-146acd28ed58
	google.golang.org/genproto v0.0.0-20180427144745-86e600f69ee4
	google.golang.org/grpc v1.11.3
)

require (
	github.com/PuerkitoBio/purell v1.1.0 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/go-openapi/jsonpointer v0.17.1 // indirect
	github.com/go-openapi/jsonreference v0.17.1 // indirect
	github.com/go-openapi/swag v0.17.1 // indirect
	github.com/mailru/easyjson v0.0.0-20180823135443-60711f1a8329 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.0 // indirect
	gopkg.in/yaml.v2 v2.2.1 // indirect
)
