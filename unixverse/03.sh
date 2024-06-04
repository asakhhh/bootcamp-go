#!/bin/bash

awk -F',' '{print $5}' <<< $(tail -1 movies.csv)