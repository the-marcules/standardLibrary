#!/bin/bash
(cd docker && podman-compose down -v)

(rm src/*.zip)

(cd ../../../koelner-holzschmied/woody-products/src && zip -r woody-products.zip * && mv woody-products.zip ../../../standardLibrary/wordpress/local-dev-env/src)
(cd ../../../koelner-holzschmied/woodyWoodyVodka/src && zip -r woodyWoodyVodka.zip * && mv woodyWoodyVodka.zip ../../../standardLibrary/wordpress/local-dev-env/src)

# Start the WordPress development environment
cd docker && podman-compose up --build -d

# Wait for wordpress to be ready
counter=0
until curl -s --head --request GET http://localhost:8080/ready/ok.txt | grep "200"; do
    counter=$((counter + 1))
    echo "[testing] ⏳  Warte auf die Einrichtung von WordPress... ($counter)"
    sleep 3
done

echo "[testing] ✅ WordPress ist online: http://localhost:8080/wordpress"

echo "[testing] 🚀 Starte die Tests..."
cd ../src/testing
npm install -y
npx cypress run



cmd='wait'
while [[ $cmd != 'exit' ]]; do

    read -p "[testing] ❓ Soll die Testumgebung heruntergefahren werden? (y/n) " yn
    case $yn in
        [Yy]* ) cmd='exit';;
        [Nn]* ) cmd='wait';;
        * ) echo "[testing] ⚠️  Bitte mit y oder n antworten.";;
    esac
    timeout=$((timeout -1))
done


# shut everything down
echo "[testing] 🗑️  Shutting down the WordPress development environment..."
cd ../../docker
podman-compose down -v

echo "[testing] ✅ Testumgebung wurde heruntergefahren. Bye! 👋"