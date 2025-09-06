import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { apiFetch } from "../../lib/api";

export default function Login() {
  const navigate = useNavigate();
  const [identifier, setIdentifier] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const onSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    try {
      // Replace with your backend auth endpoint
      const json = await apiFetch("/auth/login", {
        method: "POST",
        body: JSON.stringify({ identifier, password }),
      });
      // store token/session if returned
      if (json.token) localStorage.setItem("auth_token", json.token);
      // fetch and cache onboarded IDs for prefills
      try {
        const me = await apiFetch("/me");
        if (me) {
          if (me.waba_id) localStorage.setItem("waba_id", me.waba_id);
          if (me.phone_number_id) localStorage.setItem("phone_number_id", me.phone_number_id);
        }
      } catch (_) {}
      navigate("/dashboard");
    } catch (e) {
      setError(e.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-container">
      <form className="card form auth-card" onSubmit={onSubmit}>
        <div className="auth-title">Welcome back</div>
        <div className="auth-subtitle">Sign in to your account</div>
        <div>
          <label>Email or Username</label>
          <input value={identifier} onChange={(e) => setIdentifier(e.target.value)} placeholder="you@example.com or username" required />
        </div>
        <div>
          <label>Password</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="••••••••" required />
        </div>
        <div className="actions-row">
          <button className="btn block" type="submit" disabled={loading}>{loading ? "Signing in..." : "Login"}</button>
          {error && <span className="error">{error}</span>}
        </div>
        <div className="muted" style={{marginTop:8}}>Don't have an account? <Link to="/auth/signup">Sign up</Link></div>
        <div className="muted" style={{marginTop:4}}><Link to="/register">Or run WhatsApp Embedded Signup</Link></div>
      </form>
    </div>
  );
}


