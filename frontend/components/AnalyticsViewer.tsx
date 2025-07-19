'use client';
import { useEffect, useState } from "react";

// Dummy slug list; you can fetch from API or props!

type ClickDetail = {
    timestamp: string;
    ip: string;
    user_agent: string;
    referrer: string;
};

type AnalyticsResponse = {
    slug: string;
    click_count: number;
    analytics: ClickDetail[];
};

export default function AnalyticsViewer() {
    const [slug, setSlug] = useState<string>("");
    const [data, setData] = useState<AnalyticsResponse | null>(null);
    const [error, setError] = useState<string>("");
    const [loading, setLoading] = useState(false);

    const fetchAnalytics = async (chosenSlug = slug) => {
        setError(""); setData(null); setLoading(true);
        try {
            const res = await fetch(`http://localhost:8080/analytics/${chosenSlug}`);
            if (!res.ok) throw new Error("No one clicked your link or it does not exist");
            const json: AnalyticsResponse = await res.json();
            setData(json);
        } catch (err: any) {
            setError(err.message || "Unknown error");
        } finally {
            setLoading(false);
        }
    };




    return (
        <div className="min-h-screen bg-gradient-to-br from-blue-100 via-white to-purple-100 flex flex-col items-center justify-center px-2">
            <div className="w-full max-w-4xl rounded-3xl shadow-2xl bg-white/70 backdrop-blur-md p-8 border border-blue-100 relative animate-fadein">
                <div className="absolute right-8 top-8 text-blue-400 font-mono text-xs opacity-70">
                    <span className="inline-block animate-pulse">🔎</span> Live Analytics
                </div>
                <h2 className="text-3xl font-extrabold text-purple-700 mb-2 drop-shadow">URL Analytics Dashboard</h2>
                <p className="mb-6 text-gray-600">Track click stats for your short URLs in real time 🚀</p>

                <div className="flex flex-col md:flex-row gap-4 items-center mb-5">
                    <input
                        type="text"
                        className="border rounded-xl p-2 min-w-[180px] bg-blue-50 font-mono focus:ring-2 ring-purple-200"
                        placeholder="Enter slug (e.g. mygithub3)"
                        value={slug}
                        onChange={e => setSlug(e.target.value)}
                        onKeyDown={e => {
                            if (e.key === "Enter") fetchAnalytics();
                        }}
                    />
                    <button
                        className="bg-gradient-to-tr from-blue-500 to-purple-500 text-white px-6 py-2 rounded-xl shadow hover:scale-105 transition-transform"
                        onClick={() => fetchAnalytics()}
                        disabled={!slug || loading}
                    >
                        {loading ? "Fetching..." : "Get Analytics"}
                    </button>
                </div>


                {error && (
                    <div className="bg-red-50 border-l-4 border-red-400 p-3 text-red-700 mb-4 rounded-xl">
                        <b>Oops:</b> {error}
                    </div>
                )}

                {!data && !error && (
                    <div className="text-gray-400 text-center py-12">
                        <span className="text-5xl mb-4 block">📊</span>
                        <div>Select a slug and click 'Get Analytics' to view data.</div>
                    </div>
                )}

                {data && (
                    <div className="animate-fadein2">
                        <div className="flex flex-wrap justify-between items-center bg-blue-50 rounded-xl px-4 py-3 mb-4 shadow-inner">
                            <div className="flex items-center gap-2 font-bold text-lg text-blue-700">
                                <span className="font-mono text-base">🔗</span> {data.slug}
                            </div>
                            <div className="font-mono text-purple-600 text-md">
                                <span className="font-bold text-2xl">{data.click_count}</span> clicks
                            </div>
                        </div>
                        <div>
                            <h3 className="font-semibold text-blue-600 mt-2 mb-3">Recent Clicks</h3>
                            <div className="overflow-x-auto max-h-60 border rounded-xl bg-white shadow">
                                <table className="min-w-full text-xs md:text-sm">
                                    <thead className="bg-blue-100">
                                        <tr>
                                            <th className="py-2 px-3 text-left">Timestamp</th>
                                            <th className="py-2 px-3 text-left">IP</th>
                                            <th className="py-2 px-3 text-left">User Agent</th>
                                            <th className="py-2 px-3 text-left">Referrer</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        {data.analytics.length === 0 && (
                                            <tr>
                                                <td colSpan={4} className="py-3 text-gray-400 text-center">
                                                    <span className="text-2xl">😶</span> No analytics data yet.
                                                </td>
                                            </tr>
                                        )}
                                        {data.analytics.map((item, i) => (
                                            <tr
                                                key={i}
                                                className={i % 2 ? "bg-blue-50" : "bg-white hover:bg-blue-100 transition"}
                                            >
                                                <td className="py-2 px-3 font-mono">{new Date(item.timestamp).toLocaleString()}</td>
                                                <td className="py-2 px-3">{item.ip}</td>
                                                <td className="py-2 px-3 truncate max-w-[180px]">{item.user_agent}</td>
                                                <td className="py-2 px-3 truncate max-w-[120px]">{item.referrer || <span className="opacity-60">–</span>}</td>
                                            </tr>
                                        ))}
                                    </tbody>
                                </table>
                            </div>
                            <div className="text-xs text-gray-400 mt-3 text-right">Last updated: {new Date().toLocaleTimeString()}</div>
                        </div>
                    </div>
                )}
            </div>
            {/* Add a soft floating gradient blob for effect */}
            <div className="fixed -z-10 top-[-12vh] right-[-10vw] w-[340px] h-[340px] bg-purple-300 opacity-20 rounded-full blur-3xl pointer-events-none"></div>
        </div>
    );
}


