#!/usr/bin/env bash

#backup log file
default_nohup_file_name=nohup.out
current_time=$(date "+%Y-%m-%d_%H.%M.%S")
nohup_backup_file_name=$default_nohup_file_name.$current_time".txt"
cp $default_nohup_file_name $nohup_backup_file_name

#backup last build file
service_name="./gitlab-webhook"
current_time=$(date "+%Y-%m-%d_%H.%M.%S")
backup_to=$service_name"."$current_time
mv $service_name $backup_to

#run services
go build
sh ./kill.sh
nohup ./gitlab-webhook > $default_nohup_file_name 2>&1 &
echo $! > pid.txt