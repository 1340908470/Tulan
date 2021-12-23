# 处理模块 handle

**处理模块 handle** 用于对数据进行一次简单处理，或执行**一个**元操作，`handle`之后可以接`handle`，表示通过若干元操作的组合执行一个复杂操作，当前`handle`执行完进入哪一个`handle`是由`next_handle_index`字段决定的；`handle`之后也可以接`guide`，此处的`guide`应该当广义的指引来理解，其可以是提示用户进行下一步输入，也可以是handle的处理结果的显示。下一个`guide`的索引为`success_guide_index`或`failed_guide_index`。
