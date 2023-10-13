import { Game, Shape } from "./model/rps_game";

task1();
task2();

async function task1() {
  const task1InputFile = Bun.file("./input/day_2_task_1.txt");
  const fileContent = (await task1InputFile.text()).split("\n");
  const game = new Game();
  for (const line of fileContent) {
    const parameters = line.split(" ");
    const opponentShape = translateInputToOpponentShape(parameters[0]);
    const myShape = translateInputToMyShape(parameters[1]);
    game.addRoundUsingFirstTaskStrategy(opponentShape!, myShape!);
  }
  console.log("task 1 score: " + game.getGameScore());
}

async function task2() {
  const task1InputFile = Bun.file("./input/day_2_task_1.txt");
  const fileContent = (await task1InputFile.text()).split("\n");
  const game = new Game();
  for (const line of fileContent) {
    const parameters = line.split(" ");
    const opponentShape = translateInputToOpponentShape(parameters[0]);
    const myShape = parameters[1];
    game.addRoundUsingSecondTaskStrategy(opponentShape!, myShape!);
  }
  console.log("task 2 score: " + game.getGameScore());
}

function translateInputToOpponentShape(input: string) {
  switch (input) {
    case "A":
      return Shape.ROCK;
    case "B":
      return Shape.PAPER;
    case "C":
      return Shape.SCISSOR;
  }
}

function translateInputToMyShape(input: string) {
  switch (input) {
    case "Y":
      return Shape.PAPER;
    case "X":
      return Shape.ROCK;
    case "Z":
      return Shape.SCISSOR;
  }
}
