#!/bin/bash

log_file_path="./tmp/chromedriver.log"
another_path="./tmp/another.log"
touch $log_file_path
chromedriver --port=9516 --adb-port=9517 --log-path=$log_file_path &
tail -f $log_file_path | tee -a $another_path
