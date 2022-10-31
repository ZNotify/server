# Notify Server

Notify 服务端。

## 启动参数

```shell
NAME:
   Notify API - This is Znotify api server.

USAGE:
   server [global options] command [command options] [arguments...]


COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE  Load configuration from FILE, or use ENV to load from environment variable CONFIG. (default: "data/config.yaml")
   --help, -h              show help (default: false)
```

## 配置文件

```yaml
server:
  port: 14444
  host: 127.0.0.1
database:
  type: sqlite
  dsn: data/notify.db
senders:
  WebSocketHost:
  FCM:
    FCMCredential: |
      {
        "type": "service_account",
        "project_id": "",
        "private_key_id": "",
        "private_key": "",
        "client_email": "",
        "client_id": "",
        "auth_uri": "",
        "token_uri": "",
        "auth_provider_x509_cert_url": "",
        "client_x509_cert_url": ""
      }
  WebPush:
    VAPIDPrivateKey: |
      PRIVATE_KEY
    VAPIDPublicKey: |
      PUBLIC_KEY
```

数据库文件名称为 `notify.db`, `sqlite3` 格式。
启动后，服务器将会监听 `http://0.0.0.0:14444` 

在 `Docker` 中启动服务，如果想将数据库映射到宿主机，需要预先在宿主机上 `touch notify.db`, 并使用 `-v` 映射。

## 请求参数
```
POST https://host/{user_id}/send

@path   user_id  用户 ID
@param  title    推送标题
@param  content  推送内容
@param  long     传送到客户端的长内容, 需要点击查看

已移除
@query  dry      传递此参数时，不会真正发送推送 (请使用测试模式)
```

`long` 支持 markdown 格式， 支持使用表格扩展。

> 完整的请求参数请参考 [API 文档](https://push.learningman.top/docs)

## 构建
构建前端后，将构建产物复制到 `/web/static` 文件夹，执行
```shell
go build -v
```

## 前端
请查看 [Frontend](https://github.com/ZNotify/frontend)
