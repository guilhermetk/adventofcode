import { Food } from "./food";

export class Elf {
  constructor(private readonly items: Array<Food<number>>) {}

  public addFood(calories: number) {
    if (calories) {
      this.items.push(new Food(calories));
    }
  }

  public calculateCalories() {
    let amount = 0;
    for (const food of this.items) {
      amount += food.value;
    }
    return amount;
  }
}
