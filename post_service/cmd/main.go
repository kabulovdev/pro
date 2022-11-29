package main

import (
	"net"
	"gitlab.com/pro/post_service/config"
	pb "gitlab.com/pro/post_service/genproto/post_proto"
	"gitlab.com/pro/post_service/pkg/db"
	"gitlab.com/pro/post_service/pkg/logger"
	"gitlab.com/pro/post_service/service"
	"gitlab.com/pro/post_service/kafka"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	CustomerCreateTopic := kafka.NewKafkaConsumer(connDB, &cfg, log, "customer.customer")
	go CustomerCreateTopic.Start()

	postService := service.NewPostService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
