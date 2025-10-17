#!/bin/bash
set -e



unzip -o /var/www/html/wordpress/woody-products.zip -d /var/www/html/wordpress/wp-content/plugins/woody-products/
unzip -o /var/www/html/wordpress/woodyWoodyVodka.zip -d /var/www/html/wordpress/wp-content/themes/woodyWoodyVodka/



service apache2 restart

# Warte, bis die Datenbank erreichbar ist
counter=0
until mysql -h database -u root -p"$MYSQL_ROOT_PASSWORD" -e "SELECT 1;" "$MYSQL_DATABASE" --skip-ssl >/dev/null 2>&1; do
  counter=$((counter + 1))
  echo "($counter) ⏳  Warte auf die Datenbank..."
  sleep 3
done

echo "✅ Datenbank ist erreichbar."
echo "🏗️  Richte WordPress ein..."

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
wp option update wopro_paypal_client_id_sandbox $PAYPAL_SANDBOX_CLIENT_ID --allow-root
wp option update wopro_paypal_client_secret_sandbox $PAYPAL_SANDBOX_CLIENT_SECRET --allow-root
wp option update wopro_paypal_sandbox_toggle 1 --allow-root
wp option update wopro_shipping_price 888 --allow-root


HEADER_MENU_NAME="Haupt"
HEADER_MENU_LOCATION="header-menu"

FOOTER_MENU_NAME="Footer"
FOOTER_MENU_LOCATION="footer-menu"

wp menu create "$HEADER_MENU_NAME" --allow-root
wp menu location assign "$HEADER_MENU_NAME" $HEADER_MENU_LOCATION --allow-root

wp menu create "$FOOTER_MENU_NAME" --allow-root
wp menu location assign "$FOOTER_MENU_NAME" $FOOTER_MENU_LOCATION --allow-root

# Seite "Shop" erstellen und ID speichern
SHOP_ID=$(wp post create --post_type=page --post_title="Shop" --post_content="" --post_status=publish --porcelain --allow-root)
# Template setzen
wp post meta update $SHOP_ID _wp_page_template all-wopro_product.php --allow-root
wp menu item add-post "$HEADER_MENU_NAME" $SHOP_ID --allow-root

# Seite "Deine Bestellung"
ORDER_ID=$(wp post create --post_type=page --post_title="Deine Bestellung" --post_content="" --post_status=publish --porcelain --allow-root)
wp post meta update $ORDER_ID _wp_page_template wopro-order-confirm.php --allow-root

# Seite "Warenkorb"
CART_ID=$(wp post create --post_type=page --post_title="Warenkorb" --post_content="" --post_status=publish --porcelain --allow-root)
wp post meta update $CART_ID _wp_page_template wopro-cart.php --allow-root

# Seite "Kontakt"
CART_ID=$(wp post create --post_type=page --post_title="Kontakt" --post_content="" --post_status=publish --porcelain --allow-root)
wp post meta update $CART_ID _wp_page_template contact-form.php --allow-root
wp menu item add-post "$HEADER_MENU_NAME" $CART_ID --allow-root

# Seite "Individuelle Fertigung"
INDIVIDUAL_ID=$(wp post create --post_type=page --post_title="Individuelle Fertigung" --post_content="" --post_status=publish --porcelain --allow-root)
wp menu item add-post "$HEADER_MENU_NAME" $INDIVIDUAL_ID --allow-root

# Seite "Individuelle Fertigung"
ABOUT_ID=$(wp post create --post_type=page --post_title="Über mich" --post_content="" --post_status=publish --porcelain --allow-root)
wp menu item add-post "$HEADER_MENU_NAME" $ABOUT_ID --allow-root


# Seite "AGB"
AGB_ID=$(wp post create --post_type=page --post_title="AGB" --post_content="" --post_status=publish --porcelain --allow-root)
wp menu item add-post "$FOOTER_MENU_NAME" $AGB_ID --allow-root

# Seite "Datenschutz"
DATENSCHUTZ_ID=$(wp post create --post_type=page --post_title="Datenschutz" --post_content="" --post_status=publish --porcelain --allow-root)
wp menu item add-post "$FOOTER_MENU_NAME" $DATENSCHUTZ_ID --allow-root

# Seite "Impressum"
IMPRESSUM_ID=$(wp post create --post_type=page --post_title="Impressum" --post_content="" --post_status=publish --porcelain --allow-root)
wp menu item add-post "$FOOTER_MENU_NAME" $IMPRESSUM_ID --allow-root

# Testprodukt anlegen
TEST_PRODUCT=$(wp post create --post_type=wopro_product --post_title="Testprodukt" --post_content="Dies ist ein Testprodukt." --post_status=publish --porcelain --allow-root)
wp post meta update $TEST_PRODUCT price 25.99 --allow-root

wp eval-file /var/www/html/wordpress/bulk_stock_import.php $TEST_PRODUCT --allow-root
# wp post meta update $TEST_PRODUCT stock 'a:1:{s:7:"default";s:1:"0";}"' --allow-root

mkdir -p /var/www/html/ready
touch /var/www/html/ready/ok.txt
echo "✅ Wordpress erfolgreich eingerichtet: 🔗 '$WORDPRESS_URL'"

exec bash || exec sh