Technical Guide
===

- **Project Title**: PlusVasis
- **Student 1:** Jason Henderson, 19309916
- **Student 2:** Conor Joyce, 19425804
- **Supervisor:** Dr Stephen Blott
- **Date Completed:** 05/05/23


# Table of Contents

[TOC]


# Introduction

## Overview/Motivation

PlusVasis is a container management and orchestration platform that provides developers with a modern and intuitive interface to manage their containers. The platform consists of a Svelte-based frontend and a backend REST API built using the Golang Echo framework. PlusVasis communicates with a personal Nomad server to create and manage container jobs, making container management a breeze for developers.

The platform features seamless networking, security, and scalability, allowing developers to focus on building their applications instead of worrying about complex backend processes. PlusVasis is built using modern frameworks and infrastructure solutions, ensuring optimal performance and reliability.

Developers can easily manage their containers using PlusVasis's intuitive interface, which streamlines container management and orchestration. PlusVasis enables developers to focus on building their applications, while the platform takes care of the container management and orchestration.

The name "PlusVasis" is derived from the Latin words "plus" meaning "more" and "vasis" meaning "containers". This name reflects the platform's goal of simplifying container management and enabling developers to manage more containers with ease.

The idea for this project came from real problems encountered by Conor, who was a Systems Administrator in DCU's NetSoc - Redbrick. It was often requested by users of Redbrick's infrastructure if they could host their websites, game servers, personal projects etc. on Redbrick, but this required manual effort from the admins. It would have been great to allow users to deploy whatever they like, whenever they like, all by themselves.

Both of us felt motivated by this idea and believed we could create something easy to use that met this goal sufficiently. We decided to use Nomad as our job orchestrator/runner because of its ability to scale horizontally, which is very important for a project where users can potentially use a lot of resources such as memory and CPU time.


## Glossary

**SvelteKit**: A web application framework based on the Svelte JavaScript framework, designed to simplify the process of building web applications.

**Vite**: A build tool and development server that enables fast and efficient web development, often used in conjunction with SvelteKit.

**TypeScript**: A programming language that adds static typing to JavaScript, providing enhanced tooling, code organization, and error detection.

**Backend**: The server-side component of your application, responsible for processing requests, handling data, and communicating with other services.

**REST API**: An architectural style for designing networked applications, where the API uses standard HTTP methods (such as GET, POST, PUT, and DELETE) to perform actions on resources.

**API middleware**: Software components that sit between an application and an API, providing additional functionality or processing, such as authentication, caching, or logging.

**Go (Golang)**: A programming language developed by Google that is known for its simplicity, efficiency, and strong support for concurrent programming.

**Echo framework**: A lightweight web framework for Go that simplifies the development of RESTful APIs by providing routing, middleware, and request-handling features.

**Frontend**: The client-side component of your application, responsible for rendering the user interface and handling user interactions.

**Testing**: The process of verifying the correctness and functionality of your application through various tests, such as unit tests, integration tests, and end-to-end tests.

**Vitest**: A testing framework or library used for testing SvelteKit applications, providing capabilities for testing components and simulating user interactions.

**Error handling**: The process of detecting, managing, and recovering from errors or exceptions that occur during the execution of your application.

**HTTP status codes**: Standard codes used to indicate the status of HTTP requests and responses, providing information about the success or failure of a request.

**Cypress**: A JavaScript end-to-end testing framework that allows you to write and execute tests in a browser-like environment, simulating user interactions and verifying the behaviour of your application.

**Docker**: A platform for developing, shipping, and running applications in containers, which are lightweight, portable, and self-sufficient. Docker allows you to package an application with all its dependencies and configurations, ensuring consistency across environments.

**Docker images**: Read-only templates that contain a set of instructions for creating a container that can run on the Docker platform. Images include application code, libraries, tools, dependencies, and other files needed to make an application run. They consist of multiple layers, each one originating from the previous layer but being different from it. Images can be stored in private or public repositories, such as Docker Hub.

**Docker containers**: Instances of Docker images, which are lightweight, portable, and self-sufficient. Containers represent the execution of a single application, process, or service and consist of the contents of a Docker image, an execution environment, and a standard set of instructions.

**Docker volumes**: Writable filesystems that containers can use, allowing programs to access a writable filesystem since images are read-only. Volumes live in the host system and are managed by Docker.

**docker-compose**: A command-line tool and YAML file format for defining and running multi-container applications. Developers can define a single application based on multiple images with one or more .yml files and deploy the whole multi-container application with a single command (`docker-compose up`).

**Nomad**: A flexible workload orchestrator that enables an organization to automate the deployment of any application on any infrastructure at any scale. 

**Nomad jobs**: Configuration files or definitions that describe how containers should be run and managed within the Nomad cluster, specifying parameters, resources, dependencies, and other details.

**Nomad allocations**: Individual instances of tasks that are scheduled by Nomad and run within a Nomad client. Allocations are the smallest unit of work in a Nomad cluster.

**Traefik**: A modern HTTP reverse proxy and load balancer that makes deploying microservices easy by dynamically managing routing, load balancing, and SSL/TLS termination.

**SSL certificates**: Digital certificates that provide secure and encrypted communication between a client and a server, ensuring the authenticity of the server and the integrity of the transmitted data.

**Cloudflare**: A global cloud network platform that offers various services, including DNS resolution, DDoS protection, and web security.

**A records**: DNS records that map a domain name to an IPv4 address, allowing users to access a website or service using a domain name instead of an IP address.

**CNAME records**: DNS records that map one domain name to another, effectively creating an alias for the target domain name.

**Webhooks**: User-defined HTTP callbacks that allow real-time notifications and updates when an event occurs in a web application or service.


## Research

Developing PlusVasis required extensive research to identify the right tools and technologies to use. We began by evaluating different backend frameworks and ultimately chose Echo for its ease of use, flexibility in terms of middleware support, and the new learning opportunities it provided. Further research was required when developing our frontend framework, particularly concerning server-side rendering in Svelte. Throughout the development process, we encountered challenges such as difficulty finding relevant resources and unexpected roadblocks in implementing certain features.

To ensure the quality of the platform, we also invested time in researching the best testing practices. This included writing unit and integration tests for our REST API, component testing for our frontend, and implementing a series of end-to-end system tests using Cypress. One of the final research efforts was integrating Docker Compose and learning how to create nomad jobs for containers by using Docker Compose.

Overall, the research we conducted enabled us to make informed decisions about technology choices and improve the quality of the platform. For example, using Echo was very powerful, yet still provided a good developer experience, while testing best practices helped ensure a high level of reliability and robustness for PlusVasis.


# High-Level Design

## User Login/Register


<img src="https://hedgedoc.plusvasis.xyz/uploads/6abeb466-8a70-40ba-975d-b3846d1cd821.png" width=400>

## Container Creation


<img src="https://hedgedoc.plusvasis.xyz/uploads/c4e76e0b-4f1a-4919-85c7-85dfda58ea34.png" width=400>

## Container Update


<img src="https://hedgedoc.plusvasis.xyz/uploads/e56c09d7-a166-4fd8-b1b5-709b78caf1e2.png" width=400>

## Container Stop


<img src="https://hedgedoc.plusvasis.xyz/uploads/474c4397-87ab-4326-b398-205195ff2e34.png" width=400>

## Container Start


<img src="https://hedgedoc.plusvasis.xyz/uploads/e5eecd6c-145a-40e8-a109-7374b9e2bc3f.png" width=400>

## Container Restart


<img src="https://hedgedoc.plusvasis.xyz/uploads/9586eecd-8879-4208-a44b-be3afa9c57d9.png" width=400>

## Container Deletion


