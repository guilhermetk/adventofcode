import { Elf } from "./elf";

export class Expedition {
  readonly elves: Elf[];
  constructor() {
    this.elves = [];
  }

  public getTopNElvesWithMostCalories(n: number): Elf[] {
    this.elves.sort((a, b) => b.calculateCalories() - a.calculateCalories());
    const topElves: Elf[] = [];
    for (let index = 0; index < n; index++) {
      topElves.push(this.elves[index]);
    }
    return topElves;
  }

  public processInput(input: string) {
    const expeditionSupply = input.split("\n\n");
    for (const elfSupply of expeditionSupply) {
      const elfObj = new Elf([]);
      elfSupply
        .split("\n")
        .forEach((item) => elfObj.addFood(Number.parseInt(item)));
      this.elves.push(elfObj);
    }
  }
}
