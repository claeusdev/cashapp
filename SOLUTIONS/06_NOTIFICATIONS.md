# Solution: Notifications System

## Context

Users need to know when money arrives or leaves immediately.

## Proposed Architecture

We should move from a synchronous request-response model to an event-driven model for notifications using a **Message Queue**.

### Step-by-Step Implementation

1.  **Event Emission**:

    - When the Ledger Service successfully completes a transfer, it publishes an event to a message broker (RabbitMQ/Kafka/Redis Streams).
    - Topic: `transactions`
    - Payload: `{ type: "transfer_success", to_user_id: 101, amount: 5000, currency: "USD" }`

2.  **Notification Service (Consumer)**:

    - A new microservice subscribes to the `transactions` topic.
    - It looks up the user's device tokens (FCM/APNS tokens) from a `UserDevices` table.

3.  **Delivery Channel**:

    - **Push**: Send payload to Apple APNS / Google FCM.
    - **Email**: Send receipt via SendGrid/AWS SES.
    - **SMS**: Send text via Twilio.

4.  **In-App Feed**:
    - The service also writes an entry to the `ActivityFeed` table (`user_id`, `message`, `icon`, `deep_link`, `read_status`).
    - Client polls `GET /feed` or listens via WebSocket for real-time updates.

## References & Open Source

- **Task Queue**: [Asynq](https://github.com/hibiken/asynq) - Simple, reliable, and efficient distributed task queue for Go based on Redis.
- **Broker**: [RabbitMQ](https://www.rabbitmq.com/) or [NATS JetStream](https://nats.io/) (Go native).
- **Library**: [Gorilla WebSocket](https://github.com/gorilla/websocket) - For real-time in-app updates.
