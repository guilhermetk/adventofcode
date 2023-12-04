export enum Shape {
  ROCK = 1,
  PAPER = 2,
  SCISSOR = 3,
}

enum Result {
  LOSE = 0,
  DRAW = 3,
  WIN = 6,
}

export class Game {
  score = 0;

  constructor() {}

  addRoundUsingFirstTaskStrategy(opponentShape: Shape, myShape: Shape) {
    let roundScore = myShape;
    if (opponentShape === myShape) {
      this.score += roundScore + Result.DRAW;
      return;
    }
    let roundResult;
    switch (opponentShape) {
      case Shape.ROCK:
        roundResult = myShape === Shape.PAPER ? Result.WIN : Result.LOSE;
        break;
      case Shape.PAPER:
        roundResult = myShape === Shape.SCISSOR ? Result.WIN : Result.LOSE;
        break;
      case Shape.SCISSOR:
        roundResult = myShape === Shape.ROCK ? Result.WIN : Result.LOSE;
        break;
    }
    this.score += roundScore + roundResult;
  }

  addRoundUsingSecondTaskStrategy(
    opponentShape: Shape,
    myMoveParapeter: string
  ) {
    let myShape;
    switch (myMoveParapeter) {
      case "Y":
        myShape = opponentShape;
        break;
      case "X":
        myShape = this.getShapeToLose(opponentShape);
        break;
      case "Z":
        myShape = this.getShapeToWin(opponentShape);
        break;
    }
    this.addRoundUsingFirstTaskStrategy(opponentShape, myShape!);
  }

  private getShapeToLose(opponentShape: Shape): Shape {
    if (opponentShape === Shape.ROCK) {
      return Shape.SCISSOR;
    } else if (opponentShape === Shape.PAPER) {
      return Shape.ROCK;
    } else {
      return Shape.PAPER;
    }
  }

  private getShapeToWin(opponentShape: Shape): Shape {
    if (opponentShape === Shape.ROCK) {
      return Shape.PAPER;
    } else if (opponentShape === Shape.PAPER) {
      return Shape.SCISSOR;
    } else {
      return Shape.ROCK;
    }
  }

  getGameScore() {
    return this.score;
  }
}
