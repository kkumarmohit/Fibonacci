# CE-PE Technical Assessment

This technical assessment forms part of the interview process for the role
inside Central Engineering - Productivity Engineering (CEPE) at Arm.
It is based around a small demo program that exposes an HTTP API to return
requested values from the Fibonacci sequence.

We have packaged it up with a Docker build environment to make it easier for
you to build and run. In order to run it, you'll need to have Docker
and Docker Compose V2 installed.

You will also need a way of making requests to an HTTP API.
On Unix systems such as Linux and OSX, the `curl` command line tool should
be available. On Windows systems, you will need to install curl or another
tool (such as [Postman](https://www.postman.com/downloads/)) to interact
with the HTTP API.

Before you start the tasks, it would be worth familiarising yourself with
the demo application. The following shows how to start it up, and a couple
of example queries:

```sh
# Starting the demo app. The first time will be slower as it needs to build.
# From this directory:
docker compose up -d

# Making requests to the demo app
curl http://localhost:8080/fibonacci?order=4
{"order":4,"value":3}
curl http://localhost:8080/fibonacci?order=10
{"order":10,"value":55}

# After making code changes, you will need to force docker compose to rebuild
# the image rather than just running the one that has already been built
docker compose up --build -d

# Shutting down the demo app
docker compose down
```

## Tasks

Now that you are familiar with the demo app and have managed to build and run
it, here are some tasks for you to attempt. We would like to see an attempt
at one of them. Have a read through and choose the task that interests you the
most, and/or the task which will best let you exibit your skills.

Once you have completed your task, please re-compress the entire `fibonacci-api`
directory and send it back to us. Ahead of the interview, we will look through
your code, and during the interview we will go through your solution together
and ask some questions about it.

### Add Recursive Endpoint to Demo App

Currently, the API exposes an endpoint `/fibonacci` that uses an iterative
method for calculating the value of the requested order of the Fibonacci
sequence. Please extend the program to add a second endpoint:
`/recursive-fibonacci`. This endpoint should use a recursive method for
calculating the value.

### Create LRU Cache for Demo App

Create a companion app which implements a Least Recently Used (LRU) cache.
The cache app should have the following properties:

- It should be a separate program, not just another endpoint on the demo app
- It should store a maximum of 5 results in the cache
- It should implement a 'least recently used' cache eviction policy.
- If the requested sequence order is already in the cache, it should be returned
- If the requested sequence order is not in the cache then it should be fetched from
  the demo app, stored, and returned.
- It should be built and run using Docker
  - It should be built as a Docker image
  - It should be included in the docker-compose.yml file
- It should listen on port 8081
- It should expose a `/fibonacci` endpoint which caches, and proxies the results
  of the `/fibonacci` endpoint of the demo app.

The message format of the responses from the cache application should be of the form:

```sh
# Request:
curl http://localhost:8081/fibonacci?order=4
# Response (don't worry about indentation, this is just to make it easier to read)
{
    "fibonacci": {
        "order": 4,
        "value": 3
    },
    "cached": true (or false)
}
```

### Re-implement the Demo App in a Language of Your Choice

Implement an app which has equivalent functionality in a language of your choice.
Your re-implementation should satisfy the following constraints:

- It should expose the same HTTP API with an identical `/fibonacci` endpoint
- It should be built and run using Docker
  - It should be built as a Docker image
  - It should be included in the docker-compose.yml file

You may also want to add the `/recursive-fibonacci` endpoint to your implementation.

## A Note on Expected Duration

Whichever task you have chosen, we don't expect you to spend more than a couple
of hours on your solution: it should fit comfortably into one evening, so please
don't spend a week refactoring your code!

## A note for Windows Users

Installing Docker Desktop is a little more involved on Windows as it requires the
use of the 'Windows Subsystem for Linux' (WSL2) and also some virtualisation features
to be enabled in the BIOS. The
[docker docs](https://docs.docker.com/docker-for-windows/install/) are helpful but you
may prefer to set up a linux virtual machine on your computer (e.g. using VirtualBox),
or use a [free AWS EC2 instance](https://aws.amazon.com/free) instead.

## A note on Docker Compose

The instructions in this guide use [Docker Compose V2](https://docs.docker.com/compose/migrate/).
This comes bundled with the latest versions of Docker Desktop.
Manual instructions for installing on Linux can be found
[here](https://docs.docker.com/compose/install/linux/#install-using-the-repository).

If you have the deprecated Docker Compose V1 tool installed these instructions will
continue to work but you will need to change the commands `docker compose` to
`docker-compose`.


## A note on solution

I don't have much experience with Go, hence I did the implmentation in Python. Then using copilot, I converted the same code for Go with some refactoring.