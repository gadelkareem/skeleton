<p align="center">
    <a href="https://skeleton-gadelkareem.onrender.com/">
        <img src="./binary/logo/logo.svg" width="400" alt="Skeleton">
    </a>
</p>

# [Skeleton](https://skeleton-gadelkareem.onrender.com/)
[![CI](https://github.com/gadelkareem/skeleton/actions/workflows/ci.yml/badge.svg)](https://github.com/gadelkareem/skeleton/actions/workflows/ci.yml) [![GitHub](https://img.shields.io/badge/GitHub-gadelkareem%2Fskeleton-blue?logo=github)](https://github.com/gadelkareem/skeleton) [![GitLab](https://img.shields.io/badge/GitLab-gadelkareem%2Fskeleton-orange?logo=gitlab)](https://gitlab.com/gadelkareem/skeleton)


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
- [Complete CI/CD pipelines](https://github.com/gadelkareem/skeleton/actions) including tests using [GitHub Actions workflow](.github/workflows/ci.yml) file.
- [Deploy to Render](#deploy-to-render) using few easy steps.
- Automated development initialization using [Docker compose](./docker-compose.yml) for containerized setup or native macOS setup scripts.
- Application Cache using [Cachita](https://github.com/gadelkareem/cachita) with support for memory, Redis, database and file cache.
- [Dependency injection](./src/backend/di/Container.go).
- Backend API integration and unit tests.
- Frontend [Jest](https://github.com/facebook/jest) tests.
- [Queue management system](./src/backend/queue) using [Que](https://github.com/gadelkareem/que).

## Development

There are 2 methods to run Skeleton locally:

### Method 1: Native macOS Development
Install required dependencies locally on macOS:
```bash
./setup-mac.sh
```
Then run the frontend and backend servers:
```bash
./start-dev.sh
```

### Method 2: Docker Development
Run the full stack on Docker (works on any OS):
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

### Prerequisites
1. [Fork the Skeleton repository on GitHub](https://github.com/gadelkareem/skeleton/fork)
2. Create accounts for:
   - [Render](https://render.com/) - for hosting
   - [Mailgun](https://www.mailgun.com/) - for sending emails
   - (Optional) Social auth providers: GitHub, Google, Facebook, LinkedIn
   - (Optional) [Twilio](https://www.twilio.com/) - for SMS verification

### Setup Steps

#### 1. Configure Production Secrets
Create your production secret file from the template:
```bash
cp src/backend/conf/app.prod.ini.secret.example src/backend/conf/app.prod.ini.secret
```

Edit `src/backend/conf/app.prod.ini.secret` and replace all placeholders:
- Generate HMAC key: `openssl rand -hex 32`
- Add your Mailgun SMTP credentials
- Add social auth credentials (if using)
- Add Twilio credentials (if using SMS)

#### 2. Add GitHub Secrets
Go to your repository Settings → Secrets and variables → Actions, and add:

**Required:**
- `PROD_CONFIG_SECRET_FILE`: Base64 encoded production secret file
  ```bash
  cat src/backend/conf/app.prod.ini.secret | base64 -w 0
  ```

**Optional (for Docker Hub):**
- `DOCKER_HUB_USER`: Your Docker Hub username
- `DOCKER_HUB_PASSWORD`: Your Docker Hub password

#### 3. Deploy to Render
- Create a new Web Service on [Render](https://render.com/) and connect it to your forked repository
- Use the provided `render.yaml` to automatically create:
  - PostgreSQL database (free tier)
  - Redis cache (free tier)
  - Web service with auto-deploy enabled
- Trigger the GitHub Actions workflow from the Actions tab to build and push the Docker image
- Render will automatically pull and deploy the image

#### 4. Verify Deployment
- Check your Render dashboard for the deployment status
- Visit your deployed URL (e.g., https://your-app.onrender.com)
- Check logs in Render dashboard if there are any issues

### Architecture
The CI/CD pipeline (see [.github/workflows/ci.yml](.github/workflows/ci.yml)):
1. Builds and tests backend (Go)
2. Builds and tests frontend (Nuxt.js)
3. Combines both into a single Docker image
4. Pushes to Docker Hub (optional)
5. Render pulls and deploys the image

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


