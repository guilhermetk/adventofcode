import { Expedition } from "./model/expedition";
import { INPUT, TEST_INPUT } from "./inputs";

const expedition = new Expedition();
expedition.processInput(INPUT);

let calories = 0;
expedition
  .getTopNElvesWithMostCalories(3)
  .forEach((elf) => (calories += elf.calculateCalories()));
console.log(calories);
