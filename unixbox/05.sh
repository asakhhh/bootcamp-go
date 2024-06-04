#!/bin/bash
grep $(sort poem.txt | uniq -c | sort | head -1 | tail -c 4) poem.txt