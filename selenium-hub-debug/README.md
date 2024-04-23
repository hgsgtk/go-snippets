# Selenium Hub/Node Debugging

This directory contains docker compose files and scripts to reveal the specification between selenium hub and nodes.

## Selenium Hub 3

[seleniarm/node-chronium](https://hub.docker.com/r/seleniarm/node-chromium/tags?page=9&page_size=&ordering=-last_updated&name=) doesn't provide selenium hub 3. So, we need to start selenium hub from java files.

### Node Spec Debug Setup

- Debug hub startup

```bash
go run debug-hub/main.go
```

- Chrome node startup

```bash
java -Dwebdriver.chrome.driver="./chromedriver.exe" -jar selenium-server-standalone-3.4.0.jar -role node -hub http://127.0.0.1:4444/grid/register -browser "browserName=chrome,maxInstances=1,platform=LINUX"
```

### Hub Spec Debug Setup

- Selenium Hub startup

```bash
java -jar selenium-server-standalone-3.4.0.jar -role hub
```

- Run the debug node script

```bash
go run debug-node/main.go
```

## Selenium Grid 4



## References

- [Selenium 3 archives](https://selenium-release.storage.googleapis.com/index.html?path=3.4/)
