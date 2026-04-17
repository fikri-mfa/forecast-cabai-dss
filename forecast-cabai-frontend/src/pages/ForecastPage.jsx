import { useState } from "react";
import { useAuth } from "../context/AuthContext";
import { runForecast } from "../services/forecastService";
import { useNavigate } from "react-router-dom";

const DEFAULT_FORM = {
  alpha: 0.2,
  beta: 0.1,
  gamma: 0.1,
  periods: 3,
  season_length: 12,
};

export default function ForecastPage() {
  const { token, logout } = useAuth();
  const navigate = useNavigate();
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
    setError("");
    setResult(null);
    setLoading(true);
    try {
      const data = await runForecast(form, token);
      setResult(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    logout();
    navigate("/auth");
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-green-50 via-white to-emerald-50">
      {/* Navbar */}
      <nav className="bg-white border-b border-gray-100 shadow-sm px-6 py-4 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <div className="w-8 h-8 bg-green-600 rounded-lg flex items-center justify-center">
            <svg className="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <span className="font-bold text-gray-800 text-sm">DSS Forecast Cabai</span>
        </div>
        <button
          onClick={handleLogout}
          className="text-sm text-gray-500 hover:text-red-500 font-medium transition-colors"
        >
          Logout
        </button>
      </nav>

      <div className="max-w-2xl mx-auto px-4 py-10">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-2xl font-bold text-gray-800">Prediksi Harga Cabai</h1>
          <p className="text-sm text-gray-500 mt-1">
            Gunakan Triple Exponential Smoothing untuk memprediksi harga ke depan
          </p>
        </div>

        {/* Form Card */}
        <div className="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 mb-6">
          <h2 className="text-sm font-semibold text-gray-700 mb-5 uppercase tracking-wide">
            Parameter Forecast
          </h2>
          <form onSubmit={handleSubmit} className="space-y-5">
            {/* Alpha, Beta, Gamma */}
            <div className="grid grid-cols-3 gap-4">
              {["alpha", "beta", "gamma"].map((key) => (
                <div key={key}>
                  <label className="block text-sm font-medium text-gray-700 mb-1.5 capitalize">
                    {key}
                    <span className="text-xs text-gray-400 ml-1">(0–1)</span>
                  </label>
                  <input
                    type="number"
                    name={key}
                    value={form[key]}
                    onChange={handleChange}
                    min={0}
                    max={1}
                    step={0.01}
                    required
                    className="w-full px-3 py-2.5 rounded-xl border border-gray-200 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition-all"
                  />
                </div>
              ))}
            </div>

            {/* Periods & Season Length */}
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1.5">
                  Periods
                  <span className="text-xs text-gray-400 ml-1">(jumlah prediksi)</span>
                </label>
                <input
                  type="number"
                  name="periods"
                  value={form.periods}
                  onChange={handleChange}
                  min={1}
                  max={24}
                  required
                  className="w-full px-3 py-2.5 rounded-xl border border-gray-200 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition-all"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1.5">
                  Season Length
                  <span className="text-xs text-gray-400 ml-1">(default 12)</span>
                </label>
                <input
                  type="number"
                  name="season_length"
                  value={form.season_length}
                  onChange={handleChange}
                  min={2}
                  required
                  className="w-full px-3 py-2.5 rounded-xl border border-gray-200 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition-all"
                />
              </div>
            </div>

            {/* Error */}
            {error && (
              <div className="px-4 py-3 bg-red-50 border border-red-200 rounded-xl flex items-center gap-2">
                <svg className="w-4 h-4 text-red-500 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
                </svg>
                <p className="text-sm text-red-600">{error}</p>
              </div>
            )}

            <button
              type="submit"
              disabled={loading}
              className="w-full py-3 bg-green-600 hover:bg-green-700 text-white text-sm font-semibold rounded-xl transition-all disabled:opacity-60 disabled:cursor-not-allowed shadow-sm"
            >
              {loading ? "Memproses..." : "Jalankan Forecast"}
            </button>
          </form>
        </div>

        {/* Result Card */}
        {result && (
          <div className="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 space-y-6">
            <h2 className="text-sm font-semibold text-gray-700 uppercase tracking-wide">
              Hasil Forecast
            </h2>

            {/* Evaluation */}
            <div className="grid grid-cols-2 gap-4">
              <div className="bg-green-50 rounded-2xl p-4 text-center">
                <p className="text-xs text-gray-500 mb-1">MAPE</p>
                <p className="text-2xl font-bold text-green-700">
                  {result.evaluation.mape.toFixed(2)}%
                </p>
              </div>
              <div className="bg-blue-50 rounded-2xl p-4 text-center">
                <p className="text-xs text-gray-500 mb-1">RMSE</p>
                <p className="text-2xl font-bold text-blue-700">
                  {result.evaluation.rmse.toFixed(0)}
                </p>
              </div>
            </div>

            {/* Forecast values */}
            <div>
              <p className="text-sm font-medium text-gray-700 mb-3">Prediksi Harga (Rp)</p>
              <div className="space-y-2">
                {result.forecast.map((val, i) => (
                  <div
                    key={i}
                    className="flex items-center justify-between px-4 py-3 bg-gray-50 rounded-xl"
                  >
                    <span className="text-sm text-gray-500">Periode {i + 1}</span>
                    <span className="text-sm font-semibold text-gray-800">
                      Rp {val.toLocaleString("id-ID", { maximumFractionDigits: 0 })}
                    </span>
                  </div>
                ))}
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}