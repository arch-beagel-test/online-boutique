---
status: proposed
date: 2026-10-14
title: Currency rates are fetched once a day and cached
---

# ADR 0003: Currency rates are fetched once a day and cached

## Context

`currencyservice` (Node.js) converts prices using European Central Bank reference rates. Today it reads a static
JSON file shipped with the image, so rates go stale between releases.

## Decision

Fetch the ECB rates once a day at 16:30 CET, cache them in memory, and fall back to the shipped file when the fetch
fails. Conversions round half to even to two decimal places.

## Consequences

Prices can move once a day. A failed fetch is logged at warning level and retried hourly.
