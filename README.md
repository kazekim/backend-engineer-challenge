# Backend Engineer Coding Challenge

> This repository contains the coding challenge for backend engineers.

**Note:** Please don't fork this repository, create a pull request against it, or use GuardRails in the repo name. Otherwise other candidates may take inspiration from it. Once the coding challenge is completed, you can submit it via this [Google form](https://forms.gle/i5nZWZKoUnTWj3td9).

## Description

Build a simple code scanning application that detects sensitive keywords in public git repos.
The application must fulfil the following requirements:
- A user can CRUD repositories. A repository contains a name and a link to the repo
- A user can trigger a scan against a repository
- A user can view the Security Scan Result ("Result") List

How to do a scan:
- Just keep it simple by iterating the words on the codebase to detect secrets findings
- A secret starts with the prefix `public_key` or `private_key`

The Result entity should have the following properties and be stored in a database of your choice:
- Id: any type of unique id
- Status: "Queued" | "In Progress" | "Success" | "Failure"
- RepositoryName: string
- RepositoryUrl: string
- Findings: JSONB, see [example](example-findings.json)
- QueuedAt: timestamp
- ScanningAt: timestamp
- FinishedAt: timestamp


**What we want to see:**
- Project Structure: Clear organization and structure of folders, code and functionality
- Containerized app
- Clean Code: Code consistency, use of linters, formatting, error handling, and anything else that shows your skills. Simple is better than complex
- Stack Knowledge: Proper use of Golang and selected frameworks / libraries
- Implementation: The implementation has to work according to the specs
- Unit Tests: Covering the core functionality with unit tests
- Proper Documentation: 
    - A High-Level Design for the components / infrastructure if any
    - Describe how you came up with the solution and what makes it a good one for the use-case
    - Describe how to configure the project, how to start it, how to test it etc

**Bonus points for:**
- SQL schema
- API documentation
- Use of appropriate design patterns
- Microservice Architecture
- Any extra feature (just write it in your documentation)

**Things you don't have to worry about:**

- Authentication / Authorization
- CI configuration / Deployment
- APM


## Scoring

| General                | Points |
|------------------------|--------|
| Project structure      | 0-5    |
| Containerized app      | 0-2    |
| Clean / clear code     | 0-5    |
| Stack Knowledge        | 0-5    |
| Implementation         | 0-5    |
| Unit tests             | 0-10   |
| README & Documentation | 0-5    |
| Commit messages        | 0-2    |
| Libraries              | 0-2    |
| Code Comments          | 0-2    |


| Bonus             | Points |
|-------------------|--------|
| SQL Schema        | 0-2    |
| API Documentation | 0-2    |
| Design Patterns   | 0-2    |
| Architecture      | 0-2    |
| Extra Features    | 0-2    |


## Concept

This project will build 5 containers for running the application. It will have

- Challenge API Service : web service for calling api
- Scanning Worker Service : worker service for doing the job (scanning the git repository)
- PostgreSQL Service : for storing data
- Message Queue Service (Kafka) : for manage job queue

You can see image below to see the application structure.

![](document/guardrails-challenge.jpg?raw=true)

This is only the initiate structure. In the future, it will need to scale up and will have 
more services to add into it.

## How it works?

- Users will call request to 'Challenge API' to add their git repositories name and url. 
- They can edit any information of the repository or delete some through the api too.
- Users can list and get information of repository through API
- They can have a Security Scan by calling API. The scanning process will take time 
depends on how big the repository it is. Users can go back to see the scanning status 
through the API. Then, when the scanning process have beed done and have status 'Success', 
they can see the scanning result via the Scan Result API.

## Implementation

**Requirement:**

Please make sure that you have installed these applications before you can run the project
- Docker
- Docker-Compose
- Makefile
- Postman (for opening API document)
- Golang (only if you need to have local develop)

**How to run project:**

Before you start, please note that this setup will be used for local build only. If you need to deploy it somewhere, 
you need to change some configurations first before you can do it.

- To initial the enviroment, please run this command
> make init

It will initiate all required files and config including creating docker environments for running the project

- Start the Kafka and PostgreSQL docker by running this command

> make service-start

This command will help you create database for the project include migrate some initial database schemas into it.

**Important:** Before go to the next step please make sure that the Kafka have finished starting. Kafka take quite a long time 
before it can be finished. So you can use this command to check if it was finished starting or not.

> make log-kafka

- This step is to run Challenge API and  Challenge Worker. If you run in the first time, you should build the containers
before. Please use this command to build the containers and start the containers automatically.

> make api-build-start

But if you don't need to start docker service yet and need to build the containers only, you can use this command.

> make build

Later, if you have already built the containers or is not the first time you run the containers, you can use this command 
to start the api and worker container directly. (This way will be faster because we don't need to build the containers)

> make api-start

You can see log of API with this command

> make log-api

You can see log of Worker with this command

> make log-worker

## How to use API?:

You should import postman file into the Postman application and open it. You will see list of API like this

![](document/api-list.png?raw=true)


[Download Postman Collection](postman_collection.json)

[DownloadPostman Environment](postman_environment.json)

There are set of API request group by use cases. (You can see in Postman file to look for example request)

**Repository**

- Create : api for creating git repository data
- Update : api for updating git repository data
- Delete : api for deleting git repository data
- Get : api for getting git repository data by repository id
- List : api for listing git repository data

**Repository > Scanning**

- Start Scan by Repository ID : api for start scanning git repository by repository id
- Scan Result List by Repository ID : api for list the scan result base on repository id  with filters

**Scan Result**

- Get : api for getting scan result data by result id
- List : api for listing scan all result data in system with filters


## Local Development

To work in local development, you need to know where you can edit the config file and execute the project

**Challenge API**

- Config file is stored in grchallengeapi/configs/config.yml
- Go's main file is stored in grchallengeapi/cmd/grchallengeapi/main.go

**Scanning Worker**

- Config file is stored in grscanningworker/configs/config.yml
- Go's main file is stored in grscanningworker/cmd/grscanningworker/main.go