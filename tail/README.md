# 🧪 Backend Engineer Interview Test: tail -f-like Log Reader

* You're designing a tool similar to tail -f, which is used to monitor a log file in real time. The tool should be able to read new content appended to a log file as it's written, and display it to the user in near real time.

## 📝 Prompt

Design a backend system (or internal component) that behaves like:

```bash
tail -f /var/log/app.log
```

Your design should support:

* Reading appended content from a file continuously.
* Handling log rotation (when the original file is renamed or replaced).
* Efficient resource usage — avoid busy waiting.
* Optionally: support for keyword filtering (e.g., tail -f app.log | grep ERROR).
