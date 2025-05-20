<p align="center">
    <a href="https://skeleton-gadelkareem.onrender.com/">
        <img src="./binary/logo/logo.svg" width="400" alt="Skeleton">
    </a>
</p>

# [Skeleton](https://skeleton-gadelkareem.onrender.com/)
[![pipeline status](https://gitlab.com/gadelkareem/skeleton/badges/master/pipeline.svg)](https://gitlab.com/gadelkareem/skeleton/commits/master) <a href="https://github.com/gadelkareem/skeleton"><img src="https://github.githubassets.com/images/modules/logos_page/Octocat.png" width="25" height="25" alt="Github Mirror"></a> <a href="https://gitlab.com/gadelkareem/skeleton"><img src="https://about.gitlab.com/images/press/logo/png/gitlab-icon-rgb.png" width="30" height="30" alt="Github Mirror"></a>


A complete Golang and Nuxt boilerplate for your project with Subscription management system, backend API, frontend, tests and CI/CD pipelines.

## [Demo](https://skeleton-gadelkareem.onrender.com/)

## Features
- Subscription management system using [Stripe](https://stripe.com/) API.
- Backend written in [Golang](https://golang.org/) using [Beego framework](https://beego.me/).
- Frontend written in [NodeJS](https://nodejs.org/en/) using [NUXT](https://nuxtjs.org/) and [Vue.js](https://vuejs.org/) frameworks.
- Material design using [Vuetify](https://vuetifyjs.com/).
- JSON REST API based on [jsonapi.org](https://jsonapi.org/) standard.
- Fully featured user registration, login, password reminder, profile update, 2FA authentication, SMS based mobile confirmation, audit logs ... etc.
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
- [Beautiful home page](./src/frontend/src/pages/index.vue) based on [Veluxi Starter](https://github.com/ilhammeidi/veluxi-starter).
- [Complete CI/CD pipelines](https://gitlab.com/gadelkareem/skeleton/-/pipelines) including tests using [GitLab .gitlab-ci.yml](.gitlab-ci.yml) file.
- [Deploy to Render](#deploy-to-render) using few easy steps.
- Automated development initialization using [Docker compose](./docker-compose.yml) and [init file](./init.sh).
- Application Cache using [Cachita](https://github.com/gadelkareem/cachita) with support for memory, Redis, database and file cache.
- [Dependency injection](./src/backend/di/Container.go).
- Backend API integration and unit tests.
- Frontend [Jest](https://github.com/facebook/jest) tests.
- [Queue management system](./src/backend/queue) using [Que](https://github.com/gadelkareem/que).

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


## Deploy to Render
- [Fork the Skeleton repository on Gitlab](https://gitlab.com/gadelkareem/skeleton/-/forks/new)
- Add the following CI/CD environment variables in [your Gitlab's CI/CD settings section](https://gitlab.com/help/ci/variables/README#custom-environment-variables):
    - PROD_CONFIG_SECRET_FILE (optional): Base64 encoded string of the `./src/backend/conf/app.prod.ini.secret` file. Use the [./src/backend/conf/app.dev.ini.secret.example](./src/backend/conf/app.dev.ini.secret.example) as an example.
    ```bash
    cat "$PROD_CONFIG_SECRET_FILE" | base64
    ```
- [Run Gitlab pipeline](https://docs.gitlab.com/ee/ci/pipelines/#run-a-pipeline-manually).
- Create a new Web Service on [Render](https://render.com/) and connect it to your repository.
- Use the provided `render.yaml` to create the required services on Render.
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
#colourfull
go test -v ./... -count=1 | sed ''/PASS/s//(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//(printf "\033[31mFAIL\033[0m")/''
```
## Frontend
```bash
cd src/frontend
yarn test
```


