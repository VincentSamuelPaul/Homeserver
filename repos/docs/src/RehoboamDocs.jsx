import { useState, useEffect } from "react";

const mono = { fontFamily: "'JetBrains Mono', 'Fira Code', 'Courier New', monospace" };

const styles = `
  @import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@300;400;500;700&display=swap');

  * { box-sizing: border-box; margin: 0; padding: 0; }

  body {
    background: #0d0d0d;
    color: #d4d4d4;
    font-family: 'JetBrains Mono', 'Fira Code', monospace;
    font-size: 13px;
    line-height: 1.7;
    -webkit-font-smoothing: antialiased;
  }

  .container {
    max-width: 860px;
    margin: 0 auto;
    padding: 48px 32px 96px;
  }

  .divider {
    border: none;
    border-top: 1px solid #222;
    margin: 36px 0;
  }

  .section-label {
    color: #555;
    font-size: 11px;
    font-weight: 400;
    letter-spacing: 0.08em;
    margin-bottom: 20px;
  }

  .section-label span {
    color: #888;
  }

  h1 { font-size: 22px; font-weight: 500; color: #e8e8e8; }
  h2 { font-size: 14px; font-weight: 500; color: #e8e8e8; margin-bottom: 12px; }

  .muted { color: #555; }
  .dim { color: #888; }
  .bright { color: #e8e8e8; }

  .grid-2 {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1px;
    background: #222;
    border: 1px solid #222;
    margin-bottom: 24px;
  }

  .grid-cell {
    background: #0d0d0d;
    padding: 16px 20px;
  }

  .grid-cell .label { color: #555; font-size: 11px; margin-bottom: 4px; }
  .grid-cell .value { color: #d4d4d4; font-size: 13px; }

  .service-row {
    display: grid;
    grid-template-columns: 160px 1fr 200px;
    gap: 0;
    border-bottom: 1px solid #1a1a1a;
    padding: 10px 0;
    align-items: start;
  }

  .service-row:last-child { border-bottom: none; }
  .service-row .name { color: #e8e8e8; }
  .service-row .desc { color: #888; font-size: 12px; padding-right: 24px; }
  .service-row .url { color: #555; font-size: 11px; text-align: right; }

  .block {
    background: #111;
    border: 1px solid #1e1e1e;
    padding: 16px 20px;
    margin-bottom: 12px;
  }

  .block .block-title { color: #e8e8e8; margin-bottom: 6px; font-weight: 500; }
  .block .block-body { color: #777; font-size: 12px; line-height: 1.8; }

  .pipe-row {
    display: flex;
    align-items: center;
    gap: 0;
    margin: 20px 0;
  }

  .pipe-node {
    background: #111;
    border: 1px solid #222;
    padding: 10px 18px;
    font-size: 12px;
    color: #d4d4d4;
    white-space: nowrap;
  }

  .pipe-arrow {
    color: #333;
    padding: 0 8px;
    font-size: 16px;
    flex-shrink: 0;
  }

  .tag {
    display: inline-block;
    border: 1px solid #2a2a2a;
    color: #666;
    font-size: 11px;
    padding: 2px 8px;
    margin: 3px 3px 3px 0;
  }

  .cursor {
    display: inline-block;
    width: 10px;
    height: 16px;
    background: #d4d4d4;
    margin-left: 4px;
    vertical-align: middle;
    animation: blink 1.1s step-end infinite;
  }

  @keyframes blink { 0%,100%{opacity:1} 50%{opacity:0} }

  .uptime { color: #e8e8e8; font-variant-numeric: tabular-nums; }

  .header-meta {
    display: flex;
    justify-content: space-between;
    align-items: flex-end;
    margin-bottom: 8px;
  }

  .fs-tree {
    color: #666;
    font-size: 12px;
    line-height: 2;
    padding-left: 8px;
  }

  .fs-tree .fs-dir { color: #d4d4d4; }
  .fs-tree .fs-comment { color: #3a3a3a; }

  .roadmap-item {
    display: flex;
    gap: 12px;
    padding: 8px 0;
    border-bottom: 1px solid #1a1a1a;
    align-items: flex-start;
    font-size: 12px;
  }

  .roadmap-item:last-child { border-bottom: none; }
  .roadmap-item .status { color: #333; flex-shrink: 0; }
  .roadmap-item .item-text { color: #888; }
  .roadmap-item .item-tag { color: #444; margin-left: auto; flex-shrink: 0; }

  .security-item {
    display: flex;
    gap: 16px;
    padding: 8px 0;
    border-bottom: 1px solid #1a1a1a;
    font-size: 12px;
  }

  .security-item:last-child { border-bottom: none; }
  .security-item .layer { color: #d4d4d4; width: 140px; flex-shrink: 0; }
  .security-item .impl { color: #777; }

  @media (max-width: 600px) {
    .container { padding: 24px 16px 64px; }
    .grid-2 { grid-template-columns: 1fr; }
    .service-row { grid-template-columns: 1fr; gap: 4px; }
    .service-row .url { text-align: left; }
    .header-meta { flex-direction: column; align-items: flex-start; gap: 8px; }
  }
`;

