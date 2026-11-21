# ADR 0002: Cart state lives in Redis

- **Status:** Accepted
- **Date:** 2026-03-09

## Context

`cartservice` (C#) must keep a shopper's cart between page views and survive a pod restart. Carts are small,
read often, and can be lost after 48 hours without harm.

## Decision

Carts are stored in Redis, keyed by the user's session ID, with a time to live of 48 hours. `cartservice` is the
only service that reads or writes the cart store. Alloy DB and Spanner were considered and rejected as too heavy
for data we may lose.

## Consequences

- A Redis outage empties carts but doesn't break checkout of carts already loaded.
- See the runbook "Cart Redis outage" for recovery.
