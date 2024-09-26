// grpc开启服务与grpc初始化
package start

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/mooncake9527/orange-core/common/utils/text"
	"github.com/mooncake9527/orange-core/core"
	"github.com/mooncake9527/orange-core/grpc/pb/health"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func grpcInit() {
	grpcServer = grpc.NewServer()
	//注册grpc服务
	health.RegisterHealthServer(grpcServer, &health.HealthServerImpl{}) //健康检测服务
	//service.RegisterGreeterServer(grpcServer, &impl.TempimplementedGreeterServer{})

	//注册服务完成
	grpcAddr := fmt.Sprintf("%s:%d", core.Cfg.GrpcServer.GetHost(), core.Cfg.GrpcServer.GetPort())
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		slog.Error("failed to listen", err)
		log.Fatal("failed to listen:", err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("failed to serve", err)
			log.Fatal("failed to serve:", err)
		}
	}()
	fmt.Println(text.Green("orange grpc server started ,Listen: ") + text.Red("[ "+grpcAddr+" ]"))
}

func closeGrpc() {
	grpcServer.GracefulStop()
}
