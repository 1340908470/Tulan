---
description: 了解图蓝机器人配置文件的整体结构
---

# 整体配置

首先需要知道的是，图蓝机器人的能力是由若干个`process`组成的，每一个`process`可以让机器人对用户的一组操作进行响应，并通过若干交替的`guide`和`handle`，指引用户进行输入或点选操作，处理用户的输入，最终完成某一预设事务。首先，可以通过一个示例配置文件了解配置文件的整体结构：

```json
{
    "processes": [
        {
            // process 1
            "name": "中文转英文",
            "trigger": {
                "keywords": [
                    "中文转英文",
                    "中转英",
                    "中文转英语",
                    "中文翻译为英文",
                    "中文翻译到英文"
                ],
                "guide_index": 1,
                "trigger_card": {
                    // message card
                },
                "trigger_cancel_card": {
                    // message card
                },
                "trigger_confirm_card": {
                    // message card
                }
            },
            "guides": [
                {
                    "index": 1,
                    "guide_card": {
                        // message card
                    },
                    "success_handle_index": 1
                },
                {
                    "index": 2,
                    "guide_card": {
                        // message card
                    },
                    "success_handle_index": 1
                }
            ],
            "handles": [
                {
                    "index": 1,
                    "handler": "TranslateFromZhToEn",
                    "params": [
                        "guide_1_response"
                    ],
                    "success_guide_index": 2,
                    "failed_guide_index": 0,
                    "next_guide_index": 0
                }
            ]
        },
        {
            // process 2
        },
        {
            // process 3
        }
    ]
}
```

首先可以看到，配置文件的最外层是一个名为`processes`的对象数组，其数组内的每一个元素是一个`process`；每一个`process`由四个字段组成：名称 name、[chu-fa-qi-tigger.md](chu-fa-qi-tigger.md "mention")、[zhi-yin-mo-kuai-guide.md](zhi-yin-mo-kuai-guide.md "mention")和[chu-li-mo-kuai-handle.md](chu-li-mo-kuai-handle.md "mention")。

**名称 name** 没有什么好说的，就是`process`的名称，需要注意的是，`process`是以名称为主键的，如果你在数组的最后追加了一个与数组前面`process`同名的`process`，那么很有可能在你使用机器人的时候，机器人无法进入预期的事务。

**触发器 Tigger** 是`process`的入口，`keywords`字段指明了当用户的输入文本中包含哪些字符串时，会触发机器人的该`process`，触发`process`后，图蓝机器人会向用户发送消息卡片`trigger_card`进行确认，若用户确定进入该事务，则消息卡片会变为`trigger_confirm_card`，并进入索引为`guide_index`的首个指引模块；若用户拒绝进入该事务，则消息卡片会变为`trigger_cancel_card`。

**指引模块 guide** 用于对用户进行指引，该模块在用户侧具体表现为消息卡片，卡片内容由`guide_card`定义，引导用户输入一段文字、点击某个按钮、选择某个日期等，在指引过程中用户进行的各类输入，都会以键值对的形式保存在上下文中，供之后的其他`guide`或`handle`使用。在用户完成输入后，程序会调用索引为`success_handle_index`的`handle`，对数据进行一定的处理。

{% hint style="info" %}
**关于上下文：**程序为每一个处于事务中的用户维护一个上下文，上下文中保存了诸如当前用户在执行哪一个`process`、当前用户处于哪一个`guide`或哪一个`handle`，此前用户的输入以及handle的处理结果等。
{% endhint %}

**处理模块 handle** 用于对数据进行一次简单处理，或执行**一个**元操作，`handle`之后可以接`handle`，表示通过若干元操作的组合执行一个复杂操作，当前`handle`执行完进入哪一个`handle`是由`next_handle_index`字段决定的；`handle`之后也可以接`guide`，此处的`guide`应该当广义的指引来理解，其可以是提示用户进行下一步输入，也可以是handle的处理结果的显示。下一个`guide`的索引为`success_guide_index`或`failed_guide_index`。

{% hint style="info" %}
**值得注意：**`next_handle_index`的优先级高于`guide_index`，即当`next_handle_index`非零时，会接着执行下一个`handle`，只有当该字段为0时，才会进入指引。
{% endhint %}
