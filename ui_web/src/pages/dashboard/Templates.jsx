import { useEffect, useState } from "react";
import { apiFetch } from "../../lib/api";

export default function Templates() {
  const [wabaId, setWabaId] = useState("");
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const load = async () => {
    if (!wabaId) return;
    setLoading(true);
    setError("");
    try {
      const json = await apiFetch(`/templates/${wabaId}`);
      setItems(json.templates || []);
    } catch (e) {
      setError(e.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    // Do not auto-expose IDs here; user can view them in Profile
  }, []);

  return (
    <div className="card">
      <h3>Templates</h3>
      <div className="actions-row" style={{marginBottom: 8}}>
        <a className="btn" href="/dashboard/templates/create">Create Template</a>
      </div>
      <div className="row">
        <div>
          <label>WABA ID</label>
          <input value={wabaId} onChange={(e) => setWabaId(e.target.value)} placeholder="e.g. 318794741325676" />
        </div>
        <div style={{alignSelf:'end'}}>
          <button className="btn" onClick={load} disabled={loading}>{loading ? "Loading..." : "Load"}</button>
        </div>
      </div>
      {error && <div className="error" style={{marginTop:8}}>{error}</div>}
      <div style={{marginTop:12}}>
        <table className="table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Language</th>
              <th>Status</th>
              <th>Category</th>
            </tr>
          </thead>
          <tbody>
            {items.map((t) => (
              <tr key={t.template_id}>
                <td>{t.template_name}</td>
                <td>{t.template_language}</td>
                <td>{t.status}</td>
                <td>{t.category}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}


