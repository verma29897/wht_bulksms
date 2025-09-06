import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { apiFetch } from "../../lib/api";

export default function Signup() {
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const onSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setSuccess("");
    try {
      // Replace with your backend signup endpoint
      const json = await apiFetch("/auth/signup", {
        method: "POST",
        body: JSON.stringify({ name, email, username, password }),
      });
      setSuccess("Account created. You can login now.");
      setTimeout(() => navigate("/auth/login"), 800);
    } catch (e) {
      setError(e.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-container">
      <form className="card form auth-card" onSubmit={onSubmit}>
        <div className="auth-title">Create account</div>
        <div className="auth-subtitle">Start using the dashboard</div>
        <div>
          <label>Name</label>
          <input value={name} onChange={(e) => setName(e.target.value)} placeholder="Your name" required />
        </div>
        <div>
          <label>Email</label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="you@example.com" required />
        </div>
        <div>
          <label>Username</label>
          <input value={username} onChange={(e) => setUsername(e.target.value)} placeholder="username" required />
        </div>
        <div>
          <label>Password</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="••••••••" required />
        </div>
        <div className="actions-row">
          <button className="btn block" type="submit" disabled={loading}>{loading ? "Creating..." : "Create account"}</button>
          {success && <span className="success">{success}</span>}
          {error && <span className="error">{error}</span>}
        </div>
        <div className="muted" style={{marginTop:8}}>Already have an account? <Link to="/auth/login">Login</Link></div>
      </form>
    </div>
  );
}


