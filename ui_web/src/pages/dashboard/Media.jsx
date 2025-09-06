import { useState } from "react";
import { getApiBase } from "../../lib/api";

export default function Media() {
  const [phoneNumberId, setPhoneNumberId] = useState("");
  const [file, setFile] = useState(null);
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const upload = async () => {
    if (!file || !phoneNumberId) return;
    setLoading(true);
    setResult(null);
    setError("");
    try {
      const form = new FormData();
      form.append("phone_number_id", phoneNumberId);
      form.append("file", file);
      const res = await fetch(`${getApiBase()}/upload/header`, { method: "POST", body: form, headers: { Authorization: `Bearer ${localStorage.getItem('auth_token') || ''}` } });
      const json = await res.json();
      if (!res.ok) throw new Error(json.error || "Upload failed");
      setResult(json);
    } catch (e) {
      setError(e.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="card">
      <h3>Media Upload</h3>
      <div className="row">
        <div>
          <label>Phone Number ID</label>
          <input value={phoneNumberId} onChange={(e) => setPhoneNumberId(e.target.value)} />
        </div>
        <div>
          <label>File</label>
          <input type="file" onChange={(e) => setFile(e.target.files?.[0] || null)} />
        </div>
      </div>
      <div className="actions-row" style={{marginTop: 8}}>
        <button className="btn" onClick={upload} disabled={loading || !file || !phoneNumberId}>{loading ? "Uploading..." : "Upload"}</button>
        {result && <span className="success">Uploaded</span>}
        {error && <span className="error">{error}</span>}
      </div>
      {result && <pre style={{marginTop:12}}>{JSON.stringify(result, null, 2)}</pre>}
    </div>
  );
}


