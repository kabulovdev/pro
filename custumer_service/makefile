CURRENT_DIR=$(shell pwd)

pull_submodule:
	git submodule update --init --recursive

update_submodule:
	git submodule update --remote --merge
	
build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

run_script:
	./script/gen-proto.sh

proto-gen:
	bash ${CURRENT_DIR}/script/gen-proto.sh
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
server:
	go run cmd/main.go

migrate_up:
	migrate -path migration/ -database postgres://abduazim:1234@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com/custumer_db up

migrate_down:
	migrate -path migration/ -database postgres://abduazim:1234@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com/custumer_db down

migrate_force:
	migrate -path migration/ -database postgres://abduazim:1234@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com/custumerdb force 2