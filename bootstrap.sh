#!/bin/bash

if [[ ${THS_TIER} == "test" ]];then
  echo "that env is test "
  rm -rf /opt/web-screenshot/conf/app.ini
	cp /opt/web-screenshot/conf/app_test.ini /opt/web-screenshot/conf/app.ini
fi

if [[ ${THS_TIER} == "prod" ]];then
  echo "that env is test "

  rm -rf /opt/web-screenshot/conf/app.ini
	cp /opt/web-screenshot/conf/app_prod.ini /opt/web-screenshot/conf/app.ini
fi

/opt/web-screenshot/web-screenshot-service