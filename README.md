# GitDeployer
[![Build Status](https://travis-ci.com/ZacharyGroff/GitDeployer.svg?branch=master)](https://travis-ci.com/ZacharyGroff/GitDeployer) 
[![Hits-of-Code](https://hitsofcode.com/github/ZacharyGroff/GitDeployer)](https://hitsofcode.com/view/github/ZacharyGroff/GitDeployer)

> Automation service that hooks into your GitHub repository's merge events and 
automatically runs a deploy script when a branch is merged to master.

Originally the plan for this project was to build a suite of tools around the
events that GitHub's webhooks supported. After diving more into this project we
found that much of the functionality we were hoping to implement is already 
natively available in the GitHub platform. Because of this we have just decided
to release the auto deployment functionality on its own.

## Installation

Before installing **GitDeployer** you will need to install Golang [here](https://golang.org/doc/install).

You can then install the binary by running this command:
```
go get github.com/ZacharyGroff/GitDeployer
```

## Running the Code

After installing you will need to create a `conf.json` file in the same 
directory as the binary that you just downloaded.

Here is the skeleton that you should use:
```
{
	"port": ":8080",
	"route": "/webhook",
	"scriptPath": "test.sh",
	"validate": true
}
```

The port and route are used for running the webhook endpoint that GitHub will 
post events to. The script is the path to the script that you want the service 
to execute when a branch is merged to master.

## Contributing

Please feel free to open an issue if you wish to add any features to the 
codebase!

## License

