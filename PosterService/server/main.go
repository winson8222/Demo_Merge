package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"posterservice/constants"
	postersapi "posterservice/kitex_gen/postersapi/posterservice"
)

func main() {

	r, err := etcd.NewEtcdRegistry([]string{constants.ETCD_ADDRESS}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	ip, _ := net.ResolveTCPAddr("tcp", constants.PosterService.UPSTREAM_URL)

	svr := postersapi.NewServer(
		new(PosterServiceImpl),
		server.WithServiceAddr(ip),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.PosterService.SERVICE_NAME}),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
