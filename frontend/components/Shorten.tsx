'use client';

import React, { useState, useRef } from 'react';
import { DotPattern } from './magicui/dot-pattern';
import { FaGithub, FaLinkedin, FaTwitter, FaChartBar } from 'react-icons/fa';
import { QRCodeCanvas } from "qrcode.react";
import { useRouter } from 'next/navigation';
import Image from 'next/image';

const GITHUB_REPO = 'https://github.com/Shobhit150/url_shortner';
const LINKEDIN = 'https://www.linkedin.com/in/shobhit150/';
const X_URL = 'https://x.com/Shobhit_codes';

const Shorten = () => {
    const [inputUrl, setInputUrl] = useState('');
    const [customUrl, setCustomUrl] = useState('');
    const [shortResult, setShortResult] = useState('');
    const [loading, setLoading] = useState(false);
    const [expiry, setExpiry] = useState('');
    const [error, setError] = useState('');
    const [showCopied, setShowCopied] = useState(false);

    const qrRef = useRef<HTMLDivElement>(null);
    const router = useRouter();

    // Download QR as PNG
    const downloadQrCode = () => {
        if (!qrRef.current) return;
        const canvas = qrRef.current.querySelector('canvas');
        if (!canvas) return;
        const url = canvas.toDataURL("image/png");
        const link = document.createElement('a');
        link.href = url;
        link.download = "qrcode.png";
        link.click();
    };
    const handleShorten = async () => {
        if (!inputUrl) return;
        setLoading(true); setError('');
        try {
            const payload = {
                url: inputUrl,
                custom_slug: customUrl,
                expires_at: expiry ? new Date(expiry).toISOString() : null,
            };
            const resp = await fetch(`http://localhost:8080/shorten`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload),
            });

            const data = await resp.json();

            // Handle error from backend
            if (!resp.ok) {
                setShortResult('');
                setError(data.error || "Failed to shorten URL2");
                return;
            }

            setShortResult(data.shorten_url);
        } catch (err) {
            setShortResult('');
            setError('Failed to shorten URL12');
        } finally {
            setLoading(false);
        }
    };


    const handleCopy = () => {
        if (!shortResult) return;
        navigator.clipboard.writeText(shortResult);
        setShowCopied(true);
        setTimeout(() => setShowCopied(false), 1200);
    };

    // Parse slug from result for analytics link
    const getSlugFromResult = () => {
        try {
            const url = new URL(shortResult);
            return url.pathname.replace(/^\//, '');
        } catch {
            return '';
        }
    };

    return (
        <div className="relative min-h-screen overflow-hidden text-white font-sans">
            {/* Background gradient */}
            <div className="absolute inset-0 bg-gradient-to-br from-[#1a1e22] via-[#13161a] to-[#141920] z-0" />
            {/* Dot animation layer */}
            <div className="absolute inset-0 z-10 pointer-events-none">
                <DotPattern className="text-white/25" glow />
            </div>

            {/* Foreground content */}
            <div className="relative z-20 flex flex-col min-h-screen">
                {/* Navbar */}
                <nav className="flex items-center justify-between px-7 py-4 bg-black/70 shadow-lg border-b border-white/10">
                    <div className="flex items-center gap-2 text-xl font-bold tracking-tight">
                        <span className="animate-pulse">🔗</span>
                        <span className="tracking-wider bg-clip-text text-transparent bg-gradient-to-r from-sky-300 via-purple-300 to-pink-400">
                            URL Shortener
                        </span>
                    </div>
                    <div className="flex items-center gap-4">
                        <a href={GITHUB_REPO} target="_blank" rel="noopener noreferrer"
                            className="text-lg text-gray-400 hover:text-white transition" title="GitHub Repo">
                            <FaGithub />
                        </a>
                        <a href={LINKEDIN} target="_blank" rel="noopener noreferrer"
                            className="text-lg text-blue-400 hover:text-white transition" title="LinkedIn">
                            <FaLinkedin />
                        </a>
                        <a href={X_URL} target="_blank" rel="noopener noreferrer"
                            className="text-lg text-gray-400 hover:text-white transition" title="X (Twitter)">
                            <FaTwitter />
                        </a>
                        <a href={GITHUB_REPO} target="_blank" rel="noopener noreferrer"
                            className="text-sm hidden md:inline hover:underline underline-offset-4 text-blue-400 ml-2">
                            View on GitHub
                        </a>
                    </div>
                </nav>

                {/* Main Content */}
                <main className="flex flex-row items-center justify-around flex-grow p-6">
                    <div className="w-full max-w-xl bg-white/10 p-8 rounded-3xl  border border-white/20">
                        <div className="flex flex-col items-center mb-6">
                            <h1 className="text-4xl font-extrabold text-center mb-2 bg-clip-text text-transparent bg-gradient-to-r from-pink-300 via-purple-200 to-sky-300 select-none drop-shadow">
                                Shorten Your URL
                            </h1>
                        </div>
                        <div className="flex flex-col gap-4">
                            <input
                                type="text"
                                value={inputUrl}
                                onChange={(e) => setInputUrl(e.target.value)}
                                placeholder="Paste your long URL here..."
                                className="bg-white/20 border border-white/30 rounded-lg px-4 py-3 text-base text-white outline-none focus:ring-2 focus:ring-blue-400/60 transition placeholder:italic"
                            />
                            <input
                                type="text"
                                value={customUrl}
                                onChange={(e) => setCustomUrl(e.target.value)}
                                placeholder="Optional: Custom slug (e.g., my-url)"
                                className="bg-white/20 border border-white/30 rounded-lg px-4 py-3 text-base text-white outline-none focus:ring-2 focus:ring-pink-400/60 transition placeholder:italic"
                            />
                            <input
                                type="datetime-local"
                                value={expiry}
                                onChange={(e) => setExpiry(e.target.value)}
                                onClick={(e) => (e.target as HTMLInputElement).showPicker()}
                                className="bg-white/20 border border-white/30 rounded-lg px-4 py-3 text-base text-white outline-none focus:ring-2 focus:ring-sky-400/60 transition placeholder:italic"
                            />
                            <button
                                onClick={handleShorten}
                                disabled={loading}
                                className="bg-gradient-to-r from-blue-500 to-purple-500  disabled:opacity-60 px-6 py-3 rounded-lg font-semibold shadow-lg cursor-pointer focus:outline-none focus:ring-4 focus:ring-purple-300"
                            >
                                {loading ? (
                                    <span className="flex items-center gap-2">
                                        <span className="animate-spin inline-block w-4 h-4 border-2 border-white border-t-transparent rounded-full" />
                                        Shortening...
                                    </span>
                                ) : (
                                    'Shorten URL'
                                )}
                            </button>
                        </div>

                        {error && (
                            <div className="mt-6 text-center bg-red-600/30 text-red-200 py-2 rounded-xl border border-red-400/40 shadow">
                                {error}
                            </div>
                        )}

                        {shortResult && (
                            <div className="mt-8 text-center space-y-2 animate-fadein2">
                                <p className="text-base text-blue-200 font-medium">Your shortened URL:</p>
                                <div className="flex flex-col md:flex-row md:items-center gap-2 justify-center">
                                    <a
                                        href={shortResult}
                                        target="_blank"
                                        rel="noopener"
                                        className="text-xl text-blue-400 hover:underline break-all font-mono transition"
                                    >
                                        {shortResult}
                                    </a>
                                    <button
                                        onClick={handleCopy}
                                        className="ml-1 px-2 py-1 text-xs bg-blue-400/80 hover:bg-blue-500 text-white rounded transition"
                                    >
                                        {showCopied ? "Copied!" : "Copy"}
                                    </button>
                                </div>
                                <div className="flex flex-col items-center mt-6 space-y-2">
                                    <div ref={qrRef}>
                                        {/* @ts-ignore */}
                                        <QRCodeCanvas value={shortResult} size={180} level="H" includeMargin={true} />
                                    </div>
                                    <button
                                        onClick={downloadQrCode}
                                        className="mt-2 bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded shadow transition"
                                    >
                                        Download QR Code
                                    </button>
                                    <span className="text-xs text-gray-400">Scan or download QR code</span>
                                </div>

                                {/* --- Analytics Button --- */}
                                <div className="flex justify-center mt-6">
                                    <button
                                        className="flex items-center gap-2 px-5 py-2 rounded-xl bg-gradient-to-r from-purple-600 to-pink-500 hover:scale-105 shadow-md transition-all text-white font-semibold text-base"
                                        onClick={() => router.push(`/analysis?slug=${getSlugFromResult()}`)}
                                    >
                                        <FaChartBar className="text-lg" />
                                        View Analytics
                                    </button>
                                </div>
                            </div>
                        )}
                    </div>
                    <div className="md:w-1/2 flex flex-col items-center text-center gap-3 p-4 bg-white/5 backdrop-blur-lg rounded-2xl border border-white/10 shadow-xl">
                        <h3 className="text-lg font-semibold text-blue-200 tracking-wide">
                            Background Architecture
                        </h3>
                        <Image
                            src="/architecture.webp"
                            alt="System Architecture"
                            width={650}
                            height={650}
                            className="rounded-xl shadow-2xl border border-white/20 hover:scale-[1.02] transition-transform duration-300 ease-out"
                        />
                    </div>

                </main>

                {/* Footer */}
                <footer className="text-xs text-gray-400 text-center py-4 mt-2 select-none">
                    <span>
                        Made with <span className="text-pink-400">♥</span> by Shobhit · Powered by Go + Next.js
                    </span>
                    <br />
                    <span className="text-[10px]">Icons by react-icons · Dot animation by MagicUI</span>
                </footer>
            </div>

        </div>
    );
};

export default Shorten;
