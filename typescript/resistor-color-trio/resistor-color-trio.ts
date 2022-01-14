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

const humanize = (value: number): string => {
  if (value > 1000) return `${Math.floor(value / 1000)} kiloohms`;
  return `${value} ohms`;
};

export function decodedResistorValue(values: string[]) {
  const exponent = bandColors[values[2]];
  const value =
    parseInt(`${bandColors[values[0]]}${bandColors[values[1]]}`) *
    Math.pow(10, exponent);

  return humanize(value);
}
