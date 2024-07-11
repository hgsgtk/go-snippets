#!/bin/sh
# Ensure the log directory exists and has the correct permissions
mkdir -p /var/log/selenium
chown -R 1000:1000 /var/log/selenium
chmod -R 775 /var/log/selenium

# Execute the passed command
exec "$@"
