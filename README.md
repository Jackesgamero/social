# Go Social Media Backend

A REST API built with **Go** and **PostgreSQL**, designed for scalability and deployed on **Google Cloud**.

---

## 📌 Overview
This project is a production-grade social media backend developed with a focus on **system design, performance, and reliability**. It implements a composable architecture and a fully automated CI/CD pipeline, handling everything from low-level networking foundations to cloud-scale deployment.

## 🛠 Tech Stack
*   **Language:** Go (Standard Library focus)
*   **Database:** PostgreSQL
*   **Infrastructure:** Google Cloud (GCP)
*   **DevOps:** GitHub Actions (CI/CD), Docker
*   **Tools:** Redis, Gin

## 🏗 Key Engineering Concepts
*   **Composable Architecture:** Modular design utilizing Go interfaces to ensure the system is maintainable, testable, and loosely coupled.
*   **Advanced Routing & Middleware:** Custom-built request handling flow, including rate limiting, authentication, and structured logging.
*   **Database Design:** Optimized PostgreSQL schema with efficient connection pooling and relational integrity.
*   **Production-Ready CI/CD:** Automated testing and deployment workflows to Google Cloud, ensuring high availability and seamless updates.

## 📂 Project Structure

```text
├── cmd/
│   ├── api/            # API entry point & dependency injection
│   └── migrate/        # Database migration CLI tool
├── internal/           # Private application logic
│   ├── auth/           # JWT implementation & Auth providers
│   ├── db/             # Database connection & seeding logic
│   ├── env/            # Environment configuration management
│   ├── mailer/         # Email services (SendGrid/Mailtrap) & templates
│   ├── ratelimiter/    # Fixed-window rate limiting implementation
│   └── store/          # Data Layer: PostgreSQL & Redis (Cache) logic
├── web/                # Static assets, frontend components or templates
├── docs/               # Swagger / OpenAPI documentation (YAML/JSON)
├── scripts/            # Infrastructure & DB initialization scripts
├── bin/                # Compiled binaries & build logs
├── Dockerfile          # Production container configuration
├── docker-compose.yml  # Local dev environment orchestration
├── Makefile            # Automation (build, migrate, test)
└── go.mod              # Go module definition
