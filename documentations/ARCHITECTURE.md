# Architecture

This document describes the architectural structure of the project, following the principles of Clean Architecture in Golang. The goal is to maintain separation of concerns, ensure testability, and allow flexibility for future changes.

## Folder Structure

```unset
/api                 → API layer (HTTP handlers, controllers)
  ├── routes/        → Routes to the REST APIs
  ├── utils/         → Utilities for the APIs (API helpers, middlewares, etc.)
/config              → Configuration files (environment variables, settings)
/features            → Business logic (use cases, services)
/documentations      → Necessary README files for the understanding of this project
/infrastructure      → External dependencies (databases, caching, external APIs, repositories)
  ├── cache/         → Cache handling (e.g., Redis)
  ├── database/      → Database connections and migrations
  ├── repositories/  → Data access layer (repositories for various entities)
  ├── external/      → External service integrations (e.g As-Sunnah Foundation Hadith SDK)
/models              → Entity definitions (domain models)
/scripts            → Utility scripts (migrations, database seeding, etc.)
/shared             → Common utilities (constants, enums, error messages, etc.)
```
