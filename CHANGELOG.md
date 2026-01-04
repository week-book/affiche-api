# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).
This project does not yet have a stable production release.

---

## [Unreleased]

### Added
- Initial implementation of the Events API.
- POST /events endpoint for creating events.
- Validation for required event fields (`text`, `photo`, `date`).
- PostgreSQL-based event repository.
- Event service layer with business validation.
- Automated tests for handler, service, and repository layers.

### Changed
- POST /events now returns the full Event resource instead of only the generated id.
  The response includes `id`, `photo`, `text`, and `date`.

### Internal
- Introduced HTTP response DTOs to separate API contracts from domain models.
- Improved test coverage for event creation flow.
- Refactored event handler to delegate business logic to the service layer.

---

## [Planned]
- GET /events/{id} endpoint.
- GET /events endpoint for listing events.
- Pagination and filtering for events list.
- API versioning strategy (v1).
