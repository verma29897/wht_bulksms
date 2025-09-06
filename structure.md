+-------------------+      +----------------+      +-------------------+
|                   |      |                |      |                   |
|   User / Client   |----->|   API Server   |----->|    Job Queue      |
| (Submits CSV/JSON)|      |    (GoLang)    |      | (RabbitMQ / NATS) |
+-------------------+      +----------------+      +-------------------+
        ^       |                                             |
        |       | (Job ID)                                    | (Message Tasks)
        |       |                                             V
        |       |  +-------------------------------------------------+
        |       |  |                                                 |
        |       +--|                 Database (PostgreSQL)           |
        |          | (Stores Campaign & Message Status)              |
        |          +-------------------------------------------------+
        |                                 ^       |
        |                                 |       | (Update Status)
        | (Check Status)                  |       |
        |                                 |       V
        |  +------------------+         +------------------+
        |  |                  |         |                  |
        +->|  Webhook Handler |<--------|   Worker Pool    |
           |     (GoLang)     |         |     (GoLang)     |
           +------------------+         +------------------+
                    ^                           |
                    |                           | (API Calls)
                    | (Status Updates)          V
                    |                  +------------------+
                    +------------------|   Meta WhatsApp  |
                                       |     Cloud API    |
                                       +------------------+