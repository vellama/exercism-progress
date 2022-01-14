interface BandColor {
  [key: string]: number;
}

const bandColors: BandColor = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9,
};

export function decodedValue(values: string[]) {
  return parseInt(`${bandColors[values[0]]}${bandColors[values[1]]}`);
}
