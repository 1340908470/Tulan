---
description: 在本节，你将按照步骤为Tulan添加第一个功能：中文转英文；
---

# 为图蓝机器人添加功能

## 进入配置文件

Tulan未来预计提供多种配置方式，但核心是一样的，即通过编辑json文件定义机器人的功能。目前，暂无用户友好的可视化编辑器，开发者可以通过编辑 `~/def/def.json`对其进行定义。

![配置文件在项目文件夹中的位置](<../.gitbook/assets/image (3).png>)

## 为`processes`添加元素

### `name`

在本快速入门的例子中，我们以“将中文翻译为英文”功能为例，因此将`name`字段设置为“中文转英文”

### `trigger`

考虑一个典型场景，你希望Tulan帮你翻译一段文字，或者一个词语，于是你打开了与Tulan的对话框，那么如何让Tulan知道你的需求呢，没错，通过发送消息。Tulan的处理是以事务为单位的，而对如何触发事务的定义就是在`trigger`中完成的。

#### `keywords`

`keywords`用于设定触发该事务的关键词，其本身是一个有字符串组成的数组，在本例中，设置为与“中文转英文”相似的可能关键词，如："中文转英文", "中转英", "中文转英语", "中文翻译为英文", "中文翻译到英文"，当用户发送的消息中包含`keywords`中任意关键词时，Tulan即会触发对应事务。

![ 关键词数组](<../.gitbook/assets/image (1).png>)

#### `guide_index`

表示在用户确认进入事务后，Tulan进入的第一个指引（对用户来说，表现为Tulan向用户发送的消息卡片）。

#### `trigger_card`

`trigger_card`是当用户对机器人发送的消息中包含某个关键词触发了`trigger`后，在正式进入`process`前的确认卡片；`trigger_card`也属于[Broken link](broken-reference "mention")，而图蓝机器人的消息卡片是基于飞书机器人的，因此你可以通过[飞书消息卡片搭建工具](https://open.feishu.cn/tool/cardbuilder?from=howtoguide)快捷搭建消息卡片的基础部分。

![trigger\_card结构](<../.gitbook/assets/image (4).png>)

#### `trigger_cancel_card`

`trigger_cancel_card`是用户取消后显示的卡片，用户发送消息触发某一个事务后，Tulan机器人会向用户发送一个用于确认是否进入事务的卡片`trigger_card`，当用户点击按钮触发`trigger_action="no"`后，其就会变为`trigger_cancel_card`，如在本例子里，当用户点击取消按钮后，卡片标题就会变为：“取消图蓝事务”，内容就会变为：“您已结束图蓝事务：中文转英文”。

![trigger\_cancel\_card的例子](<../.gitbook/assets/image (6).png>)

#### `trigger_confirm_card`

`trigger_confirm_card`是用户确认后显示的卡片，用户发送消息触发某一个事务后，Tulan机器人会向用户发送一个用于确认是否进入事务的卡片`trigger_card`，当用户点击按钮触发`trigger_action="yes"`后，其就会变为`trigger_confirm_card`，如在本例子里，当用户点击确定按钮后，卡片标题就会变为：“开启图蓝事务”，内容就会变为：“您开启了图蓝事务：中文转英文”。

![trigger\_confirm\_card的例子](<../.gitbook/assets/image (5).png>)
