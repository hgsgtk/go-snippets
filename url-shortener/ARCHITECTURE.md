# Architecture

## Basic requirements

* The service supports only HTTP and HTTPS for the original URL's protocol.
* Users will interact via HTTP APIs not though a web UI.
* Basic API requirements:
    * 1. Post /shorten
        * This API creates a short URL from a long URL.
        * Request body:
            ```json
            {
                "url": "https://www.google.com"
            }
            ```
        * Response:
            * Status code: 201
            * Body:
                ```json
                {
                    "short_url": "https://short.url/123456"
                }
                ```

    * 2. GET /{short_code}
        * This API redirects to the original long URL.
        * Response
            * Status code: 302
            * Redirect to the original long URL.          
* If the same long URL is submitted again, the service should return the existing short code instead of creating a new one.
* The length of the URL is limited to 2048 characters.

## Short code generation

There are a couple of options for short code generation. We'll use random alphanumeric string (Base62) for now.

* Auto-incrementing counter - Uses a counter to generate unique numbers that are encoded into 6-character strings using digits 0-9, A-Z, and a-z
    * e.g., 000001, 000002
    * Pros: Extremely simple to implement
    * Cons:
        * The service scale is guessed easily.
* Hashing with Base62 encoding
    * e.g., MD5, SHA-1, SHA-256, SHA-512 -> Base62 encoding -> 6-character string
    * Pros:
        * Same URL will always generate the same short code.
    * Cons:
        * Collision risk
* Random alphanumeric string (Base62)
    * e.g., 123456, 789012
    * Pros:
        * Simple
    * Cons:
        * Checking for uniqueness
* Timestamp-based string
    * Epoc + random string
* UUID - Universally Unique Identifier
* NanoID - a tiny UUID
* ULID - Universally Unique Lexicographically Sortable Identifier

## Expiration

* The expiration process is handled by a separate thread.
* The expiration time is decided by the TTL (time-to-live) interval given by the CLI option or user's input.
* Even a stored URL is accessed, the expiration time is not extended.

## System Components

### PoC Stage

In the PoC stage, we'll use a single server that stores all the data in memory.

#### In-memory key-value database

We assume the data is stored in a in-memory key-vale database first. We'll maintain two KV databases:

1. short_code -> long_url
2. long_url -> short_code

### Production Stage

We'll consider the following considerations:

* Data persistence - in-memory is not enough.
* Code generation must be tolerant in distributed environment.

#### MySQL database

The MySQL database has the one table that stores the short code and the corresponding long URL.

```sql
CREATE TABLE url_shortener (
    id INT AUTO_INCREMENT PRIMARY KEY,
    short_code VARCHAR(6) UNIQUE,
    long_url VARCHAR(2048)
    expired_at TIMESTAMP NULL
);
```

## Later considerations

* Redundant data volume: the data volume is doubled. It is not cost-effective.
* Analytics: track number of hits per short URL
    * 301 (Permanent Redirect) vs 302 (Temporary Redirect)
* Custom aliases (e.g. short.ly/testing-doc)
    * The custom alias can be implemented as another unique field in the table.
* Do short URLs live forever? Would you consider adding TTL (time-to-live) to each mapping? How would expired entries be cleaned up in memory?
    * Expiration can be implemented as a separate thread.
    * Table-based approach with proper index is better than KV-based approach considering the expiration needs to search for expired entries. 
* Could users generate malicious short URLs (e.g., phishing links)?
    * Would rate-limiting or domain allowlists help?
