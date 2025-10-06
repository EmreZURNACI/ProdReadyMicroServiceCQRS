
# ProdReadyMicroServiceCQRS

## CQRS-Based Microservice Architecture - Auth & User Services

This project is built using the CQRS (Command Query Responsibility Segregation) architectural pattern and consists of two independent microservices: `Auth` and `User`. Each service is fully isolated at both the application and infrastructure levels.

### Key Architectural Decisions:
- **Command operations (writes)** are handled using **PostgreSQL**.
- **Query operations (reads)** are served from **MongoDB**.
- Data synchronization between these two databases is managed via **Confluent Kafka Connect**.

### Data Streaming with Connectors:
- **Source Connector**: Captures changes from PostgreSQL using the **Debezium PostgreSQL Connector**.
- **Sink Connector**: Propagates these changes into MongoDB using the **MongoDB Sink Connector**.

This setup allows for high scalability, separation of concerns, and improved read performance without impacting write operations.

## Architecture Overview

This project follows the **CQRS** pattern, which separates **write** and **read** operations into different data models and storage mechanisms.

### 🔹 Command Side (Write)
- All write operations (e.g., user registration, updates, deletions) are performed via PostgreSQL.
- Each microservice maintains its own isolated command-side data schema.

### 🔹 Query Side (Read)
- All read operations (e.g., fetching user data) are handled via MongoDB.
- MongoDB is kept in sync with PostgreSQL through **Change Data Capture (CDC)** mechanisms.

### 🔄 Data Synchronization via Confluent Kafka Connect

| Layer               | Description                                                   |
|---------------------|---------------------------------------------------------------|
| **Kafka Connect**   | Manages the data pipeline between source and sink             |
| **Source Connector**| Uses Debezium to capture row-level changes from PostgreSQL    |
| **Sink Connector**  | Writes those changes to MongoDB in near real-time             |

By decoupling reads and writes, the system achieves better performance, scalability, and resilience.

