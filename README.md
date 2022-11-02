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
   --host HOST             Set host to HOST. (default: "127.0.0.1")
   --port PORT             Set port to PORT. (default: 14444)
   --address ADDRESS       Set listen address to ADDRESS.
   --test                  Enable test mode (default: false)
   --help, -h              show help (default: false)
```

## 配置文件

```yaml
users:
  - test
database:
  type: sqlite
  dsn: "file:./data/notify.db?cache=shared&mode=rwc&_journal_mode=WAL"
server:
  host: "127.0.0.1"
  mode: development
  port: 14444
senders:
  WebSocketHost:
  FCM:
    FCMCredential: ""
  WebPush:
    VAPIDPrivateKey: ""
    VAPIDPublicKey: ""
  TelegramHost:
    BotToken: ""
```

配置文件应当被放置在 `data/config.yaml` 中, 或者通过参数 `config` 指定。

可以使用 [JSON Schema](https://raw.githubusercontent.com/ZNotify/server/master/data/schema.json) 来验证配置文件。

## 请求参数
```
POST https://host/{user_id}/send

@path   user_id  用户 ID
@param  title    推送标题
@param  content  推送内容
@param  long     传送到客户端的长内容, 需要点击查看
```

`long` 支持 markdown 格式， 支持使用表格扩展。

> 完整的请求参数请参考 [API 文档](https://push.learningman.top/docs)

## 构建
构建[前端](https://github.com/ZNotify/frontend)后，将构建产物复制到 `/web/static` 文件夹，执行
```shell
go build -v
```
