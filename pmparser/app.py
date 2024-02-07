""" App entry point """
from app.server import server
from app.utilities.log import setupLogger

if __name__ == '__main__':
  DEBUG = True
  logger = setupLogger(name='root', debug=DEBUG)
  # app.run(debug=getEnv('DEV', False),
  #         port=getEnv('PORT', 5000))
  server()
