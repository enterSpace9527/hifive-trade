#!/bin/bash

service_name="hifive-trade"

# 检查服务是否存在
if pgrep -x "$service_name" > /dev/null
then
    # 如果服务存在，杀掉服务
    echo "Service $service_name is running. Killing the existing instance..."
    pkill -f "$service_name"
    sleep 2
fi

sleep 1
cur_dir="$PWD"
nohup ./hifive-trade -c=$cur_dir/etc/test_tradeapi.yaml > /dev/null 2>&1 &

if pgrep -x "$service_name" > /dev/null
then
    echo "start service $service_name success"
else
    echo "start service $service_name failed"
fi