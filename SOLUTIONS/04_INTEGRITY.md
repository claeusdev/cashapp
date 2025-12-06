# Solution: Transaction Integrity & Reliability

## Context

In distributed systems, networks fail. We need to ensure money is never created or destroyed accidentally (ACID) and that requests are processed exactly once (Idempotency).

## Proposed Architecture

### 1. Idempotency Key Pattern

- **Concept**: Clients generate a unique key (UUID v4) for every mutating request (`POST /payments`).
- **Middleware Step-by-Step**:
  1.  Extract `Idempotency-Key` header.
  2.  Check Redis: `GET idempotency:{key}`.
  3.  **Hit**: Return the _cached response_ immediately. Do not process logic.
  4.  **Miss**:
      - Set Redis key with status `processing` (with TTL).
      - Execute business logic.
      - Update Redis key with status `completed` and the serialized response body.
- **Outcome**: If the client retries the request 5 times due to timeouts, the money is moved only once.

### 2. Distributed Locking (Redlock)

- **Problem**: A user taps "Send" on two devices simultaneously.
- **Solution**: Lock the wallet during processing.
- **Implementation**:
  - Acquire Lock: `Redlock.Lock("wallet:{id}", 200ms)`.
  - If failed, return 429 (Try again).
  - If success, proceed with transaction. Unlock immediately after.

### 3. Reconciliation (The Source of Truth)

- **Problem**: Software bugs or database corruption might cause the sum of all internal wallets to diverge from the money held in the application's actual bank account.
- **Solution**: Daily Reconciliation Job.
- **Step-by-Step**:
  1.  **Internal Sum**: `SELECT SUM(balance) FROM wallets`.
  2.  **External Balence**: API Call to Stripe/Bank -> `GetBalance()`.
  3.  **Compare**: They must match perfectly (or within a known float).
  4.  **Alert**: If `Internal > External`, we are insolvent. CRITICAL ALERT.
  5.  **Ledger Replay**: If mismatch, script replays all `transaction_events` to rebuild balances and find the erroneous transaction.

## References & Open Source

- **Article**: [Stripe: Designing robust APIs for idempotency](https://stripe.com/blog/idempotency).
- **Library**: [go-redsync](https://github.com/go-redsync/redsync) - Distributed mutual exclusion lock using Redis (implementation of Redlock).
- **Pattern**: [Saga Pattern](https://microservices.io/patterns/data/saga.html) - Managing distributed transactions across microservices (User Svc <-> Ledger Svc).
