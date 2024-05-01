import { Controller } from "../stimulus.js"

export default class extends Controller {
  connect() {
    if (this.element.dataset.isInProgress === "true") {
      this.startTimer()
    }
  }

  disconnect() {
    if (this.interval) {
      this.stopTimer()
    }
  }

  startTimer() {
    this.interval = setInterval(() => {
      const duration = this.element.innerHTML;
      const [_, hoursString, minutesString, secondsString] = /^(\d+):(\d+):(\d+)$/.exec(duration);
      let hours = parseInt(hoursString);
      let minutes = parseInt(minutesString);
      let seconds = parseInt(secondsString);
      const total_seconds = hours * 3600 + minutes * 60 + seconds + 1;
      hours = Math.floor(total_seconds / 3600);
      minutes = Math.floor((total_seconds - hours * 3600) / 60);
      seconds = total_seconds - hours * 3600 - minutes * 60;
      this.element.innerHTML = `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    }, 1000)
  }

  stopTimer() {
    clearInterval(this.interval)
  }
}
