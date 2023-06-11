#!/bin/bash

# Format
# ./scripts/curl/getfaved.sh [user_id] [page] [size] [pretty]
# ./scripts/curl/getfaved.sh efd8cec6-3e45-4fb1-b0d7-3a1be9cfae2c 1 12 pretty

# Editable variables
HOSTNAME="localhost"
PORT="8080"
USER="efd8cec6-3e45-4fb1-b0d7-3a1be9cfae2c

# Default values for page and size
DEFAULT_PAGE="1"
DEFAULT_SIZE="12"

# Get the values of page and size from command-line arguments
PAGE=${2:-$DEFAULT_PAGE}
SIZE=${3:-$DEFAULT_SIZE}

# API endpoint URL
URL="http://$HOSTNAME:$PORT/api/v1/$USER/faved?page=$PAGE&size=$SIZE"

# Echo the URL
echo "Request URL: $URL"

# Check if 'pretty' argument is provided
if [ "$1" == "pretty" ]; then
  # Make CURL request and pretty-print JSON response using jq
  curl "$URL" | jq
else
  # Make CURL request and display raw JSON response
  curl "$URL"
fi
