# Age-of-Empires-II-CLI
A CLI appllication that outputs infromation about units in Age of Empires II

## Steps to run in project folder:
* create a .env file with a single entry: `DATABASE_URI = <local_mongodb_ur>`
* or alternatively run `export DATABASE_URI = <local_mongodb_ur>`
* run `go build`
* run `./Age-of-Empires-II-CLI`
Note this only works on linux/uni for windows


## Steps to install application systemwide:
* run `go install`
* In a new terminal window, or the current if you want, run `export DATABASE_URI = <local_mongodb_ur>`
* run `Age-of-Empires-II-CLI`

Note: 
* Building steps work for linux and mac only.
* You can set the DATABASE_URI in ~/.bashrc but I am not sure if that's very secure :thinking:
