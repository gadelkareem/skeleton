# macOS Development Setup Guide

This guide will help you set up the Skeleton project for development on macOS.

## Prerequisites

- macOS 10.14 (Mojave) or later
- Command Line Tools (will be installed with Homebrew if not present)

## Quick Start

### 1. Initial Setup (First Time Only)

Run the setup script to install all dependencies:

```bash
chmod +x setup-mac.sh
./setup-mac.sh
```

This script will:
- Install Homebrew (if not already installed)
- Install PostgreSQL 16
- Install Redis
- Install Go 1.24+
- Install Node.js 14
- Install Yarn
- Install MailHog (email testing)
- Create and initialize the PostgreSQL database
- Run database migrations
- Install Go dependencies
- Install frontend dependencies
- Create configuration files

**Note:** The script will prompt you if the database already exists.

### 2. Start Development Servers

After setup is complete, start the development servers:

```bash
chmod +x start-dev.sh
./start-dev.sh
```

This will start:
- **Backend API** on `http://localhost:8081`
- **Frontend App** on `http://localhost:8080`
- **MailHog UI** on `http://localhost:8025` (for testing emails)

Press `Ctrl+C` to stop all servers.

## Services

The following services will be installed and configured:

| Service      | Port | URL                        | Purpose                    |
|--------------|------|----------------------------|----------------------------|
| Frontend     | 8080 | http://localhost:8080      | Nuxt.js web application    |
| Backend API  | 8081 | http://localhost:8081/api/v1 | Go/Beego REST API        |
| PostgreSQL   | 5432 | localhost:5432             | Database                   |
| Redis        | 6379 | localhost:6379             | Cache & session store      |
| MailHog SMTP | 1025 | localhost:1025             | Email testing (SMTP)       |
| MailHog UI   | 8025 | http://localhost:8025      | Email testing (Web UI)     |

## Configuration

### Backend Configuration

Backend configuration files are located in `src/backend/conf/`:

- `app.default.ini` - Default configuration
- `app.dev.ini` - Development environment settings
- `app.dev.ini.secret` - Secret keys and API credentials (not in git)

After running `setup-mac.sh`, edit `src/backend/conf/app.dev.ini.secret` to configure:

```ini
# Social Authentication
[social]
githubClientID = your_github_client_id
githubClientSecret = your_github_secret
googleClientID = your_google_client_id
googleClientSecret = your_google_secret
# ... etc

# SMS Service (Twilio)
[sms]
accountSID = your_twilio_account_sid
authToken = your_twilio_auth_token
ownNumber = your_twilio_phone_number
```

### Frontend Configuration

Frontend configuration is in `src/frontend/src/nuxt.config.js`.

## Manual Commands

If you prefer to run commands manually instead of using `start-dev.sh`:

### Backend

```bash
cd src/backend
export DATABASE_URL="postgres://skeleton_backend:dev_awTf9d2GceKRNzhkCb4H5B8nfmq@localhost/skeleton_backend?sslmode=disable"
export REDIS_URL="redis://localhost:6379"
export BEEGO_RUNMODE=dev
export CACHE_TYPE=redis
go run main.go
```

### Frontend

```bash
cd src/frontend
yarn serve
```

## Database Management

### Run Migrations

```bash
cd src/backend
export DATABASE_URL="postgres://skeleton_backend:dev_awTf9d2GceKRNzhkCb4H5B8nfmq@localhost/skeleton_backend?sslmode=disable"
export BEEGO_RUNMODE=dev
./skeleton migrate up
```

### Rollback Migrations

```bash
cd src/backend
./skeleton migrate down
```

### Reset Database

```bash
# Stop the dev server first (Ctrl+C if running)
psql -U postgres -c "DROP DATABASE skeleton_backend;"
psql -U postgres -c "DROP USER skeleton_backend;"
bash ./docker/db/db.sh
cd src/backend
./skeleton migrate up
```

## Managing Services

### Start/Stop Services

```bash
# PostgreSQL
brew services start postgresql@16
brew services stop postgresql@16

# Redis
brew services start redis
brew services stop redis

# MailHog
brew services start mailhog
brew services stop mailhog
```

### Check Service Status

```bash
brew services list
```

## Testing

### Backend Tests

```bash
cd src/backend
export BEEGO_RUNMODE=test
export DATABASE_URL="postgres://skeleton_backend:dev_awTf9d2GceKRNzhkCb4H5B8nfmq@localhost/skeleton_backend?sslmode=disable"
export CACHE_TYPE=memory
go test ./... -v
```

### Frontend Tests

```bash
cd src/frontend
yarn test
```

## Troubleshooting

### Port Already in Use

If you see "port already in use" errors:

```bash
# Find and kill process on port 8080 (frontend)
lsof -ti:8080 | xargs kill -9

# Find and kill process on port 8081 (backend)
lsof -ti:8081 | xargs kill -9
```

### Database Connection Issues

1. Ensure PostgreSQL is running:
   ```bash
   brew services list | grep postgresql
   ```

2. Check database exists:
   ```bash
   psql -U postgres -l | grep skeleton_backend
   ```

3. Test connection:
   ```bash
   psql -U skeleton_backend -h localhost skeleton_backend
   # Password: dev_awTf9d2GceKRNzhkCb4H5B8nfmq
   ```

### Redis Connection Issues

```bash
# Check if Redis is running
brew services list | grep redis

# Test Redis connection
redis-cli ping
# Should return: PONG
```

### Node/Yarn Issues

If you encounter Node.js version issues:

```bash
# Install nvm (Node Version Manager)
brew install nvm

# Install Node 14
nvm install 14
nvm use 14

# Reinstall dependencies
cd src/frontend
rm -rf node_modules
yarn install
```

### Go Module Issues

```bash
cd src/backend
go clean -modcache
go mod download
go mod tidy
```

## Logs

Development logs are stored in the `logs/` directory:

- `logs/backend.log` - Backend server logs
- `logs/frontend.log` - Frontend server logs

View logs in real-time:

```bash
# Backend logs
tail -f logs/backend.log

# Frontend logs
tail -f logs/frontend.log

# Both
tail -f logs/*.log
```

## Docker Alternative

If you prefer using Docker instead of native services:

```bash
# Start all services with Docker Compose
docker-compose up

# Stop all services
docker-compose down
```

## Useful Resources

- **Backend Framework:** [Beego Documentation](https://beego.vip/)
- **Frontend Framework:** [Nuxt.js Documentation](https://nuxtjs.org/)
- **Go Documentation:** [golang.org](https://golang.org/doc/)
- **Vue.js Documentation:** [vuejs.org](https://vuejs.org/)

## Additional Commands

### Build for Production

```bash
# Backend
cd src/backend
go build -o skeleton

# Frontend
cd src/frontend
yarn generate
```

### Clean Build Artifacts

```bash
# Backend
cd src/backend
rm -f skeleton
go clean

# Frontend
cd src/frontend
rm -rf .nuxt dist node_modules/.cache
```

## Next Steps

1. Visit `http://localhost:8080` to see the frontend
2. Check `http://localhost:8081/api/v1/health` for backend health check
3. View emails at `http://localhost:8025` (MailHog)
4. Start coding! 🚀
