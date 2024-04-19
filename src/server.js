import dotenv from 'dotenv';
dotenv.config();

import express from 'express';
import path from 'path';

import HomePage from "./views/pages/HomePage.js";

const dir = import.meta.dirname;
const app = express();
const PORT = process.env.PORT || 3000;

const ASSETS_IMMUTABLE = process.env.ASSETS_IMMUTABLE || false;
const ASSETS_CACHING_TIME = process.env.ASSETS_CACHING_TIME || 0;

app.use((req, res, next) => {
    console.log(`[${req.method}] ${req.url}`);
    next();
})

app.use('/assets', express.static(path.join(dir, '../assets'), {
    immutable: ASSETS_IMMUTABLE,
    maxAge: ASSETS_CACHING_TIME
}));

app.all('/', (req, res) => {
    res.send(HomePage());
})

app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
