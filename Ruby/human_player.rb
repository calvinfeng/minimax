class HumanPlayer
    attr_reader :name, :mark

    def initialize(name, mark)
        @name = name
        @mark = mark
    end

    def get_move(board)
        puts "\nWhere would you like to place your mark? (row, col)"
        move = gets.chomp.split(",")
        puts 
        move.map {|el| el.to_i}
    end
end
