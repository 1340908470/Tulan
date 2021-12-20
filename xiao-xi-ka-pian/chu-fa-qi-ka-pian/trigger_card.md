# trigger\_card

`trigger_card`是当用户对机器人发送的消息中包含某个关键词触发了`trigger`后，在正式进入`process`前的确认卡片；下面是一个`trigger_card`的例子：

```json
{
    "config": {
        "wide_screen_mode": true
    },
    "elements": [
        {
            "tag": "div",
            "text": {
                "content": "您触发了图蓝事务：**@@process_name@@**，是吗？",
                "tag": "lark_md"
            }
        },
        {
            "tag": "action",
            "actions": [
                {
                    "tag": "button",
                    "text": {
                        "content": "😁  是的",
                        "tag": "plain_text"
                    },
                    "type": "default",
                    "value": {
                        "key": "trigger_action",
                        "value": "yes"
                    }
                },
                {
                    "tag": "button",
                    "text": {
                        "content": "😢  不是",
                        "tag": "plain_text"
                    },
                    "type": "default",
                    "value": {
                        "key": "trigger_action",
                        "value": "no"
                    }
                }
            ]
        }
    ],
    "header": {
        "template": "turquoise",
        "title": {
            "content": "🤖️ 触发图蓝事务",
            "tag": "plain_text"
        }
    }
    }
```

{% hint style="info" %}
在典型场景中，`trigger_card`应**至少包括两个按钮**，分别用于确定进入事务和取消进入事务，在`def.json`中，解释器会自动将`value.key = "trigger_action"`的 `button` 判定为`trigger_card`的确认按钮，当用户点击`value.value = "yes"`的按钮后，图蓝会将当前`trigger_card`卡片内容替换为`trigger_confirm_card`的内容，如指引用户输入处理当前事务所需的信息等，当用户点击`value.value = "no"`的按钮后，图蓝会将当前`trigger_card`卡片内容替换为`trigger_cancel_card`的内容，如显示当前事务已被取消等。
{% endhint %}

上方卡片在消息界面呈现的效果：

![](../../.gitbook/assets/image.png)

点击“是的”后，消息卡片会变换为`trigger_confirm_card`：

![](<../../.gitbook/assets/image (3) (1).png>)

点击“不是”后，消息卡片会变换为`trigger_cancel_card`：

![](<../../.gitbook/assets/image (1).png>)
