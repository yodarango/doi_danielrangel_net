# DOI

Self-hosted DOI system to preserve locality of my academic papers. The format followed to create them is <MMDDYYYY><Alphabetical sequence (for posts published the same day)><Language code>. Example -> July 2nd, 1977 (A) Italian -> `02071977AITA`

## DEV

1. `GOOS=linux GOARCH=amd64 go build -o doi_danielrangel_net`
2. `scp ./doi.json, doi_danielrangel_net user@<pp>:/var/www/repos/doi_danielrangel_net/app`

## TODO

- Auto-generate the doi.json file from my posts project
