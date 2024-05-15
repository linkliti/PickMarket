""" Server Module """
import atexit
import logging
import os
import socket
# import threading
from concurrent import futures
from concurrent.futures import ThreadPoolExecutor
from multiprocessing import Process

import grpc
from app.grpcs.categoryParserServicer import PMCategoryParserServicer
from app.grpcs.itemParserServicer import PMItemParserServicer
from app.protos import categories_pb2_grpc as categPBgrpc
from app.protos import items_pb2_grpc as itemsPBgrpc
from app.selen.seleniumMarkets import SeleniumWorkerMarkets
from app.selen.seleniumPool import browserQueue
from app.utilities.log import setupLogger
# from flask import Flask, Response
from grpc._server import _Server
from grpc_health.v1 import health, health_pb2, health_pb2_grpc

# from waitress import serve as waitress_serve

DEBUG = bool(os.environ.get('DEBUG', False))
setupLogger(name='root', debug=DEBUG, filename='pmparser.log')

log = logging.getLogger(__name__)
workerCount: int = 4
browserCount: int = 4


def serve(bindAddress: str) -> None:
  """ Server """
  options: tuple = (("grpc.so_reuseport", 1),)
  server: _Server = grpc.server(thread_pool=futures.ThreadPoolExecutor(),
                                options=options)
  itemsPBgrpc.add_ItemParserServicer_to_server(servicer=PMItemParserServicer(), server=server)
  categPBgrpc.add_CategoryParserServicer_to_server(servicer=PMCategoryParserServicer(),
                                                   server=server)
  # Health server and workers
  configureHealthServer(server=server)
  server.add_insecure_port(address=bindAddress)
  server.start()
  log.info("Server started", extra={"bindAddress": bindAddress})
  server.wait_for_termination()


def reservePort():
  """Find and reserve a port for all subprocesses to use."""
  sock = socket.socket(family=socket.AF_INET, type=socket.SOCK_STREAM)
  sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEPORT, 1)
  if sock.getsockopt(socket.SOL_SOCKET, socket.SO_REUSEPORT) == 0:
    raise RuntimeError("Failed to set SO_REUSEPORT.")
  address: str = os.environ.get('PARSER_ADDR', default="localhost:1111")
  host, port = address.split(sep=":")
  sock.bind((host, int(port)))
  return sock.getsockname()


def configureHealthServer(server: grpc.Server) -> None:
  """ Configure health server """
  healthServicer = health.HealthServicer(
    experimental_non_blocking=True,
    experimental_thread_pool=futures.ThreadPoolExecutor(),
  )
  health_pb2_grpc.add_HealthServicer_to_server(healthServicer, server)
  healthServicer.set("parser", health_pb2.HealthCheckResponse.SERVING)


def setCookies(w: SeleniumWorkerMarkets) -> None:
  """ Call setOzonAdultCookies """
  w.setOzonAdultCookies()


def closeBrowsers() -> None:
  """ Close browsers """
  for _ in range(workerCount):
    worker: SeleniumWorkerMarkets = browserQueue.get()
    worker.driver.quit()


def main() -> None:
  """ Main """
  addr = reservePort()
  addrStr: str = addr[0] + ':' + str(addr[1])
  workers: list[Process] = []
  # Browsers
  log.info("Starting browsers", extra={'browserCount': browserCount})
  for _ in range(browserCount):
    worker = SeleniumWorkerMarkets()
    browserQueue.put(item=worker)
  with ThreadPoolExecutor() as executor:
    executor.map(setCookies, browserQueue.queue)
  log.info("Browsers started", extra={'browserCount': browserCount})
  atexit.register(closeBrowsers)
  # Workers
  log.info("Starting workers", extra={'workerCount': workerCount})
  if DEBUG:
    log.info("Debug enabled")
  for _ in range(workerCount):
    worker = Process(target=serve, args=(addrStr,))
    worker.start()
    workers.append(worker)
  for worker in workers:
    worker.join()


if __name__ == '__main__':
  main()
