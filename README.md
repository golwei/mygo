# mygo
Go 关键字

break	（终止跳出循环）
case	（与switch 使用）
chan  （定义一个channel变量）
const	（定义常量）
continue	（跳过本次循环继续）
------------------
default		（默认语句）
defer	（延迟语句）
else		（选择语句）
fallthrough
for		（循环语句）
-------------------
func	(用于定义函数和方法）
go	（用于并行）
goto	（跳转语句）
if	（测试语句）
import	（导入外部包）
-------------------
interface	（定义接口）
map		（声明内部数据结构）
package	（定义包名）
range	（迭代）
return		（返回函数值）
----------------
select		（用于并发，监听channel上输入的数据）
struct		（定义结构）
switch		（开关语句）
type	（定义类型）
var	（声明变量）

==============================================================

Go 中的预定义函数

close	(关闭channel）
delete	（用于在map中删除实例）
len		（用于返回字符串，切片长度）
cap		（返回切片的最大容量）
new	（用于各种类型的内存分配）
make	（用于内建类型的内存分配）
append	（向slice追加零值或其他值，并返回追加后新的与slice同类型的值）
copy	（从源slice复制元素到目标，并返回复制元素的个数）
panic	（用于异常处理机制，恐慌）
recover		（用于异常处理机制，回复）
print		（底层打印函数，不用引入fmt包）
println		（底层打印函数，不用引入fmt包）
complex		（用于处理复数）
real		（用于处理复数）
imag		（用于处理复数）
