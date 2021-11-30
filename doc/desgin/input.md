```json
{
  "name": "中文转英文",
  "trigger": {
    "keywords": [
      "中文转英文", "中转英", "中文转英语", "中文翻译为英文", "中文翻译到英文"
    ],
    "guide_index": 1,
    "trigger_card": {},
    "trigger_cancel_card": {},
    "trigger_confirm_card": {}
  },
  "guides": [
    {
      "index": 1,
      "guide_card": {},
      "success_handle_index": 1
    },
    {
      "index": 2,
      "guide_card": {},
      "success_handle_index": 1
    }
  ],
  "handles": [
    {
      "index": 1,
      "handler": "TranslateFromZhToEn",
      "params": ["guide_1_response"],
      "success_guide_index": 2,
      "failed_guide_index": 0
    }
  ]
}
```



举办活动 - 卡片 - 获得活动信息 - ~~创建活动~~ - 卡片 - 群发信息

预置参数：{{process_name}}

