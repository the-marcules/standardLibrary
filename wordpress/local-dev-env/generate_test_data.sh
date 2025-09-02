#!/bin/bash

# Testdaten-Generator für WordPress Post Meta
# Verwendung: ./generate_test_data.sh [anzahl_posts]

POSTS_COUNT=${1:-10}  # Standard: 10 Posts wenn nicht angegeben
META_KEY="stock"
META_VALUE='s:28:"a:1:{s:7:\"default\";s:1:\"1\";}\""'

echo "Erstelle $POSTS_COUNT Test-Posts mit Stock-Meta..."

for i in $(seq 1 $POSTS_COUNT); do
    # Erstelle einen Test-Post
    POST_ID=$(wp post create --post_title="Test Produkt $i" --post_content="Test Inhalt für Produkt $i" --post_status=publish --post_type=product --porcelain)
    
    if [ $? -eq 0 ]; then
        # Setze Stock Meta für den erstellten Post
        wp post meta set $POST_ID $META_KEY "$META_VALUE"
        echo "✓ Post ID $POST_ID erstellt und Stock-Meta gesetzt"
    else
        echo "✗ Fehler beim Erstellen von Post $i"
    fi
done

echo ""
echo "Testdaten-Generierung abgeschlossen!"
echo "Erstellt: $POSTS_COUNT Posts mit Stock-Meta"
echo ""
echo "Zum Überprüfen verwenden Sie:"
echo "wp post meta list [POST_ID] --keys=$META_KEY"
