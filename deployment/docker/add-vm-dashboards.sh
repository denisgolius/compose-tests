#!/bin/sh
set -ex

mkdir -p /tmp/dashboards

cp ./dashboards/victoriametrics.json /tmp/dashboards/victoriametrics-vm-ds.json
cp ./dashboards/vmagent.json /tmp/dashboards/vmagent-vm-ds.json
cp ./dashboards/vmalert.json /tmp/dashboards/vmalert-vm-ds.json

for file in /tmp/dashboards/*.json; do sed -i 's/\"id\"\:\ \"prometheus\"/\"id\"\:\ \"victoriametrics-datasource\"/g' "$file"; done
for file in /tmp/dashboards/*.json; do sed -i  's/\"type\"\:\ \"prometheus\"/\"type\"\:\ \"victoriametrics-datasource\"/g' "$file"; done
for file in /tmp/dashboards/*.json; do sed -i  's/\"name\"\:\ \"Prometheus\"/\"name\"\:\ \"VictoriaMetrics\"/g' "$file"; done

cp -r /tmp/dashboards/*.json ./dashboards/