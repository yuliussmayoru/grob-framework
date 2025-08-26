
# Grob Framework
## Overview

Grob created based on grob company needs that the entire web apps backend were served from one repository. this framework created to serve better security, high performance and fast development wether for creating mvp or creating production ready api service.

Grob is a powerful and scalable Go framework inspired by the modular architecture of Nest.js. It is designed to build efficient and reliable server-side applications, with first-class support for managing multiple web applications within a single monorepo.

The framework leverages a robust dependency injection system powered by [dig](go.uber.org/dig) and a high-performance HTTP router from [gin](github.com/gin-gonic/gin). This combination provides a solid foundation for creating maintainable and extensible services.

## Features
* Modular Architecture: Organize your code into reusable modules, each with its own controllers, services, and dependencies.

* Multi-App Support: Develop and manage multiple distinct web applications within a single project structure.

* Powerful CLI: Comes with the grob CLI, a command-line tool for scaffolding applications and modules, automating boilerplate code generation.

* Dependency Injection: Built-in dependency injection container makes managing dependencies simple and clean.

* Scalable by Design: The architecture is designed to grow with your needs, from a simple proof-of-concept to a large-scale enterprise system.

* Database Agnostic: Capable of handling multiple database connections, allowing you to work with different databases across your applications.

## Getting Started

### The grob CLI
The grob CLI is the primary tool for scaffolding and managing your Grob projects.
1. Create a New Project
To start, create a new Grob project. This command will create a new directory with the complete project structure and all necessary files.

```
$ grob new my-grob-project
```
This will create a folder named my-grob-project.

2. Create a New Application api
Navigate into your new project directory and use the create-app command to scaffold a new web application.

```
$ cd my-grob-project
$ grob create-app webapp1
```
This generates a new application directory inside the internal/ folder, complete with a main entry point file.

3. To add a new feature to an application, use the create-module command. This creates a folder for your module containing the controller, service, and module files.
```
grob create-module webapp1 auth
```
This command creates an auth module within webapp1, providing a clean structure for your authentication-related logic.

## Project Structure
When you create a new project, Grob generates the following structure to keep your applications and modules clearly separated.
```
grob-framework/
├── cmd/
│   └── grob/
│       └── main.go         # Entry point for the grob CLI
├── internal/
│   ├── webapp1/
│   │   ├── auth/
│   │   │   ├── auth.controller.go
│   │   │   ├── auth.module.go
│   │   │   └── auth.service.go
│   │   └── webapp1_main.go
│   ├── webapp2/
│   │   └── webapp2_main.go
│   └── main.go             # Centralized entry point for all web apps
├── pkg/
│   └── framework/          # Core framework cod
│       ├── core.go
│       └── module.go
├── go.mod
└── go.sum
```
## Architecture concept
### Modules
Modules are the core building blocks of a Grob application. Each module encapsulates a specific feature or domain of your application. A module is defined by a Module interface and is responsible for registering its components (controllers, services) with the dependency injection container.
```
// pkg/framework/module.go
package framework

import "go.uber.org/dig"

type Module interface {
	Register(container *dig.Container) error
}
```
### Dependency injection
The framework uses [dig](go.uber.org/dig) to manage dependencies. When you create a module, you provide its services and controllers to the container, which then handles instantiation and injection where needed.
#### Application Entry point
The internal/main.go file serves as the centralized runner for all your web applications. It discovers and runs every application defined within the internal directory.
```
// internal/main.go
package main

import (
    "log"
    "sync"

    "my-grob-project/internal/webapp1"
    "my-grob-project/internal/webapp2"
    // Import other web apps here
)

// AppRunner defines the interface for a runnable application.
type AppRunner interface {
    Run()
}

func main() {
    // A map of application names to their runner functions.
    // This would ideally be populated automatically by the CLI.
    apps := map[string]AppRunner{
        "webapp1": webapp1.App{}, // Assuming webapp1.App implements AppRunner
        "webapp2": webapp2.App{}, // Assuming webapp2.App implements AppRunner
    }

    var wg sync.WaitGroup

    for name, app := range apps {
        wg.Add(1)
        
        // Launch each application in its own goroutine.
        go func(appName string, runner AppRunner) {
            defer wg.Done()
            log.Printf("Starting application: %s", appName)
            runner.Run() // This Run method will start the Gin server on a specific port.
        }(name, app)
    }

    log.Println("All applications are starting...")
    wg.Wait() // Wait for all applications to finish.
    log.Println("All applications have been shut down.")
}
```
To run all applications, you simply execute:
```
$ grob run
```
### Why a Separate Port for Each App?
Running each application on its own unique port is a core principle of microservices architecture that provides several key benefits:
```
* ✅ No Port Conflicts: The most basic reason. Only one process can listen on a port at a time. Assigning unique ports ensures all your applications can run simultaneously without errors.
* 🛡️ Isolation & Fault Tolerance: If one application crashes, it doesn't affect the others. The failure is contained, making your overall system more resilient and reliable.
* ⚖️ Independent Scaling: You can scale each application independently based on its specific load. If your auth service is under heavy use, you can allocate more resources to it without scaling less-used services.
* 🚀 Independent Deployment: You can update, restart, or deploy a single application without taking the entire system offline, which is essential for continuous integration and deployment (CI/CD).
```
This port configuration can be managed via environment variables or a central configuration file.

## Routing
Routing is handled by [gin](gin-gonic/gin). The framework's structure is designed to map directly to your API endpoints. For example, a login method in the AuthController of the auth module within webapp1 would correspond to the following endpoint:
```
/api/webapp1/auth/login
```
This is achieved by creating nested router groups that mirror your application and module structure, ensuring your API endpoints are consistent and predictable.