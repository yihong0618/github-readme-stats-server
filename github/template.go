package github

var Template = `
<html>
  <head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
    <meta
      name="viewport"
      content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no"
    />
    <meta name="description" content="" />
    <meta name="keywords" content="" />
    <link
      href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css"
      rel="stylesheet"
      integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3"
      crossorigin="anonymous"
    />
    <meta name="viewport" content="width=device-width, initial-scale=1, minimal-ui">
    <title>{{.Title}}</title>
    <style>
      body {
        min-width: 200px;
        max-width: 790px;
        margin: 0 auto;
        padding: 30px;
      }
.main-link {
  text-decoration: none;
  color: inherit;
}
.markdown-body {
  -webkit-text-size-adjust: 100%;
  -ms-text-size-adjust: 100%;
  text-size-adjust: 100%;
  color: #333;
  overflow: hidden;
  font-family: "Helvetica Neue", Helvetica, "Segoe UI", Arial, freesans, sans-serif;
  font-size: 16px;
  line-height: 1.6;
  word-wrap: break-word;
}
.markdown-body a {
  background-color: transparent;
}
.markdown-body a:active,
.markdown-body a:hover {
  outline: 0;
}
.markdown-body strong {
  font-weight: bold;
}
.markdown-body h1 {
  font-size: 2em;
  margin: 0.67em 0;
}
.markdown-body hr {
  box-sizing: content-box;
  height: 0;
}
.markdown-body table {
  border-collapse: collapse;
  border-spacing: 0;
}
.markdown-body td,
.markdown-body th {
  padding: 0;
}
.markdown-body * {
  box-sizing: border-box;
}
.markdown-body a {
  color: #4078c0;
  text-decoration: none;
}
.markdown-body a:hover,
.markdown-body a:active {
  text-decoration: underline;
}
.markdown-body hr {
  height: 0;
  margin: 15px 0;
  overflow: hidden;
  background: transparent;
  border: 0;
  border-bottom: 1px solid #ddd;
}
.markdown-body hr:before {
  display: table;
  content: "";
}
.markdown-body hr:after {
  display: table;
  clear: both;
  content: "";
}
.markdown-body h1,
.markdown-body h2,
.markdown-body h1 {
  font-size: 30px;
}
.markdown-body h2 {
  font-size: 21px;
}
.markdown-body blockquote {
  margin: 0;
}
.markdown-body dd {
  margin-left: 0;
}
.markdown-body code {
  font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
  font-size: 12px;
}
.markdown-body pre {
  margin-top: 0;
  margin-bottom: 0;
  font: 12px Consolas, "Liberation Mono", Menlo, Courier, monospace;
}
.markdown-body .octicon {
  font: normal normal normal 16px/1 octicons-anchor;
  display: inline-block;
  text-decoration: none;
  text-rendering: auto;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
.markdown-body .octicon-link:before {
  content: '\f05c';
}
.markdown-body>*:first-child {
  margin-top: 0 !important;
}
.markdown-body>*:last-child {
  margin-bottom: 0 !important;
}
.markdown-body a:not([href]) {
  color: inherit;
  text-decoration: none;
}
.markdown-body .anchor {
  position: absolute;
  top: 0;
  left: 0;
  display: block;
  padding-right: 6px;
  padding-left: 30px;
  margin-left: -30px;
}
.markdown-body .anchor:focus {
  outline: none;
}
.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4,
.markdown-body h5,
.markdown-body h6 {
  position: relative;
  margin-top: 1em;
  margin-bottom: 16px;
  font-weight: bold;
  line-height: 1.4;
}
.markdown-body h1 .octicon-link,
.markdown-body h2 .octicon-link,
.markdown-body h3 .octicon-link,
.markdown-body h4 .octicon-link,
.markdown-body h5 .octicon-link,
.markdown-body h6 .octicon-link {
  display: none;
  color: #000;
  vertical-align: middle;
}
.markdown-body h1:hover .anchor,
.markdown-body h2:hover .anchor,
.markdown-body h3:hover .anchor,
.markdown-body h4:hover .anchor,
.markdown-body h5:hover .anchor,
.markdown-body h6:hover .anchor {
  padding-left: 8px;
  margin-left: -30px;
  text-decoration: none;
}
.markdown-body h1:hover .anchor .octicon-link,
.markdown-body h2:hover .anchor .octicon-link,
.markdown-body h3:hover .anchor .octicon-link,
.markdown-body h4:hover .anchor .octicon-link,
.markdown-body h5:hover .anchor .octicon-link,
.markdown-body h6:hover .anchor .octicon-link {
  display: inline-block;
}
.markdown-body h1 {
  padding-bottom: 0.3em;
  font-size: 2.25em;
  line-height: 1.2;
  border-bottom: 1px solid #eee;
}
.markdown-body h1 .anchor {
  line-height: 1;
}
.markdown-body h2 {
  padding-bottom: 0.3em;
  font-size: 1.75em;
  line-height: 1.225;
  border-bottom: 1px solid #eee;
}
.markdown-body h2 .anchor {
  line-height: 1;
}
.markdown-body h3 {
  font-size: 1.5em;
  line-height: 1.43;
}
.markdown-body h3 .anchor {
  line-height: 1.2;
}
.markdown-body h4 {
  font-size: 1.25em;
}
.markdown-body h4 .anchor {
  line-height: 1.2;
}
.markdown-body h5 {
  font-size: 1em;
}
.markdown-body h5 .anchor {
  line-height: 1.1;
}
.markdown-body h6 {
  font-size: 1em;
  color: #777;
}
.markdown-body h6 .anchor {
  line-height: 1.1;
}
.markdown-body p,
.markdown-body blockquote,
.markdown-body ul,
.markdown-body ol,
.markdown-body dl,
.markdown-body table,
.markdown-body pre {
  margin-top: 0;
  margin-bottom: 16px;
}
.markdown-body hr {
  height: 4px;
  padding: 0;
  margin: 16px 0;
  background-color: #e7e7e7;
  border: 0 none;
}
.markdown-body ul,
.markdown-body ol {
  padding-left: 2em;
}
.markdown-body ul ul,
.markdown-body ul ol,
.markdown-body ol ol,
.markdown-body ol ul {
  margin-top: 0;
  margin-bottom: 0;
}
.markdown-body li>p {
  margin-top: 16px;
}
.markdown-body dl {
  padding: 0;
}
.markdown-body dl dt {
  padding: 0;
  margin-top: 16px;
  font-size: 1em;
  font-style: italic;
  font-weight: bold;
}
.markdown-body dl dd {
  padding: 0 16px;
  margin-bottom: 16px;
}
.markdown-body blockquote {
  padding: 0 15px;
  color: #777;
  border-left: 4px solid #ddd;
}
.markdown-body blockquote>:first-child {
  margin-top: 0;
}
.markdown-body blockquote>:last-child {
  margin-bottom: 0;
}
.markdown-body table {
  display: block;
  width: 100%;
  overflow: auto;
  word-break: normal;
  word-break: keep-all;
}
.markdown-body table th {
  font-weight: bold;
}
.markdown-body table th,
.markdown-body table td {
  padding: 6px 13px;
  border: 1px solid #ddd;
}
.markdown-body table tr {
  background-color: #fff;
  border-top: 1px solid #ccc;
}
.markdown-body table tr:nth-child(2n) {
  background-color: #f8f8f8;
}
.markdown-body img {
  max-width: 100%;
  box-sizing: border-box;
}
.markdown-body code {
  padding: 0;
  padding-top: 0.2em;
  padding-bottom: 0.2em;
  margin: 0;
  font-size: 85%;
  background-color: rgba(0,0,0,0.04);
  border-radius: 3px;
}
.markdown-body code:before,
.markdown-body code:after {
  letter-spacing: -0.2em;
  content: "\00a0";
}
.markdown-body pre>code {
  padding: 0;
  margin: 0;
  font-size: 100%;
  word-break: normal;
  white-space: pre;
  background: transparent;
  border: 0;
}
.markdown-body .highlight {
  margin-bottom: 16px;
}
.markdown-body .highlight pre,
.markdown-body pre {
  padding: 16px;
  overflow: auto;
  font-size: 85%;
  line-height: 1.45;
  background-color: #f7f7f7;
  border-radius: 3px;
}
.markdown-body .highlight pre {
  margin-bottom: 0;
  word-break: normal;
}
.markdown-body pre {
  word-wrap: normal;
}
.markdown-body pre code {
  display: inline;
  max-width: initial;
  padding: 0;
  margin: 0;
  overflow: initial;
  line-height: inherit;
  word-wrap: normal;
  background-color: transparent;
  border: 0;
}
.markdown-body pre code:before,
.markdown-body pre code:after {
  content: normal;
}
}
    </style>
  </head>
<body>
	<div class="container p-4">
	<main>
	<div class="heading text-center">
		<h1 class="display-1">
		<a
			href="https://github.com/yihong0618/github-readme-stats-server"
			class="main-link"
			>GitHub README Stats</a
		>
		</h1>
		<p class="lead">Generate GitHub User README Profile</p>
	</div>
	<hr />
	<form action="/generate" method="post">
		<div class="input-group mb-4">
		<input
			class="form-control form-control-lg"
			type="text"
			placeholder="GitHub username"
			aria-label="Input GitHub username"
			name="r"
			id="r"
		/>
		<button class="btn btn-primary">GO!</button>
		</div>
	</form>
	</div>
	</main>
  <div class="tips">
    <h4>Tips</h4>
    <ul>
      <li>
        Add query string <code>--refresh</code> to clear the caches
      </li>
    </ul>
  </div>
  </main>
  <footer>
  <h4>Credits</h4>
  <ul>
    <li>
      <a href="https://github.com/frostming">frostming</a>'s
      <a href="https://github.com/frostming/tokei-pie-cooker">tokei-pie-cooker</a> project
    </li>
  </ul>
  </footer>
	</div>
  <article class="markdown-body">
    {{.Body}}
  </article>
</body>
</html>
`