<img src="https://hedgedoc.plusvasis.xyz/uploads/6df2060a-d4e6-41b9-9445-b6f1a43f75c2.png" width=400>

## Create and Expose a Web Server


<img src="https://hedgedoc.plusvasis.xyz/uploads/8352935e-1676-431a-9c13-9a0aabde3ae2.png" width=400>

## Import Containers from docker-compose file


<img src="https://hedgedoc.plusvasis.xyz/uploads/2baed630-98fb-473b-a483-3c57843cf901.png" width=400>

## Start a Live Terminal Session


<img src="https://hedgedoc.plusvasis.xyz/uploads/b2e7f362-d80e-4083-9e1c-35bdd1bf0f67.png" width=400>

## Real-Time Log Streaming


<img src="https://hedgedoc.plusvasis.xyz/uploads/0c17ef59-34f4-467c-82ae-03f3c5f16f58.png" width=400>


# System Architecture


<img src="https://hedgedoc.plusvasis.xyz/uploads/d0f88b26-f5ae-4fc1-a881-0f7a3a8281e0.jpg" width=400>

PlusVasis consists of the following components:
- **Go REST API**: Acts as a middleman between Nomad and the frontend. The aliveness of this can be checked at https://api.plusvasis.xyz/health/
- **SvelteKit frontend**: Provides the web UI using Svelte components, SvelteKit for routing and server-side rendering, and Typescript for a type-safe implementation. Accessible at https://app.plusvasis.xyz/
- **Nomad**: Orchestrates jobs on the cluster by communicating with clients and the Docker API for each host/node. Also provides functionality for job management and real-time monitoring.
- **Traefik**: Reverse proxy for the whole system. Routes requests to the api, frontend, nomad and any containers that users expose on PlusVasis. Provides integration with both Nomad and Docker to do all of this.
- **Docker**: Driver for executing Nomad jobs. A docker image specified in a Nomad job will then be started on one of the hosts on the Nomad cluster using Docker.
- **Cloudflare**: DNS resolution for the whole system. Provides SSL certificates for all the web-facing components and user containers. Also acts as a proxy to minimize threats from bad actors, e.g. DDoS protection.
- **Firebase**: Authentication layer used by both the frontend and backend to allow separation between user resources and security.


# Implementation

## Backend

The backend consists of a REST API written in Go using the Echo framework. The backend can be seen as the middleman of our application. It allows communication between our frontend and Nomad by communicating with them both directly. Go is a fast and efficient programming language, making it well-suited for building high-performance APIs. Additionally, Echo is a lightweight and easy-to-use web framework that provides a simple and intuitive way to create RESTful APIs.

It includes features such as routing, middleware, and request handling that make it easy to build complex APIs with minimal boilerplate code. Echo also has excellent documentation and a large community of developers, making it easy to find help and support when needed. Finally, Echo is a highly scalable framework, which means it can handle large amounts of traffic and is suitable for building enterprise-grade applications.

Overall, the combination of Go and Echo makes for a powerful and flexible API development platform that can meet the needs of a wide range of use cases.

### Structure


<img src="https://hedgedoc.plusvasis.xyz/uploads/d2559b80-8702-4027-8517-ac03cf5f4684.png" width=400>

### Main


<img src="https://hedgedoc.plusvasis.xyz/uploads/191b66c2-24b2-41de-88ca-598c93dd1c3a.png" width=400>

Here we set up both the middlewares and routes used in our API.

### Routes


<img src="https://hedgedoc.plusvasis.xyz/uploads/8df76a38-c237-4330-9e75-9185ad0cd25e.png" width=400>

The setupRoutes function in our main.go file is where we provide our different route setup functions, organised by purpose - health, nomad jobs and nomad proxy.
Also here we add a route for our swagger api documentation - which can be accessed at the following address: https://api.plusvasis.xyz/swagger/index.html

#### Swagger


<img src="https://hedgedoc.plusvasis.xyz/uploads/d9c2d030-2f7d-4c65-83bf-cffb0c926934.png" width=400>

Swagger allows us to specify schemas and documentation for all of our API routes that can be viewed in a nice web UI. This allows us to manually test our endpoints and ensure that for different inputs, we get a response that is defined in our schema and is therefore an expected response.

#### HealthRoutes


<img src="https://hedgedoc.plusvasis.xyz/uploads/38e5d4bf-5ece-469c-9688-d384bd6895e0.png" width=400>

Simply a basic route that we can use to prove the aliveness of the API easily.

#### NomadJobs


<img src="https://hedgedoc.plusvasis.xyz/uploads/d3af5996-38fb-4211-8cbf-4e8c6447e978.png" width=400>

Here we define all of our routes relating specifically to Nomad jobs and the handler functions that run whenever an endpoint receives a particular request.

#### NomadProxy


<img src="https://hedgedoc.plusvasis.xyz/uploads/3fd789dd-d1ff-47d7-95e5-6a4e2f14d7fe.png" width=400>

Here we define all our two routes that provide a real-time proxy layer between Nomad and our frontend. `/job/:id/exec` for example is a middleman for a two-way websocket communication channel between the frontend and Nomad, used for the terminal on our container page. `/job/:id/logs` provides a stream of data consisting of container logs that can be read from Nomad and then written to the client's buffer in real-time.

### Middleware


<img src="https://hedgedoc.plusvasis.xyz/uploads/6dd8fc9c-f376-4229-98a2-da470b2b6131.png" width=400>

The setupRoutes function in our main.go file is where we set up all of our middleware. Some are native to Echo and others we wrote ourselves.

#### Echo Middleware
- CORS: allows us to define CORS rules such as allowed request origins and allowed headers. This helps us improve security and reduce the likelihood of bad actors trying to access our API.
- Recover: recovers from panics anywhere in the chain and prints a stack trace, which is both useful during development and allows higher availability of the API.
- Secure: protects against cross-site scripting (XSS) attacks, content type sniffing, clickjacking, insecure connections and other code injection attacks.
- RateLimiter: middleware for limiting the number of requests to the API from a particular IP address within a period. Used to increase security and protect against DDoS attacks.

#### Custom Middleware

#### Logger


<img src="https://hedgedoc.plusvasis.xyz/uploads/db1e09ef-ddea-46ab-9b87-44b27ea778c0.png" width=400>

