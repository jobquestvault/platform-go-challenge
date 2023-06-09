#!/bin/bash

# Editable variables
HOSTNAME="localhost"
PORT="8080"
USER="c03dc326-7160-4b63-ac36-7105a4c96fa3"

# Default values for action and asset ID
DEFAULT_ACTION="fav"
DEFAULT_ASSET_ID="32f92457-de96-4f0f-bfd1-9a382a198fd2"

# Command-line parameters
ACTION=${1:-$DEFAULT_ACTION}
ASSET_ID=${2:-$DEFAULT_ASSET_ID}

# API endpoint URL
URL="http://$HOSTNAME:$PORT/api/v1/$USER/assets/$ASSET_ID"

# JSON body
JSON_BODY=$(cat <<EOF
{
  "type": "chart",
  "action": "$ACTION",
  "name": "Asset Name"
}
EOF
)

# Make CURL request
curl -X PUT -H "Content-Type: application/json" -d "$JSON_BODY" "$URL"
