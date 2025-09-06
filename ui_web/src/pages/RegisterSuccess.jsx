import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { apiFetch } from "../lib/api";

const RegisterSuccess = () => {
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    const params = new URLSearchParams(location.search);

    const waba_id = params.get("waba_id");
    const access_token = params.get("access_token");
    const phone_number = params.get("phone_number");

    if (waba_id && access_token) {
      apiFetch("/store-onboarding", {
        method: "POST",
        body: JSON.stringify({ waba_id, access_token, phone_number })
      })
        .then(() => navigate('/dashboard'))
        .catch((err) => console.error("Error:", err));
    }
  }, [location]);

  return <h2>Signup Completed. You can close this tab.</h2>;
};

export default RegisterSuccess;
