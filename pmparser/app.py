""" Server Module """
import logging
import os
from concurrent import futures

import grpc
from app.grpcs.categoryParserServicer import PMCategoryParserServicer
from app.grpcs.itemParserServicer import PMItemParserServicer
from app.protos import categories_pb2_grpc as categPBgrpc
from app.protos import items_pb2_grpc as itemsPBgrpc
from app.utilities.log import setupLogger

DEBUG = bool(os.environ.get('DEBUG', False))
log: logging.Logger = setupLogger(name='root', debug=DEBUG)


def serve() -> None:
  """ Server """
  server = grpc.server(thread_pool=futures.ThreadPoolExecutor(max_workers=10))
  itemsPBgrpc.add_ItemParserServicer_to_server(servicer=PMItemParserServicer(), server=server)
  categPBgrpc.add_CategoryParserServicer_to_server(servicer=PMCategoryParserServicer(),
                                                   server=server)
  address = "localhost:50051"
  server.add_insecure_port(address=address)
  log.info("Server started, listening on %s", address)
  if DEBUG:
    log.info("Debug enabled")
  server.start()
  server.wait_for_termination()


if __name__ == '__main__':
  serve()
