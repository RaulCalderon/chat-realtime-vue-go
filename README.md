Chat App with Go, Vuejs and Socket.io

1.- Install Go on your computer from the website: https://go.dev/

2.- Add it to your enviroment variables:

    - Windows + S
    - Search 'system enviroment variables'.
    - Select 'Enviroment variables'.
    - On system variables select 'Path' and click on 'Edit'.
    - Add a New line (if not exists) with the path of the .exe file ( My case: 'C:\Program Files\Go\bin')
    - Save changes and open a new terminal. (You may have to restart the whole editor too.)

3.- You can verify if Go was installed correctly by checking the version.

    - Run this on terminal: 'go version'
    - You should see something like this: 'go version go1.24.3 windows/amd64'

4.- Install Go dependencies from: https://github.com/feederco/go-socket.io

    - Run this on terminal: 'go get github.com/googollee/go-socket.io'

    NOTE: If you get an error about 'go get'. This is because some commands are deprecated on the latest version of Go. 
    SOLUTION: Run this on yor terminal: 'go mod init my-module'. This will create a module that Go needs to proceed. 
              Once installed, execute the same command: 'go get github.com/googollee/go-socket.io' and it will be installed.
    




