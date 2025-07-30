# Delivery service with microservices
## Table of contents
1. [Technology Stack](#technology-stack)
2. [Structure](#structure)
3. [Usage](#usage)

## Technology Stack
1. Go v1.24
2. gRPC v1.74.2
3. PostgreSQL v17.5

## Structure
- __env_examples__ - examples of all required env files
- __src__ - sources folder
    - __api_gateway__ - web service api gateway
        - __clients__ - gRPC clients for communications with microservices
        - __handlers__ - logical joined url groups of API handlers
        - __model__ - GO type struct data objects with their transformations to different formats
        - __router__ - connections url requests with their handlers
        - __util__ - shared useful functions
        - __main.go__ - api gateway entry point
    - __database__ - all DB configuration and init scripts
        - __db_scripts__ - sql files for init PgSQL
    - __proto__ - services .proto files
    - __services__ - microservices
        - __user__ - user microservice
            - __model__ - GO type struct user data object with him transformations to gRPC and DB row formats
            - __repository__ - communicators between DB and microservice
            - __util__ - shared useful functions
            - __main.go__ - user microservice entry point
            - __user.go__ - init file gRPC user microservice
- __docker-compose.yaml__ - docker compose start file

## Usage
1. Install git
```bash
sudo apt-get update
sudo apt-get install git
```

2. Clone repository
```bash
    git clone https://github.com/SkiStalker/DeliveryServiceGO
```
3. Copy __ALL__ env files from [env_examples](./env_examples) folder to project root and specify your own values in them

4. Install docker <br>
See [how to install docker](https://docs.docker.com/desktop/setup/install/linux/)

5. Start docker compose
```bash
docker compose up
```