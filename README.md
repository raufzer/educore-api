# EduCore API

EduCore API is a backend service built with Golang using the Gin framework. It provides robust functionality for managing courses, assignments, student enrollments, and collaboration tools, implementing modern software practices such as layered architecture.

## Table of Contents
- [EduCore API](#educore-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Setup and Installation](#setup-and-installation)
    - [Installation Steps](#installation-steps)
  - [Running the Application](#running-the-application)
  - [Project Structure](#project-structure)
  - [Troubleshooting](#troubleshooting)
  - [Contributing](#contributing)
  - [License](#license)

## Features
- RESTful API built with Gin framework
- PostgreSQL database integration
- Course and assignment management
- Student and instructor role-based access

## Technologies Used
- **Programming Language**: Golang
- **Web Framework**: Gin
## Setup and Installation

### Installation Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/educore-api.git
   cd educore-api
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Create a `.env` file in the root directory with required variables

## Running the Application
Start the server:
```bash
go run cmd/server/main.go
```
The server will be accessible at http://localhost:8080.

## Project Structure
```
educore-api/
├── cmd/server
│   └── main.go          # Application entry point
├── internal/
│   ├── bootstrap/       # Application bootstrapping
│   ├── controllers/     # API request handlers  
│   ├── dto/             # Data transfer objects
│   ├── integrations/    # External service integrations
│   ├── middlewares/     # Middleware logic
│   ├── models/          # Domain models
│   ├── repositories/    # Data access layer
│   └── services/        # Business logic
```

## Troubleshooting
- Verify environment variables
- Check Docker and Go versions
- Ensure network ports are available
- Review logs for specific errors

## Contributing
1. Fork the repository
2. Create a feature branch
3. Commit changes
4. Push branch
5. Open pull request

## License
Abd Raouf Zerkhef
zerkhefraouf90@gmail.com
