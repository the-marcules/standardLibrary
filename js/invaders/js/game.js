window.onload = function game() {
  const gameName = "Triangular Space Invaders from space";
  let enemys = [];
  let projectiles = [];
  let result = "Game over.";
  let enemyFireInterval;
  let stars = [];
  const enemyFireRate = 3000;
  const canvas = document.getElementById("game");
  const ctx = canvas.getContext("2d");
  const screen = {
    w: canvas.clientWidth,
    h: canvas.clientHeight,
  };
  const colors = ["orange", "green", "blue", "yellow", "red"];
  let colorsCurrentColorIndex = 0;
  const ship = new Ship(ctx, screen);
  let animationFrameId;

  let continueGame = true;

  let activeMode = "startScreen";

  document.addEventListener("keydown", (event) => {
    let code = event.code;
    switch (activeMode) {
      case "startScreen":
        ship.reset();
        result = "Game over.";
        continueGame = true;
        projectiles = [];
        enemys = [];
        let enemyCount = 3;
        for (let i = 0; i < enemyCount; i++) {
          enemys.push(new Enemy(ctx, screen, { x: i * 100, y: 50 }));
        }
        enemyFireInterval = setInterval(enemyFireHandler, enemyFireRate);
        animationFrameId = window.requestAnimationFrame(gameLoop);

        break;
      case "gameLoop":
        if (code.includes("Arrow"))
          ship.updateCoords(code.substring(5), screen.w);
        else if (code == "Escape") continueGame = false;
        else if (code == "Space") ship.fire(projectiles);
        break;
      case "scoreBoard":
        if (code == "Enter") {
          Frames.stop().reset();
          window.cancelAnimationFrame(animationFrameId);
          startScreen();
        }
        break;
    }
  });

  function drawGameName() {
    ctx.beginPath();
    ctx.font = "48px Arcade Interlaced";
    ctx.fillStyle = "orange";
    ctx.fillText("***", 380, 50);
    ctx.fillText("triangular ", 220, 100);
    ctx.fillText("invaders", 260, 150);
    ctx.fillText("***", 380, 200);
  }

  startScreen();

  function defineStarsPositions() {
    const starsCount = 25;
    for (let i = 0; i < starsCount; i++) {
      let x = Math.random() * 1000;
      let y = Math.random() * 1000;
      stars.push({ x, y });
    }
  }
  function gameBackGround() {
    stars.forEach((star) => {
      ctx.beginPath();
      ctx.fillStyle = "white";
      ctx.arc(star.x, star.y, 2, 0, 2 * Math.PI);
      ctx.fill();
    });
  }
  function startScreen() {
    defineStarsPositions();
    gameBackGround();
    activeMode = "startScreen";

    ctx.clearRect(0, 0, screen.w, screen.h);
    gameBackGround();
    drawGameName();
    ctx.fillStyle = "red";
    ctx.font = "32px Arcade Interlaced";
    ctx.fillText("press any key to start", 100, 400);
  }

  function enemyHandler() {
    enemys.forEach((enemy) => {
      enemy.move();
      enemy.draw();
    });
  }
  function enemyFireHandler() {
    enemys.forEach((enemy) => {
      enemy.fire(projectiles);
    });
  }

  function projectileHandler() {
    let removeKeys = [];
    projectiles.forEach((projectile, i) => {
      if (projectile.removeMe) removeKeys.push(i);
      projectile.move();
      projectile.draw();
      enemys.forEach((enemy, i) => {
        if (projectile.checkCollosion(enemy.getBoundarys())) {
          enemy.destroy();
          enemys.splice(i, 1);
          ship.enemysTerminated++;
        }
      });

      if (projectile.origin.type != "ship")
        if (projectile.checkCollosion(ship.getBoundarys())) {
          continueGame = false;
          result = "Game over.";
        }
    });
    removeKeys.forEach((key) => projectiles.splice(key, 1));
  }

  function gameLoop() {
    activeMode = "gameLoop";
    ctx.clearRect(0, 0, screen.w, screen.h);
    gameBackGround();

    ship.draw(ctx);
    enemyHandler();
    projectileHandler();
    if (enemys.length < 1) {
      continueGame = false;
      result = "You won!";
    }
    if (continueGame) window.requestAnimationFrame(gameLoop);
    else scroreBoard();
  }

  function switchColors() {
    if (colorsCurrentColorIndex >= colors.length) colorsCurrentColorIndex = 0;
    if (Frames.trueEveryXseconds(0.1)) colorsCurrentColorIndex++;
    return colors[colorsCurrentColorIndex];
  }

  function scroreBoard() {
    window.cancelAnimationFrame(animationFrameId);
    if (Frames.isRunning) Frames.count();
    else Frames.startCounter().count();

    clearInterval(enemyFireInterval);
    ctx.fillStyle = "white";
    activeMode = "scoreBoard";

    ctx.clearRect(0, 0, screen.w, screen.h);
    gameBackGround();
    drawGameName();
    ctx.fillStyle = switchColors();
    ctx.font = "32px Arcade Interlaced";
    ctx.fillText(result, 300, 300);

    ctx.fillStyle = "white";
    ctx.font = "18px Arcade Interlaced";
    ctx.fillText("You fired " + ship.fireCount + " times.", 50, 350);
    ctx.fillText(
      "You Terminated " + ship.enemysTerminated + " enemys.",
      50,
      400
    );

    ctx.font = "18px Arcade Interlaced";
    ctx.fillText("Press 'Enter' to restart the game.", 50, 550);

    animationFrameId = window.requestAnimationFrame(scroreBoard);
  }
};
