import React, { useState, useEffect } from 'react';
import '../styles/Gallery.css';

function Gallery() {
    const [images, setImages] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8009/api/images')
            .then(response => response.json())
            .then(data => setImages(data))
            .catch(error => console.error('Error:', error));
    }, []);

    return (
        <div className="gallery">
            <h1>图片展示</h1>
            <div className="image-grid">
                {images.map((image, index) => (
                    <div key={index} className="image-item">
                        <img
                            src={`http://localhost:8009${image}`}
                            alt={`图片 ${index + 1}`}
                        />
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Gallery;