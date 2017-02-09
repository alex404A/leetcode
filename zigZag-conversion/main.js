/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
  if (numRows === 1) {
    return s;
  }

  var i, j;
  var row = 0, col =0;
  var cnt = 0;
  var dirFlag = true;
  var map = []

  for (var i = 0; i < numRows; i++) {
    map[i] = [];
  }

  for (var i = 0; i < s.length; i++) {
    if (dirFlag && cnt <= numRows) {
      map[row][col] = s[i];
      row++;
      cnt++;
    } else {
      console.log(row, col, i);
      map[row][col] = s[i];
      row--;
      col++;
      cnt++;
    }

    if (dirFlag && cnt === numRows) {
      if (numRows === 2) {
        row = 0;
        cnt = 0;
        col++;
      } else {
        row -= 2;
        col++;
        cnt = 0;
        dirFlag = false;
      }
    } else if (!dirFlag && cnt === numRows - 2) {
      row = 0;
      cnt = 0;
      dirFlag = true;
    }
  }

  var result = '';

  for (i = 0; i < numRows; i++) {
    for (j = 0; j < map[i].length; j++) {
      if (map[i][j]) {
        result += map[i][j];
      }
    }
  }

  return result;

};

console.log(convert('abcd', 2));
