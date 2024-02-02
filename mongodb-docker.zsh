#!/bin/zsh

# Name of the MongoDB container
CONTAINER_NAME="mongodb-teams-secure-notes"

# Name of the MongoDB database
DATABASE_NAME="teams-secure-notes-local"

# Generate random username and password
MONGO_USERNAME=$(openssl rand -base64 12)
MONGO_PASSWORD=$(openssl rand -base64 12)

# Create a Docker network for MongoDB if it doesn't exist
docker network create mongodb_network || true

# Start a new MongoDB container with port mapping
docker run --name $CONTAINER_NAME \
    -e MONGO_INITDB_ROOT_USERNAME=$MONGO_USERNAME \
    -e MONGO_INITDB_ROOT_PASSWORD=$MONGO_PASSWORD \
    -e MONGO_INITDB_DATABASE=$DATABASE_NAME \
    --network mongodb_network \
    -p 27017:27017 \
    -d mongo

# Print the credentials and correct MongoDB URI
echo "MongoDB container started."
echo "Database: $DATABASE_NAME"
echo "Username: $MONGO_USERNAME"
echo "Password: $MONGO_PASSWORD"
echo "To connect to your database, use the following URI:"
echo "mongodb://$MONGO_USERNAME:$MONGO_PASSWORD@localhost:27017/$DATABASE_NAME?authSource=admin"

