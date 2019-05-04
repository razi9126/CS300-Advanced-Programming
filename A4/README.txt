I used sir's code and mostly random code from internet to implement the game. For checking the winning conditions I copied the
code for all 3 cases(horizontal, vertical and diagonal win) from the internet. I hope the puropose of the assignment was to 
have us learn React and not writing the winning conditions. At the time I'm writing this, I have partially implemented
the feature that any number of players in multiples of 2 can play(still need to convert some variables to arrays so that i 
can completely implement this feature).

What works.
1) Basic 7x6 grid that takes only valid moves and displays them to both users (Part1)
2) Users move will only be displayed at his turn
3) On winning one user gets winning message and the other the losing message (Part2)
4) If more than 2 players they get their own board(but the moves get mixed up in boards, working to fix that) (Part4 partial)

What doesn't work
1) Didn't have the boxes coloured green on winning move
2) My part4 works only partially, working to fix that.
3) My code mostly came from https://github.com/bryanbraun/connect-four/blob/gh-pages/js/functions.js for the conditions. The rest
   of the code was done through internet and mostly by myself( the client side is almost completely my own). Line 40-60 of 
   server.js are also copied from github. They implement the waiting functionality, which allows the game to start only after 
   2 players have joined



I couldn't find a method to extract value of a button thats why i made seven separate handlers for click on each button

