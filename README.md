# golangToPHP
通过golang的cgo特性，结合PHP的扩展开发的所需C库，编译链接成PHP扩展，提供给PHP脚本使用

## 使用golang开发PHP扩展
### 环境
- golang 1.18 （低版本没尝试，应该也可以）
- Linux
- PHP7.4 源码安装 （PHP8.x PHP5.X 没有尝试）

> 代码组成
   config.m4    
   function.go
   main.go

#####  config.m4文件：PHP脚手架ext_skel生成

## main.go

```go
package main

//#cgo CFLAGS: -g -I /home/php7/install/include/php -I /home/php7/install/include/php/main -I /home/php7/install/include/php/TSRM -I /home/php7/install/include/php/Zend  -I /home/php7/install/include/php/ext -I /home/php7/install/include/php/ext/date/lib -DHAVE_CONFIG_H
//#cgo LDFLAGS: -Wl,--export-dynamic -Wl,--unresolved-symbols=ignore-all
/*
#ifdef HAVE_CONFIG_H
#include "config.h"
#endif
#include "php.h"
#include "php_ini.h"
#include "ext/standard/info.h"
static int le_go2php;
PHP_MINIT_FUNCTION(go2php)
{
    return SUCCESS;
}
PHP_MSHUTDOWN_FUNCTION(go2php)
{
    return SUCCESS;
}
PHP_RINIT_FUNCTION(go2php)
{
    return SUCCESS;
}
PHP_RSHUTDOWN_FUNCTION(go2php)
{
    return SUCCESS;
}
PHP_MINFO_FUNCTION(go2php)
{
    php_info_print_table_start();
    php_info_print_table_header(2, "go2php support", "enabled");
    php_info_print_table_end();
}
PHP_FUNCTION(go2php_print)
{
    zend_long a,b;
    ZEND_PARSE_PARAMETERS_START(1, 1)
        Z_PARAM_LONG(a)
    ZEND_PARSE_PARAMETERS_END();
    b = calcFib(a);
    RETURN_LONG(b);
}
ZEND_BEGIN_ARG_INFO(null, 0)
ZEND_END_ARG_INFO()
const zend_function_entry go2php_functions[] = {
    ZEND_FE(go2php_print, null)
    PHP_FE_END
};
zend_module_entry go2php_module_entry = {
    STANDARD_MODULE_HEADER,
    "go2php",
    go2php_functions,
    PHP_MINIT(go2php),
    PHP_MSHUTDOWN(go2php),
    PHP_RINIT(go2php),
    PHP_RSHUTDOWN(go2php),
    PHP_MINFO(go2php),
    "0.1.0",
    STANDARD_MODULE_PROPERTIES
};
#ifdef COMPILE_DL_GO2PHP
ZEND_GET_MODULE(go2php)
#endif
*/
import "C"

func main() {}
```
> CFLAGS -g -I
g参数是支持gdb调试信息
I参数是指定导入头文件路径 （根据你PHP环境安装的位置，请替换{/home/php7/install/include/php}中的{/home/php7/install}路径。防止编译的时候C找不到头文件）

## function.go
这个是用golang语法写的C语言接口的函数，通过export关键字，导出为C的函数，注意参数和返回值要是C语言友好的类型，比如C.int *C.char等，相关内容可以去系统的学习下CGO的语法。

```go
package main
import  "C"
//export calcFib
func  calcFib(i int) int {
 if i < 2 {
 	return i
 }
 return  calcFib(i-1) + calcFib(i-2)
}
```
使用golang写内部功能和业务逻辑比用Zend宏命令和C来写PHP扩展来的简单一些，不需要去理解zend的过多宏命令，只需要了解PHP的扩展声明，方法声明，参数获取和类型转换。虽然CGO性能相比较纯golang和php的原生扩展，性能会有一些下降，但是开发效率和运行效率并不是很差，够用。

## 以上代码和环境准备好之后，进入编译链接阶段
- cd 到当前的源码目录
- 执行 {你的php安装目录}/bin/phpize （phpize是PHP环境的bin目录下的phpize工具）
- 执行完phpize命令之后，你会看到下面多出了很多文件
```shell
ls
autom4te.cache  build  config.h.in  config.m4  configure  configure.ac  function.go  go.mod  main.go  run-tests.php  test.php
```
- 执行 ./configure --with-php-config={你的php安装目录}/bin/php-config  PHP扩展安装的常规操作步骤之一
- 执行完./configure命令之后，你会看到下面又多出了很多文件，其中config.h是之前main.go 源码里需要的一个头文件，就在这一步生成的。
```shell
ls
autom4te.cache  config.h     config.log  config.nice    configure     function.go  include  main.go   Makefile.fragments  modules        test.php
build           config.h.in  config.m4   config.status  configure.ac  go.mod       libtool  Makefile  Makefile.objects    run-tests.php
```
- 最终阶段生成动态链接库SO
go build -gcflags "-l" -buildmode=c-shared -o go2php.so *.go
上面操作没有错误的话，将会生成so文件，这个就是PHP扩展文件了，然后将so扩展copy到php的扩展安装位置，然后去php.ini 增加extension=go2php
然后php -m 查看是否成功安装扩展。
- 测试
写一个php脚本
```php
<?php
echo go2php_print(10);
```
运行上面脚本看是否正常输出你期待的结果。

## 源码地址
https://github.com/529124368/golangToPHP
