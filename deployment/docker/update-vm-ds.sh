#!/bin/sh

VM_DS_PATH='deployment/docker/grafana-plugins/victoriametrics-datasource'
PLUGIN_PATH='deployment/docker/grafana-plugins'
# VM_DS_PATH='grafana-plugins/victoriametrics-datasource'
# PLUGIN_PATH='grafana-plugins'
#SIZE=$(du -sB 1 $VM_DS_PATH | cut -f 1)

if [[ ! -d ${VM_DS_PATH} ]] or [[ $(du -sB 1 $VM_DS_PATH | cut -f 1) -le 104857600 ]]; then
    echo "Victoriametrics datasource is not installed. Installing datasource..."
    mkdir -p ${VM_DS_PATH}
    wget -nc https://github.com/VictoriaMetrics/grafana-datasource/releases/download/v${VM_DS_VER}/victoriametrics-datasource-v${VM_DS_VER}.tar.gz -O ${VM_DS_PATH}-v${VM_DS_VER}.tar.gz
    tar -xzf ${VM_DS_PATH}-v${VM_DS_VER}.tar.gz -C ${PLUGIN_PATH}
    echo "Victoriametrics datasource has been installed."
else
    if [[ $(du -sB 1 $VM_DS_PATH | cut -f 1) -le 104857600 ]] or ; then
	    echo "Victoriametrics datasource is already installed.1"
    else
        if [ ! -f ${VM_DS_PATH}-v${VM_DS_VER}.tar.gz ]; then
            rm -rf ${VM_DS_PATH}/*
            wget -nc https://github.com/VictoriaMetrics/grafana-datasource/releases/download/v${VM_DS_VER}/victoriametrics-datasource-v${VM_DS_VER}.tar.gz -O ${VM_DS_PATH}-v${VM_DS_VER}.tar.gz
            tar -xzf ${VM_DS_PATH}-v${VM_DS_VER}.tar.gz -C ${PLUGIN_PATH}
	        echo "Victoriametrics datasource has been installed.2"
        else
            tar -xzf ${VM_DS_PATH}-v${VM_DS_VER}.tar.gz -C ${PLUGIN_PATH}
            echo "Victoriametrics datasource has been installed.3"
        fi
    fi
fi
