# github-readme-stats-server
Server for [github-readme-stats](https://github.com/yihong0618/github-readme-stats)


---

[Demo Website](https://github-readme-stats-d7t3.onrender.com/) 🚀

## Start locally

- go 1.14 with go mod
- go run main.go
- with token GITHUB_TOKE=xxxx go run main.go

### As a container

```bash
docker build . -t gin_readme
docker run gin_readme -p 8080:8080 

# Visit http://localhost:8080
```

## Credits

- [tokei-pie-cooker](https://github.com/frostming/tokei-pie-cooker)

## Special Thanks 

- [frostming](https://github.com/frostming)
