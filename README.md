
# go-popular-repos

Just a learning Go app to list Github repos sorted by stars, that can be deployed to Heroku.

Based on [Heroku tutorial](https://devcenter.heroku.com/articles/getting-started-with-go).

example heroku app: https://powerful-beach-30230.herokuapp.com

## Running Locally

### Setup
* Install [Go](http://golang.org/doc/install) 1.6+: `brew install go`
* install [Heroku Toolbelt](https://toolbelt.heroku.com/) for heroku deployment.
* install godep: `go get -u github.com/tools/godep`

```sh
$ go get -u github.com/alexchee/go-popular-repos/...
$ cd $GOPATH/src/github.com/alexchee/go-popular-repos
$ godep restore
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

### Notes (for first Go/Gin timers):
* Need to rebuild/install cmd and restart server to see latest changes to code:
  in `cmd/go-popular-repos/` directory: `$ go install ./...`
  in project's root directory: `$ heroku local`
* Do not need to restart server for template changes.
* If packages are added/removed, need to save new dependencies:
  in project's root directory: `$ godep save ./...`

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```
or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## TODO
* truncate description
* add search for repos
* style
