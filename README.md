# Habibi Salon

## A system to manage everything in a beauty salon.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

1. Clone the repository: 

```bash
 git clone https://github.com/RobTov/habibi-salon
```

2. Setup The Database (default port is 5498, you can change it on the docker-compose.yaml file
but keep in mind that you will need to change it on the environment variables or in cmd/api/main.go and cmd/migrate/seed/main.go files):

```docker
 docker-compose up -d 
```

3. Make the migrations to the database (this will create all the tables, funcions, triggers, etc...):
```bash
 make migrate-up 
```
4. Seed the database (this will populate the database with default values):
```bash
 make seed 
```
## Usage

1. Run the backend:

```makefile
make run
```

2. Run the frontend:

```makefile
make run-frontend
```

3. Create a new Migration:

```makefile
make migration <<migration_name>>
```
