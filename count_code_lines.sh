#!/bin/bash



# Check if the number of arguments is zero
if [ "$#" -eq 0 ]; then
  echo "usage: ./count_code_lines dir_path"
  exit 0
fi

# Check if an argument is provided

# Check if the argument is a directory
if [ ! -d "$1" ]; then
    if [ -e "$1" ]; then
        echo "error: $1 is not a directory"
    else
        echo "error: directory $1 not found"
    fi
    exit 0
fi

# Define directories to exclude
EXCLUDED_DIRS="node_modules|build|dest|.git"

# Find all files, excluding the specified directories, and count the lines of code
find "$1" -type d \( -path "*/node_modules" -o -path "*/build" -o -path "*/dest" -o -path "*/.git" \) -prune -false -o -type f -exec grep -v '^\s*$' {} + | wc -l