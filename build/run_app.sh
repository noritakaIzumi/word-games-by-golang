#!/usr/bin/env bash

SCRIPT_DIR=$(cd "$(dirname "$0")" || exit; pwd)
ROOT_DIR="$(dirname "${SCRIPT_DIR}")"
DIR_NAME="word_games"
APP_NAME=$1

if [[ -z $1 ]]; then
  echo "empty app name."
  exit 1
fi

if ! ls "$(dirname "${SCRIPT_DIR}")"/cmd/"${APP_NAME}"/main.go > /dev/null; then
  echo "${APP_NAME} does not exist."
  exit 1
fi

cp "${SCRIPT_DIR}"/Dockerfile "${ROOT_DIR}"

cd "${ROOT_DIR}" || exit

TAG_NAME="${DIR_NAME}-${APP_NAME}"

docker build -t "${TAG_NAME}" . \
  && docker run --rm -it "${TAG_NAME}" "${APP_NAME}" \

rm ./Dockerfile
