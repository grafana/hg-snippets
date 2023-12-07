<img src="doc/grappet.jpg" style="max-height: 700px; max-width: 700px;">

Grappet is a snippet manager for Hosted Grafana that allows us to share snippets via the [hg-snippets-config](https://github.com/grafana/hg-snippets-config) repo.
It is a fork of [Pet](https://github.com/knqyf263/pet).

### Quickstart
Clone the repo, run the setup, and list the commands:
```bash
git clone git@github.com:grafana/hg-snippets.git
cd hg-snippets
make setup
./grappet list
```

### Syncing snippets
All our snippets are stored in a private repo, [hg-snippets-config](https://github.com/grafana/hg-snippets-config). This repo is nested within this one. To update snippets, just pull in changes from that repo:
```bash
cd ./hg-snippets-config && git pull
```

To add your snippets to shared repo, just commit them, push, and make a PR:
```bash
cd ./hg-snippets-config && git push
```
