class Projectile extends screenObject {
  origin; // entity which did fire
  radius;
  removeMe;

  constructor(origin, ctx, screen) {
    super(ctx, screen);
    this.origin = origin;
    this.removeMe = false;
    this.radius = 5;
    let startPosition = this.origin.getPos();
    startPosition.y +=
      this.getProjectileDirection() * (this.radius + this.origin.height);
    this.setPos(startPosition);
  }

  move() {
    if (this.y <= this.screen.h && this.y > 0)
      this.y += this.origin.projectileSpeed;
    else this.removeMe = true;
  }
  draw() {
    this.ctx.fillStyle = this.origin.projectileColor;
    this.ctx.beginPath();
    this.ctx.arc(this.x, this.y, 5, 0, 2 * Math.PI);
    this.ctx.fill();
  }

  checkCollosion(boundarys) {
    let Xn = Math.max(boundarys.x1, Math.min(this.x, boundarys.x2));
    let Yn = Math.max(boundarys.y1, Math.min(this.y, boundarys.y2));
    let Dx = Xn - this.x;
    let Dy = Yn - this.y;
    return Dx * Dx + Dy * Dy <= this.radius * this.radius;
  }
}
