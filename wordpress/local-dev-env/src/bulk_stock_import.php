<?php
/**
 * Bulk Stock Meta Import Script für WordPress
 * Verwendung: wp eval-file bulk_stock_import.php [POST_ID1] [POST_ID2] ...
 */

if (!empty($args)) {
    $post_ids = $args;
} else {
    echo "❌ Bitte geben Sie mindestens eine Post-ID als Argument an.\n";
    exit(1);
}
// Der serialisierte Stock-Wert
$stock_value = array("default" => "3");

foreach ($post_ids as $post_id) {
    // Prüfen ob Post existiert
    if (get_post($post_id)) {
        $result = update_post_meta($post_id, 'stock', $stock_value);
        
        if ($result) {
            echo "✅  Stock Meta für Post ID {$post_id} erfolgreich gesetzt\n";
        } else {
            echo "❌  Fehler beim Setzen der Stock Meta für Post ID {$post_id}\n";
        }
    } else {
        echo "⚠️  Post ID {$post_id} existiert nicht\n";
    }
}

echo "✅ Bulk Import abgeschlossen!\n\n";
?>
