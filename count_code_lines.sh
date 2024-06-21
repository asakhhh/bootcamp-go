#!/bin/bash

# Function to display usage message
usage() {
  echo "usage: $0 dir_path"
  exit 1
}

# Check if an argument is provided
if [ $# -eq 0 ]; then
  usage
fi

DIR=$1

# Check if the argument is a directory
if [ ! -d "$DIR" ]; then
  if [ -e "$DIR" ]; then
    echo "error: $DIR is not a directory"
  else
    echo "error: directory $DIR not found"
  fi
  exit 1
fi

# Directories to exclude from the count
EXCLUDE_DIRS="node_modules build dest .git"

# Find all files in the directory, excluding the specified directories
find "$DIR" -type d \( $(printf -- "-name %s -o " $EXCLUDE_DIRS | sed 's/ -o $//') \) -prune -o -type f -print |
while IFS= read -r file; do
  # Count non-empty lines in the file
  grep -v '^\s*$' "$file"
done | wc -l
