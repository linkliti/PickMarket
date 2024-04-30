import WhiteBlock from "@/components/base/WhiteBlock";
import CategoryMarketSelector from "@/components/categories/CategoryMarketSelector";
import CategorySelect from "@/components/categories/CategorySelect";
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { useCategoryStore } from "@/store/categoryStore";
import { Category, CategoryStore, Marketplace } from "@/types/categoryTypes";
import { TriangleAlert } from "lucide-react";
import { ReactElement } from "react";
import { NavigateFunction, useNavigate } from "react-router-dom";

export default function Categories(): ReactElement {
  const { toast } = useToast();
  const navigate: NavigateFunction = useNavigate();

  const [selectedMarket, selectedCategory] = useCategoryStore(
    (state: CategoryStore): [Marketplace, Category | null] => [
      state.selectedMarket,
      state.selectedCategory,
    ],
  );

  function redirectToFilters(): void {
    if (!selectedMarket || !selectedCategory) {
      toast({
        className: "p-4 border border-border",
        action: (
          <div className="flex w-full items-center">
            <TriangleAlert className="mr-2" />
            <span className="first-letter:capitalize">Не выбрана категория или маркетплейс</span>
          </div>
        ),
      });
      return;
    }
    const searchParams = new URLSearchParams();
    searchParams.set("market", selectedMarket.shortLabel);
    searchParams.set("category", selectedCategory.url);
    const targetURL: string = `/filters?${searchParams.toString()}`;
    navigate(targetURL);
  }

  return (
    <>
      <WhiteBlock className="finline-flex w-full justify-center">
        <h1 className="pb-4 text-2xl font-bold">Выберите маркетплейс:</h1>
        <div className="inline-flex flex-wrap">
          <CategoryMarketSelector
            className="me-4 pb-4"
            listClassNames="w-[200px] bg-white text-primary-foreground hover:bg-secondary"
          />
          <Button
            onClick={redirectToFilters}
            // disabled={!selectedCategory}
            className="border-border w-[150px] border sm:w-[200px]"
          >
            Поиск
          </Button>
        </div>
      </WhiteBlock>
      {selectedMarket && (
        <WhiteBlock className="w-full grow">
          <h1 className="pb-4 text-2xl font-bold">Выберите категорию:</h1>
          <CategorySelect marketplace={selectedMarket} />
        </WhiteBlock>
      )}
    </>
  );
}
