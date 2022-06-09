package board_test

import (
    "github.com/tomlitton/boggle-my-mind/pkg/word"
)

var PATHS2X2 = []word.Path {
    word.Path{ 0 },
    word.Path{ 1 },
    word.Path{ 2 },
    word.Path{ 3 },

    word.Path{ 0, 1 },
    word.Path{ 0, 2 },
    word.Path{ 0, 3 },
    word.Path{ 1, 3 },
    word.Path{ 1, 2 },
    word.Path{ 1, 0 },
    word.Path{ 2, 1 },
    word.Path{ 2, 3 },
    word.Path{ 2, 0 },
    word.Path{ 3, 0 },
    word.Path{ 3, 1 },
    word.Path{ 3, 2 },

    word.Path{ 0, 1, 2 },
    word.Path{ 0, 1, 3 },
    word.Path{ 0, 2, 3 },
    word.Path{ 0, 2, 1 },
    word.Path{ 0, 3, 1 },
    word.Path{ 0, 3, 2 },

    word.Path{ 1, 0, 2 },
    word.Path{ 1, 0, 3 },
    word.Path{ 1, 2, 0 },
    word.Path{ 1, 2, 3 },
    word.Path{ 1, 3, 0 },
    word.Path{ 1, 3, 2 },

    word.Path{ 2, 0, 1 },
    word.Path{ 2, 0, 3 },
    word.Path{ 2, 1, 0 },
    word.Path{ 2, 1, 3 },
    word.Path{ 2, 3, 0 },
    word.Path{ 2, 3, 1 },

    word.Path{ 3, 0, 1 },
    word.Path{ 3, 0, 2 },
    word.Path{ 3, 1, 0 },
    word.Path{ 3, 1, 2 },
    word.Path{ 3, 2, 0 },
    word.Path{ 3, 2, 1 },

    word.Path{ 0, 1, 2, 3 },
    word.Path{ 0, 1, 3, 2 },
    word.Path{ 0, 2, 1, 3 },
    word.Path{ 0, 2, 3, 1 },
    word.Path{ 0, 3, 1, 2 },
    word.Path{ 0, 3, 2, 1 },

    word.Path{ 1, 0, 2, 3 },
    word.Path{ 1, 0, 3, 2 },
    word.Path{ 1, 2, 0, 3 },
    word.Path{ 1, 2, 3, 0 },
    word.Path{ 1, 3, 0, 2 },
    word.Path{ 1, 3, 2, 0 },

    word.Path{ 2, 0, 1, 3 },
    word.Path{ 2, 0, 3, 1 },
    word.Path{ 2, 1, 0, 3 },
    word.Path{ 2, 1, 3, 0 },
    word.Path{ 2, 3, 1, 0 },
    word.Path{ 2, 3, 0, 1 },

    word.Path{ 3, 0, 1, 2 },
    word.Path{ 3, 0, 2, 1 },
    word.Path{ 3, 1, 0, 2 },
    word.Path{ 3, 1, 2, 0 },
    word.Path{ 3, 2, 0, 1 },
    word.Path{ 3, 2, 1, 0 },
}

// This is a few valid cases.  There are too many to manually role
// @GTODO:  Need a few more cases
var PATH4X4_VALID = []word.Path {
    word.Path{0,1,2,3},
    word.Path{0,5,10,15},
    word.Path{0,1,2,3,7,11,15,14,10,6,5,9,13,12,8,4},
    word.Path{0,1,2,3,7,11,15,14,13,12,8,4,5,9,10,6},
}

// This is a few valid cases.  There are too many to manually role
// @GTODO:  Need a few more cases
var PATH4X4_INVALID = []word.Path {
    word.Path{0,2,1,3},
    word.Path{0,8,5,12},
    word.Path{0,1,2,3,7,11,15,14,10,6,5,9,0,12,8,4},
}