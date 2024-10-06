class screenObject {
  x;
  y;
  width;
  height;
  screen;
  ctx;
  projectileColor;
  fireCount;

  constructor(ctx, screen) {
    this.ctx = ctx;
    this.screen = screen;
    this.fireCount = 0;
  }

  getPos() {
    return { x: this.x, y: this.y };
  }

  setPos(coords) {
    if (coords.x && coords.x < this.screen.w && coords.x > 0) this.x = coords.x;
    if (coords.y && coords.y <= this.screen.h && coords.y > 0)
      this.y = coords.y;
  }

  getProjectileDirection() {
    if (this.projectileSpeed < 0) return -1;
    else return 1;
  }

  fire(projectilesList) {
    projectilesList.push(new Projectile(this, this.ctx, this.screen));
    this.fireCount++;
  }
}
