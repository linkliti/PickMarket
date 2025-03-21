import Logo from "@/components/base/Logo";
import WhiteBlock from "@/components/base/WhiteBlock";
import { Button } from "@/components/ui/button";
import { ReactElement, useEffect } from "react";
import { Link } from "react-router-dom";

export default function Hero(): ReactElement {
  useEffect((): void => {
    document.title = "PickMarket";
  }, []);

  return (
    <WhiteBlock className="flex w-full grow flex-col items-center justify-center">
      <div className="mx-auto max-w-xl text-center">
        <Logo className="text-4xl" />

        <p className="mt-4 text-lg">Удобная подборка товаров</p>

        <div className="mt-8 flex flex-wrap justify-center gap-4">
          <Button
            asChild
            className=""
          >
            <Link
              to="/categories"
              className="px-5 py-3"
            >
              Перейти к категориям
            </Link>
          </Button>
        </div>
      </div>
    </WhiteBlock>
  );
}
