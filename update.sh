#!/bin/bash

cd metascoop
echo "::group::Building metascoop executable"
go build -o metascoop
echo "::endgroup::"

echo "Running metascoop..."
./metascoop -ap=../apps.yaml -rd=../fdroid/repo -pat="$GH_ACCESS_TOKEN" $1
EXIT_CODE=$?

cd ..
echo "=== Scoop / metascoop exited with code: $EXIT_CODE ==="

if [ $EXIT_CODE -eq 2 ]; then
    echo "No significant changes detected. Nothing to do."
    exit 0
elif [ $EXIT_CODE -eq 0 ] || [ $EXIT_CODE -eq 1 ]; then
    echo "Significant changes detected (Repo.version bump is normal). Committing..."
    git config --global user.name 'github-actions'
    git config --global user.email '41898282+github-actions[bot]@users.noreply.github.com'
    
    git add .
    if git commit -m "Automated F-Droid repo update" > /dev/null 2>&1; then
        echo "Changes committed successfully."
        git push
    else
        echo "No new changes to commit (already up-to-date)."
    fi
    exit 0
else
    echo "Unexpected error from metascoop (exit code $EXIT_CODE)"
    exit $EXIT_CODE
fi
