#!/bin/bash

awk -F',' '{a=$2;$2=$4;$4=a; print $1, $2, $3, $4, $5}' movies.csv