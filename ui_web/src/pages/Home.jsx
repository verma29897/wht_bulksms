import { Link } from 'react-router-dom';

export default function Home() {
  return (
    <div style={{ padding: '2rem' }}>
      <h2>Welcome</h2>
      <div style={{ display:'flex', gap:12 }}>
        <Link to="/register"><button>Embedded Signup</button></Link>
        <Link to="/auth/login"><button>Login</button></Link>
        <Link to="/auth/signup"><button>Sign up</button></Link>
        <Link to="/dashboard"><button>Dashboard</button></Link>
      </div>
    </div>
  );
}
