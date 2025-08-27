from locust import events
from locust.stats import stats_printer, stats_history
from locust.env import Environment
from locust_users.casual_browser import *

import gevent

def main():
    print("Locust GUI running on http://localhost:8089")
    print("Press Ctrl+C to stop")
    
    # setup Environment and Runner
    env = Environment(user_classes=[CasualBrowser])
    runner = env.create_local_runner()
    
    # start a WebUI instance
    web_ui = env.create_web_ui("127.0.0.1", 8089)

    # execute init event handlers
    env.events.init.fire(environment=env, runner=runner, web_ui=web_ui)

    # start a greenlet that periodically outputs the current stats
    gevent.spawn(stats_printer(env.stats))

    # start a greenlet that save current stats to history
    gevent.spawn(stats_history, runner)
    
    # start the test
    runner.start(10, spawn_rate=2)

    # wait for the greenlets
    runner.greenlet.join()

    # stop the web server for good measures
    web_ui.stop()

if __name__ == "__main__":
    main()