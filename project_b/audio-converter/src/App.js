import React, { useState, useRef } from 'react';
import './App.css';

function App() {
  const [audioUrl, setAudioUrl] = useState(null);
  const [loading, setLoading] = useState(false);
  const audioRef = useRef(null);

  const handleFileUpload = async (event) => {
    const file = event.target.files[0];
    if (!file) return;

    setLoading(true);
    const formData = new FormData();
    formData.append('audio', file);

    try {
      const response = await fetch('http://localhost:8009/convert', {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('转换失败');
      }

      const blob = await response.blob();
      const url = URL.createObjectURL(blob);
      setAudioUrl(url);

      // 自动播放转换后的音频
      if (audioRef.current) {
        audioRef.current.load();
        audioRef.current.play();
      }
    } catch (error) {
      console.error('错误:', error);
      alert('音频转换失败');
    } finally {
      setLoading(false);
    }
  };

  return (
      <div className="App">
        <header className="App-header">
          <h1>PCM 转 MP3 工具</h1>

          <div className="upload-container">
            <input
                type="file"
                accept=".pcm"
                onChange={handleFileUpload}
                disabled={loading}
            />
          </div>

          {loading && <p>正在转换中...</p>}

          {audioUrl && (
              <div className="audio-player">
                <audio ref={audioRef} controls>
                  <source src={audioUrl} type="audio/mp3" />
                  您的浏览器不支持音频播放
                </audio>
              </div>
          )}
        </header>
      </div>
  );
}

export default App;