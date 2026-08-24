# Go URL Shortener --- MVP Requirements

## Goal

Build a standalone URL shortener service in Go.

The main objective is to practice backend development and reinforce Go
fundamentals without relying on generated implementation code.

## Core Requirements

### 1. Create a short URL

Implement an endpoint:

`POST /urls`

It should:

-   Accept a destination URL.
-   Validate the input.
-   Generate a unique short code.
-   Store the short code and destination URL.
-   Return the generated short URL.

Example request:

``` json
{
  "url": "https://example.com/some/very/long/url"
}
```

### 2. Redirect

Implement an endpoint:

`GET /{code}`

It should:

-   Find the URL associated with the short code.
-   Redirect the client to the original URL.
-   Return an appropriate error when the code does not exist.

### 3. PostgreSQL

Use PostgreSQL for persistence.

Store at minimum:

-   ID
-   Short code
-   Destination URL
-   Created timestamp
-   Expiration timestamp (optional)
-   Click count

### 4. sqlc

Use `sqlc` for database queries.

Write the SQL queries yourself and generate the Go database layer using
sqlc.

### 5. Click Counter

Every successful redirect should increment the link's click count.

Add an endpoint for retrieving information about a shortened URL:

`GET /urls/{code}`

Return at minimum:

-   Short code
-   Destination URL
-   Click count
-   Created timestamp

### 6. URL Expiration

Allow a URL to optionally expire.

An expired short URL must no longer redirect to its destination.

### 7. Configuration

Read configuration from environment variables.

Examples:

-   Server port
-   PostgreSQL connection string
-   Base URL
-   Redis address

Provide a `.env.example`.

### 8. Graceful Shutdown

Use Go's `context` package to gracefully stop the HTTP server and close
resources when the application receives a shutdown signal.

### 9. Docker

Provide a Docker setup that allows the project and its dependencies to
run locally.

Use Docker Compose for infrastructure such as PostgreSQL and Redis.

------------------------------------------------------------------------

## Phase 2 --- Redis Cache

After the PostgreSQL-only MVP works, add Redis.

For redirects:

1.  Look for the short code in Redis.
2.  If found, use the cached destination.
3.  If not found, retrieve it from PostgreSQL.
4.  Store the result in Redis for future requests.

Use a TTL for cached entries.

The application must still work correctly when a value is not present in
the cache.

------------------------------------------------------------------------

## Error Handling

Handle at minimum:

-   Invalid request body
-   Invalid URL
-   Unknown short code
-   Expired URL
-   Database errors
-   Cache errors
-   Duplicate/generated-code collisions

Do not expose internal implementation details in HTTP error responses.

------------------------------------------------------------------------

## Testing

Add tests for important application behavior.

At minimum consider:

-   URL validation
-   Short-code generation
-   Creating a URL
-   Looking up an existing URL
-   Looking up an unknown URL
-   Expiration behavior

------------------------------------------------------------------------

## Concepts to Practice

This project should reinforce:

-   Go packages
-   Structs
-   Interfaces where they provide a real benefit
-   Error handling
-   `context.Context`
-   HTTP handlers
-   JSON encoding/decoding
-   PostgreSQL
-   SQL
-   sqlc
-   Environment configuration
-   Graceful shutdown
-   Docker
-   Redis
-   Caching
-   TTL
-   Cache misses
-   Database/cache interaction
-   Basic API design

## Constraint

Build the implementation yourself.

AI assistance should primarily be used for:

-   Explaining concepts
-   Go syntax questions
-   Documentation questions
-   Hints when blocked
-   Code review after you implement something

Avoid requesting complete implementations unless you explicitly want to
inspect a reference solution.

## Definition of Done

The project is complete when:

-   A URL can be shortened.
-   The generated short URL redirects correctly.
-   Data persists in PostgreSQL.
-   Clicks are counted.
-   Expired URLs stop working.
-   Configuration comes from environment variables.
-   The service shuts down gracefully.
-   PostgreSQL and Redis can run through Docker Compose.
-   Redis caches URL lookups.
-   Important behavior has tests.
-   A README explains how to run and use the project.

## Optional Extensions

Only after the MVP is complete:

-   Custom short codes
-   Rate limiting
-   Link deletion
-   Link analytics
-   QR-code generation
-   Authentication
-   Per-user links
-   Cache invalidation strategy
-   Benchmarking the API with your Go Load Tester
