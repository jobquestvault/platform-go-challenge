#!/bin/bash

# Editable variables
HOSTNAME="localhost"
PORT="8080"
USER="c03dc326-7160-4b63-ac36-7105a4c96fa3"

# API endpoint URL
URL="http://$HOSTNAME:$PORT/api/v1/$USER/assets"

# Check if 'pretty' argument is provided
if [ "$1" == "pretty" ]; then
  # Make CURL request and pretty-print JSON response using jq
  curl "$URL" | jq
else
  # Make CURL request and display raw JSON response
  curl "$URL"
fi
