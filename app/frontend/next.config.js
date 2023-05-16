/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  output: "export",
  distDir: "../backend/views",
  assetPrefix: process.env.STATIC_FILE_URL,
};

module.exports = nextConfig;
