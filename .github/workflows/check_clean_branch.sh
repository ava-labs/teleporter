#!/bin/bash
# Exits if any uncommitted changes are found.

set -o errexit
set -o nounset
set -o pipefail

# Modifies the index based on the current index.
# Checks to see if merges or updates are needed by checking stat() information.
git update-index --refresh > /dev/null
# Checks to see if there are any differences from the index to the current HEAD commit
git diff-index --quiet HEAD