#!/bin/sh

if [ "$1" = "" ] 
then
    echo "Usage: $0 <source>"
    exit 1
fi

go run main.go "$1"
nasm -f elf64 test.asm
ld test.o -o test -lc -dynamic-linker /lib/ld-linux-x86-64.so.2
