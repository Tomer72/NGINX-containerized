Assignment
By: Tomer David Dahan

----- Summary -----

This project implements a containerized environment featuring two NGINX servers and an automated testing script. The NGINX servers are configured to respond on different ports with specific HTTP status codes and custom HTML. A Go-based testing script, also containerized, is orchestrated to run post-startup to validate the expected behavior of the servers. All components are grouped and managed via Docker Compose.

To fulfill the non-functional requirement of keeping the image sizes as small as possible, optimizations were made across both components:

Testing Stage: I conducted an analysis between Python script and Go script. By utilizing Go with a multi-stage build and a scratch base image, I reduced the testing image size to 11.9MB (compared to an 88.5MB Python equivalent).

NGINX Stage: there was a functional requirement to use the ubuntu base image, I implemented cleanups (removing apt-cache, unused lists, and manuals), reducing the final NGINX image to 150MB.


----- Project Architecture -----

1. NGINX Service (nginx folder)
Configuration (nginx.conf): Defines two distinct virtual servers:

    Server #1 (Port 8081): Responds with a custom HTML response and an HTTP 200 (OK) status code.

    Server #2 (Port 8082): Simulates a failure by responding with an HTTP 503 (Service Unavailable) status code.

    Dockerfile: Built on top of ubuntu:24.04. It installs NGINX, performs cleanups to minimize layer size, copies the configuration file, and exposes the required ports.

2. Testing Service (testing folder)
Testing Script (main.go): Sends HTTP GET requests to both NGINX servers. It asserts the responses against the expected status codes. It exits with Code 0 (Success) if all assertions pass, or Code 1 (Failure) if any unexpected behavior occurs.

Dockerfile: Utilizes a multi-stage build. The first stage compiles the Go binary, and the second stage uses the minimal scratch image to execute it.


----- Startup Instructions -----

- Ensure that ports 8081 and 8082 are available on your local machine.
- Use .env.example file to set the URL for each nginx server on the real .env file, using the ports mentioned above.

Clone the repository and navigate to the project's root directory.

Build and spin up the environment by running:

    docker compose up --build

To stop and remove the containers, run:

    docker compose down
