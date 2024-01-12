
## Backend services of Admin Course
### Services
  This is an entry service of our application. 
  This service will accept all requests and collect all required info from other services.<br/>
  Only this service will be with HTTP support. 

[//]: # (* <b>Authorization service</b><br/>)

[//]: # (  This is one of the most important part of every application. In our app this is implemented as RPC service.<br/>)

[//]: # (  You can read more about service in documentation.)
## Guide to run app
### 1. Clone repo to local directory
### 2. Make sure you have installed `Docker` and `docker-compose`
### 2. Execute `docker-compose up` in project folder

#### Documentation is in `docs/gateway.md`

## Guide to database migrations
### Before make sure you have installed golang-migrate CLI tool.
### 1. To create migration execute `make create-migration name='migration_name'`
### 2. Apply migration `make migrate`
### 3. Rollback migration `make migrate-rollback`