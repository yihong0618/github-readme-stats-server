package github

import (
	"context"
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

var htmlTemplate = Template

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
			fmt.Println(username, "Something wrong to get repos", err)
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

func makeCreatedRepos(repos []*github.Repository) ([]myRepoInfo, int) {
	totalCount := 0
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
		}
	}
	return myRepos, totalCount
}

func fetchAllPrIssues(username string, client *github.Client) []*github.Issue {
	nowPage := 100
	opt := &github.SearchOptions{ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	var allIssues []*github.Issue
	for {
		result, _, err := client.Search.Issues(context.Background(), fmt.Sprintf("is:pr author:%s", username), opt)
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
		if nowPage >= 1000 {
			// api only support first 1000
			break
		}
	}
	return allIssues
}

func makePrRepos(issues []*github.Issue) ([]myPrInfo, int) {
	totalPrCount := 0
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
		totalPrCount++
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
	return myPrs, totalPrCount
}

func fetchRecentStared(username string, client *github.Client) []*github.StarredRepository {
	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{Page: 1, PerPage: 100},
	}
	var allStared []*github.StarredRepository
	repos, _, err := client.Activity.ListStarred(context.Background(), username, opt)
	if err != nil {
		fmt.Println("Something wrong to get stared", err)
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

func makeCreatedString(repos []myRepoInfo, userName string, total int) string {
	starsData := [][]string{}
	for i, repo := range repos {
		starsData = append(starsData, []string{strconv.Itoa(i + 1), repo.mdName(), repo.create, repo.update, repo.lauguage, strconv.Itoa(repo.star)})
	}
	starsData = append(starsData, []string{"sum", "", "", "", "", strconv.Itoa(total)})
	myStarsString := makeMdTable(starsData, []string{"ID", "Repo", "Start", "Update", "Lauguage", "Stars"})
	myCreatedTitle := fmt.Sprintf("## The repos %s created\n", userName)
	return myCreatedTitle + myStarsString + "\n"
}

func makeContributedString(myPRs []myPrInfo, userName string, total int) string {
	prsData := [][]string{}
	for i, pr := range myPRs {
		prsData = append(prsData, []string{strconv.Itoa(i + 1), pr.mdName(), pr.fisrstDate, pr.lasteDate, fmt.Sprintf("[%d](%s)", pr.prCount, getAllPrLinks(pr, userName))})
	}
	prsData = append(prsData, []string{"sum", "", "", "", strconv.Itoa(total)})
	myContributedTitle := fmt.Sprintf("## The repos %s contributed to\n", userName)
	myPrString := makeMdTable(prsData, []string{"ID", "Repo", "firstDate", "lasteDate", "prCount"})
	return myContributedTitle + myPrString + "\n"
}

func makeStaredString(myStars []myStaredInfo, starNumber int, userName string) string {
	myStaredTitle := fmt.Sprintf("## The repos %s recent stared (random %s)", userName, strconv.Itoa(starNumber)) + "\n"
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
	client := github.NewClient(nil)
	repos := fetchAllCreatedRepos(UserName, client)
	myRepos, totalCount := makeCreatedRepos(repos)
	// change sort logic here
	sort.Slice(myRepos[:], func(i, j int) bool {
		return myRepos[j].star < myRepos[i].star
	})

	issues := fetchAllPrIssues(UserName, client)
	myPRs, totalPrCount := makePrRepos(issues)
	// change sort logic here
	sort.Slice(myPRs[:], func(i, j int) bool {
		return myPRs[j].prCount < myPRs[i].prCount
	})
	myStaredString := ""
	if withStared {
		stars := fetchRecentStared(UserName, client)
		myStared := makeStaredRepos(stars)
		myStaredString = makeStaredString(myStared, staredNumber, UserName)
	}

	myCreatedString := makeCreatedString(myRepos, UserName, totalCount)
	myPrString := makeContributedString(myPRs, UserName, totalPrCount)

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
