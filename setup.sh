# /bin/bash

# Clone the repo holding shared snippets
if [ ! -d "hg-snippets-config" ]; then
    git clone git@github.com:grafana/hg-snippets-config.git
else
    echo "hg-snippets-config repo already exists"
fi
