const BASE_URL = "http://localhost:9090";

export async function runForecast({ alpha, beta, gamma, periods, season_length }, token) {
  const res = await fetch(`${BASE_URL}/forecast`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`,
    },
    body: JSON.stringify({ alpha, beta, gamma, periods, season_length }),
  });
  const data = await res.json();
  if (!res.ok) throw new Error(data.message || "Forecast gagal");
  return data;
}

export async function getForecastHistory(token) {
  const res = await fetch(`${BASE_URL}/forecast/history`, {
    headers: { "Authorization": `Bearer ${token}` },
  });
  const data = await res.json();
  if (!res.ok) throw new Error(data.message || "Gagal ambil history");
  return data;
}

export async function getDashboardStats(token) {
  const res = await fetch(`${BASE_URL}/dashboard/stats`, {
    headers: { "Authorization": `Bearer ${token}` },
  });
  const data = await res.json();
  if (!res.ok) throw new Error("Gagal ambil stats");
  return data.data;
}

export async function getDashboardChart(token) {
  const res = await fetch(`${BASE_URL}/dashboard/chart`, {
    headers: { "Authorization": `Bearer ${token}` },
  });
  const data = await res.json();
  if (!res.ok) throw new Error("Gagal ambil chart");
  return data.data;
}

// GET /dashboard/chart/all
// Backend return semua data historis dengan fitted TES values
// Tidak ada future dates — hanya data aktual + nilai fitted model
export async function getAllChartData(token) {
  const res = await fetch(`${BASE_URL}/dashboard/chart/all`, {
    headers: { "Authorization": `Bearer ${token}` },
  });
  
  // Handle jika response bukan JSON (misal 404 HTML)
  const text = await res.text();
  let data;
  try {
    data = JSON.parse(text);
  } catch {
    throw new Error("Endpoint /dashboard/chart/all belum tersedia di backend");
  }
  
  if (!res.ok) throw new Error("Gagal ambil semua data chart");
  const rows = (data.data || []).filter((d) => d.tanggal && d.aktual > 0);
  rows.sort((a, b) => new Date(a.tanggal) - new Date(b.tanggal));
  return rows;
}