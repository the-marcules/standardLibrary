#!/bin/sh
cd /app
npm install

until curl -s --head --request GET http://app:80/ready/ok.txt | grep "200"; do
    echo "Waiting for WordPress site to be available"
    sleep 2
done
echo "WordPress site is up!"
echo "current dir $(pwd)"
npx cypress run --e2e 