import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export default function EmbeddedSignup() {
  const navigate = useNavigate();

  useEffect(() => {
    window.fbAsyncInit = function () {
      window.FB.init({
        appId: import.meta.env.VITE_FB_APP_ID || "1050568876584504",
        cookie: true,
        xfbml: true,
        version: "v20.0",
      });
    };

    const script = document.createElement("script");
    script.src = "https://connect.facebook.net/en_US/sdk.js";
    script.async = true;
    script.defer = true;
    document.body.appendChild(script);
  }, []);

  const launchEmbeddedSignup = () => {
    const configId = import.meta.env.VITE_WHATSAPP_EMBEDDED_SIGNUP_CONFIG_ID;
    if (!configId) {
      alert("Missing VITE_WHATSAPP_EMBEDDED_SIGNUP_CONFIG_ID");
      return;
    }

    window.FB.login(
      (response) => {
        // Expecting response.authResponse.code
        const code = response?.authResponse?.code;
        if (code) {
          fetch(`/api/onboard/callback?auth_code=${encodeURIComponent(code)}`)
            .then(() => navigate("/register/success"))
            .catch((e) => {
              console.error(e);
              alert("Onboarding failed");
            });
        } else {
          alert("Signup canceled or failed");
        }
      },
      {
        config_id: configId,
        response_type: "code",
        override_default_response_type: true,
      }
    );
  };

  const fbReady = typeof window !== 'undefined' && !!window.FB;
  return (
    <div className="auth-container">
      <div className="card auth-card">
        <div className="auth-title">WhatsApp Embedded Signup</div>
        <div className="auth-subtitle">Onboard your WABA to start messaging</div>
        <button className="btn block" onClick={launchEmbeddedSignup} disabled={!fbReady}>
          {fbReady ? 'Start WhatsApp Embedded Signup' : 'Loading Facebook SDK...'}
        </button>
      </div>
    </div>
  );
}
