# Solution: Funding & Withdrawals (On/Off Ramps)

## Context

Users currently cannot get real money into the system. We need to integrate with banking rails.

## Proposed Architecture

We will implement a "Payment Method" abstraction allowing users to link Cards and Bank Accounts. We will likely use a provider like **Stripe** or **Plaid** to handle the complexity of PCI compliance and bank communication.

### Step-by-Step Implementation

#### 1. Linking a Funding Source (Tokenization)

1.  **Client-Side**: User input credit card details into a regulated SDK (e.g., Stripe Elements).
2.  **Exchange**: SDK returns a temporary `token` or `payment_method_id` to the frontend.
3.  **Backend API**: `POST /wallets/funding-sources`.
    - Input: `payment_method_id`.
    - Action: Save this ID in a `funding_sources` table associated with the User. DO NOT save raw card numbers.

#### 2. Cash In (Deposit)

1.  **API**: `POST /wallets/deposit`.
2.  **Input**: `amount`, `source_id`.
3.  **Process**:
    - Create a database transaction `pending`.
    - Call Provider API (e.g., `stripe.PaymentIntents.Create(confirm=true)`).
    - **Synchronous Success**: If provider returns success immediately, credit User Wallet and update DB status to `success`.
    - **Asynchronous (ACH/SEPA)**: Provider returns `processing`. Keep DB status `pending`. Listen for webhook `payment_intent.succeeded` to credit wallet later.

#### 3. Cash Out (Withdrawal)

1.  **API**: `POST /wallets/withdraw`.
2.  **Checks**: Ensure `wallet.balance >= amount`.
3.  **Process**:
    - Debit User Wallet immediately (prevent double spend).
    - Create Payout request to provider (e.g., `stripe.Payouts.Create`).
    - If Provider fails request, roll back the debit (refund user).

## References & Open Source

- **Library**: [Stripe Go](https://github.com/stripe/stripe-go) - Official Go client for Stripe.
- **Standard**: [ISO 20022](https://www.iso20022.org/) - The global standard for financial messaging (relevant if building direct bank integrations).
- **Open Banking**: [Plaid Link](https://plaid.com/docs/link/) - Industry standard for connecting bank accounts securely.