function useUptime(startMs) {
  const [elapsed, setElapsed] = useState(Date.now() - startMs);
  useEffect(() => {
    const id = setInterval(() => setElapsed(Date.now() - startMs), 1000);
    return () => clearInterval(id);
  }, [startMs]);

  const s = Math.floor(elapsed / 1000);
  const d = Math.floor(s / 86400);
  const h = Math.floor((s % 86400) / 3600);
  const m = Math.floor((s % 3600) / 60);
  const sec = s % 60;
  return `${d}d ${String(h).padStart(2,"0")}h ${String(m).padStart(2,"0")}m ${String(sec).padStart(2,"0")}s`;
}

const BOOT_TIME = new Date("2026-07-01T00:00:00+05:30").getTime();

export default function RehoboamDocs() {
  const uptime = useUptime(BOOT_TIME);

  return (
    <>
      <style>{styles}</style>
      <div className="container">

        {/* Header */}
        <div className="header-meta">
          <span className="muted" style={{ fontSize: 11 }}>rehoboam · infrastructure docs · rev 1.0</span>
          <span className="muted" style={{ fontSize: 11 }}>vincents.systems</span>
        </div>
        <hr className="divider" style={{ marginTop: 0, marginBottom: 24 }} />

        <h1>
          rehoboam<span className="cursor" />
        </h1>
        <p style={{ color: "#666", marginTop: 8, marginBottom: 32, fontSize: 12 }}>
          Production-grade self-hosted homeserver · Arch Linux · Docker · Cloudflare Zero Trust
        </p>

        {/* Status */}
        <div className="grid-2">
          <div className="grid-cell">
            <div className="label">host</div>
            <div className="value">forge</div>
          </div>
          <div className="grid-cell">
            <div className="label">os</div>
            <div className="value">arch linux x86_64</div>
          </div>
          <div className="grid-cell">
            <div className="label">cpu</div>
            <div className="value">AMD Ryzen 5 PRO 1500 · 8t · 3.5GHz</div>
          </div>
          <div className="grid-cell">
            <div className="label">ram</div>
            <div className="value">16 GB DDR4</div>
          </div>
          <div className="grid-cell">
            <div className="label">storage</div>
            <div className="value">single SSD · 201 GB home</div>
          </div>
          <div className="grid-cell">
            <div className="label">uptime (session)</div>
            <div className="value uptime">{uptime}</div>
          </div>
        </div>

        <hr className="divider" />

        {/* Services */}
        <div className="section-label">[ <span>SERVICES</span> ]</div>

        <div style={{ borderTop: "1px solid #1a1a1a" }}>
          {[
            { name: "portfolio", desc: "Personal portfolio — React + Vite + Nginx, multi-stage Docker build", url: "portfolio.vincents.systems", port: ":8081" },
            { name: "nextcloud", desc: "Self-hosted file storage — Apache + PostgreSQL + Redis", url: "files.vincents.systems", port: ":8082" },
            { name: "grafana", desc: "Observability dashboards — log volume, SSH, error rate, tunnel health", url: "grafana.vincents.systems", port: ":3000" },
            { name: "loki", desc: "Log aggregation — TSDB storage, LogQL queries, retained forever", url: "internal", port: ":3100" },
            { name: "fluent-bit", desc: "Log collection — Docker fluentd-driver + systemd journal input", url: "internal", port: ":24224" },
            { name: "sshd", desc: "Remote access — ed25519 keys only, password auth disabled", url: "ssh.vincents.systems", port: ":22" },
            { name: "postgres:16", desc: "Nextcloud metadata database — isolated on nextcloud network", url: "internal", port: ":5432" },
            { name: "redis", desc: "Nextcloud caching layer — alpine image, memory-only", url: "internal", port: ":6379" },
          ].map(s => (
            <div className="service-row" key={s.name}>
              <span className="name">{s.name}</span>
              <span className="desc">{s.desc}</span>
              <span className="url">{s.url} <span className="muted">{s.port}</span></span>
            </div>
          ))}
        </div>

        <hr className="divider" />

        {/* Architecture */}
        <div className="section-label">[ <span>TRAFFIC PATH</span> ]</div>

        <div className="pipe-row">
          <div className="pipe-node">internet</div>
          <div className="pipe-arrow">→</div>
          <div className="pipe-node">cloudflare edge<br /><span style={{ color: "#555", fontSize: 11 }}>DDoS · TLS · DNS</span></div>
          <div className="pipe-arrow">→</div>
          <div className="pipe-node">cloudflared<br /><span style={{ color: "#555", fontSize: 11 }}>:443 quic · outbound-only</span></div>
          <div className="pipe-arrow">→</div>
          <div className="pipe-node">services<br /><span style={{ color: "#555", fontSize: 11 }}>no open ports</span></div>
        </div>

        <p style={{ color: "#555", fontSize: 11, marginTop: 8 }}>
          All inbound traffic routes through Cloudflare Zero Trust tunnel. The router has zero open ports.
          Cloudflare terminates TLS at the edge; traffic reaches the server over an outbound QUIC connection established by cloudflared.
        </p>

        <hr className="divider" />

        {/* Log pipeline */}
        <div className="section-label">[ <span>LOG PIPELINE</span> ]</div>

        <div className="pipe-row">
          <div className="pipe-node" style={{ fontSize: 11 }}>docker containers<br /><span style={{ color: "#555" }}>fluentd-driver :24224</span></div>
          <div className="pipe-arrow">→</div>
          <div className="pipe-node" style={{ fontSize: 11 }}>systemd journal<br /><span style={{ color: "#555" }}>SSH · cloudflared</span></div>
          <div className="pipe-arrow">→</div>
          <div className="pipe-node" style={{ fontSize: 11 }}>fluent-bit v5<br /><span style={{ color: "#555" }}>filter · tag · ship</span></div>
          <div className="pipe-arrow">→</div>
          <div className="pipe-node" style={{ fontSize: 11 }}>loki<br /><span style={{ color: "#555" }}>tsdb chunks</span></div>
          <div className="pipe-arrow">→</div>
          <div className="pipe-node" style={{ fontSize: 11 }}>grafana<br /><span style={{ color: "#555" }}>LogQL · alerts</span></div>
        </div>

        <p style={{ color: "#555", fontSize: 11, marginTop: 8 }}>
          Docker daemon configured globally with the fluentd logging driver — every container ships logs automatically with no per-service config.
          Fluent Bit tags logs by container name and injects host/environment metadata before shipping to Loki.
        </p>

        <hr className="divider" />

        {/* Filesystem */}
        <div className="section-label">[ <span>FILESYSTEM</span> ]</div>

        <div className="fs-tree">
          <div><span className="fs-dir">/srv/rehoboam/</span></div>
          <div>├── <span className="fs-dir">apps/</span> <span className="fs-comment">          # docker-compose per service</span></div>
          <div>│   ├── portfolio/</div>
          <div>│   ├── nextcloud/</div>
          <div>│   └── logging/</div>
          <div>├── <span className="fs-dir">repos/</span> <span className="fs-comment">         # source code · local clone</span></div>
          <div>│   └── portfolio/</div>
          <div>├── <span className="fs-dir">infra/</span> <span className="fs-comment">         # shared config</span></div>
          <div>│   ├── cloudflared/</div>
          <div>│   ├── fluent-bit/</div>
          <div>│   └── loki-config.yml</div>
          <div>├── <span className="fs-dir">data/</span> <span className="fs-comment">          # persistent volumes · gitignored</span></div>
          <div>│   ├── loki/ <span className="fs-comment">      # ~26MB · chunks · index · WAL</span></div>
          <div>│   ├── grafana/</div>
          <div>│   └── nextcloud/</div>
          <div>├── <span className="fs-dir">backups/</span> <span className="fs-comment">       # scheduled snapshots · gitignored</span></div>
          <div>└── .env <span className="fs-comment">           # secrets · never committed</span></div>
        </div>

        <hr className="divider" />

        {/* Docker networks */}
        <div className="section-label">[ <span>DOCKER NETWORKS</span> ]</div>

        {[
          { name: "logging_logging", members: "loki · fluent-bit · grafana", note: "log pipeline isolation" },
          { name: "nextcloud_nextcloud", members: "nextcloud · nextcloud-db · nextcloud-redis", note: "file stack isolation" },
          { name: "portfolio_default", members: "portfolio", note: "standalone nginx" },
        ].map(n => (
          <div className="block" key={n.name}>
            <div className="block-title">{n.name}</div>
            <div className="block-body">{n.members} · <span style={{ color: "#444" }}>{n.note}</span></div>
          </div>
        ))}

        <hr className="divider" />

        {/* Security */}
        <div className="section-label">[ <span>SECURITY</span> ]</div>

        <div style={{ borderTop: "1px solid #1a1a1a" }}>
          {[
            { layer: "SSH", impl: "ed25519 key auth only · root login disabled · password auth disabled" },
            { layer: "Network exposure", impl: "zero open inbound ports · all traffic via outbound Cloudflare tunnel" },
            { layer: "DDoS", impl: "Cloudflare edge absorbs attacks before reaching origin" },
            { layer: "Container isolation", impl: "separate Docker bridge network per service stack" },
            { layer: "Secrets", impl: ".env file · gitignored · never baked into images or committed" },
            { layer: "Logging", impl: "systemd journal persisted to disk · collected and retained in Loki" },
          ].map(s => (
            <div className="security-item" key={s.layer}>
              <span className="layer">{s.layer}</span>
              <span className="impl">{s.impl}</span>
            </div>
          ))}
        </div>

        <hr className="divider" />

        {/* Stack */}
        <div className="section-label">[ <span>STACK</span> ]</div>

        <div style={{ marginBottom: 24 }}>
          {["Arch Linux", "Docker", "Docker Compose", "Cloudflare Zero Trust", "cloudflared",
            "Nginx", "Apache", "PostgreSQL 16", "Redis", "Fluent Bit v5", "Grafana Loki",
            "Grafana", "React", "Vite", "TypeScript", "Tailwind CSS", "Node.js 20"].map(t => (
            <span className="tag" key={t}>{t}</span>
          ))}
        </div>

        <hr className="divider" />

        {/* Roadmap */}
        <div className="section-label">[ <span>ROADMAP</span> ]</div>

        <div style={{ borderTop: "1px solid #1a1a1a" }}>
          {[
            { done: false, text: "Tailscale VPN — direct peer-to-peer access without tunnel dependency", tag: "networking" },
            { done: false, text: "Prometheus + Node Exporter — metrics to complement the logging stack", tag: "observability" },
            { done: false, text: "GitHub Actions CI/CD — auto-deploy portfolio on push to main", tag: "automation" },
            { done: false, text: "Rclone — automated encrypted backup to Backblaze B2", tag: "storage" },
            { done: false, text: "k3s — lightweight Kubernetes for orchestration", tag: "infra" },
            { done: false, text: "WhatsApp bot — file backup via Baileys → Nextcloud WebDAV", tag: "integration" },
            { done: false, text: "Dedicated GPU — low-profile RDNA for local LLM inference", tag: "hardware" },
          ].map((r, i) => (
            <div className="roadmap-item" key={i}>
              <span className="status">{r.done ? "✓" : "○"}</span>
              <span className="item-text">{r.text}</span>
              <span className="item-tag">{r.tag}</span>
            </div>
          ))}
        </div>

        <hr className="divider" />

        {/* Footer */}
        <div style={{ display: "flex", justifyContent: "space-between", color: "#333", fontSize: 11 }}>
          <span>github.com/VincentSamuelPaul/Homeserver</span>
          <span>MIT</span>
        </div>

      </div>
    </>
  );
}
