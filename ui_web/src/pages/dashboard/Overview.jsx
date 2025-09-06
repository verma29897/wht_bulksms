export default function Overview() {
  return (
    <div className="card-grid">
      <div className="card">
        <div className="muted">Messages Sent (today)</div>
        <div style={{ fontSize: 28, fontWeight: 700, marginTop: 6 }}>—</div>
      </div>
      <div className="card">
        <div className="muted">Templates</div>
        <div style={{ fontSize: 28, fontWeight: 700, marginTop: 6 }}>—</div>
      </div>
      <div className="card">
        <div className="muted">Media Uploads</div>
        <div style={{ fontSize: 28, fontWeight: 700, marginTop: 6 }}>—</div>
      </div>
      <div className="card">
        <div className="muted">Accounts Onboarded</div>
        <div style={{ fontSize: 28, fontWeight: 700, marginTop: 6 }}>—</div>
      </div>
    </div>
  );
}


