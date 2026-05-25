impl Solution {
    pub fn num_rook_captures(board: Vec<Vec<char>>) -> i32 {
        let (rook_i, rook_j) = Self::find_rook_pos(&board);

        let mut result = 0;
        let mut i = rook_i - 1;
        while i < board.len() && board[i][rook_j] == '.' {
            i -= 1;
        }
        if i < board.len() && board[i][rook_j] == 'p' {
            result += 1;
        }

        i = rook_i + 1;
        while i < board.len() && board[i][rook_j] == '.' {
            i += 1;
        }
        if i < board.len() && board[i][rook_j] == 'p' {
            result += 1;
        }

        let mut j = rook_j - 1;
        while j < board[rook_i].len() && board[rook_i][j] == '.' {
            j -= 1;
        }
        if j < board[rook_i].len() && board[rook_i][j] == 'p' {
            result += 1;
        }

        j = rook_j + 1;
        while j < board[rook_i].len() && board[rook_i][j] == '.' {
            j += 1;
        }
        if j < board[rook_i].len() && board[rook_i][j] == 'p' {
            result += 1;
        }

        result
    }

    fn find_rook_pos(board: &Vec<Vec<char>>) -> (usize, usize) {
        let mut rook_i = 0;
        let mut rook_j = 0;
        
        for (i, row) in board.iter().enumerate() {
            for (j, c) in row.iter().enumerate() {
                if *c == 'R' {
                    return (i, j);
                }
            }
        }

        unreachable!()
    }
}
