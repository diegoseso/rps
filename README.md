# RPS gameserver
Rock, Paper, Scissors Game Server

## Running the docker 

 docker run -itd -p 80:80 -p 8080:8080 -v /home/diegoseso/go/src/github.com/diegoseso/sps-gameserver/server:/root/go/src/github.com/diegoseso/sps-gameserver -v /home/diegoseso/go/src/github.com/diegoseso/sps-gameserver/frontend:/var/www/html/sps diegoseso/sps

## Setting and running the frontend

To download the dependencies you only need to run the following command inside the docker:

``` bower install --allow-root ```

Once installed all the dependencies, to be able to test the frontend project you need to: 
``` /etc/init.d/nginx start ```

To run the server side, simply run: 

``` cd ~/go/src/github.com/diegoseso/sps; go run main.go```


