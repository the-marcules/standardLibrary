#!/bin/zsh

if [[ ! $1 || ! $3 ]]; then
    echo "Error: Missing Parameters"
    exit 1
fi

if [[ ! -e $3 ]]; then
    echo "Error: file does not exist."
    exit 1
fi

input_file="$3"


if [[ ! $2 ]]; then
    directory="$HOME/reports"
else
    directory="$2"
fi

month="$1"


mkdir -p "$directory"
if grep -- "$month" "$input_file" > "$directory/${month}_report.txt";
then
    echo "created a report for month $month to directory $directory"
else 
    echo "Did not find the needle in the hey stack"
fi
