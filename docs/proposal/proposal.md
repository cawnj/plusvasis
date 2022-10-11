# School of Computing &mdash; Year 4 Project Proposal Form

## SECTION A

|                     |                   |
|---------------------|-------------------|
|Project Title:       | Continens         |
|Student 1 Name:      | Jason Henderson   |
|Student 1 ID:        | 19309916          |
|Student 2 Name:      | Conor Joyce       |
|Student 2 ID:        | 19425804          |
|Project Supervisor:  | Stephen Blott     |

## SECTION B

### Introduction

Continens will be an easy-to-use interface for Docker container deployments and the infrastructure that enables this. Networking, security and scalability will be  the primary areas covered by the project. Other areas include both frontend and backend development, and user management.

### Outline

Continens will consist of a frontend dashboard where users can spin up containers on demand and connect to them via SSH in the browser. There will also be a settings section where users can expose a given container for public access - for example, a web server.
The frontend will talk exclusively to a backend REST API that will handle container management, user settings and authentication.

### Background

The idea originally came from a desire from Redbrick users to have the ability to host their own projects on Redbrick's servers easily. Docker is an easier to manage approach to deploying projects compared to bare metal, so we figured a web UI that does all of this for the user would be beneficial.

From both of our past experiences using Docker for our 3rd year project, we spent a lot of time setting up infrastructure to deploy our backend with Docker, and making our REST API publicly accessible - so a tool like this may have been useful for us during that time.

The SoC's Termcast tool was an inspiration for this project also as it allows user to easily SSH into SoC servers to do their programming labs from their browser.

Continens will be an all-in-one tool that lets users do any of these tasks, plus whatever else they can imagine, as long as what they want to do can be containerised.

### Achievements

The project will provide an easy-to-use system for any developer - the ability to pull existing Docker images or even create their own.
Users can spin up a database for example, or host their own website, or the backend for their mobile app, all via Docker containers.
Users will be able to pick and choose what container to expose publicly on a subdomain matching their username, for example the user 'johndoe' could host their website at the URL 'johndoe.continens.xyz'.

### Justification

Continens will be useful for anyone who uses Docker on a daily basis and wants a nice interface listing their container instances, or to create new ones with - like a hub or dashboard for all their containers. Continens will make calls to the backend API, which will prove to be useful in doing all the work for the user. For example, when they click a given container and then are instantly connected to that instance via SSH in their browser, without having to manually SSH from their own machines.

### Programming language(s)

The backend REST API of Continens will be written in Go, along with a frontend written in Javascript. The current plan is to use the web framework Svelte to help with development, and to enable the site to be a performant single-page application.

### Programming tools / Tech stack

- SvelteKit: Web framework for Svelte
- Node.js: JavaScript runtime to host server-side JS
- TailwindCSS: CSS framework that integrates well with component based JS frameworks like Svelte
- Firebase: User authentication and database for user information
- Traefik: Reverse proxy that will handle all HTTP requests to the frontend
- Docker: Tool for running software containers
- Wetty: SSH in the browser
- Nomad: Container orchestration tool for availability and scalibility
- Consul: Key-value store to help with service networking

### Hardware

- Home server: Will be used as a server node and for container orchestration
- Raspberry Pi: Will be used solely as a server node

### Learning Challenges

- Svelte: How best to create a single-paged application in a component-based web framework like Svelte
- TailwindCSS: How to efficiently use utility CSS frameworks, like TailwindCSS, to improve both the design of our website and the developer experience of implementing our design
- Docker: How to use more advanced features of Docker to separate users from one another and interacting with the Docker API
- Traefik: How to deal with constantly changing routing rules for each user and their given subdomain
- Nomad: Everything involving this, integration with Consul, and how container orchestration even works
- Go REST API: Potentially will have to learn a framework or router for our API

### Breakdown of work

#### Student 1 (Jason)

We're planning on splitting the frontend evenly, as we both have an interest in where our development could lead us and share innovative ideas that will be ideal in pair programming scenarios. However I will be taking charge of the auth on the frontend. In regards to the backend, I will be focusing primarily on the API and also in setting up Docker in collaboration with Conor, as I'm keen to learn more on it and could learn from Conor's past experiences using it. Relative to the frontend auth, I'll be responsible for the setup and integration of Firebase. I plan to focus heavily on using a test driven development approach in the dev of our project, so for any work I carry out, there will be a unit test to follow it. Finally, as well as unit tests, I will be developing the Integration tests for the overall project as I have previous experience in Integration Testing.

#### Student 2 (Conor)

I will be focusing a lot more on the infrastructure behind our project. For example, setting up Docker, Nomad, Consul and Traefik. A lot of the work required to do this is new to us, but because I have the most experience, I'll be the one focusing on them the most.
I will also be the one setting up CI and CD for our project, as early on as possible, as we really want to push for a test-driven development approach. We both will be writing either unit tests or integration tests for all the code we write, so a good CI/CD setup will benefit us here greatly.
We are sharing the work on the frontend due to both of us wanting to gain more experience using the web technologies we have chosen, and because the frontend is the best showcase of what our project can/will do, having both of us working equally on this, we hope will end with a better result that we are both happy with.
I will also be focusing on authentication in the backend and working on deploying and configuring Wetty so that it can easily integrate with the frontend aspects Jason will be setting up.

