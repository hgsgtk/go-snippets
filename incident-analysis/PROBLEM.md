# 🛠️ Pair Programming: Incident Aggregator & Alert Scorer

## Problem Statement

A production environment spans multiple regions and thousands of services. One of the core responsibilities of the Enterprise SRE team is to ensure incident response is fast, reliable, and intelligent.

You're tasked with building a simplified tool that ingests incident alerts, aggregates them, and assigns a reliability score to help prioritize resolution.

## 🧩 Requirements

Write a program that:

- Accepts a stream of alerts, each with:
  - `service_name`: string
  - `timestamp`: RFC3339 format
  - `severity`: integer (1 = highest severity, 5 = lowest severity)
  - `region`: string (e.g. ap-northeast-1, us-east-1)

- Aggregates alerts by `service_name` and `region`.

- For each `(service_name, region)` pair, compute:
  - `total_alerts`
  - `most_recent_alert_timestamp`
  - `average_severity` (float, rounded to 2 decimal places)
  - `reliability_score` (lower severity + recent alerts = worse score)

- Sort and output the top N least reliable services based on the computed `reliability_score`.

## 🧠 Scoring Logic (You define it!)

You should design the `reliability_score` function. A simple example:

```go
reliability_score = average_severity * log(total_alerts + 1) + recentness_penalty
```

Where `recentness_penalty` could be:

```go
recentness_penalty = max(0, 100 - minutes_since_last_alert)
```

You can adjust this based on your judgment.

## 🧪 Example Input

```json
[
  {
    "service_name": "checkout-api",
    "timestamp": "2025-07-08T06:00:00Z",
    "severity": 2,
    "region": "ap-northeast-1"
  },
  {
    "service_name": "checkout-api",
    "timestamp": "2025-07-08T06:10:00Z",
    "severity": 1,
    "region": "ap-northeast-1"
  },
  {
    "service_name": "inventory-worker",
    "timestamp": "2025-07-08T05:00:00Z",
    "severity": 3,
    "region": "us-east-1"
  }
]
```

And a call like:

```bash
get_unreliable_services(alerts, topN=2)
```

## ✅ Output

Your program should return:

```json
[
  {
    "service_name": "checkout-api",
    "region": "ap-northeast-1",
    "total_alerts": 2,
    "most_recent_alert_timestamp": "2025-07-08T06:10:00Z",
    "average_severity": 1.5,
    "reliability_score": 98.2
  },
  {
    "service_name": "inventory-worker",
    "region": "us-east-1",
    "total_alerts": 1,
    "most_recent_alert_timestamp": "2025-07-08T05:00:00Z",
    "average_severity": 3.0,
    "reliability_score": 47.1
  }
]
```

## 🧱 Constraints

- Time parsing and timezones matter.
- You should handle out-of-order timestamps.
- Alerts could come in batches or one at a time (you can choose the interface).
- You must support thousands of services (optimize your approach for memory or latency where needed). 
