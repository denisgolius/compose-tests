#!/bin/sh
set -ex

mkdir -p /tmp/dashboards

cp ./dashboards/victoriametrics.json /tmp/dashboards/victoriametrics-vm-ds.json
cp ./dashboards/vmagent.json /tmp/dashboards/vmagent-vm-ds.json
cp ./dashboards/vmalert.json /tmp/dashboards/vmalert-vm-ds.json

for file in /tmp/dashboards/*.json; do sed -i 's/\"id\"\:\ \"prometheus\"/\"id\"\:\ \"victoriametrics-datasource\"/g' "$file"; done
for file in /tmp/dashboards/*.json; do sed -i  's/\"type\"\:\ \"prometheus\"/\"type\"\:\ \"victoriametrics-datasource\"/g' "$file"; done
for file in /tmp/dashboards/*.json; do sed -i  's/\"name\"\:\ \"Prometheus\"/\"name\"\:\ \"VictoriaMetrics\"/g' "$file"; done

for file in /tmp/dashboards/*.json; do sed -i  's/\"query\"\:\ \"prometheus\"/\"query\"\:\ \"victoriametrics-datasource\"/g' "$file"; done

for file in /tmp/dashboards/*.json; do sed -i  's/\"uid\"\:\ \"wNf0q_kZk\"/\"uid\"\:\ \"wNf0q_kZkVM\"/g' "$file"; done
for file in /tmp/dashboards/*.json; do sed -i  's/\"uid\"\:\ \"G7Z9GzMGz\"/\"uid\"\:\ \"G7Z9GzMGzVM\"/g' "$file"; done
for file in /tmp/dashboards/*.json; do sed -i  's/\"uid\"\:\ \"LzldHAVnz\"/\"uid\"\:\ \"LzldHAVnzVM\"/g' "$file"; done

for file in /tmp/dashboards/*.json; do sed -i  's/\"title\"\:\ \"VictoriaMetrics\"/\"title\"\:\ \"VictoriaMetrics - VM DS\"/g' "$file"; done
for file in /tmp/dashboards/*.json; do sed -i  's/\"title\"\:\ \"VictoriaMetrics\ -\ vmagent\"/\"title\"\:\ \"VictoriaMetrics\ -\ vmagent\ VM \DS\"/g' "$file"; done
for file in /tmp/dashboards/*.json; do sed -i  's/\"title\"\:\ \"VictoriaMetrics\ -\ vmalert\"/\"title\"\:\ \"VictoriaMetrics\ -\ vmalert\ VM \DS\"/g' "$file"; done

cp -r /tmp/dashboards/*.json ./dashboards/