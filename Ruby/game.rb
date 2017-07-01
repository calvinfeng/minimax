require_relative 'board'
require_relative 'human_player'
require_relative 'computer_player'

class Game
    attr_reader :board
    def initialize
        @p1 = HumanPlayer.new("Human", :X)
        @p2 = ComputerPlayer.new("Computer", :O)
        @current_player = @p1
        @board = Board.new
    end

    def current_player
        @current_player
    end

    def switch_players!
        if @current_player == @p1
            @current_player = @p2
        else
            @current_player = @p1
        end
    end

    def play_turn
        puts @board
        move = @current_player.get_move(@board)
        mark = @current_player.mark
        @board.place_mark(move, mark)
        switch_players!
    end

    def start
        puts "______________________________________"
        puts "Let's play Tic Tac Toe with Minimax AI"
        puts "______________________________________"
        until @board.over?
            play_turn
        end

        if @board.winner.nil?
            puts "Game over, it's a tied"
        else
            puts "Game over, and winner is #{@board.winner}"
        end
    end
end

game = Game.new
game.start
