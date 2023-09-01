package constants

type Service struct {
	SERVICE_NAME        string
	UPSTREAM_URL        string
	LOAD_BALANCING_TYPE string
}

// information about etcd instance
const (
	ETCD_ADDRESS = "0.0.0.0:20000"
)

// var name is [ServiceName]
var ViewerService = Service{
	SERVICE_NAME:        "ViewerService", //take straight from the idl file
	UPSTREAM_URL:        "0.0.0.0:8001",
	LOAD_BALANCING_TYPE: "ROUND_ROBIN",
}
