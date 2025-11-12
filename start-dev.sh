#!/bin/bash

# Skeleton Project - Development Server Startup Script
# This script starts both backend and frontend development servers

set -e

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

# Check if services are running
check_service() {
    local service_name=$1
    local port=$2

    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null ; then
        return 0
    else
        return 1
    fi
}

echo ""
echo "╔════════════════════════════════════════════════════════════════╗"
echo "║         Skeleton Project - Development Server                 ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Verify prerequisites
print_step "Checking prerequisites..."

if ! command_exists go; then
    print_error "Go is not installed. Please run ./setup-mac.sh first"
    exit 1
fi

if ! command_exists node; then
    print_error "Node.js is not installed. Please run ./setup-mac.sh first"
    exit 1
fi

if ! command_exists yarn; then
    print_error "Yarn is not installed. Please run ./setup-mac.sh first"
    exit 1
fi

# Check PostgreSQL
if ! check_service "PostgreSQL" 5432; then
    print_warning "PostgreSQL is not running. Starting it now..."
    brew services start postgresql@16
    sleep 3
fi
print_success "PostgreSQL is running"

# Check Redis
if ! check_service "Redis" 6379; then
    print_warning "Redis is not running. Starting it now..."
    brew services start redis
    sleep 2
fi
print_success "Redis is running"

# Check MailHog
if ! check_service "MailHog" 1025; then
    print_warning "MailHog is not running. Starting it now..."
    brew services start mailhog
    sleep 2
fi
print_success "MailHog is running"

# Check if ports 8080 and 8081 are available
if check_service "Backend" 8081; then
    print_error "Port 8081 is already in use. Please stop the running process."
    echo "Run: lsof -ti:8081 | xargs kill -9"
    exit 1
fi

if check_service "Frontend" 8080; then
    print_error "Port 8080 is already in use. Please stop the running process."
    echo "Run: lsof -ti:8080 | xargs kill -9"
    exit 1
fi

# Function to cleanup background processes on exit
cleanup() {
    echo ""
    print_step "Shutting down servers..."

    if [ ! -z "$BACKEND_PID" ]; then
        kill $BACKEND_PID 2>/dev/null || true
    fi

    if [ ! -z "$FRONTEND_PID" ]; then
        kill $FRONTEND_PID 2>/dev/null || true
    fi

    # Kill any remaining processes on these ports
    lsof -ti:8081 | xargs kill -9 2>/dev/null || true
    lsof -ti:8080 | xargs kill -9 2>/dev/null || true

    print_success "Servers stopped"
    exit 0
}

# Register cleanup function
trap cleanup SIGINT SIGTERM

# Create log directory
mkdir -p logs

# Start Backend
print_step "Starting Backend server..."
cd src/backend

# Set environment variables
export DATABASE_URL="postgres://skeleton_backend:dev_awTf9d2GceKRNzhkCb4H5B8nfmq@localhost/skeleton_backend?sslmode=disable"
export REDIS_URL="redis://localhost:6379"
export BEEGO_RUNMODE=dev
export CACHE_TYPE=redis
export HTTP_ADDR=localhost
export MAIL_HOST=localhost

# Build if binary doesn't exist
if [ ! -f "skeleton" ]; then
    print_step "Building backend binary..."
    go build -o skeleton
fi

# Start backend in background
./skeleton > ../../logs/backend.log 2>&1 &
BACKEND_PID=$!

cd "$SCRIPT_DIR"

# Wait for backend to start
print_step "Waiting for backend to start..."
for i in {1..30}; do
    if check_service "Backend" 8081; then
        print_success "Backend started (PID: $BACKEND_PID)"
        print_success "Backend API: http://localhost:8081"
        break
    fi
    sleep 1
    if [ $i -eq 30 ]; then
        print_error "Backend failed to start. Check logs/backend.log for details"
        cat logs/backend.log
        cleanup
    fi
done

# Start Frontend
print_step "Starting Frontend server..."
cd src/frontend

# Start frontend in background
yarn serve > ../../logs/frontend.log 2>&1 &
FRONTEND_PID=$!

cd "$SCRIPT_DIR"

# Wait for frontend to start
print_step "Waiting for frontend to start..."
for i in {1..60}; do
    if check_service "Frontend" 8080; then
        print_success "Frontend started (PID: $FRONTEND_PID)"
        print_success "Frontend app: http://localhost:8080"
        break
    fi
    sleep 1
    if [ $i -eq 60 ]; then
        print_error "Frontend failed to start. Check logs/frontend.log for details"
        cat logs/frontend.log
        cleanup
    fi
done

echo ""
echo "╔════════════════════════════════════════════════════════════════╗"
echo "║              All Servers Running! 🚀                           ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""
echo "Services:"
echo "  • Frontend:      http://localhost:8080"
echo "  • Backend API:   http://localhost:8081/api/v1"
echo "  • MailHog UI:    http://localhost:8025"
echo "  • PostgreSQL:    localhost:5432"
echo "  • Redis:         localhost:6379"
echo ""
echo "Logs:"
echo "  • Backend:       tail -f logs/backend.log"
echo "  • Frontend:      tail -f logs/frontend.log"
echo ""
print_warning "Press Ctrl+C to stop all servers"
echo ""

# Monitor logs in real-time (optional - comment out if you prefer silent mode)
tail -f logs/backend.log logs/frontend.log 2>/dev/null &
TAIL_PID=$!

# Wait forever (until Ctrl+C)
wait
