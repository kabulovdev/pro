package kafka

import (
	"fmt"
	"gitlab.com/pro/post_service/config"
	pp "gitlab.com/pro/post_service/genproto/post_proto"

	//pp "pro/post_service/genproto/custumer_proto"
	"gitlab.com/pro/post_service/pkg/logger"
	"gitlab.com/pro/post_service/storage"
)

type KafkaHandler struct {
	config  config.Config
	storage storage.IStorage
	log     logger.Logger
}

func NewKafkaHandlerFunc(config config.Config, storage storage.IStorage, log logger.Logger) *KafkaHandler {
	return &KafkaHandler{
		storage: storage,
		config:  config,
		log:     log,
	}
}

func (h *KafkaHandler) Handle(value []byte) error {
	user := pp.PostForCreate{}
	err := user.Unmarshal(value)
	if err != nil {
		return err
	}
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Println(user)
	_,err = h.storage.Post().Create(&user)
	if err != nil {
		return err
	}
	return nil
}