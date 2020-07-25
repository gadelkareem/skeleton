#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"
DIR="$(pwd)"
BUILD_DIR="$DIR/build"
FRONTEND_DIR="$DIR/src/frontend"
BACKEND_DIR="$DIR/src/backend"

rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR/conf/"

docker-compose up -d db cache mail

## generate frontend
cd "$FRONTEND_DIR"
yarn install
yarn test
export API_URL=/api/v1
yarn generate
mv src/dist "$BUILD_DIR/"

## build backend
cd "$BACKEND_DIR"
go mod tidy
go get ./...
# migration for tests
go build -o skeleton && chmod +x skeleton && ./skeleton migrate up
go test -v ./... -count=1
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o "$BUILD_DIR/skeleton"
go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o "skeleton"
chmod +x skeleton
mv skeleton "$BUILD_DIR/"
cp -r conf "$BUILD_DIR/"

cd "$BUILD_DIR"
export BEEGO_RUNMODE=prod
export REDIS_URL=redis://localhost:6379
export DATABASE_URL=postgres://skeleton_backend:dev_awTf9d2GceKRNzhkCb4H5B8nfmq@localhost:5432/skeleton_backend?sslmode=disable
./skeleton

