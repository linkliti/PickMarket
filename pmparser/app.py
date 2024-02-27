""" Server Module """
import logging
from app.parsers.ozon.ozonParserCategories import OzonParserCategories
from app.protos.items_pb2_grpc import ItemParserServicer

# from parsers.ozon.ozonParser import OzonParser

log = logging.getLogger(__name__)

# class ItemParserServicer(items_pb2_grpc.ItemParserServicer):
#   """ Implementation of proto """
#   def GetItems(self, request, context):
#     req: items_pb2.ItemsRequest = request
#     print(req.market)

class PMItemParserServicer(ItemParserServicer):
  def GetItems(self, request, context):
    pass


def serve() -> None:
  """ Server """
  # p = OzonParser()
  o = PMItemParserServicer()
  try:
    o.GetItemCharacteristics(None, None)
  except AttributeError:
    print("No such attribute")

  # print(p.api)
  p = OzonParserCategories()
  for i in p.getRootCategories():
    print(i)


#   server: _Server = grpc.server(thread_pool=futures.ThreadPoolExecutor(max_workers=10))
#   items_pb2_grpc.add_ItemParserServicer_to_server(servicer=ItemParserServicer(), server=server)
#   categories_pb2_grpc.add_CategoryParserServicer_to_server(
#     servicer=categories_pb2_grpc.CategoryParserServicer(), server=server)
#   server.add_insecure_port(address='[::]:50051')
#   server.start()
#   server.wait_for_termination()

if __name__ == '__main__':
  serve()
