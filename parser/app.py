from application import app
from application.utils import getEnv
from application.server import server
import os

if __name__ == '__main__':
  # app.run(debug=getEnv('DEV', False),
  #         port=getEnv('PORT', 5000))
  server()
