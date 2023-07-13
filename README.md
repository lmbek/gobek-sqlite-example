# gobek (Local GUI / Go Chrome Framework)

gobek is a simple framework made for developing localhosted software that can reuse chrome/chromium or embed chromium (in future releases). Both available in deployment for the applications.

This framework uses Chrome (Windows) or Chromium (Linux) as frontend by opening them with cmd/terminal and hosting a localhost webserver, while opening chrome/chromium with --app and --user-directory arguments. The frontend can be changed by the user in runtime, while the backend needs to be compiled/build. The API can be decoupled in future versions, so every part of the application is changeable - Sustainable development. Frontends is easy to change. Alternatives to this is embedding a chromium or webview framework into the project, which will require more space. I chose to depend on Chrome/Chromium, as they are market leaders and html/css/javascript technology frontrunners.

Feel free to use this piece of software, I will be happy to assist you <br>
I am currently working on this project, it will be updated and maintained. I consider it production ready. <br>
<br>
This project is used by Beksoft ApS for projects such as:
* BekCMS
* PingPong Game made in Three.js
* Several local webbased software projects
  <br><br>
Write to me at lars@beksoft.dk if you want to have your project listed<br>
Be aware that i am mostly targetting developers who is doing programming on windows, so i mostly do not focus on linux instructions in this README. Like with gcc, i focus on how to do it with windows. <br>

## Contributors
Lars Morten Bek (https://github.com/lmbek) <br>
Ida Marcher Jensen (gobek-example) (https://github.com/notHooman996) <br>
<br><br>
Feel free to open an issue or pull request with feature requests or if you have any issues<br>

## Requirements to developers
Go 1.20+ <br>
Chrome (Windows) or Chromium (Linux) <br>
macOS/DARWIN NOT SUPPORTED YET <br>
Knowing how to install go-sqlite from https://github.com/mattn/go-sqlite3

## Requirements for users
Chrome (Windows) or Chromium (Linux) <br>
macOS/DARWIN NOT SUPPORTED YET <br> 

## How to use (download example project)
The best way to start using the project is to download the example project at: <br>
https://github.com/lmbek/gobek-sqlite-example <br>
and then using:

	go get github.com/lmbek/gobek
This example project uses this package and combines it with a local api and database. <br>
Then the Go api is being developed and customized by you together with the frontend (JavaScript, HTML, CSS) <br>

## Installing gcc and sqlite3 third-party project required
When you want to use sqlite together with gobek, you need to preinstall a few things<br>
First install gcc, you can download it with one of the TDM-GCC Toolchains that can be found at https://jmeubank.github.io/tdm-gcc/ and find an installer (for advanced developers)<br>
Once you have installed it and set the Path in system environments, then you can try to confirm if it is installed by using gcc --version in CMD
This project already uses a dependency for the github project, but you can run it by using: <br>

	go get github.com/mattn/go-sqlite3
## How to test

	go test ./tests/...

## How to run
make your own main.go by following https://github.com/lmbek/gobek-example

## How to apply manifest and logo to executible
Use something like goversioninfo: https://github.com/josephspurrier/goversioninfo

## How to build

	go build -ldflags -H=windowsgui -o NewProjectName.exe
