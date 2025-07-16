'use client';

import React, { useState } from 'react';
import { DotPattern } from './magicui/dot-pattern';
import {
  FaGithub,
  FaLinkedin,
  FaTwitter,
} from 'react-icons/fa';

const GITHUB_REPO = 'https://github.com/Shobhit150/url_shortner';
const LINKEDIN = 'https://www.linkedin.com/in/shobhit-verma-13947b235/';
const X_URL = 'https://x.com/vshobhit150';

const Shorten = () => {
  const [inputUrl, setInputUrl] = useState('');
  const [customUrl, setCustomUrl] = useState('');
  const [shortResult, setShortResult] = useState('');
  const [loading, setLoading] = useState(false);


  const handleShorten = async () => {
    if (!inputUrl) return;
    setLoading(true);
    try {
      const resp = await fetch(`http://localhost:8080/shorten`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ url: inputUrl, custom_slug: customUrl }),
      });
      const data = await resp.json();
      setShortResult(data.shorten_url);
    } catch (err) {
      console.error('Failed to shorten URL', err);
    } finally {
      setLoading(false);
    }
  };

  // One-time "Experience" animation


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
            <a
              href={GITHUB_REPO}
              target="_blank"
              rel="noopener noreferrer"
              className="text-lg text-gray-400 hover:text-white transition"
              title="GitHub Repo"
            >
              <FaGithub />
            </a>
            <a
              href={LINKEDIN}
              target="_blank"
              rel="noopener noreferrer"
              className="text-lg text-blue-400 hover:text-white transition"
              title="LinkedIn"
            >
              <FaLinkedin />
            </a>
            <a
              href={X_URL}
              target="_blank"
              rel="noopener noreferrer"
              className="text-lg text-gray-400 hover:text-white transition"
              title="X (Twitter)"
            >
              <FaTwitter />
              {/* If you want the "X" SVG instead, replace <FaTwitter /> with the SVG I shared earlier */}
            </a>
            <a
              href={GITHUB_REPO}
              target="_blank"
              rel="noopener noreferrer"
              className="text-sm hidden md:inline hover:underline underline-offset-4 text-blue-400 ml-2"
            >
              View on GitHub
            </a>
          </div>
        </nav>

        {/* Main Content */}
        <main className="flex flex-col items-center justify-center flex-grow p-6">
          <div
            className={`w-full max-w-xl bg-white/10 backdrop-blur-md p-8 rounded-3xl shadow-2xl border border-white/20 transition-all duration-300`}
          >
            <div className="flex flex-col items-center mb-6">
              <h1 className="text-4xl font-extrabold text-center mb-2 bg-clip-text text-transparent bg-gradient-to-r from-pink-300 via-purple-200 to-sky-300 select-none">
                Shorten Your URL
              </h1>
              
            </div>

            <div className="flex flex-col gap-4">
              <input
                type="text"
                value={inputUrl}
                onChange={(e) => setInputUrl(e.target.value)}
                placeholder="Paste your long URL here..."
                className="bg-white/20 border border-white/30 rounded-lg px-4 py-3 text-base text-white outline-none focus:ring-2 focus:ring-blue-400/60 transition"
              />
              <input
                type="text"
                value={customUrl}
                onChange={(e) => setCustomUrl(e.target.value)}
                placeholder="Optional: Custom slug (e.g., my-url)"
                className="bg-white/20 border border-white/30 rounded-lg px-4 py-3 text-base text-white outline-none focus:ring-2 focus:ring-pink-400/60 transition"
              />
              <button
                onClick={handleShorten}
                disabled={loading}
                className="bg-gradient-to-r from-blue-500 to-purple-500 hover:from-pink-500 hover:to-blue-400 disabled:opacity-60 px-6 py-3 rounded-lg font-semibold shadow-lg transition-all duration-200"
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

            {shortResult && (
              <div className="mt-8 text-center space-y-2">
                <p className="text-base text-blue-200 font-medium">Your shortened URL:</p>
                <a
                  href={shortResult}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-xl text-blue-400 hover:underline break-all font-mono"
                >
                  {shortResult}
                </a>
              </div>
            )}
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

      {/* Animation CSS */}
      
    </div>
  );
};

export default Shorten;
