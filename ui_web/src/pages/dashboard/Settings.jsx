import { useState } from "react";

export default function Settings() {
  const [apiBase, setApiBase] = useState("http://localhost:8080");
  const [fbAppId, setFbAppId] = useState(import.meta.env.VITE_FB_APP_ID || "");
  const [configId, setConfigId] = useState(import.meta.env.VITE_WHATSAPP_EMBEDDED_SIGNUP_CONFIG_ID || "");

  return (
    <div className="card">
      <h3>Settings</h3>
      <div className="form">
        <div className="row">
          <div>
            <label>API Base URL</label>
            <input value={apiBase} onChange={(e) => setApiBase(e.target.value)} />
          </div>
          <div>
            <label>FB App ID</label>
            <input value={fbAppId} onChange={(e) => setFbAppId(e.target.value)} />
          </div>
        </div>
        <div className="row">
          <div>
            <label>Embedded Signup Config ID</label>
            <input value={configId} onChange={(e) => setConfigId(e.target.value)} />
          </div>
        </div>
        <div className="muted">Environment-driven; this page is a placeholder if you later persist settings client-side.</div>
      </div>
    </div>
  );
}


