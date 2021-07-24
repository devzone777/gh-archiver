#!/bin/bash

#curl -X GET https://api.github.com/legacy/repos/search/language:go > gh_full_repo_list.json

curl -u devzone777:$token -X GET https://api.github.com/legacy/repos/search/keyword:$1 > gh_full_repo_list.json

go run getrepolist_legacy.go > repolist.txt

sed -n "0~3 p" repolist.txt > repoqueue.dat

go run repomonitor_legacy.go
