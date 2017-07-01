class ComputerPlayer
    attr_reader :name, :mark

    def initialize(name, mark)
        @name = name
        @mark = mark
    end

    def get_move(board)
        puts "Computer is making a move..."
        minimax(board, @mark, 0)[:move]
    end

    def minimax(board, current_mark, depth)
        if board.over?
            score = Hash.new
            if board.winner == @mark
                score[:value] = 10 - depth
                return score
            else
                score[:value] = -10 + depth
                return score
            end
        end

        score_list = []
        board.available_moves.each do | move |
            new_board = board.get_copy
            new_board.place_mark(move, current_mark)

            if current_mark == :X
                score = minimax(new_board, :O, depth + 1)
                score[:move] = move
                score_list << score
            else
                score = minimax(new_board, :X, depth + 1)
                score[:move] = move
                score_list << score
            end
        end

        if current_mark == @mark
            score_list.max_by { |score| score[:value] }
        else
            score_list.min_by { |score| score[:value] }
        end
    end
end
