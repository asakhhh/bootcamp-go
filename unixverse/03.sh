#!/bin/bash

f=$(tail -1 movies.csv)

awk -F',' '{print $5}' $f