// @ts-check
const withPlugins = require('next-compose-plugins');
const path = require('path');
const withImages = require('next-images');
const imagePlugin = withImages({
  webpack(config) {
    return config;
  },
  images: {
    domains: ['picsum.photos'],
  },
});

/**
 * @type {import('next/dist/next-server/server/config').NextConfig}
 **/
const nextConfig = withPlugins([
  imagePlugin,
  {
    reactStrictMode: true,
    sassOptions: {
      includePaths: [path.join(__dirname, 'styles')],
    },
    experimental: {
      optimizeFonts: true,
    },
    eslint: {
      dirs: ['pages', 'utils'],
    },
  },
]);

module.exports = nextConfig;
