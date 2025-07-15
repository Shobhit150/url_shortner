'use client';
import React, { useState } from 'react'

const Shorten = () => {
    const [inputUrl, setinputUrl] = useState<string>("")
    const [customUrl, setcustomUrl] = useState<string>("")
    // const [slug, setslug] = useState<string>("")
    const [shortResult, setShortResult] = useState<string>("");

    const handleShorten = async () => {
        const resp = await fetch(`http://localhost:8080/shorten`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({url: inputUrl, custom_slug: customUrl})
        });
        const data = await resp.json();
        setShortResult(data.shorten_url);
    }
    return (
        <div>
            <h1>Shorten a URL</h1>
            <input
                value={inputUrl}
                onChange={e => setinputUrl(e.target.value)}
                placeholder="Enter long URL"
                className='border p-2 m-2'
            />
            <input
                value={customUrl}
                onChange={e => setcustomUrl(e.target.value)}
                placeholder="Enter custom Slug"
                className='border p-2 m-2'
            />
            <button onClick={handleShorten} className='px-3 py-2'>Shorten</button>
            <div>
                <div>
                    Short URL: <a href={shortResult} target="_blank" rel="noopener noreferrer">{shortResult}</a>
                </div>
            </div>
        </div>
        
    )
}

export default Shorten