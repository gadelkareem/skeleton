<p align="center">
    <img src="./binary/logo/logo.svg" width="400" alt="Skeleton">
</p>

# Skeleton

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


## Gitlab CI
Check [.gitlab-ci.yml](.gitlab-ci.yml) to review how the production container is being generated in the pipelines.














# Services
## [Mailhog](https://github.com/mailhog/MailHog) 
[http://localhost:8025/](http://localhost:8025/)



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



## Heroku
https://devcenter.heroku.com/articles/build-docker-images-heroku-yml#creating-your-app-from-setup
