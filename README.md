<p align="center">
    <a href="https://skeleton-gadelkareem.herokuapp.com/">
        <img src="./binary/logo/logo.svg" width="400" alt="Skeleton">
    </a>
</p>

# [Skeleton](https://skeleton-gadelkareem.herokuapp.com/)
[![pipeline status](https://gitlab.com/gadelkareem/skeleton/badges/master/pipeline.svg)](https://gitlab.com/gadelkareem/skeleton/commits/master)

A complete Golang and Nuxt boilerplate for your project with backend API, frontend, tests and CI/CD pipelines.

## Build & Run

There are 2 methods of running Skeleton locally

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


