#!/usr/bin/env bash

DIR_NAME="word_games"
APP_NAME=$1

if ! ls "${APP_NAME}" > /dev/null; then
    echo "${APP_NAME} does not exist."
    exit 1
fi

cp ./"${APP_NAME}"/Dockerfile ../

cd ../

TAG_NAME="${DIR_NAME}-${APP_NAME}"

docker build -t "${TAG_NAME}" . \
  && docker run --rm -it "${TAG_NAME}" \

rm ./Dockerfile
