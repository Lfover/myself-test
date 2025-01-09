const express = require('express');
const cors = require('cors');
const path = require('path');

const app = express();
const port = 8009;

// 启用 CORS
app.use(cors());

// 将 images 目录设置为静态文件目录
app.use('/images', express.static(path.join(__dirname, 'images')));

// 获取所有图片列表的 API
app.get('/api/images', (req, res) => {
    const fs = require('fs');
    const imagesDir = path.join(__dirname, 'images');

    fs.readdir(imagesDir, (err, files) => {
        if (err) {
            return res.status(500).json({ error: '无法读取图片目录' });
        }

        const images = files.filter(file =>
            ['.jpg', '.jpeg', '.png', '.gif'].includes(path.extname(file).toLowerCase())
        );

        res.json(images.map(image => `/images/${image}`));
    });
});

app.listen(port, () => {
    console.log(`服务器运行在 http://localhost:${port}`);
});