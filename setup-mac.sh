#!/bin/bash

# Skeleton Project - macOS Setup Script
# This script installs all dependencies and sets up the development environment

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper functions
print_step() {
    echo -e "${BLUE}==>${NC} ${1}"
}

print_success() {
    echo -e "${GREEN}✓${NC} ${1}"
}

print_error() {
    echo -e "${RED}✗${NC} ${1}"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} ${1}"
}

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Get script directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPT_DIR"

echo ""
echo "╔════════════════════════════════════════════════════════════════╗"
echo "║         Skeleton Project - macOS Setup Script                 ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Step 1: Install Homebrew
print_step "Step 1: Checking Homebrew installation"
if ! command_exists brew; then
    print_warning "Homebrew not found. Installing Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

    # Add Homebrew to PATH for Apple Silicon Macs
    if [[ $(uname -m) == 'arm64' ]]; then
        echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile
        eval "$(/opt/homebrew/bin/brew shellenv)"
    fi
    print_success "Homebrew installed"
else
    print_success "Homebrew already installed"
fi

# Update Homebrew
print_step "Updating Homebrew..."
brew update

# Step 2: Install PostgreSQL
print_step "Step 2: Installing PostgreSQL"
if ! command_exists psql; then
    brew install postgresql@16
    print_success "PostgreSQL installed"
else
    print_success "PostgreSQL already installed"
fi

# Start PostgreSQL service
print_step "Starting PostgreSQL service..."
brew services start postgresql@16 || brew services restart postgresql@16
sleep 3
print_success "PostgreSQL service started"

# Step 3: Install Redis
print_step "Step 3: Installing Redis"
if ! command_exists redis-server; then
    brew install redis
    print_success "Redis installed"
else
    print_success "Redis already installed"
fi

# Start Redis service
print_step "Starting Redis service..."
brew services start redis || brew services restart redis
print_success "Redis service started"

# Step 4: Install Go
print_step "Step 4: Installing Go 1.24+"
if ! command_exists go; then
    brew install go
    print_success "Go installed"
else
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    print_success "Go already installed (version: $GO_VERSION)"
fi

# Step 5: Install Node.js and Yarn
print_step "Step 5: Installing Node.js and Yarn"
if ! command_exists node; then
    brew install node@14
    brew link node@14
    print_success "Node.js 14 installed"
else
    NODE_VERSION=$(node --version)
    print_success "Node.js already installed (version: $NODE_VERSION)"
fi

if ! command_exists yarn; then
    brew install yarn
    print_success "Yarn installed"
else
    print_success "Yarn already installed"
fi

# Step 6: Install MailHog (for email testing)
print_step "Step 6: Installing MailHog"
if ! command_exists mailhog; then
    brew install mailhog
    print_success "MailHog installed"
else
    print_success "MailHog already installed"
fi

# Start MailHog service
print_step "Starting MailHog service..."
brew services start mailhog || brew services restart mailhog
print_success "MailHog service started (Web UI: http://localhost:8025)"

# Step 7: Setup Database
print_step "Step 7: Setting up PostgreSQL database"
export PGHOST=localhost
export PGUSER=postgres

# Check if database already exists
if psql -lqt | cut -d \| -f 1 | grep -qw skeleton_backend; then
    print_warning "Database 'skeleton_backend' already exists"
    read -p "Do you want to recreate it? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        print_step "Dropping existing database..."
        psql -c "DROP DATABASE IF EXISTS skeleton_backend;" || true
        psql -c "DROP USER IF EXISTS skeleton_backend;" || true
        bash ./docker/db/db.sh
        print_success "Database recreated"
    fi
else
    bash ./docker/db/db.sh
    print_success "Database created"
fi

# Step 8: Setup Backend
print_step "Step 8: Setting up Backend"

# Create secret config if it doesn't exist
if [ ! -f "src/backend/conf/app.dev.ini.secret" ]; then
    print_warning "Creating app.dev.ini.secret from example..."
    cp src/backend/conf/app.dev.ini.secret.example src/backend/conf/app.dev.ini.secret
    print_success "Created app.dev.ini.secret (you may need to configure API keys later)"
fi

# Install Go dependencies
print_step "Installing Go dependencies..."
cd src/backend
go mod download
print_success "Go dependencies installed"

# Build backend
print_step "Building backend binary..."
go build -o skeleton
print_success "Backend binary built"

# Run migrations
print_step "Running database migrations..."
export DATABASE_URL="postgres://skeleton_backend:dev_awTf9d2GceKRNzhkCb4H5B8nfmq@localhost/skeleton_backend?sslmode=disable"
export BEEGO_RUNMODE=dev
export CACHE_TYPE=memory
./skeleton migrate up
print_success "Database migrations completed"

cd "$SCRIPT_DIR"

# Step 9: Setup Frontend
print_step "Step 9: Setting up Frontend"
cd src/frontend

print_step "Installing frontend dependencies (this may take a few minutes)..."
yarn install --frozen-lockfile
print_success "Frontend dependencies installed"

cd "$SCRIPT_DIR"

# Done!
echo ""
echo "╔════════════════════════════════════════════════════════════════╗"
echo "║                   Setup Complete! 🎉                           ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""
print_success "All dependencies installed and services configured"
echo ""
echo "Services running:"
echo "  • PostgreSQL: localhost:5432"
echo "  • Redis: localhost:6379"
echo "  • MailHog Web UI: http://localhost:8025"
echo ""
echo "Next steps:"
echo "  1. Run './start-dev.sh' to start the backend and frontend servers"
echo "  2. Backend will be available at: http://localhost:8081"
echo "  3. Frontend will be available at: http://localhost:8080"
echo ""
print_warning "Note: Edit src/backend/conf/app.dev.ini.secret to configure:"
echo "  • Social authentication (GitHub, Google, Facebook, LinkedIn)"
echo "  • SMS service (Twilio)"
echo "  • Other API keys as needed"
echo ""
