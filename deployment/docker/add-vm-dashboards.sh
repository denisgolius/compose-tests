#!/bin/sh
set -ex

cp -r ./../../dashboards /tmp
FILES="/tmp/dashboards"
for file in $FILES; do sed -i 's/\"id\"\:\ \"prometheus\"/\"id\"\:\ \"victoriametrics-datasource\"/g' ${FILE}; done
for file in $FILES; do sed -i 's/\"type\"\:\ \"prometheus\"/\"type\"\:\ \"victoriametrics-datasource\"/g'${FILE}; done
for file in $FILES; do sed -i 's/\"name\"\:\ \"Prometheus\"/\"name\"\:\ \"VictoriaMetrics\"/g' ${FILE}; done