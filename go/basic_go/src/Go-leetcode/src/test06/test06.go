package main

func convert(s string, numRows int) string {
	n := len(s)
	var rows [][]byte
	dir := 1
	row := 0
	col := 0
	varNum := 0
	for i := 0; i < n; i++ {
		rows[row][col] = s[i]
		row++
		col++
		if row == numRows {

		}
	}
}

func main() {

}
