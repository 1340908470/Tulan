```json
{
    "process": [
        {
            "name": "",
            "trigger": {
                "keywords": [
                    "群发通知",
                    "群发"
                ],
                "guide_index": 1
            },
            "guide": [
                {
                    "index": 1,
                    "response_type": "message_box / text",
                    "response_para": [
                        "para_1",
                        "para_2"
                    ],
                    "success_handle_index": 1
                }
            ],
            "handle": [
                {
                    "index": 1,
                    "handler": "对于用户输入的处理方式",
                    "handler_paras": [
                        "value_name_1",
                        "value_name_2"
                    ],
                    "value_name": "处理后值的变量名",
                    "success_guide_index": 2,
                    "failed_guide_index": 0
                }
            ]
        }
    ]
}
```



举办活动 - 卡片 - 获得活动信息 - ~~创建活动~~ - 卡片 - 群发信息

预置参数：{{process_name}}

