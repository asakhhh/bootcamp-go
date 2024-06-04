#!/bin/bash

awk -F',' '{print $2}' <<< $(cat movies.csv | tr [:lower:] [:upper:])