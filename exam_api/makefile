pull_submodule:
	git submodule update --init --recursive

update_submodule:
	git submodule update --remote --merge
run:
	go run cmd/main.go

run_script:
	./scripts/gen-proto.sh

swag:
	swag init -g ./api/router.go -o api/docs

proto-gen:
	bash ${CURRENT_DIR}/scripts/gen-proto.sh
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"	