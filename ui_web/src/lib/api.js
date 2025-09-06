export function getApiBase() {
  // Prefer env base; otherwise use Vite proxy via /api
  return import.meta.env.VITE_API_BASE || '';
}

export async function apiFetch(path, options = {}) {
  const token = localStorage.getItem('auth_token');
  const headers = new Headers(options.headers || {});
  if (token) headers.set('Authorization', `Bearer ${token}`);
  if (!headers.has('Content-Type') && options.body && !(options.body instanceof FormData)) {
    headers.set('Content-Type', 'application/json');
  }
  const base = getApiBase();
  let url = '';
  if (base) {
    url = `${base}${path}`;
  } else {
    url = path.startsWith('/api') ? path : `/api${path}`;
  }
  const res = await fetch(url, { ...options, headers });
  const contentType = res.headers.get('content-type') || '';
  const body = contentType.includes('application/json') ? await res.json() : await res.text();
  if (!res.ok) {
    const errMsg = (body && body.error) ? body.error : (typeof body === 'string' ? body : 'Request failed');
    if (res.status === 401) {
      // Optional: redirect to login
    }
    throw new Error(errMsg);
  }
  return body;
}


