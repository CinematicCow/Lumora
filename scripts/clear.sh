#!/bin/bash

# Get the directory to delete
directory="$HOME/.lumora"

# Check if the user wants to bypass the prompt
if [[ "$1" == "-y" ]]; then
  # Delete the directory
  rm -rf "$directory"
  echo "Lumora has been wiped clean."
else
  # Prompt the user for confirmation
  read -p "Are you sure you want to delete the '$directory' directory? (y/N) " answer

  # Check the user's response
  if [[ "$answer" == "y" ]]; then
    # Delete the directory
    rm -rf "$directory"
    echo "Lumora has been wiped clean."
  else
    # Exit the script
    echo "Lumora is intact."
    exit 1
  fi
fi
