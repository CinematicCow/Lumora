#!/bin/bash

# Set the directory and file names.
dir_name="$HOME/.lumora"
file_name="lumora.gob"

# Check if the directory or file already exists.
if [[ -d "$dir_name" || -f "$dir_name/$file_name" ]]; then
  # Ask the user if they want to overwrite the directory or file.
  echo "The directory or file already exists. Do you want to overwrite it? (y/n)"
  read overwrite

  # If the user says yes, delete the directory or file and create a new one.
  if [[ "$overwrite" == "y" ]]; then
    rm -rf "$dir_name"
    mkdir -p "$dir_name"
    touch "$dir_name/$file_name"
    echo "Lumora has been initialized."
  else
    echo "Lumora could not be initialized."
  fi
else
  # Create the directory and file.
  mkdir -p "$dir_name"
  touch "$dir_name/$file_name"
  echo "Lumora has been initialized."
fi

# Exit with the OK code.
exit 0
