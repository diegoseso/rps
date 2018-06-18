var isUserSet;
var username;
var httpLogin = "http://localhost/login";
var server = {address: "ws://localhost/ws"};

function isUserSet(){
    console.log('isUserSet function');
    if (Cookies.get('username') === undefined){
       return false;
    } else {
       currentUsername = Cookies.get('username')
       return true;
    }
}

function LoginScreen(){
    console.log('LoginScreen function');
    $('#login-form').show();
    $('#login-play').click(function(){
        username = $('#username').val();
        // Here we make an http request to validate against server
        LoginServer(username);
        $('#login-form').hide();
        Login(username);
    });
}

function LoginServer(username){
    console.log('Trying to log to the server');
    axios.get(httpLogin, {
        params: {
            username: username
        }
    })
        .then(function (response) {
            console.log(response);
            var bodyResponse = JSON.parse(response.data)
            if( bodyResponse.success = true ){
                username = bodyResponse.data;
                console.log(username);
            }
        })
        .catch(function (error) {
            console.log(error);
        });
}


function Login(username){
    if (!sconn) {
        return false;
    }
    if (sconn.readyState != 1){

    }
    var newLogin = {message:'login', username : username}
    sconn.send(JSON.stringify(newLogin));
    Listen();
    currentUsername = Cookies.set('username', username);
    $('#my-user').html(username);
    // Login process to the server
    $('#whole-game').show();
    $('#whole-game').css('display', 'inline-flex');
}


function Listen(){
    sconn.onmessage = function (event){
        console.log('Message received');
        msg = JSON.parse(event.data);
        console.log(msg)
        if(msg.type == 'onlinePLayers'){
            DisplaysUsersOnChat(msg.data);
        }
    };
}

function ServerConnect(){
    return new Promise(function(resolve, reject) {
        sconn = new WebSocket(server.address);
        sconn.onopen = function () {
            resolve(sconn)
        };
        sconn.onclose = function (evt) {
            reject();
        };
    });
}

function ServerUnavailable(){
    console.log('The server is unavailable');
}


function playIntroAudio(){
    welcomeLoop = self.game.add.audio('welcome');
    welcomeLoop.play();
}

var LoginState = {
    preload: function(){
        console.log('loginState preload');
        playIntroAudio();
    },
    create: function(){
        console.log('loginState create');
        ServerConnect().then(function() {
            //If there is no cookie with a username then ask for a name to login, otherwise use the same username in the cookie
            if (!isUserSet()){
                LoginScreen();
            }
            if (isUserSet()){
                Login(currentUsername);
            }
            game.state.start('PlayState');
        }).catch(function(err) {
            console.error(err)
        });
        console.log(sconn);
    }
}