#/bin/bash
mkdir -p /opt/logs/apps
# 如果没有debug变量则默认为false
if [ -z "$DEBUG" ]; then
    DEBUG=false
fi
miniprogram --debug=$DEBUG  --log_file /opt/logs/apps/app.log --config_file /root/config.yaml