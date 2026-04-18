import { useEffect, useState } from "react";
import { useAuth } from "../context/AuthContext";
import { getDashboardStats, getDashboardChart } from "../services/forecastService";
import {
  LineChart, Line, XAxis, YAxis, CartesianGrid,
  Tooltip, Legend, ResponsiveContainer
} from "recharts";

function StatCard({ title, value, sub, color }) {
  return (
    <div className="bg-white dark:bg-gray-900 rounded-2xl p-4 sm:p-5 shadow-sm border border-gray-100 dark:border-gray-800">
      <p className="text-xs text-gray-400 uppercase tracking-wide mb-1 leading-tight">{title}</p>
      <p className={`text-xl sm:text-2xl font-bold break-words ${color || "text-gray-800 dark:text-white"}`}>{value}</p>
      {sub && <p className="text-xs text-gray-400 mt-1">{sub}</p>}
    </div>
  );
}

export default function DashboardPage() {
  const { token } = useAuth();
  const [stats, setStats] = useState(null);
  const [chart, setChart] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function load() {
      try {
        const [s, c] = await Promise.all([
          getDashboardStats(token),
          getDashboardChart(token),
        ]);
        setStats(s);
        setChart(c || []);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, [token]);

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <p className="text-gray-400 text-sm">Memuat data...</p>
      </div>
    );
  }

  return (
    <div className="p-4 sm:p-6 lg:p-8">
      <div className="mb-6 sm:mb-8">
        <h1 className="text-xl sm:text-2xl font-bold text-gray-800 dark:text-white">Dashboard</h1>
        <p className="text-sm text-gray-400 mt-1">Ringkasan sistem forecast harga cabai</p>
      </div>

      {/* Grid: 2 kolom di mobile, 4 kolom di desktop */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-3 sm:gap-4 mb-6 sm:mb-8">
        <StatCard
          title="Total Data Harga"
          value={stats?.total_data_harga ?? "-"}
          sub="data historis"
        />
        <StatCard
          title="Total Forecast"
          value={stats?.total_forecasts ?? "-"}
          sub="perhitungan dilakukan"
          color="text-blue-600"
        />
        <StatCard
          title="Akurasi Sistem"
          value={stats ? `${stats.akurasi_sistem.toFixed(1)}%` : "-"}
          sub="rata-rata akurasi"
          color="text-green-600"
        />
        <StatCard
          title="Rekomendasi"
          value={
            stats
              ? `Rp ${stats.rekomendasi_hari_ini.toLocaleString("id-ID", { maximumFractionDigits: 0 })}`
              : "-"
          }
          sub="prediksi 1 periode ke depan"
          color="text-emerald-600"
        />
      </div>

      <div className="bg-white dark:bg-gray-900 rounded-2xl p-4 sm:p-6 shadow-sm border border-gray-100 dark:border-gray-800">
        <h2 className="text-sm font-semibold text-gray-700 dark:text-gray-300 mb-4 sm:mb-6">
          Harga vs Prediksi (7 Data Terakhir)
        </h2>
        {chart.length === 0 ? (
          <p className="text-sm text-gray-400 text-center py-10">Belum ada data chart</p>
        ) : (
          <ResponsiveContainer width="100%" height={260}>
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
              <Line type="monotone" dataKey="aktual" name="Harga Aktual" stroke="#16a34a" strokeWidth={2} dot={{ r: 3 }} />
              <Line type="monotone" dataKey="prediksi" name="Prediksi TES" stroke="#3b82f6" strokeWidth={2} strokeDasharray="5 5" dot={{ r: 3 }} />
            </LineChart>
          </ResponsiveContainer>
        )}
      </div>
    </div>
  );
}
