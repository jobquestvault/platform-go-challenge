#!/bin/bash

# Format
# ./scripts/curl/updateasset.sh [asset_type] [asset_id] -n [name] -d [description]
# ./scripts/curl/updateasset.sh chart 22732116-c180-44ae-8da8-2c9644ed19f4 -n "Faved Chart" -d "Shows cool data"

# Editable variables
HOSTNAME="localhost"
PORT="8080"
USER="efd8cec6-3e45-4fb1-b0d7-3a1be9cfae2c"

# Default values for action, asset ID, name, and description
DEFAULT_ACTION="update"
DEFAULT_ASSET_ID="15372e01-2ff5-4ec1-a2b8-9cb183933606"
DEFAULT_NAME="Asset Name"
DEFAULT_DESCRIPTION="Asset Description"

# Default values for command-line parameters
ASSET_TYPE="insight"
ASSET_ID=$DEFAULT_ASSET_ID
NAME=$DEFAULT_NAME
DESCRIPTION=$DEFAULT_DESCRIPTION

# Command-line parameters
if [[ $# -ge 2 ]]; then
  ASSET_TYPE=$1
  ASSET_ID=$2
fi

# Flags
shift 2
while getopts ":n:d:" flag; do
  case $flag in
    n) NAME=$OPTARG;;
    d) DESCRIPTION=$OPTARG;;
    \?) echo "Invalid option -$OPTARG" >&2
        exit 1;;
  esac
done

# API endpoint URL
URL="http://$HOSTNAME:$PORT/api/v1/$USER/assets/$ASSET_ID"

# JSON body
JSON_BODY=$(cat <<EOF
{
  "type": "$ASSET_TYPE",
  "action": "$DEFAULT_ACTION",
  "name": "$NAME",
  "description": "$DESCRIPTION"
}
EOF
)

# Echo the URL
echo "Request URL: $URL"
echo "JSON Body: $JSON_BODY"

# Make CURL request
curl -X PUT -H "Content-Type: application/json" -d "$JSON_BODY" "$URL"
