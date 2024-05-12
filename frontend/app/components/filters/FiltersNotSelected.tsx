import WhiteBlock from "@/components/base/WhiteBlock";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";

export default function FiltersNotSelected() {
  return (
    <>
      <WhiteBlock className="flex w-full flex-col">
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
