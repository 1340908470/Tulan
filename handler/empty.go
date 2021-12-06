package handler

// EmptyHandler 空Handler，什么都不处理；
// 使用场景举例：在执行某一操作前，需要用户进行多次输入（如创建账号操作至少需要：输入账号+输入密码 两步），
// 而在Tulan的设计中，Guide之后跟的是Handler，因此可以通过使用空操作，接着执行Guide操作
func (h Handler) EmptyHandler() {}
