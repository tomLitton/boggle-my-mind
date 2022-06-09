# boggle-my-mind
GoLang implementation of boggle

## Calling the API
This api works based on a simple list of letters, along with the rows and columns 
of the board.  Note:  Only a 4 X 4 board is supported at this time.

The api has the form: board/{rows}/{columns}/{letters}

For example:  `http://localhost:8080/board/4/4/naigorlydaisoiuf` which translates to a board of
n a i g
o r l y
d a i s
o i u f

This also does not yet have a minimum letter word.

## Testing
The unit tests usees genmock to generate mock implementations.

The generator must be run before the tests by:
```shell
go genereate .\..
```