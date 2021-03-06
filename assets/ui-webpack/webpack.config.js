const path = require('path');
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const MiniCSSExtractPlugin = require('mini-css-extract-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const TerserJSPlugin = require('terser-webpack-plugin');

module.exports = (env, options) => {
    const config = {
        entry: {
            edit: './src/js/edit.js',
            config: './src/js/config.js',
            list: './src/js/list.js',
            signin: './src/js/signin.js',
            plugins: './src/js/plugins.js',
            plugin: './src/js/plugin.js',
            profile: './src/js/profile.js',
            users: './src/js/users.js'
        },
        output: {
            filename: '[name].bundle.js',
            path: path.resolve(__dirname, 'dist')
        },
        module: {
            rules: [
                {
                    test: /\.css$/i,
                    use: [
                        MiniCSSExtractPlugin.loader,
                        'css-loader'
                    ]
                }, {
                    test: /\.png$/,
                    loader: 'file-loader'
                }, {
                    test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
                    loader: "url-loader?limit=10000&mimetype=application/font-woff"
                }, {
                    test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
                    loader: "file-loader"
                }, {
                    test: /\.html$/,
                    include: [
                        path.resolve(__dirname, "src/html")
                    ],
                    loader: "html-loader",
                    options: {
                        minimize: true
                    }
                }
            ]
        },
        plugins: [
            new CleanWebpackPlugin({}),
            new MiniCSSExtractPlugin({
                filename: '[name].bundle.css'
            })
        ]
    };

    if (options.mode === "production") {
        config.optimization = {
            minimizer: [
                new TerserJSPlugin({}),
                new OptimizeCSSAssetsPlugin({})
            ]
        }
    }

    return config;
};

