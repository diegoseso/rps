var BootState = {
    preload: function(){
        console.log('Enter Boot state preload');
        game.load.image('black_stone', 'img/stone.png');
        game.load.image('black_scissors', 'img/scissors.png');
        game.load.image('black_paper', 'img/paper.png');
        //this.load.audio('welcome', ['audio/welcome.ogg', 'audio/welcome.mp3']);

    },
    create: function(){
        this.state.start('LoginState');
    }
};