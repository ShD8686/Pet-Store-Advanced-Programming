# Pet-Store-Advanced-Programming
This is the first milestone for the Advanced Programming 1 final project. 
We are building a comprehensive Pet Store management system using a monolithic architecture in Go.

## Team Members
*   **[Nurzhaina Kuralbay]** - Project Manager (Proposal & Planning)
*   **[Araizhan Tazhimova]** - System Architect (Diagrams & Design)
*   **[Damir Sheneussizov]** - Backend Developer (Go Implementation & Setup)

## 1. Project Proposal

## Relevance 
In 2026, pet owners want an easy and convenient way to buy pet supplies and adopt pets online. Our system helps pet stores manage their inventory automatically and allows customers to browse and order pets and products from home. It also connects local shelters with people who want to adopt animals.

##  Project Overview
The Pet Store API provides a robust foundation for managing pets, users, and orders. 
This milestone focuses on the core architecture, documentation, and the initial code skeleton.

### Competitor Analysis
*   **PetSmart / Petco:** Large-scale competitors with complex systems. Our solution is more lightweight, focusing on high performance and a simplified API for mobile integration.
*   **Local Pet Stores:** Many lack digital presence. Our project provides an affordable, ready-to-use "Monolith-as-a-Service" architecture.

### Planned Features
- **Pet Catalog:** Browse pets by category and price (Implemented in Skeleton).
- **User Management:** Track customer profiles and admin roles.
- **Order System:** Process purchases and track order history.
- **Authentication:** Secure login for users and admins (Planned).
- **Admin Dashboard:** Interface for adding/deleting inventory (Planned).

### Target Users
*   **Pet Owners:** People looking to buy food, accessories, or adopt pets.
*   **Shelter Administrators:** Staff managing pet listings and availability.
*   **Veterinarians:** Who can track pet sales history and health records.

## Architecture
The project follows a **Monolithic** structure with a clear separation of concerns using the **Repository Pattern**.

*   **Models:** Defined in `internal/models`, representing our database entities.
*   **Repository:** Uses interfaces in `internal/repository` to decouple data storage logic. Currently implemented as a **Mock Repository** (In-Memory).
*   **Handlers:** HTTP handlers in `internal/handlers` to process API requests.

##  Project Structure
```text
.
├── main.go               # Entry point of the application
├── internal/
│   ├── models/           # Data structures (Pet, User, Order)
│   ├── repository/       # Interfaces and Mock Data Implementation
│   └── handlers/         # HTTP request handlers
├── docs/                 # PDF Documentation & Diagrams
├── go.mod                # Go module file
└── README.md             # Project overview 
``` 

## 3. Project Plan 

Our development is scheduled across 4 key weeks:
*   **Week 7:** Initial Design, ERD, and Repository Setup (Current Status).
*   **Week 8:** API Implementation for Pets and Users.
*   **Week 9:** Database Integration (PostgreSQL) and Order Logic.
*   **Week 10:** Final Polish, Testing, and Deployment.

## The detailed **Gantt Chart** with task assignments is available in the `/docs` folder.

## Tech Stack
Language: Go (Golang)
Framework: Standard Library (net/http)
Format: JSON 

## How to Run
1)Clone the repository.
2)Open terminal in the project root.
3)Run the following command:
go run .
The server will start at http://localhost:8080.

### For branch changing 
1)Clone the repository.
2)Open in Visual Studio Code.
3)add branch. 
4)check with command:git branch.
5)add/change code and save with Ctrl+s.
6)commit command in terminal for git.
7)git push.

### Documentation
All diagrams (Use-Case, ERD, UML, Gantt Chart) and the Project Proposal are located in the docs/ folder.