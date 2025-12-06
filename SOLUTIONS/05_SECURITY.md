# Solution: Security & Fraud Prevention

## Context

Financial apps are primary targets for attackers. We need robust authentication and proactive fraud detection.

## Proposed Architecture

### 1. Authentication (JWT + Refresh Tokens)

- **Current**: No auth.
- **Proposal**:
  - `POST /login`: Returns `access_token` (15 min life) and `refresh_token` (7 day life, HTTPOnly Cookie).
  - **Middleware**: Validates JWT signature (RSA/ECDSA) on every request.
  - **Revocation**: store `refresh_token` hash in DB. If account compromised, delete hash to force logout on all devices.

### 2. Multi-Factor Authentication (MFA)

- **Triggers**: Login on new device, High value transfer (>$500).
- **Methods**: TOTP (Google Authenticator) or SMS (Twilio).
- **Flow**:
  1.  User initiates transfer.
  2.  Backend detects check: `NeedsMFA`. Returns `403 Forbidden` with `error_code: mfa_required`.
  3.  Client prompts user for code.
  4.  Client retries req with header `X-MFA-Code: 123456`.

### 3. Fraud Detection Engine (Rule-Based)

- **Concept**: Synchronous checks before approving a transaction.
- **Step-by-Step**:
  1.  Transaction hits `PaymentService`.
  2.  Service calls `FraudEngine.Check(transaction)`.
  3.  **Rule 1 (Velocity)**: `Sum(transactions_last_24h) > DailyLimit`? -> RESTRICT.
  4.  **Rule 2 (Impossible Travel)**: Last login IP in London, current IP in Tokyo, time diff < 2 hours? -> BLOCK.
  5.  **Rule 3 (New Device High Value)**: Device is new AND amount > $200? -> REQUIRE_MFA.
- **Outcome**: Transaction is Approved, Rejected, or Flagged for Review.

## References & Open Source

- **Library**: [Golang-JWT](https://github.com/golang-jwt/jwt) - Standard for token handling.
- **Engine**: [Grule Rule Engine](https://github.com/hyperjumptech/grule-rule-engine) - Rule engine for Go. Allows defining rules in a simple DSL (Domain Specific Language) separate from code.
- **Algorithm**: **Leaky Bucket** for Rate Limiting/Velocity checks.
- **Security Standard**: [OWASP Top 10](https://owasp.org/www-project-top-ten/) and [NIST Digital Identity Guidelines](https://pages.nist.gov/800-63-3/).
