""" Server Module """
import logging
import os
from concurrent import futures
from concurrent.futures import ThreadPoolExecutor

import grpc
from app.grpcs.categoryParserServicer import PMCategoryParserServicer
from app.grpcs.itemParserServicer import PMItemParserServicer
from app.protos import categories_pb2_grpc as categPBgrpc
from app.protos import items_pb2_grpc as itemsPBgrpc
from app.selen.seleniumMarkets import SeleniumWorkerMarkets
from app.selen.seleniumPool import browserQueue
from app.utilities.log import setupLogger

DEBUG = bool(os.environ.get('DEBUG', False))
log: logging.Logger = setupLogger(name='root', debug=DEBUG)


def serve() -> None:
  """ Server """
  server = grpc.server(thread_pool=futures.ThreadPoolExecutor())
  itemsPBgrpc.add_ItemParserServicer_to_server(servicer=PMItemParserServicer(), server=server)
  categPBgrpc.add_CategoryParserServicer_to_server(servicer=PMCategoryParserServicer(),
                                                   server=server)
  address = "localhost:50051"
  server.add_insecure_port(address=address)
  if DEBUG:
    log.info("Debug enabled")

  workerCount = 1
  for _ in range(workerCount):
    worker = SeleniumWorkerMarkets()
    browserQueue.put(item=worker)
  with ThreadPoolExecutor() as executor:
    executor.map(setCookies, browserQueue.queue)
  log.info("Browsers started. Worker count: %d", workerCount)

  server.start()
  log.info("Server started, listening on %s", address)
  server.wait_for_termination()


def setCookies(worker: SeleniumWorkerMarkets) -> None:
  """ Call setOzonAdultCookies """
  worker.setOzonAdultCookies()


if __name__ == '__main__':
  serve()
