<p align="center">
    <a href="https://skeleton-gadelkareem.onrender.com/">
        <img src="./binary/logo/logo.svg" width="400" alt="Skeleton">
    </a>
</p>

# [Skeleton](https://skeleton-gadelkareem.onrender.com/)
[![CI](https://github.com/gadelkareem/skeleton/actions/workflows/ci.yml/badge.svg)](https://github.com/gadelkareem/skeleton/actions/workflows/ci.yml) [![GitHub](https://img.shields.io/badge/GitHub-gadelkareem%2Fskeleton-blue?logo=github)](https://github.com/gadelkareem/skeleton) [![GitLab](https://img.shields.io/badge/GitLab-gadelkareem%2Fskeleton-orange?logo=gitlab)](https://gitlab.com/gadelkareem/skeleton)


A complete Golang and Nuxt boilerplate for your project with Subscription management system, backend API, frontend, tests and CI/CD pipelines.

## [Demo](https://skeleton-gadelkareem.onrender.com/)

## Quick Start

```bash
# macOS Setup (one-time)
./setup-mac.sh

# Start Development Servers
./start-dev.sh

# Run Tests
cd src/backend && go test -v ./... -count=1
cd src/frontend && yarn test

# Build for Production
./build.sh
```

**Servers:**
- Frontend: http://localhost:8080
- Backend API: http://localhost:8081/api/v1
- MailHog (email testing): http://localhost:8025

## Table of Contents
- [Features](#features)
- [Development](#development)
- [Deploy to Render](#deploy-to-render)
- [Configuration Files](#configuration-files)
- [Project Structure](#project-structure)
- [Development Scripts](#development-scripts)
- [Services](#services)
- [Tests](#tests)
- [API Documentation](#api-documentation)
- [Security Features](#security-features)
- [Tech Stack](#tech-stack)
- [Performance Optimization](#performance-optimization)
- [Contributing](#contributing)

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

### Troubleshooting Render Deployment

**Build fails with "Dockerfile.ci not found":**
- Make sure you're deploying from the latest master branch
- The `docker/Dockerfile.ci` file was added to fix this issue

**Application returns 403 or 500 errors:**
- Check Render logs for detailed error messages
- Verify `PROD_CONFIG_SECRET_FILE` is set correctly in GitHub secrets
- Ensure your production secret file has valid credentials (Mailgun, etc.)

**Database connection errors:**
- Render automatically provides `DATABASE_URL` environment variable
- Check that migrations ran successfully in the logs
- Verify PostgreSQL service is running in Render dashboard

**Redis connection errors:**
- Render automatically provides `REDIS_URL` environment variable
- Check Redis service status in Render dashboard
- The app falls back to memory cache if Redis is unavailable

**Email sending fails:**
- Verify Mailgun credentials in production secret file
- Check Mailgun domain is verified and active
- Review email logs in Mailgun dashboard

## Repository Mirrors

This repository is mirrored on both GitHub and GitLab:
- **GitHub** (primary): https://github.com/gadelkareem/skeleton
- **GitLab** (mirror): https://gitlab.com/gadelkareem/skeleton

Both repositories are kept in sync automatically. You can clone from either location.

## Configuration Files

### Development Configuration
- `src/backend/conf/app.dev.ini` - Development settings (ports, local database, etc.)
- `src/backend/conf/app.dev.ini.secret.example` - Template for development secrets
- Create `src/backend/conf/app.dev.ini.secret` from the example for local development

### Production Configuration
- `src/backend/conf/app.prod.ini` - Production settings (uses environment variables)
- `src/backend/conf/app.prod.ini.secret.example` - Template for production secrets
- Production secrets are injected via GitHub Actions from `PROD_CONFIG_SECRET_FILE`

### Default Configuration
- `src/backend/conf/app.default.ini` - Base configuration included by both dev and prod

### Environment Variables
The following environment variables are used in production:
- `DATABASE_URL` - PostgreSQL connection string (auto-provided by Render)
- `REDIS_URL` - Redis connection string (auto-provided by Render)
- `BEEGO_RUNMODE` - Set to `prod` for production
- `CACHE_TYPE` - Cache backend type (memory, redis, database, file)

## Project Structure

```
skeleton/
├── .github/
│   └── workflows/
│       └── ci.yml              # GitHub Actions CI/CD pipeline
├── docker/
│   ├── Dockerfile.ci           # Production Dockerfile for CI/CD
│   ├── Dockerfile.gitlab       # GitLab deployment Dockerfile
│   └── db/                     # Database initialization scripts
├── src/
│   ├── backend/                # Go backend application
│   │   ├── conf/               # Configuration files
│   │   ├── controllers/        # API controllers
│   │   ├── di/                 # Dependency injection
│   │   ├── limiter/            # Rate limiting
│   │   ├── migrations/         # Database migrations
│   │   ├── models/             # Data models
│   │   ├── queue/              # Job queue system
│   │   ├── rbac/               # Role-based access control
│   │   ├── services/           # Business logic services
│   │   └── utils/              # Utility functions
│   └── frontend/               # Nuxt.js frontend application
│       └── src/
│           ├── api/            # API client
│           ├── components/     # Vue components
│           ├── layouts/        # Page layouts
│           ├── pages/          # Application pages
│           ├── plugins/        # Nuxt plugins
│           └── store/          # Vuex store
├── build.sh                    # Build script for production
├── setup-mac.sh                # macOS development setup script
├── start-dev.sh                # Start development servers
├── docker-compose.yml          # Docker Compose configuration
└── render.yaml                 # Render deployment configuration
```

## Development Scripts

### setup-mac.sh
Installs all required dependencies for macOS development:
- Homebrew (if not installed)
- PostgreSQL 16
- Redis
- Go 1.24+
- Node.js 14 and Yarn
- MailHog (for email testing)

Creates database, runs migrations, and installs all dependencies.

**Usage:**
```bash
./setup-mac.sh
```

### start-dev.sh
Starts both backend and frontend development servers:
- Ensures PostgreSQL, Redis, and MailHog are running
- Starts backend on http://localhost:8081
- Starts frontend on http://localhost:8080
- Logs output to `logs/backend.log` and `logs/frontend.log`
- Press Ctrl+C to stop all servers

**Usage:**
```bash
./start-dev.sh
```

### build.sh
Builds production-ready artifacts:
- Generates static frontend files
- Builds statically-linked backend binary
- Runs all tests
- Creates production build directory

**Usage:**
```bash
./build.sh
```

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

## API Documentation

The backend provides a RESTful JSON API following the [JSON API specification](https://jsonapi.org/).

### Base URL
- Development: `http://localhost:8081/api/v1`
- Production: `https://your-app.onrender.com/api/v1`

### Authentication
The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

### Key Endpoints
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/logout` - User logout
- `GET /api/v1/users/me` - Get current user profile
- `PUT /api/v1/users/me` - Update user profile
- `POST /api/v1/auth/2fa/enable` - Enable 2FA
- `POST /api/v1/auth/verify-mobile` - Verify mobile number
- `GET /api/v1/subscriptions` - List subscriptions
- `POST /api/v1/subscriptions` - Create subscription

For complete API documentation, explore the [controllers](./src/backend/controllers) directory.

## Security Features

### Authentication & Authorization
- **JWT-based authentication** with secure token storage
- **Role-based access control (RBAC)** for fine-grained permissions
- **Social OAuth** integration (GitHub, Google, Facebook, LinkedIn)
- **Two-factor authentication (2FA)** using TOTP (Time-based One-Time Passwords)
- **SMS verification** for mobile confirmation (via Twilio)

### Security Best Practices
- **Password hashing** using bcrypt
- **HMAC signing** for sensitive operations
- **Rate limiting** to prevent abuse
- **CORS protection** with configurable origins
- **SQL injection protection** via parameterized queries
- **XSS protection** in frontend rendering
- **Secure session management**
- **Audit logging** for sensitive operations

### Configuration Security
- Secrets stored in separate `.secret` files (never committed to git)
- Production secrets injected via CI/CD environment variables
- Database credentials use environment variables
- API keys encrypted in transit and at rest

## Tech Stack

### Backend
- **Language:** Go 1.24+
- **Framework:** Beego
- **Database:** PostgreSQL 12+
- **Cache:** Redis
- **Queue:** Custom queue management system
- **Authentication:** JWT, OAuth 2.0
- **Email:** Mailgun SMTP
- **SMS:** Twilio
- **Payments:** Stripe API

### Frontend
- **Framework:** Nuxt.js 2.x (Vue.js)
- **UI Library:** Vuetify (Material Design)
- **State Management:** Vuex
- **HTTP Client:** Axios
- **Testing:** Jest
- **Build:** Webpack (via Nuxt)

### DevOps
- **CI/CD:** GitHub Actions
- **Containerization:** Docker
- **Hosting:** Render
- **Version Control:** Git (GitHub/GitLab mirrors)
- **Database Migrations:** sql-migrate

## Performance Optimization

### Caching Strategy
- **Multiple cache backends:** Memory, Redis, Database, File
- **Configurable cache type** via environment variable
- **Automatic cache invalidation** on data updates
- **Query result caching** for expensive operations
- **Pagination with caching** for large datasets

### Frontend Optimization
- **Static site generation** for improved SEO and performance
- **Lazy loading** of components and routes
- **Code splitting** for smaller bundle sizes
- **Image optimization** and lazy loading
- **Service Worker** for offline support (PWA ready)

### Backend Optimization
- **Connection pooling** for database and Redis
- **Goroutines** for concurrent operations
- **Efficient JSON serialization** using standard library
- **Database query optimization** with proper indexing
- **Static binary compilation** for faster startup

## Contributing

We welcome contributions! Here's how you can help:

### Reporting Issues
- Use the [GitHub Issues](https://github.com/gadelkareem/skeleton/issues) page
- Check if the issue already exists before creating a new one
- Include detailed steps to reproduce the problem
- Provide environment information (OS, Go version, Node version)

### Pull Requests
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature-name`
3. Make your changes
4. Run tests: `go test ./...` and `yarn test`
5. Commit your changes with clear messages
6. Push to your fork: `git push origin feature/your-feature-name`
7. Create a Pull Request

### Development Guidelines
- Follow Go best practices and idiomatic code style
- Use `gofmt` for Go code formatting
- Write tests for new features
- Update documentation as needed
- Keep commits atomic and well-described
- Ensure all tests pass before submitting PR

## License

This project is open source and available under the [MIT License](LICENSE).

## Support

- **Documentation:** This README and inline code comments
- **Issues:** [GitHub Issues](https://github.com/gadelkareem/skeleton/issues)
- **Discussions:** [GitHub Discussions](https://github.com/gadelkareem/skeleton/discussions)

## Credits

Built with:
- [Beego](https://beego.me/) - Go web framework
- [Nuxt.js](https://nuxtjs.org/) - Vue.js framework
- [Vuetify](https://vuetifyjs.com/) - Material Design components
- [Stripe](https://stripe.com/) - Payment processing
- [Render](https://render.com/) - Cloud hosting

## Changelog

See [GitHub Releases](https://github.com/gadelkareem/skeleton/releases) for version history and changes.

---

**Star this repository if you find it useful! ⭐**
