#!/bin/bash
(cd docker && docker-compose down -v)

(rm src/*.zip)

(cd ../../../koelner-holzschmied/woody-products/src && zip -r woody-products.zip * && mv woody-products.zip ../../../standardLibrary/wordpress/local-dev-env/src)
(cd ../../../koelner-holzschmied/woodyWoodyVodka/src && zip -r woodyWoodyVodka.zip * && mv woodyWoodyVodka.zip ../../../standardLibrary/wordpress/local-dev-env/src)

# Start the WordPress development environment
cd docker && docker-compose up --build



# Open the WordPress site in the default web browser
# xdg-open http://localhost:8000
