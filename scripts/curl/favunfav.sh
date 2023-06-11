#!/bin/bash

# Format
# ./scripts/curl/favunfav.sh [action] [asset_type] [asset_id]
# ./scripts/curl/favunfav.sh fav insight 146012c3-5fac-491f-8764-84741e223231
# ./scripts/curl/favunfav.sh unfav insight 146012c3-5fac-491f-8764-84741e223231

# Editable variables
HOSTNAME="localhost"
PORT="8080"
USER="efd8cec6-3e45-4fb1-b0d7-3a1be9cfae2c"

# Default values for action and asset ID
DEFAULT_ACTION="fav"
DEFAULT_ASSET_ID="15372e01-2ff5-4ec1-a2b8-9cb183933606"

# Command-line parameters
ACTION=${1:-$DEFAULT_ACTION}
ASSET_TYPE=${2:-"chart"}
ASSET_ID=${3:-$DEFAULT_ASSET_ID}

# API endpoint URL
URL="http://$HOSTNAME:$PORT/api/v1/$USER/assets/$ASSET_ID"

# JSON body
JSON_BODY=$(cat <<EOF
{
  "type": "$ASSET_TYPE",
  "action": "$ACTION",
  "name": "Asset Name",
  "description": "Asset Description"
}
EOF
)

# Echo the URL
echo "Request URL: $URL"
echo "JSON Body: $JSON_BODY"

# Make CURL request
curl -X PUT -H "Content-Type: application/json" -d "$JSON_BODY" "$URL"
