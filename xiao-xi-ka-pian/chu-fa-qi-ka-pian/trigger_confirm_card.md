# trigger\_confirm\_card

`trigger_confirm_card`是用户确认后显示的卡片，用户发送消息触发某一个事务后，Tulan机器人会向用户发送一个用于确认是否进入事务的卡片`trigger_card`，当用户点击按钮触发`trigger_action="yes"`后，其就会变为`trigger_confirm_card`，如在快速入门中的例子里，当用户点击确定按钮后，卡片标题就会变为：“开启图蓝事务”，内容就会变为：“您开启了图蓝事务：中文转英文”。

下面是一个`trigger_confirm_card`的例子：

```json
"trigger_confirm_card": {
    "config": {
        "wide_screen_mode": true
    },
    "elements": [
        {
            "tag": "div",
            "text": {
                "content": "您开启了图蓝事务：**@@process_name@@**",
                "tag": "lark_md"
            }
        }
    ],
    "header": {
        "template": "blue",
        "title": {
            "content": "🤖️  开启图蓝事务",
            "tag": "plain_text"
        }
    }
}
```

{% hint style="info" %}
本例子中，使用了存储于上下文中的参数：process\_name，表示当前所处的事务的名称，用@@包裹的内容表示引用上下文参数列表中的参数，其可以是如process\_name的Tulan默认提供的参数，也可以是用户输入或者每次handler处理后得到的结果。
{% endhint %}
