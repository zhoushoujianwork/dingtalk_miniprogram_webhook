## 项目介绍
在使用钉钉企业开发事件回调时候发现需要部署回调接口暴露在公网，所以拆分出来一个简单的服务，用于转发钉钉回调事件到内网的服务。避免了内网服务需要暴露在公网的情况。

---
支持自定义 config.yaml 配置文件，来设定关注的事件，过滤的条件，最终执行的动作。
### 关注的事件：
    钉钉的事件类型，比如审批事件的 bpms_instance_change、bpms_task_change；
- 已经支持的事件：
    1. bpms_instance_change
### 过滤条件：
    钉钉事件的具体内容，比如审批事件的 processInstanceId、taskId；
### 执行动作：
    钉钉事件触发后，执行的动作，比如发送钉钉消息、发送邮件、执行脚本等；
    已经支持的动作：
    - http_post：发送 http post 请求

## 部署方式
### 1. 本地部署
```shell
# 直接运行 miniprogram
./miniprogram

# 或者使用 docker 运行
docker run -d -p 8080:8080 -v /path/to/config.yaml:/app/config.yaml --name miniprogram miniprogram

# 或者使用 docker-compose 运行
sh build.sh
# 配置docker-compose.yaml中的环境变量
docker-compose up -d
```