# Twitter Ualá

Welcome to the Twitter Ualá!

This application is designed to work with a PostgreSQL database and is containerized using Docker. Follow the steps below to get the application up and running on your local machine.

## Prerequisites

- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: [Install Docker Compose](https://docs.docker.com/compose/install/)

## Running the Application
1. First of all you will need to create a `.env` file. Copying `.sample.env` and renaming it is the simplest way to getting started.

2. Now, from root project folder `twitter-uala` just run it with docker-compose

```bash
docker-compose up
```

## Migrations
Database migrations will be automatically applied on app´s startup.
It includes creation of some fake data for testing purposes.

