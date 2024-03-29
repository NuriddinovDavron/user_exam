CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go


proto-gen:
	./scripts/genproto.sh

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs