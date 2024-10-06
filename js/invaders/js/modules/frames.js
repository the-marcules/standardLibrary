let Frames = {
  startTime: 0,
  frames: 0,
  isRunning: false,
  startCounter() {
    this.reset();
    this.isRunning = true;
    this.startTime = new Date().getTime();
    return this;
  },
  count() {
    this.frames++;
  },
  getCurrentFPS() {
    if (!this.isRunning || this.isNull()) return -1;
    else return Math.round(this.frames / this.getElapsedTimeInSeconds());
  },
  getFPS() {
    if (this.isNull()) return -1;
    else return Math.round(this.frames / this.getElapsedTimeInSeconds());
  },
  isNull() {
    return this.startTime == 0 || this.frames == 0;
  },
  getElapsedTimeInSeconds() {
    if (this.isNull()) return -1;
    return (new Date().getTime() - this.startTime) / 1000;
  },
  trueEveryXseconds(seconds) {
    const fps = this.getCurrentFPS();
    if (fps == -1 || seconds == undefined || typeof seconds != "number") {
      return -1;
    }
    if (this.frames % (fps * seconds) == 0) {
      return true;
    }
  },
  stop() {
    this.isRunning = false;
    return this;
  },
  reset() {
    this.startTime = 0;
    this.frames = 0;
    this.isRunning = false;
  },
};
