# Verwende ein minimales Linux-Image
FROM ubuntu/apache2:latest

RUN apt update

RUN apt install -y php8.3 php8.3-mysql php8.3-xml php8.3-curl php8.3-gd php8.3-mbstring php8.3-zip php8.3-intl php8.3-bcmath libapache2-mod-php8.3 wget unzip mysql-client curl nano


RUN mkdir -p /var/www/html/wordpress

WORKDIR /var/www/html/wordpress

RUN wget https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar
RUN chmod +x wp-cli.phar
RUN mv wp-cli.phar /usr/local/bin/wp
RUN wp --info
RUN wp core download --allow-root


COPY src/woody-products.zip /var/www/html/wordpress/woody-products.zip
COPY src/woodyWoodyVodka.zip /var/www/html/wordpress/woodyWoodyVodka.zip
COPY docker/webserver/entrypoint.sh /entrypoint.sh
COPY src/bulk_stock_import.php /var/www/html/wordpress/bulk_stock_import.php

RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]