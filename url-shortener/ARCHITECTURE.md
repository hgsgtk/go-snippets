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

## System Components

### In-memory key-value database

We assume the data is stored in a in-memory key-vale database first. We'll maintain two KV databases:

1. short_code -> long_url
2. long_url -> short_code

### API Server


## Later considerations

* 301 (Permanent Redirect) vs 302 (Temporary Redirect)
