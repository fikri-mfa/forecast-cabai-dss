import { useState } from "react";
import { useAuth } from "../context/AuthContext";
import { runForecast } from "../services/forecastService";

const DEFAULT_FORM = { alpha: 0.2, beta: 0.1, gamma: 0.1, periods: 3, season_length: 12 };

export default function ForecastPage() {
  const { token } = useAuth();
  const [form, setForm] = useState(DEFAULT_FORM);
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: parseFloat(value) || 0 }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(""); setResult(null); setLoading(true);
    try {
      const data = await runForecast(form, token);
      setResult(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-4 sm:p-6 lg:p-8">
      <div className="mb-6 sm:mb-8">
        <h1 className="text-xl sm:text-2xl font-bold text-gray-800 dark:text-white">Prediksi Harga Cabai</h1>
        <p className="text-sm text-gray-500 dark:text-gray-400 mt-1">
          Gunakan Triple Exponential Smoothing untuk memprediksi harga ke depan
        </p>
      </div>

      <div className="bg-white dark:bg-gray-900 rounded-2xl sm:rounded-3xl shadow-sm border border-gray-100 dark:border-gray-800 p-4 sm:p-6 mb-5">
        <h2 className="text-xs sm:text-sm font-semibold text-gray-700 dark:text-gray-300 mb-4 sm:mb-5 uppercase tracking-wide">
          Parameter Forecast
        </h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Alpha, Beta, Gamma — 1 kolom di mobile, 3 di desktop */}
          <div className="grid grid-cols-1 sm:grid-cols-3 gap-3 sm:gap-4">
            {["alpha", "beta", "gamma"].map((key) => (
              <div key={key}>
                <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1.5 capitalize">
                  {key} <span className="text-xs text-gray-400 normal-case">(0–1)</span>
                </label>
                <input
                  type="number"
                  name={key}
                  value={form[key]}
                  onChange={handleChange}
                  min={0} max={1} step={0.01} required
                  className="w-full px-3 py-2.5 rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-gray-800 dark:text-white text-sm focus:outline-none focus:ring-2 focus:ring-green-500 transition-all"
                />
              </div>
            ))}
          </div>

          {/* Periods + Season Length — 2 kolom di semua ukuran, lebih mudah di HP */}
          <div className="grid grid-cols-2 gap-3 sm:gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1.5">
                Periods <span className="text-xs text-gray-400">(prediksi)</span>
              </label>
              <input
                type="number"
                name="periods"
                value={form.periods}
                onChange={handleChange}
                min={1} max={24} required
                className="w-full px-3 py-2.5 rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-gray-800 dark:text-white text-sm focus:outline-none focus:ring-2 focus:ring-green-500 transition-all"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1.5">
                Season Length <span className="text-xs text-gray-400">(default 12)</span>
              </label>
              <input
                type="number"
                name="season_length"
                value={form.season_length}
                onChange={handleChange}
                min={2} required
                className="w-full px-3 py-2.5 rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 text-gray-800 dark:text-white text-sm focus:outline-none focus:ring-2 focus:ring-green-500 transition-all"
              />
            </div>
          </div>

          {error && (
            <div className="px-4 py-3 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl">
              <p className="text-sm text-red-600 dark:text-red-400">{error}</p>
            </div>
          )}

          <button
            type="submit"
            disabled={loading}
            className="w-full py-3 bg-green-600 hover:bg-green-700 text-white text-sm font-semibold rounded-xl transition-all disabled:opacity-60"
          >
            {loading ? "Memproses..." : "Jalankan Forecast"}
          </button>
        </form>
      </div>

      {result && (
        <div className="bg-white dark:bg-gray-900 rounded-2xl sm:rounded-3xl shadow-sm border border-gray-100 dark:border-gray-800 p-4 sm:p-6 space-y-5">
          <h2 className="text-xs sm:text-sm font-semibold text-gray-700 dark:text-gray-300 uppercase tracking-wide">
            Hasil Forecast
          </h2>

          {/* Metrik MAPE + RMSE */}
          <div className="grid grid-cols-2 gap-3 sm:gap-4">
            <div className="bg-green-50 dark:bg-green-900/20 rounded-2xl p-4 text-center">
              <p className="text-xs text-gray-500 mb-1">MAPE</p>
              <p className="text-xl sm:text-2xl font-bold text-green-700 dark:text-green-400">
                {result.evaluation.mape.toFixed(2)}%
              </p>
            </div>
            <div className="bg-blue-50 dark:bg-blue-900/20 rounded-2xl p-4 text-center">
              <p className="text-xs text-gray-500 mb-1">RMSE</p>
              <p className="text-xl sm:text-2xl font-bold text-blue-700 dark:text-blue-400">
                {result.evaluation.rmse.toFixed(0)}
              </p>
            </div>
          </div>

          {/* List prediksi per periode */}
          <div>
            <p className="text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">Prediksi Harga (Rp)</p>
            <div className="space-y-2">
              {result.forecast.map((val, i) => (
                <div
                  key={i}
                  className="flex items-center justify-between px-4 py-3 bg-gray-50 dark:bg-gray-800 rounded-xl"
                >
                  <span className="text-sm text-gray-500 dark:text-gray-400">Periode {i + 1}</span>
                  <span className="text-sm font-semibold text-gray-800 dark:text-white">
                    Rp {val.toLocaleString("id-ID", { maximumFractionDigits: 0 })}
                  </span>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
