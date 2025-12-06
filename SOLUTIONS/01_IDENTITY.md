# Solution: Identity Verification & Compliance (KYC/AML)

## Context

Our current system lacks identity verification. To operate legally as a financial service, we must verify users (KYC) and screen them against sanctions lists (AML).

## Proposed Architecture

We will implement a 3-tier identity system:

1.  **Tier 0 (Unverified)**: Created on sign-up. Can browse but cannot transact.
2.  **Tier 1 (Basic)**: Phone + Email Verified. Low limits ($500/month).
3.  **Tier 2 (Full)**: Government ID + Selfie Verified. High limits ($10k/month).

### Step-by-Step Implementation

1.  **Data Modeling Changes**:

    - Add `kyc_level` (int), `kyc_status` (pending, verified, rejected), and `risk_score` (int) to the `User` model.
    - Create a `IdentityDocuments` table to store metadata (URLs to secure S3 buckets) of uploaded ID images, although relying on a 3rd party provider is safer.

2.  **Integration with Provider (e.g., Stripe Identity / SumSub/ Onfido)**:

    - **Backend**: Create an endpoint `POST /verification/session`.
    - **Logic**: Call provider API to create a verification session. Returns a `url` or `client_secret`.
    - **Frontend**: Redirect user to this URL or open SDK modal to capture ID photos and Selfie.

3.  **Webhook Handling**:

    - **Endpoint**: `POST /webhooks/identity`.
    - **Logic**:
      - Receive payload from provider (`verification.passed` or `verification.failed`).
      - Validate webhook signature (security).
      - Update `user.kyc_status` to `verified`.
      - If `failed`, notify user via email with reason.

4.  **Sanctions Screening (AML)**:

    - Most providers bundle this. If `verification.passed` but `sanctions.match` is true, set user status to `manual_review` and alert admin.

5.  **Access Control**:
    - Middleware `RequireKYC(level int)` on transaction endpoints.

## References & Open Source

- **Open Source Identity Server**: [Ory Kratos](https://github.com/ory/kratos) - Handles user management, flow logic, and can integrate with ID providers.
- **Research**: _'Anti-Money Laundering in Bitcoin: Experimenting with Graph Convolutional Networks for Financial Forensics'_ (arXiv:1908.02591) - Advanced but relevant for understanding risk scoring.
- **Reference Implementation**: Stripe Identity Docs (Excellent standard for flow design).
