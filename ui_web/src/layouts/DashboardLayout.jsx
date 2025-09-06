import { Link, Outlet, useLocation } from "react-router-dom";
import "../assets/dashboard.css";

export default function DashboardLayout() {
  const location = useLocation();
  const isActive = (path) => (location.pathname.startsWith(path) ? "active" : "");

  return (
    <div className="dash-container">
      <aside className="dash-sidebar">
        <div className="brand">BulkSMS</div>
        <nav>
          <Link className={`nav-link ${isActive("/dashboard")}`} to="/dashboard">Overview</Link>
          <Link className={`nav-link ${isActive("/dashboard/profile")}`} to="/dashboard/profile">Profile</Link>
          <Link className={`nav-link ${isActive("/dashboard/bulk")}`} to="/dashboard/bulk">Bulk Send</Link>
          <Link className={`nav-link ${isActive("/dashboard/templates")}`} to="/dashboard/templates">Templates</Link>
          <Link className={`nav-link ${isActive("/dashboard/media")}`} to="/dashboard/media">Media</Link>
          <Link className={`nav-link ${isActive("/dashboard/accounts")}`} to="/dashboard/accounts">Accounts</Link>
        </nav>
      </aside>
      <main className="dash-main">
        <header className="dash-topbar">
          <div className="title">WhatsApp Panel</div>
          <div className="actions">
            <a className="btn" href="/register">Embedded Signup</a>
            <button className="btn secondary" onClick={() => { localStorage.removeItem('auth_token'); window.location.href = '/auth/login'; }}>Logout</button>
          </div>
        </header>
        <section className="dash-content">
          <Outlet />
        </section>
      </main>
    </div>
  );
}


