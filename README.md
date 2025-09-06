# Bulk WhatsApp Panel (React + Go)

A full-stack panel to onboard WhatsApp Business Accounts (Embedded Signup), upload media/templates, and send bulk messages using Meta WhatsApp Cloud API.

## Tech Stack
- Backend: Go (Gin), JWT auth, PostgreSQL (GORM + SQL), concurrent workers for bulk send
- Frontend: React (Vite), React Router, simple CSS

## Repository Structure
```
meta_data/
  api/                 # Go API
    handlers/          # HTTP handlers (auth, templates, media, send, onboarding)
    models/            # Data models (User, Account, Templates, Messages)
    routers/           # Route registration
    db/                # DB init (GORM) + SQL schema
    utils/             # Helpers (media types)
    main.go            # API entrypoint
    curl.md            # Example curl requests
  ui_web/              # React UI (Vite)
    src/               # Pages, layout, lib/api helper, styles
  structure.md         # High-level architecture diagram
  README.md            # This file
```

## Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 14+
- Meta Developer App (Business) with WhatsApp product

## Environment Variables
### API
- `PORT`: API port (default 8081 recommended)
- `DB_URL`: PostgreSQL DSN
  - TCP example: `postgres://USER:PASS@localhost:5432/DBNAME?sslmode=disable`
  - Socket example: `postgres:///DBNAME?host=/var/run/postgresql`
- `JWT_SECRET`: secret to sign JWTs
- `META_ACCESS_TOKEN`: long-lived token for WhatsApp Cloud API calls
- `APP_ID`, `APP_SECRET`, `REDIRECT_URI`: required to exchange Embedded Signup `auth_code`

### UI
- `VITE_API_BASE` (optional). If not set, Vite dev proxy `/api` → `http://localhost:8081` is used
- `VITE_FB_APP_ID`: Meta App ID
- `VITE_WHATSAPP_EMBEDDED_SIGNUP_CONFIG_ID`: Embedded Signup config id

Create `ui_web/.env` (dev example):
```
VITE_FB_APP_ID=YOUR_APP_ID
VITE_WHATSAPP_EMBEDDED_SIGNUP_CONFIG_ID=YOUR_CONFIG_ID
# optional for prod builds (bypass proxy)
# VITE_API_BASE=http://localhost:8081
```

## Database Setup (pick one)
### A) TCP with password
```
sudo -u postgres psql -c "CREATE ROLE \"your_user\" LOGIN PASSWORD 'STRONG_PASS';"
sudo -u postgres createdb bulksms
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE bulksms TO \"your_user\";"
export DB_URL='postgres://your_user:STRONG_PASS@localhost:5432/bulksms?sslmode=disable'
```

### B) Local Unix socket (no TCP)
```
sudo -u postgres createdb bulksms
export DB_URL='postgres:///bulksms?host=/var/run/postgresql'
```

Load schema (one-time):
```
psql "$DB_URL" -f ./api/db/query.sql
```

## Run (development)
### API
```
cd api
PORT=8081 DB_URL="postgres://USER:PASS@localhost:5432/DB?sslmode=disable" JWT_SECRET=dev-secret go run .
```

### UI (Vite dev server with proxy to 8081)
```
cd ui_web
npm install
npm run dev -- --host
# open https://localhost:5173/
```

## API Endpoints (summary)
- Public
  - POST `/auth/signup` → create user (name, email, username, password)
  - POST `/auth/login` → returns `{ token }`
  - GET `/onboard/callback?auth_code=...` → exchanges Meta `auth_code` (Embedded Signup)
- Protected (Bearer token)
  - GET `/templates/:waba_id` → fetch templates from Meta
  - POST `/templates` → create a template on Meta
  - POST `/upload` → create upload session (header handle)
  - POST `/upload/header` → upload media to get `media_id`
  - POST `/send` → bulk send using a template
  - POST `/store-onboarding` → persist onboarded account (used by success page)
  - GET `/users` → list users (GORM)
  - GET `/accounts` → list accounts (GORM)

See `api/curl.md` for request examples.

## Embedded Signup (UI)
- `/register` launches `FB.login` with `config_id` and `response_type=code`
- UI calls `/api/onboard/callback?auth_code=...` via proxy → API exchanges the code
- `/register/success` posts to `/store-onboarding` (protected) to persist account

Checklist:
- Configure App Domains and Valid OAuth Redirect URIs
- Permissions: `whatsapp_business_management`, `whatsapp_business_messaging`

## Bulk Send Payload
```
POST /send (Bearer token)
{
  "phone_number_id": "<NUMBER_ID>",
  "template_name": "<TEMPLATE_NAME>",
  "language": "en_US",
  "media_type": "IMAGE|DOCUMENT|VIDEO|AUDIO|",  // optional
  "media_id": "<MEDIA_ID>",                     // optional
  "contact_list": ["919000000000", "919111111111"]
}
```

## Troubleshooting
- UI: `Unexpected token '<'`
  - Ensure UI calls `/api/...` (proxy) or set `VITE_API_BASE` to API URL
- DB: `pq: password authentication failed` or `connect: connection refused`
  - Fix `DB_URL` credentials, or start Postgres and verify it listens
- Port already in use
  - API: change `PORT` or kill previous process; UI: stop existing Vite
- Embedded signup fails
  - Verify `APP_ID`, `APP_SECRET`, `REDIRECT_URI`, Meta config; check browser console

## Security Notes
- Keep secrets out of source; use env vars or a secrets manager
- Use a strong `JWT_SECRET` in production
- Add rate limiting and strict validation for production

## License
MIT (or as you prefer)
