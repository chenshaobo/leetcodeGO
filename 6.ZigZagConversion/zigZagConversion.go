package zigZagConversion

import "bytes"

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

string convert(string text, int nRows);
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

 */

func Convert(s string, numRows int) string {
	length := len(s)
	if numRows == 1 || numRows > length{
		return s
	}

	rows := make([][]rune,numRows)
	row,step := 0,0
	for _,c := range s{
		rows[row] = append(rows[row],c)
		if row == 0{
			step =1
		}else if row == numRows -1{
			step = -1
		}
		row += step
	}
	var out bytes.Buffer
	for _, r := range rows {
		out.WriteString(string(r))
	}
	return out.String()
}