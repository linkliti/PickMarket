import WhiteBlock from "@/components/base/WhiteBlock";
import { Button } from "@/components/ui/button";
import { ReactElement } from "react";
import { Link } from "react-router-dom";

export default function Hero(): ReactElement {
  return (
    <WhiteBlock className="flex w-full grow flex-col items-center justify-center">
      <div className="mx-auto max-w-xl text-center">
        <h1 className="text-4xl font-extrabold">PickMarket</h1>

        <p className="mt-4 text-lg">Удобная подборка товаров</p>

        <div className="mt-8 flex flex-wrap justify-center gap-4">
          <Button
            asChild
            className=""
          >
            <Link
              to="/categories"
              className="text-md block w-full rounded-lg px-5 py-3 font-medium sm:w-auto"
            >
              Перейти к категориям
            </Link>
          </Button>
        </div>
      </div>
    </WhiteBlock>
  );
}
