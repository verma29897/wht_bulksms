import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Register from './pages/Register';
import RegisterSuccess from './pages/RegisterSuccess';
import DashboardLayout from './layouts/DashboardLayout';
import Overview from './pages/dashboard/Overview';
import BulkSend from './pages/dashboard/BulkSend';
import Templates from './pages/dashboard/Templates';
import TemplateCreate from './pages/dashboard/TemplateCreate';
import Media from './pages/dashboard/Media';
import Accounts from './pages/dashboard/Accounts';
import Settings from './pages/dashboard/Settings';
import Profile from './pages/dashboard/Profile';
import Login from './pages/auth/Login';
import Signup from './pages/auth/Signup';
import ProtectedRoute from './components/ProtectedRoute';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/register/success" element={<RegisterSuccess />} />
        <Route path="/auth/login" element={<Login />} />
        <Route path="/auth/signup" element={<Signup />} />
        <Route element={<ProtectedRoute />}>
          <Route path="/dashboard" element={<DashboardLayout />}>
            <Route index element={<Overview />} />
            <Route path="profile" element={<Profile />} />
            <Route path="bulk" element={<BulkSend />} />
            <Route path="templates" element={<Templates />} />
            <Route path="templates/create" element={<TemplateCreate />} />
            <Route path="media" element={<Media />} />
            <Route path="accounts" element={<Accounts />} />
            {false && <Route path="settings" element={<Settings />} />}
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
