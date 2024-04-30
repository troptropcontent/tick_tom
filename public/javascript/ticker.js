
function startTimer(timer) {
  const timerInterval = setInterval(increaseTimer, 1000);

  function increaseTimer() {
    const duration = timer.innerHTML;
    const [_, hoursString, minutesString, secondsString] = /^(\d+):(\d+):(\d+)$/.exec(duration);
    let hours = parseInt(hoursString);
    let minutes = parseInt(minutesString);
    let seconds = parseInt(secondsString);
    const total_seconds = hours * 3600 + minutes * 60 + seconds + 1;
    hours = Math.floor(total_seconds / 3600);
    minutes = Math.floor((total_seconds - hours * 3600) / 60);
    seconds = total_seconds - hours * 3600 - minutes * 60;
    timer.innerHTML = `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
  }
}

timer = document.getElementById('timer');
if (timer && timer.dataset.isInProgress === 'true') {
  startTimer(timer);
}

