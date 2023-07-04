#!/bin/sh
set -ex
for entry in $(ls ../../dashboards/)
do echo $entry; done
# for FILE in ./../../dashboards/; do sed -i 's/\"id\"\:\ \"prometheus\"/\"id\"\:\ \"victoriametrics-datasource\"/g' ${FILE} >; done
# for FILE in ./../../dashboards/; do sed -i 's/\"type\"\:\ \"prometheus\"/\"type\"\:\ \"victoriametrics-datasource\"/g'${FILE}; done
# for FILE in ./../../dashboards/; do sed -i 's/\"name\"\:\ \"Prometheus\"/\"name\"\:\ \"VictoriaMetrics\"/g' ${FILE}; done
# cp -r /var/lib/grafana/vm /var/lib/grafana/dashboards
# sed -i 's/type\:\ prometheus/type\:\ victoriametrics-datasource/g' /var/lib/grafana/vm/provisioning/datasources/datasources.yml
# cp -r /var/lib/grafana/vm/provisioning /var/lib/grafana/dashboards