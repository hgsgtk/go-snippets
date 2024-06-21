# Log filtering program

This programreads a log file and filters out the lines that contain a specific string.

## Log sample

```
$ ./sample.sh
[1718993769.410][INFO]: Starting ChromeDriver 105.0.5195.52 (412c95e518836d8a7d97250d62b29c2ae6a26a85-refs/branch-heads/5195@{#853}) on port 9515
[1718993769.410][INFO]: Please see https://chromedriver.chromium.org/security-considerations for suggestions on keeping ChromeDriver safe.
[1718993769.411][SEVERE]: bind() failed: Address already in use (48)
[1718993769.411][INFO]: listen on IPv4 failed with error ERR_ADDRESS_IN_USE
Starting ChromeDriver 105.0.5195.52 (412c95e518836d8a7d97250d62b29c2ae6a26a85-refs/branch-heads/5195@{#853}) on port 9515
Only local connections are allowed.
Please see https://chromedriver.chromium.org/security-considerations for suggestions on keeping ChromeDriver safe.
[1718994331.459][INFO]: Starting ChromeDriver 105.0.5195.52 (412c95e518836d8a7d97250d62b29c2ae6a26a85-refs/branch-heads/5195@{#853}) on port 9515
[1718994331.459][INFO]: Please see https://chromedriver.chromium.org/security-considerations for suggestions on keeping ChromeDriver safe.
IPv4 port not available. Exiting...
[1718994331.459][SEVERE]: bind() failed: Address already in use (48)
[1718994331.459][INFO]: listen on IPv4 failed with error ERR_ADDRESS_IN_USE
```

### Create a session

```
$ curl -X POST http://localhost:9516/session -d '{
    "desiredCapabilities": {
        "browserName": "chrome",
        "version": "",
        "platform": "ANY"
    }
}'
{"sessionId":"a070fb6fa4b5070bc1fbb20f4220c541","status":33,"value":{"message":"session not created: This version of ChromeDriver only supports Chrome version 105\nCurrent browser version is 126.0.6478.63 with binary path /Applications/Google Chrome.app/Contents/MacOS/Google Chrome\n  (Driver info: chromedriver=105.0.5195.52 (412c95e518836d8a7d97250d62b29c2ae6a26a85-refs/branch-heads/5195@{#853}),platform=Mac OS X 12.5.1 arm64)"}}
```

Here is the log written by chromedriver.

```
[1718995185.053][INFO]: [d535695a1035bcf58c266565a6a7cd1e] COMMAND InitSession {
}
[1718995185.054][INFO]: [d535695a1035bcf58c266565a6a7cd1e] RESPONSE InitSession ERROR invalid argument: 'capabilities' must be a JSON object
[1718995323.842][INFO]: [a070fb6fa4b5070bc1fbb20f4220c541] COMMAND InitSession {
   "desiredCapabilities": {
      "browserName": "chrome",
      "platform": "ANY",
      "version": ""
   }
}
[1718995323.843][INFO]: Populating Preferences file: {
   "alternate_error_pages": {
      "enabled": false
   },
   "autofill": {
      "enabled": false
   },
   "browser": {
      "check_default_browser": false
   },
   "distribution": {
      "import_bookmarks": false,
      "import_history": false,
      "import_search_engine": false,
      "make_chrome_default_for_user": false,
      "skip_first_run_ui": true
   },
   "dns_prefetching": {
      "enabled": false
   },
   "profile": {
      "content_settings": {
         "pattern_pairs": {
            "https://*,*": {
               "media-stream": {
                  "audio": "Default",
                  "video": "Default"
               }
            }
         }
      },
      "default_content_setting_values": {
         "geolocation": 1
      },
      "default_content_settings": {
         "geolocation": 1,
         "mouselock": 1,
         "notifications": 1,
         "popups": 1,
         "ppapi-broker": 1
      },
      "password_manager_enabled": false
   },
   "safebrowsing": {
      "enabled": false
   },
   "search": {
      "suggest_enabled": false
   },
   "translate": {
      "enabled": false
   }
}
[1718995323.843][INFO]: Populating Local State file: {
   "background_mode": {
      "enabled": false
   },
   "ssl": {
      "rev_checking": {
         "enabled": false
      }
   }
}
[1718995323.844][INFO]: Launching chrome: /Applications/Google Chrome.app/Contents/MacOS/Google Chrome --allow-pre-commit-input --disable-background-networking --disable-client-side-phishing-detection --disable-default-apps --disable-hang-monitor --disable-popup-blocking --disable-prompt-on-repost --disable-sync --enable-automation --enable-blink-features=ShadowDOMV0 --enable-logging --log-level=0 --no-first-run --no-service-autorun --password-store=basic --remote-debugging-port=0 --test-type=webdriver --use-mock-keychain --user-data-dir=/var/folders/bx/jq4_sc7n17xgvhd89gnjxkym0000gn/T/.com.google.Chrome.iKJ2jh data:,
[1718995329.763][INFO]: Failed to connect to Chrome. Attempting to kill it.
[1718995330.245][INFO]: [a070fb6fa4b5070bc1fbb20f4220c541] RESPONSE InitSession ERROR session not created: This version of ChromeDriver only supports Chrome version 105
Current browser version is 126.0.6478.63 with binary path /Applications/Google Chrome.app/Contents/MacOS/Google Chrome
```

### Delete a session

```
$ curl -X DELETE http://localhost:9516/session/a070fb6fa4b5070bc1fbb20f4220c541
{"value":null}
```
