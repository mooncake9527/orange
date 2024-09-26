#!/bin/bash
project="open-platform"
chmod +x ./${project}
echo "kill ${project} service"
killall ${project}
nohup ./${project} start -c config.yaml >> access.log 2>&1 &
echo "run ${project} success"