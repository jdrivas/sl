lib := sl
repo := github.com/jdrivas/$(lib)

help:
	@echo release \# push master branch to github and then do a local go update.

release: check
	@echo Pushing $(repo) to git and getting local copy of library to go env.
	git push
	go get -u $(repo)
