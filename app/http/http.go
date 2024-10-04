package http

import (
	"context"
	"fmt"
	"net"

	"flag"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"base-api/config"

	"base-api/infra/redis"
	//-------------------------------------------------------------------------
)

var (
	grpcCMD = &cobra.Command{
		Use:   "serve-grpc",
		Short: "Run grpc server",
		Long:  "API",
		RunE:  runGRPC,
	}
)

func runGRPC(cmd *cobra.Command, args []string) error {
	// initial config
	ctx := context.Background()
	cfg := config.InitConfig()

	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	// db, err := db.Open(&cfg.DB)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// initial redis server
	redisServer := redis.NewRedisServer(&cfg.Redis)
	_, err := redisServer.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*time.Duration(cfg.Server.GraceFulTimeout), "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Maximum Received Message Bytes, will use later from env
	// maxReceivedMessageSize := grpc.MaxRecvMsgSize(1024 * 1024 * cfg.Server.MaxReceivedMessageSize)

	// RPC Server Option for later use, like interceptor or something that will be added to server option
	// serverOption := []grpc.ServerOption{
	// 	maxReceivedMessageSize,
	// }

	// init repo ctx
	// repoCtx := initRepoCtx(db, &cfg.S3)

	// init module ctx
	// moduleCtx := initModuleCtx(repoCtx, &cfg)

	// init service ctx
	// serviceCtx := initServiceCtx(repoCtx, moduleCtx, &cfg)
	// Open and Listen Server to configured Address
	listener, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		log.Panicf("Failed listen port %s => %v", cfg.Server.Addr, err)
		os.Exit(0)
	}
	fmt.Printf("RPC Listening on %s", cfg.Server.Addr)
	srv := grpc.NewServer()
	reflection.Register(srv)
	if err := srv.Serve(listener); err != nil {
		fmt.Println(err)
		log.Panicf("Failed SERVE gRPC => %v", err)
	}
	fmt.Printf("RPC Listening on %s", cfg.Server.Addr)
	return nil
}

// ServeHTTP return instance of serve HTTP command object
func ServeGRPC() *cobra.Command {
	return grpcCMD
}
