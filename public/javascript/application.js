// Stimulus
import { Application } from "./stimulus.js"

import TimerController from "./controllers/timer_controller.js"

window.Stimulus = Application.start()
Stimulus.register("timer", TimerController)