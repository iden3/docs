# docs
iden3 documentation

[Work in progress]


## Offline use
In order to use and visualize the result in a local machine follow this two steps:


- install Docsifyjs cli:
```
npm i docsify-cli -g
```

- run local server:
```
docsify serve docs
```
This will serve the documentation website in a local port, reloading the webpage each time that a file is updated.

More details: https://docsify.js.org/#/quickstart

## Deploying on server
- just need to download this repository
```
git clone https://github.com/iden3/docs.git
```

- then run:
```
./gitconnect
```
This will automatically pull new commits added to this repo, to have always in the server the last version of this docs repositoy.

- run local server:
```
docsify serve docs
```
This will serve the documentation website in a local port, reloading the webpage each time that a file is updated.
