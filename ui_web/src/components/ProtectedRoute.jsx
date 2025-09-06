import { Navigate, Outlet } from 'react-router-dom';

export default function ProtectedRoute() {
  const token = localStorage.getItem('auth_token');
  if (!token) return <Navigate to="/auth/login" replace />;
  return <Outlet />;
}


