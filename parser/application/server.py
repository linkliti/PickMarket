from seleniumbase import Driver
from seleniumbase import undetected


def server():
  driver: undetected.Chrome = Driver(uc=True)
  try:
    driver.get('https://ozon.ru')
    driver.save_screenshot('nowsecure.png')
  finally:
    driver.quit()
