# Solution: Support & Administration

## Context

When things go wrong, humans need tools to fix data. Direct SQL access is dangerous and lacks audit trails.

## Proposed Architecture

### 1. Admin API

- A separate set of endpoints guarded by strict roles (`role: super_admin`, `role: support_agent`).
- **Capabilities**:
  - `GET /admin/users/{id}/history`: View full audit log.
  - `POST /admin/users/{id}/freeze`: Lock account.
  - `POST /admin/transactions/{id}/refund`: Reverse a transaction.

### 2. Audit Logging

- **Requirement**: Every action taken by an admin must be logged.
- **Table**: `AdminAuditLog` (`admin_id`, `action`, `target_resource`, `changes_json`, `timestamp`, `ip_address`).
- **Why**: If an internal employee steals money or data, we need to trace it.

### 3. Dispute Resolution System

- **Flow**:
  1.  User clicks "Report Issue" on transaction.
  2.  System creates a `Dispute` ticket (`status: open`).
  3.  Funds associated with the transaction are put in "Escrow" or "Frozen" state if possible.
  4.  Support Agent reviews evidence on Admin Dashboard.
  5.  Agent decision: `Resolve (Refund)` or `Reject`.

## References & Open Source

- **Dashboard Framework**: [Retool](https://retool.com/) - Rapidly build internal tools connecting to REST APIs and Postgres.
- **Admin Framework**: [Go-Admin](https://github.com/GoAdminGroup/go-admin) - Data visualization and admin panel framework for Golang.
- **Role Based Access Control (RBAC)**: [Casbin](https://github.com/casbin/casbin) - An authorization library that supports access control models like ACL, RBAC, ABAC in Go.
