# OpenKF Deploy

Using the OpenKF deploy tools, you can deploy OpenKF to docker clusters.

## Build images

### Build images for web

Build and run web image

```bash
# pwd = /deploy/docker
docker build -f Dockerfile.web -t openkf-web ../../.

# optional: run web image
docker run -d -p 8080:80 openkf-web
```

### Build images for server

```bash
# build openkf-server
docker build -f docker/Dockerfile.server -t openkf-server:latest ../.
```

You can check the images by running the following commands:
```bash
$ docker images

REPOSITORY                                            TAG                 IMAGE ID       CREATED              SIZE
openkf-server                                         latest              f1676becc5ce   About a minute ago   60.5MB
...
```

## Getting Started Guide: Docker Compose with Profiles

Welcome to our quick start guide designed to help you get up and running with our backend services using Docker Compose. By utilizing Docker Compose profiles, you can easily manage and run services tailored to different environments (development, testing, and production) from a single `docker-compose.yaml` file. This approach simplifies configuration management and enhances your workflow. Let's dive in!

### Prerequisites

- Docker and Docker Compose installed on your machine.
- Basic understanding of Docker concepts.
- Our `docker-compose.yaml` file downloaded to your project directory.

### Understanding Profiles

In our `docker-compose.yaml`, services are organized into profiles named `dev`, `test`, and `prod`. These profiles correspond to development, testing, and production environments, respectively. You'll select a profile to run the services relevant to your current needs.

### Starting Services with Profiles

1. **Development Environment**

   To start the services in a development environment, where you might need live code reloading and debugging tools, use the following command:

   ```bash
   docker-compose --profile dev up -d
   ```

   This command starts all services marked with the `dev` profile in detached mode, allowing you to continue using your terminal.

2. **Testing Environment**

   For running services in a testing environment, which could include test databases and mock services, execute:

   ```bash
   docker-compose --profile test up -d
   ```

   Services configured under the `test` profile will be initiated, facilitating your testing processes.

3. **Production Environment**

   To launch services in a production environment, optimized for performance and security, run:

   ```bash
   docker-compose --profile prod up -d
   ```

   This will start all services designated for production use, ensuring your application runs smoothly in a live setting.

### Checking Container Status

To view the status of running containers, use:

```bash
docker-compose ps
```

This command lists all containers initiated by Docker Compose, providing you with a quick overview of their statuses.

### Viewing Service Logs

To monitor the logs of a specific service for troubleshooting or performance analysis, execute:

```bash
docker-compose logs [service_name]
```

Replace `[service_name]` with the actual name of the service whose logs you wish to view.

### Stopping Services

When you're ready to stop all services, simply run:

```bash
docker-compose down
```

This command stops and removes all containers, networks, and volumes created by `up`.