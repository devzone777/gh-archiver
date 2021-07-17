#!/bin/bash

curl -X GET https://api.github.com/legacy/repos/search/language:go > gh_full_repo_list.json &

