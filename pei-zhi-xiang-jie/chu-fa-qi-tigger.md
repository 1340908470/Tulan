# 触发器 Tigger

**触发器 Tigger** 是`process`的入口，`keywords`字段指明了当用户的输入文本中包含哪些字符串时，会触发机器人的该`process`，触发`process`后，图蓝机器人会向用户发送消息卡片`trigger_card`进行确认，若用户确定进入该事务，则消息卡片会变为`trigger_confirm_card`，并进入索引为`guide_index`的首个指引模块；若用户拒绝进入该事务，则消息卡片会变为`trigger_cancel_card`。
