#/usr/bin/env bash

clean() {
    docker stop $1 >/dev/null 2>&1 && \
        echo "Stopped container $1"
    docker rm $1 >/dev/null 2>&1 && \
        echo "Removed container $1"
    docker image rm $1 >/dev/null 2>&1 && \
        echo "Removed image $1@latest"
}

start() {
    docker build -t $1 . && \
    docker run -d -p 8080:8080 --name $1 $1 || \
    return 1
}

test "$2" = "-r" && \
    echo "Option error: use -r before app name" && \
    exit 1

APP_NAME=${2:-gohttp}

if test "$1" = "-r"; then
    clean $APP_NAME
    exit 0
fi

APP_NAME=${1:-gohttp}

if !(start $APP_NAME); then
    clean $APP_NAME
    start $APP_NAME
fi
