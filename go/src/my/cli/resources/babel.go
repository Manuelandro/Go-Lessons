package resources

var BabelPath string = ".babelrc"

var Babel string = `{
    "presets": [
        [
            "@babel/preset-env",
            {
                "targets": {
                    "browsers": ["last 4 versions", "ie >= 9"],
                    "node": "current"
                },
                "useBuiltIns": true,
                "debug": true
            }
        ],
        "react"
    ],
    "ignore": ["node_modules"],
    "plugins": [
        "transform-object-rest-spread",
        "transform-es2015-arrow-functions"
    ],
    "sourceMaps": true,
    "retainLines": true
}
`
