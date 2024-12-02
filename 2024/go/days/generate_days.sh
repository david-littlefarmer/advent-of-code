#!/bin/bash

# Create folders
for ((i = 1; i <= 25; i++)); do
	folder_name="d$(printf "%02d" $i)"
	mkdir -p "days/$folder_name"

	# Copy the specific file (e.g., example.txt) to the created folder
	cp -nr days/d00/* "days/$folder_name/"
	find "days/$folder_name" -type f -exec sed -i "s/d00/$folder_name/g" {} +

done
