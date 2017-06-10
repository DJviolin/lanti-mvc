#!/bin/bash

#echo "Hello, World!" >> out.txt

# GAWK tutorials
# https://www.tutorialspoint.com/awk/awk_workflow.htm
# https://www.tutorialspoint.com/awk/awk_workflow.htm

# USAGE
# $ LANG=C LC_ALL=C ./dcheck.sh

# set -e making the commands if they were like &&
# set -x putting + before every line
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TIMESTAMP=$(date +%s)

mkdir -p $DIR/dl

wget "http://www.pornhub.com/pornhub.com-db.zip" -O "$DIR/dl/pornhub.com-db-$TIMESTAMP.zip"
wget "http://www.pornhub.com/deleted.csv" -O "$DIR/dl/deleted-$TIMESTAMP.csv"
