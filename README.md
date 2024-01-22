# Google Search Service

## Requirements

Before getting started, make sure you have Docker and Docker Compose installed on your machine.

- Docker: [Docker Installation](https://docs.docker.com/get-docker/)
- Docker Compose: [Docker Compose Installation](https://docs.docker.com/compose/install/)

## Settings

1. Clone the repository to your local machine.
2. Create a `.env` file in the project's root directory with the necessary variables. An example can be found in the `.env.example` file.
3. To use the Google Custom Search API, you need to generate an API key and a Custom Search Engine ID (cx). Follow the steps at [this link](https://developers.google.com/custom-search/v1/introduction?hl=en) to obtain these credentials, and set them up in your project's .env file.

## Initialization

To start the application, follow these steps:

1. Open a terminal in the project's root directory.
2. Execute the following command to start the Docker containers:

```
docker-compose up -d
```

This will start the application, and you can access it at port 5443.

## Additional Notes

- Ensure that the necessary ports are available and not being used by other services.

To stop the application, use the following command:

```
docker-compose down
```
