package resources

var EslintPath string = ".eslintrc.yml"

var Eslint string = `extends:
	- airbnb
	- 'plugin:redux-saga/recommended'
parser: babel-eslint
plugins:
	- react
	- jsx-a11y
	- import
	- redux-saga
	- jest
settings:
	import/resolver:
		webpack:
			config: webpack/webpack.client.js
rules:
	class-methods-use-this:
		- 0
	comma-dangle:
		- 2
		-
			imports: always-multiline
			exports: always-multiline
			arrays: always-multiline
			objects: always-multiline
			functions: never
	import/named: 2
	import/no-unresolved:
		- 0
	import/no-extraneous-dependencies:
		- 0
	import/extensions:
		- 0
	max-len:
		- 1
		- 80
		- 2
		-
			ignoreComments: true
	indent:
		- error
		- 4
		-
			SwitchCase: 1
	jsx-a11y/anchor-is-valid:
		- 0
	semi:
		- error
		- never
	padded-blocks:
		- error
		-
			switches: never
	no-bitwise:
		- error
		-
			allow:
				- '~'
	no-param-reassign:
		- error
		-
			props: false
	no-console: off
	dot-notation: off
	no-extra-boolean-cast: 2
	no-unused-expressions: off
	func-names: off
	eol-last: off
	react/prop-types:
		- 1
		-
			ignore:
				- t
	react/jsx-indent:
		- 1
		- 4
	react/jsx-indent-props:
		- 0
	react/jsx-filename-extension:
		- 1
		-
			extensions:
				- .js
				- .jsx
globals:
	__PRODUCTION__: false
	__DEV__: false
	__CLIENT__: false
	__SERVER__: false
env:
	browser: true
	node: true
	es6: true
	jest/globals: true`
