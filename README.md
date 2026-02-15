# Open Source ABIS Prototype

A scalable backend prototype for an Automated Biometric Identification System (ABIS) built using Go.

## Features

- REST API endpoints
- Thread-safe in-memory storage
- 1:N biometric-style matching
- Cosine similarity calculation
- Configurable similarity threshold
- Performance measurement (search time tracking)

## Endpoints

### Health Check
GET /health

### Enroll User
POST /enroll

```json
{
  "id": "user1",
  "embedding": [0.1, 0.2, 0.3]
}
