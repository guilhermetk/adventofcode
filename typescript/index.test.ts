import { describe, expect, test } from "bun:test";
import { Game, Shape } from "./model/rps_game";
import { checkSectionOverlap } from "./day_4";

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

test("day_3 - task_1 - complete overlapping cleaning sections", () => {
  const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`;
  const pairs_split = input.split("\n");
  expect(checkSectionOverlap(pairs_split[0].split(","), true)).toBe(false);
  expect(checkSectionOverlap(pairs_split[1].split(","), true)).toBe(false);
  expect(checkSectionOverlap(pairs_split[2].split(","), true)).toBe(false);
  expect(checkSectionOverlap(pairs_split[3].split(","), true)).toBe(true);
  expect(checkSectionOverlap(pairs_split[4].split(","), true)).toBe(true);
  expect(checkSectionOverlap(pairs_split[5].split(","), true)).toBe(false);
});

test("day_3 - task_2 - some overlapping cleaning sections", () => {
  const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`;
  const pairs_split = input.split("\n");
  expect(checkSectionOverlap(pairs_split[0].split(","), false)).toBe(false);
  expect(checkSectionOverlap(pairs_split[1].split(","), false)).toBe(false);
  expect(checkSectionOverlap(pairs_split[2].split(","), false)).toBe(true);
  expect(checkSectionOverlap(pairs_split[3].split(","), false)).toBe(true);
  expect(checkSectionOverlap(pairs_split[4].split(","), false)).toBe(true);
  expect(checkSectionOverlap(pairs_split[5].split(","), false)).toBe(true);
});
