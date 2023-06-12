package main

import (
	"fmt"
	"net"

	"gitlab.com/pro/custumer_service/config"
	pb "gitlab.com/pro/custumer_service/genproto/custumer_proto"
	"gitlab.com/pro/custumer_service/kafka"
	"gitlab.com/pro/custumer_service/pkg/db"
	"gitlab.com/pro/custumer_service/pkg/logger"
	"gitlab.com/pro/custumer_service/pkg/messagebroker"
	"gitlab.com/pro/custumer_service/service"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {



	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "")
	
	conf := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 10,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}

	closer, err := conf.InitGlobalTracer(
		"user-service",
	)
	if err != nil {
		fmt.Println(err)
	}
	defer closer.Close()
	
	defer logger.Cleanup(log)

	log.Info("main:sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("datbase", cfg.PostgresDatabase))
	connDb, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	//kafka, connCloseFunc, err := producer.NewKafka(cfg)
	//if err != nil {
	//	log.Fatal("Error while connecting to kafka", logger.Error(err))
	//}
	//defer connCloseFunc()
	//kafka
	produceMap := make(map[string]messagebroker.Producer)
	topic := "customer.customer"
	customerTopicproduce := kafka.NewKafkaProducer(cfg, log, topic)
	defer func() {
		err := customerTopicproduce.Stop()
		if err != nil {
			log.Fatal("Failed to stopping Kafka", logger.Error(err))
		}
	}()
	produceMap["customer"] = customerTopicproduce

	//kafka end
	storeService := service.NewCustumService(connDb, log, produceMap)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
	fmt.Println("hi from custumer _service")
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCustomServiceServer(s, storeService)
	log.Info("main: server runing",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
