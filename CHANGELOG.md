# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).
This project does not yet have a stable production release.

---

## [Unreleased]

### Added
- Initial implementation of the Events API.
- POST /events endpoint for creating events.
- GET /events endpoint for listing events.
- Validation for required event fields (`text`, `photo`, `date`).
- PostgreSQL-based event repository for event creation.
- Event service layer with business validation.
- HTTP response DTO (`EventResponse`) to represent API responses.
- Automated tests for handler, service, and repository layers.
- Test helpers for creating events in handler tests.

### Changed
- POST /events now returns the full Event resource instead of only the generated `id`.
  The response includes `id`, `photo`, `text`, and `date`.
- Refactored handlers to return structured JSON responses instead of plain text.
- Improved test structure to better reflect real HTTP request/response flow.

### Internal
- Separated domain models from HTTP-layer DTOs.
- Refactored event handler to delegate business logic to the service layer.
- Improved test coverage for the event creation flow.
- Began migration from `http.ServeMux` to `gorilla/mux` for REST-style routing.
- Introduced UUID-based identifiers for events.

---

## [In Progress]

---

## [Planned]
- Pagination and filtering for events list.
- Proper HTTP error mapping (400 / 404 / 422).
- API versioning strategy (v1).
- OpenAPI / Swagger documentation.
