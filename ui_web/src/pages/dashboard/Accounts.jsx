export default function Accounts() {
  return (
    <div className="card">
      <h3>Accounts</h3>
      <div className="muted">After Embedded Signup, accounts are stored in the backend DB. Build a list view by adding an API endpoint to fetch accounts.</div>
      <div style={{marginTop:12}}>
        <a className="btn" href="/register">Run Embedded Signup</a>
      </div>
    </div>
  );
}


