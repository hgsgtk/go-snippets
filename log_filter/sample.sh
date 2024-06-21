#!/bin/bash

log_file_path="./chromedriver.log"
another_path="./another.log"
touch $log_file_path
chromedriver --log-path=$log_file_path &
tail -f $log_file_path | tee -a $another_path
