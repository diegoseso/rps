var onlinePlayers;
var play;
var leaderBoard;

var PlayState = {
    preload: function(){
    },
    create: function(){
    board = game.add.sprite(game.width, game.height, 'board');
    board.anchor.set(1, 1);

    var boardScaleWidth = 600 / board.texture.frame.width;
    var boardScaleHeight = boardScaleWidth;
    board.scale.setTo(boardScaleWidth, boardScaleHeight);

    red01 = game.add.sprite(boardSpecial['red'][0][0], boardSpecial['red'][0][1], 'red');
    red01.anchor.set(0.5, 0.5);
    red01.scale.setTo(.6, .6);
    resetPlay();
    initGame();
    },

    update: function(){
        if (!gameEnds) {
            showDices();
        }
    }
}


function initGame() {
}
