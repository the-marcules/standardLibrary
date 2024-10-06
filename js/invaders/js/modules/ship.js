class Ship extends screenObject {
  radius;
  guns;
  speed;
  projectileSpeed;

  constructor(ctx, screen) {
    super(ctx, screen);
    this.type = "ship";
    this.x = 150;
    this.y = 580;
    this.height = 20;
    this.radius = 20;
    this.guns = 1;
    this.speed = 20;
    this.projectileSpeed = -30;
    this.projectileColor = "cyan";
    this.enemysTerminated = 0;
  }

  reset() {
    this.x = this.screen.w / 2;
    this.y = 580;
    this.enemysTerminated = 0;
    this.fireCount = 0;
  }

  draw(ctx) {
    ctx.beginPath();
    ctx.fillStyle = "white";

    ctx.arc(this.x, this.y, this.radius, 0, 2 * Math.PI);
    ctx.fill();
    ctx.fillRect(this.x - 4, this.y + this.radius - 1, 8, 23);
    ctx.fillRect(this.x - this.radius, this.y + this.radius + 5, 6, 20);
    ctx.fillRect(this.x + this.radius - 6, this.y + this.radius + 5, 6, 20);
    ctx.fillRect(
      this.x - this.radius,
      this.y + this.radius + 15,
      2 * this.radius,
      6
    );
  }
  getBoundarys() {
    return {
      x1: this.x - this.radius,
      y1: this.y - this.radius,
      x2: this.x + this.radius,
      y2: this.y + this.radius,
    };
  }

  updateCoords(direction) {
    switch (direction) {
      case "Left":
        if (this.x - this.radius > 0) this.x -= this.speed;
        break;
      case "Right":
        if (this.x + this.radius < this.screen.w) this.x += this.speed;
        break;
        0;
    }
  }

  hasBeenHit() {}
}
