<p align="center">
    <a href="https://skeleton-gadelkareem.herokuapp.com/">
        <img src="./binary/logo/logo.svg" width="400" alt="Skeleton">
    </a>
</p>

# [Skeleton](https://skeleton-gadelkareem.herokuapp.com/)
[![pipeline status](https://gitlab.com/gadelkareem/skeleton/badges/master/pipeline.svg)](https://gitlab.com/gadelkareem/skeleton/commits/master)

A complete Golang and Nuxt boilerplate for your project with backend API, frontend, tests and CI/CD pipelines.

## Features
- Backend written in [Golang](https://golang.org/) using [Beego framework](https://beego.me/).
- Frontend written in [NodeJS](https://nodejs.org/en/) using [NUXT](https://nuxtjs.org/) and [Vue.js](https://vuejs.org/) frameworks.
- Material design using [Vuetify](https://vuetifyjs.com/).
- JSON REST API based on [jsonapi.org](https://jsonapi.org/) standard.
- Fully featured user registration, login, password reminder, profile update, 2FA authentication, SMS based mobile confirmation, ... etc.
- [JSON Web Tokens (JWT)](https://jwt.io/) based authentication.
- [Social login](./src/backend/services/SocialAuthService.go) using Facebook, LinkedIn, Google, Github.
- [SEO friendly](https://en.wikipedia.org/wiki/Search_engine_optimization) thanks to [the NUXT Static Generated (Pre Rendering)](https://nuxtjs.org/guide/#static-generated-pre-rendering).
- [SQL migration](./src/backend/migrations/sql) using [sql-migrate](https://github.com/rubenv/sql-migrate).
- [Rate limiter](./src/backend/limiter) for API routes to easily set a rate limit per IP for one or more routes.
- [Role-based access control (RBAC)](./src/backend/rbac) for API routes and permissions.
- [Pagination](./src/backend/utils/paginator) implementation for API and frontend with caching.
- [Multi-factor authentication](./src/backend/services/AuthenticatorService.go) using [One Time Passwords](https://github.com/pquerna/otp) and mobile SMS code.
- [Nice Email templates](./src/backend/services/EmailService.go) using [Hermes](https://github.com/matcornic/hermes).
- [Fully featured admin dashboard](./src/frontend/src/pages/dashboard) based on [Vuetify Material Dashboard](https://demos.creative-tim.com/vuetify-material-dashboard/?partner=116160&ref=vuetifyjs.com#/).
- [Complete CI/CD pipelines](https://gitlab.com/gadelkareem/skeleton/-/pipelines) including tests using [GitLab .gitlab-ci.yml](.gitlab-ci.yml) file.
- [Deploy to Heroku](#deploy-to-heroku) using few easy steps.
- Automated development initialization using [Docker compose](./docker-compose.yml) and [init file](./init.sh).
- Application Cache using [Cachita](https://github.com/gadelkareem/cachita) with support for memory, Redis, database and file cache.
- [Dependency injection](./src/backend/di/Container.go).
- Backend API integration and unit tests.
- Frontend [Jest](https://github.com/facebook/jest) tests.

## Development

There are 2 methods to run Skeleton locally

- Install required libs locally on OSX:
```bash
./init.sh init
```
Then Run the frontend and backend servers:
```bash
./init.sh
```
-- OR --
- Run the full stack on Docker:
Note that `yarn install` might take some time.
```bash
docker-compose up
# run tests
docker exec -it skeleton_backend_1 /bin/bash -c "go test -v ./... -count=1 | sort -u"
```





## Generate Nuxt static files
```bash
./build.sh
```


## Deploy to Heroku
- [Fork the Skeleton repository on Gitlab](https://gitlab.com/gadelkareem/skeleton/-/forks/new)
- Add the following CI/CD environment variables in [your Gitlab's CI/CD settings section](https://gitlab.com/help/ci/variables/README#custom-environment-variables):
    - HEROKU_API_KEY: Your free [Heroku API KEY](https://dashboard.heroku.com/account)
    - PROD_CONFIG_SECRET_FILE: Base64 encoded string of the `./src/backend/conf/app.prod.ini.secret` file. Use the [./src/backend/conf/app.dev.ini.secret.example](./src/backend/conf/app.dev.ini.secret.example) as an example. 
    ```bash 
    echo "$PROD_SECRET_FILE_CONTENTS" | base64
    ```
- [Run Gitlab pipeline](https://docs.gitlab.com/ee/ci/pipelines/#run-a-pipeline-manually).
- Navigate to [your Heroku apps](https://dashboard.heroku.com/apps) to open your app URL.
- For more information, check [.gitlab-ci.yml](.gitlab-ci.yml) to review how the production container is being generated in the pipelines.
- Optionally you can also push your final image to Docker Hub by adding your username and password as CI/CD environment variables:
    - DOCKER_HUB_USER: docker hub username
    - DOCKER_HUB_PASSWORD: docker hub password

# Services
## Mail service
Skeleton works in development with [MailHog](https://github.com/mailhog/MailHog) in docker which you can access via [http://localhost:8025/](http://localhost:8025/). To use [MailTrap](https://mailtrap.io/) instead, change [./src/backend/conf/app.dev.ini](./src/backend/conf/app.dev.ini) SMTP config to match MailTrap settings.




# Tests 
## Backend
```bash
cd src/backend
go test -v ./... -count=1 | sort -u
#extra
go test -v backend/controllers -count=1 -debug=7 -run  TestUserController_VerifyMobile
```
## Frontend
```bash
cd src/frontend
yarn test
```


