# Solution: P2P Transactions & Social Context

## Context

Transactions are currently just "User A sent money to User B". Real-world apps thrive on the social context (notes, emojis, request flows).

## Proposed Architecture

### 1. Payment Requests

A "Pull" mechanism where A asks B for money.

- **New Table**: `PaymentRequests` (`id`, `requester_id`, `payer_id`, `amount`, `status` [pending, paid, declined], `expires_at`).
- **API**: `POST /requests`.
- **Notification**: Payer receives a push notification.
- **Action**: Payer clicks "Pay". This triggers the standard `SendMoney` flow but references the `request_id` to mark it as `paid`.

### 2. Social Graph & Visibility

- **Privacy Settings**: Add `default_privacy` to User (public, friends, private). Override per transaction.
- **Feed Service**: A separate read-heavy service or optimized query to fetch "Friends' Activity".
  - Query: `SELECT * FROM transactions WHERE (sender_id IN my_friends OR receiver_id IN my_friends) AND privacy = 'public'`.

### 3. Split Bill Logic

- **Concept**: A "Split" is a parent container for multiple Payment Requests.
- **Step-by-Step**:
  1.  User selects a past transaction (e.g., $100 Dinner).
  2.  User selects 3 friends to split with.
  3.  Backend creates 3 `PaymentRequests` for $25 each.
  4.  Original transaction remains, but UI shows "Split with 3 others".

### 4. Recipient Validation (Confirmation of Payee)

- **Problem**: Mistyping a username causes loss of funds.
- **Solution**: Two-step send.
  1.  `GET /users/lookup?phone=+12345` -> Returns partial info: "Bob S. (Picture)".
  2.  User confirms visually.
  3.  `POST /payments` includes the `recipient_id` derived from step 1.

## References & Open Source

- **Graph Database**: [Neo4j](https://neo4j.com/) - Excellent for storing and querying social connections (Friends of Friends).
- **Library**: [Cayley](https://github.com/cayleygraph/cayley) - An open-source graph database written in Go.
- **Paper**: _'Social Network Analysis regarding P2P Money Transfers'_ - Analyzing transaction graphs to suggest friends.
