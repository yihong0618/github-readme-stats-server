package github

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/google/go-github/v41/github"
	"github.com/olekukonko/tablewriter"
	"github.com/shurcooL/github_flavored_markdown"
)

var (
	staredNumber int  = 10
	withStared   bool = true
)

type contextHTML struct {
	Title string
	Body  string
}

var baseURL = "https://github.com/"
var myCreatedTitle = "## The repos I created\n"
var myContributedTitle = "## The repos I contributed to\n"

var htmlTemplate = `
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, minimal-ui">
    <title>{{.Title}}</title>
    <style>
      body {
        min-width: 200px;
        max-width: 790px;
        margin: 0 auto;
        padding: 30px;
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
  <article class="markdown-body">
    {{.Body}}
  </article>
</html>
`

type myRepoInfo struct {
	star     int
	name     string
	HTMLURL  string
	create   string
	update   string
	lauguage string
}

func (r *myRepoInfo) mdName() string {
	return "[" + r.name + "]" + "(" + r.HTMLURL + ")"
}

type myPrInfo struct {
	name       string
	repoURL    string
	fisrstDate string
	lasteDate  string
	prCount    int
}

func (p *myPrInfo) mdName() string {
	return "[" + p.name + "]" + "(" + p.repoLink() + ")"
}

func (p *myPrInfo) repoLink() string {
	q := strings.Split(p.repoURL, "/")
	return baseURL + q[len(q)-2] + "/" + q[len(q)-1]
}

func getAllPrLinks(p myPrInfo, userName string) string {
	url := fmt.Sprintf("%s/pulls?q=is:pr+author:%s", p.repoLink(), userName)
	return "https://" + strings.ReplaceAll(strings.Split(url, "https://")[1], ":", "%3A")
}

type myStaredInfo struct {
	staredDate string
	desc       string
	myRepoInfo
}

func getRepoName(RepositoryURL string) string {
	q := strings.Split(RepositoryURL, "/")
	return q[len(q)-1]
}

