#!/bin/bash
set -e



unzip -o /var/www/html/wordpress/woody-products.zip -d /var/www/html/wordpress/wp-content/plugins/woody-products/
unzip -o /var/www/html/wordpress/woodyWoodyVodka.zip -d /var/www/html/wordpress/wp-content/themes/woodyWoodyVodka/



service apache2 restart

# Warte, bis die Datenbank erreichbar ist
until mysql -h database -u root -p"$MYSQL_ROOT_PASSWORD" -e "SELECT 1;" "$MYSQL_DATABASE" >/dev/null 2>&1; do
  echo "Warte auf die Datenbank..."
  sleep 3
done

wp core config --dbname="$MYSQL_DATABASE" --dbuser=root --dbpass="$MYSQL_ROOT_PASSWORD" --dbhost=database --allow-root --skip-check
wp core install --url="$WORDPRESS_URL" --title=wordpress --admin_user=admin --admin_password=admin --admin_email=marcules@gmail.com --allow-root

wp config set WP_DEBUG true --raw --type=constant --allow-root
wp config set WP_DEBUG_LOG true --raw --type=constant --allow-root
wp config set WP_DEBUG_DISPLAY true --raw --type=constant --allow-root

wp plugin install classic-editor --activate --allow-root
wp plugin uninstall akismet --allow-root
wp plugin uninstall hello --allow-root

wp theme activate woodyWoodyVodka --allow-root
wp plugin activate woody-products --allow-root || true

wp option update wopro_shop_toggle 1 --allow-root
wp option update wopro_paypal_client_id 888 --allow-root
wp option update wopro_paypal_client_secret 888 --allow-root
wp option update wopro_paypal_client_id_sandbox 888 --allow-root
wp option update wopro_paypal_client_secret_sandbox 888 --allow-root
wp option update wopro_paypal_sandbox_toggle 1 --allow-root
wp option update wopro_shipping_price 888 --allow-root

# Seite "Shop" erstellen und ID speichern
SHOP_ID=$(wp post create --post_type=page --post_title="Shop" --post_content="" --post_status=publish --porcelain --allow-root)
# Template setzen
wp post meta update $SHOP_ID _wp_page_template all-wopro_product.php --allow-root

# Seite "Deine Bestellung"
ORDER_ID=$(wp post create --post_type=page --post_title="Deine Bestellung" --post_content="" --post_status=publish --porcelain --allow-root)
wp post meta update $ORDER_ID _wp_page_template wopro-order-confirm.php --allow-root

# Seite "Warenkorb"
CART_ID=$(wp post create --post_type=page --post_title="Warenkorb" --post_content="" --post_status=publish --porcelain --allow-root)
wp post meta update $CART_ID _wp_page_template wopro-cart.php --allow-root


echo "Wordpress erfolgreich eingerichtet: 🔗 '$WORDPRESS_URL'"

exec bash || exec sh