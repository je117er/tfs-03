const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
// const isDevelopment = process.env.NODE_ENV === 'development'

module.exports = {
    entry: [
        './src/style.scss'
    ],
    module: {
        rules: [
            {
                test: /\.(sa|sc)ss$/,
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader',
                    'sass-loader'
                ]
            },
        ]
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: '../dist/style.min.css',
        })
    ]
};