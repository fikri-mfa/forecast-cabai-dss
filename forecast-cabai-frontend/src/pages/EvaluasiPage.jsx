import { useEffect, useState } from "react";
import { useAuth } from "../context/AuthContext";
import { getForecastHistory } from "../services/forecastService";

export default function EvaluasiPage() {
  const { token } = useAuth();
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function load() {
      try {
        const res = await getForecastHistory(token);
        setData(res.data || []);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, []);

  const avgMAPE =
    data.length > 0 ? data.reduce((sum, d) => sum + (d.mape || 0), 0) / data.length : null;

  return (
    <div className="p-4 sm:p-6 lg:p-8">
      <div className="mb-6 sm:mb-8">
        <h1 className="text-xl sm:text-2xl font-bold text-gray-800 dark:text-white">Evaluasi Akurasi</h1>
        <p className="text-sm text-gray-400 mt-1">Riwayat hasil forecast dan tingkat akurasi</p>
      </div>

      {/* Stat cards — 2 kolom di mobile */}
      {avgMAPE !== null && (
        <div className="grid grid-cols-2 gap-3 sm:gap-4 mb-5 sm:mb-6">
          <div className="bg-green-50 dark:bg-green-900/20 rounded-2xl p-4 sm:p-5 border border-green-100 dark:border-green-900">
            <p className="text-xs text-gray-500 dark:text-gray-400 uppercase tracking-wide mb-1">Rata-rata Akurasi</p>
            <p className="text-2xl sm:text-3xl font-bold text-green-700 dark:text-green-400">
              {(100 - avgMAPE).toFixed(1)}%
            </p>
          </div>
          <div className="bg-blue-50 dark:bg-blue-900/20 rounded-2xl p-4 sm:p-5 border border-blue-100 dark:border-blue-900">
            <p className="text-xs text-gray-500 dark:text-gray-400 uppercase tracking-wide mb-1">Total Forecast</p>
            <p className="text-2xl sm:text-3xl font-bold text-blue-700 dark:text-blue-400">{data.length}</p>
          </div>
        </div>
      )}

      {/* Tabel dengan scroll horizontal */}
      <div className="bg-white dark:bg-gray-900 rounded-2xl shadow-sm border border-gray-100 dark:border-gray-800 overflow-hidden">
        {loading ? (
          <p className="text-sm text-gray-400 text-center py-10">Memuat data...</p>
        ) : (
          <div className="overflow-x-auto">
            <table className="w-full text-sm min-w-[560px]">
              <thead>
                <tr className="bg-gray-50 dark:bg-gray-800 border-b border-gray-100 dark:border-gray-700">
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">No</th>
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Tanggal</th>
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">α / β / γ</th>
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Periods</th>
                  <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Hasil Forecast</th>
                </tr>
              </thead>
              <tbody>
                {data.map((row, i) => (
                  <tr
                    key={row.id}
                    className="border-b border-gray-50 dark:border-gray-800 hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors"
                  >
                    <td className="px-4 sm:px-6 py-3 text-gray-400">{i + 1}</td>
                    <td className="px-4 sm:px-6 py-3 text-gray-600 dark:text-gray-300 whitespace-nowrap">
                      {row.created_at?.substring(0, 10)}
                    </td>
                    <td className="px-4 sm:px-6 py-3 text-gray-700 dark:text-gray-300 font-mono text-xs whitespace-nowrap">
                      {row.alpha} / {row.beta} / {row.gamma}
                    </td>
                    <td className="px-4 sm:px-6 py-3 text-gray-700 dark:text-gray-300">{row.periods}</td>
                    <td className="px-4 sm:px-6 py-3">
                      <div className="flex flex-wrap gap-1">
                        {row.result?.map((v, j) => (
                          <span
                            key={j}
                            className="inline-block bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 rounded px-2 py-0.5 text-xs whitespace-nowrap"
                          >
                            Rp {v.toLocaleString("id-ID", { maximumFractionDigits: 0 })}
                          </span>
                        ))}
                      </div>
                    </td>
                  </tr>
                ))}
                {data.length === 0 && (
                  <tr>
                    <td colSpan={5} className="text-center text-gray-400 py-10">Belum ada data forecast</td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}
