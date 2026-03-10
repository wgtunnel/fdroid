#!/bin/bash
set -e

cd metascoop
echo "::group::Building metascoop executable"
go build -o metascoop
echo "::endgroup::"

./metascoop -ap=../apps.yaml -rd=../fdroid/repo -pat="$GH_ACCESS_TOKEN" $1
EXIT_CODE=$?

cd ..
echo "Scoop had an exit code of $EXIT_CODE"

if [ $EXIT_CODE -eq 2 ]; then
    # Exit code 2 means that there were no significant changes
    echo "This means that there were no significant changes"
    exit 0
elif [ $EXIT_CODE -eq 0 ] || [ $EXIT_CODE -eq 1 ]; then
    # Exit code 0 or 1 is normal (Repo.version bump from fdroid update)
    echo "This means that we now have changes we should push"
    git config --global user.name 'github-actions'
    git config --global user.email '41898282+github-actions[bot]@users.noreply.github.com'
    git add .
    git commit -m "Automated update" || true   # safe if nothing changed
    git push
    exit 0
else
    echo "This is an unexpected error"
    exit $EXIT_CODE
fi
