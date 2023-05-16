/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  output: "export",
  distDir: "../backend/views",
};

module.exports = nextConfig;
