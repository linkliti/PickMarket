from application import app
import os

if __name__ == '__main__':
  app.run(debug=os.environ.get('DEV', False),
          port=os.environ.get('PORT', 5000))
