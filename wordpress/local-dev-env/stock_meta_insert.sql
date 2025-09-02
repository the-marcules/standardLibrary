-- SQL Script für direktes Einfügen von Stock Meta
-- VORSICHT: Backup der Datenbank vor Ausführung empfohlen!

-- Beispiel für direktes INSERT in wp_postmeta
-- Ersetzen Sie 123, 124, 125 mit Ihren tatsächlichen Post-IDs

INSERT INTO wp_postmeta (post_id, meta_key, meta_value) VALUES
(123, 'stock', 's:28:"a:1:{s:7:"default";s:1:"1";}";'),
(124, 'stock', 's:28:"a:1:{s:7:"default";s:1:"1";}";'),
(125, 'stock', 's:28:"a:1:{s:7:"default";s:1:"1";}";');

-- Alternative: UPDATE falls Meta bereits existiert
-- UPDATE wp_postmeta 
-- SET meta_value = 's:28:"a:1:{s:7:"default";s:1:"1";}";" 
-- WHERE meta_key = 'stock' AND post_id IN (123, 124, 125);

-- Für WooCommerce Produkte (falls relevant):
-- UPDATE wp_postmeta 
-- SET meta_value = 's:28:"a:1:{s:7:"default";s:1:"1";}";" 
-- WHERE meta_key = '_stock' AND post_id IN (123, 124, 125);
