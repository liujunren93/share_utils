#!/bin/zsh
source ~/.zshrc
go env -w GOPRIVATE=github.com/liujunren93/share
for (( i = 0; i < 100; i++ )); do
     a=$( go get -u github.com/liujunren93/share)
#    echo 11
  if test ! -z $a ; then
    echo $a
    break
    else
            echo 11222
    fi

done
