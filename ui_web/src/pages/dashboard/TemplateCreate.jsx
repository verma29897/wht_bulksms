import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { apiFetch } from "../../lib/api";

export default function TemplateCreate() {
  const navigate = useNavigate();
  const [wabaId, setWabaId] = useState("");
  const [templateName, setTemplateName] = useState("");
  const [language, setLanguage] = useState("en_US");
  const [category, setCategory] = useState("MARKETING");
  const [headerType, setHeaderType] = useState("");
  const [headerContent, setHeaderContent] = useState("");
  const [bodyText, setBodyText] = useState("");
  const [footerText, setFooterText] = useState("");
  const [callButtonText, setCallButtonText] = useState("");
  const [phoneNumber, setPhoneNumber] = useState("");
  const [urlButtonText, setUrlButtonText] = useState("");
  const [websiteUrl, setWebsiteUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const onSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setSuccess("");
    try {
      const payload = {
        waba_id: wabaId.trim(),
        template_name: templateName.trim(),
        language: language.trim(),
        category,
        header_type: headerType,
        header_content: headerContent,
        body_text: bodyText,
        footer_text: footerText,
        call_button_text: callButtonText,
        phone_number: phoneNumber,
        url_button_text: urlButtonText,
        website_url: websiteUrl,
      };
      await apiFetch('/templates', { method: 'POST', body: JSON.stringify(payload) });
      setSuccess('Template created.');
      setTimeout(() => navigate('/dashboard/templates'), 800);
    } catch (e) {
      setError(e.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="card">
      <h3>Create Template</h3>
      <form className="form" onSubmit={onSubmit}>
        <div className="row">
          <div>
            <label>WABA ID</label>
            <input value={wabaId} onChange={(e) => setWabaId(e.target.value)} placeholder="318794741325676" required />
          </div>
          <div>
            <label>Template Name</label>
            <input value={templateName} onChange={(e) => setTemplateName(e.target.value)} placeholder="my_template" required />
          </div>
        </div>
        <div className="row">
          <div>
            <label>Language</label>
            <input value={language} onChange={(e) => setLanguage(e.target.value)} placeholder="en_US" />
          </div>
          <div>
            <label>Category</label>
            <select value={category} onChange={(e) => setCategory(e.target.value)}>
              <option value="MARKETING">MARKETING</option>
              <option value="UTILITY">UTILITY</option>
              <option value="AUTHENTICATION">AUTHENTICATION</option>
            </select>
          </div>
        </div>
        <div className="row">
          <div>
            <label>Header Type</label>
            <select value={headerType} onChange={(e) => setHeaderType(e.target.value)}>
              <option value="">None</option>
              <option value="headerText">Text</option>
              <option value="headerImage">Image</option>
              <option value="headerVideo">Video</option>
              <option value="headerDocument">Document</option>
              <option value="headerAudio">Audio</option>
            </select>
          </div>
          <div>
            <label>Header Content</label>
            <input value={headerContent} onChange={(e) => setHeaderContent(e.target.value)} placeholder={headerType === 'headerText' ? 'Header text' : 'header_handle from Media upload'} />
          </div>
        </div>
        <div>
          <label>Body Text</label>
          <textarea value={bodyText} onChange={(e) => setBodyText(e.target.value)} placeholder="Body content" required />
        </div>
        <div>
          <label>Footer Text</label>
          <input value={footerText} onChange={(e) => setFooterText(e.target.value)} placeholder="Optional footer" />
        </div>
        <div className="row">
          <div>
            <label>Call Button Text</label>
            <input value={callButtonText} onChange={(e) => setCallButtonText(e.target.value)} placeholder="Call Us" />
          </div>
          <div>
            <label>Phone Number</label>
            <input value={phoneNumber} onChange={(e) => setPhoneNumber(e.target.value)} placeholder="+911234567890" />
          </div>
        </div>
        <div className="row">
          <div>
            <label>URL Button Text</label>
            <input value={urlButtonText} onChange={(e) => setUrlButtonText(e.target.value)} placeholder="Visit" />
          </div>
          <div>
            <label>Website URL</label>
            <input value={websiteUrl} onChange={(e) => setWebsiteUrl(e.target.value)} placeholder="https://example.com" />
          </div>
        </div>
        <div className="actions-row">
          <button className="btn" type="submit" disabled={loading}>{loading ? 'Creating...' : 'Create Template'}</button>
          {success && <span className="success">{success}</span>}
          {error && <span className="error">{error}</span>}
        </div>
      </form>
    </div>
  );
}


