# Notify Server

Notify 服务端。

## 启动参数

```shell
server
```

环境变量

```env
MiPushSecret # MiPush Token
FCMCredential # FCM Credential 文件内容
VAPIDPublicKey # Web Push VAPID Public Key
VAPIDPrivateKey # Web Push VAPID Private Key
```

数据库文件名称为 `notify.db`, `sqlite3` 格式。
启动后，服务器将会监听 `http://0.0.0.0:14444` 

在 `Docker` 中启动服务，如果想将数据库映射到宿主机，需要预先在宿主机上 `touch notify.db`, 并使用 `-v` 映射。

## 请求参数
```
POST http://host/{user_id}/send

@path   user_id  用户 ID
@param  title    推送标题
@param  content  推送内容
@param  long     传送到客户端的长内容, 需要点击查看

已移除
@query  dry      传递此参数时，不会真正发送推送 (请使用测试模式)
```

`long` 支持 markdown 格式， 支持使用表格扩展。

## 构建
构建前端后，将构建产物复制到 `/static` 文件夹，执行
```shell
go build -v github.com/ZNotify/server
```

## 前端
请查看 [Frontend](https://github.com/ZNotify/frontend)
