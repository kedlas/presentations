#!/usr/bin/env python
from __future__ import with_statement, print_function

import os
from pyrabbit.api import Client

def get_queue_depths(host, username, password, vhost):
    """ Fetches queue depths from rabbitmq instance."""
    cl = Client(host, username, password)
    if not cl.is_alive():
        raise Exception("Failed to connect to rabbitmq")
    depths = {}
    queues = [q['name'] for q in cl.get_queues(vhost=vhost)]
    for queue in queues:
        if queue == "aliveness-test": #pyrabbit
            continue
        elif queue.startswith('amq.gen-'): #Anonymous queues
            continue
        depths[queue] = cl.get_queue_depth(vhost, queue)
    return depths

def send_stats(host, username, password, vhost, namespace):
    depths = get_queue_depths(host, username, password, vhost)
    for queue in depths:
        # Here send data about queue lenght somewhere
        print("namespace=%s queue=%s length=%i" % (namespace, queue, depths[queue]))

if __name__ == "__main__":
    send_stats(
        "localhost:5672",
        "guest",
        "guest",
        "/",
        "rabbitmq_depth")
