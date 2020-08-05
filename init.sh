#!/usr/bin/env bash

set -euo pipefail


cd "$(dirname "${BASH_SOURCE[0]}")"
CUR_DIR="$(pwd)"

function die() {
  killall backend  || true
  killall bee  || true
}
trap die INT

function prerequisites() {
  bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
  brew cask install virtualbox
  brew install node yarn golang docker docker-compose
  go get -u github.com/beego/bee
  echo "export GOPATH=\$GOPATH:${CUR_DIR}" >> ${HOME}/.bashrc
  export GOPATH=${GOPATH}:${CUR_DIR}
  xcode-select --install
}

function run() {
  export BEEGO_RUNMODE=dev
  docker-compose up -d db cache mail
  cd "${CUR_DIR}"/src/backend
  cp conf/app.dev.ini.secret.example conf/app.dev.ini.secret
  go mod vendor
  go build -o skeleton
  chmod +x skeleton
  ./skeleton migrate up
  bee run &
  go test -v ./... -count=1 | sort -u &


  cd "${CUR_DIR}"/src/frontend
  yarn install
  yarn serve &

  sleep 5
  open http://localhost:8080
  wait
}


if [ "${1-}" == "init" ];
then
    prerequisites
    exit
fi

run

