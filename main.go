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
