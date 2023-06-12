#!/bin/bash

retry_count=0
max_retries=5
retry_delay=5

until psql -h 127.0.0.1 -U ak -d pgc -c "SELECT 1" > /dev/null 2>&1; do
    if [ $retry_count -eq $max_retries ]; then
        echo "Unable to connect to PostgreSQL container after $max_retries retries. Exiting..."
        exit 1
    fi

    echo "Waiting for PostgreSQL container to be ready..."
    sleep $retry_delay
    ((retry_count++))
done

echo "Setup completed. Starting the application..."
