FROM cypress/included:latest

# Example: Install curl using apt-get (Debian/Ubuntu)
RUN apt-get update && apt-get install -y curl


COPY docker/testing/entrypoint.sh  /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN mkdir -p /app
# Set working directory
WORKDIR /app

# Copy all files from current directory
COPY docker/testing /app

# Install Cypress and dependencies
RUN npm install
RUN npx cypress install


ENTRYPOINT ["/entrypoint.sh"]