func fetchAllCreatedRepos(username string, client *github.Client) []*github.Repository {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(context.Background(), username, opt)
		if err != nil {
			fmt.Println("Something wrong to get repos")
			continue
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allRepos
}

func makeCreatedRepos(repos []*github.Repository) ([]myRepoInfo, int, int) {
	totalCount := 0
	longest := 0
	myRepos := []myRepoInfo{}
	for _, repo := range repos {
		if !*repo.Fork {
			create := (*repo.CreatedAt).String()[:10]
			update := (*repo.UpdatedAt).String()[:10]
			lauguage := "md"
			if repo.Language != nil {
				lauguage = *repo.Language
			}
			myRepos = append(myRepos, myRepoInfo{
				star:     *repo.StargazersCount,
				name:     *repo.Name,
				create:   create,
				update:   update,
				lauguage: lauguage,
				HTMLURL:  *repo.HTMLURL,
			})
			totalCount = totalCount + *repo.StargazersCount
			if len(*repo.Name) > longest {
				longest = len(*repo.Name)
			}
		}
	}
	return myRepos, totalCount, longest
}

func fetchAllPrIssues(username string, client *github.Client) []*github.Issue {
	nowPage := 100
	opt := &github.SearchOptions{ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	var allIssues []*github.Issue
	for {
		result, _, err := client.Search.Issues(context.Background(), fmt.Sprintf("is:pr author:%s is:closed is:merged", username), opt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		allIssues = append(allIssues, result.Issues...)
		if nowPage >= result.GetTotal() {
			break
		}
		opt.Page = opt.Page + 1
		nowPage = nowPage + 100
	}
	return allIssues
}

func makePrRepos(issues []*github.Issue) []myPrInfo {
	prMap := make(map[string]map[string]interface{})
	for _, issue := range issues {
		if *issue.AuthorAssociation == "OWNER" {
			continue
		}
		repoName := getRepoName(*issue.RepositoryURL)
		if len(prMap[repoName]) == 0 {
			prMap[repoName] = make(map[string]interface{})
			prMap[repoName]["prCount"] = 1
			prMap[repoName]["fisrstDate"] = (*issue.CreatedAt).String()[:10]
			prMap[repoName]["lasteDate"] = (*issue.CreatedAt).String()[:10]
			prMap[repoName]["repoURL"] = *issue.RepositoryURL
		} else {
			prMap[repoName]["prCount"] = prMap[repoName]["prCount"].(int) + 1
			if prMap[repoName]["fisrstDate"].(string) > (*issue.CreatedAt).String()[:10] {
				prMap[repoName]["fisrstDate"] = (*issue.CreatedAt).String()[:10]
			}
			if prMap[repoName]["lasteDate"].(string) < (*issue.CreatedAt).String()[:10] {
				prMap[repoName]["lasteDate"] = (*issue.CreatedAt).String()[:10]
			}
		}
	}
	myPrs := []myPrInfo{}
	for k, v := range prMap {
		myPrs = append(myPrs, myPrInfo{
			name:       k,
			repoURL:    v["repoURL"].(string),
			fisrstDate: v["fisrstDate"].(string),
			lasteDate:  v["lasteDate"].(string),
			prCount:    v["prCount"].(int),
		})
	}
	return myPrs
}

func fetchRecentStared(username string, client *github.Client) []*github.StarredRepository {
	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{Page: 1, PerPage: 100},
	}
	var allStared []*github.StarredRepository
	repos, _, err := client.Activity.ListStarred(context.Background(), username, opt)
	if err != nil {
		fmt.Println("Something wrong to get stared")
	}
	allStared = append(allStared, repos...)
	return allStared
}

func makeStaredRepos(stars []*github.StarredRepository) []myStaredInfo {
	myStars := []myStaredInfo{}
	for _, star := range stars {
		repo := *star.Repository
		lauguage := "md"
		if repo.Language != nil {
			lauguage = *repo.Language
		}
		desc := ""
		if repo.Description != nil {
			desc = *repo.Description
		}

		myStars = append(myStars, myStaredInfo{
			staredDate: (*star.StarredAt).String()[:10],
			desc:       desc,
			myRepoInfo: myRepoInfo{
				name:     *repo.Name,
				create:   (*repo.CreatedAt).String()[:10],
				update:   (*repo.UpdatedAt).String()[:10],
				lauguage: lauguage,
				HTMLURL:  *repo.HTMLURL,
			},
		})
	}
	// shffle to get random array
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(myStars), func(i, j int) { myStars[i], myStars[j] = myStars[j], myStars[i] })
	return myStars
}

func makeMdTable(data [][]string, header []string) string {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	return tableString.String()
}

func makeCreatedString(repos []myRepoInfo) string {
	starsData := [][]string{}
	for i, repo := range repos {
		starsData = append(starsData, []string{strconv.Itoa(i + 1), repo.mdName(), repo.create, repo.update, repo.lauguage, strconv.Itoa(repo.star)})
	}
	myStarsString := makeMdTable(starsData, []string{"ID", "Repo", "Start", "Update", "Lauguage", "Stars"})
	return myCreatedTitle + myStarsString + "\n"
}

func makeContributedString(myPRs []myPrInfo, userName string) string {
	prsData := [][]string{}
	for i, pr := range myPRs {
		prsData = append(prsData, []string{strconv.Itoa(i + 1), pr.mdName(), pr.fisrstDate, pr.lasteDate, fmt.Sprintf("[%d](%s)", pr.prCount, getAllPrLinks(pr, userName))})
	}
	myPrString := makeMdTable(prsData, []string{"ID", "Repo", "firstDate", "lasteDate", "prCount"})
	return myContributedTitle + myPrString + "\n"
}

func makeStaredString(myStars []myStaredInfo, starNumber int) string {
	myStaredTitle := fmt.Sprintf("## The repos I stared (random %s)", strconv.Itoa(starNumber)) + "\n"
	starsData := [][]string{}
	// maybe a better way in golang?
	if (len(myStars)) < starNumber {
		starNumber = len(myStars)
	}
	for i, star := range myStars[:starNumber] {
		repo := star.myRepoInfo
		starsData = append(starsData, []string{strconv.Itoa(i + 1), repo.mdName(), star.staredDate, repo.lauguage, repo.update})
	}
	myStaredString := makeMdTable(starsData, []string{"ID", "Repo", "staredDate", "Lauguage", "LatestUpdate"})
	return myStaredTitle + myStaredString + "\n"
}

func GenerateNewFile(UserName string) {
	flag.Parse()
	client := github.NewClient(nil)
	repos := fetchAllCreatedRepos(UserName, client)
	myRepos, totalCount, longest := makeCreatedRepos(repos)
	fmt.Println(totalCount, longest)
	// change sort logic here
	sort.Slice(myRepos[:], func(i, j int) bool {
		return myRepos[j].star < myRepos[i].star
	})

	issues := fetchAllPrIssues(UserName, client)
	myPRs := makePrRepos(issues)
	// change sort logic here
	sort.Slice(myPRs[:], func(i, j int) bool {
		return myPRs[j].prCount < myPRs[i].prCount
	})
	myStaredString := ""
	if withStared {
		stars := fetchRecentStared(UserName, client)
		myStared := makeStaredRepos(stars)
		myStaredString = makeStaredString(myStared, staredNumber)
	}

	myCreatedString := makeCreatedString(myRepos)
	myPrString := makeContributedString(myPRs, UserName)

	newContentString := myCreatedString + myPrString
	if withStared {
		newContentString = newContentString + myStaredString
	}
	tmpl, err := template.New("markdown").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("")
	}
	result := github_flavored_markdown.Markdown([]byte(newContentString))
	outputFile, _ := os.Create("templates/" + UserName + ".html")
	err = tmpl.Execute(outputFile, contextHTML{Title: UserName, Body: string(result)})
	if err != nil {
		log.Fatal(err)
	}
}
