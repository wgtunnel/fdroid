#!/bin/bash

cd metascoop
echo "::group::Building metascoop executable"
go build -o metascoop
echo "::endgroup::"

./metascoop -ap=../apps.yaml -rd=../fdroid/repo -pat="$GH_ACCESS_TOKEN" $1
EXIT_CODE=$?
cd ..

echo "Scoop had an exit code of $EXIT_CODE"

set -e

if [ $EXIT_CODE -eq 2 ]; then
    # Exit code 2 means that there were no significant changes
    echo "This means that there were no significant changes"
    exit 0
elif [ $EXIT_CODE -eq 0 ]; then
    # Exit code 0 means that we can commit everything & push
    echo "This means that we now have changes we should push"

    git config --global user.name 'github-actions'
    git config --global user.email 'dev@zaneschepke.com'

    # Configure the remote to use the PAT
    git remote set-url origin https://x-access-token:$GH_ACCESS_TOKEN@github.com/wgtunnel/fdroid.git

    git add .
    git commit -m "Automated update"
    git push
else 
    echo "This is an unexpected error"
    exit $EXIT_CODE
fi