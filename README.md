## message-operator 简易型集群内消息中心
![](https://github.com/googs1025/message-operator/blob/main/image/%E6%B5%81%E7%A8%8B%E5%9B%BE%20(1).jpg?raw=true)
### 项目思路与设计
设计背景：本项目基于k8s的扩展功能，实现Message的自定义资源，实现一个集群内消息下发渠道的controller应用。调用方可在cluster中部署与启动相关配置即可使用。

思路：当应用启动后，会启动一个controller，controller会监听所需的资源，并执行相应的业务逻辑(如：消息发送)。

### 项目功能
1. 支持消息配置的热加载、热更新等功能。
2. 实现deployment、pod与service变更与删除后的消息通知。
3. 提供qq邮箱的发送功能、飞书机器人发送功能。

- 自定义资源如下所示
```yaml
apiVersion: api.practice.com/v1alpha1
kind: Message
metadata:
  name: mymessage
spec:
  sender:
    open: true
    remote: smtp.qq.com
    port:  25
    email: 2539512760@qq.com
    password: xxxxx
    targets: 3467320690@qq.com
  feishu:
    open: true
    webhook: xxxxxx  # 飞书 webhook ip
    type: card   # 支持 text card 推送模式
```