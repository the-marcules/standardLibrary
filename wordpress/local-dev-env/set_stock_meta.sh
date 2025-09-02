#!/bin/bash

# Beispiel für das Setzen von Post Meta mit WP-CLI
# Annahme: Post ID = 123 (ersetzen Sie mit Ihrer tatsächlichen Post ID)

POST_ID=123
META_KEY="stock"
META_VALUE="s:28:\"a:1:{s:7:\"default\";s:1:\"1\";}\""

# WP-CLI Befehl zum Setzen der Post Meta
wp post meta set $POST_ID $META_KEY "$META_VALUE"

echo "Post Meta für Post ID $POST_ID mit Key $META_KEY gesetzt"

