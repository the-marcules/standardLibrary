# Verwende ein minimales Linux-Image
FROM php:8.3-apache 
#ubuntu/apache2:latest

RUN apt update

RUN apt-get update && apt-get install -y \
    libxml2-dev \
    libcurl4-openssl-dev \
    libgd-dev \
    libzip-dev \
    libicu-dev \
    libonig-dev \
    unzip \
    nano \
    wget \
    curl \
    default-mysql-client

RUN docker-php-ext-install mysqli bcmath

RUN mkdir -p /var/www/html/wordpress

WORKDIR /var/www/html/wordpress

RUN wget https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar
RUN chmod +x wp-cli.phar
RUN mv wp-cli.phar /usr/local/bin/wp
RUN wp --info
RUN php -d memory_limit=512M /usr/local/bin/wp core download --allow-root


COPY src/woody-products.zip /var/www/html/wordpress/woody-products.zip
COPY src/woodyWoodyVodka.zip /var/www/html/wordpress/woodyWoodyVodka.zip
COPY docker/webserver/entrypoint.sh /entrypoint.sh
COPY src/bulk_stock_import.php /var/www/html/wordpress/bulk_stock_import.php

RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]