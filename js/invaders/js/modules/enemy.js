class Enemy extends screenObject {
  #shots;
  #guns;
  #direction;
  #speed;
  projectileSpeed;

  constructor(ctx, screen, initialOffset) {
    super(ctx, screen);
    this.x = 0 + initialOffset.x;
    this.y = 0 + initialOffset.y;
    this.height = 30;
    this.width = 40;
    this.#shots = 0;
    this.#guns = 1;
    this.projectileColor = "red";
    this.projectileSpeed = 6;
    this.#speed = -4;
    this.#direction = "left";
  }

  draw() {
    this.ctx.fillStyle = "red";
    this.ctx.beginPath();
    this.ctx.moveTo(this.x, this.y);
    this.ctx.lineTo(this.x + this.width, this.y);
    this.ctx.lineTo(this.x + this.width / 2, this.y + this.height);
    this.ctx.lineTo(this.x, this.y);
    this.ctx.fill();
  }

  getBoundarys() {
    return {
      x1: this.x,
      y1: this.y,
      x2: this.x + this.width,
      y2: this.y + this.height,
    };
  }

  move() {
    this.setSpeed();
    if (this.needToSwap()) this.swapDirection();

    this.setPos({
      x: this.x + this.#speed,
      y: this.y,
    });
  }

  swapDirection() {
    if (this.#direction == "left") this.#direction = "right";
    else this.#direction = "left";
  }

  needToSwap() {
    switch (this.#direction) {
      case "left":
        if (this.x <= Math.abs(this.#speed)) return true;
        break;
      case "right":
        if (this.x + this.width >= this.screen.w) return true;
        break;
    }
    return false;
  }

  setSpeed() {
    switch (this.#direction) {
      case "left":
        this.#speed = -1 * Math.abs(this.#speed);
        break;
      case "right":
        this.#speed = Math.abs(this.#speed);
    }
  }

  destroy() {}

  upgrade() {
    this.#guns++;
  }
}
