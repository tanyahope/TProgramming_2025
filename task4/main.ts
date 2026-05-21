const xn = 0.11;
const xk = 0.36;
const step = 0.05;

function calcYnew(x: number): number {
  return (Math.sin(x) ** 3 + Math.cos(x) ** 3) * Math.log(x);
}

console.log("x\t\ty");
console.log("------------------------");

for (let x = xn; x <= xk + 1e-9; x += step) {
  const y = calcYnew(x);
  console.log(`${x.toFixed(2)}\t\t${y.toFixed(6)}`);
}
