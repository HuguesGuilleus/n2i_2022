class snake{
    constructor(x,y){
        this.non_game_over=true;
        this.non_change=true;
        this.non_pause=true;
        this.sens="ArrowUp";
        this.taille=0
        this.score=0;
        this.vie=true;
        this.color="red";
        this.pos_x=parseInt(x);
        this.pos_y=parseInt(y);
        this.liste_pos=[[this.pos_x,this.pos_y]]
    }

    move(){
        if (this.vie){
            if (this.sens=="ArrowUp"){
                this.pos_y=this.pos_y-1;
                this.pos_x=this.pos_x;
            }else if (this.sens=="ArrowDown"){
                this.pos_y=this.pos_y+1;
                this.pos_x=this.pos_x;
            }else if (this.sens=="ArrowLeft"){
                this.pos_y=this.pos_y;
                this.pos_x=this.pos_x-1;
            }else if (this.sens=="ArrowRight"){
                this.pos_y=this.pos_y;
                this.pos_x=this.pos_x+1;
            }
        }
    }
    out(side){    
        if (side=="left"){
            this.pos_x=79;
        }else if (side=="right"){
            this.pos_x=1;
        }else if (side=="up"){
            this.pos_y=Math.floor(size_can/10)-1;
        }else if (side=="down"){
            this.pos_y=1;
        }

    }
}

class food{
    constructor(x,y){
        this.pos_x=parseInt(x)
        this.pos_y=parseInt(y)
    }
}



function my_event_key(sens) {
    // console.log(sens,typeof sens)
    if (my_snake.non_change){
        if (my_snake.non_pause){
            my_snake.non_change=false;
            if (sens=="ArrowUp" && my_snake.sens!="ArrowDown"){
                my_snake.sens="ArrowUp";
            } 
            if (sens=="ArrowDown" && my_snake.sens!="ArrowUp"){
                my_snake.sens="ArrowDown";
            } 
            if (sens=="ArrowLeft" && my_snake.sens!="ArrowRight"){
                my_snake.sens="ArrowLeft";
            } 
            if (sens=="ArrowRight" && my_snake.sens!="ArrowLeft"){
                my_snake.sens="ArrowRight";
            }   
            if (sens==" "){
                my_snake.non_pause=false;
            }            
        }
        else{
            if (sens=" "){
                my_snake.non_pause=true;
            }

        }

    }
}

function rand(min,max){
    return Math.floor(Math.random()*(max-min+1))+min;
}




var size_can=screen.height-300;



var my_snake= new snake(rand(5,Math.floor(size_can/10)-5),rand(5,Math.floor(size_can/10)-5))
var my_food= new food(rand(5,75),rand(5,Math.floor(size_can/10)-5))
var input, button, pause,over;
var start=false;

console.log("width",size_can,Math.floor(size_can/10),typeof size_can)
console.log(my_food,my_snake)
function setup(){
    var can=createCanvas(800,size_can);
    can.parent("div_game")
    console.log("width",size_can,Math.floor(size_can/10),typeof size_can)
    background('#000')
    console.log("can",can)
    frameRate(15)
}

function action(){
    document.getElementById("div_game").removeChild(document.getElementById('button_start'));
    start=true;
}

function canvas_display(msg) {
    document.getElementById('espace_msg').innerText=msg;
}

function mise_en_pause(argument) {
    if (my_snake.non_pause){
        my_snake.non_pause=false;
    }else{
        my_snake.non_pause=true;
    }
        
}

function draw(){
    
   
    
    if (start){ 
        document.getElementById('score').innerText="Score : "+my_snake.score;
        document.onkeydown=  function(e){e=e||window.event;my_event_key(e.key)} 
        background('#000');
        my_snake.non_change=true;
        if (my_snake.non_pause==false){
            canvas_display("jeu en pause");

        }else if (my_snake.non_game_over){
            canvas_display("")
            console.log("x",my_food.pos_x,"y",my_food.pos_y,my_food)
            // console.log("x",my_snake.pos_x,"y",my_snake.pos_y)
            fill(255,0,0);
            rect(my_snake.pos_x*10,my_snake.pos_y*10,10,10);
            fill(0,255,0);
            rect(my_food.pos_x*10,my_food.pos_y*10,10,10);
            my_snake.move();
            my_snake.liste_pos.push([my_snake.pos_x,my_snake.pos_y]);
            /*
            if (my_snake.liste_pos[my_snake.liste_pos.length-1][0]!=my_snake.pos_x || my_snake.liste_pos[my_snake.liste_pos.length-1][1]!=my_snake.pos_y ){
                my_snake.liste_pos.push([my_snake.pos_x,my_snake.pos_y]);
            }
            */
            
            if (my_snake.taille<my_snake.score){
                my_snake.taille+=1;
            }
            var reverse_liste_pos=my_snake.liste_pos.reverse();
            fill(0,0,255);
            for (var d=2;d<my_snake.taille+2;d++){                
                rect(my_snake.liste_pos[my_snake.liste_pos.length-d][0]*10,my_snake.liste_pos[my_snake.liste_pos.length-d][1]*10,10,10);
                if (my_snake.pos_x ==reverse_liste_pos[d+2][0] && my_snake.pos_y ==reverse_liste_pos[d+2][1]){
                    my_snake.non_game_over=false;
                }
            }
            if (my_snake.pos_x==my_food.pos_x && my_snake.pos_y==my_food.pos_y){
                my_snake.score+=2;
                my_food= new food(rand(5,75),rand(5,Math.floor(size_can/10)-5));
            }else if (my_snake.pos_x<=0){
                my_snake.out("left");
            }else if (my_snake.pos_x>=79){
                my_snake.out("right");
            }else if (my_snake.pos_y<=0){
                my_snake.out("up");
            }else if (my_snake.pos_y>=Math.floor(size_can/10)){
                my_snake.out("down");
            }
        }else{
            canvas_display("Game Over");
            start=false;

        } 
    }else{
        background('#83ffe7');
    }

}

