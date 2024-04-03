""" Server Module """
import atexit
import logging
import os
import socket
# import threading
from concurrent import futures
from concurrent.futures import ThreadPoolExecutor
from multiprocessing import Process, cpu_count

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
# from waitress import serve as waitress_serve

DEBUG = bool(os.environ.get('DEBUG', False))
setupLogger(name='root', debug=DEBUG)

log = logging.getLogger(__name__)
workerCount: int = cpu_count() // 4
browserCount: int = workerCount // 2


def serve(bindAddress: str) -> None:
  """ Server """
  options: tuple = (("grpc.so_reuseport", 1),)
  server: _Server = grpc.server(thread_pool=futures.ThreadPoolExecutor(max_workers=2),
                                options=options)
  itemsPBgrpc.add_ItemParserServicer_to_server(servicer=PMItemParserServicer(), server=server)
  categPBgrpc.add_CategoryParserServicer_to_server(servicer=PMCategoryParserServicer(),
                                                   server=server)
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
  # Flask app for health check
  # app = Flask(import_name=__name__)

  # @app.route(rule='/health', methods=['GET'])
  # def healthCheck() -> Response:
  #   return Response(response="OK", status=200)

  # # Run Flask app in a separate thread
  # threading.Thread(target=waitress_serve,
  #                  kwargs={
  #                    'app': app,
  #                    'host': addr[0],
  #                    'port': 9999
  #                  }).start()
  # log.info("Health check server started")
  for worker in workers:
    worker.join()


if __name__ == '__main__':
  main()
