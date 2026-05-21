const xValues = [0.2, 0.3, 0.38, 0.43, 0.57];

function calcY(x: number): number {
  return (Math.sin(x) ** 3 + Math.cos(x) ** 3) * Math.log(x);
}

console.log("x\t\ty");
console.log("------------------------");

xValues.forEach((x, i) => {
  const y = calcY(x);
  console.log(`x${i + 1} = ${x.toFixed(2)}\t${y.toFixed(6)}`);
});
