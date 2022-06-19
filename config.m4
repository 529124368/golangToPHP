

PHP_ARG_WITH(go2php, for go2php support,
  [  --with-go2php             Include go2php support]
)

if test "$PHP_GO2PHP" != "no"; then


  PHP_NEW_EXTENSION(go2php, go2php.c, $ext_shared)
fi
