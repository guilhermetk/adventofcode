import { describe, expect, test } from "bun:test";
import { Game, Shape } from "./model/rps_game";

test("day_2 - task_1 - rock_paper_scissor", () => {
  const game = new Game();
  game.addRoundUsingFirstTaskStrategy(Shape.ROCK, Shape.PAPER);
  expect(game.getGameScore()).toBe(8);
  game.addRoundUsingFirstTaskStrategy(Shape.PAPER, Shape.ROCK);
  expect(game.getGameScore()).toBe(9);
  game.addRoundUsingFirstTaskStrategy(Shape.SCISSOR, Shape.SCISSOR);
  expect(game.getGameScore()).toBe(15);
});

test("day_2 - task_2 - rock_paper_scissor", () => {
  const game = new Game();
  game.addRoundUsingSecondTaskStrategy(Shape.ROCK, "Y");
  expect(game.getGameScore()).toBe(4);
  game.addRoundUsingSecondTaskStrategy(Shape.PAPER, "X");
  expect(game.getGameScore()).toBe(5);
  game.addRoundUsingSecondTaskStrategy(Shape.SCISSOR, "Z");
  expect(game.getGameScore()).toBe(12);
});
