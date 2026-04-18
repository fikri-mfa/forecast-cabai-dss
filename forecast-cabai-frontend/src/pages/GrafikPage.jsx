import { useEffect, useState } from "react";
import { useAuth } from "../context/AuthContext";
import { getAllChartData } from "../services/forecastService";
import {
  LineChart, Line, XAxis, YAxis, CartesianGrid,
  Tooltip, Legend, ResponsiveContainer
} from "recharts";

export default function GrafikPage() {
  const { token } = useAuth();
  const [chart, setChart] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function load() {
      try {
        const res = await getAllChartData(token);
        setChart(res || []);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, []);

  return (
    <div className="p-4 sm:p-6 lg:p-8">
      <div className="mb-6 sm:mb-8">
        <h1 className="text-xl sm:text-2xl font-bold text-gray-800 dark:text-white">Grafik Detail</h1>
        <p className="text-sm text-gray-400 mt-1">Perbandingan harga aktual vs prediksi TES</p>
      </div>

      {/* Stat cards — 1 kolom di mobile, 3 di desktop */}
      {chart.length > 0 && (
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-3 sm:gap-4 mb-5 sm:mb-6">
          <div className="bg-white dark:bg-gray-900 rounded-2xl p-4 sm:p-5 shadow-sm border border-gray-100 dark:border-gray-800">
            <p className="text-xs text-gray-400 uppercase tracking-wide mb-1">Data Ditampilkan</p>
            <p className="text-xl sm:text-2xl font-bold text-gray-800 dark:text-white">{chart.length} periode</p>
          </div>
          <div className="bg-white dark:bg-gray-900 rounded-2xl p-4 sm:p-5 shadow-sm border border-gray-100 dark:border-gray-800">
            <p className="text-xs text-gray-400 uppercase tracking-wide mb-1">Harga Tertinggi</p>
            <p className="text-xl sm:text-2xl font-bold text-red-600">
              Rp {Math.max(...chart.map((d) => d.aktual)).toLocaleString("id-ID", { maximumFractionDigits: 0 })}
            </p>
          </div>
          <div className="bg-white dark:bg-gray-900 rounded-2xl p-4 sm:p-5 shadow-sm border border-gray-100 dark:border-gray-800">
            <p className="text-xs text-gray-400 uppercase tracking-wide mb-1">Harga Terendah</p>
            <p className="text-xl sm:text-2xl font-bold text-green-600">
              Rp {Math.min(...chart.map((d) => d.aktual)).toLocaleString("id-ID", { maximumFractionDigits: 0 })}
            </p>
          </div>
        </div>
      )}

      {/* Chart */}
      <div className="bg-white dark:bg-gray-900 rounded-2xl p-4 sm:p-6 shadow-sm border border-gray-100 dark:border-gray-800 mb-5 sm:mb-6">
        <h2 className="text-sm font-semibold text-gray-700 dark:text-gray-300 mb-4 sm:mb-6">
          Harga Aktual vs Prediksi TES
        </h2>
        {loading ? (
          <p className="text-sm text-gray-400 text-center py-10">Memuat grafik...</p>
        ) : chart.length === 0 ? (
          <p className="text-sm text-gray-400 text-center py-10">Belum ada data</p>
        ) : (
          <ResponsiveContainer width="100%" height={280}>
            <LineChart data={chart}>
              <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
              <XAxis
                dataKey="tanggal"
                tick={{ fontSize: 10, fill: "#9ca3af" }}
                tickFormatter={(val) => val?.substring(0, 7)}
              />
              <YAxis
                tick={{ fontSize: 10, fill: "#9ca3af" }}
                tickFormatter={(val) => `${(val / 1000).toFixed(0)}k`}
                width={36}
              />
              <Tooltip
                contentStyle={{ backgroundColor: "#1f2937", border: "none", borderRadius: "12px", color: "#f9fafb", fontSize: "12px" }}
                formatter={(val) => `Rp ${val.toLocaleString("id-ID", { maximumFractionDigits: 0 })}`}
                labelFormatter={(label) => label?.substring(0, 7)}
              />
              <Legend wrapperStyle={{ fontSize: "12px" }} />
              <Line type="monotone" dataKey="aktual" name="Harga Aktual" stroke="#16a34a" strokeWidth={2} dot={{ r: 3 }} activeDot={{ r: 5 }} />
              <Line type="monotone" dataKey="prediksi" name="Prediksi TES" stroke="#3b82f6" strokeWidth={2} strokeDasharray="5 5" dot={{ r: 3 }} activeDot={{ r: 5 }} />
            </LineChart>
          </ResponsiveContainer>
        )}
      </div>

      {/* Tabel detail dengan scroll horizontal */}
      <div className="bg-white dark:bg-gray-900 rounded-2xl shadow-sm border border-gray-100 dark:border-gray-800 overflow-hidden">
        <div className="px-4 sm:px-6 py-4 border-b border-gray-100 dark:border-gray-800">
          <h2 className="text-sm font-semibold text-gray-700 dark:text-gray-300">Detail Per Periode</h2>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full text-sm min-w-[480px]">
            <thead>
              <tr className="bg-gray-50 dark:bg-gray-800 border-b border-gray-100 dark:border-gray-700">
                <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Periode</th>
                <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Harga Aktual</th>
                <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Prediksi TES</th>
                <th className="text-left px-4 sm:px-6 py-3 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Selisih</th>
              </tr>
            </thead>
            <tbody>
              {chart.map((row, i) => {
                const selisih = row.aktual - row.prediksi;
                return (
                  <tr key={i} className="border-b border-gray-50 dark:border-gray-800 hover:bg-gray-50 dark:hover:bg-gray-800">
                    <td className="px-4 sm:px-6 py-3 text-gray-600 dark:text-gray-300 whitespace-nowrap">
                      {row.tanggal?.substring(0, 7)}
                    </td>
                    <td className="px-4 sm:px-6 py-3 font-medium text-gray-800 dark:text-white whitespace-nowrap">
                      Rp {row.aktual.toLocaleString("id-ID", { maximumFractionDigits: 0 })}
                    </td>
                    <td className="px-4 sm:px-6 py-3 text-blue-600 whitespace-nowrap">
                      Rp {row.prediksi.toLocaleString("id-ID", { maximumFractionDigits: 0 })}
                    </td>
                    <td className={`px-4 sm:px-6 py-3 font-medium whitespace-nowrap ${selisih >= 0 ? "text-green-600" : "text-red-500"}`}>
                      {selisih >= 0 ? "+" : ""}Rp {selisih.toLocaleString("id-ID", { maximumFractionDigits: 0 })}
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
