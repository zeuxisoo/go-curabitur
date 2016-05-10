var webpack = require('webpack');
var discardComments = require('postcss-discard-comments')
var ExtractTextPlugin = require("extract-text-webpack-plugin");

var config = require('./webpack.config');

config.output.filename = 'bundle.js';
config.devtool         = 'hidden-source-map';

config.plugins.push(
    new ExtractTextPlugin("bundle.css", {
        disable: false
    }),
    new webpack.optimize.OccurenceOrderPlugin(),
    new webpack.DefinePlugin({
        '__DEV__': false,
        'process.env': {
            'NODE_ENV': JSON.stringify('production')
        }
    }),
    new webpack.optimize.UglifyJsPlugin({
        output: {
            comments: false
        },
        compressor: {
            screw_ie8: true,
            warnings: false
        }
    })
);

config.postcss = function() {
    return [
        discardComments({
            removeAll: true
        })
    ]
}

module.exports = config;
