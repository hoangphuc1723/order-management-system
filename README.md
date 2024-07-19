# Order Management System with MongoDB and Mosquitto

This project sets up an Order Management System using Docker Compose. It includes a web application, a MongoDB database, and a Mosquitto MQTT broker.

## Prerequisites

Make sure you have the following installed on your system:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Mosquitto Configuration

Ensure you have a `mosquitto.conf` file inside the `mosquitto` directory. Here is an example configuration:

# mosquitto.conf

persistence true <br />
persistence_location /mosquitto/data/ <br />
log_dest file /mosquitto/log/mosquitto.log <br />

listener 1883 <br />
allow_anonymous true

## How to Run
Follow these steps to set up and run the project:

- Clone the repository:

git clone https://github.com/hoangphuc1723/order-management-system <br />
cd order-management-system

- Build and start the containers:

docker compose up -d <br />
This command will build the Docker image for the web application, if needed, and start all the services in the background.

## Access the web application:
The web application will be accessible at http://localhost:8080/web.

## Stopping the Services
To stop the running services, use the following command:  <br />
docker compose down <br />

## Logs and Data
MongoDB data is stored in a Docker volume named mongo-data.
Mosquitto data and logs are stored in the ./mosquitto/data and ./mosquitto/log directories, respectively.

## Troubleshooting
If you encounter issues, you can check the logs of the individual services:

docker-compose logs order-management-system-app <br />
docker-compose logs mongo <br />
docker-compose logs mosquitto

## Additional Notes
Ensure the mosquitto directory and its subdirectories (data and log) have the appropriate permissions.
You can customize the Mosquitto configuration by editing the mosquitto.conf file.
