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
                    "failed_guide_index": 0
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

首先可以看到，配置文件的最外层是一个名为`processes`的对象数组，其数组内的每一个元素是一个`process`；每一个`process`由四个字段组成：名称name、[chu-fa-qi-tigger.md](chu-fa-qi-tigger.md "mention")、[zhi-yin-mo-kuai-guide.md](zhi-yin-mo-kuai-guide.md "mention")和[chu-li-mo-kuai-handle.md](chu-li-mo-kuai-handle.md "mention")
