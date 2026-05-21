class Pistol {
  constructor(
    private model: string,
    private caliber: string,
    private magazineCapacity: number,
    private roundsInMagazine: number = 0
  ) {}

  static create(
    model: string,
    caliber: string,
    magazineCapacity: number
  ): Pistol {
    return new Pistol(model, caliber, magazineCapacity);
  }

  info(): string {
    return `Model: ${this.model}, Caliber: ${this.caliber}, Capacity: ${this.magazineCapacity}, Rounds: ${this.roundsInMagazine}`;
  }

  reload(rounds: number): void {
    const loaded = Math.max(0, Math.min(rounds, this.magazineCapacity));
    this.roundsInMagazine = loaded;
  }

  shoot(): boolean {
    if (this.roundsInMagazine === 0) {
      return false;
    }
    this.roundsInMagazine--;
    return true;
  }
}

const pistol = Pistol.create("Glock 17", "9mm", 17);

console.log(pistol.info());

pistol.reload(17);
console.log("After reload:", pistol.info());

for (let i = 1; i <= 3; i++) {
  if (pistol.shoot()) {
    console.log(`Shot ${i} fired. ${pistol.info()}`);
  }
}
