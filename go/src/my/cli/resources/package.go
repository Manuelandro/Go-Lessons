package resources

var PackagePath string = "package.json"

var Package string = `{
    "name": "project",
    "version": "1.0.0",
    "scripts": {
        "pack": "webpack --env.beep",
        "production": "NODE_ENV=production webpack"
    },
    "devDependencies": {
        "@babel/preset-env": "^7.0.0-beta.34",
        "babel-cli": "^6.18.0",
        "babel-core": "^6.9.0",
        "babel-eslint": "^8.0.1",
        "babel-loader": "^7.1.0",
        "babel-plugin-module-resolver": "^2.7.1",
        "babel-plugin-syntax-flow": "^6.18.0",
        "babel-plugin-transform-es2015-arrow-functions": "^6.22.0",
        "babel-plugin-transform-object-rest-spread": "^6.26.0",
        "babel-polyfill": "^6.9.1",
        "babel-preset-react": "^6.16.0",
        "brotli-webpack-plugin": "^0.4.1",
        "compression-webpack-plugin": "^1.0.1",
        "eslint": "^4.7.2",
        "eslint-config-airbnb": "^16.0.0",
        "eslint-import-resolver-webpack": "^0.8.3",
        "eslint-plugin-import": "^2.7.0",
        "eslint-plugin-jsx-a11y": "^6.0.2",
        "eslint-plugin-react": "^7.4.0",
        "webpack": "^3.7.1",
        "webpack-beep-plugin": "^1.0.0"
    },
    "dependencies": {
        "localforage": "^1.4.3",
        "lodash": "^4.17.4",
        "react": "^16.0.0",
        "react-dom": "^16.0.0",
        "react-redux": "^5.0.2",
        "redux": "^3.7.2",
        "redux-saga": "^0.15.6",
        "reduxsauce": "^0.7.0"
    }
}`
