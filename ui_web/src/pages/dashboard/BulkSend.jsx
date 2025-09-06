import { useEffect, useState } from "react";
import { apiFetch } from "../../lib/api";

export default function BulkSend() {
  const [phoneNumberId, setPhoneNumberId] = useState("");
  const [templateName, setTemplateName] = useState("");
  const [language, setLanguage] = useState("en_US");
  const [contacts, setContacts] = useState("");
  const [sending, setSending] = useState(false);
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    // Do not auto-expose IDs here; user can view them in Profile
  }, []);

  const onSubmit = async (e) => {
    e.preventDefault();
    setSending(true);
    setError("");
    setResult(null);
    try {
      const body = {
        phone_number_id: phoneNumberId.trim(),
        template_name: templateName.trim(),
        language: language.trim(),
        media_type: "", // optional
        media_id: null,
        contact_list: contacts.split(/\r?\n|,\s*/).map((s) => s.trim()).filter(Boolean),
      };
      const json = await apiFetch("/send", { method: "POST", body: JSON.stringify(body) });
      setResult(json);
    } catch (err) {
      setError(err.message);
    } finally {
      setSending(false);
    }
  };

  return (
    <div className="card">
      <h3>Bulk Send</h3>
      <form className="form" onSubmit={onSubmit}>
        <div className="row">
          <div>
            <label>Phone Number ID</label>
            <input value={phoneNumberId} onChange={(e) => setPhoneNumberId(e.target.value)} placeholder="e.g. 1234567890" />
          </div>
          <div>
            <label>Language</label>
            <input value={language} onChange={(e) => setLanguage(e.target.value)} placeholder="en_US" />
          </div>
        </div>
        <div className="row">
          <div>
            <label>Template Name</label>
            <input value={templateName} onChange={(e) => setTemplateName(e.target.value)} placeholder="template_name" />
          </div>
        </div>
        <div>
          <label>Contacts (comma or newline separated)</label>
          <textarea value={contacts} onChange={(e) => setContacts(e.target.value)} placeholder="919000000000, 919111111111" />
        </div>
        <div className="actions-row">
          <button className="btn" disabled={sending} type="submit">{sending ? "Sending..." : "Send"}</button>
          {result && <span className="success">Sent successfully</span>}
          {error && <span className="error">{error}</span>}
        </div>
      </form>
    </div>
  );
}


