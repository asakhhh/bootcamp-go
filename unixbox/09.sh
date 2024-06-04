#!/bin/bash

sz=$(cat poem.txt | wc -l)
num1=$(expr 1 + $RANDOM % $sz)
num2=$(expr 1 + $RANDOM % $sz)
echo $(head -n $num1 poem.txt | tail -1)
echo $(head -n $num2 poem.txt | tail -1)
