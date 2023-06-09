#!/bin/bash

# Editable variables
HOSTNAME="localhost"
PORT="8080"
USER="c03dc326-7160-4b63-ac36-7105a4c96fa3"

# Default values for action and asset ID
DEFAULT_ACTION="update"
DEFAULT_ASSET_ID="9b9a27f1-5957-430e-8916-e1df3a79c13b"
DEFAULT_NAME="Default Name"

# Command-line parameters
ACTION=${1:-$DEFAULT_ACTION}
ASSET_ID=${2:-$DEFAULT_ASSET_ID}

# Generate a random 8-character alphanumeric name
generate_random_name() {
  cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 8 | head -n 1
}

# Use the provided name or generate a random name
NAME=${3:-$(generate_random_name)}
if [ -z "$NAME" ]; then
  NAME=$DEFAULT_NAME
fi

# API endpoint URL
URL="http://$HOSTNAME:$PORT/api/v1/$USER/faved/$ASSET_ID"

# JSON body
JSON_BODY=$(cat <<EOF
{
  "type": "chart",
  "action": "$ACTION",
  "name": "$NAME"
}
EOF
)

# Make CURL request
curl -X PUT -H "Content-Type: application/json" -d "$JSON_BODY" "$URL"
