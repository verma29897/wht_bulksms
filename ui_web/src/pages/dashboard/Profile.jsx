import { useEffect, useState } from "react";
import { apiFetch } from "../../lib/api";

export default function Profile() {
  const [wabaId, setWabaId] = useState(localStorage.getItem('waba_id') || "");
  const [phoneId, setPhoneId] = useState(localStorage.getItem('phone_number_id') || "");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    const load = async () => {
      setLoading(true);
      setError("");
      try {
        const me = await apiFetch('/me');
        if (me?.waba_id) {
          setWabaId(me.waba_id);
          localStorage.setItem('waba_id', me.waba_id);
        }
        if (me?.phone_number_id) {
          setPhoneId(me.phone_number_id);
          localStorage.setItem('phone_number_id', me.phone_number_id);
        }
      } catch (e) {
        setError(e.message);
      } finally {
        setLoading(false);
      }
    };
    load();
  }, []);

  return (
    <div className="card">
      <h3>Profile</h3>
      <div className="muted" style={{marginBottom: 12}}>Your onboarded WhatsApp Business details</div>
      {error && <div className="error" style={{marginBottom:8}}>{error}</div>}
      <div className="form">
        <div>
          <label>WABA ID</label>
          <input value={wabaId} readOnly />
        </div>
        <div>
          <label>Phone Number ID</label>
          <input value={phoneId} readOnly />
        </div>
        <div className="muted">{loading ? 'Refreshing...' : 'Loaded'}</div>
      </div>
    </div>
  );
}


