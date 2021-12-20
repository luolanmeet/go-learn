## panic
### 作用
* 停止当前函数执行
* 一直往上返回，执行每一层的defer
** 如果没有遇见recover，程序退出
* panic应该少用

### recover
* 仅在defer调用中使用
* 可以获取panic的值
* 如果无法处理，可重新panic