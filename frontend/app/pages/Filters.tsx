import WhiteBlock from "@/components/base/WhiteBlock";
import { Button } from "@/components/ui/button";
import { ReactElement } from "react";
import { Link, useSearchParams } from "react-router-dom";

export default function Filters(): ReactElement {
  const [searchParams] = useSearchParams();

  const market = searchParams.get("market");
  const categoryURL = searchParams.get("category");
  if (!market || !categoryURL) {
    return (
      <>
        <WhiteBlock className="w-full">
          <p className="pb-4">Не выбрана категория или маркетплейс</p>
          <Button
            asChild
            className="max-w-[200px]"
          >
            <Link
              to="/categories"
              className="text-md block w-full rounded-lg px-5 py-3 font-medium sm:w-auto"
            >
              Перейти к категориям
            </Link>
          </Button>
        </WhiteBlock>
      </>
    );
  }

  return (
    <WhiteBlock className="w-full">
      <>
        {market}
        {categoryURL}
      </>
    </WhiteBlock>
  );
}
