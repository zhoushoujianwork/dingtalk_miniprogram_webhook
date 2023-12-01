FROM centos:7

ENV TZ "Asia/Shanghai"

COPY ./miniprogram /usr/local/bin/miniprogram
COPY ./entrypoint.sh /root/entrypoint.sh
COPY config.yaml /root/config.yaml

WORKDIR /root
RUN chmod +x /usr/local/bin/miniprogram && \
    chmod +x /root/entrypoint.sh 

USER root

ENTRYPOINT /root/entrypoint.sh