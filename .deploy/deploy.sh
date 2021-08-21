#!/bin/bash

WORKER_DIR=/docker/web-screenshot
DOCKER_IMAGE_NAME=debug.icu/web-screenshot-service:latest
DOCKER_RUN_CUSTOM_PARAM="-v /docker/web-screenshot:/opt/web-screenshot -p 8500:8500"
DOCKER_NODE_NAME=web-screenshot


cd $WORKER_DIR
if [[ "x$@" != "xno_build" ]];then
   echo "restart build image : $DOCKER_IMAGE_NAME"
   docker build -t $DOCKER_IMAGE_NAME .
fi


function ClearNode(){
   CLEAR_NAME=$1
   runing=`docker ps --filter name=$CLEAR_NAME -q`
   echo "check runing : "$runing
   if [[ "x"$runing != "x" ]]; then
       echo "runing pod $runing will stop and rm before restart!"
       docker stop $runing
       docker rm -f $runing
   fi
   stoping=`docker ps -a --filter name=$CLEAR_NAME -q`
   echo "check stoping :"$stoping
   if [[ "x"$stoping != "x" ]]; then
      echo "stoping pod $runing will stop and rm before restart!"
      docker rm -f $stoping
   fi
}

ClearNode $DOCKER_NODE_NAME

echo "about to be executed:"
echo "docker run -d $DOCKER_RUN_CUSTOM_PARAM --name $DOCKER_NODE_NAME  $IMAGE_NAME"
docker run -d $DOCKER_RUN_CUSTOM_PARAM --name $DOCKER_NODE_NAME  $IMAGE_NAME
if [[ $? -ne 0 ]];then
  echo "start docker node failed! exit status [$?]"
  exit $?
fi

echo "start docker node success!"
exit 0


