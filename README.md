# Rehoboam

> A production-grade self-hosted homeserver infrastructure — exposing services securely to the internet via Cloudflare Zero Trust tunnels, with a full observability pipeline and containerized service deployment.

[![Arch Linux](https://img.shields.io/badge/OS-Arch_Linux-1793D1?style=flat&logo=arch-linux&logoColor=white)](https://archlinux.org/)
[![Docker](https://img.shields.io/badge/Runtime-Docker-2496ED?style=flat&logo=docker&logoColor=white)](https://docker.com/)
[![Cloudflare](https://img.shields.io/badge/Tunnel-Cloudflare-F38020?style=flat&logo=cloudflare&logoColor=white)](https://cloudflare.com/)
[![Grafana](https://img.shields.io/badge/Observability-Grafana-F46800?style=flat&logo=grafana&logoColor=white)](https://grafana.com/)

---

## Overview

Rehoboam is a self-hosted homeserver designed around production infrastructure principles — separation of concerns between source, deployment config, and persistent state; secrets management via environment variables; containerized services with isolated networks; and a centralized observability stack collecting logs from every layer of the system.

All services are exposed to the internet without opening a single port on the router, using Cloudflare Zero Trust tunnels that establish outbound-only connections to Cloudflare's edge network.

---

## Hardware

| Component | Specification |
|-----------|--------------|
| CPU | AMD Ryzen 5 PRO 1500 — 4 cores / 8 threads @ 3.5GHz |
| RAM | 16GB DDR4 |
| Storage | Single SSD |
| GPU | AMD OLAND 2GB VRAM |
| OS | Arch Linux (x86_64) |
| Network | Gigabit Ethernet (DHCP reserved) |

---

## Architecture

```
                        ┌─────────────────────────────────┐
                        │        Cloudflare Edge           │
                        │  (DDoS protection, TLS, DNS)     │
                        └────────────────┬────────────────┘
                                         │ HTTPS
                        ┌────────────────▼────────────────┐
                        │         cloudflared              │
                        │    (Zero Trust tunnel daemon)    │
                        └──┬──────┬──────┬──────┬─────────┘
                           │      │      │      │
               ┌───────────▼─┐ ┌──▼──┐ ┌▼────┐ ┌▼──────────┐
               │  Portfolio  │ │ NC  │ │ GF  │ │    SSH     │
               │  :8081      │ │8082 │ │3000 │ │   :22      │
               └─────────────┘ └─────┘ └─────┘ └────────────┘

NC = Nextcloud    GF = Grafana
```

### Network Isolation

Each service stack runs on its own Docker bridge network. Containers communicate internally by name — no inter-stack exposure unless explicitly configured.

```
logging_logging:    loki ←→ fluent-bit ←→ grafana
nextcloud_nextcloud: nextcloud ←→ nextcloud-db ←→ nextcloud-redis
portfolio_default:  portfolio (nginx, standalone)
```

---

## Services

| Service | Domain | Stack | Description |
|---------|--------|-------|-------------|
| Portfolio | `portfolio.vincents.systems` | React + Vite + Nginx | Personal portfolio, multi-stage Docker build |
| Nextcloud | `files.vincents.systems` | Nextcloud + PostgreSQL + Redis | Self-hosted file storage and sync |
| Grafana | `grafana.vincents.systems` | Grafana | Observability dashboards |
| Loki | internal | Grafana Loki | Log aggregation and storage |
| Fluent Bit | internal | Fluent Bit v5 | Log collection from Docker and systemd |
| PostgreSQL | internal | Postgres 16 | Nextcloud metadata database |
| Redis | internal | Redis Alpine | Nextcloud caching layer |

---

## Observability Stack

A three-layer logging pipeline collects logs from every part of the system automatically — no per-service configuration required for new containers.

```
┌─────────────────────────────────────────────────────┐
│                    Fluent Bit                        │
│                                                      │
│  [INPUT: forward]     ← Docker containers (fluentd  │
│                         logging driver on :24224)   │
│                                                      │
│  [INPUT: systemd]     ← systemd journal              │
│                         (SSH, cloudflared, kernel)  │
│                                                      │
│  [FILTER: grep]       ← excludes fluent-bit self-   │
│                         logs to prevent feedback    │
│                                                      │
│  [FILTER: record_modifier] ← injects host=forge,   │
│                               environment=homelab  │
│                                                      │
│  [OUTPUT: loki]       → ships to Loki on :3100      │
└─────────────────────────────────────────────────────┘
                              │
                    ┌─────────▼──────────┐
                    │        Loki         │
                    │                     │
                    │  Labels indexed:    │
                    │  - container_name  │
                    │  - SYSLOG_IDENTIFIER│
                    │  - job             │
                    │  - host            │
                    │                     │
                    │  Storage:           │
                    │  /data/loki/chunks  │
                    │  (retained forever) │
                    └─────────┬──────────┘
                              │
                    ┌─────────▼──────────┐
                    │       Grafana       │
                    │                     │
                    │  Dashboards:        │
                    │  - Log volume/svc  │
                    │  - SSH logins      │
                    │  - Error rate      │
                    │  - Tunnel health   │
                    └─────────────────────┘
```

**Docker logging driver** is configured globally in `/etc/docker/daemon.json` — every container automatically ships logs to Fluent Bit's forward input on `172.17.0.1:24224` without any per-container configuration.

**Systemd journal** logs are read natively by Fluent Bit's systemd input plugin, capturing SSH sessions, cloudflared events, and all system-level activity.

---

## Filesystem Structure

Designed around the principle of separating **source code**, **deployment config**, **runtime state**, and **secrets** — mirroring how production systems separate these concerns across different systems.

```
/srv/rehoboam/
│
├── apps/                        # Deployment manifests (docker-compose per service)
│   ├── portfolio/
│   │   └── docker-compose.yml
│   ├── nextcloud/
│   │   └── docker-compose.yml
│   └── logging/
│       └── docker-compose.yml
│
├── repos/                       # Source code (local clone, disposable)
│   └── portfolio/               # Vite + React + TypeScript
│       ├── Dockerfile           # Multi-stage: node:alpine → nginx:alpine
│       └── src/
│
├── infra/                       # Cross-cutting infrastructure config
│   ├── cloudflared/             # Tunnel config (credentials gitignored)
│   ├── fluent-bit/
│   │   ├── fluent-bit.yaml      # Collection pipeline config
│   │   └── parsers.conf
│   └── loki-config.yml
│
├── data/                        # Persistent volumes — gitignored
│   ├── loki/                    # Log chunks, index, WAL
│   ├── grafana/                 # Dashboards, alert configs
│   └── nextcloud/               # Files, DB, config
│
├── backups/                     # Scheduled snapshots — gitignored
├── logs/                        # Centralized log output — gitignored
│
├── .env                         # Secrets — gitignored
├── .env.example                 # Secret template — committed
└── .gitignore
```

**Key design decision:** `data/` is completely separate from `apps/`. Wiping and redeploying a service never touches persistent state. Backing up the entire server state means backing up only `data/` — one directory, one command.

---

## Security

| Layer | Implementation |
|-------|---------------|
| SSH | ed25519 keys only, password auth disabled, root login disabled |
| Network exposure | Zero open inbound ports — Cloudflare tunnel is outbound-only |
| DDoS protection | Cloudflare edge absorbs attacks before they reach the server |
| Container isolation | Each stack on its own Docker bridge network |
| Secrets | Environment variables via `.env`, never committed to git |
| Systemd journal | Persistent to disk, collected and retained in Loki |

---

## Setup

### Prerequisites

- Arch Linux (or any systemd-based Linux)
- Docker + Docker Compose
- A domain on Cloudflare
- `cloudflared` installed

### 1. Clone

```bash
git clone https://github.com/VincentSamuelPaul/Homeserver.git /srv/rehoboam
cd /srv/rehoboam
```

### 2. Configure secrets

```bash
cp .env.example .env
nano .env
```

### 3. Configure Cloudflare tunnel

```bash
cloudflared tunnel login
cloudflared tunnel create homeserver
sudo cp ~/.cloudflared/homeserver.json /etc/cloudflared/
sudo cp ~/.cloudflared/cert.pem /etc/cloudflared/
sudo cloudflared service install
sudo systemctl enable --now cloudflared
```

### 4. Configure Docker logging driver

```bash
sudo nano /etc/docker/daemon.json
```

```json
{
  "log-driver": "fluentd",
  "log-opts": {
    "fluentd-address": "172.17.0.1:24224",
    "fluentd-async": "true",
    "tag": "docker.{{.Name}}"
  }
}
```

```bash
sudo systemctl restart docker
```

### 5. Start the logging stack first

```bash
cd /srv/rehoboam/apps/logging
docker compose up -d
```

### 6. Start remaining services

```bash
cd /srv/rehoboam/apps/nextcloud && docker compose up -d
cd /srv/rehoboam/apps/portfolio && docker compose up -d --build
```

### 7. Configure DNS routes

```bash
cloudflared tunnel route dns homeserver portfolio.vincents.systems
cloudflared tunnel route dns homeserver files.vincents.systems
cloudflared tunnel route dns homeserver grafana.vincents.systems
cloudflared tunnel route dns homeserver ssh.vincents.systems
```

---

## Environment Variables

Copy `.env.example` to `.env` and fill in values. Never commit `.env`.

```bash
# Nextcloud
NEXTCLOUD_POSTGRES_PASSWORD=
NEXTCLOUD_ADMIN_PASSWORD=

# Grafana
GRAFANA_ADMIN_PASSWORD=
```

---

## Deployment Notes

**Start order matters.** The logging stack must start before other services — Docker's fluentd logging driver will discard logs if Fluent Bit isn't listening on `:24224` when containers start.

**Cloudflare tunnel restarts drop SSH.** Since SSH routes through the tunnel, restarting cloudflared terminates active sessions. Always have a local network fallback (direct IP SSH) before restarting the tunnel service.

**Persistent journal.** Systemd journal is configured to persist to `/var/log/journal` — logs survive reboots and are collected by Fluent Bit's systemd input plugin.

---

## Grafana Dashboards

Access Grafana at `https://grafana.vincents.systems`

**Rehoboam Dashboard** includes:

- **Log Volume by Service** — `sum by (container_name) (count_over_time({job="fluent-bit"}[1m]))` — time series of log throughput per container
- **SSH Login Attempts** — `{SYSLOG_IDENTIFIER="sshd-session"}` — raw log stream of every SSH session
- **Error Rate** — `sum(count_over_time({job="fluent-bit"} | json | detected_level="error" [1m]))` — error frequency over time
- **Cloudflare Tunnel Logs** — `{SYSLOG_IDENTIFIER="cloudflared"}` — tunnel connectivity events

---

## Portfolio Build Pipeline

The portfolio uses a multi-stage Docker build — Node.js never runs on the host.

```dockerfile
# Stage 1: Build
FROM node:20-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build          # tsc -b && vite build → /app/dist

# Stage 2: Serve
FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
```

The build context is `/srv/rehoboam/repos/portfolio` — source lives in `repos/`, deployment config lives in `apps/`. Rebuilding means `docker compose up -d --build` from `apps/portfolio/`.

---

## License

MIT
EOF