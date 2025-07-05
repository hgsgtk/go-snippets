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

## System Components

### In-memory key-value database

We assume the data is stored in a in-memory key-vale database first. We'll maintain two KV databases:

1. short_code -> long_url
2. long_url -> short_code

### API Server


## Later considerations

* 301 (Permanent Redirect) vs 302 (Temporary Redirect)
