import WhiteBlock from "@/components/base/WhiteBlock";
import CategoryMarketSelector from "@/components/categories/CategoryMarketSelector";
import CategorySelect from "@/components/categories/CategorySelect";
import { Button } from "@/components/ui/button";
import { ReactElement } from "react";

export default function Categories(): ReactElement {
  // const [text, setText] = useState<string>("");

  // function handleSubmit(e: React.FormEvent<HTMLFormElement>): void {
  //   e.preventDefault();
  //   terminal.log(selectedMarketplace?.url, text);
  // }

  return (
    <>
      <WhiteBlock className="w-full">
        <h1 className="pb-4 text-2xl font-bold">Выберите маркетплейс:</h1>
        <CategoryMarketSelector
          className="pb-4"
          listClassNames="w-[200px] bg-white text-primary-foreground hover:bg-secondary"
        />
        <Button
          type="submit"
          className="border-border border"
        >
          Поиск
        </Button>
      </WhiteBlock>
      <WhiteBlock className="w-full grow">
        <h1 className="pb-4 text-2xl font-bold">Выберите категорию:</h1>
        <CategorySelect />
      </WhiteBlock>
    </>
  );
}
