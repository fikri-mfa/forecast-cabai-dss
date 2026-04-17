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