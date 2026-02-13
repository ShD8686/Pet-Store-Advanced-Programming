# DNA — Digital Pet Ecosystem 

**DNA** (from Kazakh *DNA* — "mark" or "stamp") is a modern Full Stack platform designed for comprehensive pet registration and owner services. The project bridges the gap between state-level animal identification and private sector needs, such as veterinary care, marketplaces, and lost-and-found services.

### Team Members
[Nurzhaina Kuralbay] - Project Manager (Proposal & Planning)
[Araizhan Tazhimova] - System Architect (Diagrams & Design)
[Damir Sheneussizov] - Backend Developer (Go Implementation & Setup)

---

##  Tech Stack

### Backend
- **Language:** Go (Golang) — chosen for its performance, type safety, and efficient handling of concurrent processes.
- **Database:** SQLite — a lightweight relational database used via the `modernc.org/sqlite` library (CGO-free), ensuring easy portability.
- **Architecture:** Clean RESTful API using the standard `net/http` library to demonstrate a deep understanding of core web protocols.

### Frontend
- **UI/UX:** HTML5 & Tailwind CSS. Features a "breathable" modern design with rounded corners, high-quality typography (Inter font), and a focus on usability.
- **Client Logic:** Vanilla JavaScript (Fetch API). The frontend behaves like a single-page application (SPA), updating content dynamically without full page reloads.
- **Icons:** Custom SVG icons for a lightweight and crisp interface.

---

##  Key Features

### 1. Advanced Authentication & Role System
- **Smart Role Assignment:** The system automatically determines the user's role during registration. If the email contains the substring `admin` (e.g., `admin@tanba.kz`), the account is granted **Administrator** privileges. All other users are registered as standard **Users**.
- **Session Management:** Uses `localStorage` to persist login states across the application.

### 2. Veterinary Service Management
- **Online Booking:** A dedicated interface to schedule appointments with specific specialists (Therapist, Surgeon, etc.).
- **User Dashboard:** Users can track their appointment history and real-time statuses (`pending`, `confirmed`).

### 3. Marketplace (Pet Shop)
- **Product Catalog:** Items are categorized into Food, Medicine, and Toys.
- **Dynamic Filtering:** Instant category switching without page refreshes.
- **Admin Integration:** Any product added via the Admin Dashboard appears instantly in the shop.

### 4. Lost & Found (Classifieds)
- **Ad Types:** Supports "Lost", "Found", "Sell", and "Give Away".
- **Reward System:** Users can specify reward amounts for missing pets.
- **Moderation:** Admins can delete outdated or inappropriate listings.

### 5. Open Data & Real-Time Statistics
- **Analytics Engine:** Live counters for total registered animals, including breakdown by species (Cats/Dogs) based on the actual database state.

### 6. Community News & Blog
- **Content Delivery:** A news feed for tips, local event announcements, and vaccination updates.

---

##  Administrative Controls (CMS)

Admins have access to a secure `/admin` area allowing them to:
- **Manage Registry:** Remove animals from the identification database.
- **Inventory Management:** Add new products to the shop with pricing and imagery.
- **Moderation:** Clear the platform of spam or invalid classified ads.

---

##  Project Structure

```text
.
├── handlers/           # HTTP Request Logic (Auth, API, Pages)
├── models/             # Data structures (User, Pet, Product, Listing)
├── repository/         # Database persistence layer (SQL queries)
├── web/templates/      # Frontend (HTML pages with Tailwind)
├── main.go             # Application entry point
└── petstore.db         # SQLite database file
```

---

##  Getting Started

1. Clone the repository.
2. Ensure you have **Go 1.18+** installed.
3. Run the application:
   ```bash
   go run main.go
   ```
4. Navigate to `http://localhost:8080`.

**Default Admin Credentials:**
- **Email:** `admin@tanba.kz`
- **Password:** `admin`

