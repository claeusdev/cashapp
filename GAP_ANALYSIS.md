# Gap Analysis: Current vs. Real-World

Based on the `REAL_WORLD_SCENARIOS.md`, here is a breakdown of what is currently implemented versus what is missing in our application.

## 1. User Onboarding & Identity (KYC/AML)

- **Status: ❌ Missing**
- **Current State**: We only create a user with a `tag` (username). There is no authentication (password/token) and no identity verification.
- **Missing**:
  - ID Verification (KYC).
  - Sanctions Screening.
  - Risk Tiering.

## 2. Funding & Withdrawals

- **Status: ❌ Missing**
- **Current State**: Money can only be "seeded" via our CLI or direct database manipulation. There is no connection to the outside banking world.
- **Missing**:
  - Bank/Card integrations (Stripe, Plaid, etc.).
  - Webhooks for payment status.

## 3. P2P Transactions

- **Status: ⚠️ Partial**
- **Current State**: We can transfer money between users (`SendMoney`).
- **Missing**:
  - **Context**: We don't have "Request Money", notes/emojis, or visibility permissions (public/private).
  - **Split Bill**: No logic to handle multi-user splits.
  - **Validation**: We check if a user exists, but we don't show the user's full name for confirmation before sending.

## 4. Transaction Integrity

- **Status: ⚠️ Partial**
- **Current State**: We utilize database transactions (ACID) for ensuring money isn't lost _within_ the SQL transaction.
- **Missing**:
  - **Idempotency**: The API does not accept or check unique `Idempotency-Keys`. A network retry could result in double-spending.
  - **Distributed Locking**: We rely on Postgres row locking. For high concurrency across distributed instances, a Redis-based distributed lock (Redlock) might be needed.
  - **Reconciliation**: No background workers checking ledgers against a source of truth.

## 5. Security & Fraud

- **Status: ❌ Missing**
- **Current State**: No limits. A user can transfer any amount they have. No authentication.
- **Missing**:
  - Auth Middleware (JWT/OAuth).
  - Velocity Limits (e.g., max $1000/day).
  - MFA Challenges.
  - Anomaly Detection.

## 6. Notifications

- **Status: ❌ Missing**
- **Current State**: Synchronous API responses only. The receiver doesn't know they got money unless they check their balance.
- **Missing**:
  - Push/Email/SMS notifications service.

## 7. Support

- **Status: ❌ Missing**
- **Current State**: Only raw SQL access for admins.
- **Missing**:
  - Admin API/Dashboard.
  - Dispute resolution flows.