An extension to the default echo logger using [logrus](https://github.com/sirupsen/logrus) and custom output. This was done to improve the readability and flexibility of our API logs. A notable addition is the user id attached to each request, which can potentially allow incident reporting and limiting for users who abuse PlusVasis. This feature is also a nice-to-have for testing and debug purposes.

#### Firebase Authentication


<img src="https://hedgedoc.plusvasis.xyz/uploads/9bf58e28-ff8e-485a-8810-7a118c3370a0.png" width=400>

A custom firebase authentication middleware allows us to easily protect every route in our API with the separation of concerns from our route handlers. JWT tokens are provided in each request either within an `Authorization` header or as a query parameter called `access_token` - this was required because additional headers can not be added to a WebSocket request.

The Skipper function allows us to define endpoints that can bypass Authentication. In our case, we skip authentication for the `/health` endpoint and any endpoints ending with `/swagger`.

This middleware also sets the current user's id in the context, so that we can later check what user made the request and return appropriate responses based on that.

### Handlers/Controllers

#### NomadClient


<img src="https://hedgedoc.plusvasis.xyz/uploads/7b7387db-48c9-45e8-b7c3-e901bde89216.png" width=400>

Provides an interface for communicating with Nomad's API. We use an interface for the communication for separation of concerns and for mocking requests to Nomad in our unit tests - testing of our API should not depend on Nomad.

#### NomadController


<img src="https://hedgedoc.plusvasis.xyz/uploads/daf487a7-5115-4573-b3d7-a485f2b5db3e.png" width=400>

Implements a NomadClient and contains the methods for each API route handler.


<img src="https://hedgedoc.plusvasis.xyz/uploads/53ade090-2ddb-46d9-bf2f-87c8b2134c5c.png" width=400>

The GetJobs handler requests all jobs from Nomad and returns only the jobs that belong to the user that made the request.

Annotated on the method we see our swagger specification for this API route. We specify the type of request (get, post, delete, etc.), and all the possible response codes alongside the types for the structs that then get marshalled before being returned as JSON - for example, on a successful request, this route returns HTTP code 200 and a JSON representation of an array of nomad.JobListStub structs.

Note: the nomad.JobListStub struct comes straight from Nomad's source, and is the same type used by Nomad when generating the response in the first place. Nomad is open-source and its API is written in Go, therefore we used this to our advantage to provide fully type-safe marshalling/unmarshaling of all our requests to and from Nomad.


<img src="https://hedgedoc.plusvasis.xyz/uploads/6ad93a45-5ac1-4880-a489-8aaf933d1634.png" width=400>

The CreateJob handler takes JSON as input, unmarshals to a NomadJob object (this is one of our types, defined in templates), generates a Job JSON object that Nomad understands and sends that to Nomad in a POST request. We then return the response from Nomad to the user.

Notable in the swagger spec is the definition of the NomadJob parameter in the POST body.


<img src="https://hedgedoc.plusvasis.xyz/uploads/db7fe810-5024-4f65-a524-cc6decc8f213.png" width=400>

The UpdateJob handler takes the same input as CreateJob but also checks that the user making the request is the owner of the job that is to be updated. If this check passes, we send the updated job to Nomad and return its response.

Notable in the swagger spec here is the job ID string as a path parameter - `/job/:id/`

From here forward we will omit the swagger annotations as there are no more notable attributes that differ greatly from the examples already given.


<img src="https://hedgedoc.plusvasis.xyz/uploads/89ef2a98-ac22-4d6b-ae2a-d78397cc4dd8.png" width=400>

The ReadJob handler takes a jobId as a path parameter and gets that job from Nomad and returns it to the user if they are the owner.


<img src="https://hedgedoc.plusvasis.xyz/uploads/9a9e27f3-c6da-4c4a-aeb2-edd847f27de6.png" width=400>

StopJob takes a jobId as a path parameter and a query parameter boolean called purge as inputs and will essentially forward this request to Nomad, which has the same inputs for its API endpoint.

Instead of forwarding the request outright, we only check if the `purge` query param is set to true and add that to our request to Nomad. The reason for this is to disallow poisoning of the request to Nomad with unwanted query parameters. For the functionality of PlusVasis, we only care about the purge parameter and therefore it is the only one we send to Nomad. This is done purely for security and exploit mitigation reasons.

We also perform the same check as above that the requesting user owns the job relating to the provided jobId. This check is done in every API handler that performs a read/write/edit on Nomad jobs and will not be mentioned in future to avoid repetition in this document.


<img src="https://hedgedoc.plusvasis.xyz/uploads/0df83c1d-e944-4e18-a157-6c1af0ba9d8e.png" width=400>

The RestartJob handler takes a jobId as input and sends a request to Nomad to Restart the currently running allocation (basically a deployment) for the related job. We send a POST request with an empty body as Nomad handles this by restarting all "tasks" for the given allocation - which for PlusVasis and not going into too much detail on Nomad's API, is exactly what we want to do here. We then return Nomad's response to the user.


<img src="https://hedgedoc.plusvasis.xyz/uploads/86ca1752-4fa9-4e4f-a55f-7bffd8b74a7d.png" width=400>

The StartJob handler takes a jobId as input and starts the related job. Nomad doesn't provide a Start Job endpoint, therefore we read the job the same way as in the ReadJob handler, set the attribute `job.Stop` to false and then update the job - similarly to the UpdateJob handler - before returning Nomad's response.

In Nomad's web UI, they provide a Start Job button, yet through research via the network tab in Chrome, we discovered how Nomad actually does this themselves even though they do not provide an endpoint for it!


<img src="https://hedgedoc.plusvasis.xyz/uploads/3f5090da-bccd-4c3b-84ed-8be9ebf66694.png" width=400>

This method takes a uid and jobId as input and checks that the given uid matches that of the owner of the job relating to the given jobId.

It will return either the error generated in NomadClient for the request to Nomad, a code 500 internal server error if we fail to unmarshal the response from Nomad, a code 401 unauthorized error if the user does not own this job, or nil (no error), meaning the user does own and therefore is allowed to access this job.


<img src="https://hedgedoc.plusvasis.xyz/uploads/f8d0e787-0ad9-4371-9643-f74fcd518ab7.png" width=400>

This method takes a jobId as input and requests the allocations for the related job from Nomad. We then check if any of the allocations are running or pending (still in the process of starting) and return that allocation or a code 404 not found error if one does not exist.

#### NomadProxyController


<img src="https://hedgedoc.plusvasis.xyz/uploads/6d7f301e-7523-48d3-bd27-368785d22016.png" width=400>

The NomadProxyController deals with any handlers/routes that act as a **real-time** middleman or proxy between the user and Nomad.

Here we provide interfaces for WebSockets - so that we can mock them during unit testing - and other configuration options relating to WebSocket connections.


<img src="https://hedgedoc.plusvasis.xyz/uploads/6d6352fa-0eb9-4bf1-b9f4-c4696fa6ab2f.png" width=400>

The AllocExec handler takes a jobId and command - `/bin/sh`, `/bin/bash`, etc. - and builds the appropriate request to send to Nomad to initiate a WebSocket connection to the running alloc for the related job. This is used for the terminal component in the frontend.

This is the most complex handler in the API so therefore we will explain it more in-depth, step-by-step:
- Get jobId and command from the path and query parameters respectively
    - If command was not provided or is empty, return an error
- Check that the requesting user is allowed to access the current job
    - If unauthorized, return an error
- Get the currently running allocation for the current job
    - If there is not one running, return an error
- Build the correct request to send to Nomad to initiate the WebSocket connection
    - path: allocId
    - query: command, task, tty, ws_handshake
- Initiate WebSocket connection with Nomad
    - If the connection fails, return an error
    - Defer closing of this connection (in case of an error later)
- Upgrade the client's request from HTTP to WebSocket, initiating a WebSocket connection with the client
    - If the connection fails, return an error
    - Defer closing of this connection (in case of an error later)
- Set deadlines to close the connection if no messages are received within the idleTimeout period (client should send empty heartbeat messages to keep the connection alive)
    - If this fails for whatever reason, return an error
- Start a go routine to forward messages received from the client to Nomad
- Start a go routine to forward messages received from Nomad to the client
- Wait until both of these go routines have finished executing
- Exit


<img src="https://hedgedoc.plusvasis.xyz/uploads/15b17180-bb18-445f-aac0-cb107ee6ca35.png" width=400>

This method is run in a go routine and is responsible for forwarding messages from one WebSocket connection to another, simply by reading a message from the source connection and writing it to the destination connection.

This method also updates the read deadline (timeout) whenever a message is successfully forwarded.


<img src="https://hedgedoc.plusvasis.xyz/uploads/8b79e08c-3ddb-4276-8757-417fa05f84f1.png" width=400>

The StreamLogs handler takes a jobId as a path parameter and a logType as a query parameter. The logType will be either `stdout` or `stderr`. This handler is used to stream logs in real time from Nomad to the client.

It functions similarly to the AllocExec handler in how it builds the appropriate request to Nomad and therefore will not be explained in depth again here. Where it differs is in the streamResponse method, explained below.


<img src="https://hedgedoc.plusvasis.xyz/uploads/6beed27e-ce14-4876-8c54-864601a1634b.png" width=400>

The streamResponse method takes the current context and the response from Nomad as inputs. It will first check if the request from Nomad is gzip encoded and if so, provides a reader for this. It then creates a buffer from which it will read the response body from Nomad (that contains the current state of the logs as a stream) and then writes this buffer to the response that is sent back to the client.

There are checks for if the data stream has ended, failed, or otherwise does not exist anymore, and appropriate returns and errors for each of these cases.

Additionally, the response to the client is flushed on every loop because writers in Go (such as the http ResponseWriter used here) are buffered until the handler returns, which is not what we want to happen here in a real-time log viewer. Flushing the writer periodically solves this issue, and we decided to flush on every loop to ensure the logs can be viewed by the client in as real-time as possible.

### Templates


<img src="https://hedgedoc.plusvasis.xyz/uploads/a3a41b65-acf6-458c-8fdf-f731c8f93912.png" width=400>

Here is the definition of a NomadJob struct type. Each field is validated using [validator](https://github.com/go-playground/validator/) for the correct type and content.


<img src="https://hedgedoc.plusvasis.xyz/uploads/aab3e6e2-b341-4a4d-8474-890982226a98.png" width=400>

The CreateJobJson method takes a NomadJob struct as input and returns the correct JSON representation of this job that can then be sent to Nomad.

The steps taken by this method to achieve this are commented on clearly and of which's implementation will be explained below.


<img src="https://hedgedoc.plusvasis.xyz/uploads/8c2833c1-c1bf-444a-bc16-019306520956.png" width=400>

The trimEnv method takes a pointer to a NomadJob struct and manipulates the Env field to be in a format that we can use in templating later and will be accepted by Nomad. This involves trimming spaces and trimming quotes if the environment variable value is surrounded by them.


<img src="https://hedgedoc.plusvasis.xyz/uploads/cc42ccb5-15d9-4b23-b0a9-1f2d5446e4c8.png" width=400>

The parseEnv method is where we deal with referencing other jobs in environment variables for inter-container communication. This is done with the `{{jobName}}` syntax and allows a sidecar database container to be accessed by a web server, for example.

This is done via regex matching the environment variable values and checking if a jobName with the `{{jobName}}` syntax exists within. For each instance of this (called a field) we check that only one job is specified per value (if specified at all).

For example `HOT_TAKE="{{nginx}} is better than {{apache}}"` is not allowed, but
`FACT="{{nginx}} == {{nginx}}"` is allowed. This is because Nomad does not easily deal with the former case, as it adds a lot of complexity to make work using Go templates - Nomad also uses Go templates for this functionality!

We ensure this isn't the case by storing the first field value, looping through each field, and returning an error if any of the fields are different.

We then update the templated EnvString field for the NomadJob accordingly, which is explained below.

<img src="https://hedgedoc.plusvasis.xyz/uploads/d20e9af1-b675-46c7-9b51-ee2436e34fdb.png" width=400>

To add some context here, this is one-way Nomad deals with inter-job communication, and is the method we use when creating our Nomad jobs:

<img src="https://hedgedoc.plusvasis.xyz/uploads/f3ee9bcd-be17-479f-8855-d27c8342d27b.png" width=400>

Nomad uses "service discovery" to find the job matching the name given, of which we can access its `Address` and `Port` fields - this will look familiar or will make more sense later, as Nomad uses Go structs to represent its jobs and `Address` and `Port` are fields on that Nomad Job struct.

In this above example, we pass a postgres database URL as an environment variable to a hedgedoc instance.

So in the generateTemplatedEnv method, we take in the key and value for an environment variable, and the name of the job referenced with the `{{jobName}}` syntax within the environment variable's value.

We then replace any instance of `{{jobName}}` with `{{ .Address }}:{{ .Port }}` and use this in a new string called templatedEnv that reflects the example given above, and also escapes any problematic characters.

This templatedEnvString will then get appended to the job.EnvString field in the parseEnv method.


<img src="https://hedgedoc.plusvasis.xyz/uploads/ddc14864-3fdd-4e47-a545-b8a1705d00b6.png" width=400>

Here is the Go template for a JSON representation of a Nomad job, that Nomad can understand and accept. Each field of a NomadJob object can be applied to this template with the `{{.FieldName}}` syntax.

We also store extra information on the current job in the `Meta` section, such as the user that owns the job, the shell they use and all other required metadata of a given job - the reason it is stored here in such a convoluted way is so that PlusVasis does not require a database to store and manage all of this metadata!

Other sections of note here:
- `Job.TaskGroups.Tasks[0].Config.mount`: here we mount a per-user volume to the /userdata path on every container belonging to a user, this is to make it easy to transfer files between containers.
- `Job.TaskGroups.Tasks[0].Templates[0].EmbeddedTmpl`: here is where we input the templated job.EnvString field, explained in depth previously.
- `Job.TaskGroups.Services[0].Tags`: here is where we set the container labels that allow our reverse proxy Traefik to discover the container and expose it publicly on the internet.
    - `${NOMAD_PORT_port}` is a variable set by Nomad that contains the port number of the current service


## Frontend

The frontend of our application is a web app created using SvelteKit, a routing framework build on top of the Svelte JavaScript framework. SvelteKit simplifies the process of building web applications by providing a streamlined development experience and built-in features such as server-side rendering and automatic code splitting.

Our front end project also utilizes Vite, a fast and modern build tool for web applications, which helps to improve the development and build process. 

TypeScript is used throughout the frontend project, providing a statically typed and robust codebase that is easier to maintain and debug. Overall, our front end is designed to provide a fast, intuitive, and user-friendly experience for managing containers and orchestrating applications.

### Structure


<img src="https://hedgedoc.plusvasis.xyz/uploads/c640aa7f-8b41-4251-81d2-139bfb4083b9.png" width=400>

### Routes


<img src="https://hedgedoc.plusvasis.xyz/uploads/b28799c7-112e-4f9d-80df-2e6e87f9d445.png" width=400>

From the get-go of the SvelteKit project initialisation, we already had a boiler-plated format from which we could build upon. SvelteKit utilises the routes directory for all of the pages that a user will physically see. Each of these pages are "+page.svelte" files. For example, the "+page.svelte" file at the top level of the routes directory correlates to the index page of the web app whereas the "+page.svelte" files under the login and logout subfolders represent the /login and /logout pages of the web app respectively.

To explain some basics before describing the implementation further, all Svelte files (.svelte) allow for using a script to specify TypeScript or JavaScript, HTML for what the user will see on that given page and CSS. For example, take a look at our "+page.svelte" file under the login sub-directory.

<img src="https://hedgedoc.plusvasis.xyz/uploads/05cab440-3e85-4bf8-be5a-44da090f14d9.png" width=400>

We also made use of Sveltes layout functionalities. We created a "+layout.svelte" file to specify our CSS across all the routes.

<img src="https://hedgedoc.plusvasis.xyz/uploads/709e2ba7-ff7a-47dc-8b42-2a877a9165b4.png" width=400>

This worked by using a script to import our global CSS which was configured with Tailwind and uses a slot, which is where the current page contents are substituted in.

As well as having layout functionalities, Svelte also has its own error capabilities, so we made a "+error.svelte" file to create an error message to be displayed to users if they stumbled across an invalid URL or wrong page.

<img src="https://hedgedoc.plusvasis.xyz/uploads/7c8c92f7-81bb-4a8f-8079-65e9f456217f.png" width=400>

We wanted to limit the amount of logic stored in these routing files, which led to the use of creating components for functionality, so this left our "+page.svelte" files looking tidy and maximising readability.

For example, here's our index page file:

<img src="https://hedgedoc.plusvasis.xyz/uploads/95fa6772-ef88-44b2-ad55-05b586ef33b3.png" width=400>

And here's our file under the container/[id] subdirectory:

<img src="https://hedgedoc.plusvasis.xyz/uploads/ea35eb55-32b6-450e-b83a-8bebfd26ca8c.png" width=400>

Another cool feature of the routing was for the example above with the container page routing, we could use "[id]" as a naming placeholder for dynamic routes that would be specific for different container names and whatnot.

One last functionality of SvelteKit we used in the routing process, was its server-side load abilities which meant we could render things server-side before a user even hits a page. This eliminated any load times that a user could experience making navigation between pages seamless.

This server-side rendering means PlusVasis does not have any loading spinners when fetching data from our API, which greatly improves the user experience.

In our testing, we noticed improvements in load times up to 3x!

We used a server-side layout file "+layout.server.ts" to check the auth of a current user by searching for a valid token in the cookies. If there was no token present the user would be redirected to the login screen with a code 307 temporary redirect.


This is the "+page.server.ts" file for our index page, which will fetch a user's current containers and have them displayed on the dashboard instantaneously when a user logs in or navigates to the home page.

<img src="https://hedgedoc.plusvasis.xyz/uploads/97feb033-cf3d-439d-ade8-e04ce6e82672.png" width=400>

### Components


<img src="https://hedgedoc.plusvasis.xyz/uploads/eaa74d32-adde-49b6-b33a-d1e16d72bd6f.png" width=400>

Here we stored all the direct page logic in the lib/components directory and our component tests using Vitest.

To start we made use of a LoginForm component.

<img src="https://hedgedoc.plusvasis.xyz/uploads/69d74adb-707c-409a-911a-cb78c2f58327.png" width=400>

This was the component used for the login and signup screens and granted users the option to sign in or create an account with an email or via GitHub. This was done by setting up a Firebase app which will be explained later.

Here's our NavBar component which we used across all the routing.

<img src="https://hedgedoc.plusvasis.xyz/uploads/9ce38162-6a0f-4251-b15a-0abc7578d2ec.png" width=400>

Next, we have the IndexPage component for our home page.

<img src="https://hedgedoc.plusvasis.xyz/uploads/ee57dc90-5b5e-4e4c-986c-6d6212c05282.png" width=400>

This component exports a prop called data, which is then passed in from the server side loading in the "+page.server.ts" file, which allowed us to list the fetched jobs to a user instantly. If no containers were present an alert would be displayed and if there was an error encountered during the server-side fetch we would display the error along with a retry button.

Moving onto the ContainerPage component.

<img src="https://hedgedoc.plusvasis.xyz/uploads/c996bdad-8251-4399-8a0a-1824d0ae36b7.png" width=400>

This component is used for the individual container pages and in itself utilises 5 other components. It uses the Tabs component to flick between three different components while staying on the same route.

Here's a look at the Tabs component.

<img src="https://hedgedoc.plusvasis.xyz/uploads/6d70111f-6ba3-40ad-881f-5fe9c022b734.png" width=400>

It simply creates a Tab bar with a list of TabItems which in our case relate to components. The components we loop through are the NomadController for the container shell, the LogController for the container logs and the SettingsController for the container settings.

Here we have the NomadController.

<img src="https://hedgedoc.plusvasis.xyz/uploads/cf71888d-d5e7-4f23-84f2-d2b409df2a2f.png" width=400>

The NomadController sets current job attributes imported from our stores and also sets the WebSocket URL with the setExecUrl function.

We then use the ExecController Component to display the terminal to a user.

<img src="https://hedgedoc.plusvasis.xyz/uploads/34a88eb7-19f7-473b-93ea-6053eff69608.png" width=400>

This ExecController component is simply used for managing the terminal displayed to a user on the Shell tab of the Container page.

For viewing the container logs we have the LogController component.

<img src="https://hedgedoc.plusvasis.xyz/uploads/db92e20f-8e4b-48a9-bc61-a25b891423c9.png" width=400>

Here we use a stream to read either the standard output or standard error logs of that given container.

For the container settings, we have the SettingsController component.

<img src="https://hedgedoc.plusvasis.xyz/uploads/417d2232-85c7-4cae-a423-65a802cb6b51.png" width=400>

We set the current job attributes by using our stores like before on the NomadController and use the JobForm component with the type "update" to update any of the given current container settings.

Here's the JobForm component.

<img src="https://hedgedoc.plusvasis.xyz/uploads/a1ac897f-1c5e-4221-beb6-8f8e1c629179.png" width=400>

This component is used for creating new jobs or updating jobs and uses a submittable form to do so that either calls our fetchJobCreate or fetchJobUpdate methods on the job and then re-routes the user to either the homepage or the shell tab of the container page.

Now we have our ContainerOptions component.

<img src="https://hedgedoc.plusvasis.xyz/uploads/9658b5ed-177c-437b-8a77-a36f00d3a4cd.png" width=400>

This is used for the four buttons on the right-hand side of the page under the NavBar on the container page, which allows users to start, stop, restart or delete the container by calling methods from our utils to do the fetch requests to our API to make these possible.

Finally, we have the Editor component.

<img src="https://hedgedoc.plusvasis.xyz/uploads/620d9e63-c4a2-40d1-a0e1-575098046ab7.png" width=400>

This is the component used for an in-browser text editor, where users can paste or write a docker-compose file which we then convert into Nomad jobs.

### Utils


<img src="https://hedgedoc.plusvasis.xyz/uploads/f1de22b5-a0da-467f-a6b9-db95a28a8a67.png" width=400>

Here we stored any TypeScript utils that were used across components or the routes pages.


<img src="https://hedgedoc.plusvasis.xyz/uploads/6555bee4-5e74-4626-b658-a7cdd044e8ef.png" width=400>

The Base64Util.ts file is used for message encoding. The reason for this is to allow control characters to be JSON serializable in our live terminal.


<img src="https://hedgedoc.plusvasis.xyz/uploads/c0ed261f-7443-44ee-9176-73b9f680b838.png" width=400>

The ExecSocketAdapter.ts file is used to manage the WebSocket connections - also for the live terminal.


<img src="https://hedgedoc.plusvasis.xyz/uploads/958fc42f-6750-4438-846a-964689638aef.png" width=400>

--

<img src="https://hedgedoc.plusvasis.xyz/uploads/74b5dabc-4201-43da-9b2f-009671815843.png" width=400>

The MakeJob.ts file is used for creating and updating jobs and setting the new attributes from the JobForm component previously mentioned or creating multiple jobs from an imported docker-compose file.


<img src="https://hedgedoc.plusvasis.xyz/uploads/da520aab-49f8-41c3-9707-902b8abfb71e.png" width=400>

--

<img src="https://hedgedoc.plusvasis.xyz/uploads/fa1cc343-8832-4750-bebd-abead297e22a.png" width=400>

The NomadClient.ts file is used to store all the async functions that perform the fetch requests to our API, and is used throughout our frontend.


<img src="https://hedgedoc.plusvasis.xyz/uploads/5d77d544-a593-480e-9bcf-1c44357116d4.png" width=400>

Lastly, the StreamLogs.ts file is used for getting the stream for the container logs.

### Types

In our types directory we have one file Types.ts.

<img src="https://hedgedoc.plusvasis.xyz/uploads/99406a29-5186-4b63-9a65-cabbb97cb4f1.png" width=400>

Here we specify the typings of a Job, the job we wish to create on Nomad for the user's container, JobFields, which we use to define the inputs for a user creating a job, Tab for the tabs displayed to the user, Service, which relates to the service section of a docker-compose and of a nomad job and finally DockerCompose - there can be multiple services in a DockerCompose.

### Firebase

#### Structure


<img src="https://hedgedoc.plusvasis.xyz/uploads/59f1fea1-dec1-4c4d-b2ef-4fd908597299.png" width=400>

#### Admin


<img src="https://hedgedoc.plusvasis.xyz/uploads/6c28912b-64fb-4a70-8f20-96d3192d72b4.png" width=400>

The admin file allows us to get a Firebase auth instance using an admin service account. The credentials for this account are stored in a .env file. These are then used to verify users during any server-side rendered requests that perform a fetch and require the user's id to be sent to the API - this is all without client interaction.

#### Client


<img src="https://hedgedoc.plusvasis.xyz/uploads/460b3692-0be1-4ae9-bfcb-498636e05f47.png" width=400>

The client file is used in any client-side code and is what allows users to log in or register accounts with Firebase. The data for this is also stored in a .env file for convenience's sake, but is PUBLIC as this code runs on the client while the above admin code only runs on the server.

### Stores


<img src="https://hedgedoc.plusvasis.xyz/uploads/4f1dcdea-75bb-4479-bfe2-20b1322f4628.png" width=400>

We use stores in Svelte to store variables or constants that can be accessible across the whole project and from any file (except for server-side fetch requests!)


<img src="https://hedgedoc.plusvasis.xyz/uploads/aa80e8bb-4588-45c6-8bb5-e8ce65683e58.png" width=400>

This is our authStore.ts and it's used to store the user data if they're logged in and it stores the user's token. We also keep the functions corresponding to authentication here such as signing in/ signing out or creating an account etc.


<img src="https://hedgedoc.plusvasis.xyz/uploads/323a252f-ba41-4d3e-a25c-b90bc6682231.png" width=400>

This is our environmentStore.ts used for storing the hostname for our API, which will either be localhost during development or the production API endpoint when the frontend is also deployed in production.


<img src="https://hedgedoc.plusvasis.xyz/uploads/d653b8b1-c35c-44bc-bdb1-9f04aaac7c7a.png" width=400>

Lastly, this is our nomadStore.ts and is used for storing the current job, the current jobs ID and a boolean to determine if the current job is stopped.


## Nomad

https://github.com/hashicorp/nomad
Nomad is an open-source product by HashiCorp (who make Terraform, Vault and Consul also) and is a simple and flexible scheduler and orchestrator for managing containers at scale.

We decided to use Nomad instead of directly communicating with the Docker API ourselves, as it provides scalability and availability by default, and for a project like PlusVasis where users can use a lot of resources, we thought it was important to be scalable from the get-go.

Nomad exposes a HTTP API to provide interaction with jobs, allocations, volumes, etc - and provides great documentation for this, which can be found here: https://developer.hashicorp.com/nomad/api-docs

It is this API that our backend Go REST API sends requests to, to perform all the different operations and features supported by PlusVasis.

### Topology

Here is a screenshot from Nomad's Web UI showing our topology:

<img src="https://hedgedoc.plusvasis.xyz/uploads/15d9198f-a41c-4a30-91d9-31427ebd54fb.png" width=400>

We can see here that we have two nodes - named "leo" and "sonic" - with 11 allocations deployed on "leo" and 1 allocation deployed on "sonic".

### Server-Client

Here is the Nomad config for "leo":

<img src="https://hedgedoc.plusvasis.xyz/uploads/5e53fc61-c494-4ea3-83da-7a43a4de3aab.png" width=400>

We see that the server stanza is enabled, so this node will run as a master server, that does all the orchestrating across the Nomad cluster.

The client stanza is enabled also, so this node will allow allocations to be deployed on it. We specify the IP address of the master server we connect to, in this case, it is the localhost IP address.

We also enable volume support for the docker plugin - this simply allows any docker containers deployed to mount docker volumes.

### Client

Here is the Nomad config for "sonic":

<img src="https://hedgedoc.plusvasis.xyz/uploads/ebfe0e7d-6390-4c8d-8622-a2a798d695bd.png" width=400>

The client stanza is enabled here, but differs from the other config, as the server specified is the IP address of the "leo" node, and not localhost.

Docker volumes are enabled on this node also.

### Example Nomad Job

Here is an example Nomad job that runs an nginx web server:

<img src="https://hedgedoc.plusvasis.xyz/uploads/0436bd4a-bb7b-43a7-95e6-929620e3cb23.png" width=400>


## Traefik

https://github.com/traefik/traefik
Traefik is an open-source reverse proxy and load balancer, which we use for providing SSL certificates for HTTPS and exposing services within PlusVasis.

Traefik provides integration with both Docker and Nomad to discover services deployed on the network and provide routes for accessing these services over HTTP.

### Deployment

Here is our `docker-compose.yml` file for Traefik:

<img src="https://hedgedoc.plusvasis.xyz/uploads/23265dca-2c53-4e63-b84b-750265f285aa.png" width=400>

This is what we use to deploy Traefik, important things to note here are the environment variables `CF_API_EMAIL` and `CF_DNS_API_TOKEN`. These are the credentials for API access to the Cloudflare account that provides the DNS resolution and SSL certificates for PlusVasis and its related domains.

### Configuration

Here is our main `traefik.yml` config file:

<img src="https://hedgedoc.plusvasis.xyz/uploads/01a2db05-7c28-418d-9811-d69be0ac614a.png" width=400>

We see here that we have two entrypoints, `http` and `https` respectively. These are not defining URL schemes, but simply are the names of the entrypoints.

The `http` entrypoint is mapped to port 80 and simply redirects any requests to the `https` entrypoint.

The `https` entrypoint is mapped to port 443 and also defines what certificate resolver to use for the listed domains - in our case we have a "certresolver" called `cloudflare`.

The `cloudflare` certificate resolver uses the API credentials supplied via environment variables in the `docker-compose.yml` file to grab SSL certificates from Cloudflare as needed. A DNS challenge is required to prove ownership of the domain.

The `providers` block defines where Traefik will grab/watch rules and other configurations from, these are the following:
- Docker: from Docker labels for any services within the `proxy` Docker network.
    - example in Traefik's docker-compose file
- File: from a file named `traefik_dynamic.yml`
    - example to be shown below
- Nomad: from any services deployed via Nomad
    - example in the Nomad nginx job provided above

Here is our `traefik_dynamic.yml` file:

<img src="https://hedgedoc.plusvasis.xyz/uploads/234e3a97-2504-44ee-95ee-6656b306b2f0.png" width=400>

In this file we can manually define host:port mappings for a given address, alongside this we can add custom middlewares.

Here we see the rules for Nomad, which is running at `https://192.168.1.201:4646` and also has a middleware that changes the Origin header to be `https://nomad.local.cawnj.dev`. The reason for this is that Nomad's HTTP API does not allow other origins to access it. This rule was added to allow WebSockets requests from domains other than Nomad's own, which is required for our real-time terminal feature.

We also see the rules for a webhook service, this is used for the automatic deployment of PlusVasis on pushes to master from GitLab, which we'll explain next.


## Webhook

https://github.com/adnanh/webhook
We use an open-source project called "webhook" to provide a HTTP endpoint that we can use for continuous deployment purposes.

Gitlab allows us to send a HTTP request on changes to our repository. We use this feature to send a request to the webhook service on any push to the master branch of our project - this includes when any merge requests get merged.

<img src="https://hedgedoc.plusvasis.xyz/uploads/f43812d0-06c2-4f9e-8693-fcefd9c1cef8.png" width=400>

### Configuration

The webhook service is configured as such:

<img src="https://hedgedoc.plusvasis.xyz/uploads/0956406e-2c5f-4d67-bf84-afaed998a70a.png" width=400>

The webhook service allows us to provide trigger rules, which are the following:
- The secret token supplied in the `X-Gitlab-Token` request header matches a given value, redacted here
- The `object_kind` key in the request payload matches the value "push"
- The `red` key in the request payload matches the value "refs/heads/master"

These rules ensure that only we can trigger a deployment, through the use of the secret token, and that the deployment will only ever run on pushes to the master branch.


## Cloudflare

Cloudflare provides both DNS resolution and SSL certificates for PlusVasis and its related domains, as mentioned previously. Cloudflare also provides a proxy layer for any sites managed by it, enabling DDoS protection, analytics and other useful features.

Here are the DNS records exported from Cloudflare:

<img src="https://hedgedoc.plusvasis.xyz/uploads/9e9d51da-dc02-4d22-a843-4070edfda04a.png" width=400>

### A records
- plusvasis.xyz -> the public IPv4 address where PlusVasis is deployed, which is redacted here
- local.plusvasis.xyz -> the local IPv4 address of PlusVasis, for internal use
- *.local.plusvasis.xyz -> a wildcard record that maps any subdomain of local.plusvasis.xyz to the local IPv4 address above

### CNAME records
- *.plusvasis.xyz -> a wildcard record that maps any subdomain of plusvasis.xyz to plusvasis.xyz
    - This is used to provide the DNS resolution for any/all services deployed by PlusVasis.


# Problems Solved

## Framework Choice

When selecting a framework for our backend, we sought a solution that would enable us to build a performant and scalable REST API while minimizing the amount of boilerplate code. We evaluated several options and ultimately decided on Echo, a lightweight and feature-rich web framework for Go.

Echo stood out for its simplicity, ease of use, and excellent documentation. It provides a robust set of features for building RESTful APIs, including routing, middleware, request handling, and more, all while keeping the codebase concise and organized. Echo also has a large and active community, which made it easy to find resources and support when needed.

Moreover, Echo is highly performant and scalable, making it a great choice for building large-scale applications that can handle a high volume of requests. Overall, our choice of Echo as the framework for our backend was driven by its combination of ease of use, feature set, performance, and community support.

## Docker Compose

Support for docker-compose files was something we wanted for the project straight away, as users that are familiar with Docker, are most likely familiar with docker-compose also, as it is often the simplest way to create idempotent deployments with Docker.

While Docker is used to run our containers/jobs, we are not directly communicating with the Docker API, but with Nomad's API instead, therefore we needed to parse docker-compose and convert them to a format that Nomad will understand, not Docker.

Nomad allows the most important features of Docker, but has replacements or its own implementation for features that do not inherently work in a distributed environment. Therefore Nomad job specifications and docker-compose files are not directly translatable, even though on the surface it may seem that they perform the exact same task - container deployment.

The way we dealt with this was by instead of focusing on Docker containers, focusing on Nomad jobs instead. Our backend API's CreateJob endpoint accepts given fields, that we then use to create a Nomad job. What we then did is wrote a parser in the frontend that will grab the data that aligns with these fields from a docker-compose file.

For example, in a docker-compose file that looks like this:

<img src="https://hedgedoc.plusvasis.xyz/uploads/95d2c82e-b8d5-4974-914f-1ce72de81093.png" width=400>


We can convert this into the following JSON blobs (with type-safety of course on both the frontend and backend, which is not as important to explain here):


<img src="https://hedgedoc.plusvasis.xyz/uploads/a53221d4-6324-49dd-9b61-827f2e55e0e8.png" width=400>

and

<img src="https://hedgedoc.plusvasis.xyz/uploads/b9a43e52-b9b4-41f1-ba8f-10f8e60c12a7.png" width=400>


Our API can then accept both of these and create Nomad jobs based on them.
All the options supported in PlusVasis in the API NomadJob type are therefore supported in docker-compose files also.

## Networking

Two features we wanted to provide for users were:
- the ability to publicly exposes their containers
- allow communication between their containers

Enabling both of these features while maintaining the goal of simplicity for users was not going to be an easy task to accomplish, but luckily both Traefik and Nomad provide features that helped us accomplish both of these.

### Exposing containers

Traefik does the work for us here with routing, but we had to implement a way for Traefik to know what to route and where to.

We decided to only allow the use of our https entrypoint in Traefik for simplicity and security purposes. Therefore only HTTP services can be publicly exposed. This means that a database cannot be accessed publicly, but a HTTP service can - but this service can consume data from a database, for example, if it is also deployed on PlusVasis. We felt this is a good compromise without drastically increasing the complexity of both our implementation and user experience.

The way we accomplished this was by allowing the user to enable exposing of their service, and specifying what port to expose.

We used these options in our job template, to allow users to expose their container at a domain like `user123-website.plusvasis.xyz`, for example. Nomad maps the provided port on the container to a random port on the host, and this is then passed to Traefik which can deal with the routing.

### Inter-container communication

For this, we took advantage of Nomad's service discovery and templating features. We use the service discovery to get the IP address and dynamic port of a service and use templating to create an environment variable containing this data.

With this we can set environment variables dynamically, for example, if a container has an environment variable that should contain the connection address to a database, we can resolve this to the correct IP address and port for the container that then contains that database.


## Testing

Testing was a crucial aspect of our project, and we encountered a few challenges while implementing it. Initially, we tested our backend using integration and unit tests, using the Testify library to write assertions. 

However, when we turned to testing our frontend, we faced some issues. We started with Jest due to its mocking capabilities, which would have been helpful for our planned component tests. However, we found it challenging to configure Jest with our SvelteKit project built on Vite and TypeScript.

After some research, we switched to Vitest, which offered similar capabilities to Jest while being better suited for our project. Vitest does not have as much community documentation due to being a less popular project than Jest. But with Vitest, we were able to efficiently test all the components of our frontend.

We later implemented end-to-end system tests using Cypress, leveraging our experience with testing our SvelteKit project.


## Learning New Things

We encountered several new technologies during the development of our project, including SvelteKit, Echo, and Cypress. These were all new to us, and we had to learn how to use them effectively. We were able to overcome this challenge by spending time researching and practising with these tools. 

We started by reading documentation and watching tutorial videos, and then we began experimenting with them in our own code - writing simple proof of concepts to help learn the foundations for these technologies and how we could apply them in our project. We found that by working together and sharing knowledge on this, we were able to learn quickly and effectively.

We also made use of online resources, such as forum posts and discussion pages, to get insight into errors we faced or to see other developers' opinions on certain topics.

Ultimately, we were able to become proficient in these technologies, and we feel that the experience has made us better developers overall.


## Error Handling

Error handling was an important consideration in our project, especially for the backend REST API written in Go using the Echo framework.

We followed the Go convention of emphasizing good error handling, and spent time testing various edge cases and implementing flexible error-handling statements that could handle each scenario.
For example, we utilized HTTP status codes and Echo's wrappers for these status codes to ensure that the API could properly communicate errors to the frontend and so that these errors could be logged effectively in our backend.

In addition, we also utilized Svelte's error-handling capabilities for the frontend, ensuring that errors were caught and displayed to the user in a clear and informative manner.

Overall, we prioritized good error-handling practices throughout our project, both in the backend and frontend, to ensure that our application could both handle and communicate errors properly and provide a smooth user experience.


# Testing

## Integration/ Unit Testing

For testing the backend we wrote a set of integration/ unit tests that test each of the endpoints and controllers functionalities through the use of valid mocks and assertions. We used the testify library for the assertions. For the nomad controller, we were able to mock the nomad client like so.

<img src="https://hedgedoc.plusvasis.xyz/uploads/4f9063e1-5cc5-4df9-80ab-4ffcfd762ae0.png" width=400>

We could then create a setup method to create a new Echo app with the mocked nomad controller and client for testing our requests.

<img src="https://hedgedoc.plusvasis.xyz/uploads/b34cb133-e521-4dfa-b888-b5c89f906bce.png" width=400>

Then for the actual tests, it was just a case of specifying the requests to make and mocking them as they would be called in the production environment. Take a look at these two tests for example, testing the "/jobs" endpoint for getting or creating jobs with either a "GET" or "POST" request.

<img src="https://hedgedoc.plusvasis.xyz/uploads/48614636-55c7-461e-8328-261d5720e82f.png" width=400>

Then for the proxy controller, we were able to do the same mocking for the nomad client but had additional setup and mocking to do, this complex mocking was all related to our use of WebSockets in the proxy controller, and was needed to ensure we were testing these methods effectively.

<img src="https://hedgedoc.plusvasis.xyz/uploads/bb0c6fdd-7acd-4086-9bb3-052df818fea5.png" width=400>

Then to test the proxy endpoints we could use the same structure as used before in the previous tests and test the "/exec" and "/logs" endpoints.

<img src="https://hedgedoc.plusvasis.xyz/uploads/a8857991-d27e-4349-a7bd-7c6738931306.png" width=400>

## Component Testing

To ensure the reliability and functionality of our frontend SvelteKit web app, we employed component testing. For this purpose, we chose to use Vitest, a powerful testing framework that facilitated efficient and comprehensive component testing.

With Vitest, we were able to systematically test each individual component of our web app, ensuring that they functioned correctly in isolation and in conjunction with other components. This allowed us to identify and address any issues or bugs early on, improving the overall quality and stability of our application.

Through component testing, we were able to simulate user interactions, input various scenarios, and validate the expected behaviour of our components. This testing approach provided us with valuable insights into the functionality, responsiveness, and user experience of our web app.

Additionally, Vitest offered a rich set of testing capabilities, such as mocking dependencies and simulating asynchronous operations, enabling us to create robust and reliable tests. It provided us with a comprehensive toolkit for writing assertions, making it easier to verify the correctness of component states, DOM manipulation, and event handling.

By incorporating component testing into our development process, we gained confidence in the stability and correctness of our frontend code. It helped us detect and resolve issues early on, minimizing the chances of encountering unexpected bugs or regressions during the development lifecycle.

Overall, component testing with Vitest played a vital role in ensuring the quality and reliability of our frontend SvelteKit web app, providing us with a solid foundation for delivering a seamless and user-friendly experience to our users.

For some examples take a look at how we tested the IndexPage and ContainerPage components.

This is the IndexPage component test.

<img src="https://hedgedoc.plusvasis.xyz/uploads/c7723ea4-6974-40e0-97fb-b111c55807e4.png" width=400>

This is the ContainerPage component test.

<img src="https://hedgedoc.plusvasis.xyz/uploads/eea10a93-3c82-4d3c-bc16-ffe8f29f3b2e.png" width=400>

## E2E System Testing


<img src="https://hedgedoc.plusvasis.xyz/uploads/da52d4c8-786b-4e55-a770-2387e840dece.png" width=400>

To ensure the overall functionality and integration of our application, we conducted end-to-end (E2E) system testing using the Cypress testing framework. Cypress provided a robust and intuitive platform for simulating real user interactions and validating the behaviour of our application as a whole.

With Cypress, we were able to write comprehensive test scenarios that covered critical user flows and key functionalities. These tests involved simulating user actions, such as clicking buttons, entering data into forms, and navigating between different pages. Cypress provided a user-friendly interface for visually inspecting the application's state during test execution, making it easier to debug and troubleshoot any issues that arose.

Using Cypress's powerful API, we could interact with the DOM elements, make assertions, and verify expected outcomes. We wrote a series of test scripts that encompassed various user journeys, ensuring that our application performed seamlessly under different scenarios.

By executing these E2E system tests, we were able to identify and address any issues related to data flow, navigation, and overall application behaviour. Cypress also provided valuable features like snapshot testing, network stubbing, and time travel debugging, which further enhanced our testing capabilities.

Through E2E system testing, we gained confidence in the reliability and stability of our application across different browsers and devices. It allowed us to validate the end-to-end user experience, ensuring that all components, APIs, and integrations functioned harmoniously.

Overall, E2E system testing with Cypress played a crucial role in validating the integrity and performance of our application. It helped us deliver a robust and user-friendly solution that met the expectations of our users and provided a seamless experience in real-world scenarios.

When running the system tests you'll encounter the nice UI that Cypress offers and be greeted with the test specs.

<img src="https://hedgedoc.plusvasis.xyz/uploads/7b40618e-bec5-45f7-ab69-f0aa215a43e3.png" width=400>

When you run a spec, Cypress allows you to physically view the real-time testing on the site. So let's take the login spec for example, this is what it looks like when run.

<img src="https://hedgedoc.plusvasis.xyz/uploads/c4fc2f07-c227-474c-9e2c-6f855817a2d1.png" width=400>

This is the actual Typescript file then for that login spec.

<img src="https://hedgedoc.plusvasis.xyz/uploads/2fa5a287-2ba1-49fd-8c2e-7f341ef014fd.png" width=400>

We used the same structure and ideology across all the tests and tested the main functionality of the site, such as logging in, creating containers and updating such containers.


# Future Work

While the current implementation of our container management and orchestration platform, PlusVasis, provides a solid foundation for developers, several areas can be further improved and expanded upon in future iterations.

One area of focus for future work is enhancing the platform's scalability and performance. As the demand for containerization continues to grow, optimizing PlusVasis to handle larger workloads, scale, and manage resources effectively will be crucial. This can involve fine-tuning the backend infrastructure, and implementing advanced caching and load-balancing techniques.

Additionally, incorporating advanced monitoring and logging capabilities into PlusVasis would enable developers to gain deeper insights into their containerized applications. Integrating with tools like Prometheus and Grafana can provide real-time performance metrics, while centralized logging solutions such as ELK (Elasticsearch, Logstash, and Kibana) can offer comprehensive log analysis and troubleshooting capabilities.

Furthermore, expanding PlusVasis's integration capabilities with other popular development tools and services can enhance its usability and versatility. Integrating with popular CI/CD platforms like Jenkins or GitLab CI/CD can streamline the deployment pipeline, while adding optional integration with cloud providers like AWS or Azure can enable seamless deployment and scaling of containers for users who want to take advantage of cloud infrastructure.

Lastly, ongoing improvements in security measures will be essential to ensure the protection of containerized applications. Implementing container image vulnerability scanning, enhancing access control mechanisms, and keeping up with the latest security best practices will help safeguard PlusVasis and the applications it manages.

By addressing these areas of future work, PlusVasis can continue to evolve and meet the evolving needs of developers, providing them with a robust and comprehensive container management and orchestration platform.


# References

https://echo.labstack.com/
https://developer.hashicorp.com/nomad/docs?product_intent=nomad
https://kit.svelte.dev/docs/introduction
https://github.com/vitejs/vite
https://github.com/stretchr/testify
https://github.com/vitest-dev/vitest
https://firebase.google.com/docs
https://docs.cypress.io/guides/overview/why-cypress
https://github.com/hashicorp/nomad
https://github.com/traefik/traefik
https://github.com/adnanh/webhook
