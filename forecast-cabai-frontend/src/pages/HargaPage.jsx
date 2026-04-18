import { useEffect, useState } from "react";
import { useAuth } from "../context/AuthContext";
import { getAllHarga, createHarga, updateHarga, deleteHarga } from "../services/hargaService";

export default function HargaPage() {
  const { token } = useAuth();
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [modal, setModal] = useState(null);
  const [selected, setSelected] = useState(null);
  const [form, setForm] = useState({ tanggal: "", harga: "" });
  const [error, setError] = useState("");
  const [saving, setSaving] = useState(false);

  async function load() {
    try {
      const res = await getAllHarga(token);
      setData(res || []);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => { load(); }, []);

  function openAdd() {
    setForm({ tanggal: "", harga: "" });
    setError("");
    setModal("add");
  }

  function openEdit(row) {
    setSelected(row);
    setForm({ tanggal: row.tanggal?.substring(0, 10), harga: row.harga });
    setError("");
    setModal("edit");
  }

  async function handleSave() {
    if (!form.tanggal || !form.harga) { setError("Semua field wajib diisi"); return; }
    setSaving(true);
    try {
      if (modal === "add") await createHarga({ tanggal: form.tanggal, harga: parseFloat(form.harga) }, token);
      else await updateHarga(selected.id, { tanggal: form.tanggal, harga: parseFloat(form.harga) }, token);
      setModal(null);
      load();
    } catch (err) {
      setError(err.message);
    } finally {
      setSaving(false);
    }
  }

  async function handleDelete(id) {
    if (!confirm("Hapus data ini?")) return;
    try { await deleteHarga(id, token); load(); }
    catch (err) { alert(err.message); }
  }

  return (
    <div className="p-4 sm:p-6 lg:p-8">
      {/* Header */}
      <div className="flex items-start sm:items-center justify-between gap-3 mb-6 sm:mb-8">
        <div>
          <h1 className="text-xl sm:text-2xl font-bold text-gray-800 dark:text-white">Kelola Data Harga</h1>
          <p className="text-sm text-gray-400 mt-1">Tambah, edit, atau hapus data harga cabai</p>
        </div>
        <button
          onClick={openAdd}
          className="flex-shrink-0 px-3 sm:px-4 py-2 bg-green-600 hover:bg-green-700 text-white text-sm font-semibold rounded-xl transition-all"
        >
          + Tambah
        </button>
      </div>

      {/* Tabel dengan scroll horizontal di mobile */}
      <div className="bg-white dark:bg-gray-900 rounded-2xl shadow-sm border border-gray-100 dark:border-gray-800 overflow-hidden">
        {loading ? (
          <p className="text-sm text-gray-400 text-center py-10">Memuat data...</p>
        ) : (
          <div className="overflow-x-auto">
            <table className="w-full text-sm min-w-[400px]">
              <thead>
                <tr className="bg-gray-50 dark:bg-gray-800 border-b border-gray-100 dark:border-gray-700">
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">No</th>
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Tanggal</th>
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Harga (Rp)</th>
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Aksi</th>
                </tr>
              </thead>
              <tbody>
                {data.map((row, i) => (
                  <tr key={row.id} className="border-b border-gray-50 dark:border-gray-800 hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors">
                    <td className="px-4 sm:px-6 py-3 text-gray-400">{i + 1}</td>
                    <td className="px-4 sm:px-6 py-3 text-gray-700 dark:text-gray-300 whitespace-nowrap">{row.tanggal?.substring(0, 10)}</td>
                    <td className="px-4 sm:px-6 py-3 text-gray-700 dark:text-gray-300 font-medium whitespace-nowrap">
                      {row.harga.toLocaleString("id-ID")}
                    </td>
                    <td className="px-4 sm:px-6 py-3">
                      <div className="flex gap-1.5">
                        <button
                          onClick={() => openEdit(row)}
                          className="px-2.5 py-1 text-xs font-medium text-blue-600 bg-blue-50 dark:bg-blue-900/30 hover:bg-blue-100 rounded-lg transition-all"
                        >
                          Edit
                        </button>
                        <button
                          onClick={() => handleDelete(row.id)}
                          className="px-2.5 py-1 text-xs font-medium text-red-600 bg-red-50 dark:bg-red-900/30 hover:bg-red-100 rounded-lg transition-all"
                        >
                          Hapus
                        </button>
                      </div>
                    </td>
                  </tr>
                ))}
                {data.length === 0 && (
                  <tr>
                    <td colSpan={4} className="text-center text-gray-400 py-10">Belum ada data</td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        )}
      </div>

      {/* Modal */}
      {modal && (
        <div className="fixed inset-0 bg-black/40 flex items-end sm:items-center justify-center z-50 p-0 sm:p-4">
          <div className="bg-white dark:bg-gray-900 rounded-t-3xl sm:rounded-2xl shadow-xl p-6 w-full sm:max-w-md border border-gray-100 dark:border-gray-800">
            {/* Handle bar mobile */}
            <div className="w-10 h-1 bg-gray-200 dark:bg-gray-700 rounded-full mx-auto mb-5 sm:hidden" />
            <h2 className="text-lg font-bold text-gray-800 dark:text-white mb-5">
              {modal === "add" ? "Tambah Data Harga" : "Edit Data Harga"}
            </h2>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1.5">Tanggal</label>
                <input
                  type="date"
                  value={form.tanggal}
                  onChange={(e) => setForm((p) => ({ ...p, tanggal: e.target.value }))}
                  className="w-full px-4 py-2.5 rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-gray-800 dark:text-white text-sm focus:outline-none focus:ring-2 focus:ring-green-500"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1.5">Harga (Rp)</label>
                <input
                  type="number"
                  value={form.harga}
                  onChange={(e) => setForm((p) => ({ ...p, harga: e.target.value }))}
                  placeholder="contoh: 45000"
                  className="w-full px-4 py-2.5 rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-gray-800 dark:text-white text-sm focus:outline-none focus:ring-2 focus:ring-green-500"
                />
              </div>
              {error && <p className="text-sm text-red-500">{error}</p>}
            </div>
            <div className="flex gap-3 mt-6">
              <button
                onClick={() => setModal(null)}
                className="flex-1 py-2.5 text-sm font-medium text-gray-600 dark:text-gray-300 bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 rounded-xl transition-all"
              >
                Batal
              </button>
              <button
                onClick={handleSave}
                disabled={saving}
                className="flex-1 py-2.5 text-sm font-semibold text-white bg-green-600 hover:bg-green-700 rounded-xl transition-all disabled:opacity-60"
              >
                {saving ? "Menyimpan..." : "Simpan"}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
