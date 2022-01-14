export class Matrix {
  private matrix: string;

  constructor(matrix: string) {
    this.matrix = matrix;
  }

  get rows(): Array<Array<number>> {
    if (!this.matrix) throw new Error('matrix not defined');
    const rowsArr = this.matrix.split('\n');
    const res: Array<Array<number>> = [];

    rowsArr.forEach((row, i) => {
      res[i] = [];
      const rowItems = row.split(' ');
      rowItems.forEach((item) => {
        res[i].push(parseInt(item));
      });
    });

    return res;
  }

  get columns(): Array<Array<number>> {
    if (!this.matrix) throw new Error('matrix not defined');
    const res: Array<Array<number>> = [];

    this.rows.forEach((row) => {
      row.forEach((item, j) => {
        if (!Array.isArray(res[j])) res[j] = [];
        res[j].push(item);
      });
    });

    return res;
  }
}
