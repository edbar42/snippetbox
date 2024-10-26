#!/bin/bash

ROOT_DIR="$(cd "$(dirname "$0")/../../" && pwd)"
ENV_FILE="$ROOT_DIR/.env"

if [ ! -f "$ENV_FILE" ]; then
    echo "Creating $ENV_FILE file..."

    cat <<EOL > "$ENV_FILE"
PORT=":42069"
DB_USER="user"
DB_PASS="passwd"
DB_NAME="snippetbox"
DSN=${DB_USER}:${DB_PASS}@/${DB_NAME}?parseTime=true
EOL

    echo "$ENV_FILE file created successfully."
else
    echo "$ENV_FILE file already exists."
fi

