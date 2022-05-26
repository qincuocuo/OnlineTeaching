BUILD_VERSION := 1.10.0
BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
BUILD_TIME := $(shell date "+%F %T")
BUILD_PATH := ..
BUILD_DATE := $(shell date "+%F %T")
BUILD_GIT_COMMIT := $(shell git rev-parse HEAD 2>/dev/null)
BUILD_LDFLAGS := -s -w
BUILD_LDFLAGS += -X 'ds_server/app.BuildDate=${BUILD_DATE}'
BUILD_LDFLAGS += -X 'ds_server/app.BuildGitCommit=${BUILD_GIT_COMMIT}'

webapi-srv:
	${BUILD_ENV} go build -mod=vendor -ldflags "${BUILD_LDFLAGS}" -o ${BUILD_PATH}/bin/webapi_srv ${BUILD_PATH}/cmd/webapi_srv.go
