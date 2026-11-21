# ADR 0001: gRPC between services

- **Status:** Accepted
- **Date:** 2026-03-02
- **Deciders:** Platform team

## Context and problem statement

The boutique has eleven services written in Go, C#, Node.js, Python and Java. They call each other on every
page view. We need one way for them to talk that has typed contracts and works in every language.

## Considered options

1. REST with JSON and an OpenAPI file per service
2. gRPC with Protocol Buffers, one `demo.proto` shared by all services
3. A message broker for every call

## Decision outcome

Option 2. Every service-to-service call is gRPC, defined in `protos/demo.proto`. The frontend is the only HTTP
entry point for users. Asynchronous messaging is not used for request paths.

## Consequences

- Good: one contract file, generated clients in five languages.
- Bad: browsers can't call services directly; everything goes through the frontend.
