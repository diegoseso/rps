var configuration;
var mode;
var sconn;

mode = getMode();
configuration = getConfiguration(mode);

function getMode(){
    rawMode = JSON.parse('../mode.json');
    return rawMode.mode; 
}

function getConfiguration(mode){
    return JSON.parse('../config/' + mode + '.json')
}

var game = new Phaser.Game(600, 600, Phaser.AUTO, 'game');
game.state.add('BootState', BootState);
game.state.add('LoginState', LoginState);
game.state.add('PlayState', PlayState);
game.state.start('BootState');
