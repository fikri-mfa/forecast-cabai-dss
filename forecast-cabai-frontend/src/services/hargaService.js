const BASE_URL = "http://localhost:9090";

export async function getAllHarga(token) {
  const res = await fetch(`${BASE_URL}/harga`, {
    headers: { "Authorization": `Bearer ${token}` },
  });
  const data = await res.json();
  if (!res.ok) throw new Error("Gagal ambil data harga");
  return data.data;
}

export async function createHarga({ tanggal, harga }, token) {
  const res = await fetch(`${BASE_URL}/harga`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`,
    },
    body: JSON.stringify({ tanggal, harga }),
  });
  if (!res.ok) throw new Error("Gagal tambah data harga");
}

export async function updateHarga(id, { tanggal, harga }, token) {
  const res = await fetch(`${BASE_URL}/harga/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`,
    },
    body: JSON.stringify({ tanggal, harga }),
  });
  if (!res.ok) throw new Error("Gagal update data harga");
}

export async function deleteHarga(id, token) {
  const res = await fetch(`${BASE_URL}/harga/${id}`, {
    method: "DELETE",
    headers: { "Authorization": `Bearer ${token}` },
  });
  if (!res.ok) throw new Error("Gagal hapus data harga");
}