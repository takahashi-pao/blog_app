import * as fs from 'fs';
import * as https from 'https';
import express, { Express } from 'express';

const app: Express = express();

// TLS証明書と秘密鍵のパス
const tlsOptions: https.ServerOptions = {
  key: fs.readFileSync('path/to/private-key.pem'),
  cert: fs.readFileSync('path/to/certificate.pem'),
};

app.get('/', (req, res) => {
  res.send('Hello, secured world!');
});

const server: https.Server = https.createServer(tlsOptions, app);

// ポート番号
const port: number = 443;

server.listen(port, () => {
  console.log(`Server is running on https://localhost:${port}`);
});
